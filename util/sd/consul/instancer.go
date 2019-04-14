package consul

import (
	"errors"
	"fmt"
	"github.com/LongMarch7/higo/util/pool"
	"strconv"
	"strings"
	"time"
	consul "github.com/hashicorp/consul/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/LongMarch7/higo/util/sd/internal/instance"
	"github.com/go-kit/kit/util/conn"
)

const defaultIndex = 0

// errStopped notifies the loop to quit. aka stopped via quitc
var errStopped = errors.New("quit and closed consul instancer")

// Instancer yields instances for a service in Consul.
type Instancer struct {
	cache       *instance.Cache
	client      Client
	logger log.Logger
	service     string
	tags        []string
	passingOnly bool
	quitc       chan struct{}
	updatePool   pool.UpdatePool
}

// NewInstancer returns a Consul instancer that publishes instances for the
// requested service. It only returns instances for which all of the passed tags
// are present.
func NewInstancer(client Client, logger log.Logger, service string, tags []string, passingOnly bool, pool  pool.UpdatePool) *Instancer {
	s := &Instancer{
		cache:       instance.NewCache(),
		client:      client,
		logger: logger,
		service:     service,
		tags:        tags,
		passingOnly: passingOnly,
		quitc:       make(chan struct{}),
		updatePool: pool,
	}

	instances, index, err := s.getInstances(defaultIndex, nil)
	if err == nil {
		s.logger.Log("NewInstancer--", len(instances))
	} else {
		s.logger.Log("NewInstancer--", len(instances))
	}

	s.cache.Update(sd.Event{Instances: instances, Err: err})
	go s.loop(index)
	return s
}

// Stop terminates the instancer.
func (s *Instancer) Stop() {
	close(s.quitc)
}

func (s *Instancer) loop(lastIndex uint64) {
	var (
		instances []string
		err       error
		d         time.Duration = 10 * time.Millisecond
	)
	for {
		instances, lastIndex, err = s.getInstances(lastIndex, s.quitc)
		switch {
		case err == errStopped:
			return // stopped via quitc
		case err != nil:
			s.logger.Log("NewInstancer--loop--", err)
			time.Sleep(d)
			d = conn.Exponential(d)
			s.cache.Update(sd.Event{Err: err})
		default:
			s.cache.Update(sd.Event{Instances: instances})
			d = 10 * time.Millisecond
		}
	}
}

func (s *Instancer) getInstances(lastIndex uint64, interruptc chan struct{}) ([]string, uint64, error) {
	tag := ""
	if len(s.tags) > 0 {
		tag = s.tags[0]
	}

	// Consul doesn't support more than one tag in its service query method.
	// https://github.com/hashicorp/consul/issues/294
	// Hashi suggest prepared queries, but they don't support blocking.
	// https://www.consul.io/docs/agent/http/query.html#execute
	// If we want blocking for efficiency, we must filter tags manually.

	type response struct {
		instances []string
		index     uint64
	}

	var (
		errc = make(chan error, 1)
		resc = make(chan response, 1)
	)

	go func() {
		entries, meta, err := s.client.Service(s.service, tag, s.passingOnly, &consul.QueryOptions{
			WaitIndex: lastIndex,
		})
		if err != nil {
			errc <- err
			return
		}
		if len(s.tags) > 1 {
			entries = filterEntries(entries, s.tags[1:]...)
		}
		instances, count := makeInstances(entries)
		s.updatePool( s.service, instances, uint32(count))
		resc <- response{
			instances: instances,
			index:     meta.LastIndex,
		}
	}()

	select {
	case err := <-errc:
		return nil, 0, err
	case res := <-resc:
		return res.instances, res.index, nil
	case <-interruptc:
		return nil, 0, errStopped
	}
}

// Register implements Instancer.
func (s *Instancer) Register(ch chan<- sd.Event) {
	s.cache.Register(ch)
}

// Deregister implements Instancer.
func (s *Instancer) Deregister(ch chan<- sd.Event) {
	s.cache.Deregister(ch)
}

func filterEntries(entries []*consul.ServiceEntry, tags ...string) []*consul.ServiceEntry {
	var es []*consul.ServiceEntry

ENTRIES:
	for _, entry := range entries {
		ts := make(map[string]struct{}, len(entry.Service.Tags))
		for _, tag := range entry.Service.Tags {
			ts[tag] = struct{}{}
		}

		for _, tag := range tags {
			if _, ok := ts[tag]; !ok {
				continue ENTRIES
			}
		}
		es = append(es, entry)
	}

	return es
}

func makeInstances(entries []*consul.ServiceEntry) ([]string, int) {
	instances := make([]string, len(entries))
	var count = 64
	for i, entry := range entries {
		addr := entry.Node.Address
		if entry.Service.Address != "" {
			addr = entry.Service.Address
		}
		instances[i] = fmt.Sprintf("%s:%d", addr, entry.Service.Port)
		countTag := entry.Service.Tags[2]
		countTag = strings.Replace(countTag, "maxThreadCount=", "", -1)
		count,_ = strconv.Atoi(countTag)
	}
	return instances, count
}
package etcdv3

import (
	"github.com/go-kit/kit/sd"
	"github.com/LongMarch7/higo/util/sd/internal/instance"
	"github.com/LongMarch7/higo/util/pool"
	"google.golang.org/grpc/grpclog"
	"strconv"
)

// Instancer yields instances stored in a certain etcd keyspace. Any kind of
// change in that keyspace is watched and will update the Instancer's Instancers.
type Instancer struct {
	cache  *instance.Cache
	client Client
	prefix string
	quitc  chan struct{}
	updatePool   pool.UpdatePool
}



// NewInstancer returns an etcd instancer. It will start watching the given
// prefix for changes, and update the subscribers.
func NewInstancer(c Client, prefix string, pool  pool.UpdatePool) (*Instancer, error) {
	s := &Instancer{
		client: c,
		prefix: prefix,
		cache:  instance.NewCache(),
		quitc:  make(chan struct{}),
		updatePool: pool,
	}

	instances, err := s.client.GetEntries(s.prefix)
	if err == nil {
		grpclog.Info("prefix", s.prefix, "instances", len(instances))
	} else {
		grpclog.Error("prefix", s.prefix, "err", err)
	}
	s.cache.Update(sd.Event{Instances: instances, Err: err})

	go s.loop()
	return s, nil
}

func (s *Instancer) loop() {
	ch := make(chan struct{})
	go s.client.WatchPrefix(s.prefix, ch)

	for {
		select {
		case <-ch:
			instances, err := s.client.GetEntries(s.prefix)
			if err != nil {
				grpclog.Error("msg", "failed to retrieve entries", "err", err)
				s.cache.Update(sd.Event{Err: err})
				continue
			}
			count, _ :=  strconv.Atoi(string(s.client.GetKV(s.prefix + "thread")))
			s.updatePool( s.prefix, instances, uint32(count) )
			s.cache.Update( sd.Event{Instances: instances} )

		case <-s.quitc:
			return
		}
	}
}

// Stop terminates the Instancer.
func (s *Instancer) Stop() {
	close(s.quitc)
}

// Register implements Instancer.
func (s *Instancer) Register(ch chan<- sd.Event) {
	s.cache.Register(ch)
}

// Deregister implements Instancer.
func (s *Instancer) Deregister(ch chan<- sd.Event) {
	s.cache.Deregister(ch)
}

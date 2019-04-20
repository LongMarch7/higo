package pool

import (
    sq "github.com/LongMarch7/higo/util/queue"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "sync"
    "time"
)


type Pool struct {
    Queue                *sq.EsQueue
    prefix               string
    InvalidateDeadline   time.Time
}

type ConnectManager struct {
    Conn                 *grpc.ClientConn
    InvalidateDeadline   time.Time
    Timeout              time.Duration
}

var poolManager = make( map[string] *Pool, 128 )

type UpdatePool func(string, []string, uint32)

var InvalidateTimeout = time.Minute
var lock sync.Mutex
var timer *time.Timer

func Init(){
    timer = time.NewTimer(InvalidateTimeout)
    go func(t *time.Timer) {
        for {
            <-t.C
            grpclog.Info("update micro service")
            for _,value :=range poolManager{
                if time.Now().After(value.InvalidateDeadline){
                    for value.Queue.Quantity() != 0 {
                        val, ok, _ := value.Queue.Get()
                        if !ok {
                            val.(*ConnectManager).Conn.Close()
                            val.(*ConnectManager).Conn = nil
                            val = nil
                        }
                    }
                }
            }
            t.Stop()
        }
    }(timer)
}


func Update( prefix string, instances []string, count uint32){
    go func(){
        lock.Lock()
        for _, v := range instances {
            if pool, ok := poolManager[v]; ok{
                if pool.prefix != prefix{
                    for pool.Queue.Quantity() != 0 {
                        val, ok, _ := pool.Queue.Get()
                        if !ok {
                            val.(*grpc.ClientConn).Close()
                            val = nil
                        }
                    }
                    pool.prefix = prefix
                }
            }else{
                if count <= 64 {
                    count = 64
                }
                pool = &Pool{
                    Queue: sq.NewQueue(count),
                    prefix: prefix,
                }
                poolManager[v] = pool
            }
            poolManager[v].InvalidateDeadline = time.Now().Add(InvalidateTimeout*2)
            grpclog.Info( "[prefix]=", prefix,",[instance]=",v)
        }
        timer.Reset(InvalidateTimeout)
        lock.Unlock()
    }()
}

func GetConnect(key string) (*Pool, bool){
    pool, ok := poolManager[key]
    return pool, ok
}

var ConnectInvalidateTimeout = time.Minute * 3

func GetConnectFromPool(addr string, p *Pool, opt ...grpc.DialOption) (*ConnectManager, error){
    cManager, ok, _ := p.Queue.Get()
    if !ok {
        time.Sleep(time.Microsecond * 100)
        cManager, ok, _ =  p.Queue.Get()
    }
    if !ok{
        conn, err := grpc.Dial(addr, opt...)
        if err == nil {
            cManager = new(ConnectManager)
            cManager.(*ConnectManager).Conn = conn
            cManager.(*ConnectManager).InvalidateDeadline = time.Now().Add(ConnectInvalidateTimeout)
        }
        return cManager.(*ConnectManager), err
    }
    return cManager.(*ConnectManager), nil
}

func PutConnectToPool(manager *ConnectManager, p *Pool) {
    var ok = true

    defer func(){
        if !ok {
            manager.Conn.Close()
            manager.Conn = nil
            manager = nil
        }
    }()
    if time.Now().After(manager.InvalidateDeadline) {
        ok = false
        return
    }
    manager.InvalidateDeadline = time.Now().Add(ConnectInvalidateTimeout)
    ok, _ = p.Queue.Put(manager)
    if !ok {
        time.Sleep(time.Microsecond)
        ok, _ =  p.Queue.Put(manager)
    }

}

func Destroy(){
    lock.Lock()
    for _,value :=range poolManager{
        for value.Queue.Quantity() != 0 {
           val, ok, _ := value.Queue.Get()
           if !ok {
               val.(*ConnectManager).Conn.Close()
               val.(*ConnectManager).Conn = nil
               val = nil
           }
        }
        value.Queue = nil
        value = nil
    }
    poolManager = nil
    lock.Unlock()
}
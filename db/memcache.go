package db

import (
    "github.com/bradfitz/gomemcache/memcache"
    "sync"
)


var initOnce sync.Once
var MemcacheClient *memcache.Client
const DefaultExpiration = 60*60*12

func MemcacheInit(conn int , server []string){
    initOnce.Do(func() {
        mem :=memcache.New(server...)
        mem.MaxIdleConns = conn
        MemcacheClient = mem
    })
}

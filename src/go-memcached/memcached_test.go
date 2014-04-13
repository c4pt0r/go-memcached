package gomemcached

import (
    "testing"
    "github.com/bradfitz/gomemcache/memcache"
)

func Test_Server(t *testing.T) {
    mc := memcache.New("127.0.0.1:8088")
    mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
}

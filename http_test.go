package gomemcached

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHTTPPool_ServeHTTP(t *testing.T) {
	var dbs = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	gee := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := dbs[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "127.0.0.1:9999"
	peers := NewHTTPPool(addr)
	log.Println("gomemcached is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
	if view, err := gee.Get("unknown"); err != nil {
		fmt.Errorf("the value of unkown shoud be empty, but %s got", view)
	}
	// http://localhost:9999/_gomemcached/scores/Tom
}

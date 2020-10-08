package session

import (
	"github.com/dashenwo/go-library/session/storage"
	redisDb "github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"testing"
)

var (
	sessionManage *Manager
)

func TestFlashes(t *testing.T) {
	rdb := redisDb.NewClusterClient(&redisDb.ClusterOptions{
		Addrs:    []string{"192.168.3.4:7001", "192.168.3.4:7002", "192.168.3.4:7003", "192.168.3.4:7004", "192.168.3.4:7005", "192.168.3.4:7006"},
		Password: "Liuqin76624291",
	})
	sessionManage = NewSessionManage(
		Storage(
			storage.NewRedisStorage(
				rdb,
			),
		),
	)
	http.HandleFunc("/", foo)
	http.ListenAndServe(":4000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	session := sessionManage.New(
		Request(r),
		Writer(w),
	)
	log.Print(session)
}

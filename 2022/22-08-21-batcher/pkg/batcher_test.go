package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

const (
	batcherSize     = 100
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second
)

type InReq struct {
	UserId    int64
	ProductId int64
}

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

func TestNew(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	b := New(
		WithSize(batcherSize),
		WithBuffer(batcherBuffer),
		WithWorker(batcherWorker),
		WithInterval(batcherInterval),
	)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % batcherWorker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*KafkaData
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*KafkaData))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			fmt.Printf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		fmt.Printf("kd:%v\n", string(kd))
	}
	s := b
	go addMsg(s)
	s.Start()
	wg.Wait()
}

func addMsg(s *Batcher) {
	in := InReq{
		ProductId: 111,
		UserId:    123456,
	}
	if err := s.Add(strconv.FormatInt(in.ProductId, 10), &KafkaData{Uid: in.UserId, Pid: in.ProductId}); err != nil {
		fmt.Printf("l.batcher.Add uid: %d pid: %d error: %v", in.UserId, in.ProductId, err)
	}
}

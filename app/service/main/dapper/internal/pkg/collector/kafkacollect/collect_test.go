package kafkacollector

import (
	"angrymiao-go/app/service/main/dapper/internal/model"
	"angrymiao-go/app/service/main/dapper/internal/pkg/process"
	"context"
	"flag"
	"fmt"
	"angrymiao-go/punk/log"
	"testing"
	"time"
)

func TestKafkaCollect(t *testing.T) {
	flag.Parse()
	log.Init(nil)
	clt, err := New("lancer_main_dapper_collector", []string{"172.18.33.163:9092", "172.18.33.164:9092", "172.18.33.165:9092"})
	if err != nil {
		t.Fatal(err)
	}
	m := process.MockProcess(func(ctx context.Context, protoSpan *model.ProtoSpan) error {
		fmt.Printf("%v\n", protoSpan)
		return nil
	})
	clt.RegisterProcess(m)
	if err := clt.Start(); err != nil {
		t.Fatal(err)
	}
	defer clt.Close()
	time.Sleep(time.Minute)
}

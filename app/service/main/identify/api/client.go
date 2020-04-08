package v1

import (
	"context"
	"fmt"

	"angrymiao-go/punk/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "TODO: ADD APP ID"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewDemoClient(cc), nil
}

// 生成 gRPC 代码
//go:generate punk tool protoc --grpc --bm api.proto

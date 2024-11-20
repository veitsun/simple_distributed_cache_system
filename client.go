package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "sunwei.top/cacheserver/cache"
)

func setupClient() {
	var opts []grpc.DialOption
	var err error
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn[0], err = grpc.Dial(address[2], opts...)
	if err != nil {
		fmt.Println("fail to dial: %v", err)
	}
	fmt.Println("Set up client for", address[2])

	conn[1], err = grpc.Dial(address[3], opts...)
	if err != nil {
		fmt.Println("fail to dial: %v", err)
	}
	fmt.Println("Set up client for", address[3])

	client[0] = pb.NewCacheClient(conn[0])
	client[1] = pb.NewCacheClient(conn[1])

}

// rpc client Set request
func CacheSet(client pb.CacheClient, req *pb.SetRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.SetCache(ctx, req)
	if err != nil {
		fmt.Println("client.SetCache failed.")
	}
}

// rpc client Get request
func CacheGet(client pb.CacheClient, req *pb.GetRequest) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	getval, err := client.GetCache(ctx, req) //getval is a pointer
	return getval.Value, err
}

// rpc client Delete request
func CacheDelete(client pb.CacheClient, req *pb.DeleteRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.DeleteCache(ctx, req)
	if err != nil {
		fmt.Println("client.DeleteCache failed.")
	}
	return err
}
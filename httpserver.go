package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gogrpc/helper"
	"gogrpc/services"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main()  {
	//必须是runtime
	gwmux := runtime.NewServeMux()
	//添加设置客户端creds
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	//注册服务，并指定后端的grpc服务地址，及相关配置
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(),gwmux,"localhost:8081",opt)
	if err != nil {
		log.Fatal(err)
	}
	//new httpserver，监听8080端口，http1.x请求->httpserver:8080->grpc-gateway转发->grpcserver:8081
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
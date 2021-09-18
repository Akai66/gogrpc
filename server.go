package main

import (
	"gogrpc/helper"
	"gogrpc/services"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	////加载证书
	//creds,err := credentials.NewServerTLSFromFile("keys/ssl.crt","keys/nopss_ssl.key")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//创建grpc server实例
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	//注册prod service server
	services.RegisterProdServiceServer(rpcServer,&services.ProdService{})
	//注册order service server
	services.RegisterOrderServiceServer(rpcServer,&services.OrderService{})
	//设置监听协议及端口
	lis,_ := net.Listen("tcp",":8081")
	//启动
	rpcServer.Serve(lis)

	////支持http服务
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request)
	//	rpcServer.ServeHTTP(writer,request)
	//})
	//httpServer := &http.Server{
	//	Addr:":8081",
	//	Handler: mux,
	//}
	//httpServer.ListenAndServeTLS("keys/ssl.crt","keys/nopss_ssl.key")
}

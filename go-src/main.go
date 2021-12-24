package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/hideyoshidan/grpc/server"
	"github.com/hideyoshidan/grpc/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// PORT use as server
	PORT string = ":1323"
)

// main
func main() {
	listenPort, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("filed to listen: %v", err)
	}

	fmt.Printf("%#v\n", "============SERVER START===============")

	// instance of gRPC server
	server := grpc.NewServer()

	// 自動生成された関数に、サーバと実際に処理を行うメソッドを実装したハンドラを設定します。
	// protoファイルで定義した`RockPaperScissorsService`に対応しています。
	pb.RegisterSampleServer(server, test.NewTest())

	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(server)

	// start server
	server.Serve(listenPort)
}

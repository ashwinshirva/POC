package stream

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	_ "google.golang.org/grpc/encoding/gzip"
	pb "pb"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"unsafe"
)

const (
	address     = "localhost:50051"
)


var (
	InterceptorOpts = []grpc_retry.CallOption{
		grpc_retry.WithCodes(codes.Unavailable),
		grpc_retry.WithMax(3),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithPerRetryTimeout(1000*time.Second),
	}
)

func CallNoStream() {
	client, conn := createClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	respFinal2, err0 := client.NoStreaming(ctx, &pb.HelloRequest{Name: "test"})
	if err0 != nil {
		log.Println("Response Size: ", unsafe.Sizeof(respFinal2))
		log.Fatalf("could not greet: %v", err0)
	}

	log.Println("Success Non Streaming")
}


func createClient() (pb.GreeterClient, *grpc.ClientConn) {
	conn := createConnection()

	client := pb.NewGreeterClient(conn)
	return client, conn
}

func createConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(InterceptorOpts...)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

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

func CallCompression() {
	client, conn := createClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	t1 := time.Now()
	_, err0 := client.ServerCompression(ctx, &pb.HelloRequest{Name: "test"}, grpc.UseCompressor("gzip"))
	if err0 != nil {
		t2 := time.Now()
		log.Println("Time taken Streaming: ", t2.Sub(t1))
		log.Fatalf("could not greet: %v", err0)
	}

	log.Println("Success Non Streaming: ")
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

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
	"constants"
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

	_, err0 := client.ServerCompression(ctx, &pb.HelloRequest{Name: "test"}, grpc.UseCompressor("gzip"))
	if err0 != nil {
		log.Fatalf("could not greet: %v", err0)
	}
	log.Println("Success Compression: ")
}


func createClient() (pb.GreeterClient, *grpc.ClientConn) {
	conn := createConnection()

	client := pb.NewGreeterClient(conn)
	return client, conn
}

func createConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(constants.Address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(InterceptorOpts...)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

package stream

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"constants"

	_ "google.golang.org/grpc/encoding/gzip"
	pb "pb"
	"io"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"encoding/binary"
	"common/helpers"
)


var (
	InterceptorOpts = []grpc_retry.CallOption{
		grpc_retry.WithCodes(codes.Unavailable),
		grpc_retry.WithMax(3),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithPerRetryTimeout(1000*time.Second),
	}
)

func CallStreaming() {
	client, conn := createClient()
	defer conn.Close()

	serverResp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "test"}, grpc.UseCompressor("gzip"))

	respFinal := pb.HashCheckResult{}
	featureParamsList := make([]*pb.FeatureHashParams, 0, 1)
	imsiList := make([]string, 0, 1)

	for i:=0;;i++ {
		resp, err := serverResp.Recv()
		if err == io.EOF {
			break
		}

		featureParamsList = append(featureParamsList, resp.GetFeatureParams()...)
		imsiList = append(imsiList, resp.GetImsi()...)

		buff, err := helpers.ConvertStructToBytes(resp)
		log.Println("Streaming rpc Response Chunk Size: ", binary.Size(buff.Bytes()))
	}

	if err != nil {
		log.Println("Client Error: ", err)
	}

	respFinal.FeatureParams = featureParamsList
	respFinal.Imsi = imsiList

	buff, err := helpers.ConvertStructToBytes(respFinal)
	log.Println("Streaming rpc Response Size: ", binary.Size(buff.Bytes()))
	log.Println("Success Streaming")
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

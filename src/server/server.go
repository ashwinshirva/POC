package main

import (
	"context"

	_ "google.golang.org/grpc/encoding/gzip"
	pb "pb"
	"net"
	"google.golang.org/grpc"
	"encoding/binary"
	"common/helpers"
	log "github.com/sirupsen/logrus"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(in *pb.HelloRequest, stream pb.Greeter_SayHelloServer) error {
	log.Info("Server: Streaming rpc called...")




	featureParamsList := make([]*pb.FeatureHashParams, 0, 1)
	imsiList := make([]string, 0, 1)

	epc := pb.Epc{HwId:"00:07:32:42:70:B3", Identifier:"SubscriptionStatusE2EPnmIdTest4"}
	hashCheckParams := &pb.FeatureHashParams{Imsi: "111111111111111", Imei:"111111111111111", StaticIp:"10.1.1.1", Profile:pb.HssProfile_Default, QosProfileName:"test_name"}

	for i:=0; i<1000000; i++ {
		resp := &pb.HashCheckResult{Result:true, Epc:&epc}

		featureParamsList = append(featureParamsList, hashCheckParams)
		imsiList = append(imsiList, "111111111111111")

		if i % 50000 == 0 {
			resp.FeatureParams = featureParamsList
			resp.Imsi = imsiList
			imsiList = make([]string, 0, 1)

			featureParamsList = make([]*pb.FeatureHashParams, 0, 1)
			if err := stream.Send(resp); err != nil {
				log.Println("Server error: ", err)

				return err
			}
		}

	}

	return nil
}

func (s *server) NoStreaming(ctx context.Context, in *pb.HelloRequest) (*pb.HashCheckResult, error) {
	log.Info("Server: Generic rpc called...")


	featureParamsList := make([]*pb.FeatureHashParams, 0, 1)
	imsiList := make([]string, 0, 1)

	epc := &pb.Epc{HwId:"00:07:32:42:70:B3", Identifier:"blrsyvegreenui@nokia.com::804f7318-79ce-4f3b-b227-56b2b7a5597b", }
	hashCheckParams := &pb.FeatureHashParams{Imsi:"111111111111111", StaticIp:"10.1.1.1", Imei:"111111111111111", Profile:pb.HssProfile_Default, QosProfileName:"test_name"}
	resp := &pb.HashCheckResult{Result:true}

	for i:=0; i<1000000; i++ {
		featureParamsList = append(featureParamsList, hashCheckParams)
		imsiList = append(imsiList, "111111111111111")
	}

	resp.FeatureParams = featureParamsList
	resp.Imsi = imsiList
	resp.Epc = epc


	buff, err := helpers.ConvertStructToBytes(resp)
	if err != nil {
		return &pb.HashCheckResult{}, err
	}
	log.Println("No Streaming rpc Response Size: ", binary.Size(buff.Bytes()))


	return resp, nil
}

func (s *server) ServerCompression(ctx context.Context, in *pb.HelloRequest) (*pb.HashCheckResult, error) {
	log.Info("Server: Compression rpc called...")


	featureParamsList := make([]*pb.FeatureHashParams, 0, 1)
	imsiList := make([]string, 0, 1)

	epc := &pb.Epc{HwId:"00:07:32:42:70:B3", Identifier:"blrsyvegreenui@nokia.com::804f7318-79ce-4f3b-b227-56b2b7a5597b"}
	resp := &pb.HashCheckResult{Result:true}
	hashCheckParams := &pb.FeatureHashParams{Imsi:"111111111111111", StaticIp:"10.1.1.1", Imei:"111111111111111", Profile:pb.HssProfile_Default, QosProfileName:"test_name"}

	for i:=0; i<1000000; i++ {
		featureParamsList = append(featureParamsList, hashCheckParams)
		imsiList = append(imsiList, "111111111111111")
	}

	resp.FeatureParams = featureParamsList
	resp.Imsi = imsiList
	resp.Epc = epc

	//gzip.SetLevel(10)

	buff, _ := helpers.ConvertStructToBytes(resp)
	log.Println("Server: Compression rpc Response Size: ", binary.Size(buff.Bytes()))

	return resp, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

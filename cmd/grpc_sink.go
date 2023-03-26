package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cikupin/learn-firehose/config"
	pb "github.com/cikupin/learn-firehose/payload"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcSinkCmd = &cobra.Command{
	Use:   "grpc-sink",
	Short: "sink for GRPC message",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv()
		initGrpcSink()
	},
}

func init() {
	rootCmd.AddCommand(grpcSinkCmd)
}

type SinkGrpc struct {
	pb.UnimplementedGrpcSinkServer
}

func (s *SinkGrpc) Receive(ctx context.Context, flag *pb.Flag) (*pb.FlagReply, error) {
	fmt.Println("======= [DATA] ==========")
	fmt.Println("id:", flag.GetId())
	fmt.Println()

	return &pb.FlagReply{
		Message: fmt.Sprintf("successfully receive data with id %s", flag.GetId()),
	}, nil
}

func initGrpcSink() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.App.Port))
	if err != nil {
		log.Fatalln(err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterGrpcSinkServer(server, &SinkGrpc{})

	fmt.Printf("server listening at %v\n", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Fatalln(err.Error())
	}
}

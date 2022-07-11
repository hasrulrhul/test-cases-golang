package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/hasrulrhul/test-cases/proto_model"
	"github.com/lovoo/goka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port   = flag.Int("port", 2000, "The server port")
	wallet []*pb.Wallet

	brokers             = []string{"localhost:29092"}
	topic   goka.Stream = "example-visit-clicks-input"
	group   goka.Group  = "example-visit-group"

	tmc *goka.TopicManagerConfig
)

func init() {
	// This sets the default replication to 1. If you have more then one broker
	// the default configuration can be used.
	tmc = goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1
}

func initWallet() {
	wallet1 := &pb.Wallet{
		WalletId: 1,
		Amount:   1000.00,
	}
	wallet2 := &pb.Wallet{
		WalletId: 2,
		Amount:   15000.00,
	}

	wallet = append(wallet, wallet1)
	wallet = append(wallet, wallet2)
}

type walletServer struct {
	pb.UnimplementedWalletServicesServer
}

func main() {
	initWallet()
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return
	}

	// Craete New Server Instances
	grpc := grpc.NewServer()

	// Register Service nya di sini
	pb.RegisterWalletServicesServer(grpc, &walletServer{})
	reflection.Register(grpc)

	if e := grpc.Serve(list); e != nil {
		panic(e)
	}

	log.Printf("Server Running on :%v", list.Addr().String())
}

func (s *walletServer) Deposit(ctx context.Context, in *pb.Wallet) (*pb.Wallet, error) {
	res := pb.Wallet{}
	res.WalletId = in.WalletId
	res.Amount = in.Amount
	if in.WalletId == 0 {
		return nil, nil
	}
	wallet = append(wallet, in)
	return &res, nil
}

func (s *walletServer) Details(ctx context.Context, id *pb.WalletId) (*pb.Detail, error) {
	var newList []*pb.Wallet
	for _, wallets := range wallet {
		if wallets.WalletId == id.WalletId {
			// log.Printf("Received: %v", wallets)
			newList = append(newList, wallets)
		}
	}

	wallets := &pb.Detail{
		List: newList,
	}
	return wallets, nil
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/hasrulrhul/test-cases/proto_model"
	"google.golang.org/grpc"
)

const (
	address = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewWalletServicesClient(conn)

	r := gin.Default()
	r.POST("deposit", func(ctx *gin.Context) {
		var wallet pb.Wallet
		err := ctx.BindJSON(&wallet)
		if err != nil {
			panic(err)
		}
		if response, err := client.Deposit(ctx, &wallet); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": response,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.GET("details", func(ctx *gin.Context) {
		var wallet pb.WalletId
		err := ctx.BindJSON(&wallet)
		if err != nil {
			panic(err)
		}
		if response, err := client.Details(ctx, &wallet); err == nil {
			ctx.JSON(http.StatusOK, response)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.Run(":1234")
}

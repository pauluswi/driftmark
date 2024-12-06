package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pauluswi/driftmark/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewFundTransferServiceClient(conn)

	// Create a transfer request
	req := &pb.TransferRequest{
		TransactionId:      "TXN12345",
		SourceAccount:      "1234567890",
		DestinationAccount: "9876543210",
		Amount:             100.0,
		Currency:           "USD",
		TransferType:       "debit",
	}

	log.Printf("Initiating fund transfer: %v", req)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ProcessFundTransfer(ctx, req)
	if err != nil {
		log.Fatalf("Error during fund transfer: %v", err)
	}

	log.Printf("Fund Transfer Response: TransactionID=%s, Status=%s, Message=%s",
		res.TransactionId, res.Status, res.Message)
}

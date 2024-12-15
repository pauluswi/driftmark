package main

import (
	"context"
	"time"

	pb "github.com/pauluswi/driftmark/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Configure zap logger
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "./logs/log.json"}
	logger, _ := config.Build()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
	}
	defer conn.Close()

	client := pb.NewFundTransferServiceClient(conn)

	req := &pb.TransferRequest{
		TransactionId:      "TXN12345",
		SourceAccount:      "1234567890",
		DestinationAccount: "9876543210",
		Amount:             100.0,
		Currency:           "USD",
		TransferType:       "debit",
	}

	logger.Info("Sending fund transfer request",
		zap.String("transaction_id", req.TransactionId),
		zap.String("source_account", req.SourceAccount),
		zap.String("destination_account", req.DestinationAccount),
		zap.Float64("amount", req.Amount),
		zap.String("currency", req.Currency),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ProcessFundTransfer(ctx, req)
	if err != nil {
		logger.Error("Fund transfer failed", zap.Error(err))
		return
	}

	logger.Info("Received response",
		zap.String("transaction_id", res.TransactionId),
		zap.String("status", res.Status),
		zap.String("message", res.Message),
	)
}

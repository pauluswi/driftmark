package main

import (
	"context"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/pauluswi/driftmark/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FundTransferServiceServer struct {
	pb.UnimplementedFundTransferServiceServer
	accounts map[string]float64
	mutex    sync.Mutex
	logger   *zap.Logger
}

func NewFundTransferServiceServer() *FundTransferServiceServer {
	// Configure zap logger to write to logs/log.json
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "./logs/log.json"}

	logger, _ := config.Build()
	return &FundTransferServiceServer{
		accounts: map[string]float64{
			"1234567890": 1000.0,
			"9876543210": 500.0,
		},
		logger: logger,
	}
}

func (s *FundTransferServiceServer) ProcessFundTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	s.logger.Info("Processing fund transfer",
		zap.String("transaction_id", req.TransactionId),
		zap.String("source_account", req.SourceAccount),
		zap.String("destination_account", req.DestinationAccount),
		zap.Float64("amount", req.Amount),
		zap.String("currency", req.Currency),
		zap.String("transfer_type", req.TransferType),
	)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Validate accounts
	sourceBalance, srcExists := s.accounts[req.SourceAccount]
	_, destExists := s.accounts[req.DestinationAccount]

	if !srcExists {
		s.logger.Error("Source account not found", zap.String("account", req.SourceAccount))
		return nil, status.Errorf(codes.NotFound, "Source account %s not found", req.SourceAccount)
	}

	if !destExists {
		s.logger.Error("Destination account not found", zap.String("account", req.DestinationAccount))
		return nil, status.Errorf(codes.NotFound, "Destination account %s not found", req.DestinationAccount)
	}

	// Check sufficient funds
	if req.TransferType == "debit" && sourceBalance < req.Amount {
		s.logger.Error("Insufficient funds", zap.String("account", req.SourceAccount), zap.Float64("balance", sourceBalance))
		return nil, status.Errorf(codes.FailedPrecondition, "Insufficient funds in account %s", req.SourceAccount)
	}

	// Process fund transfer
	if req.TransferType == "debit" {
		s.accounts[req.SourceAccount] -= req.Amount
		s.accounts[req.DestinationAccount] += req.Amount
	} else {
		s.accounts[req.SourceAccount] += req.Amount
		s.accounts[req.DestinationAccount] -= req.Amount
	}

	s.logger.Info("Fund transfer successful",
		zap.String("transaction_id", req.TransactionId),
		zap.String("source_account", req.SourceAccount),
		zap.String("destination_account", req.DestinationAccount),
		zap.Float64("amount", req.Amount),
	)

	return &pb.TransferResponse{
		TransactionId: req.TransactionId,
		Status:        "SUCCESS",
		Message:       "Fund transfer processed successfully",
	}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Failed to listen: " + err.Error())
	}

	server := NewFundTransferServiceServer()
	grpcServer := grpc.NewServer()

	pb.RegisterFundTransferServiceServer(grpcServer, server)
	server.logger.Info("Fund Transfer Service Server running on port 50051...")

	if err := grpcServer.Serve(lis); err != nil {
		server.logger.Fatal("Failed to serve gRPC server", zap.Error(err))
	}
}

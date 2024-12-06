package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/pauluswi/driftmark/proto"
	"google.golang.org/grpc"
)

// FundTransferServiceServer implements the gRPC service
type FundTransferServiceServer struct {
	pb.UnimplementedFundTransferServiceServer
	accounts map[string]float64 // Mock account balances
	mutex    sync.Mutex         // To handle concurrent updates
}

// NewFundTransferServiceServer initializes the server with mock data
func NewFundTransferServiceServer() *FundTransferServiceServer {
	return &FundTransferServiceServer{
		accounts: map[string]float64{
			"1234567890": 1000.0, // Source account
			"9876543210": 500.0,  // Destination account
		},
	}
}

// ProcessFundTransfer handles fund transfers
func (s *FundTransferServiceServer) ProcessFundTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	log.Printf("Processing transfer: TransactionID=%s, Source=%s, Destination=%s, Amount=%.2f %s",
		req.TransactionId, req.SourceAccount, req.DestinationAccount, req.Amount, req.Currency)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Validate source and destination accounts
	sourceBalance, srcExists := s.accounts[req.SourceAccount]
	_, destExists := s.accounts[req.DestinationAccount]

	if !srcExists {
		return &pb.TransferResponse{
			TransactionId: req.TransactionId,
			Status:        "FAILED",
			Message:       "Source account not found",
		}, nil
	}

	if !destExists {
		return &pb.TransferResponse{
			TransactionId: req.TransactionId,
			Status:        "FAILED",
			Message:       "Destination account not found",
		}, nil
	}

	// Check sufficient funds for debit transactions
	if req.TransferType == "debit" && sourceBalance < req.Amount {
		return &pb.TransferResponse{
			TransactionId: req.TransactionId,
			Status:        "FAILED",
			Message:       "Insufficient funds in source account",
		}, nil
	}

	// Process the fund transfer
	if req.TransferType == "debit" {
		s.accounts[req.SourceAccount] -= req.Amount
		s.accounts[req.DestinationAccount] += req.Amount
	} else {
		s.accounts[req.SourceAccount] += req.Amount
		s.accounts[req.DestinationAccount] -= req.Amount
	}

	log.Printf("Transfer successful: TransactionID=%s", req.TransactionId)
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
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFundTransferServiceServer(grpcServer, NewFundTransferServiceServer())

	log.Println("Fund Transfer Service Server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

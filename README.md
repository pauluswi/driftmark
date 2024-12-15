# driftmark - Fund Transfer Service (gRPC)

[![Build Server](https://github.com/pauluswi/driftmark/actions/workflows/go.yml/badge.svg)](https://github.com/pauluswi/driftmark/actions/workflows/go.yml)

A **gRPC-based Fund Transfer Service** implemented in Go for real-time transaction processing between accounts. This project demonstrates structured logging with **Zap**, robust error handling with gRPC status codes, and unit tests for ensuring reliability.

---

## Features

- **Real-Time Fund Transfer**: Process fund transfers between accounts with transaction details.
- **Structured Logging**: Logs using [Zap](https://github.com/uber-go/zap) for better observability.
- **Error Handling**: Implements gRPC error codes for precise error reporting.
- **Proto Definitions**: Based on `Protocol Buffers` for defining the service and messages.

---

## Project Structure

```python
.
├── proto
│   └── fund_transfer.proto    # Protocol Buffer definitions for the Fund Transfer Service
├── server
│   ├── main.go                # Server implementation for the gRPC service
├── client
│   └── main.go                # Client implementation to test the service
├── logs
│   └── logs.json              # (Optional) Sample structured logs generated by the service
├── go.mod                     # Go module dependencies
├── go.sum                     # Go module checksums
├── README.md                  # Documentation for the project
└── LICENSE                    # License for the project
```


---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (1.21 or later)
- [Protocol Buffers Compiler](https://grpc.io/docs/protoc-installation/)
- Install `protoc-gen-go` and `protoc-gen-go-grpc`:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
  
### Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/golang-grpc-portfolio.git
   cd golang-grpc-portfolio
   ```
2. **Generate the gRPC code**:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/fund_transfer.proto
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

### Running the Application
1. **Start the Server**:
   ```bash
   go run server/main.go
   ```
2. **Run the Client**:
   ```bash
   go run client/main.go
   ```

## API Definition
The Fund Transfer Service is defined in proto/fund_transfer.proto:

### **Service**: FundTransferService
#### **RPC Method**: ProcessFundTransfer

- **Request**:
  - ***transaction_id***: Unique ID for the transaction.
  - ***source_account***: Source account number.
  - ***destination_account***: Destination account number.
  - ***amount***: Amount to transfer.
  - ***currency***: Currency of the transaction (e.g., USD).
  - ***transfer_type***: Type of transfer (debit or credit).
- **Response**:
  - ***transaction_id***: Same as the request ID.
  - ***status***: SUCCESS or FAILED.
  - ***message***: Additional details about the transfer.

#### **Sample Request/Response**:

- **Request**:
   ```json
      {
         "transaction_id": "TXN12345",
         "source_account": "1234567890",
         "destination_account": "9876543210",
         "amount": 100.0,
         "currency": "USD",
         "transfer_type": "debit"
      }
   ```

- **Response**:
   ```json
      {
         "transaction_id": "TXN12345",
         "status": "SUCCESS",
         "message": "Fund transfer processed successfully"
      }
   ```

## Logging
Structured logs are implemented using Zap and stored in the logs/log.json file. Logs include transaction details, statuses, and errors.

Example logs/log.json:
```json
{"level":"info","msg":"Processing fund transfer","transaction_id":"TXN12345","source_account":"1234567890","destination_account":"9876543210","amount":100,"currency":"USD","transfer_type":"debit"}
{"level":"info","msg":"Fund transfer successful","transaction_id":"TXN12345","source_account":"1234567890","destination_account":"9876543210","amount":100}
{"level":"info","msg":"Received response","transaction_id":"TXN12345","status":"SUCCESS","message":"Fund transfer processed successfully"}
      
```

## License
This project is licensed under the Apache License. 



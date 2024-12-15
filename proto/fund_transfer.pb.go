// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/fund_transfer.proto

package fundtransfer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for fund transfer
type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId      string  `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	SourceAccount      string  `protobuf:"bytes,2,opt,name=source_account,json=sourceAccount,proto3" json:"source_account,omitempty"`
	DestinationAccount string  `protobuf:"bytes,3,opt,name=destination_account,json=destinationAccount,proto3" json:"destination_account,omitempty"`
	Amount             float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`                             // e.g., "USD", "IDR"
	TransferType       string  `protobuf:"bytes,6,opt,name=transfer_type,json=transferType,proto3" json:"transfer_type,omitempty"` // e.g., "debit", "credit"
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_proto_fund_transfer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fund_transfer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_fund_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *TransferRequest) GetSourceAccount() string {
	if x != nil {
		return x.SourceAccount
	}
	return ""
}

func (x *TransferRequest) GetDestinationAccount() string {
	if x != nil {
		return x.DestinationAccount
	}
	return ""
}

func (x *TransferRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferRequest) GetTransferType() string {
	if x != nil {
		return x.TransferType
	}
	return ""
}

// Response message for fund transfer
type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status        string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`   // e.g., "SUCCESS", "FAILED"
	Message       string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // Additional details about the transfer
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	mi := &file_proto_fund_transfer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fund_transfer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_proto_fund_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *TransferResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransferResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_fund_transfer_proto protoreflect.FileDescriptor

var file_proto_fund_transfer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x75, 0x6e,
	0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0xe9, 0x01, 0x0a, 0x0f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x6b, 0x0a, 0x13, 0x46, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x46, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61,
	0x75, 0x6c, 0x75, 0x73, 0x77, 0x69, 0x2f, 0x64, 0x72, 0x69, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fund_transfer_proto_rawDescOnce sync.Once
	file_proto_fund_transfer_proto_rawDescData = file_proto_fund_transfer_proto_rawDesc
)

func file_proto_fund_transfer_proto_rawDescGZIP() []byte {
	file_proto_fund_transfer_proto_rawDescOnce.Do(func() {
		file_proto_fund_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fund_transfer_proto_rawDescData)
	})
	return file_proto_fund_transfer_proto_rawDescData
}

var file_proto_fund_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_fund_transfer_proto_goTypes = []any{
	(*TransferRequest)(nil),  // 0: fundtransfer.TransferRequest
	(*TransferResponse)(nil), // 1: fundtransfer.TransferResponse
}
var file_proto_fund_transfer_proto_depIdxs = []int32{
	0, // 0: fundtransfer.FundTransferService.ProcessFundTransfer:input_type -> fundtransfer.TransferRequest
	1, // 1: fundtransfer.FundTransferService.ProcessFundTransfer:output_type -> fundtransfer.TransferResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_fund_transfer_proto_init() }
func file_proto_fund_transfer_proto_init() {
	if File_proto_fund_transfer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_fund_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fund_transfer_proto_goTypes,
		DependencyIndexes: file_proto_fund_transfer_proto_depIdxs,
		MessageInfos:      file_proto_fund_transfer_proto_msgTypes,
	}.Build()
	File_proto_fund_transfer_proto = out.File
	file_proto_fund_transfer_proto_rawDesc = nil
	file_proto_fund_transfer_proto_goTypes = nil
	file_proto_fund_transfer_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: history.proto

package coins_service

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

type HistoryUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *HistoryUserRequest) Reset() {
	*x = HistoryUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryUserRequest) ProtoMessage() {}

func (x *HistoryUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryUserRequest.ProtoReflect.Descriptor instead.
func (*HistoryUserRequest) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type HistoriesUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status            string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	ConfirmImg        string `protobuf:"bytes,4,opt,name=confirm_img,json=confirmImg,proto3" json:"confirm_img,omitempty"`
	CoinAmount        string `protobuf:"bytes,5,opt,name=coin_amount,json=coinAmount,proto3" json:"coin_amount,omitempty"`
	CoinPrice         string `protobuf:"bytes,6,opt,name=coin_price,json=coinPrice,proto3" json:"coin_price,omitempty"`
	AllPrice          string `protobuf:"bytes,7,opt,name=all_price,json=allPrice,proto3" json:"all_price,omitempty"`
	Address           string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	CardNumber        string `protobuf:"bytes,9,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardName          string `protobuf:"bytes,10,opt,name=card_name,json=cardName,proto3" json:"card_name,omitempty"`
	DateTime          string `protobuf:"bytes,11,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	TransactionStatus string `protobuf:"bytes,12,opt,name=transaction_status,json=transactionStatus,proto3" json:"transaction_status,omitempty"`
	CoinId            string `protobuf:"bytes,13,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	UserId            string `protobuf:"bytes,14,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *HistoriesUser) Reset() {
	*x = HistoriesUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoriesUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoriesUser) ProtoMessage() {}

func (x *HistoriesUser) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoriesUser.ProtoReflect.Descriptor instead.
func (*HistoriesUser) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{1}
}

func (x *HistoriesUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HistoriesUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HistoriesUser) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HistoriesUser) GetConfirmImg() string {
	if x != nil {
		return x.ConfirmImg
	}
	return ""
}

func (x *HistoriesUser) GetCoinAmount() string {
	if x != nil {
		return x.CoinAmount
	}
	return ""
}

func (x *HistoriesUser) GetCoinPrice() string {
	if x != nil {
		return x.CoinPrice
	}
	return ""
}

func (x *HistoriesUser) GetAllPrice() string {
	if x != nil {
		return x.AllPrice
	}
	return ""
}

func (x *HistoriesUser) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HistoriesUser) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *HistoriesUser) GetCardName() string {
	if x != nil {
		return x.CardName
	}
	return ""
}

func (x *HistoriesUser) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *HistoriesUser) GetTransactionStatus() string {
	if x != nil {
		return x.TransactionStatus
	}
	return ""
}

func (x *HistoriesUser) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *HistoriesUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type HistoryUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Histories []*HistoriesUser `protobuf:"bytes,1,rep,name=histories,proto3" json:"histories,omitempty"`
	Count     int64            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *HistoryUserResponse) Reset() {
	*x = HistoryUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryUserResponse) ProtoMessage() {}

func (x *HistoryUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryUserResponse.ProtoReflect.Descriptor instead.
func (*HistoryUserResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryUserResponse) GetHistories() []*HistoriesUser {
	if x != nil {
		return x.Histories
	}
	return nil
}

func (x *HistoryUserResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TransactionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TransactionStatus) Reset() {
	*x = TransactionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStatus) ProtoMessage() {}

func (x *TransactionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStatus.ProtoReflect.Descriptor instead.
func (*TransactionStatus) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionStatus) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TransactionStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransactionStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HistoryUserWithStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoryUser   *HistoriesUser     `protobuf:"bytes,1,opt,name=history_user,json=historyUser,proto3" json:"history_user,omitempty"`
	HistoryStatus *TransactionStatus `protobuf:"bytes,2,opt,name=history_status,json=historyStatus,proto3" json:"history_status,omitempty"`
}

func (x *HistoryUserWithStatus) Reset() {
	*x = HistoryUserWithStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryUserWithStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryUserWithStatus) ProtoMessage() {}

func (x *HistoryUserWithStatus) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryUserWithStatus.ProtoReflect.Descriptor instead.
func (*HistoryUserWithStatus) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryUserWithStatus) GetHistoryUser() *HistoriesUser {
	if x != nil {
		return x.HistoryUser
	}
	return nil
}

func (x *HistoryUserWithStatus) GetHistoryStatus() *TransactionStatus {
	if x != nil {
		return x.HistoryStatus
	}
	return nil
}

type HistoryMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoryWithStatus []*HistoryUserWithStatus `protobuf:"bytes,1,rep,name=history_with_status,json=historyWithStatus,proto3" json:"history_with_status,omitempty"`
}

func (x *HistoryMessageResponse) Reset() {
	*x = HistoryMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryMessageResponse) ProtoMessage() {}

func (x *HistoryMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryMessageResponse.ProtoReflect.Descriptor instead.
func (*HistoryMessageResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryMessageResponse) GetHistoryWithStatus() []*HistoryUserWithStatus {
	if x != nil {
		return x.HistoryWithStatus
	}
	return nil
}

var File_history_proto protoreflect.FileDescriptor

var file_history_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9f, 0x03, 0x0a, 0x0d, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x5f, 0x69, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x6d, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x69, 0x6e,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x69,
	0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x13, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xa1, 0x01, 0x0a, 0x15, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0e, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x6e, 0x0a, 0x16, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x13, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x11, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xe4, 0x02, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c,
	0x12, 0x14, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x21, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_history_proto_rawDescOnce sync.Once
	file_history_proto_rawDescData = file_history_proto_rawDesc
)

func file_history_proto_rawDescGZIP() []byte {
	file_history_proto_rawDescOnce.Do(func() {
		file_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_history_proto_rawDescData)
	})
	return file_history_proto_rawDescData
}

var file_history_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_history_proto_goTypes = []any{
	(*HistoryUserRequest)(nil),     // 0: coins_service.HistoryUserRequest
	(*HistoriesUser)(nil),          // 1: coins_service.HistoriesUser
	(*HistoryUserResponse)(nil),    // 2: coins_service.HistoryUserResponse
	(*TransactionStatus)(nil),      // 3: coins_service.TransactionStatus
	(*HistoryUserWithStatus)(nil),  // 4: coins_service.HistoryUserWithStatus
	(*HistoryMessageResponse)(nil), // 5: coins_service.HistoryMessageResponse
	(*Empty)(nil),                  // 6: coins_service.Empty
}
var file_history_proto_depIdxs = []int32{
	1, // 0: coins_service.HistoryUserResponse.histories:type_name -> coins_service.HistoriesUser
	1, // 1: coins_service.HistoryUserWithStatus.history_user:type_name -> coins_service.HistoriesUser
	3, // 2: coins_service.HistoryUserWithStatus.history_status:type_name -> coins_service.TransactionStatus
	4, // 3: coins_service.HistoryMessageResponse.history_with_status:type_name -> coins_service.HistoryUserWithStatus
	0, // 4: coins_service.HistoryService.HistoryUser:input_type -> coins_service.HistoryUserRequest
	6, // 5: coins_service.HistoryService.HistoryUserAll:input_type -> coins_service.Empty
	0, // 6: coins_service.HistoryService.HistoryMessage:input_type -> coins_service.HistoryUserRequest
	0, // 7: coins_service.HistoryService.UpdateHistoryRead:input_type -> coins_service.HistoryUserRequest
	2, // 8: coins_service.HistoryService.HistoryUser:output_type -> coins_service.HistoryUserResponse
	2, // 9: coins_service.HistoryService.HistoryUserAll:output_type -> coins_service.HistoryUserResponse
	5, // 10: coins_service.HistoryService.HistoryMessage:output_type -> coins_service.HistoryMessageResponse
	6, // 11: coins_service.HistoryService.UpdateHistoryRead:output_type -> coins_service.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_history_proto_init() }
func file_history_proto_init() {
	if File_history_proto != nil {
		return
	}
	file_coin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_history_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HistoryUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_history_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HistoriesUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_history_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HistoryUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_history_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_history_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*HistoryUserWithStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_history_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*HistoryMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_history_proto_goTypes,
		DependencyIndexes: file_history_proto_depIdxs,
		MessageInfos:      file_history_proto_msgTypes,
	}.Build()
	File_history_proto = out.File
	file_history_proto_rawDesc = nil
	file_history_proto_goTypes = nil
	file_history_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: coin.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{0}
}

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CoinIcon      string           `protobuf:"bytes,3,opt,name=coin_icon,json=coinIcon,proto3" json:"coin_icon,omitempty"`
	CoinBuyPrice  string           `protobuf:"bytes,4,opt,name=coin_buy_price,json=coinBuyPrice,proto3" json:"coin_buy_price,omitempty"`
	CoinSellPrice string           `protobuf:"bytes,5,opt,name=coin_sell_price,json=coinSellPrice,proto3" json:"coin_sell_price,omitempty"`
	Halfcoins     []*HalfCoinPrice `protobuf:"bytes,6,rep,name=halfcoins,proto3" json:"halfcoins,omitempty"`
	Address       string           `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	CardNumber    string           `protobuf:"bytes,8,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Status        string           `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	ImageId       string           `protobuf:"bytes,10,opt,name=imageId,proto3" json:"imageId,omitempty"`
	CreatedAt     string           `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string           `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{1}
}

func (x *Coin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Coin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coin) GetCoinIcon() string {
	if x != nil {
		return x.CoinIcon
	}
	return ""
}

func (x *Coin) GetCoinBuyPrice() string {
	if x != nil {
		return x.CoinBuyPrice
	}
	return ""
}

func (x *Coin) GetCoinSellPrice() string {
	if x != nil {
		return x.CoinSellPrice
	}
	return ""
}

func (x *Coin) GetHalfcoins() []*HalfCoinPrice {
	if x != nil {
		return x.Halfcoins
	}
	return nil
}

func (x *Coin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Coin) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *Coin) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Coin) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Coin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Coin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CoinPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CoinPrimaryKey) Reset() {
	*x = CoinPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinPrimaryKey) ProtoMessage() {}

func (x *CoinPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinPrimaryKey.ProtoReflect.Descriptor instead.
func (*CoinPrimaryKey) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{2}
}

func (x *CoinPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CoinBuyPrice  string           `protobuf:"bytes,2,opt,name=coin_buy_price,json=coinBuyPrice,proto3" json:"coin_buy_price,omitempty"`
	CoinSellPrice string           `protobuf:"bytes,3,opt,name=coin_sell_price,json=coinSellPrice,proto3" json:"coin_sell_price,omitempty"`
	Address       string           `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Halfcoins     []*HalfCoinPrice `protobuf:"bytes,5,rep,name=halfcoins,proto3" json:"halfcoins,omitempty"`
	CardNumber    string           `protobuf:"bytes,6,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Status        string           `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	ImageId       string           `protobuf:"bytes,8,opt,name=imageId,proto3" json:"imageId,omitempty"`
}

func (x *CreateCoin) Reset() {
	*x = CreateCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoin) ProtoMessage() {}

func (x *CreateCoin) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoin.ProtoReflect.Descriptor instead.
func (*CreateCoin) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCoin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCoin) GetCoinBuyPrice() string {
	if x != nil {
		return x.CoinBuyPrice
	}
	return ""
}

func (x *CreateCoin) GetCoinSellPrice() string {
	if x != nil {
		return x.CoinSellPrice
	}
	return ""
}

func (x *CreateCoin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateCoin) GetHalfcoins() []*HalfCoinPrice {
	if x != nil {
		return x.Halfcoins
	}
	return nil
}

func (x *CreateCoin) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *CreateCoin) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateCoin) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

type UpdateCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CoinIcon      string           `protobuf:"bytes,3,opt,name=coin_icon,json=coinIcon,proto3" json:"coin_icon,omitempty"`
	CoinBuyPrice  string           `protobuf:"bytes,4,opt,name=coin_buy_price,json=coinBuyPrice,proto3" json:"coin_buy_price,omitempty"`
	CoinSellPrice string           `protobuf:"bytes,5,opt,name=coin_sell_price,json=coinSellPrice,proto3" json:"coin_sell_price,omitempty"`
	Address       string           `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Halfcoins     []*HalfCoinPrice `protobuf:"bytes,7,rep,name=halfcoins,proto3" json:"halfcoins,omitempty"`
	CardNumber    string           `protobuf:"bytes,8,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Status        string           `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateCoin) Reset() {
	*x = UpdateCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoin) ProtoMessage() {}

func (x *UpdateCoin) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoin.ProtoReflect.Descriptor instead.
func (*UpdateCoin) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCoin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCoin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCoin) GetCoinIcon() string {
	if x != nil {
		return x.CoinIcon
	}
	return ""
}

func (x *UpdateCoin) GetCoinBuyPrice() string {
	if x != nil {
		return x.CoinBuyPrice
	}
	return ""
}

func (x *UpdateCoin) GetCoinSellPrice() string {
	if x != nil {
		return x.CoinSellPrice
	}
	return ""
}

func (x *UpdateCoin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateCoin) GetHalfcoins() []*HalfCoinPrice {
	if x != nil {
		return x.Halfcoins
	}
	return nil
}

func (x *UpdateCoin) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *UpdateCoin) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetListCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListCoinRequest) Reset() {
	*x = GetListCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCoinRequest) ProtoMessage() {}

func (x *GetListCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCoinRequest.ProtoReflect.Descriptor instead.
func (*GetListCoinRequest) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{5}
}

func (x *GetListCoinRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListCoinRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListCoinRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Coins []*Coin `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins,omitempty"`
}

func (x *GetListCoinResponse) Reset() {
	*x = GetListCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCoinResponse) ProtoMessage() {}

func (x *GetListCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCoinResponse.ProtoReflect.Descriptor instead.
func (*GetListCoinResponse) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{6}
}

func (x *GetListCoinResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListCoinResponse) GetCoins() []*Coin {
	if x != nil {
		return x.Coins
	}
	return nil
}

type HalfCoinPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HalfCoinAmount string `protobuf:"bytes,1,opt,name=halfCoinAmount,proto3" json:"halfCoinAmount,omitempty"`
	HalfCoinPrice  string `protobuf:"bytes,2,opt,name=halfCoinPrice,proto3" json:"halfCoinPrice,omitempty"`
}

func (x *HalfCoinPrice) Reset() {
	*x = HalfCoinPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalfCoinPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalfCoinPrice) ProtoMessage() {}

func (x *HalfCoinPrice) ProtoReflect() protoreflect.Message {
	mi := &file_coin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalfCoinPrice.ProtoReflect.Descriptor instead.
func (*HalfCoinPrice) Descriptor() ([]byte, []int) {
	return file_coin_proto_rawDescGZIP(), []int{7}
}

func (x *HalfCoinPrice) GetHalfCoinAmount() string {
	if x != nil {
		return x.HalfCoinAmount
	}
	return ""
}

func (x *HalfCoinPrice) GetHalfCoinPrice() string {
	if x != nil {
		return x.HalfCoinPrice
	}
	return ""
}

var File_coin_proto protoreflect.FileDescriptor

var file_coin_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f,
	0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0xfc, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x42, 0x75, 0x79, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x6c,
	0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x69, 0x6e, 0x53, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x68, 0x61, 0x6c, 0x66, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x48, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x09, 0x68,
	0x61, 0x6c, 0x66, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x69, 0x6e,
	0x5f, 0x62, 0x75, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x42, 0x75, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x6c,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x3a, 0x0a, 0x09, 0x68, 0x61, 0x6c, 0x66, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x09, 0x68, 0x61, 0x6c, 0x66, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22,
	0xaa, 0x02, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x63, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x42, 0x75, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x73, 0x65,
	0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x68, 0x61, 0x6c, 0x66, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x61, 0x6c, 0x66, 0x43,
	0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x09, 0x68, 0x61, 0x6c, 0x66, 0x63, 0x6f,
	0x69, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73,
	0x22, 0x5d, 0x0a, 0x0d, 0x48, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x61, 0x6c, 0x66, 0x43,
	0x6f, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x6c,
	0x66, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x68, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32,
	0xe6, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coin_proto_rawDescOnce sync.Once
	file_coin_proto_rawDescData = file_coin_proto_rawDesc
)

func file_coin_proto_rawDescGZIP() []byte {
	file_coin_proto_rawDescOnce.Do(func() {
		file_coin_proto_rawDescData = protoimpl.X.CompressGZIP(file_coin_proto_rawDescData)
	})
	return file_coin_proto_rawDescData
}

var file_coin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_coin_proto_goTypes = []any{
	(*Empty)(nil),               // 0: coins_service.Empty
	(*Coin)(nil),                // 1: coins_service.Coin
	(*CoinPrimaryKey)(nil),      // 2: coins_service.CoinPrimaryKey
	(*CreateCoin)(nil),          // 3: coins_service.CreateCoin
	(*UpdateCoin)(nil),          // 4: coins_service.UpdateCoin
	(*GetListCoinRequest)(nil),  // 5: coins_service.GetListCoinRequest
	(*GetListCoinResponse)(nil), // 6: coins_service.GetListCoinResponse
	(*HalfCoinPrice)(nil),       // 7: coins_service.HalfCoinPrice
}
var file_coin_proto_depIdxs = []int32{
	7, // 0: coins_service.Coin.halfcoins:type_name -> coins_service.HalfCoinPrice
	7, // 1: coins_service.CreateCoin.halfcoins:type_name -> coins_service.HalfCoinPrice
	7, // 2: coins_service.UpdateCoin.halfcoins:type_name -> coins_service.HalfCoinPrice
	1, // 3: coins_service.GetListCoinResponse.coins:type_name -> coins_service.Coin
	3, // 4: coins_service.CoinsService.Create:input_type -> coins_service.CreateCoin
	2, // 5: coins_service.CoinsService.GetById:input_type -> coins_service.CoinPrimaryKey
	5, // 6: coins_service.CoinsService.GetList:input_type -> coins_service.GetListCoinRequest
	4, // 7: coins_service.CoinsService.Update:input_type -> coins_service.UpdateCoin
	2, // 8: coins_service.CoinsService.Delete:input_type -> coins_service.CoinPrimaryKey
	2, // 9: coins_service.CoinsService.Create:output_type -> coins_service.CoinPrimaryKey
	1, // 10: coins_service.CoinsService.GetById:output_type -> coins_service.Coin
	6, // 11: coins_service.CoinsService.GetList:output_type -> coins_service.GetListCoinResponse
	1, // 12: coins_service.CoinsService.Update:output_type -> coins_service.Coin
	0, // 13: coins_service.CoinsService.Delete:output_type -> coins_service.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_coin_proto_init() }
func file_coin_proto_init() {
	if File_coin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_coin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Coin); i {
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
		file_coin_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CoinPrimaryKey); i {
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
		file_coin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCoin); i {
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
		file_coin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCoin); i {
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
		file_coin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetListCoinRequest); i {
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
		file_coin_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetListCoinResponse); i {
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
		file_coin_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*HalfCoinPrice); i {
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
			RawDescriptor: file_coin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coin_proto_goTypes,
		DependencyIndexes: file_coin_proto_depIdxs,
		MessageInfos:      file_coin_proto_msgTypes,
	}.Build()
	File_coin_proto = out.File
	file_coin_proto_rawDesc = nil
	file_coin_proto_goTypes = nil
	file_coin_proto_depIdxs = nil
}

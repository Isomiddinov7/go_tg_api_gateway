// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: coin_nft.proto

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

type CoinNFT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NftImg     string `protobuf:"bytes,2,opt,name=nft_img,json=nftImg,proto3" json:"nft_img,omitempty"`
	NftPrice   string `protobuf:"bytes,3,opt,name=nft_price,json=nftPrice,proto3" json:"nft_price,omitempty"`
	NftAddress string `protobuf:"bytes,4,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	NftName    string `protobuf:"bytes,5,opt,name=nft_name,json=nftName,proto3" json:"nft_name,omitempty"`
	CreatedAt  string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CoinNFT) Reset() {
	*x = CoinNFT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_nft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinNFT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinNFT) ProtoMessage() {}

func (x *CoinNFT) ProtoReflect() protoreflect.Message {
	mi := &file_coin_nft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinNFT.ProtoReflect.Descriptor instead.
func (*CoinNFT) Descriptor() ([]byte, []int) {
	return file_coin_nft_proto_rawDescGZIP(), []int{0}
}

func (x *CoinNFT) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoinNFT) GetNftImg() string {
	if x != nil {
		return x.NftImg
	}
	return ""
}

func (x *CoinNFT) GetNftPrice() string {
	if x != nil {
		return x.NftPrice
	}
	return ""
}

func (x *CoinNFT) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *CoinNFT) GetNftName() string {
	if x != nil {
		return x.NftName
	}
	return ""
}

func (x *CoinNFT) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CoinNFT) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CoinNFTCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NftImg     string `protobuf:"bytes,2,opt,name=nft_img,json=nftImg,proto3" json:"nft_img,omitempty"`
	NftPrice   string `protobuf:"bytes,3,opt,name=nft_price,json=nftPrice,proto3" json:"nft_price,omitempty"`
	NftAddress string `protobuf:"bytes,4,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	NftName    string `protobuf:"bytes,5,opt,name=nft_name,json=nftName,proto3" json:"nft_name,omitempty"`
}

func (x *CoinNFTCreate) Reset() {
	*x = CoinNFTCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_nft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinNFTCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinNFTCreate) ProtoMessage() {}

func (x *CoinNFTCreate) ProtoReflect() protoreflect.Message {
	mi := &file_coin_nft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinNFTCreate.ProtoReflect.Descriptor instead.
func (*CoinNFTCreate) Descriptor() ([]byte, []int) {
	return file_coin_nft_proto_rawDescGZIP(), []int{1}
}

func (x *CoinNFTCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoinNFTCreate) GetNftImg() string {
	if x != nil {
		return x.NftImg
	}
	return ""
}

func (x *CoinNFTCreate) GetNftPrice() string {
	if x != nil {
		return x.NftPrice
	}
	return ""
}

func (x *CoinNFTCreate) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *CoinNFTCreate) GetNftName() string {
	if x != nil {
		return x.NftName
	}
	return ""
}

type CoinNFTPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CoinNFTPrimaryKey) Reset() {
	*x = CoinNFTPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_nft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinNFTPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinNFTPrimaryKey) ProtoMessage() {}

func (x *CoinNFTPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_coin_nft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinNFTPrimaryKey.ProtoReflect.Descriptor instead.
func (*CoinNFTPrimaryKey) Descriptor() ([]byte, []int) {
	return file_coin_nft_proto_rawDescGZIP(), []int{2}
}

func (x *CoinNFTPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CoinNFTUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NftImg     string `protobuf:"bytes,2,opt,name=nft_img,json=nftImg,proto3" json:"nft_img,omitempty"`
	NftPrice   string `protobuf:"bytes,3,opt,name=nft_price,json=nftPrice,proto3" json:"nft_price,omitempty"`
	NftAddress string `protobuf:"bytes,4,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	NftName    string `protobuf:"bytes,5,opt,name=nft_name,json=nftName,proto3" json:"nft_name,omitempty"`
}

func (x *CoinNFTUpdate) Reset() {
	*x = CoinNFTUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_nft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinNFTUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinNFTUpdate) ProtoMessage() {}

func (x *CoinNFTUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_coin_nft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinNFTUpdate.ProtoReflect.Descriptor instead.
func (*CoinNFTUpdate) Descriptor() ([]byte, []int) {
	return file_coin_nft_proto_rawDescGZIP(), []int{3}
}

func (x *CoinNFTUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoinNFTUpdate) GetNftImg() string {
	if x != nil {
		return x.NftImg
	}
	return ""
}

func (x *CoinNFTUpdate) GetNftPrice() string {
	if x != nil {
		return x.NftPrice
	}
	return ""
}

func (x *CoinNFTUpdate) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *CoinNFTUpdate) GetNftName() string {
	if x != nil {
		return x.NftName
	}
	return ""
}

type GetListCoinNFTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListCoinNFTRequest) Reset() {
	*x = GetListCoinNFTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_nft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListCoinNFTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCoinNFTRequest) ProtoMessage() {}

func (x *GetListCoinNFTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coin_nft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCoinNFTRequest.ProtoReflect.Descriptor instead.
func (*GetListCoinNFTRequest) Descriptor() ([]byte, []int) {
	return file_coin_nft_proto_rawDescGZIP(), []int{4}
}

func (x *GetListCoinNFTRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListCoinNFTRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListCoinNFTRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListCoinNFTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	CoinNfts []*CoinNFT `protobuf:"bytes,2,rep,name=coin_nfts,json=coinNfts,proto3" json:"coin_nfts,omitempty"`
}

func (x *GetListCoinNFTResponse) Reset() {
	*x = GetListCoinNFTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coin_nft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListCoinNFTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCoinNFTResponse) ProtoMessage() {}

func (x *GetListCoinNFTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coin_nft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCoinNFTResponse.ProtoReflect.Descriptor instead.
func (*GetListCoinNFTResponse) Descriptor() ([]byte, []int) {
	return file_coin_nft_proto_rawDescGZIP(), []int{5}
}

func (x *GetListCoinNFTResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListCoinNFTResponse) GetCoinNfts() []*CoinNFT {
	if x != nil {
		return x.CoinNfts
	}
	return nil
}

var File_coin_nft_proto protoreflect.FileDescriptor

var file_coin_nft_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x07,
	0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x66, 0x74, 0x5f, 0x69,
	0x6d, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x66, 0x74, 0x49, 0x6d, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x66, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x66, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6e, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x69, 0x6e,
	0x4e, 0x46, 0x54, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x66, 0x74,
	0x5f, 0x69, 0x6d, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x66, 0x74, 0x49,
	0x6d, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x66, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x66, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x6e, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x43,
	0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x91, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6d, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x66, 0x74, 0x49, 0x6d, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x66, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x66, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x66, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x66, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x66, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x22, 0x63, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x4e, 0x46, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x66, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x52, 0x08,
	0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x66, 0x74, 0x73, 0x32, 0xf9, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x69,
	0x6e, 0x4e, 0x46, 0x54, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e,
	0x46, 0x54, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x4e, 0x46, 0x54, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x69, 0x6e, 0x4e, 0x46, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4e,
	0x46, 0x54, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coin_nft_proto_rawDescOnce sync.Once
	file_coin_nft_proto_rawDescData = file_coin_nft_proto_rawDesc
)

func file_coin_nft_proto_rawDescGZIP() []byte {
	file_coin_nft_proto_rawDescOnce.Do(func() {
		file_coin_nft_proto_rawDescData = protoimpl.X.CompressGZIP(file_coin_nft_proto_rawDescData)
	})
	return file_coin_nft_proto_rawDescData
}

var file_coin_nft_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coin_nft_proto_goTypes = []any{
	(*CoinNFT)(nil),                // 0: coins_service.CoinNFT
	(*CoinNFTCreate)(nil),          // 1: coins_service.CoinNFTCreate
	(*CoinNFTPrimaryKey)(nil),      // 2: coins_service.CoinNFTPrimaryKey
	(*CoinNFTUpdate)(nil),          // 3: coins_service.CoinNFTUpdate
	(*GetListCoinNFTRequest)(nil),  // 4: coins_service.GetListCoinNFTRequest
	(*GetListCoinNFTResponse)(nil), // 5: coins_service.GetListCoinNFTResponse
	(*Empty)(nil),                  // 6: coins_service.Empty
}
var file_coin_nft_proto_depIdxs = []int32{
	0, // 0: coins_service.GetListCoinNFTResponse.coin_nfts:type_name -> coins_service.CoinNFT
	1, // 1: coins_service.CoinNFTService.Create:input_type -> coins_service.CoinNFTCreate
	2, // 2: coins_service.CoinNFTService.GetById:input_type -> coins_service.CoinNFTPrimaryKey
	3, // 3: coins_service.CoinNFTService.Update:input_type -> coins_service.CoinNFTUpdate
	4, // 4: coins_service.CoinNFTService.GetList:input_type -> coins_service.GetListCoinNFTRequest
	2, // 5: coins_service.CoinNFTService.Delete:input_type -> coins_service.CoinNFTPrimaryKey
	0, // 6: coins_service.CoinNFTService.Create:output_type -> coins_service.CoinNFT
	0, // 7: coins_service.CoinNFTService.GetById:output_type -> coins_service.CoinNFT
	0, // 8: coins_service.CoinNFTService.Update:output_type -> coins_service.CoinNFT
	5, // 9: coins_service.CoinNFTService.GetList:output_type -> coins_service.GetListCoinNFTResponse
	6, // 10: coins_service.CoinNFTService.Delete:output_type -> coins_service.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coin_nft_proto_init() }
func file_coin_nft_proto_init() {
	if File_coin_nft_proto != nil {
		return
	}
	file_coin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coin_nft_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CoinNFT); i {
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
		file_coin_nft_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CoinNFTCreate); i {
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
		file_coin_nft_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CoinNFTPrimaryKey); i {
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
		file_coin_nft_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CoinNFTUpdate); i {
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
		file_coin_nft_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetListCoinNFTRequest); i {
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
		file_coin_nft_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetListCoinNFTResponse); i {
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
			RawDescriptor: file_coin_nft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coin_nft_proto_goTypes,
		DependencyIndexes: file_coin_nft_proto_depIdxs,
		MessageInfos:      file_coin_nft_proto_msgTypes,
	}.Build()
	File_coin_nft_proto = out.File
	file_coin_nft_proto_rawDesc = nil
	file_coin_nft_proto_goTypes = nil
	file_coin_nft_proto_depIdxs = nil
}

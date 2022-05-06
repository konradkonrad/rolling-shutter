// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: gossip.proto

package shmsg

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

type DecryptionTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID       uint64 `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	EpochID          uint64 `protobuf:"varint,2,opt,name=epochID,proto3" json:"epochID,omitempty"`
	TransactionsHash []byte `protobuf:"bytes,3,opt,name=transactionsHash,proto3" json:"transactionsHash,omitempty"`
	Signature        []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DecryptionTrigger) Reset() {
	*x = DecryptionTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptionTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptionTrigger) ProtoMessage() {}

func (x *DecryptionTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptionTrigger.ProtoReflect.Descriptor instead.
func (*DecryptionTrigger) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{0}
}

func (x *DecryptionTrigger) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *DecryptionTrigger) GetEpochID() uint64 {
	if x != nil {
		return x.EpochID
	}
	return 0
}

func (x *DecryptionTrigger) GetTransactionsHash() []byte {
	if x != nil {
		return x.TransactionsHash
	}
	return nil
}

func (x *DecryptionTrigger) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// TODO: replace keyper index by signature
type DecryptionKeyShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID  uint64 `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	EpochID     uint64 `protobuf:"varint,2,opt,name=epochID,proto3" json:"epochID,omitempty"`
	KeyperIndex uint64 `protobuf:"varint,3,opt,name=keyperIndex,proto3" json:"keyperIndex,omitempty"`
	Share       []byte `protobuf:"bytes,4,opt,name=share,proto3" json:"share,omitempty"`
}

func (x *DecryptionKeyShare) Reset() {
	*x = DecryptionKeyShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptionKeyShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptionKeyShare) ProtoMessage() {}

func (x *DecryptionKeyShare) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptionKeyShare.ProtoReflect.Descriptor instead.
func (*DecryptionKeyShare) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *DecryptionKeyShare) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *DecryptionKeyShare) GetEpochID() uint64 {
	if x != nil {
		return x.EpochID
	}
	return 0
}

func (x *DecryptionKeyShare) GetKeyperIndex() uint64 {
	if x != nil {
		return x.KeyperIndex
	}
	return 0
}

func (x *DecryptionKeyShare) GetShare() []byte {
	if x != nil {
		return x.Share
	}
	return nil
}

type DecryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID uint64 `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	EpochID    uint64 `protobuf:"varint,2,opt,name=epochID,proto3" json:"epochID,omitempty"`
	Key        []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DecryptionKey) Reset() {
	*x = DecryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptionKey) ProtoMessage() {}

func (x *DecryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptionKey.ProtoReflect.Descriptor instead.
func (*DecryptionKey) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{2}
}

func (x *DecryptionKey) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *DecryptionKey) GetEpochID() uint64 {
	if x != nil {
		return x.EpochID
	}
	return 0
}

func (x *DecryptionKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// EonPublicKeyCandidate is sent as part of the EonPublicKey message by the
// keypers to publish the EonPublicKey for a certain eon. From the keypers point
// of view this isn't a candidate, it's the real key for the given eon. But for
// those that observe it, e.g. the collator, it's a candidate until the observer
// has seen at least threshold messages.
type EonPublicKeyCandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID        uint64 `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	PublicKey         []byte `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	ActivationBlock   uint64 `protobuf:"varint,3,opt,name=activationBlock,proto3" json:"activationBlock,omitempty"`
	KeyperConfigIndex uint64 `protobuf:"varint,6,opt,name=keyperConfigIndex,proto3" json:"keyperConfigIndex,omitempty"`
	Eon               uint64 `protobuf:"varint,7,opt,name=eon,proto3" json:"eon,omitempty"`
}

func (x *EonPublicKeyCandidate) Reset() {
	*x = EonPublicKeyCandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EonPublicKeyCandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EonPublicKeyCandidate) ProtoMessage() {}

func (x *EonPublicKeyCandidate) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EonPublicKeyCandidate.ProtoReflect.Descriptor instead.
func (*EonPublicKeyCandidate) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{3}
}

func (x *EonPublicKeyCandidate) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *EonPublicKeyCandidate) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *EonPublicKeyCandidate) GetActivationBlock() uint64 {
	if x != nil {
		return x.ActivationBlock
	}
	return 0
}

func (x *EonPublicKeyCandidate) GetKeyperConfigIndex() uint64 {
	if x != nil {
		return x.KeyperConfigIndex
	}
	return 0
}

func (x *EonPublicKeyCandidate) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

// EonPublicKey is sent by the keypers to publish the EonPublicKey for a certain eon.
type EonPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidate   *EonPublicKeyCandidate `protobuf:"bytes,7,opt,name=candidate,proto3" json:"candidate,omitempty"`
	KeyperIndex uint64                 `protobuf:"varint,4,opt,name=keyperIndex,proto3" json:"keyperIndex,omitempty"`
	Signature   []byte                 `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *EonPublicKey) Reset() {
	*x = EonPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EonPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EonPublicKey) ProtoMessage() {}

func (x *EonPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EonPublicKey.ProtoReflect.Descriptor instead.
func (*EonPublicKey) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{4}
}

func (x *EonPublicKey) GetCandidate() *EonPublicKeyCandidate {
	if x != nil {
		return x.Candidate
	}
	return nil
}

func (x *EonPublicKey) GetKeyperIndex() uint64 {
	if x != nil {
		return x.KeyperIndex
	}
	return 0
}

func (x *EonPublicKey) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x73, 0x68, 0x6d, 0x73, 0x67, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x86, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x5b, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x15, 0x45, 0x6f, 0x6e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x6b, 0x65, 0x79, 0x70, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x11, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x45, 0x6f, 0x6e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x68,
	0x6d, 0x73, 0x67, 0x2e, 0x45, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x70, 0x65,
	0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x73, 0x68, 0x6d, 0x73, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gossip_proto_rawDescOnce sync.Once
	file_gossip_proto_rawDescData = file_gossip_proto_rawDesc
)

func file_gossip_proto_rawDescGZIP() []byte {
	file_gossip_proto_rawDescOnce.Do(func() {
		file_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_gossip_proto_rawDescData)
	})
	return file_gossip_proto_rawDescData
}

var file_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gossip_proto_goTypes = []interface{}{
	(*DecryptionTrigger)(nil),     // 0: shmsg.DecryptionTrigger
	(*DecryptionKeyShare)(nil),    // 1: shmsg.DecryptionKeyShare
	(*DecryptionKey)(nil),         // 2: shmsg.DecryptionKey
	(*EonPublicKeyCandidate)(nil), // 3: shmsg.EonPublicKeyCandidate
	(*EonPublicKey)(nil),          // 4: shmsg.EonPublicKey
}
var file_gossip_proto_depIdxs = []int32{
	3, // 0: shmsg.EonPublicKey.candidate:type_name -> shmsg.EonPublicKeyCandidate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gossip_proto_init() }
func file_gossip_proto_init() {
	if File_gossip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gossip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptionTrigger); i {
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
		file_gossip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptionKeyShare); i {
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
		file_gossip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptionKey); i {
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
		file_gossip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EonPublicKeyCandidate); i {
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
		file_gossip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EonPublicKey); i {
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
			RawDescriptor: file_gossip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gossip_proto_goTypes,
		DependencyIndexes: file_gossip_proto_depIdxs,
		MessageInfos:      file_gossip_proto_msgTypes,
	}.Build()
	File_gossip_proto = out.File
	file_gossip_proto_rawDesc = nil
	file_gossip_proto_goTypes = nil
	file_gossip_proto_depIdxs = nil
}

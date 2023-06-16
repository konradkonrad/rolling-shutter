// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: gossip.proto

package p2pmsg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
	EpochID          []byte `protobuf:"bytes,2,opt,name=epochID,proto3" json:"epochID,omitempty"`
	BlockNumber      uint64 `protobuf:"varint,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TransactionsHash []byte `protobuf:"bytes,4,opt,name=transactionsHash,proto3" json:"transactionsHash,omitempty"`
	Signature        []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
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

func (x *DecryptionTrigger) GetEpochID() []byte {
	if x != nil {
		return x.EpochID
	}
	return nil
}

func (x *DecryptionTrigger) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
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

type KeyShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochID []byte `protobuf:"bytes,1,opt,name=epochID,proto3" json:"epochID,omitempty"`
	Share   []byte `protobuf:"bytes,2,opt,name=share,proto3" json:"share,omitempty"`
}

func (x *KeyShare) Reset() {
	*x = KeyShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyShare) ProtoMessage() {}

func (x *KeyShare) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use KeyShare.ProtoReflect.Descriptor instead.
func (*KeyShare) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *KeyShare) GetEpochID() []byte {
	if x != nil {
		return x.EpochID
	}
	return nil
}

func (x *KeyShare) GetShare() []byte {
	if x != nil {
		return x.Share
	}
	return nil
}

// TODO: replace keyper index by signature
type DecryptionKeyShares struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID  uint64      `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	Eon         uint64      `protobuf:"varint,4,opt,name=eon,proto3" json:"eon,omitempty"`
	KeyperIndex uint64      `protobuf:"varint,5,opt,name=keyperIndex,proto3" json:"keyperIndex,omitempty"`
	Shares      []*KeyShare `protobuf:"bytes,9,rep,name=shares,proto3" json:"shares,omitempty"`
}

func (x *DecryptionKeyShares) Reset() {
	*x = DecryptionKeyShares{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptionKeyShares) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptionKeyShares) ProtoMessage() {}

func (x *DecryptionKeyShares) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DecryptionKeyShares.ProtoReflect.Descriptor instead.
func (*DecryptionKeyShares) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{2}
}

func (x *DecryptionKeyShares) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *DecryptionKeyShares) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *DecryptionKeyShares) GetKeyperIndex() uint64 {
	if x != nil {
		return x.KeyperIndex
	}
	return 0
}

func (x *DecryptionKeyShares) GetShares() []*KeyShare {
	if x != nil {
		return x.Shares
	}
	return nil
}

type DecryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID uint64 `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	Eon        uint64 `protobuf:"varint,2,opt,name=eon,proto3" json:"eon,omitempty"`
	EpochID    []byte `protobuf:"bytes,3,opt,name=epochID,proto3" json:"epochID,omitempty"`
	Key        []byte `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DecryptionKey) Reset() {
	*x = DecryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptionKey) ProtoMessage() {}

func (x *DecryptionKey) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DecryptionKey.ProtoReflect.Descriptor instead.
func (*DecryptionKey) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{3}
}

func (x *DecryptionKey) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *DecryptionKey) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *DecryptionKey) GetEpochID() []byte {
	if x != nil {
		return x.EpochID
	}
	return nil
}

func (x *DecryptionKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// EonPublicKey is sent by the keypers to publish the EonPublicKey for a certain
// eon.  For those that observe it, e.g. the collator, it's a candidate until
// the observer has seen at least threshold messages.
type EonPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceID        uint64 `protobuf:"varint,1,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	PublicKey         []byte `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	ActivationBlock   uint64 `protobuf:"varint,3,opt,name=activationBlock,proto3" json:"activationBlock,omitempty"`
	KeyperConfigIndex uint64 `protobuf:"varint,6,opt,name=keyperConfigIndex,proto3" json:"keyperConfigIndex,omitempty"`
	Eon               uint64 `protobuf:"varint,7,opt,name=eon,proto3" json:"eon,omitempty"`
	Signature         []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
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

func (x *EonPublicKey) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *EonPublicKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *EonPublicKey) GetActivationBlock() uint64 {
	if x != nil {
		return x.ActivationBlock
	}
	return 0
}

func (x *EonPublicKey) GetKeyperConfigIndex() uint64 {
	if x != nil {
		return x.KeyperConfigIndex
	}
	return 0
}

func (x *EonPublicKey) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *EonPublicKey) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TraceContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceID    []byte `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`
	SpanID     []byte `protobuf:"bytes,2,opt,name=spanID,proto3" json:"spanID,omitempty"`
	TraceFlags []byte `protobuf:"bytes,3,opt,name=traceFlags,proto3" json:"traceFlags,omitempty"`
	TraceState string `protobuf:"bytes,4,opt,name=traceState,proto3" json:"traceState,omitempty"`
}

func (x *TraceContext) Reset() {
	*x = TraceContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceContext) ProtoMessage() {}

func (x *TraceContext) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceContext.ProtoReflect.Descriptor instead.
func (*TraceContext) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{5}
}

func (x *TraceContext) GetTraceID() []byte {
	if x != nil {
		return x.TraceID
	}
	return nil
}

func (x *TraceContext) GetSpanID() []byte {
	if x != nil {
		return x.SpanID
	}
	return nil
}

func (x *TraceContext) GetTraceFlags() []byte {
	if x != nil {
		return x.TraceFlags
	}
	return nil
}

func (x *TraceContext) GetTraceState() string {
	if x != nil {
		return x.TraceState
	}
	return ""
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string        `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Message *anypb.Any    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Trace   *TraceContext `protobuf:"bytes,3,opt,name=trace,proto3,oneof" json:"trace,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{6}
}

func (x *Envelope) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Envelope) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Envelope) GetTrace() *TraceContext {
	if x != nil {
		return x.Trace
	}
	return nil
}

var File_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x32, 0x70, 0x6d, 0x73, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3a, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x65, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x32, 0x70, 0x6d, 0x73, 0x67, 0x2e, 0x4b,
	0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x22,
	0x6d, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xd4,
	0x01, 0x0a, 0x0c, 0x45, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
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
	0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x32, 0x70, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x3b, 0x70, 0x32, 0x70, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gossip_proto_goTypes = []interface{}{
	(*DecryptionTrigger)(nil),   // 0: p2pmsg.DecryptionTrigger
	(*KeyShare)(nil),            // 1: p2pmsg.KeyShare
	(*DecryptionKeyShares)(nil), // 2: p2pmsg.DecryptionKeyShares
	(*DecryptionKey)(nil),       // 3: p2pmsg.DecryptionKey
	(*EonPublicKey)(nil),        // 4: p2pmsg.EonPublicKey
	(*TraceContext)(nil),        // 5: p2pmsg.TraceContext
	(*Envelope)(nil),            // 6: p2pmsg.Envelope
	(*anypb.Any)(nil),           // 7: google.protobuf.Any
}
var file_gossip_proto_depIdxs = []int32{
	1, // 0: p2pmsg.DecryptionKeyShares.shares:type_name -> p2pmsg.KeyShare
	7, // 1: p2pmsg.Envelope.message:type_name -> google.protobuf.Any
	5, // 2: p2pmsg.Envelope.trace:type_name -> p2pmsg.TraceContext
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
			switch v := v.(*KeyShare); i {
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
			switch v := v.(*DecryptionKeyShares); i {
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
		file_gossip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceContext); i {
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
		file_gossip_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
	file_gossip_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gossip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/relations/v1beta1/common.proto

package v1beta1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// A _Relationship_ is the realization of a _Relation_ (a string)
// between a _Resource_ and a _Subject_ or a _Subject Set_ (known as a Userset in Zanzibar).
//
// All Relationships are object-object relations.
// "Resource" and "Subject" are relative terms which define the direction of a Relation.
// That is, Relations are unidirectional.
// If you reverse the Subject and Resource, it is a different Relation and a different Relationship.
// Conventionally, we generally refer to the Resource first, then Subject,
// following the direction of typical graph traversal (Resource to Subject).
type Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *ObjectReference  `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Relation string            `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Subject  *SubjectReference `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_kessel_relations_v1beta1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Relationship) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Relationship) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *Relationship) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

// A reference to a Subject or, if a `relation` is provided, a Subject Set.
type SubjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional relation which points to a set of Subjects instead of the single Subject.
	// e.g. "members" or "owners" of a group identified in `subject`.
	Relation *string          `protobuf:"bytes,1,opt,name=relation,proto3,oneof" json:"relation,omitempty"`
	Subject  *ObjectReference `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *SubjectReference) Reset() {
	*x = SubjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectReference) ProtoMessage() {}

func (x *SubjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectReference.ProtoReflect.Descriptor instead.
func (*SubjectReference) Descriptor() ([]byte, []int) {
	return file_kessel_relations_v1beta1_common_proto_rawDescGZIP(), []int{1}
}

func (x *SubjectReference) GetRelation() string {
	if x != nil && x.Relation != nil {
		return *x.Relation
	}
	return ""
}

func (x *SubjectReference) GetSubject() *ObjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

type RequestPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit             uint32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	ContinuationToken *string `protobuf:"bytes,2,opt,name=continuation_token,json=continuationToken,proto3,oneof" json:"continuation_token,omitempty"`
}

func (x *RequestPagination) Reset() {
	*x = RequestPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPagination) ProtoMessage() {}

func (x *RequestPagination) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPagination.ProtoReflect.Descriptor instead.
func (*RequestPagination) Descriptor() ([]byte, []int) {
	return file_kessel_relations_v1beta1_common_proto_rawDescGZIP(), []int{2}
}

func (x *RequestPagination) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RequestPagination) GetContinuationToken() string {
	if x != nil && x.ContinuationToken != nil {
		return *x.ContinuationToken
	}
	return ""
}

type ResponsePagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContinuationToken string `protobuf:"bytes,1,opt,name=continuation_token,json=continuationToken,proto3" json:"continuation_token,omitempty"`
}

func (x *ResponsePagination) Reset() {
	*x = ResponsePagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePagination) ProtoMessage() {}

func (x *ResponsePagination) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePagination.ProtoReflect.Descriptor instead.
func (*ResponsePagination) Descriptor() ([]byte, []int) {
	return file_kessel_relations_v1beta1_common_proto_rawDescGZIP(), []int{3}
}

func (x *ResponsePagination) GetContinuationToken() string {
	if x != nil {
		return x.ContinuationToken
	}
	return ""
}

type ObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *ObjectType `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id   string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ObjectReference) Reset() {
	*x = ObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectReference) ProtoMessage() {}

func (x *ObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectReference.ProtoReflect.Descriptor instead.
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return file_kessel_relations_v1beta1_common_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectReference) GetType() *ObjectType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ObjectReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ObjectType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ObjectType) Reset() {
	*x = ObjectType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectType) ProtoMessage() {}

func (x *ObjectType) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_relations_v1beta1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectType.ProtoReflect.Descriptor instead.
func (*ObjectType) Descriptor() ([]byte, []int) {
	return file_kessel_relations_v1beta1_common_proto_rawDescGZIP(), []int{5}
}

func (x *ObjectType) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_kessel_relations_v1beta1_common_proto protoreflect.FileDescriptor

var file_kessel_relations_v1beta1_common_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0,
	0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x4d, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7d, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x43, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6c, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x72, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kessel_relations_v1beta1_common_proto_rawDescOnce sync.Once
	file_kessel_relations_v1beta1_common_proto_rawDescData = file_kessel_relations_v1beta1_common_proto_rawDesc
)

func file_kessel_relations_v1beta1_common_proto_rawDescGZIP() []byte {
	file_kessel_relations_v1beta1_common_proto_rawDescOnce.Do(func() {
		file_kessel_relations_v1beta1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_relations_v1beta1_common_proto_rawDescData)
	})
	return file_kessel_relations_v1beta1_common_proto_rawDescData
}

var file_kessel_relations_v1beta1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kessel_relations_v1beta1_common_proto_goTypes = []any{
	(*Relationship)(nil),       // 0: kessel.relations.v1beta1.Relationship
	(*SubjectReference)(nil),   // 1: kessel.relations.v1beta1.SubjectReference
	(*RequestPagination)(nil),  // 2: kessel.relations.v1beta1.RequestPagination
	(*ResponsePagination)(nil), // 3: kessel.relations.v1beta1.ResponsePagination
	(*ObjectReference)(nil),    // 4: kessel.relations.v1beta1.ObjectReference
	(*ObjectType)(nil),         // 5: kessel.relations.v1beta1.ObjectType
}
var file_kessel_relations_v1beta1_common_proto_depIdxs = []int32{
	4, // 0: kessel.relations.v1beta1.Relationship.resource:type_name -> kessel.relations.v1beta1.ObjectReference
	1, // 1: kessel.relations.v1beta1.Relationship.subject:type_name -> kessel.relations.v1beta1.SubjectReference
	4, // 2: kessel.relations.v1beta1.SubjectReference.subject:type_name -> kessel.relations.v1beta1.ObjectReference
	5, // 3: kessel.relations.v1beta1.ObjectReference.type:type_name -> kessel.relations.v1beta1.ObjectType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kessel_relations_v1beta1_common_proto_init() }
func file_kessel_relations_v1beta1_common_proto_init() {
	if File_kessel_relations_v1beta1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_relations_v1beta1_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Relationship); i {
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
		file_kessel_relations_v1beta1_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SubjectReference); i {
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
		file_kessel_relations_v1beta1_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RequestPagination); i {
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
		file_kessel_relations_v1beta1_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ResponsePagination); i {
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
		file_kessel_relations_v1beta1_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectReference); i {
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
		file_kessel_relations_v1beta1_common_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectType); i {
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
	file_kessel_relations_v1beta1_common_proto_msgTypes[1].OneofWrappers = []any{}
	file_kessel_relations_v1beta1_common_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kessel_relations_v1beta1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_relations_v1beta1_common_proto_goTypes,
		DependencyIndexes: file_kessel_relations_v1beta1_common_proto_depIdxs,
		MessageInfos:      file_kessel_relations_v1beta1_common_proto_msgTypes,
	}.Build()
	File_kessel_relations_v1beta1_common_proto = out.File
	file_kessel_relations_v1beta1_common_proto_rawDesc = nil
	file_kessel_relations_v1beta1_common_proto_goTypes = nil
	file_kessel_relations_v1beta1_common_proto_depIdxs = nil
}

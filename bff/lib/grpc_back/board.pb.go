// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: board.proto

package grpc_back

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	UserId               string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description          string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DepartureLatitude    float32                `protobuf:"fixed32,5,opt,name=departure_latitude,json=departureLatitude,proto3" json:"departure_latitude,omitempty"`
	DepartureLongitude   float32                `protobuf:"fixed32,6,opt,name=departure_longitude,json=departureLongitude,proto3" json:"departure_longitude,omitempty"`
	DestinationLatitude  float32                `protobuf:"fixed32,7,opt,name=destination_latitude,json=destinationLatitude,proto3" json:"destination_latitude,omitempty"`
	DestinationLongitude float32                `protobuf:"fixed32,8,opt,name=destination_longitude,json=destinationLongitude,proto3" json:"destination_longitude,omitempty"`
	Reward               string                 `protobuf:"bytes,9,opt,name=reward,proto3" json:"reward,omitempty"`
	StartTime            *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{0}
}

func (x *Board) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Board) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Board) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Board) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Board) GetDepartureLatitude() float32 {
	if x != nil {
		return x.DepartureLatitude
	}
	return 0
}

func (x *Board) GetDepartureLongitude() float32 {
	if x != nil {
		return x.DepartureLongitude
	}
	return 0
}

func (x *Board) GetDestinationLatitude() float32 {
	if x != nil {
		return x.DestinationLatitude
	}
	return 0
}

func (x *Board) GetDestinationLongitude() float32 {
	if x != nil {
		return x.DestinationLongitude
	}
	return 0
}

func (x *Board) GetReward() string {
	if x != nil {
		return x.Reward
	}
	return ""
}

func (x *Board) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Board) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Board) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type BoardList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boards []*Board `protobuf:"bytes,1,rep,name=boards,proto3" json:"boards,omitempty"`
}

func (x *BoardList) Reset() {
	*x = BoardList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardList) ProtoMessage() {}

func (x *BoardList) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardList.ProtoReflect.Descriptor instead.
func (*BoardList) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{1}
}

func (x *BoardList) GetBoards() []*Board {
	if x != nil {
		return x.Boards
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId      string                 `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CommenterId string                 `protobuf:"bytes,3,opt,name=commenter_id,json=commenterId,proto3" json:"commenter_id,omitempty"`
	Content     string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Comment) GetCommenterId() string {
	if x != nil {
		return x.CommenterId
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CommentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *CommentList) Reset() {
	*x = CommentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentList) ProtoMessage() {}

func (x *CommentList) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentList.ProtoReflect.Descriptor instead.
func (*CommentList) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{3}
}

func (x *CommentList) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type BoardID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BoardID) Reset() {
	*x = BoardID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardID) ProtoMessage() {}

func (x *BoardID) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardID.ProtoReflect.Descriptor instead.
func (*BoardID) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{4}
}

func (x *BoardID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_board_proto protoreflect.FileDescriptor

var file_board_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf7, 0x03, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x33,
	0x0a, 0x15, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x31, 0x0a, 0x09,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x22,
	0xe5, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x88, 0x04,
	0x0a, 0x0c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0c, 0x2e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x0c, 0x2e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0c, 0x2e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x0c, 0x2e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x2f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44,
	0x1a, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x37, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_board_proto_rawDescOnce sync.Once
	file_board_proto_rawDescData = file_board_proto_rawDesc
)

func file_board_proto_rawDescGZIP() []byte {
	file_board_proto_rawDescOnce.Do(func() {
		file_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_proto_rawDescData)
	})
	return file_board_proto_rawDescData
}

var file_board_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_board_proto_goTypes = []interface{}{
	(*Board)(nil),                 // 0: board.Board
	(*BoardList)(nil),             // 1: board.BoardList
	(*Comment)(nil),               // 2: board.Comment
	(*CommentList)(nil),           // 3: board.CommentList
	(*BoardID)(nil),               // 4: board.BoardID
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_board_proto_depIdxs = []int32{
	5,  // 0: board.Board.start_time:type_name -> google.protobuf.Timestamp
	5,  // 1: board.Board.created_at:type_name -> google.protobuf.Timestamp
	5,  // 2: board.Board.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 3: board.BoardList.boards:type_name -> board.Board
	5,  // 4: board.Comment.created_at:type_name -> google.protobuf.Timestamp
	5,  // 5: board.Comment.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 6: board.CommentList.comments:type_name -> board.Comment
	0,  // 7: board.BoardService.CreateBoard:input_type -> board.Board
	4,  // 8: board.BoardService.ReadBoard:input_type -> board.BoardID
	6,  // 9: board.BoardService.ReadAllBoard:input_type -> google.protobuf.Empty
	0,  // 10: board.BoardService.UpdateBoard:input_type -> board.Board
	4,  // 11: board.BoardService.DeleteBoard:input_type -> board.BoardID
	2,  // 12: board.BoardService.CreateComment:input_type -> board.Comment
	4,  // 13: board.BoardService.ReadComment:input_type -> board.BoardID
	6,  // 14: board.BoardService.ReadAllComment:input_type -> google.protobuf.Empty
	2,  // 15: board.BoardService.UpdateComment:input_type -> board.Comment
	4,  // 16: board.BoardService.DeleteComment:input_type -> board.BoardID
	0,  // 17: board.BoardService.CreateBoard:output_type -> board.Board
	0,  // 18: board.BoardService.ReadBoard:output_type -> board.Board
	1,  // 19: board.BoardService.ReadAllBoard:output_type -> board.BoardList
	0,  // 20: board.BoardService.UpdateBoard:output_type -> board.Board
	6,  // 21: board.BoardService.DeleteBoard:output_type -> google.protobuf.Empty
	2,  // 22: board.BoardService.CreateComment:output_type -> board.Comment
	2,  // 23: board.BoardService.ReadComment:output_type -> board.Comment
	3,  // 24: board.BoardService.ReadAllComment:output_type -> board.CommentList
	2,  // 25: board.BoardService.UpdateComment:output_type -> board.Comment
	6,  // 26: board.BoardService.DeleteComment:output_type -> google.protobuf.Empty
	17, // [17:27] is the sub-list for method output_type
	7,  // [7:17] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_board_proto_init() }
func file_board_proto_init() {
	if File_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardList); i {
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
		file_board_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_board_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentList); i {
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
		file_board_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardID); i {
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
			RawDescriptor: file_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_board_proto_goTypes,
		DependencyIndexes: file_board_proto_depIdxs,
		MessageInfos:      file_board_proto_msgTypes,
	}.Build()
	File_board_proto = out.File
	file_board_proto_rawDesc = nil
	file_board_proto_goTypes = nil
	file_board_proto_depIdxs = nil
}

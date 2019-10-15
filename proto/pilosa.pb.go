// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pilosa.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryPQLRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Pql                  string   `protobuf:"bytes,2,opt,name=pql,proto3" json:"pql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryPQLRequest) Reset()         { *m = QueryPQLRequest{} }
func (m *QueryPQLRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPQLRequest) ProtoMessage()    {}
func (*QueryPQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{0}
}

func (m *QueryPQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPQLRequest.Unmarshal(m, b)
}
func (m *QueryPQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPQLRequest.Marshal(b, m, deterministic)
}
func (m *QueryPQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPQLRequest.Merge(m, src)
}
func (m *QueryPQLRequest) XXX_Size() int {
	return xxx_messageInfo_QueryPQLRequest.Size(m)
}
func (m *QueryPQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPQLRequest proto.InternalMessageInfo

func (m *QueryPQLRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *QueryPQLRequest) GetPql() string {
	if m != nil {
		return m.Pql
	}
	return ""
}

type RowResponse struct {
	Headers              []*ColumnInfo     `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Columns              []*ColumnResponse `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	ErrorMessage         string            `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RowResponse) Reset()         { *m = RowResponse{} }
func (m *RowResponse) String() string { return proto.CompactTextString(m) }
func (*RowResponse) ProtoMessage()    {}
func (*RowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{1}
}

func (m *RowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowResponse.Unmarshal(m, b)
}
func (m *RowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowResponse.Marshal(b, m, deterministic)
}
func (m *RowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowResponse.Merge(m, src)
}
func (m *RowResponse) XXX_Size() int {
	return xxx_messageInfo_RowResponse.Size(m)
}
func (m *RowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RowResponse proto.InternalMessageInfo

func (m *RowResponse) GetHeaders() []*ColumnInfo {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RowResponse) GetColumns() []*ColumnResponse {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *RowResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type ColumnInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Datatype             string   `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnInfo) Reset()         { *m = ColumnInfo{} }
func (m *ColumnInfo) String() string { return proto.CompactTextString(m) }
func (*ColumnInfo) ProtoMessage()    {}
func (*ColumnInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{2}
}

func (m *ColumnInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnInfo.Unmarshal(m, b)
}
func (m *ColumnInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnInfo.Marshal(b, m, deterministic)
}
func (m *ColumnInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnInfo.Merge(m, src)
}
func (m *ColumnInfo) XXX_Size() int {
	return xxx_messageInfo_ColumnInfo.Size(m)
}
func (m *ColumnInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnInfo proto.InternalMessageInfo

func (m *ColumnInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnInfo) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

type ColumnResponse struct {
	// Types that are valid to be assigned to ColumnVal:
	//	*ColumnResponse_StringVal
	//	*ColumnResponse_UintVal
	//	*ColumnResponse_IntVal
	//	*ColumnResponse_BoolVal
	//	*ColumnResponse_BlobVal
	ColumnVal            isColumnResponse_ColumnVal `protobuf_oneof:"columnVal"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ColumnResponse) Reset()         { *m = ColumnResponse{} }
func (m *ColumnResponse) String() string { return proto.CompactTextString(m) }
func (*ColumnResponse) ProtoMessage()    {}
func (*ColumnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{3}
}

func (m *ColumnResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnResponse.Unmarshal(m, b)
}
func (m *ColumnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnResponse.Marshal(b, m, deterministic)
}
func (m *ColumnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnResponse.Merge(m, src)
}
func (m *ColumnResponse) XXX_Size() int {
	return xxx_messageInfo_ColumnResponse.Size(m)
}
func (m *ColumnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnResponse proto.InternalMessageInfo

type isColumnResponse_ColumnVal interface {
	isColumnResponse_ColumnVal()
}

type ColumnResponse_StringVal struct {
	StringVal string `protobuf:"bytes,1,opt,name=stringVal,proto3,oneof"`
}

type ColumnResponse_UintVal struct {
	UintVal uint64 `protobuf:"varint,2,opt,name=uintVal,proto3,oneof"`
}

type ColumnResponse_IntVal struct {
	IntVal int64 `protobuf:"varint,3,opt,name=intVal,proto3,oneof"`
}

type ColumnResponse_BoolVal struct {
	BoolVal bool `protobuf:"varint,4,opt,name=boolVal,proto3,oneof"`
}

type ColumnResponse_BlobVal struct {
	BlobVal []byte `protobuf:"bytes,5,opt,name=blobVal,proto3,oneof"`
}

func (*ColumnResponse_StringVal) isColumnResponse_ColumnVal() {}

func (*ColumnResponse_UintVal) isColumnResponse_ColumnVal() {}

func (*ColumnResponse_IntVal) isColumnResponse_ColumnVal() {}

func (*ColumnResponse_BoolVal) isColumnResponse_ColumnVal() {}

func (*ColumnResponse_BlobVal) isColumnResponse_ColumnVal() {}

func (m *ColumnResponse) GetColumnVal() isColumnResponse_ColumnVal {
	if m != nil {
		return m.ColumnVal
	}
	return nil
}

func (m *ColumnResponse) GetStringVal() string {
	if x, ok := m.GetColumnVal().(*ColumnResponse_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (m *ColumnResponse) GetUintVal() uint64 {
	if x, ok := m.GetColumnVal().(*ColumnResponse_UintVal); ok {
		return x.UintVal
	}
	return 0
}

func (m *ColumnResponse) GetIntVal() int64 {
	if x, ok := m.GetColumnVal().(*ColumnResponse_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (m *ColumnResponse) GetBoolVal() bool {
	if x, ok := m.GetColumnVal().(*ColumnResponse_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (m *ColumnResponse) GetBlobVal() []byte {
	if x, ok := m.GetColumnVal().(*ColumnResponse_BlobVal); ok {
		return x.BlobVal
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ColumnResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ColumnResponse_OneofMarshaler, _ColumnResponse_OneofUnmarshaler, _ColumnResponse_OneofSizer, []interface{}{
		(*ColumnResponse_StringVal)(nil),
		(*ColumnResponse_UintVal)(nil),
		(*ColumnResponse_IntVal)(nil),
		(*ColumnResponse_BoolVal)(nil),
		(*ColumnResponse_BlobVal)(nil),
	}
}

func _ColumnResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ColumnResponse)
	// columnVal
	switch x := m.ColumnVal.(type) {
	case *ColumnResponse_StringVal:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringVal)
	case *ColumnResponse_UintVal:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.UintVal))
	case *ColumnResponse_IntVal:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntVal))
	case *ColumnResponse_BoolVal:
		t := uint64(0)
		if x.BoolVal {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ColumnResponse_BlobVal:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BlobVal)
	case nil:
	default:
		return fmt.Errorf("ColumnResponse.ColumnVal has unexpected type %T", x)
	}
	return nil
}

func _ColumnResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ColumnResponse)
	switch tag {
	case 1: // columnVal.stringVal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ColumnVal = &ColumnResponse_StringVal{x}
		return true, err
	case 2: // columnVal.uintVal
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ColumnVal = &ColumnResponse_UintVal{x}
		return true, err
	case 3: // columnVal.intVal
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ColumnVal = &ColumnResponse_IntVal{int64(x)}
		return true, err
	case 4: // columnVal.boolVal
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ColumnVal = &ColumnResponse_BoolVal{x != 0}
		return true, err
	case 5: // columnVal.blobVal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ColumnVal = &ColumnResponse_BlobVal{x}
		return true, err
	default:
		return false, nil
	}
}

func _ColumnResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ColumnResponse)
	// columnVal
	switch x := m.ColumnVal.(type) {
	case *ColumnResponse_StringVal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringVal)))
		n += len(x.StringVal)
	case *ColumnResponse_UintVal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.UintVal))
	case *ColumnResponse_IntVal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntVal))
	case *ColumnResponse_BoolVal:
		n += 1 // tag and wire
		n += 1
	case *ColumnResponse_BlobVal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BlobVal)))
		n += len(x.BlobVal)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InspectRequest struct {
	Index                string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Columns              *IdsOrKeys `protobuf:"bytes,2,opt,name=columns,proto3" json:"columns,omitempty"`
	FilterFields         []string   `protobuf:"bytes,3,rep,name=filterFields,proto3" json:"filterFields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InspectRequest) Reset()         { *m = InspectRequest{} }
func (m *InspectRequest) String() string { return proto.CompactTextString(m) }
func (*InspectRequest) ProtoMessage()    {}
func (*InspectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{4}
}

func (m *InspectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectRequest.Unmarshal(m, b)
}
func (m *InspectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectRequest.Marshal(b, m, deterministic)
}
func (m *InspectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectRequest.Merge(m, src)
}
func (m *InspectRequest) XXX_Size() int {
	return xxx_messageInfo_InspectRequest.Size(m)
}
func (m *InspectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InspectRequest proto.InternalMessageInfo

func (m *InspectRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *InspectRequest) GetColumns() *IdsOrKeys {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *InspectRequest) GetFilterFields() []string {
	if m != nil {
		return m.FilterFields
	}
	return nil
}

type Ids struct {
	Vals                 []uint64 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ids) Reset()         { *m = Ids{} }
func (m *Ids) String() string { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()    {}
func (*Ids) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{5}
}

func (m *Ids) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ids.Unmarshal(m, b)
}
func (m *Ids) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ids.Marshal(b, m, deterministic)
}
func (m *Ids) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ids.Merge(m, src)
}
func (m *Ids) XXX_Size() int {
	return xxx_messageInfo_Ids.Size(m)
}
func (m *Ids) XXX_DiscardUnknown() {
	xxx_messageInfo_Ids.DiscardUnknown(m)
}

var xxx_messageInfo_Ids proto.InternalMessageInfo

func (m *Ids) GetVals() []uint64 {
	if m != nil {
		return m.Vals
	}
	return nil
}

type Keys struct {
	Vals                 []string `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keys) Reset()         { *m = Keys{} }
func (m *Keys) String() string { return proto.CompactTextString(m) }
func (*Keys) ProtoMessage()    {}
func (*Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{6}
}

func (m *Keys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keys.Unmarshal(m, b)
}
func (m *Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keys.Marshal(b, m, deterministic)
}
func (m *Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keys.Merge(m, src)
}
func (m *Keys) XXX_Size() int {
	return xxx_messageInfo_Keys.Size(m)
}
func (m *Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_Keys proto.InternalMessageInfo

func (m *Keys) GetVals() []string {
	if m != nil {
		return m.Vals
	}
	return nil
}

type IdsOrKeys struct {
	// Types that are valid to be assigned to Type:
	//	*IdsOrKeys_Ids
	//	*IdsOrKeys_Keys
	Type                 isIdsOrKeys_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IdsOrKeys) Reset()         { *m = IdsOrKeys{} }
func (m *IdsOrKeys) String() string { return proto.CompactTextString(m) }
func (*IdsOrKeys) ProtoMessage()    {}
func (*IdsOrKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{7}
}

func (m *IdsOrKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdsOrKeys.Unmarshal(m, b)
}
func (m *IdsOrKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdsOrKeys.Marshal(b, m, deterministic)
}
func (m *IdsOrKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdsOrKeys.Merge(m, src)
}
func (m *IdsOrKeys) XXX_Size() int {
	return xxx_messageInfo_IdsOrKeys.Size(m)
}
func (m *IdsOrKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_IdsOrKeys.DiscardUnknown(m)
}

var xxx_messageInfo_IdsOrKeys proto.InternalMessageInfo

type isIdsOrKeys_Type interface {
	isIdsOrKeys_Type()
}

type IdsOrKeys_Ids struct {
	Ids *Ids `protobuf:"bytes,1,opt,name=ids,proto3,oneof"`
}

type IdsOrKeys_Keys struct {
	Keys *Keys `protobuf:"bytes,2,opt,name=keys,proto3,oneof"`
}

func (*IdsOrKeys_Ids) isIdsOrKeys_Type() {}

func (*IdsOrKeys_Keys) isIdsOrKeys_Type() {}

func (m *IdsOrKeys) GetType() isIdsOrKeys_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *IdsOrKeys) GetIds() *Ids {
	if x, ok := m.GetType().(*IdsOrKeys_Ids); ok {
		return x.Ids
	}
	return nil
}

func (m *IdsOrKeys) GetKeys() *Keys {
	if x, ok := m.GetType().(*IdsOrKeys_Keys); ok {
		return x.Keys
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*IdsOrKeys) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _IdsOrKeys_OneofMarshaler, _IdsOrKeys_OneofUnmarshaler, _IdsOrKeys_OneofSizer, []interface{}{
		(*IdsOrKeys_Ids)(nil),
		(*IdsOrKeys_Keys)(nil),
	}
}

func _IdsOrKeys_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*IdsOrKeys)
	// type
	switch x := m.Type.(type) {
	case *IdsOrKeys_Ids:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ids); err != nil {
			return err
		}
	case *IdsOrKeys_Keys:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Keys); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("IdsOrKeys.Type has unexpected type %T", x)
	}
	return nil
}

func _IdsOrKeys_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*IdsOrKeys)
	switch tag {
	case 1: // type.ids
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Ids)
		err := b.DecodeMessage(msg)
		m.Type = &IdsOrKeys_Ids{msg}
		return true, err
	case 2: // type.keys
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Keys)
		err := b.DecodeMessage(msg)
		m.Type = &IdsOrKeys_Keys{msg}
		return true, err
	default:
		return false, nil
	}
}

func _IdsOrKeys_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*IdsOrKeys)
	// type
	switch x := m.Type.(type) {
	case *IdsOrKeys_Ids:
		s := proto.Size(x.Ids)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *IdsOrKeys_Keys:
		s := proto.Size(x.Keys)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InspectResponse struct {
	Fields               []*FieldValues `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InspectResponse) Reset()         { *m = InspectResponse{} }
func (m *InspectResponse) String() string { return proto.CompactTextString(m) }
func (*InspectResponse) ProtoMessage()    {}
func (*InspectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{8}
}

func (m *InspectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectResponse.Unmarshal(m, b)
}
func (m *InspectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectResponse.Marshal(b, m, deterministic)
}
func (m *InspectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectResponse.Merge(m, src)
}
func (m *InspectResponse) XXX_Size() int {
	return xxx_messageInfo_InspectResponse.Size(m)
}
func (m *InspectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InspectResponse proto.InternalMessageInfo

func (m *InspectResponse) GetFields() []*FieldValues {
	if m != nil {
		return m.Fields
	}
	return nil
}

type FieldValues struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values               *IdsOrKeys `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FieldValues) Reset()         { *m = FieldValues{} }
func (m *FieldValues) String() string { return proto.CompactTextString(m) }
func (*FieldValues) ProtoMessage()    {}
func (*FieldValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0691a44d1e275c, []int{9}
}

func (m *FieldValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValues.Unmarshal(m, b)
}
func (m *FieldValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValues.Marshal(b, m, deterministic)
}
func (m *FieldValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValues.Merge(m, src)
}
func (m *FieldValues) XXX_Size() int {
	return xxx_messageInfo_FieldValues.Size(m)
}
func (m *FieldValues) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValues.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValues proto.InternalMessageInfo

func (m *FieldValues) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldValues) GetValues() *IdsOrKeys {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryPQLRequest)(nil), "proto.QueryPQLRequest")
	proto.RegisterType((*RowResponse)(nil), "proto.RowResponse")
	proto.RegisterType((*ColumnInfo)(nil), "proto.ColumnInfo")
	proto.RegisterType((*ColumnResponse)(nil), "proto.ColumnResponse")
	proto.RegisterType((*InspectRequest)(nil), "proto.InspectRequest")
	proto.RegisterType((*Ids)(nil), "proto.Ids")
	proto.RegisterType((*Keys)(nil), "proto.Keys")
	proto.RegisterType((*IdsOrKeys)(nil), "proto.IdsOrKeys")
	proto.RegisterType((*InspectResponse)(nil), "proto.InspectResponse")
	proto.RegisterType((*FieldValues)(nil), "proto.FieldValues")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PilosaClient is the client API for Pilosa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PilosaClient interface {
	QueryPQL(ctx context.Context, in *QueryPQLRequest, opts ...grpc.CallOption) (Pilosa_QueryPQLClient, error)
	Inspect(ctx context.Context, in *InspectRequest, opts ...grpc.CallOption) (Pilosa_InspectClient, error)
}

type pilosaClient struct {
	cc *grpc.ClientConn
}

func NewPilosaClient(cc *grpc.ClientConn) PilosaClient {
	return &pilosaClient{cc}
}

func (c *pilosaClient) QueryPQL(ctx context.Context, in *QueryPQLRequest, opts ...grpc.CallOption) (Pilosa_QueryPQLClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pilosa_serviceDesc.Streams[0], "/proto.Pilosa/QueryPQL", opts...)
	if err != nil {
		return nil, err
	}
	x := &pilosaQueryPQLClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pilosa_QueryPQLClient interface {
	Recv() (*RowResponse, error)
	grpc.ClientStream
}

type pilosaQueryPQLClient struct {
	grpc.ClientStream
}

func (x *pilosaQueryPQLClient) Recv() (*RowResponse, error) {
	m := new(RowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pilosaClient) Inspect(ctx context.Context, in *InspectRequest, opts ...grpc.CallOption) (Pilosa_InspectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pilosa_serviceDesc.Streams[1], "/proto.Pilosa/Inspect", opts...)
	if err != nil {
		return nil, err
	}
	x := &pilosaInspectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pilosa_InspectClient interface {
	Recv() (*InspectResponse, error)
	grpc.ClientStream
}

type pilosaInspectClient struct {
	grpc.ClientStream
}

func (x *pilosaInspectClient) Recv() (*InspectResponse, error) {
	m := new(InspectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PilosaServer is the server API for Pilosa service.
type PilosaServer interface {
	QueryPQL(*QueryPQLRequest, Pilosa_QueryPQLServer) error
	Inspect(*InspectRequest, Pilosa_InspectServer) error
}

func RegisterPilosaServer(s *grpc.Server, srv PilosaServer) {
	s.RegisterService(&_Pilosa_serviceDesc, srv)
}

func _Pilosa_QueryPQL_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryPQLRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PilosaServer).QueryPQL(m, &pilosaQueryPQLServer{stream})
}

type Pilosa_QueryPQLServer interface {
	Send(*RowResponse) error
	grpc.ServerStream
}

type pilosaQueryPQLServer struct {
	grpc.ServerStream
}

func (x *pilosaQueryPQLServer) Send(m *RowResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Pilosa_Inspect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InspectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PilosaServer).Inspect(m, &pilosaInspectServer{stream})
}

type Pilosa_InspectServer interface {
	Send(*InspectResponse) error
	grpc.ServerStream
}

type pilosaInspectServer struct {
	grpc.ServerStream
}

func (x *pilosaInspectServer) Send(m *InspectResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Pilosa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Pilosa",
	HandlerType: (*PilosaServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryPQL",
			Handler:       _Pilosa_QueryPQL_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Inspect",
			Handler:       _Pilosa_Inspect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pilosa.proto",
}

func init() { proto.RegisterFile("pilosa.proto", fileDescriptor_ef0691a44d1e275c) }

var fileDescriptor_ef0691a44d1e275c = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x8e, 0x97, 0x34, 0x6d, 0x5e, 0xaa, 0x6d, 0x58, 0x30, 0x85, 0x1c, 0xa6, 0xe0, 0x53, 0x34,
	0xa4, 0x81, 0xca, 0x09, 0x34, 0x2e, 0x43, 0x42, 0xa9, 0x06, 0x62, 0xf3, 0xa1, 0x77, 0x77, 0x71,
	0x47, 0x84, 0x17, 0x67, 0x71, 0x52, 0xe8, 0x95, 0x3f, 0xc0, 0x0f, 0xe1, 0x4f, 0x22, 0x3b, 0x4e,
	0xdb, 0x4c, 0xc0, 0x29, 0xf6, 0xf7, 0xbd, 0xe7, 0xf7, 0xbd, 0xef, 0xe5, 0xc1, 0xb4, 0x2a, 0x84,
	0x54, 0xec, 0xbc, 0xaa, 0x65, 0x23, 0xf1, 0xc8, 0x7c, 0xc8, 0x5b, 0x38, 0xba, 0x69, 0x79, 0xbd,
	0xb9, 0xbe, 0xf9, 0x44, 0xf9, 0x43, 0xcb, 0x55, 0x83, 0x9f, 0xc2, 0xa8, 0x28, 0x73, 0xfe, 0x23,
	0x42, 0x09, 0x4a, 0x03, 0xda, 0x5d, 0xf0, 0x31, 0xb8, 0xd5, 0x83, 0x88, 0x0e, 0x0c, 0xa6, 0x8f,
	0xe4, 0x17, 0x82, 0x90, 0xca, 0xef, 0x94, 0xab, 0x4a, 0x96, 0x8a, 0xe3, 0x97, 0x30, 0xfe, 0xca,
	0x59, 0xce, 0x6b, 0x15, 0xa1, 0xc4, 0x4d, 0xc3, 0xd9, 0x93, 0xae, 0xd4, 0xf9, 0x07, 0x29, 0xda,
	0xfb, 0x72, 0x5e, 0xae, 0x24, 0xed, 0x23, 0xf0, 0x2b, 0x18, 0xdf, 0x1a, 0x58, 0x45, 0x07, 0x26,
	0xf8, 0xd9, 0x20, 0xb8, 0x7f, 0x94, 0xf6, 0x51, 0x98, 0xc0, 0x94, 0xd7, 0xb5, 0xac, 0x3f, 0x73,
	0xa5, 0xd8, 0x1d, 0x8f, 0x5c, 0x23, 0x64, 0x80, 0x91, 0x0b, 0x80, 0x5d, 0x2d, 0x8c, 0xc1, 0x2b,
	0xd9, 0x3d, 0xb7, 0x6d, 0x98, 0x33, 0x8e, 0x61, 0x92, 0xb3, 0x86, 0x35, 0x9b, 0x8a, 0xdb, 0x56,
	0xb6, 0x77, 0xf2, 0x1b, 0xc1, 0xe1, 0xb0, 0x3a, 0x3e, 0x85, 0x40, 0x35, 0x75, 0x51, 0xde, 0x2d,
	0x98, 0xe8, 0xde, 0xc9, 0x1c, 0xba, 0x83, 0x70, 0x0c, 0xe3, 0xb6, 0x28, 0x1b, 0xcd, 0xea, 0xd7,
	0xbc, 0xcc, 0xa1, 0x3d, 0x80, 0x23, 0xf0, 0x2d, 0xa5, 0xa5, 0xba, 0x99, 0x43, 0xed, 0x5d, 0x67,
	0x2d, 0xa5, 0x14, 0x9a, 0xf2, 0x12, 0x94, 0x4e, 0x74, 0x96, 0x05, 0x0c, 0x27, 0xe4, 0x52, 0x73,
	0xa3, 0x04, 0xa5, 0x53, 0xc3, 0x75, 0xc0, 0x65, 0x08, 0x41, 0xe7, 0xc6, 0x82, 0x09, 0xb2, 0x86,
	0xc3, 0x79, 0xa9, 0x2a, 0x7e, 0xdb, 0xfc, 0x7f, 0x6e, 0x67, 0xfb, 0x46, 0xa3, 0x34, 0x9c, 0x1d,
	0x5b, 0xa3, 0xe7, 0xb9, 0xfa, 0x52, 0x5f, 0xf1, 0x8d, 0x1a, 0x78, 0xbc, 0x2a, 0x44, 0xc3, 0xeb,
	0x8f, 0x05, 0x17, 0xb9, 0x8a, 0xdc, 0xc4, 0xd5, 0x1e, 0xef, 0x63, 0xe4, 0x39, 0xb8, 0xf3, 0x5c,
	0x69, 0x73, 0xd7, 0x4c, 0x74, 0x93, 0xf6, 0xa8, 0x39, 0x93, 0x18, 0x3c, 0xfd, 0xde, 0x80, 0x0b,
	0x2c, 0xb7, 0x80, 0x60, 0x5b, 0x10, 0x9f, 0x82, 0x5b, 0xe4, 0xca, 0xe8, 0x0c, 0x67, 0xb0, 0xd3,
	0x93, 0x39, 0x54, 0x13, 0xf8, 0x05, 0x78, 0xdf, 0xf8, 0xa6, 0x17, 0x1c, 0xda, 0x00, 0x9d, 0x9a,
	0x39, 0xd4, 0x50, 0x97, 0x3e, 0x78, 0x66, 0x68, 0xef, 0xe1, 0x68, 0x6b, 0x83, 0x1d, 0xda, 0x19,
	0xf8, 0xab, 0x4e, 0x7f, 0xf7, 0x1b, 0x62, 0x9b, 0x6f, 0x1a, 0x58, 0x30, 0xd1, 0x72, 0x45, 0x6d,
	0x04, 0xb9, 0x82, 0x70, 0x0f, 0xfe, 0xeb, 0x2f, 0x93, 0x82, 0xbf, 0x36, 0xec, 0x3f, 0xfd, 0xb3,
	0xfc, 0xec, 0x27, 0x02, 0xff, 0xda, 0xec, 0x18, 0x7e, 0x07, 0x93, 0x7e, 0xad, 0xf0, 0x89, 0x4d,
	0x78, 0xb4, 0x67, 0x71, 0xaf, 0x6b, 0x6f, 0x87, 0x88, 0xf3, 0x1a, 0xe1, 0x0b, 0x18, 0xdb, 0x96,
	0x70, 0xbf, 0x14, 0xc3, 0x49, 0xc7, 0x27, 0x8f, 0xe1, 0x5d, 0xf6, 0xd2, 0x37, 0xd4, 0x9b, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xf8, 0x83, 0x2f, 0xee, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log/v1/Message.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	MessageType_RAW    MessageType = 0
	MessageType_VALUES MessageType = 1
	MessageType_ERROR  MessageType = 2
	MessageType_HTTP   MessageType = 3
	MessageType_INFO   MessageType = 4
	MessageType_SYSLOG MessageType = 5
	MessageType_AUDIT  MessageType = 6
)

var MessageType_name = map[int32]string{
	0: "RAW",
	1: "VALUES",
	2: "ERROR",
	3: "HTTP",
	4: "INFO",
	5: "SYSLOG",
	6: "AUDIT",
}
var MessageType_value = map[string]int32{
	"RAW":    0,
	"VALUES": 1,
	"ERROR":  2,
	"HTTP":   3,
	"INFO":   4,
	"SYSLOG": 5,
	"AUDIT":  6,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// [START messages]
type Message struct {
	// type is the type of the object
	// it is used principally in routing to processors.
	Type MessageType `protobuf:"varint,1,opt,name=type,enum=logv1.MessageType" json:"type,omitempty"`
	// meta contains extra metadata for the message like timestamp.
	Meta        *Meta        `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Raw         *Raw         `protobuf:"bytes,100,opt,name=raw" json:"raw,omitempty"`
	Values      *Values      `protobuf:"bytes,101,opt,name=values" json:"values,omitempty"`
	Error       *Error       `protobuf:"bytes,102,opt,name=error" json:"error,omitempty"`
	HttpRequest *HttpRequest `protobuf:"bytes,103,opt,name=httpRequest" json:"httpRequest,omitempty"`
	Info        *Info        `protobuf:"bytes,104,opt,name=info" json:"info,omitempty"`
	Syslog      *Syslog      `protobuf:"bytes,105,opt,name=syslog" json:"syslog,omitempty"`
	Audit       *Audit       `protobuf:"bytes,106,opt,name=audit" json:"audit,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_RAW
}

func (m *Message) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Message) GetRaw() *Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Message) GetValues() *Values {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Message) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Message) GetHttpRequest() *HttpRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *Message) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Message) GetSyslog() *Syslog {
	if m != nil {
		return m.Syslog
	}
	return nil
}

func (m *Message) GetAudit() *Audit {
	if m != nil {
		return m.Audit
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "logv1.Message")
	proto.RegisterEnum("logv1.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("log/v1/Message.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x49, 0x9b, 0x64, 0xe0, 0x00, 0x32, 0x86, 0x07, 0x6b, 0x42, 0x62, 0x9a, 0x34, 0x98,
	0x78, 0x48, 0xb4, 0xc1, 0x17, 0x08, 0xa2, 0xb0, 0x4a, 0x85, 0xa2, 0x9b, 0xb4, 0xfc, 0x79, 0x73,
	0xa9, 0x9b, 0x06, 0xa5, 0x75, 0x88, 0xdd, 0x56, 0xfd, 0xb0, 0x7c, 0x17, 0x64, 0xe7, 0xb6, 0x6b,
	0xfb, 0x16, 0x9f, 0xdf, 0x89, 0x7d, 0xee, 0xb9, 0xe4, 0x45, 0xa5, 0x8a, 0x64, 0x7d, 0x93, 0x7c,
	0x91, 0x5a, 0x8b, 0x42, 0xc6, 0x75, 0xa3, 0x8c, 0x62, 0x41, 0xa5, 0x8a, 0xf5, 0xcd, 0xf9, 0xb3,
	0x3d, 0x34, 0xa2, 0x25, 0xe7, 0x14, 0x25, 0x10, 0x1b, 0x54, 0x9e, 0xa3, 0x32, 0x16, 0xd5, 0x4a,
	0x6a, 0x14, 0x19, 0x8a, 0xbd, 0xa6, 0x51, 0x0d, 0x6a, 0x1c, 0xb5, 0x3b, 0x63, 0x6a, 0x90, 0x7f,
	0x57, 0x52, 0x1b, 0x24, 0xbb, 0x77, 0xfa, 0xcb, 0x99, 0x3a, 0xb9, 0x35, 0xdb, 0xea, 0x4a, 0x15,
	0x27, 0xb7, 0xa6, 0xab, 0x69, 0x89, 0xff, 0x5e, 0xfe, 0xeb, 0x90, 0x33, 0x0c, 0xcf, 0x5e, 0x13,
	0xdf, 0x6c, 0x6b, 0xc9, 0xbd, 0x0b, 0xef, 0xfa, 0xe9, 0x2d, 0x8b, 0xdd, 0x14, 0x31, 0xd2, 0x7c,
	0x5b, 0x4b, 0x70, 0x9c, 0xbd, 0x22, 0xfe, 0x42, 0x1a, 0xc1, 0x3b, 0x17, 0xde, 0x75, 0x74, 0x1b,
	0xed, 0x7d, 0x46, 0x80, 0x03, 0xec, 0x25, 0xe9, 0x36, 0x62, 0xc3, 0xa7, 0x8e, 0x13, 0xe4, 0x20,
	0x36, 0x60, 0x65, 0x76, 0x45, 0xc2, 0xb5, 0x1b, 0x96, 0x4b, 0x67, 0x78, 0x82, 0x86, 0xb6, 0x01,
	0x40, 0xc8, 0x2e, 0x49, 0x20, 0xed, 0xf8, 0x7c, 0xe6, 0x5c, 0x8f, 0xd1, 0xe5, 0x2a, 0x81, 0x16,
	0xb1, 0xf7, 0x24, 0x9a, 0xdf, 0xd7, 0xc1, 0x0b, 0xe7, 0xdc, 0x05, 0x3f, 0x28, 0x0a, 0x0e, 0x6d,
	0x36, 0x7f, 0xb9, 0x9c, 0x29, 0x3e, 0x3f, 0xca, 0x6f, 0xdb, 0x03, 0x07, 0x6c, 0x42, 0xed, 0x8a,
	0xe3, 0xe5, 0x51, 0xc2, 0xb6, 0x4d, 0x40, 0x68, 0x13, 0x0a, 0x5b, 0x25, 0xff, 0x73, 0x94, 0xd0,
	0xd5, 0x0b, 0x2d, 0x7a, 0xfb, 0x83, 0x44, 0x07, 0x05, 0xb2, 0x33, 0xd2, 0x85, 0xf4, 0x3b, 0x7d,
	0xc0, 0x08, 0x09, 0xc7, 0xe9, 0x60, 0xd4, 0xcb, 0xa8, 0xc7, 0x1e, 0x91, 0xa0, 0x07, 0x30, 0x04,
	0xda, 0x61, 0x0f, 0x89, 0x7f, 0x97, 0xe7, 0xdf, 0x68, 0xd7, 0x7e, 0xf5, 0xbf, 0x7e, 0x1a, 0x52,
	0xdf, 0x5a, 0xb3, 0x9f, 0xd9, 0x60, 0xf8, 0x99, 0x06, 0xd6, 0x9a, 0x8e, 0x3e, 0xf6, 0x73, 0x1a,
	0x7e, 0x78, 0xf3, 0xeb, 0xaa, 0x28, 0x4d, 0x3c, 0xa9, 0xe4, 0x72, 0x5a, 0x89, 0x89, 0x8e, 0x7f,
	0xab, 0x45, 0xe2, 0x4e, 0x89, 0x5b, 0xad, 0x4e, 0xda, 0x75, 0x4f, 0x42, 0x77, 0x7c, 0xf7, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0xcf, 0x90, 0xd1, 0xac, 0x02, 0x00, 0x00,
}

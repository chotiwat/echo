// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log/v1/Syslog.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SyslogSeverity int32

const (
	SyslogSeverity_SYSSEV_EMERG   SyslogSeverity = 0
	SyslogSeverity_SYSSEV_ALERT   SyslogSeverity = 1
	SyslogSeverity_SYSSEV_CRIT    SyslogSeverity = 2
	SyslogSeverity_SYSSEV_ERR     SyslogSeverity = 3
	SyslogSeverity_SYSSEV_WARNING SyslogSeverity = 4
	SyslogSeverity_SYSSEV_NOTICE  SyslogSeverity = 5
	SyslogSeverity_SYSSEV_INFO    SyslogSeverity = 6
	SyslogSeverity_SYSSEV_DEBUG   SyslogSeverity = 7
)

var SyslogSeverity_name = map[int32]string{
	0: "SYSSEV_EMERG",
	1: "SYSSEV_ALERT",
	2: "SYSSEV_CRIT",
	3: "SYSSEV_ERR",
	4: "SYSSEV_WARNING",
	5: "SYSSEV_NOTICE",
	6: "SYSSEV_INFO",
	7: "SYSSEV_DEBUG",
}
var SyslogSeverity_value = map[string]int32{
	"SYSSEV_EMERG":   0,
	"SYSSEV_ALERT":   1,
	"SYSSEV_CRIT":    2,
	"SYSSEV_ERR":     3,
	"SYSSEV_WARNING": 4,
	"SYSSEV_NOTICE":  5,
	"SYSSEV_INFO":    6,
	"SYSSEV_DEBUG":   7,
}

func (x SyslogSeverity) String() string {
	return proto.EnumName(SyslogSeverity_name, int32(x))
}
func (SyslogSeverity) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// [START messages]
type Syslog struct {
	Severity SyslogSeverity `protobuf:"varint,2,opt,name=severity,enum=logv1.SyslogSeverity" json:"severity,omitempty"`
	Facility string         `protobuf:"bytes,1,opt,name=facility" json:"facility,omitempty"`
	Message  string         `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Syslog) Reset()                    { *m = Syslog{} }
func (m *Syslog) String() string            { return proto.CompactTextString(m) }
func (*Syslog) ProtoMessage()               {}
func (*Syslog) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Syslog) GetSeverity() SyslogSeverity {
	if m != nil {
		return m.Severity
	}
	return SyslogSeverity_SYSSEV_EMERG
}

func (m *Syslog) GetFacility() string {
	if m != nil {
		return m.Facility
	}
	return ""
}

func (m *Syslog) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Syslog)(nil), "logv1.Syslog")
	proto.RegisterEnum("logv1.SyslogSeverity", SyslogSeverity_name, SyslogSeverity_value)
}

func init() { proto.RegisterFile("log/v1/Syslog.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x6b, 0x7f, 0x1c, 0x35, 0xae, 0x23, 0x42, 0xf0, 0xaa, 0x08, 0x62, 0xf1, 0x22,
	0x21, 0xfa, 0x04, 0x6d, 0x5d, 0x43, 0x40, 0x53, 0xd8, 0x8d, 0x8a, 0xde, 0x48, 0x52, 0xd7, 0x25,
	0xb0, 0x75, 0xb5, 0x1b, 0x02, 0x7d, 0x1a, 0x5f, 0x55, 0xdc, 0xa4, 0x25, 0x5e, 0x7e, 0xdf, 0x9c,
	0x99, 0x81, 0x03, 0x27, 0x4a, 0xcb, 0xa0, 0x0a, 0x03, 0xbe, 0x36, 0x4a, 0x4b, 0xff, 0x6b, 0xa5,
	0x4b, 0x8d, 0x3d, 0xa5, 0x65, 0x15, 0x9e, 0x7f, 0x43, 0xbf, 0xd6, 0x18, 0xc2, 0xd0, 0x88, 0x4a,
	0xac, 0x8a, 0x72, 0xed, 0x75, 0x46, 0xce, 0xd8, 0xbd, 0x3e, 0xf5, 0x6d, 0xc6, 0xaf, 0x03, 0xbc,
	0x19, 0xb2, 0x6d, 0x0c, 0xcf, 0x60, 0xf8, 0x91, 0x2d, 0x0a, 0xf5, 0xb7, 0xe2, 0x8c, 0x9c, 0xf1,
	0x1e, 0xdb, 0x32, 0x7a, 0x30, 0x58, 0x0a, 0x63, 0x32, 0x29, 0xbc, 0xae, 0x1d, 0x6d, 0xf0, 0xea,
	0xc7, 0x01, 0xf7, 0xff, 0x49, 0x24, 0x70, 0xc0, 0x5f, 0x38, 0xa7, 0x4f, 0x6f, 0xf4, 0x81, 0xb2,
	0x88, 0xec, 0xb4, 0xcc, 0xe4, 0x9e, 0xb2, 0x94, 0x38, 0x78, 0x04, 0xfb, 0x8d, 0x99, 0xb1, 0x38,
	0x25, 0x1d, 0x74, 0x01, 0x36, 0x4b, 0x8c, 0x91, 0x2e, 0x22, 0xb8, 0x0d, 0x3f, 0x4f, 0x58, 0x12,
	0x27, 0x11, 0xd9, 0xc5, 0x63, 0x38, 0x6c, 0x5c, 0x32, 0x4f, 0xe3, 0x19, 0x25, 0xbd, 0xd6, 0x9d,
	0x38, 0xb9, 0x9b, 0x93, 0x7e, 0xeb, 0xd5, 0x2d, 0x9d, 0x3e, 0x46, 0x64, 0x30, 0xbd, 0x7c, 0xbd,
	0x90, 0x45, 0xe9, 0xe7, 0x4a, 0x7c, 0xbe, 0xab, 0x2c, 0x37, 0xfe, 0x42, 0x2f, 0x03, 0x4b, 0x81,
	0x2d, 0xcf, 0x04, 0x75, 0xa3, 0x79, 0xdf, 0xe2, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0xd4, 0xc1, 0xf1, 0x62, 0x01, 0x00, 0x00,
}

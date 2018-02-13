// Code generated by protoc-gen-go.
// source: signalling.proto
// DO NOT EDIT!

/*
Package tapdance is a generated protocol buffer package.

It is generated from these files:
	signalling.proto

It has these top-level messages:
	PubKey
	TLSDecoySpec
	ClientConf
	DecoyList
	StationToClient
	ClientToStation
	DupOvUrl
*/
package tdproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyType int32

const (
	KeyType_AES_GCM_128 KeyType = 90
	KeyType_AES_GCM_256 KeyType = 91
)

var KeyType_name = map[int32]string{
	90: "AES_GCM_128",
	91: "AES_GCM_256",
}
var KeyType_value = map[string]int32{
	"AES_GCM_128": 90,
	"AES_GCM_256": 91,
}

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}
func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}
func (x *KeyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(KeyType_value, data, "KeyType")
	if err != nil {
		return err
	}
	*x = KeyType(value)
	return nil
}
func (KeyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// State transitions of the client
type C2S_Transition int32

const (
	C2S_Transition_C2S_NO_CHANGE C2S_Transition = 0
	// C2S_SESSION_INIT = 1; // Currently MSG_INIT is NOT protobuf
	C2S_Transition_C2S_EXPECT_RECONNECT         C2S_Transition = 2
	C2S_Transition_C2S_SESSION_CLOSE            C2S_Transition = 3
	C2S_Transition_C2S_YIELD_UPLOAD             C2S_Transition = 4
	C2S_Transition_C2S_ACQUIRE_UPLOAD           C2S_Transition = 5
	C2S_Transition_C2S_EXPECT_UPLOADONLY_RECONN C2S_Transition = 6
	C2S_Transition_C2S_ERROR                    C2S_Transition = 255
)

var C2S_Transition_name = map[int32]string{
	0:   "C2S_NO_CHANGE",
	2:   "C2S_EXPECT_RECONNECT",
	3:   "C2S_SESSION_CLOSE",
	4:   "C2S_YIELD_UPLOAD",
	5:   "C2S_ACQUIRE_UPLOAD",
	6:   "C2S_EXPECT_UPLOADONLY_RECONN",
	255: "C2S_ERROR",
}
var C2S_Transition_value = map[string]int32{
	"C2S_NO_CHANGE":                0,
	"C2S_EXPECT_RECONNECT":         2,
	"C2S_SESSION_CLOSE":            3,
	"C2S_YIELD_UPLOAD":             4,
	"C2S_ACQUIRE_UPLOAD":           5,
	"C2S_EXPECT_UPLOADONLY_RECONN": 6,
	"C2S_ERROR":                    255,
}

func (x C2S_Transition) Enum() *C2S_Transition {
	p := new(C2S_Transition)
	*p = x
	return p
}
func (x C2S_Transition) String() string {
	return proto.EnumName(C2S_Transition_name, int32(x))
}
func (x *C2S_Transition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(C2S_Transition_value, data, "C2S_Transition")
	if err != nil {
		return err
	}
	*x = C2S_Transition(value)
	return nil
}
func (C2S_Transition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// State transitions of the server
type S2C_Transition int32

const (
	S2C_Transition_S2C_NO_CHANGE         S2C_Transition = 0
	S2C_Transition_S2C_SESSION_INIT      S2C_Transition = 1
	S2C_Transition_S2C_CONFIRM_RECONNECT S2C_Transition = 2
	S2C_Transition_S2C_SESSION_CLOSE     S2C_Transition = 3
	S2C_Transition_S2C_ERROR             S2C_Transition = 255
)

var S2C_Transition_name = map[int32]string{
	0:   "S2C_NO_CHANGE",
	1:   "S2C_SESSION_INIT",
	2:   "S2C_CONFIRM_RECONNECT",
	3:   "S2C_SESSION_CLOSE",
	255: "S2C_ERROR",
}
var S2C_Transition_value = map[string]int32{
	"S2C_NO_CHANGE":         0,
	"S2C_SESSION_INIT":      1,
	"S2C_CONFIRM_RECONNECT": 2,
	"S2C_SESSION_CLOSE":     3,
	"S2C_ERROR":             255,
}

func (x S2C_Transition) Enum() *S2C_Transition {
	p := new(S2C_Transition)
	*p = x
	return p
}
func (x S2C_Transition) String() string {
	return proto.EnumName(S2C_Transition_name, int32(x))
}
func (x *S2C_Transition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(S2C_Transition_value, data, "S2C_Transition")
	if err != nil {
		return err
	}
	*x = S2C_Transition(value)
	return nil
}
func (S2C_Transition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Should accompany all S2C_ERROR messages.
type ErrorReasonS2C int32

const (
	ErrorReasonS2C_NO_ERROR         ErrorReasonS2C = 0
	ErrorReasonS2C_COVERT_STREAM    ErrorReasonS2C = 1
	ErrorReasonS2C_CLIENT_REPORTED  ErrorReasonS2C = 2
	ErrorReasonS2C_CLIENT_PROTOCOL  ErrorReasonS2C = 3
	ErrorReasonS2C_STATION_INTERNAL ErrorReasonS2C = 4
	ErrorReasonS2C_DECOY_OVERLOAD   ErrorReasonS2C = 5
	ErrorReasonS2C_CLIENT_STREAM    ErrorReasonS2C = 100
	ErrorReasonS2C_CLIENT_TIMEOUT   ErrorReasonS2C = 101
)

var ErrorReasonS2C_name = map[int32]string{
	0:   "NO_ERROR",
	1:   "COVERT_STREAM",
	2:   "CLIENT_REPORTED",
	3:   "CLIENT_PROTOCOL",
	4:   "STATION_INTERNAL",
	5:   "DECOY_OVERLOAD",
	100: "CLIENT_STREAM",
	101: "CLIENT_TIMEOUT",
}
var ErrorReasonS2C_value = map[string]int32{
	"NO_ERROR":         0,
	"COVERT_STREAM":    1,
	"CLIENT_REPORTED":  2,
	"CLIENT_PROTOCOL":  3,
	"STATION_INTERNAL": 4,
	"DECOY_OVERLOAD":   5,
	"CLIENT_STREAM":    100,
	"CLIENT_TIMEOUT":   101,
}

func (x ErrorReasonS2C) Enum() *ErrorReasonS2C {
	p := new(ErrorReasonS2C)
	*p = x
	return p
}
func (x ErrorReasonS2C) String() string {
	return proto.EnumName(ErrorReasonS2C_name, int32(x))
}
func (x *ErrorReasonS2C) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorReasonS2C_value, data, "ErrorReasonS2C")
	if err != nil {
		return err
	}
	*x = ErrorReasonS2C(value)
	return nil
}
func (ErrorReasonS2C) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PubKey struct {
	// A public key, as used by the station.
	Key              []byte   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Type             *KeyType `protobuf:"varint,2,opt,name=type,enum=tapdance.KeyType" json:"type,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PubKey) Reset()                    { *m = PubKey{} }
func (m *PubKey) String() string            { return proto.CompactTextString(m) }
func (*PubKey) ProtoMessage()               {}
func (*PubKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PubKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PubKey) GetType() KeyType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return KeyType_AES_GCM_128
}

type TLSDecoySpec struct {
	// The hostname/SNI to use for this host
	//
	// The hostname is the only required field, although other
	// fields are expected to be present in most cases.
	Hostname *string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	// The 32-bit ipv4 address, in network byte order
	//
	// If the IPv4 address is absent, then it may be resolved via
	// DNS by the client, or the client may discard this decoy spec
	// if local DNS is untrusted, or the service may be multihomed.
	Ipv4Addr *uint32 `protobuf:"fixed32,2,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
	// The Tapdance station public key to use when contacting this
	// decoy
	//
	// If omitted, the default station public key (if any) is used.
	Pubkey *PubKey `protobuf:"bytes,3,opt,name=pubkey" json:"pubkey,omitempty"`
	// The maximum duration, in milliseconds, to maintain an open
	// connection to this decoy (because the decoy may close the
	// connection itself after this length of time)
	//
	// If omitted, a default of 30,000 milliseconds is assumed.
	Timeout *uint32 `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	// The maximum TCP window size to attempt to use for this decoy.
	//
	// If omitted, a default of 15360 is assumed.
	//
	// TODO: the default is based on the current heuristic of only
	// using decoys that permit windows of 15KB or larger.  If this
	// heuristic changes, then this default doesn't make sense.
	Tcpwin           *uint32 `protobuf:"varint,5,opt,name=tcpwin" json:"tcpwin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TLSDecoySpec) Reset()                    { *m = TLSDecoySpec{} }
func (m *TLSDecoySpec) String() string            { return proto.CompactTextString(m) }
func (*TLSDecoySpec) ProtoMessage()               {}
func (*TLSDecoySpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TLSDecoySpec) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *TLSDecoySpec) GetIpv4Addr() uint32 {
	if m != nil && m.Ipv4Addr != nil {
		return *m.Ipv4Addr
	}
	return 0
}

func (m *TLSDecoySpec) GetPubkey() *PubKey {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *TLSDecoySpec) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *TLSDecoySpec) GetTcpwin() uint32 {
	if m != nil && m.Tcpwin != nil {
		return *m.Tcpwin
	}
	return 0
}

type ClientConf struct {
	DecoyList        *DecoyList `protobuf:"bytes,1,opt,name=decoy_list,json=decoyList" json:"decoy_list,omitempty"`
	Generation       *uint32    `protobuf:"varint,2,opt,name=generation" json:"generation,omitempty"`
	DefaultPubkey    *PubKey    `protobuf:"bytes,3,opt,name=default_pubkey,json=defaultPubkey" json:"default_pubkey,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ClientConf) Reset()                    { *m = ClientConf{} }
func (m *ClientConf) String() string            { return proto.CompactTextString(m) }
func (*ClientConf) ProtoMessage()               {}
func (*ClientConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientConf) GetDecoyList() *DecoyList {
	if m != nil {
		return m.DecoyList
	}
	return nil
}

func (m *ClientConf) GetGeneration() uint32 {
	if m != nil && m.Generation != nil {
		return *m.Generation
	}
	return 0
}

func (m *ClientConf) GetDefaultPubkey() *PubKey {
	if m != nil {
		return m.DefaultPubkey
	}
	return nil
}

type DecoyList struct {
	TlsDecoys        []*TLSDecoySpec `protobuf:"bytes,1,rep,name=tls_decoys,json=tlsDecoys" json:"tls_decoys,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DecoyList) Reset()                    { *m = DecoyList{} }
func (m *DecoyList) String() string            { return proto.CompactTextString(m) }
func (*DecoyList) ProtoMessage()               {}
func (*DecoyList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DecoyList) GetTlsDecoys() []*TLSDecoySpec {
	if m != nil {
		return m.TlsDecoys
	}
	return nil
}

type StationToClient struct {
	// Should accompany (at least) SESSION_INIT and CONFIRM_RECONNECT.
	ProtocolVersion *uint32 `protobuf:"varint,1,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
	// There might be a state transition. May be absent; absence should be
	// treated identically to NO_CHANGE.
	StateTransition *S2C_Transition `protobuf:"varint,2,opt,name=state_transition,json=stateTransition,enum=tapdance.S2C_Transition" json:"state_transition,omitempty"`
	// The station can send client config info piggybacked
	// on any message, as it sees fit
	ConfigInfo *ClientConf `protobuf:"bytes,3,opt,name=config_info,json=configInfo" json:"config_info,omitempty"`
	// If state_transition == S2C_ERROR, this field is the explanation.
	ErrReason *ErrorReasonS2C `protobuf:"varint,4,opt,name=err_reason,json=errReason,enum=tapdance.ErrorReasonS2C" json:"err_reason,omitempty"`
	// Signals client to stop connecting for following amount of seconds
	TmpBackoff      *uint32 `protobuf:"varint,5,opt,name=tmp_backoff,json=tmpBackoff" json:"tmp_backoff,omitempty"`
	NonleafContents []byte  `protobuf:"bytes,15,opt,name=nonleaf_contents,json=nonleafContents" json:"nonleaf_contents,omitempty"`
	NonleafComplete *bool   `protobuf:"varint,16,opt,name=nonleaf_complete,json=nonleafComplete" json:"nonleaf_complete,omitempty"`
	LeafComplete    *bool   `protobuf:"varint,17,opt,name=leaf_complete,json=leafComplete" json:"leaf_complete,omitempty"`
	// Random-sized junk to defeat packet size fingerprinting.
	Padding          []byte `protobuf:"bytes,100,opt,name=padding" json:"padding,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StationToClient) Reset()                    { *m = StationToClient{} }
func (m *StationToClient) String() string            { return proto.CompactTextString(m) }
func (*StationToClient) ProtoMessage()               {}
func (*StationToClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StationToClient) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

func (m *StationToClient) GetStateTransition() S2C_Transition {
	if m != nil && m.StateTransition != nil {
		return *m.StateTransition
	}
	return S2C_Transition_S2C_NO_CHANGE
}

func (m *StationToClient) GetConfigInfo() *ClientConf {
	if m != nil {
		return m.ConfigInfo
	}
	return nil
}

func (m *StationToClient) GetErrReason() ErrorReasonS2C {
	if m != nil && m.ErrReason != nil {
		return *m.ErrReason
	}
	return ErrorReasonS2C_NO_ERROR
}

func (m *StationToClient) GetTmpBackoff() uint32 {
	if m != nil && m.TmpBackoff != nil {
		return *m.TmpBackoff
	}
	return 0
}

func (m *StationToClient) GetNonleafContents() []byte {
	if m != nil {
		return m.NonleafContents
	}
	return nil
}

func (m *StationToClient) GetNonleafComplete() bool {
	if m != nil && m.NonleafComplete != nil {
		return *m.NonleafComplete
	}
	return false
}

func (m *StationToClient) GetLeafComplete() bool {
	if m != nil && m.LeafComplete != nil {
		return *m.LeafComplete
	}
	return false
}

func (m *StationToClient) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

type ClientToStation struct {
	ProtocolVersion *uint32 `protobuf:"varint,1,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
	// The client reports its decoy list's version number here, which the
	// station can use to decide whether to send an updated one. The station
	// should always send a list if this field is set to 0.
	DecoyListGeneration *uint32         `protobuf:"varint,2,opt,name=decoy_list_generation,json=decoyListGeneration" json:"decoy_list_generation,omitempty"`
	StateTransition     *C2S_Transition `protobuf:"varint,3,opt,name=state_transition,json=stateTransition,enum=tapdance.C2S_Transition" json:"state_transition,omitempty"`
	// The position in the overall session's upload sequence where the current
	// YIELD=>ACQUIRE switchover is happening.
	UploadSync *uint64 `protobuf:"varint,4,opt,name=upload_sync,json=uploadSync" json:"upload_sync,omitempty"`
	//
	FailedDecoys []string `protobuf:"bytes,10,rep,name=failed_decoys,json=failedDecoys" json:"failed_decoys,omitempty"`
	// Random-sized junk to defeat packet size fingerprinting.
	Padding []byte `protobuf:"bytes,100,opt,name=padding" json:"padding,omitempty"`
	// (DittoTap-only)
	// Item(s) the station should add to its to-dup-overt-download queue. These
	// are assumed to be to the same decoy webserver as this session started on,
	// so rather than full URL, it should just be the path:
	// "/foo.html" instead of "https://a.b.com/foo.html", or
	// "/bar.png"  instead of "https://a.b.com/bar.png"
	OvertUrl []*DupOvUrl `protobuf:"bytes,101,rep,name=overt_url,json=overtUrl" json:"overt_url,omitempty"`
	// (DittoTap-only)
	// The HTTP Host field to use. The client only needs to specify this once,
	// at the beginning. (If we start getting the initial nonleaf URL from the
	// incomplete HTTP request, we can get this from there and drop this field).
	OvertHost        *string `protobuf:"bytes,102,opt,name=overt_host,json=overtHost" json:"overt_host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientToStation) Reset()                    { *m = ClientToStation{} }
func (m *ClientToStation) String() string            { return proto.CompactTextString(m) }
func (*ClientToStation) ProtoMessage()               {}
func (*ClientToStation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClientToStation) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

func (m *ClientToStation) GetDecoyListGeneration() uint32 {
	if m != nil && m.DecoyListGeneration != nil {
		return *m.DecoyListGeneration
	}
	return 0
}

func (m *ClientToStation) GetStateTransition() C2S_Transition {
	if m != nil && m.StateTransition != nil {
		return *m.StateTransition
	}
	return C2S_Transition_C2S_NO_CHANGE
}

func (m *ClientToStation) GetUploadSync() uint64 {
	if m != nil && m.UploadSync != nil {
		return *m.UploadSync
	}
	return 0
}

func (m *ClientToStation) GetFailedDecoys() []string {
	if m != nil {
		return m.FailedDecoys
	}
	return nil
}

func (m *ClientToStation) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

func (m *ClientToStation) GetOvertUrl() []*DupOvUrl {
	if m != nil {
		return m.OvertUrl
	}
	return nil
}

func (m *ClientToStation) GetOvertHost() string {
	if m != nil && m.OvertHost != nil {
		return *m.OvertHost
	}
	return ""
}

type DupOvUrl struct {
	// This should be the resource requested by the client for traffic shaping:
	//     e.g. "/", "/images/kittens.jpg"
	OvertUrl *string `protobuf:"bytes,1,opt,name=overt_url,json=overtUrl" json:"overt_url,omitempty"`
	// This is a boolean specifying if the client wants the data from this
	// resource or not.
	IsLeaf           *bool  `protobuf:"varint,2,opt,name=is_leaf,json=isLeaf" json:"is_leaf,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DupOvUrl) Reset()                    { *m = DupOvUrl{} }
func (m *DupOvUrl) String() string            { return proto.CompactTextString(m) }
func (*DupOvUrl) ProtoMessage()               {}
func (*DupOvUrl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DupOvUrl) GetOvertUrl() string {
	if m != nil && m.OvertUrl != nil {
		return *m.OvertUrl
	}
	return ""
}

func (m *DupOvUrl) GetIsLeaf() bool {
	if m != nil && m.IsLeaf != nil {
		return *m.IsLeaf
	}
	return false
}

func init() {
	proto.RegisterType((*PubKey)(nil), "tapdance.PubKey")
	proto.RegisterType((*TLSDecoySpec)(nil), "tapdance.TLSDecoySpec")
	proto.RegisterType((*ClientConf)(nil), "tapdance.ClientConf")
	proto.RegisterType((*DecoyList)(nil), "tapdance.DecoyList")
	proto.RegisterType((*StationToClient)(nil), "tapdance.StationToClient")
	proto.RegisterType((*ClientToStation)(nil), "tapdance.ClientToStation")
	proto.RegisterType((*DupOvUrl)(nil), "tapdance.DupOvUrl")
	proto.RegisterEnum("tapdance.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("tapdance.C2S_Transition", C2S_Transition_name, C2S_Transition_value)
	proto.RegisterEnum("tapdance.S2C_Transition", S2C_Transition_name, S2C_Transition_value)
	proto.RegisterEnum("tapdance.ErrorReasonS2C", ErrorReasonS2C_name, ErrorReasonS2C_value)
}

func init() { proto.RegisterFile("signalling.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 960 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0x5e, 0x37, 0xdd, 0x36, 0x39, 0x69, 0x93, 0xe9, 0xf4, 0x82, 0xb9, 0x6e, 0x14, 0x84, 0x14,
	0x8a, 0x54, 0x84, 0x45, 0x59, 0x1e, 0xc9, 0xba, 0xa6, 0x1b, 0x6d, 0x6a, 0x87, 0xb1, 0xbb, 0xa2,
	0xf0, 0x30, 0x72, 0xe3, 0x71, 0xb1, 0xd6, 0x9d, 0xb1, 0xec, 0x49, 0x51, 0xc4, 0x2f, 0xe1, 0x95,
	0x17, 0x5e, 0x78, 0xe4, 0x2f, 0xf1, 0x3b, 0x40, 0x33, 0xb6, 0x13, 0x67, 0xb9, 0x69, 0xdf, 0x7c,
	0xbe, 0x73, 0xfb, 0xce, 0x77, 0x8e, 0x07, 0x50, 0x91, 0xdc, 0xf1, 0x30, 0x4d, 0x13, 0x7e, 0x77,
	0x96, 0xe5, 0x42, 0x0a, 0xdc, 0x96, 0x61, 0x16, 0x85, 0x7c, 0xce, 0x86, 0x63, 0xd8, 0x99, 0x2d,
	0x6e, 0x5f, 0xb0, 0x25, 0x46, 0xd0, 0x7a, 0xc5, 0x96, 0xa6, 0x31, 0x30, 0x46, 0x7b, 0x44, 0x7d,
	0xe2, 0x8f, 0x60, 0x5b, 0x2e, 0x33, 0x66, 0x6e, 0x0d, 0x8c, 0x51, 0xcf, 0x3a, 0x38, 0xab, 0x93,
	0xce, 0x5e, 0xb0, 0x65, 0xb0, 0xcc, 0x18, 0xd1, 0xee, 0xe1, 0x2f, 0x06, 0xec, 0x05, 0x53, 0xff,
	0x82, 0xcd, 0xc5, 0xd2, 0xcf, 0xd8, 0x1c, 0xbf, 0x03, 0xed, 0x1f, 0x44, 0x21, 0x79, 0x78, 0xcf,
	0x74, 0xb9, 0x0e, 0x59, 0xd9, 0xca, 0x97, 0x64, 0x0f, 0x9f, 0x87, 0x51, 0x94, 0xeb, 0xba, 0xbb,
	0x64, 0x65, 0xe3, 0x11, 0xec, 0x64, 0x8b, 0x5b, 0x45, 0xa2, 0x35, 0x30, 0x46, 0x5d, 0x0b, 0xad,
	0x3b, 0x96, 0x1c, 0x49, 0xe5, 0xc7, 0x26, 0xec, 0xca, 0xe4, 0x9e, 0x89, 0x85, 0x34, 0xb7, 0x07,
	0xc6, 0x68, 0x9f, 0xd4, 0x26, 0x3e, 0x81, 0x1d, 0x39, 0xcf, 0x7e, 0x4c, 0xb8, 0xf9, 0x58, 0x3b,
	0x2a, 0x6b, 0xf8, 0xb3, 0x01, 0x60, 0xa7, 0x09, 0xe3, 0xd2, 0x16, 0x3c, 0xc6, 0x16, 0x40, 0xa4,
	0xf8, 0xd2, 0x34, 0x29, 0xa4, 0x26, 0xd9, 0xb5, 0x0e, 0xd7, 0xed, 0xf4, 0x2c, 0xd3, 0xa4, 0x90,
	0xa4, 0x13, 0xd5, 0x9f, 0xf8, 0x03, 0x80, 0x3b, 0xc6, 0x59, 0x1e, 0xca, 0x44, 0x70, 0x4d, 0x7e,
	0x9f, 0x34, 0x10, 0xfc, 0x14, 0x7a, 0x11, 0x8b, 0xc3, 0x45, 0x2a, 0xe9, 0xff, 0x8c, 0xb1, 0x5f,
	0xc5, 0xcd, 0x74, 0xd8, 0xf0, 0x19, 0x74, 0x56, 0x0d, 0xf1, 0x39, 0x80, 0x4c, 0x0b, 0xaa, 0xdb,
	0x16, 0xa6, 0x31, 0x68, 0x8d, 0xba, 0xd6, 0xc9, 0xba, 0x42, 0x53, 0x68, 0xd2, 0x91, 0x69, 0xa1,
	0xad, 0x62, 0xf8, 0x6b, 0x0b, 0xfa, 0xbe, 0xd4, 0x44, 0x02, 0x51, 0x0e, 0x8a, 0x3f, 0x06, 0xa4,
	0xd7, 0x3d, 0x17, 0x29, 0x7d, 0x60, 0x79, 0xa1, 0x68, 0x1b, 0x9a, 0x76, 0xbf, 0xc6, 0x5f, 0x96,
	0x30, 0xb6, 0x01, 0x15, 0x32, 0x94, 0x8c, 0xca, 0x3c, 0xe4, 0x45, 0xb2, 0x9a, 0xb0, 0x67, 0x99,
	0xeb, 0xde, 0xbe, 0x65, 0xd3, 0x60, 0xe5, 0x27, 0x7d, 0x9d, 0xb1, 0x06, 0xf0, 0x39, 0x74, 0xe7,
	0x82, 0xc7, 0xc9, 0x1d, 0x4d, 0x78, 0x2c, 0xaa, 0xe9, 0x8f, 0xd6, 0xf9, 0x6b, 0xfd, 0x09, 0x94,
	0x81, 0x13, 0x1e, 0x0b, 0xfc, 0x14, 0x80, 0xe5, 0x39, 0xcd, 0x59, 0x58, 0x08, 0xae, 0xf7, 0xb9,
	0xd1, 0xd5, 0xc9, 0x73, 0x91, 0x13, 0xed, 0xf4, 0x2d, 0x9b, 0x74, 0x58, 0x5e, 0x59, 0xf8, 0x09,
	0x74, 0xe5, 0x7d, 0x46, 0x6f, 0xc3, 0xf9, 0x2b, 0x11, 0xc7, 0xd5, 0xc2, 0x41, 0xde, 0x67, 0xcf,
	0x4a, 0x44, 0x09, 0xc0, 0x05, 0x4f, 0x59, 0x18, 0xd3, 0xb9, 0xe0, 0x92, 0x71, 0x59, 0x98, 0x7d,
	0x7d, 0xdf, 0xfd, 0x0a, 0xb7, 0x2b, 0x78, 0x33, 0xf4, 0x3e, 0x4b, 0x99, 0x64, 0x26, 0x1a, 0x18,
	0xa3, 0x76, 0x23, 0xb4, 0x84, 0xf1, 0x87, 0xb0, 0xbf, 0x19, 0x77, 0xa0, 0xe3, 0xf6, 0x36, 0x82,
	0x4c, 0xd8, 0xcd, 0xc2, 0x28, 0x4a, 0xf8, 0x9d, 0x19, 0xe9, 0x8e, 0xb5, 0x39, 0xfc, 0x63, 0x0b,
	0xfa, 0xa5, 0x12, 0x81, 0xa8, 0x36, 0xf6, 0x26, 0x9b, 0xb2, 0xe0, 0x78, 0x7d, 0xb9, 0xf4, 0x6f,
	0x07, 0x79, 0xb8, 0xba, 0xd7, 0xcb, 0xf5, 0x65, 0xfe, 0xd3, 0x76, 0x5b, 0xaf, 0xeb, 0x6c, 0x5b,
	0xfe, 0x7f, 0x6e, 0xf7, 0x09, 0x74, 0x17, 0x59, 0x2a, 0xc2, 0x88, 0x16, 0x4b, 0x3e, 0xd7, 0x7b,
	0xda, 0x26, 0x50, 0x42, 0xfe, 0x92, 0xcf, 0x95, 0x2e, 0x71, 0x98, 0xa4, 0x2c, 0xaa, 0x8f, 0x17,
	0x06, 0xad, 0x51, 0x87, 0xec, 0x95, 0x60, 0x79, 0xa7, 0xff, 0xae, 0x0b, 0xfe, 0x14, 0x3a, 0xe2,
	0x81, 0xe5, 0x92, 0x2e, 0xf2, 0xd4, 0x64, 0xfa, 0xee, 0x71, 0xe3, 0x8f, 0x5c, 0x64, 0xde, 0xc3,
	0x75, 0x9e, 0x92, 0xb6, 0x0e, 0xba, 0xce, 0x53, 0xfc, 0x3e, 0x40, 0x99, 0xa0, 0x1e, 0x17, 0x33,
	0xd6, 0x0f, 0x4d, 0x59, 0xe2, 0xb9, 0x28, 0xe4, 0xf0, 0x2b, 0x68, 0xd7, 0x49, 0xf8, 0xdd, 0x66,
	0xed, 0xea, 0x49, 0x5a, 0xd5, 0x79, 0x0b, 0x76, 0x93, 0x82, 0xaa, 0xed, 0x69, 0x0d, 0xdb, 0x64,
	0x27, 0x29, 0xa6, 0x2c, 0x8c, 0x4f, 0x3f, 0x81, 0xdd, 0xea, 0xa5, 0xc3, 0x7d, 0xe8, 0x8e, 0x1d,
	0x9f, 0x5e, 0xda, 0x57, 0xf4, 0x33, 0xeb, 0x4b, 0xf4, 0x5d, 0x13, 0xb0, 0xce, 0xbf, 0x40, 0xdf,
	0x9f, 0xfe, 0x6e, 0x40, 0x6f, 0x53, 0x42, 0x7c, 0x00, 0xfb, 0x0a, 0x71, 0x3d, 0x6a, 0x3f, 0x1f,
	0xbb, 0x97, 0x0e, 0x7a, 0x84, 0x4d, 0x38, 0x52, 0x90, 0xf3, 0xed, 0xcc, 0xb1, 0x03, 0x4a, 0x1c,
	0xdb, 0x73, 0x5d, 0xc7, 0x0e, 0xd0, 0x16, 0x3e, 0x86, 0x03, 0xe5, 0xf1, 0x1d, 0xdf, 0x9f, 0x78,
	0x2e, 0xb5, 0xa7, 0x9e, 0xef, 0xa0, 0x16, 0x3e, 0x02, 0xa4, 0xe0, 0x9b, 0x89, 0x33, 0xbd, 0xa0,
	0xd7, 0xb3, 0xa9, 0x37, 0xbe, 0x40, 0xdb, 0xf8, 0x04, 0xb0, 0x42, 0xc7, 0xf6, 0x37, 0xd7, 0x13,
	0xe2, 0xd4, 0xf8, 0x63, 0x3c, 0x80, 0xf7, 0x1a, 0xe5, 0x4b, 0xd8, 0x73, 0xa7, 0x37, 0x55, 0x27,
	0xb4, 0x83, 0x7b, 0xd0, 0xd1, 0x11, 0x84, 0x78, 0x04, 0xfd, 0x69, 0x9c, 0xfe, 0x04, 0xbd, 0xcd,
	0xdf, 0x5a, 0xb1, 0x56, 0x48, 0x93, 0xf5, 0x11, 0x20, 0x05, 0xd5, 0xdc, 0x26, 0xee, 0x24, 0x40,
	0x06, 0x7e, 0x1b, 0x8e, 0x15, 0x6a, 0x7b, 0xee, 0xd7, 0x13, 0x72, 0xf5, 0xfa, 0x30, 0xcd, 0x84,
	0x7a, 0x98, 0x1e, 0x74, 0x14, 0xbc, 0x6a, 0xfe, 0x9b, 0x01, 0xbd, 0xcd, 0xdf, 0x1b, 0xef, 0x41,
	0xdb, 0xf5, 0xaa, 0x88, 0x47, 0x5a, 0x41, 0xef, 0xa5, 0x43, 0x02, 0xea, 0x07, 0xc4, 0x19, 0x5f,
	0x21, 0x03, 0x1f, 0x42, 0xdf, 0x9e, 0x4e, 0x1c, 0x57, 0xa9, 0x37, 0xf3, 0x48, 0xe0, 0x5c, 0xa0,
	0xad, 0x06, 0x38, 0x23, 0x5e, 0xe0, 0xd9, 0xde, 0xb4, 0x94, 0xce, 0x0f, 0xc6, 0x41, 0xc9, 0x38,
	0x70, 0x88, 0x3b, 0x9e, 0xa2, 0x6d, 0x8c, 0xa1, 0x77, 0xe1, 0xd8, 0xde, 0x0d, 0x55, 0x75, 0x2b,
	0xd9, 0x54, 0x9b, 0x32, 0xbd, 0x6a, 0x13, 0xa9, 0xb0, 0x0a, 0x0a, 0x26, 0x57, 0x8e, 0x77, 0x1d,
	0x20, 0xf6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd5, 0x95, 0xa7, 0x48, 0x07, 0x00, 0x00,
}
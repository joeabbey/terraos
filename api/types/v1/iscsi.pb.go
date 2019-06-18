// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stellarproject/terraos/api/types/v1/iscsi.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Target struct {
	Iqn                  string   `protobuf:"bytes,1,opt,name=iqn,proto3" json:"iqn,omitempty"`
	ID                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Luns                 []*LUN   `protobuf:"bytes,3,rep,name=luns,proto3" json:"luns,omitempty"`
	Accepts              []string `protobuf:"bytes,4,rep,name=accepts,proto3" json:"accepts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()      { *m = Target{} }
func (*Target) ProtoMessage() {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fc248a35ef3cab8, []int{0}
}
func (m *Target) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Target.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return m.Size()
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

type LUN struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	FsSize               int64    `protobuf:"varint,3,opt,name=fs_size,json=fsSize,proto3" json:"fs_size,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LUN) Reset()      { *m = LUN{} }
func (*LUN) ProtoMessage() {}
func (*LUN) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fc248a35ef3cab8, []int{1}
}
func (m *LUN) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LUN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LUN.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LUN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LUN.Merge(m, src)
}
func (m *LUN) XXX_Size() int {
	return m.Size()
}
func (m *LUN) XXX_DiscardUnknown() {
	xxx_messageInfo_LUN.DiscardUnknown(m)
}

var xxx_messageInfo_LUN proto.InternalMessageInfo

type ISCSIState struct {
	Targets              []*Target       `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	UnallocatedLuns      map[string]*LUN `protobuf:"bytes,2,rep,name=unallocated_luns,json=unallocatedLuns,proto3" json:"unallocated_luns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ISCSIState) Reset()      { *m = ISCSIState{} }
func (*ISCSIState) ProtoMessage() {}
func (*ISCSIState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fc248a35ef3cab8, []int{2}
}
func (m *ISCSIState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ISCSIState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ISCSIState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ISCSIState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ISCSIState.Merge(m, src)
}
func (m *ISCSIState) XXX_Size() int {
	return m.Size()
}
func (m *ISCSIState) XXX_DiscardUnknown() {
	xxx_messageInfo_ISCSIState.DiscardUnknown(m)
}

var xxx_messageInfo_ISCSIState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Target)(nil), "io.stellarproject.types.v1.Target")
	proto.RegisterType((*LUN)(nil), "io.stellarproject.types.v1.LUN")
	proto.RegisterType((*ISCSIState)(nil), "io.stellarproject.types.v1.ISCSIState")
	proto.RegisterMapType((map[string]*LUN)(nil), "io.stellarproject.types.v1.ISCSIState.UnallocatedLunsEntry")
}

func init() {
	proto.RegisterFile("github.com/stellarproject/terraos/api/types/v1/iscsi.proto", fileDescriptor_9fc248a35ef3cab8)
}

var fileDescriptor_9fc248a35ef3cab8 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0xa4, 0x6b, 0xd5, 0x97, 0x03, 0x93, 0x55, 0x81, 0xd5, 0x43, 0x56, 0xf5, 0xd4,
	0x93, 0xa3, 0x30, 0x21, 0xa1, 0x8d, 0xd3, 0x80, 0x43, 0x45, 0xb5, 0x83, 0x43, 0x25, 0xc4, 0x65,
	0xb8, 0xa9, 0x9b, 0x99, 0x99, 0x38, 0xc4, 0x4e, 0xa4, 0xec, 0x04, 0x9f, 0x82, 0xaf, 0xb4, 0x23,
	0x47, 0x4e, 0x88, 0xe5, 0x93, 0xa0, 0x38, 0x54, 0x53, 0x25, 0x98, 0xe0, 0xf6, 0xda, 0xfa, 0x3d,
	0xef, 0xf3, 0xfe, 0x83, 0x93, 0x54, 0xda, 0xcb, 0x72, 0x4d, 0x13, 0xfd, 0x31, 0x34, 0x56, 0x28,
	0xc5, 0x8b, 0xbc, 0xd0, 0x1f, 0x44, 0x62, 0x43, 0x2b, 0x8a, 0x82, 0x6b, 0x13, 0xf2, 0x5c, 0x86,
	0xb6, 0xce, 0x85, 0x09, 0xab, 0x28, 0x94, 0x26, 0x31, 0x92, 0xe6, 0x85, 0xb6, 0x1a, 0x4f, 0xa4,
	0xa6, 0xfb, 0x1a, 0xea, 0x38, 0x5a, 0x45, 0x93, 0x71, 0xaa, 0x53, 0xed, 0xb0, 0xb0, 0x8d, 0x3a,
	0xc5, 0xec, 0x0b, 0x82, 0xc1, 0x1b, 0x5e, 0xa4, 0xc2, 0xe2, 0x43, 0xf0, 0xe5, 0xa7, 0x8c, 0xa0,
	0x29, 0x9a, 0x8f, 0x58, 0x1b, 0xe2, 0x47, 0xe0, 0xc9, 0x0d, 0xf1, 0xa6, 0x68, 0xee, 0x9f, 0x0d,
	0x9a, 0x1f, 0x47, 0xde, 0xe2, 0x25, 0xf3, 0xe4, 0x06, 0x1f, 0x43, 0x5f, 0x95, 0x99, 0x21, 0xfe,
	0xd4, 0x9f, 0x3f, 0x78, 0x72, 0x44, 0xff, 0xee, 0x4a, 0x97, 0xab, 0x73, 0xe6, 0x60, 0x4c, 0x60,
	0xc8, 0x93, 0x44, 0xe4, 0xd6, 0x90, 0xfe, 0xd4, 0x9f, 0x8f, 0xd8, 0xee, 0x39, 0x7b, 0x0f, 0xfe,
	0x72, 0x75, 0xfe, 0xdb, 0xcd, 0xd9, 0xef, 0xb9, 0x61, 0xe8, 0xe7, 0xdc, 0x5e, 0xba, 0x3a, 0x46,
	0xcc, 0xc5, 0xf8, 0x31, 0x0c, 0xb7, 0xe6, 0xc2, 0xc8, 0x6b, 0x41, 0xfc, 0xb6, 0x3c, 0x36, 0xd8,
	0x9a, 0x58, 0x5e, 0x0b, 0x3c, 0x86, 0x03, 0xc5, 0xd7, 0x42, 0x91, 0xbe, 0xa3, 0xbb, 0xc7, 0xec,
	0xab, 0x07, 0xb0, 0x88, 0x5f, 0xc4, 0x8b, 0xd8, 0x72, 0x2b, 0xf0, 0x73, 0x18, 0x5a, 0xd7, 0xb3,
	0x21, 0xc8, 0xb5, 0x30, 0xbb, 0xaf, 0x85, 0x6e, 0x3c, 0x6c, 0x27, 0xc1, 0x5b, 0x38, 0x2c, 0x33,
	0xae, 0x94, 0x4e, 0xb8, 0x15, 0x9b, 0x0b, 0x37, 0x09, 0xcf, 0xa5, 0x39, 0xbd, 0x2f, 0xcd, 0x9d,
	0x3f, 0x5d, 0xdd, 0xc9, 0x97, 0x65, 0x66, 0x5e, 0x65, 0xb6, 0xa8, 0xd9, 0xc3, 0x72, 0xff, 0x77,
	0x92, 0xc0, 0xf8, 0x4f, 0x60, 0xbb, 0xa7, 0x2b, 0x51, 0xef, 0xf6, 0x74, 0x25, 0x6a, 0xfc, 0x14,
	0x0e, 0x2a, 0xae, 0x4a, 0xe1, 0x46, 0xf4, 0x0f, 0x0b, 0xe9, 0xe8, 0x13, 0xef, 0x19, 0x3a, 0x7b,
	0x7d, 0x73, 0x1b, 0xf4, 0xbe, 0xdf, 0x06, 0xbd, 0xcf, 0x4d, 0x80, 0x6e, 0x9a, 0x00, 0x7d, 0x6b,
	0x02, 0xf4, 0xb3, 0x09, 0xd0, 0xbb, 0xe8, 0xff, 0xee, 0xf0, 0xb4, 0x8a, 0xde, 0xf6, 0xd6, 0x03,
	0x77, 0x55, 0xc7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0x93, 0x16, 0xab, 0xc5, 0x02, 0x00,
	0x00,
}

func (m *Target) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Target) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Iqn) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIscsi(dAtA, i, uint64(len(m.Iqn)))
		i += copy(dAtA[i:], m.Iqn)
	}
	if m.ID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIscsi(dAtA, i, uint64(m.ID))
	}
	if len(m.Luns) > 0 {
		for _, msg := range m.Luns {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintIscsi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Accepts) > 0 {
		for _, s := range m.Accepts {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LUN) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LUN) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIscsi(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Path) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintIscsi(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if m.FsSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintIscsi(dAtA, i, uint64(m.FsSize))
	}
	if len(m.Label) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintIscsi(dAtA, i, uint64(len(m.Label)))
		i += copy(dAtA[i:], m.Label)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ISCSIState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ISCSIState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Targets) > 0 {
		for _, msg := range m.Targets {
			dAtA[i] = 0xa
			i++
			i = encodeVarintIscsi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.UnallocatedLuns) > 0 {
		for k, _ := range m.UnallocatedLuns {
			dAtA[i] = 0x12
			i++
			v := m.UnallocatedLuns[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovIscsi(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovIscsi(uint64(len(k))) + msgSize
			i = encodeVarintIscsi(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintIscsi(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintIscsi(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintIscsi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Target) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Iqn)
	if l > 0 {
		n += 1 + l + sovIscsi(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovIscsi(uint64(m.ID))
	}
	if len(m.Luns) > 0 {
		for _, e := range m.Luns {
			l = e.Size()
			n += 1 + l + sovIscsi(uint64(l))
		}
	}
	if len(m.Accepts) > 0 {
		for _, s := range m.Accepts {
			l = len(s)
			n += 1 + l + sovIscsi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LUN) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovIscsi(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovIscsi(uint64(l))
	}
	if m.FsSize != 0 {
		n += 1 + sovIscsi(uint64(m.FsSize))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovIscsi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ISCSIState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Targets) > 0 {
		for _, e := range m.Targets {
			l = e.Size()
			n += 1 + l + sovIscsi(uint64(l))
		}
	}
	if len(m.UnallocatedLuns) > 0 {
		for k, v := range m.UnallocatedLuns {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovIscsi(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovIscsi(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovIscsi(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIscsi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIscsi(x uint64) (n int) {
	return sovIscsi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Target) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Target{`,
		`Iqn:` + fmt.Sprintf("%v", this.Iqn) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Luns:` + strings.Replace(fmt.Sprintf("%v", this.Luns), "LUN", "LUN", 1) + `,`,
		`Accepts:` + fmt.Sprintf("%v", this.Accepts) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LUN) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LUN{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`FsSize:` + fmt.Sprintf("%v", this.FsSize) + `,`,
		`Label:` + fmt.Sprintf("%v", this.Label) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ISCSIState) String() string {
	if this == nil {
		return "nil"
	}
	keysForUnallocatedLuns := make([]string, 0, len(this.UnallocatedLuns))
	for k, _ := range this.UnallocatedLuns {
		keysForUnallocatedLuns = append(keysForUnallocatedLuns, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForUnallocatedLuns)
	mapStringForUnallocatedLuns := "map[string]*LUN{"
	for _, k := range keysForUnallocatedLuns {
		mapStringForUnallocatedLuns += fmt.Sprintf("%v: %v,", k, this.UnallocatedLuns[k])
	}
	mapStringForUnallocatedLuns += "}"
	s := strings.Join([]string{`&ISCSIState{`,
		`Targets:` + strings.Replace(fmt.Sprintf("%v", this.Targets), "Target", "Target", 1) + `,`,
		`UnallocatedLuns:` + mapStringForUnallocatedLuns + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIscsi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Target) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIscsi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Target: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Target: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Iqn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Iqn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Luns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Luns = append(m.Luns, &LUN{})
			if err := m.Luns[len(m.Luns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accepts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accepts = append(m.Accepts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIscsi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIscsi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIscsi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LUN) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIscsi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LUN: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LUN: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FsSize", wireType)
			}
			m.FsSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FsSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIscsi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIscsi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIscsi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ISCSIState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIscsi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ISCSIState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ISCSIState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Targets = append(m.Targets, &Target{})
			if err := m.Targets[len(m.Targets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnallocatedLuns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIscsi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIscsi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UnallocatedLuns == nil {
				m.UnallocatedLuns = make(map[string]*LUN)
			}
			var mapkey string
			var mapvalue *LUN
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIscsi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIscsi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthIscsi
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthIscsi
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIscsi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthIscsi
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthIscsi
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &LUN{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipIscsi(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthIscsi
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.UnallocatedLuns[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIscsi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIscsi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIscsi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIscsi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIscsi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIscsi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIscsi
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIscsi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIscsi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIscsi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIscsi
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIscsi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIscsi   = fmt.Errorf("proto: integer overflow")
)

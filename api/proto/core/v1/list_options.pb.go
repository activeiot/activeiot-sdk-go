// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/proto/core/v1/list_options.proto

package corev1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// ListOptions ...
// @param field_mask - Fields to be returned
// @param order - Order of the results (eg. ["ASC id"])
// @param limit - max amount of items to return
// @param offset - Amount of items to skip
// @param query - Query to execute, Base64 encoded json, https://docs.mongodb.com/manual/tutorial/query-documents/
type ListOptions struct {
	FieldMask *types.FieldMask `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	Order     []string         `protobuf:"bytes,2,rep,name=order,proto3" json:"order,omitempty"`
	Limit     int64            `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int64            `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Query     string           `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
}

func (m *ListOptions) Reset()      { *m = ListOptions{} }
func (*ListOptions) ProtoMessage() {}
func (*ListOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fc7cd9d9701bd0b, []int{0}
}
func (m *ListOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptions.Merge(m, src)
}
func (m *ListOptions) XXX_Size() int {
	return m.Size()
}
func (m *ListOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptions proto.InternalMessageInfo

func (m *ListOptions) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ListOptions) GetOrder() []string {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *ListOptions) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOptions) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListOptions) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*ListOptions)(nil), "activeiot.core.v1.ListOptions")
}

func init() {
	proto.RegisterFile("api/proto/core/v1/list_options.proto", fileDescriptor_1fc7cd9d9701bd0b)
}

var fileDescriptor_1fc7cd9d9701bd0b = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x33, 0xd4, 0xcf, 0xc9,
	0x2c, 0x2e, 0x89, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x03, 0x4b, 0x09, 0x09, 0x26,
	0x26, 0x97, 0x64, 0x96, 0xa5, 0x66, 0xe6, 0x97, 0xe8, 0x81, 0x54, 0xe9, 0x95, 0x19, 0x4a, 0x29,
	0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x42, 0xf4, 0x26, 0x95, 0xa6, 0xe9, 0xa7, 0x65, 0xa6, 0xe6,
	0xa4, 0xc4, 0xe7, 0x26, 0x16, 0x67, 0x43, 0x34, 0x29, 0x2d, 0x62, 0xe4, 0xe2, 0xf6, 0xc9, 0x2c,
	0x2e, 0xf1, 0x87, 0x18, 0x25, 0x64, 0xc9, 0xc5, 0x85, 0x50, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x6d, 0x24, 0xa5, 0x07, 0x31, 0x46, 0x0f, 0x66, 0x8c, 0x9e, 0x1b, 0x48, 0x89, 0x6f, 0x62, 0x71,
	0x76, 0x10, 0x67, 0x1a, 0x8c, 0x29, 0x24, 0xc2, 0xc5, 0x9a, 0x5f, 0x94, 0x92, 0x5a, 0x24, 0xc1,
	0xa4, 0xc0, 0xac, 0xc1, 0x19, 0x04, 0xe1, 0x80, 0x44, 0x73, 0x32, 0x73, 0x33, 0x4b, 0x24, 0x98,
	0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0xb6, 0xfc, 0xb4, 0xb4, 0xe2, 0xd4,
	0x12, 0x09, 0x16, 0xb0, 0x30, 0x94, 0x07, 0x52, 0x5d, 0x58, 0x9a, 0x5a, 0x54, 0x29, 0xc1, 0xaa,
	0xc0, 0x08, 0x32, 0x03, 0xcc, 0x71, 0x9a, 0xca, 0x78, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72,
	0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81,
	0x4b, 0x34, 0x39, 0x3f, 0x57, 0x0f, 0x23, 0x50, 0x9c, 0x04, 0x90, 0xfc, 0x1b, 0x00, 0xf2, 0x51,
	0x00, 0x63, 0x14, 0x1b, 0x48, 0xb2, 0xcc, 0x70, 0x11, 0x13, 0xb3, 0xa3, 0x73, 0xc4, 0x2a, 0x26,
	0x41, 0x47, 0xb8, 0x2e, 0x67, 0x90, 0xae, 0x30, 0xc3, 0x53, 0x48, 0x62, 0x31, 0x20, 0xb1, 0x98,
	0x30, 0xc3, 0x24, 0x36, 0x70, 0x80, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xf6, 0x33,
	0x7b, 0xa0, 0x01, 0x00, 0x00,
}

func (this *ListOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListOptions)
	if !ok {
		that2, ok := that.(ListOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.FieldMask.Equal(that1.FieldMask) {
		return false
	}
	if len(this.Order) != len(that1.Order) {
		return false
	}
	for i := range this.Order {
		if this.Order[i] != that1.Order[i] {
			return false
		}
	}
	if this.Limit != that1.Limit {
		return false
	}
	if this.Offset != that1.Offset {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	return true
}
func (this *ListOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&corev1.ListOptions{")
	if this.FieldMask != nil {
		s = append(s, "FieldMask: "+fmt.Sprintf("%#v", this.FieldMask)+",\n")
	}
	s = append(s, "Order: "+fmt.Sprintf("%#v", this.Order)+",\n")
	s = append(s, "Limit: "+fmt.Sprintf("%#v", this.Limit)+",\n")
	s = append(s, "Offset: "+fmt.Sprintf("%#v", this.Offset)+",\n")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringListOptions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FieldMask != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintListOptions(dAtA, i, uint64(m.FieldMask.Size()))
		n1, err := m.FieldMask.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Order) > 0 {
		for _, s := range m.Order {
			dAtA[i] = 0x12
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
	if m.Limit != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintListOptions(dAtA, i, uint64(m.Limit))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintListOptions(dAtA, i, uint64(m.Offset))
	}
	if len(m.Query) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintListOptions(dAtA, i, uint64(len(m.Query)))
		i += copy(dAtA[i:], m.Query)
	}
	return i, nil
}

func encodeVarintListOptions(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ListOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FieldMask != nil {
		l = m.FieldMask.Size()
		n += 1 + l + sovListOptions(uint64(l))
	}
	if len(m.Order) > 0 {
		for _, s := range m.Order {
			l = len(s)
			n += 1 + l + sovListOptions(uint64(l))
		}
	}
	if m.Limit != 0 {
		n += 1 + sovListOptions(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovListOptions(uint64(m.Offset))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovListOptions(uint64(l))
	}
	return n
}

func sovListOptions(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozListOptions(x uint64) (n int) {
	return sovListOptions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListOptions{`,
		`FieldMask:` + strings.Replace(fmt.Sprintf("%v", this.FieldMask), "FieldMask", "types.FieldMask", 1) + `,`,
		`Order:` + fmt.Sprintf("%v", this.Order) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Offset:` + fmt.Sprintf("%v", this.Offset) + `,`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringListOptions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListOptions
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
			return fmt.Errorf("proto: ListOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldMask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOptions
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
				return ErrInvalidLengthListOptions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FieldMask == nil {
				m.FieldMask = &types.FieldMask{}
			}
			if err := m.FieldMask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOptions
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
				return ErrInvalidLengthListOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Order = append(m.Order, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOptions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOptions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListOptions
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
				return ErrInvalidLengthListOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListOptions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListOptions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListOptions
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipListOptions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListOptions
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
					return 0, ErrIntOverflowListOptions
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
					return 0, ErrIntOverflowListOptions
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
				return 0, ErrInvalidLengthListOptions
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthListOptions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowListOptions
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
				next, err := skipListOptions(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthListOptions
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
	ErrInvalidLengthListOptions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListOptions   = fmt.Errorf("proto: integer overflow")
)

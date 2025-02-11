// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: db/continue_list_items_of_type.proto

package db

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ContinueListItemsOfTypeRequest) CloneVT() *ContinueListItemsOfTypeRequest {
	if m == nil {
		return (*ContinueListItemsOfTypeRequest)(nil)
	}
	r := new(ContinueListItemsOfTypeRequest)
	r.SchemaVersionId = m.SchemaVersionId
	if rhs := m.TokenData; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.TokenData = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ContinueListItemsOfTypeRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ContinueListItemsOfTypeRequest) EqualVT(that *ContinueListItemsOfTypeRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.TokenData) != string(that.TokenData) {
		return false
	}
	if this.SchemaVersionId != that.SchemaVersionId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ContinueListItemsOfTypeRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ContinueListItemsOfTypeRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *ContinueListItemsOfTypeRequest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinueListItemsOfTypeRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ContinueListItemsOfTypeRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.SchemaVersionId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SchemaVersionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TokenData) > 0 {
		i -= len(m.TokenData)
		copy(dAtA[i:], m.TokenData)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.TokenData)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContinueListItemsOfTypeRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenData)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SchemaVersionId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SchemaVersionId))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ContinueListItemsOfTypeRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: ContinueListItemsOfTypeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinueListItemsOfTypeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenData = append(m.TokenData[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenData == nil {
				m.TokenData = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersionId", wireType)
			}
			m.SchemaVersionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersionId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

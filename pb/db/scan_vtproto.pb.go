// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: db/scan.proto

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

func (m *BeginScanRequest) CloneVT() *BeginScanRequest {
	if m == nil {
		return (*BeginScanRequest)(nil)
	}
	r := new(BeginScanRequest)
	r.StoreId = m.StoreId
	r.Limit = m.Limit
	r.SegmentationParams = m.SegmentationParams.CloneVT()
	r.SchemaVersionId = m.SchemaVersionId
	r.SchemaId = m.SchemaId
	if rhs := m.FilterConditions; rhs != nil {
		tmpContainer := make([]*FilterCondition, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.FilterConditions = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *BeginScanRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SegmentationParams) CloneVT() *SegmentationParams {
	if m == nil {
		return (*SegmentationParams)(nil)
	}
	r := new(SegmentationParams)
	r.TotalSegments = m.TotalSegments
	r.SegmentIndex = m.SegmentIndex
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SegmentationParams) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *BeginScanRequest) EqualVT(that *BeginScanRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.StoreId != that.StoreId {
		return false
	}
	if len(this.FilterConditions) != len(that.FilterConditions) {
		return false
	}
	for i, vx := range this.FilterConditions {
		vy := that.FilterConditions[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &FilterCondition{}
			}
			if q == nil {
				q = &FilterCondition{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.Limit != that.Limit {
		return false
	}
	if !this.SegmentationParams.EqualVT(that.SegmentationParams) {
		return false
	}
	if this.SchemaVersionId != that.SchemaVersionId {
		return false
	}
	if this.SchemaId != that.SchemaId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *BeginScanRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*BeginScanRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SegmentationParams) EqualVT(that *SegmentationParams) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TotalSegments != that.TotalSegments {
		return false
	}
	if this.SegmentIndex != that.SegmentIndex {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SegmentationParams) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SegmentationParams)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *BeginScanRequest) MarshalVT() (dAtA []byte, err error) {
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

func (m *BeginScanRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BeginScanRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.SchemaId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SchemaId))
		i--
		dAtA[i] = 0x30
	}
	if m.SchemaVersionId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SchemaVersionId))
		i--
		dAtA[i] = 0x28
	}
	if m.SegmentationParams != nil {
		size, err := m.SegmentationParams.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.Limit != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FilterConditions) > 0 {
		for iNdEx := len(m.FilterConditions) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.FilterConditions[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StoreId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.StoreId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SegmentationParams) MarshalVT() (dAtA []byte, err error) {
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

func (m *SegmentationParams) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SegmentationParams) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.SegmentIndex != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SegmentIndex))
		i--
		dAtA[i] = 0x30
	}
	if m.TotalSegments != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.TotalSegments))
		i--
		dAtA[i] = 0x28
	}
	return len(dAtA) - i, nil
}

func (m *BeginScanRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoreId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.StoreId))
	}
	if len(m.FilterConditions) > 0 {
		for _, e := range m.FilterConditions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.Limit != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Limit))
	}
	if m.SegmentationParams != nil {
		l = m.SegmentationParams.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SchemaVersionId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SchemaVersionId))
	}
	if m.SchemaId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SchemaId))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SegmentationParams) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalSegments != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.TotalSegments))
	}
	if m.SegmentIndex != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SegmentIndex))
	}
	n += len(m.unknownFields)
	return n
}

func (m *BeginScanRequest) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: BeginScanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeginScanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreId", wireType)
			}
			m.StoreId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterConditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilterConditions = append(m.FilterConditions, &FilterCondition{})
			if err := m.FilterConditions[len(m.FilterConditions)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentationParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SegmentationParams == nil {
				m.SegmentationParams = &SegmentationParams{}
			}
			if err := m.SegmentationParams.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaId", wireType)
			}
			m.SchemaId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaId |= uint64(b&0x7F) << shift
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
func (m *SegmentationParams) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: SegmentationParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SegmentationParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSegments", wireType)
			}
			m.TotalSegments = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalSegments |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentIndex", wireType)
			}
			m.SegmentIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SegmentIndex |= uint32(b&0x7F) << shift
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

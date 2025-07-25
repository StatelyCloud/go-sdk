// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: db/list.proto

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

func (m *BeginListRequest) CloneVT() *BeginListRequest {
	if m == nil {
		return (*BeginListRequest)(nil)
	}
	r := new(BeginListRequest)
	r.StoreId = m.StoreId
	r.KeyPathPrefix = m.KeyPathPrefix
	r.Limit = m.Limit
	r.AllowStale = m.AllowStale
	r.SortProperty = m.SortProperty
	r.SortDirection = m.SortDirection
	r.SchemaVersionId = m.SchemaVersionId
	r.SchemaId = m.SchemaId
	if rhs := m.FilterConditions; rhs != nil {
		tmpContainer := make([]*FilterCondition, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.FilterConditions = tmpContainer
	}
	if rhs := m.KeyConditions; rhs != nil {
		tmpContainer := make([]*KeyCondition, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.KeyConditions = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *BeginListRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListResponse) CloneVT() *ListResponse {
	if m == nil {
		return (*ListResponse)(nil)
	}
	r := new(ListResponse)
	if m.Response != nil {
		r.Response = m.Response.(interface {
			CloneVT() isListResponse_Response
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListResponse_Result) CloneVT() isListResponse_Response {
	if m == nil {
		return (*ListResponse_Result)(nil)
	}
	r := new(ListResponse_Result)
	r.Result = m.Result.CloneVT()
	return r
}

func (m *ListResponse_Finished) CloneVT() isListResponse_Response {
	if m == nil {
		return (*ListResponse_Finished)(nil)
	}
	r := new(ListResponse_Finished)
	r.Finished = m.Finished.CloneVT()
	return r
}

func (m *ListPartialResult) CloneVT() *ListPartialResult {
	if m == nil {
		return (*ListPartialResult)(nil)
	}
	r := new(ListPartialResult)
	if rhs := m.Items; rhs != nil {
		tmpContainer := make([]*Item, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Items = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListPartialResult) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListFinished) CloneVT() *ListFinished {
	if m == nil {
		return (*ListFinished)(nil)
	}
	r := new(ListFinished)
	r.Token = m.Token.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListFinished) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *KeyCondition) CloneVT() *KeyCondition {
	if m == nil {
		return (*KeyCondition)(nil)
	}
	r := new(KeyCondition)
	r.KeyPath = m.KeyPath
	r.Operator = m.Operator
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *KeyCondition) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *BeginListRequest) EqualVT(that *BeginListRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.StoreId != that.StoreId {
		return false
	}
	if this.KeyPathPrefix != that.KeyPathPrefix {
		return false
	}
	if this.Limit != that.Limit {
		return false
	}
	if this.AllowStale != that.AllowStale {
		return false
	}
	if this.SortProperty != that.SortProperty {
		return false
	}
	if this.SortDirection != that.SortDirection {
		return false
	}
	if this.SchemaVersionId != that.SchemaVersionId {
		return false
	}
	if this.SchemaId != that.SchemaId {
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
	if len(this.KeyConditions) != len(that.KeyConditions) {
		return false
	}
	for i, vx := range this.KeyConditions {
		vy := that.KeyConditions[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &KeyCondition{}
			}
			if q == nil {
				q = &KeyCondition{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *BeginListRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*BeginListRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListResponse) EqualVT(that *ListResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Response == nil && that.Response != nil {
		return false
	} else if this.Response != nil {
		if that.Response == nil {
			return false
		}
		if !this.Response.(interface {
			EqualVT(isListResponse_Response) bool
		}).EqualVT(that.Response) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListResponse_Result) EqualVT(thatIface isListResponse_Response) bool {
	that, ok := thatIface.(*ListResponse_Result)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Result, that.Result; p != q {
		if p == nil {
			p = &ListPartialResult{}
		}
		if q == nil {
			q = &ListPartialResult{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ListResponse_Finished) EqualVT(thatIface isListResponse_Response) bool {
	that, ok := thatIface.(*ListResponse_Finished)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Finished, that.Finished; p != q {
		if p == nil {
			p = &ListFinished{}
		}
		if q == nil {
			q = &ListFinished{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ListPartialResult) EqualVT(that *ListPartialResult) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Items) != len(that.Items) {
		return false
	}
	for i, vx := range this.Items {
		vy := that.Items[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Item{}
			}
			if q == nil {
				q = &Item{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListPartialResult) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListPartialResult)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListFinished) EqualVT(that *ListFinished) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Token.EqualVT(that.Token) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListFinished) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListFinished)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *KeyCondition) EqualVT(that *KeyCondition) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.KeyPath != that.KeyPath {
		return false
	}
	if this.Operator != that.Operator {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KeyCondition) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KeyCondition)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *BeginListRequest) MarshalVT() (dAtA []byte, err error) {
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

func (m *BeginListRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BeginListRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.KeyConditions) > 0 {
		for iNdEx := len(m.KeyConditions) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.KeyConditions[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x52
		}
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
			dAtA[i] = 0x4a
		}
	}
	if m.SchemaId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SchemaId))
		i--
		dAtA[i] = 0x40
	}
	if m.SchemaVersionId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SchemaVersionId))
		i--
		dAtA[i] = 0x38
	}
	if m.SortDirection != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SortDirection))
		i--
		dAtA[i] = 0x30
	}
	if m.SortProperty != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.SortProperty))
		i--
		dAtA[i] = 0x28
	}
	if m.AllowStale {
		i--
		if m.AllowStale {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Limit != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.KeyPathPrefix) > 0 {
		i -= len(m.KeyPathPrefix)
		copy(dAtA[i:], m.KeyPathPrefix)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.KeyPathPrefix)))
		i--
		dAtA[i] = 0x12
	}
	if m.StoreId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.StoreId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Response.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *ListResponse_Result) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListResponse_Result) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Result != nil {
		size, err := m.Result.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ListResponse_Finished) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListResponse_Finished) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Finished != nil {
		size, err := m.Finished.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ListPartialResult) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListPartialResult) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListPartialResult) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Items[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListFinished) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListFinished) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListFinished) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Token != nil {
		size, err := m.Token.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeyCondition) MarshalVT() (dAtA []byte, err error) {
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

func (m *KeyCondition) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *KeyCondition) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Operator != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Operator))
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyPath) > 0 {
		i -= len(m.KeyPath)
		copy(dAtA[i:], m.KeyPath)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.KeyPath)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BeginListRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoreId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.StoreId))
	}
	l = len(m.KeyPathPrefix)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Limit))
	}
	if m.AllowStale {
		n += 2
	}
	if m.SortProperty != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SortProperty))
	}
	if m.SortDirection != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SortDirection))
	}
	if m.SchemaVersionId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SchemaVersionId))
	}
	if m.SchemaId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.SchemaId))
	}
	if len(m.FilterConditions) > 0 {
		for _, e := range m.FilterConditions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.KeyConditions) > 0 {
		for _, e := range m.KeyConditions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Response.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListResponse_Result) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *ListResponse_Finished) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Finished != nil {
		l = m.Finished.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *ListPartialResult) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListFinished) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Token != nil {
		l = m.Token.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *KeyCondition) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyPath)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Operator != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Operator))
	}
	n += len(m.unknownFields)
	return n
}

func (m *BeginListRequest) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: BeginListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeginListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field KeyPathPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyPathPrefix = string(dAtA[iNdEx:postIndex])
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowStale", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowStale = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortProperty", wireType)
			}
			m.SortProperty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SortProperty |= SortableProperty(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortDirection", wireType)
			}
			m.SortDirection = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SortDirection |= SortDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
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
		case 8:
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
		case 9:
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
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyConditions", wireType)
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
			m.KeyConditions = append(m.KeyConditions, &KeyCondition{})
			if err := m.KeyConditions[len(m.KeyConditions)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *ListResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
			if oneof, ok := m.Response.(*ListResponse_Result); ok {
				if err := oneof.Result.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &ListPartialResult{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Response = &ListResponse_Result{Result: v}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finished", wireType)
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
			if oneof, ok := m.Response.(*ListResponse_Finished); ok {
				if err := oneof.Finished.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &ListFinished{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Response = &ListResponse_Finished{Finished: v}
			}
			iNdEx = postIndex
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
func (m *ListPartialResult) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListPartialResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPartialResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
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
			m.Items = append(m.Items, &Item{})
			if err := m.Items[len(m.Items)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *ListFinished) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListFinished: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListFinished: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			if m.Token == nil {
				m.Token = &ListToken{}
			}
			if err := m.Token.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *KeyCondition) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: KeyCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			m.Operator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operator |= Operator(b&0x7F) << shift
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

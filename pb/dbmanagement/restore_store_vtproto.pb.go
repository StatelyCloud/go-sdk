// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: dbmanagement/restore_store.proto

package dbmanagement

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

func (m *RestoreStoreRequest) CloneVT() *RestoreStoreRequest {
	if m == nil {
		return (*RestoreStoreRequest)(nil)
	}
	r := new(RestoreStoreRequest)
	r.StoreId = m.StoreId
	r.PitRestore = m.PitRestore.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RestoreStoreRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *PointInTimeRestore) CloneVT() *PointInTimeRestore {
	if m == nil {
		return (*PointInTimeRestore)(nil)
	}
	r := new(PointInTimeRestore)
	if m.RestoreTime != nil {
		r.RestoreTime = m.RestoreTime.(interface {
			CloneVT() isPointInTimeRestore_RestoreTime
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PointInTimeRestore) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *PointInTimeRestore_RestoreTimeUnixMillis) CloneVT() isPointInTimeRestore_RestoreTime {
	if m == nil {
		return (*PointInTimeRestore_RestoreTimeUnixMillis)(nil)
	}
	r := new(PointInTimeRestore_RestoreTimeUnixMillis)
	r.RestoreTimeUnixMillis = m.RestoreTimeUnixMillis
	return r
}

func (m *PointInTimeRestore_UseLatest) CloneVT() isPointInTimeRestore_RestoreTime {
	if m == nil {
		return (*PointInTimeRestore_UseLatest)(nil)
	}
	r := new(PointInTimeRestore_UseLatest)
	r.UseLatest = m.UseLatest
	return r
}

func (m *RestoreStoreResponse) CloneVT() *RestoreStoreResponse {
	if m == nil {
		return (*RestoreStoreResponse)(nil)
	}
	r := new(RestoreStoreResponse)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RestoreStoreResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DescribeBackupRequest) CloneVT() *DescribeBackupRequest {
	if m == nil {
		return (*DescribeBackupRequest)(nil)
	}
	r := new(DescribeBackupRequest)
	r.StoreId = m.StoreId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DescribeBackupRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DescribeBackupResponse) CloneVT() *DescribeBackupResponse {
	if m == nil {
		return (*DescribeBackupResponse)(nil)
	}
	r := new(DescribeBackupResponse)
	r.BackupInfo = m.BackupInfo.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DescribeBackupResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *PointInTimeBackup) CloneVT() *PointInTimeBackup {
	if m == nil {
		return (*PointInTimeBackup)(nil)
	}
	r := new(PointInTimeBackup)
	r.BackupsAvailable = m.BackupsAvailable
	r.EarliestAvailableUnixMillis = m.EarliestAvailableUnixMillis
	r.LatestAvailableUnixMillis = m.LatestAvailableUnixMillis
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PointInTimeBackup) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *RestoreStoreRequest) EqualVT(that *RestoreStoreRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.StoreId != that.StoreId {
		return false
	}
	if !this.PitRestore.EqualVT(that.PitRestore) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RestoreStoreRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RestoreStoreRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *PointInTimeRestore) EqualVT(that *PointInTimeRestore) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RestoreTime == nil && that.RestoreTime != nil {
		return false
	} else if this.RestoreTime != nil {
		if that.RestoreTime == nil {
			return false
		}
		if !this.RestoreTime.(interface {
			EqualVT(isPointInTimeRestore_RestoreTime) bool
		}).EqualVT(that.RestoreTime) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PointInTimeRestore) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PointInTimeRestore)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *PointInTimeRestore_RestoreTimeUnixMillis) EqualVT(thatIface isPointInTimeRestore_RestoreTime) bool {
	that, ok := thatIface.(*PointInTimeRestore_RestoreTimeUnixMillis)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.RestoreTimeUnixMillis != that.RestoreTimeUnixMillis {
		return false
	}
	return true
}

func (this *PointInTimeRestore_UseLatest) EqualVT(thatIface isPointInTimeRestore_RestoreTime) bool {
	that, ok := thatIface.(*PointInTimeRestore_UseLatest)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.UseLatest != that.UseLatest {
		return false
	}
	return true
}

func (this *RestoreStoreResponse) EqualVT(that *RestoreStoreResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RestoreStoreResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RestoreStoreResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DescribeBackupRequest) EqualVT(that *DescribeBackupRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.StoreId != that.StoreId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DescribeBackupRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DescribeBackupRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DescribeBackupResponse) EqualVT(that *DescribeBackupResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.BackupInfo.EqualVT(that.BackupInfo) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DescribeBackupResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DescribeBackupResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *PointInTimeBackup) EqualVT(that *PointInTimeBackup) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.BackupsAvailable != that.BackupsAvailable {
		return false
	}
	if this.EarliestAvailableUnixMillis != that.EarliestAvailableUnixMillis {
		return false
	}
	if this.LatestAvailableUnixMillis != that.LatestAvailableUnixMillis {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PointInTimeBackup) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PointInTimeBackup)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *RestoreStoreRequest) MarshalVT() (dAtA []byte, err error) {
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

func (m *RestoreStoreRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RestoreStoreRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.PitRestore != nil {
		size, err := m.PitRestore.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
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

func (m *PointInTimeRestore) MarshalVT() (dAtA []byte, err error) {
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

func (m *PointInTimeRestore) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PointInTimeRestore) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.RestoreTime.(interface {
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

func (m *PointInTimeRestore_RestoreTimeUnixMillis) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PointInTimeRestore_RestoreTimeUnixMillis) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(m.RestoreTimeUnixMillis))
	i--
	dAtA[i] = 0x10
	return len(dAtA) - i, nil
}
func (m *PointInTimeRestore_UseLatest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PointInTimeRestore_UseLatest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.UseLatest {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	return len(dAtA) - i, nil
}
func (m *RestoreStoreResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *RestoreStoreResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RestoreStoreResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	return len(dAtA) - i, nil
}

func (m *DescribeBackupRequest) MarshalVT() (dAtA []byte, err error) {
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

func (m *DescribeBackupRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *DescribeBackupRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.StoreId != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.StoreId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DescribeBackupResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *DescribeBackupResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *DescribeBackupResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.BackupInfo != nil {
		size, err := m.BackupInfo.MarshalToSizedBufferVT(dAtA[:i])
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

func (m *PointInTimeBackup) MarshalVT() (dAtA []byte, err error) {
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

func (m *PointInTimeBackup) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PointInTimeBackup) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LatestAvailableUnixMillis != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.LatestAvailableUnixMillis))
		i--
		dAtA[i] = 0x18
	}
	if m.EarliestAvailableUnixMillis != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.EarliestAvailableUnixMillis))
		i--
		dAtA[i] = 0x10
	}
	if m.BackupsAvailable {
		i--
		if m.BackupsAvailable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RestoreStoreRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoreId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.StoreId))
	}
	if m.PitRestore != nil {
		l = m.PitRestore.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *PointInTimeRestore) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.RestoreTime.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *PointInTimeRestore_RestoreTimeUnixMillis) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + protohelpers.SizeOfVarint(uint64(m.RestoreTimeUnixMillis))
	return n
}
func (m *PointInTimeRestore_UseLatest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *RestoreStoreResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *DescribeBackupRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoreId != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.StoreId))
	}
	n += len(m.unknownFields)
	return n
}

func (m *DescribeBackupResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BackupInfo != nil {
		l = m.BackupInfo.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *PointInTimeBackup) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BackupsAvailable {
		n += 2
	}
	if m.EarliestAvailableUnixMillis != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.EarliestAvailableUnixMillis))
	}
	if m.LatestAvailableUnixMillis != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.LatestAvailableUnixMillis))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RestoreStoreRequest) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: RestoreStoreRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestoreStoreRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field PitRestore", wireType)
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
			if m.PitRestore == nil {
				m.PitRestore = &PointInTimeRestore{}
			}
			if err := m.PitRestore.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *PointInTimeRestore) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: PointInTimeRestore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PointInTimeRestore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestoreTimeUnixMillis", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RestoreTime = &PointInTimeRestore_RestoreTimeUnixMillis{RestoreTimeUnixMillis: v}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseLatest", wireType)
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
			b := bool(v != 0)
			m.RestoreTime = &PointInTimeRestore_UseLatest{UseLatest: b}
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
func (m *RestoreStoreResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: RestoreStoreResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestoreStoreResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *DescribeBackupRequest) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: DescribeBackupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescribeBackupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *DescribeBackupResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: DescribeBackupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescribeBackupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupInfo", wireType)
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
			if m.BackupInfo == nil {
				m.BackupInfo = &PointInTimeBackup{}
			}
			if err := m.BackupInfo.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *PointInTimeBackup) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: PointInTimeBackup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PointInTimeBackup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupsAvailable", wireType)
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
			m.BackupsAvailable = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarliestAvailableUnixMillis", wireType)
			}
			m.EarliestAvailableUnixMillis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EarliestAvailableUnixMillis |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestAvailableUnixMillis", wireType)
			}
			m.LatestAvailableUnixMillis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestAvailableUnixMillis |= uint64(b&0x7F) << shift
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

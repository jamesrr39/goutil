// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profiler.proto

package profile

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Run struct {
	Name           string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Summary        string   `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	StartTimeNanos int64    `protobuf:"varint,3,opt,name=start_time_nanos,json=startTimeNanos,proto3" json:"startTimeNanos"`
	EndTimeNanos   int64    `protobuf:"varint,4,opt,name=end_time_nanos,json=endTimeNanos,proto3" json:"endTimeNanos"`
	Events         []*Event `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *Run) Reset()      { *m = Run{} }
func (*Run) ProtoMessage() {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3dbcb9d8e2f391d, []int{0}
}
func (m *Run) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Run.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return m.Size()
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Run) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Run) GetStartTimeNanos() int64 {
	if m != nil {
		return m.StartTimeNanos
	}
	return 0
}

func (m *Run) GetEndTimeNanos() int64 {
	if m != nil {
		return m.EndTimeNanos
	}
	return 0
}

func (m *Run) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Event struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeNanos int64  `protobuf:"varint,2,opt,name=time_nanos,json=timeNanos,proto3" json:"timeNanos"`
}

func (m *Event) Reset()      { *m = Event{} }
func (*Event) ProtoMessage() {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3dbcb9d8e2f391d, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetTimeNanos() int64 {
	if m != nil {
		return m.TimeNanos
	}
	return 0
}

func init() {
	proto.RegisterType((*Run)(nil), "github.com.jamesrr39.goutil.profile.Run")
	proto.RegisterType((*Event)(nil), "github.com.jamesrr39.goutil.profile.Event")
}

func init() { proto.RegisterFile("profiler.proto", fileDescriptor_e3dbcb9d8e2f391d) }

var fileDescriptor_e3dbcb9d8e2f391d = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xe3, 0xa6, 0x7f, 0x6e, 0x7d, 0xab, 0xe8, 0xca, 0x53, 0x54, 0x5d, 0x39, 0x51, 0x61,
	0xa8, 0x90, 0x48, 0x25, 0x3a, 0x75, 0x60, 0x89, 0x54, 0xc6, 0x0e, 0x01, 0x31, 0xb0, 0x54, 0x29,
	0x75, 0x43, 0x50, 0x6d, 0x57, 0x8e, 0x83, 0xc4, 0xc6, 0x23, 0xb0, 0xf2, 0x06, 0x3c, 0x0a, 0x63,
	0xc7, 0x4e, 0x11, 0x75, 0x17, 0xc4, 0xd4, 0x47, 0x40, 0x31, 0x81, 0xa6, 0x4c, 0x6c, 0xfe, 0x7d,
	0xdf, 0x77, 0x8e, 0x8e, 0x8e, 0xa1, 0xb5, 0x10, 0x7c, 0x16, 0xcf, 0x89, 0xf0, 0x16, 0x82, 0x4b,
	0x8e, 0x0e, 0xa2, 0x58, 0xde, 0xa4, 0x13, 0xef, 0x9a, 0x53, 0xef, 0x36, 0xa4, 0x24, 0x11, 0xa2,
	0x3f, 0xf0, 0x22, 0x9e, 0xca, 0x78, 0xee, 0x15, 0xb7, 0xed, 0xe3, 0xdd, 0x51, 0x2f, 0xe2, 0x11,
	0xef, 0x69, 0xed, 0x24, 0x9d, 0x69, 0xd2, 0xa0, 0x5f, 0x9f, 0x9e, 0x9d, 0xa7, 0x0a, 0x34, 0x83,
	0x94, 0xa1, 0xff, 0xb0, 0xca, 0x42, 0x4a, 0x6c, 0xe0, 0x82, 0x6e, 0xd3, 0xff, 0xa3, 0x32, 0xa7,
	0x3a, 0x0a, 0x29, 0x09, 0xf4, 0x14, 0xd9, 0xb0, 0x91, 0xa4, 0x94, 0x86, 0xe2, 0xde, 0xae, 0xe4,
	0x07, 0xc1, 0x17, 0xa2, 0x11, 0xfc, 0x97, 0xc8, 0x50, 0xc8, 0xb1, 0x8c, 0x29, 0x19, 0xb3, 0x90,
	0xf1, 0xc4, 0x36, 0x5d, 0xd0, 0x35, 0xfd, 0x43, 0x95, 0x39, 0xd6, 0x79, 0xbe, 0xbb, 0x88, 0x29,
	0x19, 0xe5, 0x9b, 0xf7, 0xcc, 0xb1, 0x92, 0xbd, 0x49, 0xf0, 0x83, 0xd1, 0x19, 0xb4, 0x08, 0x9b,
	0x96, 0xdd, 0xaa, 0xda, 0xcd, 0x55, 0x99, 0xd3, 0x1a, 0xb2, 0x69, 0xd9, 0xab, 0x45, 0x4a, 0x1c,
	0xec, 0x11, 0xf2, 0x61, 0x9d, 0xdc, 0x11, 0x26, 0x13, 0xbb, 0xe6, 0x9a, 0xdd, 0xbf, 0x27, 0x47,
	0xde, 0x2f, 0xca, 0xf3, 0x86, 0xb9, 0x24, 0x28, 0x94, 0x9d, 0x4b, 0x58, 0xd3, 0x03, 0x84, 0xca,
	0xe5, 0x14, 0x95, 0x0c, 0x20, 0x2c, 0x85, 0xac, 0xe8, 0x90, 0x6d, 0x95, 0x39, 0xcd, 0x72, 0xc2,
	0xa6, 0xfc, 0x8e, 0xb7, 0x7b, 0xfa, 0xa7, 0xcb, 0x35, 0x36, 0x56, 0x6b, 0x6c, 0x6c, 0xd7, 0x18,
	0x3c, 0x28, 0x0c, 0x9e, 0x15, 0x06, 0x2f, 0x0a, 0x83, 0xa5, 0xc2, 0xe0, 0x55, 0x61, 0xf0, 0xa6,
	0xb0, 0xb1, 0x55, 0x18, 0x3c, 0x6e, 0xb0, 0xb1, 0xdc, 0x60, 0x63, 0xb5, 0xc1, 0xc6, 0x55, 0xa3,
	0x08, 0x39, 0xa9, 0xeb, 0x9f, 0xeb, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0xf2, 0x96, 0x6a,
	0x1f, 0x02, 0x00, 0x00,
}

func (this *Run) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Run)
	if !ok {
		that2, ok := that.(Run)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Summary != that1.Summary {
		return false
	}
	if this.StartTimeNanos != that1.StartTimeNanos {
		return false
	}
	if this.EndTimeNanos != that1.EndTimeNanos {
		return false
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(that1.Events[i]) {
			return false
		}
	}
	return true
}
func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
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
	if this.Name != that1.Name {
		return false
	}
	if this.TimeNanos != that1.TimeNanos {
		return false
	}
	return true
}
func (this *Run) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&profile.Run{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Summary: "+fmt.Sprintf("%#v", this.Summary)+",\n")
	s = append(s, "StartTimeNanos: "+fmt.Sprintf("%#v", this.StartTimeNanos)+",\n")
	s = append(s, "EndTimeNanos: "+fmt.Sprintf("%#v", this.EndTimeNanos)+",\n")
	if this.Events != nil {
		s = append(s, "Events: "+fmt.Sprintf("%#v", this.Events)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Event) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&profile.Event{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "TimeNanos: "+fmt.Sprintf("%#v", this.TimeNanos)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProfiler(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Run) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Run) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProfiler(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Summary) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProfiler(dAtA, i, uint64(len(m.Summary)))
		i += copy(dAtA[i:], m.Summary)
	}
	if m.StartTimeNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProfiler(dAtA, i, uint64(m.StartTimeNanos))
	}
	if m.EndTimeNanos != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProfiler(dAtA, i, uint64(m.EndTimeNanos))
	}
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintProfiler(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProfiler(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.TimeNanos != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProfiler(dAtA, i, uint64(m.TimeNanos))
	}
	return i, nil
}

func encodeVarintProfiler(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Run) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProfiler(uint64(l))
	}
	l = len(m.Summary)
	if l > 0 {
		n += 1 + l + sovProfiler(uint64(l))
	}
	if m.StartTimeNanos != 0 {
		n += 1 + sovProfiler(uint64(m.StartTimeNanos))
	}
	if m.EndTimeNanos != 0 {
		n += 1 + sovProfiler(uint64(m.EndTimeNanos))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovProfiler(uint64(l))
		}
	}
	return n
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProfiler(uint64(l))
	}
	if m.TimeNanos != 0 {
		n += 1 + sovProfiler(uint64(m.TimeNanos))
	}
	return n
}

func sovProfiler(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProfiler(x uint64) (n int) {
	return sovProfiler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Run) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEvents := "[]*Event{"
	for _, f := range this.Events {
		repeatedStringForEvents += strings.Replace(f.String(), "Event", "Event", 1) + ","
	}
	repeatedStringForEvents += "}"
	s := strings.Join([]string{`&Run{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Summary:` + fmt.Sprintf("%v", this.Summary) + `,`,
		`StartTimeNanos:` + fmt.Sprintf("%v", this.StartTimeNanos) + `,`,
		`EndTimeNanos:` + fmt.Sprintf("%v", this.EndTimeNanos) + `,`,
		`Events:` + repeatedStringForEvents + `,`,
		`}`,
	}, "")
	return s
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`TimeNanos:` + fmt.Sprintf("%v", this.TimeNanos) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProfiler(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Run) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProfiler
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
			return fmt.Errorf("proto: Run: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Run: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
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
				return ErrInvalidLengthProfiler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProfiler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Summary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
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
				return ErrInvalidLengthProfiler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProfiler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Summary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimeNanos", wireType)
			}
			m.StartTimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTimeNanos |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTimeNanos", wireType)
			}
			m.EndTimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTimeNanos |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
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
				return ErrInvalidLengthProfiler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProfiler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProfiler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProfiler
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProfiler
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProfiler
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
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
				return ErrInvalidLengthProfiler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProfiler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeNanos", wireType)
			}
			m.TimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeNanos |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProfiler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProfiler
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProfiler
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
func skipProfiler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProfiler
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
					return 0, ErrIntOverflowProfiler
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
					return 0, ErrIntOverflowProfiler
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
				return 0, ErrInvalidLengthProfiler
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthProfiler
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProfiler
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
				next, err := skipProfiler(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthProfiler
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
	ErrInvalidLengthProfiler = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProfiler   = fmt.Errorf("proto: integer overflow")
)

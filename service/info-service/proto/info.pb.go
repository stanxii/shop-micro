// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/info.proto

/*
	Package shop_srv_info is a generated protocol buffer package.

	It is generated from these files:
		proto/info.proto

	It has these top-level messages:
		Request
		VideoDetailReq
		InfoListReq
		VideoListResp
		NewsListResp
		VideoCategorysResp
		NewsCategorysResp
*/
package shop_srv_info

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shop_srv_shopproto "shop-micro/shopproto/news"
import shop_srv_shopproto1 "shop-micro/shopproto/video"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{0} }

type VideoDetailReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *VideoDetailReq) Reset()                    { *m = VideoDetailReq{} }
func (m *VideoDetailReq) String() string            { return proto.CompactTextString(m) }
func (*VideoDetailReq) ProtoMessage()               {}
func (*VideoDetailReq) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{1} }

func (m *VideoDetailReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type InfoListReq struct {
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	PageNum  int32  `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (m *InfoListReq) Reset()                    { *m = InfoListReq{} }
func (m *InfoListReq) String() string            { return proto.CompactTextString(m) }
func (*InfoListReq) ProtoMessage()               {}
func (*InfoListReq) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{2} }

func (m *InfoListReq) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *InfoListReq) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *InfoListReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type VideoListResp struct {
	VideoList []*shop_srv_shopproto1.Video `protobuf:"bytes,1,rep,name=videoList" json:"videoList,omitempty"`
}

func (m *VideoListResp) Reset()                    { *m = VideoListResp{} }
func (m *VideoListResp) String() string            { return proto.CompactTextString(m) }
func (*VideoListResp) ProtoMessage()               {}
func (*VideoListResp) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{3} }

func (m *VideoListResp) GetVideoList() []*shop_srv_shopproto1.Video {
	if m != nil {
		return m.VideoList
	}
	return nil
}

type NewsListResp struct {
	NewsList []*shop_srv_shopproto.News `protobuf:"bytes,1,rep,name=newsList" json:"newsList,omitempty"`
}

func (m *NewsListResp) Reset()                    { *m = NewsListResp{} }
func (m *NewsListResp) String() string            { return proto.CompactTextString(m) }
func (*NewsListResp) ProtoMessage()               {}
func (*NewsListResp) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{4} }

func (m *NewsListResp) GetNewsList() []*shop_srv_shopproto.News {
	if m != nil {
		return m.NewsList
	}
	return nil
}

type VideoCategorysResp struct {
	VideoCategoryList []*shop_srv_shopproto1.VideoCategory `protobuf:"bytes,1,rep,name=VideoCategoryList" json:"VideoCategoryList,omitempty"`
}

func (m *VideoCategorysResp) Reset()                    { *m = VideoCategorysResp{} }
func (m *VideoCategorysResp) String() string            { return proto.CompactTextString(m) }
func (*VideoCategorysResp) ProtoMessage()               {}
func (*VideoCategorysResp) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{5} }

func (m *VideoCategorysResp) GetVideoCategoryList() []*shop_srv_shopproto1.VideoCategory {
	if m != nil {
		return m.VideoCategoryList
	}
	return nil
}

type NewsCategorysResp struct {
	NewsCategoryList []*shop_srv_shopproto.NewsCategory `protobuf:"bytes,1,rep,name=newsCategoryList" json:"newsCategoryList,omitempty"`
}

func (m *NewsCategorysResp) Reset()                    { *m = NewsCategorysResp{} }
func (m *NewsCategorysResp) String() string            { return proto.CompactTextString(m) }
func (*NewsCategorysResp) ProtoMessage()               {}
func (*NewsCategorysResp) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{6} }

func (m *NewsCategorysResp) GetNewsCategoryList() []*shop_srv_shopproto.NewsCategory {
	if m != nil {
		return m.NewsCategoryList
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "shop.srv.info.request")
	proto.RegisterType((*VideoDetailReq)(nil), "shop.srv.info.VideoDetailReq")
	proto.RegisterType((*InfoListReq)(nil), "shop.srv.info.InfoListReq")
	proto.RegisterType((*VideoListResp)(nil), "shop.srv.info.VideoListResp")
	proto.RegisterType((*NewsListResp)(nil), "shop.srv.info.NewsListResp")
	proto.RegisterType((*VideoCategorysResp)(nil), "shop.srv.info.VideoCategorysResp")
	proto.RegisterType((*NewsCategorysResp)(nil), "shop.srv.info.NewsCategorysResp")
}
func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *VideoDetailReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VideoDetailReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInfo(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *InfoListReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoListReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Category) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Category)))
		i += copy(dAtA[i:], m.Category)
	}
	if m.PageNum != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintInfo(dAtA, i, uint64(m.PageNum))
	}
	if m.PageSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintInfo(dAtA, i, uint64(m.PageSize))
	}
	return i, nil
}

func (m *VideoListResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VideoListResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VideoList) > 0 {
		for _, msg := range m.VideoList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NewsListResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewsListResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NewsList) > 0 {
		for _, msg := range m.NewsList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *VideoCategorysResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VideoCategorysResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VideoCategoryList) > 0 {
		for _, msg := range m.VideoCategoryList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NewsCategorysResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewsCategorysResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NewsCategoryList) > 0 {
		for _, msg := range m.NewsCategoryList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *VideoDetailReq) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovInfo(uint64(m.Id))
	}
	return n
}

func (m *InfoListReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.PageNum != 0 {
		n += 1 + sovInfo(uint64(m.PageNum))
	}
	if m.PageSize != 0 {
		n += 1 + sovInfo(uint64(m.PageSize))
	}
	return n
}

func (m *VideoListResp) Size() (n int) {
	var l int
	_ = l
	if len(m.VideoList) > 0 {
		for _, e := range m.VideoList {
			l = e.Size()
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	return n
}

func (m *NewsListResp) Size() (n int) {
	var l int
	_ = l
	if len(m.NewsList) > 0 {
		for _, e := range m.NewsList {
			l = e.Size()
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	return n
}

func (m *VideoCategorysResp) Size() (n int) {
	var l int
	_ = l
	if len(m.VideoCategoryList) > 0 {
		for _, e := range m.VideoCategoryList {
			l = e.Size()
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	return n
}

func (m *NewsCategorysResp) Size() (n int) {
	var l int
	_ = l
	if len(m.NewsCategoryList) > 0 {
		for _, e := range m.NewsCategoryList {
			l = e.Size()
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	return n
}

func sovInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func (m *VideoDetailReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VideoDetailReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VideoDetailReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func (m *InfoListReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoListReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoListReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageNum", wireType)
			}
			m.PageNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageNum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageSize", wireType)
			}
			m.PageSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func (m *VideoListResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VideoListResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VideoListResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VideoList = append(m.VideoList, &shop_srv_shopproto1.Video{})
			if err := m.VideoList[len(m.VideoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func (m *NewsListResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NewsListResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewsListResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewsList = append(m.NewsList, &shop_srv_shopproto.News{})
			if err := m.NewsList[len(m.NewsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func (m *VideoCategorysResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VideoCategorysResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VideoCategorysResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoCategoryList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VideoCategoryList = append(m.VideoCategoryList, &shop_srv_shopproto1.VideoCategory{})
			if err := m.VideoCategoryList[len(m.VideoCategoryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func (m *NewsCategorysResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NewsCategorysResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewsCategorysResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewsCategoryList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewsCategoryList = append(m.NewsCategoryList, &shop_srv_shopproto.NewsCategory{})
			if err := m.NewsCategoryList[len(m.NewsCategoryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInfo
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
				next, err := skipInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto/info.proto", fileDescriptorInfo) }

var fileDescriptorInfo = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x13, 0xa0, 0xcd, 0xa4, 0x0d, 0xed, 0x82, 0x90, 0x31, 0x60, 0xb9, 0x2b, 0x84, 0x72,
	0xe9, 0x46, 0x2a, 0x48, 0xdc, 0xa1, 0x52, 0x0a, 0x8a, 0x82, 0x30, 0x12, 0x57, 0x94, 0x9f, 0x8d,
	0xbb, 0x12, 0xf1, 0xba, 0xde, 0x75, 0x2a, 0x78, 0x07, 0xee, 0x3c, 0x12, 0xc7, 0x3e, 0x02, 0x84,
	0x17, 0x41, 0xbb, 0xb6, 0xb7, 0x8e, 0x7f, 0xc4, 0xc5, 0x9a, 0xf9, 0xf6, 0xfb, 0xbe, 0x19, 0xcf,
	0x0c, 0x1c, 0xc5, 0x09, 0x97, 0x7c, 0xc4, 0xa2, 0x15, 0x27, 0x3a, 0x44, 0x87, 0xe2, 0x92, 0xc7,
	0x44, 0x24, 0x1b, 0xa2, 0x40, 0xf7, 0xb9, 0x4a, 0x4f, 0xd7, 0x6c, 0x91, 0xf0, 0x91, 0x0a, 0x33,
	0x7e, 0x44, 0xaf, 0x85, 0xfe, 0x64, 0x22, 0xf7, 0x45, 0x23, 0x6b, 0xc3, 0x96, 0x34, 0xff, 0xe6,
	0xbc, 0xd3, 0x90, 0xc9, 0xcb, 0x74, 0x4e, 0x16, 0x7c, 0x3d, 0x0a, 0x79, 0xc8, 0x47, 0x1a, 0x9e,
	0xa7, 0x2b, 0x9d, 0x65, 0x2a, 0x15, 0x65, 0x74, 0xdc, 0x83, 0xbd, 0x84, 0x5e, 0xa5, 0x54, 0x48,
	0xec, 0xc3, 0xe0, 0xb3, 0x32, 0x3a, 0xa7, 0x72, 0xc6, 0xbe, 0x06, 0xf4, 0x0a, 0x0d, 0xa0, 0xc3,
	0x96, 0x8e, 0xed, 0xdb, 0xc3, 0x6e, 0xd0, 0x61, 0x4b, 0xfc, 0x05, 0xfa, 0xef, 0xa2, 0x15, 0x9f,
	0x30, 0x21, 0xd5, 0xb3, 0x0b, 0xfb, 0x8b, 0x99, 0xa4, 0x21, 0x4f, 0xbe, 0x69, 0x52, 0x2f, 0x30,
	0x39, 0x72, 0x60, 0x2f, 0x9e, 0x85, 0x74, 0x9a, 0xae, 0x9d, 0x8e, 0x6f, 0x0f, 0xef, 0x06, 0x45,
	0xaa, 0x54, 0x2a, 0xfc, 0xc4, 0xbe, 0x53, 0xa7, 0xab, 0x9f, 0x4c, 0x8e, 0x2f, 0xe0, 0x50, 0xb7,
	0x90, 0x55, 0x10, 0x31, 0x7a, 0x0d, 0xbd, 0x4d, 0x01, 0x38, 0xb6, 0xdf, 0x1d, 0xf6, 0xcf, 0x1e,
	0x13, 0x33, 0x3e, 0x33, 0x07, 0xa2, 0x55, 0xc1, 0x2d, 0x17, 0x9f, 0xc3, 0xc1, 0x94, 0x5e, 0x0b,
	0x63, 0xf4, 0x0a, 0xf6, 0xa3, 0x3c, 0xcf, 0x7d, 0x9c, 0x26, 0x1f, 0xa5, 0x09, 0x0c, 0x13, 0x53,
	0x40, 0xda, 0xf9, 0x6d, 0xfe, 0x5b, 0x42, 0x7b, 0x7d, 0x80, 0xe3, 0x1d, 0xb4, 0x64, 0x7a, 0xd2,
	0xda, 0x5c, 0x41, 0x0e, 0xea, 0x5a, 0x3c, 0x83, 0x63, 0x55, 0x78, 0xb7, 0xca, 0x04, 0x8e, 0xa2,
	0x12, 0x58, 0x2a, 0xe2, 0xb7, 0x75, 0x6e, 0x6a, 0xd4, 0x94, 0x67, 0x3f, 0xba, 0x70, 0x47, 0xed,
	0x0e, 0x4d, 0xe1, 0xfe, 0x98, 0xca, 0xf2, 0x1b, 0x7a, 0x44, 0x76, 0x0e, 0x92, 0xe4, 0x07, 0xe1,
	0x9e, 0x54, 0xf0, 0xfa, 0x28, 0xb0, 0x85, 0x26, 0x30, 0x18, 0x53, 0x59, 0x3a, 0x1c, 0xf4, 0xac,
	0x49, 0x66, 0x8e, 0xca, 0x6d, 0xdf, 0x1f, 0xb6, 0xd0, 0x7b, 0x38, 0x28, 0xdc, 0x74, 0x6b, 0x6e,
	0xc5, 0xab, 0x74, 0x7e, 0xee, 0xd3, 0xa6, 0x3a, 0xc5, 0xc2, 0xb1, 0x85, 0x3e, 0xc2, 0x83, 0x31,
	0x95, 0xd3, 0xca, 0x24, 0x5a, 0xff, 0xd6, 0xaf, 0xe0, 0xb5, 0x8d, 0x60, 0x0b, 0x5d, 0x40, 0x3f,
	0xb7, 0xfc, 0x6f, 0x77, 0x4f, 0x1a, 0xec, 0x6e, 0x9b, 0x7b, 0xf3, 0xf0, 0xe6, 0x8f, 0x67, 0xfd,
	0xda, 0x7a, 0xf6, 0xcd, 0xd6, 0xb3, 0x7f, 0x6f, 0x3d, 0xfb, 0xe7, 0x5f, 0xcf, 0x9a, 0xdf, 0xd3,
	0xd3, 0x78, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x93, 0x32, 0x05, 0x93, 0x34, 0x04, 0x00, 0x00,
}
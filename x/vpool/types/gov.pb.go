// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vpool/v1/gov.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CreatePoolProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// pair represents the pair of the vpool.
	Pair string `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	// quote_asset_reserve is the amount of quote asset the pool will be initialized with.
	QuoteAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=quote_asset_reserve,json=quoteAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quote_asset_reserve"`
	// base_asset_reserve is the amount of base asset the pool will be initialized with.
	BaseAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=base_asset_reserve,json=baseAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_asset_reserve"`
	Config           VpoolConfig                            `protobuf:"bytes,7,opt,name=config,proto3" json:"config"`
}

func (m *CreatePoolProposal) Reset()         { *m = CreatePoolProposal{} }
func (m *CreatePoolProposal) String() string { return proto.CompactTextString(m) }
func (*CreatePoolProposal) ProtoMessage()    {}
func (*CreatePoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a393460ab414204, []int{0}
}
func (m *CreatePoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePoolProposal.Merge(m, src)
}
func (m *CreatePoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreatePoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePoolProposal proto.InternalMessageInfo

func (m *CreatePoolProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePoolProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreatePoolProposal) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func (m *CreatePoolProposal) GetConfig() VpoolConfig {
	if m != nil {
		return m.Config
	}
	return VpoolConfig{}
}

type EditPoolConfigProposal struct {
	Title       string      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Pair        string      `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	Config      VpoolConfig `protobuf:"bytes,4,opt,name=config,proto3" json:"config"`
}

func (m *EditPoolConfigProposal) Reset()         { *m = EditPoolConfigProposal{} }
func (m *EditPoolConfigProposal) String() string { return proto.CompactTextString(m) }
func (*EditPoolConfigProposal) ProtoMessage()    {}
func (*EditPoolConfigProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a393460ab414204, []int{1}
}
func (m *EditPoolConfigProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditPoolConfigProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditPoolConfigProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditPoolConfigProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPoolConfigProposal.Merge(m, src)
}
func (m *EditPoolConfigProposal) XXX_Size() int {
	return m.Size()
}
func (m *EditPoolConfigProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPoolConfigProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EditPoolConfigProposal proto.InternalMessageInfo

func (m *EditPoolConfigProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EditPoolConfigProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EditPoolConfigProposal) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func (m *EditPoolConfigProposal) GetConfig() VpoolConfig {
	if m != nil {
		return m.Config
	}
	return VpoolConfig{}
}

// EditSwapInvariantsProposal is a governance proposal to change the swap
// invariant of the virtual pool for one or more trading pairs.
type EditSwapInvariantsProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Map from pair ID to a multiple on the swap invariant. For example, a proposal containing
	// "swap_invariant_maps": [{ "uatom:unusd": "5" }, { "uosmo:unusd": "0.9" }]
	// would mutliply the swap invariant of the ATOM and OSMO trading pairs by
	// 5 and 0.9 respectively. The price at which k changes is the instantaneous
	// mark price at the time of the proposal's execution.
	SwapInvariantMaps []EditSwapInvariantsProposal_SwapInvariantMultiple `protobuf:"bytes,5,rep,name=swap_invariant_maps,json=swapInvariantMaps,proto3" json:"swap_invariant_maps"`
}

func (m *EditSwapInvariantsProposal) Reset()         { *m = EditSwapInvariantsProposal{} }
func (m *EditSwapInvariantsProposal) String() string { return proto.CompactTextString(m) }
func (*EditSwapInvariantsProposal) ProtoMessage()    {}
func (*EditSwapInvariantsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a393460ab414204, []int{2}
}
func (m *EditSwapInvariantsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditSwapInvariantsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditSwapInvariantsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditSwapInvariantsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditSwapInvariantsProposal.Merge(m, src)
}
func (m *EditSwapInvariantsProposal) XXX_Size() int {
	return m.Size()
}
func (m *EditSwapInvariantsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EditSwapInvariantsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EditSwapInvariantsProposal proto.InternalMessageInfo

func (m *EditSwapInvariantsProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EditSwapInvariantsProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EditSwapInvariantsProposal) GetSwapInvariantMaps() []EditSwapInvariantsProposal_SwapInvariantMultiple {
	if m != nil {
		return m.SwapInvariantMaps
	}
	return nil
}

// A map between a trading pair and a desired multiplier for its swap invariant.
type EditSwapInvariantsProposal_SwapInvariantMultiple struct {
	// Pair is a string identifier for an asset pair.
	Pair string `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	// Multiplier is a number representing the desired percentage change to the
	// swap invariant of the AMM pool underlying 'pair'
	Multiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"multiplier"`
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Reset() {
	*m = EditSwapInvariantsProposal_SwapInvariantMultiple{}
}
func (*EditSwapInvariantsProposal_SwapInvariantMultiple) ProtoMessage() {}
func (*EditSwapInvariantsProposal_SwapInvariantMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a393460ab414204, []int{2, 0}
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple.Merge(m, src)
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_Size() int {
	return m.Size()
}
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_EditSwapInvariantsProposal_SwapInvariantMultiple proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreatePoolProposal)(nil), "nibiru.vpool.v1.CreatePoolProposal")
	proto.RegisterType((*EditPoolConfigProposal)(nil), "nibiru.vpool.v1.EditPoolConfigProposal")
	proto.RegisterType((*EditSwapInvariantsProposal)(nil), "nibiru.vpool.v1.EditSwapInvariantsProposal")
	proto.RegisterType((*EditSwapInvariantsProposal_SwapInvariantMultiple)(nil), "nibiru.vpool.v1.EditSwapInvariantsProposal.SwapInvariantMultiple")
}

func init() { proto.RegisterFile("vpool/v1/gov.proto", fileDescriptor_8a393460ab414204) }

var fileDescriptor_8a393460ab414204 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xed, 0x34, 0x09, 0x70, 0x19, 0xa0, 0xd7, 0x80, 0xac, 0x08, 0x39, 0x51, 0x07, 0x14,
	0x09, 0x71, 0xa7, 0x96, 0xad, 0x5b, 0x13, 0x3a, 0x30, 0x50, 0x45, 0x41, 0x62, 0x40, 0x88, 0xe8,
	0xe2, 0x1c, 0xee, 0x09, 0xc7, 0xef, 0xb8, 0x77, 0x71, 0xe0, 0x03, 0x20, 0x31, 0x32, 0x32, 0xe6,
	0xe3, 0x74, 0xec, 0x82, 0x84, 0x3a, 0x54, 0x28, 0xf9, 0x22, 0xc8, 0xe7, 0x50, 0x92, 0x52, 0x86,
	0x80, 0x98, 0xfc, 0xfc, 0xfe, 0xef, 0xfd, 0xee, 0xef, 0xbf, 0x7c, 0x84, 0x66, 0x1a, 0x20, 0xe1,
	0xd9, 0x1e, 0x8f, 0x21, 0x63, 0xda, 0x80, 0x05, 0x7a, 0x3b, 0x55, 0x43, 0x65, 0x26, 0xcc, 0x49,
	0x2c, 0xdb, 0x6b, 0xd4, 0x63, 0x88, 0xc1, 0x69, 0x3c, 0xaf, 0x8a, 0xb1, 0x46, 0xfd, 0x72, 0x15,
	0xad, 0xb0, 0xb2, 0xe8, 0xee, 0x9e, 0x97, 0x08, 0xed, 0x1a, 0x29, 0xac, 0xec, 0x01, 0x24, 0x3d,
	0x03, 0x1a, 0x50, 0x24, 0xb4, 0x4e, 0x2a, 0x56, 0xd9, 0x44, 0x06, 0x7e, 0xcb, 0x6f, 0xdf, 0xea,
	0x17, 0x2f, 0xb4, 0x45, 0x6a, 0x23, 0x89, 0x91, 0x51, 0xda, 0x2a, 0x48, 0x83, 0x92, 0xd3, 0x56,
	0x5b, 0x94, 0x92, 0xb2, 0x16, 0xca, 0x04, 0x5b, 0x4e, 0x72, 0x35, 0x7d, 0x4d, 0x76, 0xde, 0x4d,
	0xc0, 0xca, 0x81, 0x40, 0x94, 0x76, 0x60, 0x24, 0x4a, 0x93, 0xc9, 0xa0, 0x92, 0x8f, 0x74, 0xd8,
	0xe9, 0x45, 0xd3, 0x3b, 0xbf, 0x68, 0x3e, 0x88, 0x95, 0x3d, 0x99, 0x0c, 0x59, 0x04, 0x63, 0x1e,
	0x01, 0x8e, 0x01, 0x97, 0x8f, 0x47, 0x38, 0x7a, 0xcb, 0xed, 0x07, 0x2d, 0x91, 0x3d, 0x91, 0x51,
	0x7f, 0xdb, 0xa1, 0x0e, 0x73, 0x52, 0xbf, 0x00, 0xd1, 0x57, 0x84, 0x0e, 0x05, 0x5e, 0xc5, 0x57,
	0xff, 0x0a, 0x7f, 0x27, 0x27, 0xad, 0xd1, 0x0f, 0x48, 0x35, 0x82, 0xf4, 0x8d, 0x8a, 0x83, 0x1b,
	0x2d, 0xbf, 0x5d, 0xdb, 0xbf, 0xcf, 0xae, 0xc4, 0xcd, 0x5e, 0xe4, 0x45, 0xd7, 0xcd, 0x74, 0xca,
	0xf9, 0x79, 0xfd, 0xe5, 0xc6, 0xee, 0xcc, 0x27, 0xf7, 0x8e, 0x46, 0xca, 0xf6, 0x2e, 0x07, 0xfe,
	0x4b, 0xc0, 0xbf, 0x2c, 0x96, 0x37, 0xb6, 0xf8, 0xb5, 0x44, 0x1a, 0xb9, 0xc5, 0xe7, 0x53, 0xa1,
	0x9f, 0xa6, 0x99, 0x30, 0x4a, 0xa4, 0x16, 0xff, 0xd9, 0xe6, 0x94, 0xec, 0xe0, 0x54, 0xe8, 0x81,
	0xfa, 0x89, 0x1c, 0x8c, 0x85, 0xc6, 0xa0, 0xd2, 0xda, 0x6a, 0xd7, 0xf6, 0x0f, 0x7f, 0xf3, 0xf7,
	0x67, 0x07, 0x6c, 0xad, 0xfd, 0x6c, 0x92, 0x58, 0xa5, 0x13, 0xb9, 0xfc, 0x88, 0x6d, 0x5c, 0x13,
	0x85, 0xc6, 0xc6, 0x47, 0x9f, 0xdc, 0xbd, 0x76, 0xe5, 0xda, 0xe4, 0x8e, 0x09, 0x19, 0x17, 0xba,
	0x92, 0xc6, 0xa5, 0xb7, 0xf9, 0x2f, 0xb3, 0x42, 0x38, 0xb8, 0xf9, 0x69, 0xd6, 0xf4, 0xbe, 0xcc,
	0x9a, 0x5e, 0xe7, 0xe8, 0x74, 0x1e, 0xfa, 0x67, 0xf3, 0xd0, 0xff, 0x3e, 0x0f, 0xfd, 0xcf, 0x8b,
	0xd0, 0x3b, 0x5b, 0x84, 0xde, 0xb7, 0x45, 0xe8, 0xbd, 0x7c, 0xb8, 0xc2, 0x3d, 0x76, 0x39, 0x74,
	0x4f, 0x84, 0x4a, 0x79, 0x91, 0x09, 0x7f, 0xcf, 0x8b, 0x7b, 0xea, 0x0e, 0x18, 0x56, 0xdd, 0x2d,
	0x7d, 0xfc, 0x23, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xee, 0x03, 0xc5, 0xf8, 0x03, 0x00, 0x00,
}

func (m *CreatePoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.BaseAssetReserve.Size()
		i -= size
		if _, err := m.BaseAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.QuoteAssetReserve.Size()
		i -= size
		if _, err := m.QuoteAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditPoolConfigProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditPoolConfigProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditPoolConfigProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditSwapInvariantsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditSwapInvariantsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditSwapInvariantsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SwapInvariantMaps) > 0 {
		for iNdEx := len(m.SwapInvariantMaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SwapInvariantMaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Multiplier.Size()
		i -= size
		if _, err := m.Multiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreatePoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.QuoteAssetReserve.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.BaseAssetReserve.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Config.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *EditPoolConfigProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.Config.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *EditSwapInvariantsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.SwapInvariantMaps) > 0 {
		for _, e := range m.SwapInvariantMaps {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.Multiplier.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatePoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CreatePoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *EditPoolConfigProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: EditPoolConfigProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EditPoolConfigProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *EditSwapInvariantsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: EditSwapInvariantsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EditSwapInvariantsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapInvariantMaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SwapInvariantMaps = append(m.SwapInvariantMaps, EditSwapInvariantsProposal_SwapInvariantMultiple{})
			if err := m.SwapInvariantMaps[len(m.SwapInvariantMaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *EditSwapInvariantsProposal_SwapInvariantMultiple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: SwapInvariantMultiple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapInvariantMultiple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Multiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)

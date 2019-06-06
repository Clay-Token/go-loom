// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto

package coin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Economy struct {
	TotalSupply          *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Economy) Reset()         { *m = Economy{} }
func (m *Economy) String() string { return proto.CompactTextString(m) }
func (*Economy) ProtoMessage()    {}
func (*Economy) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{0}
}
func (m *Economy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Economy.Unmarshal(m, b)
}
func (m *Economy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Economy.Marshal(b, m, deterministic)
}
func (dst *Economy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Economy.Merge(dst, src)
}
func (m *Economy) XXX_Size() int {
	return xxx_messageInfo_Economy.Size(m)
}
func (m *Economy) XXX_DiscardUnknown() {
	xxx_messageInfo_Economy.DiscardUnknown(m)
}

var xxx_messageInfo_Economy proto.InternalMessageInfo

func (m *Economy) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type Account struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance              *types.BigUInt `protobuf:"bytes,2,opt,name=balance" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Account) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type Allowance struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender              *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Allowance) Reset()         { *m = Allowance{} }
func (m *Allowance) String() string { return proto.CompactTextString(m) }
func (*Allowance) ProtoMessage()    {}
func (*Allowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{2}
}
func (m *Allowance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Allowance.Unmarshal(m, b)
}
func (m *Allowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Allowance.Marshal(b, m, deterministic)
}
func (dst *Allowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Allowance.Merge(dst, src)
}
func (m *Allowance) XXX_Size() int {
	return xxx_messageInfo_Allowance.Size(m)
}
func (m *Allowance) XXX_DiscardUnknown() {
	xxx_messageInfo_Allowance.DiscardUnknown(m)
}

var xxx_messageInfo_Allowance proto.InternalMessageInfo

func (m *Allowance) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Allowance) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *Allowance) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type InitialAccount struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance              uint64         `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitialAccount) Reset()         { *m = InitialAccount{} }
func (m *InitialAccount) String() string { return proto.CompactTextString(m) }
func (*InitialAccount) ProtoMessage()    {}
func (*InitialAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{3}
}
func (m *InitialAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialAccount.Unmarshal(m, b)
}
func (m *InitialAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialAccount.Marshal(b, m, deterministic)
}
func (dst *InitialAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialAccount.Merge(dst, src)
}
func (m *InitialAccount) XXX_Size() int {
	return xxx_messageInfo_InitialAccount.Size(m)
}
func (m *InitialAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialAccount.DiscardUnknown(m)
}

var xxx_messageInfo_InitialAccount proto.InternalMessageInfo

func (m *InitialAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *InitialAccount) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type InitRequest struct {
	Accounts                   []*InitialAccount `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
	DeflationFactorNumerator   uint64            `protobuf:"varint,2,opt,name=deflation_factor_numerator,json=deflationFactorNumerator,proto3" json:"deflation_factor_numerator,omitempty"`
	DeflationFactorDenominator uint64            `protobuf:"varint,3,opt,name=deflation_factor_decimals,json=deflationFactorDecimals,proto3" json:"deflation_factor_decimals,omitempty"`
	BaseMintingAmount          uint64            `protobuf:"varint,4,opt,name=base_minting_amount,json=baseMintingAmount,proto3" json:"base_minting_amount,omitempty"`
	MintingAccount             *types.Address    `protobuf:"bytes,5,opt,name=minting_account,json=mintingAccount" json:"minting_account,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}          `json:"-"`
	XXX_unrecognized           []byte            `json:"-"`
	XXX_sizecache              int32             `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{4}
}
func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (dst *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(dst, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetAccounts() []*InitialAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *InitRequest) GetDeflationFactorNumerator() uint64 {
	if m != nil {
		return m.DeflationFactorNumerator
	}
	return 0
}

func (m *InitRequest) GetDeflationFactorDenominator() uint64 {
	if m != nil {
		return m.DeflationFactorDenominator
	}
	return 0
}

func (m *InitRequest) GetBaseMintingAmount() uint64 {
	if m != nil {
		return m.BaseMintingAmount
	}
	return 0
}

func (m *InitRequest) GetMintingAccount() *types.Address {
	if m != nil {
		return m.MintingAccount
	}
	return nil
}

type Policy struct {
	DeflationFactorNumerator   uint64         `protobuf:"varint,1,opt,name=deflation_factor_numerator,json=deflationFactorNumerator,proto3" json:"deflation_factor_numerator,omitempty"`
	DeflationFactorDenominator uint64         `protobuf:"varint,2,opt,name=deflation_factor_decimals,json=deflationFactorDecimals,proto3" json:"deflation_factor_decimals,omitempty"`
	BaseMintingAmount          *types.BigUInt `protobuf:"bytes,3,opt,name=base_minting_amount,json=baseMintingAmount" json:"base_minting_amount,omitempty"`
	MintingAccount             *types.Address `protobuf:"bytes,4,opt,name=minting_account,json=mintingAccount" json:"minting_account,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}       `json:"-"`
	XXX_unrecognized           []byte         `json:"-"`
	XXX_sizecache              int32          `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{5}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetDeflationFactorNumerator() uint64 {
	if m != nil {
		return m.DeflationFactorNumerator
	}
	return 0
}

func (m *Policy) GetDeflationFactorDenominator() uint64 {
	if m != nil {
		return m.DeflationFactorDenominator
	}
	return 0
}

func (m *Policy) GetBaseMintingAmount() *types.BigUInt {
	if m != nil {
		return m.BaseMintingAmount
	}
	return nil
}

func (m *Policy) GetMintingAccount() *types.Address {
	if m != nil {
		return m.MintingAccount
	}
	return nil
}

type MintToGatewayRequest struct {
	Amount               *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MintToGatewayRequest) Reset()         { *m = MintToGatewayRequest{} }
func (m *MintToGatewayRequest) String() string { return proto.CompactTextString(m) }
func (*MintToGatewayRequest) ProtoMessage()    {}
func (*MintToGatewayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{6}
}
func (m *MintToGatewayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintToGatewayRequest.Unmarshal(m, b)
}
func (m *MintToGatewayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintToGatewayRequest.Marshal(b, m, deterministic)
}
func (dst *MintToGatewayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintToGatewayRequest.Merge(dst, src)
}
func (m *MintToGatewayRequest) XXX_Size() int {
	return xxx_messageInfo_MintToGatewayRequest.Size(m)
}
func (m *MintToGatewayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintToGatewayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintToGatewayRequest proto.InternalMessageInfo

func (m *MintToGatewayRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type BurnRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BurnRequest) Reset()         { *m = BurnRequest{} }
func (m *BurnRequest) String() string { return proto.CompactTextString(m) }
func (*BurnRequest) ProtoMessage()    {}
func (*BurnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{7}
}
func (m *BurnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BurnRequest.Unmarshal(m, b)
}
func (m *BurnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BurnRequest.Marshal(b, m, deterministic)
}
func (dst *BurnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnRequest.Merge(dst, src)
}
func (m *BurnRequest) XXX_Size() int {
	return xxx_messageInfo_BurnRequest.Size(m)
}
func (m *BurnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BurnRequest proto.InternalMessageInfo

func (m *BurnRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *BurnRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

// read only
type TotalSupplyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalSupplyRequest) Reset()         { *m = TotalSupplyRequest{} }
func (m *TotalSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*TotalSupplyRequest) ProtoMessage()    {}
func (*TotalSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{8}
}
func (m *TotalSupplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalSupplyRequest.Unmarshal(m, b)
}
func (m *TotalSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalSupplyRequest.Marshal(b, m, deterministic)
}
func (dst *TotalSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalSupplyRequest.Merge(dst, src)
}
func (m *TotalSupplyRequest) XXX_Size() int {
	return xxx_messageInfo_TotalSupplyRequest.Size(m)
}
func (m *TotalSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TotalSupplyRequest proto.InternalMessageInfo

type TotalSupplyResponse struct {
	TotalSupply          *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TotalSupplyResponse) Reset()         { *m = TotalSupplyResponse{} }
func (m *TotalSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*TotalSupplyResponse) ProtoMessage()    {}
func (*TotalSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{9}
}
func (m *TotalSupplyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalSupplyResponse.Unmarshal(m, b)
}
func (m *TotalSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalSupplyResponse.Marshal(b, m, deterministic)
}
func (dst *TotalSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalSupplyResponse.Merge(dst, src)
}
func (m *TotalSupplyResponse) XXX_Size() int {
	return xxx_messageInfo_TotalSupplyResponse.Size(m)
}
func (m *TotalSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TotalSupplyResponse proto.InternalMessageInfo

func (m *TotalSupplyResponse) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type BalanceOfRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceOfRequest) Reset()         { *m = BalanceOfRequest{} }
func (m *BalanceOfRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceOfRequest) ProtoMessage()    {}
func (*BalanceOfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{10}
}
func (m *BalanceOfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceOfRequest.Unmarshal(m, b)
}
func (m *BalanceOfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceOfRequest.Marshal(b, m, deterministic)
}
func (dst *BalanceOfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceOfRequest.Merge(dst, src)
}
func (m *BalanceOfRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceOfRequest.Size(m)
}
func (m *BalanceOfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceOfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceOfRequest proto.InternalMessageInfo

func (m *BalanceOfRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type BalanceOfResponse struct {
	Balance              *types.BigUInt `protobuf:"bytes,1,opt,name=balance" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceOfResponse) Reset()         { *m = BalanceOfResponse{} }
func (m *BalanceOfResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceOfResponse) ProtoMessage()    {}
func (*BalanceOfResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{11}
}
func (m *BalanceOfResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceOfResponse.Unmarshal(m, b)
}
func (m *BalanceOfResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceOfResponse.Marshal(b, m, deterministic)
}
func (dst *BalanceOfResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceOfResponse.Merge(dst, src)
}
func (m *BalanceOfResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceOfResponse.Size(m)
}
func (m *BalanceOfResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceOfResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceOfResponse proto.InternalMessageInfo

func (m *BalanceOfResponse) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type AllowanceRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender              *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllowanceRequest) Reset()         { *m = AllowanceRequest{} }
func (m *AllowanceRequest) String() string { return proto.CompactTextString(m) }
func (*AllowanceRequest) ProtoMessage()    {}
func (*AllowanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{12}
}
func (m *AllowanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowanceRequest.Unmarshal(m, b)
}
func (m *AllowanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowanceRequest.Marshal(b, m, deterministic)
}
func (dst *AllowanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceRequest.Merge(dst, src)
}
func (m *AllowanceRequest) XXX_Size() int {
	return xxx_messageInfo_AllowanceRequest.Size(m)
}
func (m *AllowanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceRequest proto.InternalMessageInfo

func (m *AllowanceRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AllowanceRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

type AllowanceResponse struct {
	Amount               *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllowanceResponse) Reset()         { *m = AllowanceResponse{} }
func (m *AllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*AllowanceResponse) ProtoMessage()    {}
func (*AllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{13}
}
func (m *AllowanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowanceResponse.Unmarshal(m, b)
}
func (m *AllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowanceResponse.Marshal(b, m, deterministic)
}
func (dst *AllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceResponse.Merge(dst, src)
}
func (m *AllowanceResponse) XXX_Size() int {
	return xxx_messageInfo_AllowanceResponse.Size(m)
}
func (m *AllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceResponse proto.InternalMessageInfo

func (m *AllowanceResponse) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

// volatile
type ApproveRequest struct {
	Spender              *types.Address `protobuf:"bytes,1,opt,name=spender" json:"spender,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApproveRequest) Reset()         { *m = ApproveRequest{} }
func (m *ApproveRequest) String() string { return proto.CompactTextString(m) }
func (*ApproveRequest) ProtoMessage()    {}
func (*ApproveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{14}
}
func (m *ApproveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApproveRequest.Unmarshal(m, b)
}
func (m *ApproveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApproveRequest.Marshal(b, m, deterministic)
}
func (dst *ApproveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApproveRequest.Merge(dst, src)
}
func (m *ApproveRequest) XXX_Size() int {
	return xxx_messageInfo_ApproveRequest.Size(m)
}
func (m *ApproveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApproveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApproveRequest proto.InternalMessageInfo

func (m *ApproveRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *ApproveRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type ApproveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApproveResponse) Reset()         { *m = ApproveResponse{} }
func (m *ApproveResponse) String() string { return proto.CompactTextString(m) }
func (*ApproveResponse) ProtoMessage()    {}
func (*ApproveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{15}
}
func (m *ApproveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApproveResponse.Unmarshal(m, b)
}
func (m *ApproveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApproveResponse.Marshal(b, m, deterministic)
}
func (dst *ApproveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApproveResponse.Merge(dst, src)
}
func (m *ApproveResponse) XXX_Size() int {
	return xxx_messageInfo_ApproveResponse.Size(m)
}
func (m *ApproveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApproveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApproveResponse proto.InternalMessageInfo

type TransferRequest struct {
	To                   *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{16}
}
func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (dst *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(dst, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{17}
}
func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (dst *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(dst, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

type TransferFromRequest struct {
	From                 *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To                   *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferFromRequest) Reset()         { *m = TransferFromRequest{} }
func (m *TransferFromRequest) String() string { return proto.CompactTextString(m) }
func (*TransferFromRequest) ProtoMessage()    {}
func (*TransferFromRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{18}
}
func (m *TransferFromRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFromRequest.Unmarshal(m, b)
}
func (m *TransferFromRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFromRequest.Marshal(b, m, deterministic)
}
func (dst *TransferFromRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFromRequest.Merge(dst, src)
}
func (m *TransferFromRequest) XXX_Size() int {
	return xxx_messageInfo_TransferFromRequest.Size(m)
}
func (m *TransferFromRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFromRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFromRequest proto.InternalMessageInfo

func (m *TransferFromRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransferFromRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferFromRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferFromResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferFromResponse) Reset()         { *m = TransferFromResponse{} }
func (m *TransferFromResponse) String() string { return proto.CompactTextString(m) }
func (*TransferFromResponse) ProtoMessage()    {}
func (*TransferFromResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_a41ade0a2aff6c59, []int{19}
}
func (m *TransferFromResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFromResponse.Unmarshal(m, b)
}
func (m *TransferFromResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFromResponse.Marshal(b, m, deterministic)
}
func (dst *TransferFromResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFromResponse.Merge(dst, src)
}
func (m *TransferFromResponse) XXX_Size() int {
	return xxx_messageInfo_TransferFromResponse.Size(m)
}
func (m *TransferFromResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFromResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFromResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Economy)(nil), "Economy")
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*Allowance)(nil), "Allowance")
	proto.RegisterType((*InitialAccount)(nil), "InitialAccount")
	proto.RegisterType((*InitRequest)(nil), "InitRequest")
	proto.RegisterType((*Policy)(nil), "Policy")
	proto.RegisterType((*MintToGatewayRequest)(nil), "MintToGatewayRequest")
	proto.RegisterType((*BurnRequest)(nil), "BurnRequest")
	proto.RegisterType((*TotalSupplyRequest)(nil), "TotalSupplyRequest")
	proto.RegisterType((*TotalSupplyResponse)(nil), "TotalSupplyResponse")
	proto.RegisterType((*BalanceOfRequest)(nil), "BalanceOfRequest")
	proto.RegisterType((*BalanceOfResponse)(nil), "BalanceOfResponse")
	proto.RegisterType((*AllowanceRequest)(nil), "AllowanceRequest")
	proto.RegisterType((*AllowanceResponse)(nil), "AllowanceResponse")
	proto.RegisterType((*ApproveRequest)(nil), "ApproveRequest")
	proto.RegisterType((*ApproveResponse)(nil), "ApproveResponse")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
	proto.RegisterType((*TransferFromRequest)(nil), "TransferFromRequest")
	proto.RegisterType((*TransferFromResponse)(nil), "TransferFromResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto", fileDescriptor_coin_a41ade0a2aff6c59)
}

var fileDescriptor_coin_a41ade0a2aff6c59 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x4f, 0xdb, 0x4a,
	0x14, 0x96, 0x4d, 0x78, 0xdc, 0x93, 0xab, 0x84, 0x18, 0xee, 0xad, 0x8b, 0x10, 0x41, 0x5e, 0x21,
	0x21, 0x92, 0x8a, 0xaa, 0x8f, 0x45, 0x37, 0xb1, 0x80, 0x8a, 0x4a, 0x94, 0x2a, 0xa5, 0x2c, 0xd8,
	0x44, 0x63, 0x7b, 0x92, 0x8e, 0x6a, 0xcf, 0x31, 0x33, 0xe3, 0x46, 0xd9, 0xf5, 0x67, 0x76, 0x95,
	0x45, 0x7e, 0x49, 0x95, 0xf1, 0x23, 0x0f, 0x9a, 0x26, 0x2d, 0xea, 0x26, 0xca, 0x79, 0x7d, 0xdf,
	0x37, 0x9f, 0x3d, 0xc7, 0xf0, 0xa6, 0xc7, 0xd4, 0xe7, 0xc4, 0x6b, 0xf8, 0x18, 0x35, 0x43, 0xc4,
	0x88, 0x53, 0xd5, 0x47, 0xf1, 0xa5, 0xd9, 0xc3, 0x93, 0x71, 0xd8, 0xf4, 0x12, 0x16, 0x2a, 0xc6,
	0x9b, 0x6a, 0x10, 0x53, 0xd9, 0xf4, 0x91, 0x71, 0xfd, 0xd3, 0x88, 0x05, 0x2a, 0xdc, 0x7b, 0xb6,
	0x64, 0x3a, 0x9d, 0xd2, 0xbf, 0xd9, 0xc4, 0xc9, 0xd4, 0x44, 0x0f, 0x7b, 0xd8, 0xd4, 0x69, 0x2f,
	0xe9, 0xea, 0x48, 0x07, 0xfa, 0x5f, 0xda, 0xee, 0xbc, 0x84, 0xcd, 0x73, 0x1f, 0x39, 0x46, 0x03,
	0xeb, 0x18, 0xfe, 0x55, 0xa8, 0x48, 0xd8, 0x91, 0x49, 0x1c, 0x87, 0x03, 0xdb, 0x38, 0x34, 0x8e,
	0xca, 0xa7, 0x5b, 0x0d, 0x97, 0xf5, 0x3e, 0x5d, 0x72, 0xd5, 0x2e, 0xeb, 0xea, 0x47, 0x5d, 0x74,
	0xae, 0x60, 0xb3, 0xe5, 0xfb, 0x98, 0x70, 0x65, 0x1d, 0xc0, 0x3a, 0xf6, 0x39, 0x15, 0xc5, 0x40,
	0x2b, 0x08, 0x04, 0x95, 0xb2, 0x9d, 0xa6, 0x2d, 0x07, 0x36, 0x3d, 0x12, 0x12, 0xee, 0x53, 0xdb,
	0x9c, 0x83, 0xcc, 0x0b, 0xce, 0x3d, 0xfc, 0xd3, 0x0a, 0x43, 0xec, 0x8f, 0x83, 0x55, 0x00, 0x65,
	0x4c, 0x79, 0x40, 0x45, 0x01, 0x98, 0x77, 0xe4, 0x05, 0xeb, 0x10, 0x36, 0x48, 0x34, 0x96, 0x67,
	0xaf, 0xcd, 0x71, 0x66, 0x79, 0xe7, 0x1d, 0x54, 0x2e, 0x39, 0x53, 0x8c, 0x84, 0xab, 0x1e, 0xc4,
	0x9e, 0x3d, 0x48, 0x69, 0x22, 0xff, 0xdb, 0x1a, 0x94, 0xc7, 0x60, 0x6d, 0x7a, 0x9f, 0x50, 0xa9,
	0xac, 0x63, 0xd8, 0x22, 0x29, 0xa8, 0xb4, 0x8d, 0xc3, 0xb5, 0xa3, 0xf2, 0x69, 0xb5, 0x31, 0x4b,
	0xd6, 0x2e, 0x1a, 0xac, 0x3b, 0xd8, 0x0b, 0x68, 0x37, 0x24, 0x8a, 0x21, 0xef, 0x74, 0x89, 0xaf,
	0x50, 0x74, 0x78, 0x12, 0x51, 0x41, 0x14, 0xa6, 0x27, 0x2c, 0xb9, 0xfb, 0xa3, 0x61, 0xdd, 0x3e,
	0xcb, 0xbb, 0x2e, 0x74, 0xd3, 0xfb, 0xbc, 0xa7, 0x6d, 0x07, 0x0b, 0x2a, 0xd6, 0x1d, 0x3c, 0x7d,
	0x80, 0x1d, 0x50, 0x9f, 0x45, 0x24, 0x94, 0xda, 0x99, 0x92, 0x7b, 0x30, 0x1a, 0xd6, 0xf7, 0xe6,
	0xa0, 0xcf, 0x28, 0xc7, 0x88, 0x71, 0x0d, 0xfe, 0x24, 0x98, 0xaf, 0xa5, 0xe3, 0xd6, 0x39, 0xec,
	0x78, 0x44, 0xd2, 0x4e, 0xc4, 0xb8, 0x62, 0xbc, 0xd7, 0xc9, 0xfc, 0x2e, 0x69, 0xd4, 0xff, 0x46,
	0xc3, 0x7a, 0xcd, 0x25, 0x92, 0x5e, 0xa5, 0xd5, 0x96, 0x2e, 0xb6, 0x6b, 0xde, 0x7c, 0xca, 0x3a,
	0x87, 0x6a, 0x81, 0x90, 0x5a, 0x62, 0xaf, 0xcf, 0xfa, 0xef, 0x5a, 0xa3, 0x61, 0xbd, 0x92, 0x4f,
	0x65, 0xfe, 0x55, 0xa2, 0x99, 0xd8, 0xf9, 0x6e, 0xc2, 0xc6, 0x07, 0x0c, 0x99, 0x3f, 0x58, 0x62,
	0xa8, 0xf1, 0xf7, 0x0c, 0x35, 0x1f, 0x67, 0xe8, 0xf5, 0xcf, 0x0d, 0x9d, 0x7b, 0x81, 0x1f, 0x67,
	0x6d, 0xe9, 0x0f, 0xac, 0x7d, 0x0d, 0xbb, 0xe3, 0x8e, 0x1b, 0x7c, 0x4b, 0x14, 0xed, 0x93, 0x41,
	0xfe, 0x96, 0x4f, 0xee, 0x98, 0xb1, 0xe0, 0x8e, 0x5d, 0x43, 0xd9, 0x4d, 0x04, 0xcf, 0x07, 0x96,
	0x5d, 0xb0, 0x09, 0xa0, 0xb9, 0x00, 0x70, 0x17, 0xac, 0x9b, 0xc9, 0x16, 0xca, 0x70, 0x1d, 0x17,
	0x76, 0x66, 0xb2, 0x32, 0x46, 0x2e, 0xe9, 0xef, 0x2d, 0xb4, 0x53, 0xd8, 0x76, 0xd3, 0xdb, 0x7c,
	0xdd, 0x5d, 0x51, 0xaf, 0xf3, 0x0a, 0x6a, 0x53, 0x33, 0x19, 0xeb, 0xd4, 0xba, 0x33, 0x16, 0xad,
	0xbb, 0x5b, 0xd8, 0x2e, 0xd6, 0xdd, 0xaa, 0xe6, 0xac, 0xb0, 0xf5, 0x9c, 0x17, 0x50, 0x9b, 0xc2,
	0xcd, 0x04, 0x2d, 0x7f, 0x4c, 0xb7, 0x50, 0x69, 0xc5, 0xb1, 0xc0, 0xaf, 0x85, 0x98, 0x29, 0x32,
	0x63, 0xf9, 0x8a, 0x5d, 0xf4, 0xb4, 0x6a, 0x50, 0x2d, 0x70, 0x53, 0x31, 0xce, 0x15, 0x54, 0x6f,
	0x04, 0xe1, 0xb2, 0x4b, 0x45, 0xce, 0x65, 0x83, 0xa9, 0xf0, 0x01, 0x8d, 0xa9, 0x70, 0x05, 0x06,
	0x0b, 0xb6, 0x27, 0x70, 0x19, 0x05, 0xc2, 0x4e, 0x9e, 0xbb, 0x10, 0x18, 0xe5, 0x34, 0xfb, 0x50,
	0xea, 0x0a, 0x8c, 0x1e, 0x10, 0xe9, 0x6c, 0x26, 0xc2, 0xfc, 0xa5, 0x88, 0x45, 0x5f, 0x92, 0xff,
	0x61, 0x77, 0x96, 0x30, 0x15, 0xe2, 0x6d, 0xe8, 0x4f, 0xec, 0xf3, 0x1f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x59, 0xc4, 0x44, 0x03, 0x08, 0x00, 0x00,
}

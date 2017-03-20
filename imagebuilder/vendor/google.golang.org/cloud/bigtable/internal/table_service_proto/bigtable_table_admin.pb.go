// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_admin.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_v2 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_admin.proto

It has these top-level messages:
	CreateTableRequest
	DropRowRangeRequest
	ListTablesRequest
	ListTablesResponse
	GetTableRequest
	DeleteTableRequest
	ModifyColumnFamiliesRequest
*/
package google_bigtable_admin_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_v21 "google.golang.org/cloud/bigtable/internal/table_data_proto"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [google.bigtable.admin.v2.BigtableTableAdmin.CreateTable][google.bigtable.admin.v2.BigtableTableAdmin.CreateTable]
type CreateTableRequest struct {
	// The unique name of the instance in which to create the table.
	// Values are of the form projects/<project>/instances/<instance>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name by which the new table should be referred to within the parent
	// instance, e.g. "foobar" rather than "<parent>/tables/foobar".
	TableId string `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	// The Table to create.
	Table *google_bigtable_admin_v21.Table `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	// The optional list of row keys that will be used to initially split the
	// table into several tablets (Tablets are similar to HBase regions).
	// Given two split keys, "s1" and "s2", three tablets will be created,
	// spanning the key ranges: [, s1), [s1, s2), [s2, ).
	//
	// Example:
	//  * Row keys := ["a", "apple", "custom", "customer_1", "customer_2",
	//                 "other", "zz"]
	//  * initial_split_keys := ["apple", "customer_1", "customer_2", "other"]
	//  * Key assignment:
	//    - Tablet 1 [, apple)                => {"a"}.
	//    - Tablet 2 [apple, customer_1)      => {"apple", "custom"}.
	//    - Tablet 3 [customer_1, customer_2) => {"customer_1"}.
	//    - Tablet 4 [customer_2, other)      => {"customer_2"}.
	//    - Tablet 5 [other, )                => {"other", "zz"}.
	InitialSplits []*CreateTableRequest_Split `protobuf:"bytes,4,rep,name=initial_splits,json=initialSplits" json:"initial_splits,omitempty"`
}

func (m *CreateTableRequest) Reset()                    { *m = CreateTableRequest{} }
func (m *CreateTableRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTableRequest) ProtoMessage()               {}
func (*CreateTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTableRequest) GetTable() *google_bigtable_admin_v21.Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *CreateTableRequest) GetInitialSplits() []*CreateTableRequest_Split {
	if m != nil {
		return m.InitialSplits
	}
	return nil
}

// An initial split point for a newly created table.
type CreateTableRequest_Split struct {
	// Row key to use as an initial tablet boundary.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *CreateTableRequest_Split) Reset()                    { *m = CreateTableRequest_Split{} }
func (m *CreateTableRequest_Split) String() string            { return proto.CompactTextString(m) }
func (*CreateTableRequest_Split) ProtoMessage()               {}
func (*CreateTableRequest_Split) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Request message for [google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange][google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange]
type DropRowRangeRequest struct {
	// The unique name of the table on which to drop a range of rows.
	// Values are of the form projects/<project>/instances/<instance>/tables/<table>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Target:
	//	*DropRowRangeRequest_RowKeyPrefix
	//	*DropRowRangeRequest_DeleteAllDataFromTable
	Target isDropRowRangeRequest_Target `protobuf_oneof:"target"`
}

func (m *DropRowRangeRequest) Reset()                    { *m = DropRowRangeRequest{} }
func (m *DropRowRangeRequest) String() string            { return proto.CompactTextString(m) }
func (*DropRowRangeRequest) ProtoMessage()               {}
func (*DropRowRangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isDropRowRangeRequest_Target interface {
	isDropRowRangeRequest_Target()
}

type DropRowRangeRequest_RowKeyPrefix struct {
	RowKeyPrefix []byte `protobuf:"bytes,2,opt,name=row_key_prefix,json=rowKeyPrefix,proto3,oneof"`
}
type DropRowRangeRequest_DeleteAllDataFromTable struct {
	DeleteAllDataFromTable bool `protobuf:"varint,3,opt,name=delete_all_data_from_table,json=deleteAllDataFromTable,oneof"`
}

func (*DropRowRangeRequest_RowKeyPrefix) isDropRowRangeRequest_Target()           {}
func (*DropRowRangeRequest_DeleteAllDataFromTable) isDropRowRangeRequest_Target() {}

func (m *DropRowRangeRequest) GetTarget() isDropRowRangeRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DropRowRangeRequest) GetRowKeyPrefix() []byte {
	if x, ok := m.GetTarget().(*DropRowRangeRequest_RowKeyPrefix); ok {
		return x.RowKeyPrefix
	}
	return nil
}

func (m *DropRowRangeRequest) GetDeleteAllDataFromTable() bool {
	if x, ok := m.GetTarget().(*DropRowRangeRequest_DeleteAllDataFromTable); ok {
		return x.DeleteAllDataFromTable
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DropRowRangeRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DropRowRangeRequest_OneofMarshaler, _DropRowRangeRequest_OneofUnmarshaler, _DropRowRangeRequest_OneofSizer, []interface{}{
		(*DropRowRangeRequest_RowKeyPrefix)(nil),
		(*DropRowRangeRequest_DeleteAllDataFromTable)(nil),
	}
}

func _DropRowRangeRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DropRowRangeRequest)
	// target
	switch x := m.Target.(type) {
	case *DropRowRangeRequest_RowKeyPrefix:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RowKeyPrefix)
	case *DropRowRangeRequest_DeleteAllDataFromTable:
		t := uint64(0)
		if x.DeleteAllDataFromTable {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("DropRowRangeRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _DropRowRangeRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DropRowRangeRequest)
	switch tag {
	case 2: // target.row_key_prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Target = &DropRowRangeRequest_RowKeyPrefix{x}
		return true, err
	case 3: // target.delete_all_data_from_table
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Target = &DropRowRangeRequest_DeleteAllDataFromTable{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _DropRowRangeRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DropRowRangeRequest)
	// target
	switch x := m.Target.(type) {
	case *DropRowRangeRequest_RowKeyPrefix:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RowKeyPrefix)))
		n += len(x.RowKeyPrefix)
	case *DropRowRangeRequest_DeleteAllDataFromTable:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request message for [google.bigtable.admin.v2.BigtableTableAdmin.ListTables][google.bigtable.admin.v2.BigtableTableAdmin.ListTables]
type ListTablesRequest struct {
	// The unique name of the instance for which tables should be listed.
	// Values are of the form projects/<project>/instances/<instance>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The view to be applied to the returned tables' fields.
	// Defaults to NAME_ONLY if unspecified (no others are currently supported).
	View google_bigtable_admin_v21.Table_View `protobuf:"varint,2,opt,name=view,enum=google.bigtable.admin.v2.Table_View" json:"view,omitempty"`
	// Not yet supported.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListTablesRequest) Reset()                    { *m = ListTablesRequest{} }
func (m *ListTablesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTablesRequest) ProtoMessage()               {}
func (*ListTablesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message for [google.bigtable.admin.v2.BigtableTableAdmin.ListTables][google.bigtable.admin.v2.BigtableTableAdmin.ListTables]
type ListTablesResponse struct {
	// The tables present in the requested cluster.
	Tables        []*google_bigtable_admin_v21.Table `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
	NextPageToken string                             `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListTablesResponse) Reset()                    { *m = ListTablesResponse{} }
func (m *ListTablesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTablesResponse) ProtoMessage()               {}
func (*ListTablesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListTablesResponse) GetTables() []*google_bigtable_admin_v21.Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

// Request message for [google.bigtable.admin.v2.BigtableTableAdmin.GetTable][google.bigtable.admin.v2.BigtableTableAdmin.GetTable]
type GetTableRequest struct {
	// The unique name of the requested table.
	// Values are of the form projects/<project>/instances/<instance>/tables/<table>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The view to be applied to the returned table's fields.
	// Defaults to SCHEMA_ONLY if unspecified.
	View google_bigtable_admin_v21.Table_View `protobuf:"varint,2,opt,name=view,enum=google.bigtable.admin.v2.Table_View" json:"view,omitempty"`
}

func (m *GetTableRequest) Reset()                    { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()               {}
func (*GetTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Request message for [google.bigtable.admin.v2.BigtableTableAdmin.DeleteTable][google.bigtable.admin.v2.BigtableTableAdmin.DeleteTable]
type DeleteTableRequest struct {
	// The unique name of the table to be deleted.
	// Values are of the form projects/<project>/instances/<instance>/tables/<table>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteTableRequest) Reset()                    { *m = DeleteTableRequest{} }
func (m *DeleteTableRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTableRequest) ProtoMessage()               {}
func (*DeleteTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Request message for [google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies][google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies]
type ModifyColumnFamiliesRequest struct {
	// The unique name of the table whose families should be modified.
	// Values are of the form projects/<project>/instances/<instance>/tables/<table>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Modifications to be atomically applied to the specified table's families.
	// Entries are applied in order, meaning that earlier modifications can be
	// masked by later ones (in the case of repeated updates to the same family,
	// for example).
	Modifications []*ModifyColumnFamiliesRequest_Modification `protobuf:"bytes,2,rep,name=modifications" json:"modifications,omitempty"`
}

func (m *ModifyColumnFamiliesRequest) Reset()                    { *m = ModifyColumnFamiliesRequest{} }
func (m *ModifyColumnFamiliesRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyColumnFamiliesRequest) ProtoMessage()               {}
func (*ModifyColumnFamiliesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ModifyColumnFamiliesRequest) GetModifications() []*ModifyColumnFamiliesRequest_Modification {
	if m != nil {
		return m.Modifications
	}
	return nil
}

// A create, update, or delete of a particular column family.
type ModifyColumnFamiliesRequest_Modification struct {
	// The ID of the column family to be modified.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Types that are valid to be assigned to Mod:
	//	*ModifyColumnFamiliesRequest_Modification_Create
	//	*ModifyColumnFamiliesRequest_Modification_Update
	//	*ModifyColumnFamiliesRequest_Modification_Drop
	Mod isModifyColumnFamiliesRequest_Modification_Mod `protobuf_oneof:"mod"`
}

func (m *ModifyColumnFamiliesRequest_Modification) Reset() {
	*m = ModifyColumnFamiliesRequest_Modification{}
}
func (m *ModifyColumnFamiliesRequest_Modification) String() string { return proto.CompactTextString(m) }
func (*ModifyColumnFamiliesRequest_Modification) ProtoMessage()    {}
func (*ModifyColumnFamiliesRequest_Modification) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type isModifyColumnFamiliesRequest_Modification_Mod interface {
	isModifyColumnFamiliesRequest_Modification_Mod()
}

type ModifyColumnFamiliesRequest_Modification_Create struct {
	Create *google_bigtable_admin_v21.ColumnFamily `protobuf:"bytes,2,opt,name=create,oneof"`
}
type ModifyColumnFamiliesRequest_Modification_Update struct {
	Update *google_bigtable_admin_v21.ColumnFamily `protobuf:"bytes,3,opt,name=update,oneof"`
}
type ModifyColumnFamiliesRequest_Modification_Drop struct {
	Drop bool `protobuf:"varint,4,opt,name=drop,oneof"`
}

func (*ModifyColumnFamiliesRequest_Modification_Create) isModifyColumnFamiliesRequest_Modification_Mod() {
}
func (*ModifyColumnFamiliesRequest_Modification_Update) isModifyColumnFamiliesRequest_Modification_Mod() {
}
func (*ModifyColumnFamiliesRequest_Modification_Drop) isModifyColumnFamiliesRequest_Modification_Mod() {
}

func (m *ModifyColumnFamiliesRequest_Modification) GetMod() isModifyColumnFamiliesRequest_Modification_Mod {
	if m != nil {
		return m.Mod
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetCreate() *google_bigtable_admin_v21.ColumnFamily {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetUpdate() *google_bigtable_admin_v21.ColumnFamily {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetDrop() bool {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Drop); ok {
		return x.Drop
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModifyColumnFamiliesRequest_Modification) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModifyColumnFamiliesRequest_Modification_OneofMarshaler, _ModifyColumnFamiliesRequest_Modification_OneofUnmarshaler, _ModifyColumnFamiliesRequest_Modification_OneofSizer, []interface{}{
		(*ModifyColumnFamiliesRequest_Modification_Create)(nil),
		(*ModifyColumnFamiliesRequest_Modification_Update)(nil),
		(*ModifyColumnFamiliesRequest_Modification_Drop)(nil),
	}
}

func _ModifyColumnFamiliesRequest_Modification_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	// mod
	switch x := m.Mod.(type) {
	case *ModifyColumnFamiliesRequest_Modification_Create:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *ModifyColumnFamiliesRequest_Modification_Update:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *ModifyColumnFamiliesRequest_Modification_Drop:
		t := uint64(0)
		if x.Drop {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ModifyColumnFamiliesRequest_Modification.Mod has unexpected type %T", x)
	}
	return nil
}

func _ModifyColumnFamiliesRequest_Modification_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	switch tag {
	case 2: // mod.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_bigtable_admin_v21.ColumnFamily)
		err := b.DecodeMessage(msg)
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Create{msg}
		return true, err
	case 3: // mod.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_bigtable_admin_v21.ColumnFamily)
		err := b.DecodeMessage(msg)
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Update{msg}
		return true, err
	case 4: // mod.drop
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Drop{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ModifyColumnFamiliesRequest_Modification_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	// mod
	switch x := m.Mod.(type) {
	case *ModifyColumnFamiliesRequest_Modification_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModifyColumnFamiliesRequest_Modification_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModifyColumnFamiliesRequest_Modification_Drop:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CreateTableRequest)(nil), "google.bigtable.admin.v2.CreateTableRequest")
	proto.RegisterType((*CreateTableRequest_Split)(nil), "google.bigtable.admin.v2.CreateTableRequest.Split")
	proto.RegisterType((*DropRowRangeRequest)(nil), "google.bigtable.admin.v2.DropRowRangeRequest")
	proto.RegisterType((*ListTablesRequest)(nil), "google.bigtable.admin.v2.ListTablesRequest")
	proto.RegisterType((*ListTablesResponse)(nil), "google.bigtable.admin.v2.ListTablesResponse")
	proto.RegisterType((*GetTableRequest)(nil), "google.bigtable.admin.v2.GetTableRequest")
	proto.RegisterType((*DeleteTableRequest)(nil), "google.bigtable.admin.v2.DeleteTableRequest")
	proto.RegisterType((*ModifyColumnFamiliesRequest)(nil), "google.bigtable.admin.v2.ModifyColumnFamiliesRequest")
	proto.RegisterType((*ModifyColumnFamiliesRequest_Modification)(nil), "google.bigtable.admin.v2.ModifyColumnFamiliesRequest.Modification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for BigtableTableAdmin service

type BigtableTableAdminClient interface {
	// Creates a new table in the specified instance.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_v21.Table, error)
	// Lists all tables served from a specified instance.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// Gets metadata information about the specified table.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_v21.Table, error)
	// Permanently deletes a specified table and all of its data.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Atomically performs a series of column family modifications
	// on the specified table.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	ModifyColumnFamilies(ctx context.Context, in *ModifyColumnFamiliesRequest, opts ...grpc.CallOption) (*google_bigtable_admin_v21.Table, error)
	// Permanently drop/delete a row range from a specified table. The request can
	// specify whether to delete all rows in a table, or only those that match a
	// particular prefix.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	DropRowRange(ctx context.Context, in *DropRowRangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type bigtableTableAdminClient struct {
	cc *grpc.ClientConn
}

func NewBigtableTableAdminClient(cc *grpc.ClientConn) BigtableTableAdminClient {
	return &bigtableTableAdminClient{cc}
}

func (c *bigtableTableAdminClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_v21.Table, error) {
	out := new(google_bigtable_admin_v21.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/CreateTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/ListTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_v21.Table, error) {
	out := new(google_bigtable_admin_v21.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/GetTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/DeleteTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) ModifyColumnFamilies(ctx context.Context, in *ModifyColumnFamiliesRequest, opts ...grpc.CallOption) (*google_bigtable_admin_v21.Table, error) {
	out := new(google_bigtable_admin_v21.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/ModifyColumnFamilies", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) DropRowRange(ctx context.Context, in *DropRowRangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/DropRowRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableTableAdmin service

type BigtableTableAdminServer interface {
	// Creates a new table in the specified instance.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	CreateTable(context.Context, *CreateTableRequest) (*google_bigtable_admin_v21.Table, error)
	// Lists all tables served from a specified instance.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// Gets metadata information about the specified table.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	GetTable(context.Context, *GetTableRequest) (*google_bigtable_admin_v21.Table, error)
	// Permanently deletes a specified table and all of its data.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	DeleteTable(context.Context, *DeleteTableRequest) (*google_protobuf1.Empty, error)
	// Atomically performs a series of column family modifications
	// on the specified table.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	ModifyColumnFamilies(context.Context, *ModifyColumnFamiliesRequest) (*google_bigtable_admin_v21.Table, error)
	// Permanently drop/delete a row range from a specified table. The request can
	// specify whether to delete all rows in a table, or only those that match a
	// particular prefix.
	//
	// Caution: This rpc is experimental. The details can change and the rpc
	// may or may not be active.
	DropRowRange(context.Context, *DropRowRangeRequest) (*google_protobuf1.Empty, error)
}

func RegisterBigtableTableAdminServer(s *grpc.Server, srv BigtableTableAdminServer) {
	s.RegisterService(&_BigtableTableAdmin_serviceDesc, srv)
}

func _BigtableTableAdmin_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_DeleteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).DeleteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/DeleteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).DeleteTable(ctx, req.(*DeleteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_ModifyColumnFamilies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyColumnFamiliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).ModifyColumnFamilies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/ModifyColumnFamilies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).ModifyColumnFamilies(ctx, req.(*ModifyColumnFamiliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_DropRowRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropRowRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).DropRowRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/DropRowRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).DropRowRange(ctx, req.(*DropRowRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableTableAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.v2.BigtableTableAdmin",
	HandlerType: (*BigtableTableAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _BigtableTableAdmin_CreateTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _BigtableTableAdmin_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _BigtableTableAdmin_GetTable_Handler,
		},
		{
			MethodName: "DeleteTable",
			Handler:    _BigtableTableAdmin_DeleteTable_Handler,
		},
		{
			MethodName: "ModifyColumnFamilies",
			Handler:    _BigtableTableAdmin_ModifyColumnFamilies_Handler,
		},
		{
			MethodName: "DropRowRange",
			Handler:    _BigtableTableAdmin_DropRowRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_admin.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x48,
	0x14, 0xae, 0x9b, 0xb4, 0x9b, 0x9e, 0xa4, 0xe9, 0xee, 0x6c, 0xd5, 0x4d, 0xdd, 0x45, 0x54, 0x16,
	0xaa, 0x8a, 0x28, 0x8e, 0x14, 0x54, 0x01, 0x12, 0x17, 0x34, 0x2d, 0x05, 0x04, 0x48, 0x91, 0x29,
	0x50, 0x6e, 0xb0, 0x9c, 0x78, 0x62, 0x46, 0x1d, 0x7b, 0x8c, 0x3d, 0x69, 0x9b, 0x3b, 0x5e, 0x83,
	0xa7, 0xe1, 0x41, 0x10, 0xef, 0xc0, 0x23, 0x30, 0x3f, 0x4e, 0x1b, 0x08, 0x76, 0xd2, 0xde, 0x44,
	0x67, 0xce, 0x9c, 0xef, 0x3b, 0x3f, 0xf3, 0x1d, 0x07, 0x8e, 0x03, 0xc6, 0x02, 0x8a, 0xed, 0x80,
	0x51, 0x2f, 0x0a, 0x6c, 0x96, 0x04, 0xcd, 0x1e, 0x65, 0x03, 0xbf, 0xd9, 0x25, 0x01, 0xf7, 0xba,
	0x14, 0x37, 0x49, 0xc4, 0x71, 0x12, 0x79, 0xb4, 0xa9, 0x8e, 0x6e, 0x8a, 0x93, 0x53, 0xd2, 0xc3,
	0x6e, 0x9c, 0x30, 0xce, 0x2e, 0xa2, 0x5c, 0xfd, 0xeb, 0xf9, 0x21, 0x89, 0x6c, 0x75, 0x85, 0x1a,
	0x19, 0xf3, 0x28, 0xc4, 0xd6, 0x97, 0xa7, 0x2d, 0xf3, 0xf0, 0xaa, 0x39, 0x7d, 0x8f, 0x7b, 0x59,
	0x42, 0x4d, 0xa5, 0x6c, 0x73, 0x43, 0xf3, 0x34, 0xd5, 0xa9, 0x3b, 0xe8, 0x37, 0x71, 0x18, 0xf3,
	0xa1, 0xbe, 0xb4, 0x7e, 0x18, 0x80, 0xf6, 0x13, 0xec, 0x71, 0x7c, 0x24, 0x21, 0x0e, 0xfe, 0x34,
	0xc0, 0x29, 0x47, 0x08, 0xca, 0x91, 0x17, 0xe2, 0x86, 0xb1, 0x69, 0x6c, 0x2f, 0x39, 0xca, 0x46,
	0xeb, 0x50, 0xd1, 0x79, 0x88, 0xdf, 0x98, 0x57, 0xfe, 0xbf, 0xd4, 0xf9, 0xb9, 0x8f, 0x76, 0x61,
	0x41, 0x99, 0x8d, 0x92, 0xf0, 0x57, 0x5b, 0x37, 0xed, 0xbc, 0xa6, 0x6c, 0x9d, 0x45, 0x47, 0xa3,
	0xf7, 0x50, 0x27, 0x11, 0xe1, 0xc4, 0xa3, 0x6e, 0x1a, 0x53, 0xc2, 0xd3, 0x46, 0x79, 0xb3, 0x24,
	0xf0, 0xad, 0x7c, 0xfc, 0x64, 0xad, 0xf6, 0x6b, 0x09, 0x75, 0x96, 0x33, 0x26, 0x75, 0x4a, 0xcd,
	0x75, 0x58, 0x50, 0x16, 0xfa, 0x1b, 0x4a, 0x27, 0x78, 0xa8, 0x1a, 0xa9, 0x39, 0xd2, 0xb4, 0xbe,
	0x18, 0xf0, 0xef, 0x41, 0xc2, 0x62, 0x87, 0x9d, 0x39, 0x62, 0xb0, 0x85, 0x3d, 0x6f, 0x41, 0x3d,
	0x61, 0x67, 0xae, 0x80, 0x89, 0xc1, 0xe2, 0x3e, 0x39, 0x57, 0x9d, 0xd7, 0x9e, 0xcd, 0x39, 0x35,
	0xe1, 0x7f, 0x81, 0x87, 0x1d, 0xe5, 0x45, 0x8f, 0xc0, 0xf4, 0x31, 0xc5, 0x5c, 0xbc, 0x2d, 0xa5,
	0xfa, 0x21, 0xfa, 0x09, 0x0b, 0xdd, 0xcb, 0xa9, 0x54, 0x04, 0x66, 0x4d, 0xc7, 0xec, 0x51, 0x7a,
	0x20, 0x22, 0x0e, 0x45, 0x80, 0x6a, 0xa4, 0x5d, 0x81, 0x45, 0xee, 0x25, 0x01, 0xe6, 0xd6, 0x67,
	0x03, 0xfe, 0x79, 0x49, 0x52, 0xae, 0xfc, 0x69, 0x51, 0x65, 0x0f, 0xa0, 0x7c, 0x4a, 0xf0, 0x99,
	0xaa, 0xa7, 0xde, 0xba, 0x35, 0x65, 0xe2, 0xf6, 0x5b, 0x11, 0xeb, 0x28, 0x04, 0xba, 0x01, 0x10,
	0x7b, 0x81, 0xd0, 0x22, 0x3b, 0xc1, 0x91, 0xaa, 0x6d, 0xc9, 0x59, 0x92, 0x9e, 0x23, 0xe9, 0xb0,
	0x06, 0x80, 0xc6, 0x2b, 0x48, 0x63, 0x16, 0xa5, 0x18, 0xdd, 0x97, 0x25, 0x4a, 0x8f, 0x28, 0xa2,
	0x34, 0xcb, 0x13, 0x67, 0xe1, 0x62, 0x82, 0x2b, 0x11, 0x3e, 0xe7, 0xee, 0x58, 0x4a, 0x2d, 0x9e,
	0x65, 0xe9, 0xee, 0x5c, 0xa4, 0x75, 0x61, 0xe5, 0x29, 0xe6, 0x53, 0x45, 0x78, 0xed, 0xb6, 0xad,
	0x6d, 0x40, 0x07, 0x6a, 0xfc, 0xd3, 0x72, 0x58, 0xdf, 0xe7, 0x61, 0xe3, 0x15, 0xf3, 0x49, 0x7f,
	0xb8, 0xcf, 0xe8, 0x20, 0x8c, 0x0e, 0xbd, 0x90, 0x50, 0x52, 0xfc, 0x1c, 0x1f, 0x61, 0x39, 0x94,
	0x10, 0xd2, 0xf3, 0x38, 0x11, 0x13, 0x13, 0x05, 0xca, 0x31, 0xb5, 0xf3, 0x0b, 0x2c, 0xc8, 0xa0,
	0xef, 0x32, 0x2a, 0xe7, 0x57, 0x62, 0xf3, 0xab, 0x01, 0xb5, 0xf1, 0x7b, 0x54, 0x87, 0x79, 0xb1,
	0x91, 0xba, 0x18, 0x61, 0xa1, 0xc7, 0xb0, 0xd8, 0x53, 0x5b, 0xa2, 0x86, 0x54, 0x6d, 0x6d, 0x15,
	0x6c, 0xd3, 0x65, 0xf6, 0xa1, 0xd0, 0x67, 0x86, 0x93, 0x0c, 0x83, 0xd8, 0x97, 0x0c, 0xa5, 0xab,
	0x32, 0x68, 0x1c, 0x5a, 0x85, 0xb2, 0x2f, 0x56, 0x4c, 0xec, 0xb3, 0x56, 0xbe, 0x3a, 0xb5, 0x17,
	0xa0, 0x24, 0x7a, 0x69, 0x7d, 0x2b, 0x03, 0x6a, 0x67, 0x4c, 0xea, 0x31, 0xf6, 0x24, 0x1b, 0xfa,
	0x00, 0xd5, 0xb1, 0xed, 0x46, 0x3b, 0x57, 0xf9, 0x08, 0x98, 0xd3, 0xf4, 0x68, 0xcd, 0x21, 0x02,
	0x70, 0x29, 0x6c, 0x74, 0x27, 0x1f, 0x30, 0xb1, 0x80, 0xe6, 0xce, 0x6c, 0xc1, 0x7a, 0x57, 0x44,
	0xaa, 0x63, 0xa8, 0x8c, 0xc4, 0x8c, 0x6e, 0xe7, 0x63, 0x7f, 0x13, 0xfc, 0x2c, 0x4d, 0xbc, 0x81,
	0xea, 0x98, 0x8a, 0x8b, 0x86, 0x34, 0x29, 0x76, 0x73, 0x6d, 0x14, 0x3d, 0xfa, 0x2b, 0xb0, 0x9f,
	0xc8, 0xbf, 0x02, 0x41, 0x1b, 0xc3, 0xea, 0x9f, 0xf4, 0x88, 0x76, 0xaf, 0xa5, 0xdf, 0x59, 0x1a,
	0x79, 0x07, 0xb5, 0xf1, 0x8f, 0x30, 0xba, 0x5b, 0xd0, 0xc9, 0xe4, 0xc7, 0x3a, 0xbf, 0x95, 0xf6,
	0x43, 0xf8, 0xbf, 0xc7, 0xc2, 0x5c, 0xb6, 0xf6, 0x7f, 0x93, 0xd2, 0xeb, 0x48, 0x8e, 0x8e, 0xd1,
	0x5d, 0x54, 0x64, 0xf7, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x69, 0x2a, 0x0b, 0xee, 0x07,
	0x00, 0x00,
}

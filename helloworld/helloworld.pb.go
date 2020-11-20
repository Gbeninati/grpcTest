// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: helloworld/helloworld.proto

package helloworld

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PedidoPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor       string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda      string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioritario string `protobuf:"bytes,6,opt,name=prioritario,proto3" json:"prioritario,omitempty"`
}

func (x *PedidoPyme) Reset() {
	*x = PedidoPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedidoPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedidoPyme) ProtoMessage() {}

func (x *PedidoPyme) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedidoPyme.ProtoReflect.Descriptor instead.
func (*PedidoPyme) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *PedidoPyme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PedidoPyme) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *PedidoPyme) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *PedidoPyme) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *PedidoPyme) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *PedidoPyme) GetPrioritario() string {
	if x != nil {
		return x.Prioritario
	}
	return ""
}

type PedidoRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor    string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda   string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *PedidoRetail) Reset() {
	*x = PedidoRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedidoRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedidoRetail) ProtoMessage() {}

func (x *PedidoRetail) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedidoRetail.ProtoReflect.Descriptor instead.
func (*PedidoRetail) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *PedidoRetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PedidoRetail) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *PedidoRetail) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *PedidoRetail) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *PedidoRetail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type PaqueteCamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nro_Seguimiento string `protobuf:"bytes,2,opt,name=nro_Seguimiento,json=nroSeguimiento,proto3" json:"nro_Seguimiento,omitempty"`
	Tipo            string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor           string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos        string `protobuf:"bytes,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado          string `protobuf:"bytes,6,opt,name=estado,proto3" json:"estado,omitempty"`
	Tienda          string `protobuf:"bytes,7,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino         string `protobuf:"bytes,8,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *PaqueteCamion) Reset() {
	*x = PaqueteCamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaqueteCamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaqueteCamion) ProtoMessage() {}

func (x *PaqueteCamion) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaqueteCamion.ProtoReflect.Descriptor instead.
func (*PaqueteCamion) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *PaqueteCamion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaqueteCamion) GetNro_Seguimiento() string {
	if x != nil {
		return x.Nro_Seguimiento
	}
	return ""
}

func (x *PaqueteCamion) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *PaqueteCamion) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *PaqueteCamion) GetIntentos() string {
	if x != nil {
		return x.Intentos
	}
	return ""
}

func (x *PaqueteCamion) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *PaqueteCamion) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *PaqueteCamion) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type RespuestaPedido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seguimiento string `protobuf:"bytes,1,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
}

func (x *RespuestaPedido) Reset() {
	*x = RespuestaPedido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaPedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaPedido) ProtoMessage() {}

func (x *RespuestaPedido) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaPedido.ProtoReflect.Descriptor instead.
func (*RespuestaPedido) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{4}
}

func (x *RespuestaPedido) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

type NuevoEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Estado   string `protobuf:"bytes,2,opt,name=estado,proto3" json:"estado,omitempty"`
	Intentos string `protobuf:"bytes,3,opt,name=intentos,proto3" json:"intentos,omitempty"`
}

func (x *NuevoEstado) Reset() {
	*x = NuevoEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NuevoEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NuevoEstado) ProtoMessage() {}

func (x *NuevoEstado) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NuevoEstado.ProtoReflect.Descriptor instead.
func (*NuevoEstado) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{5}
}

func (x *NuevoEstado) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NuevoEstado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *NuevoEstado) GetIntentos() string {
	if x != nil {
		return x.Intentos
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{6}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa2, 0x01,
	0x0a, 0x0a, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72,
	0x69, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x71, 0x75,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x72, 0x6f,
	0x5f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6e, 0x72, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x6f, 0x22, 0x33, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50,
	0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x4e, 0x75, 0x65, 0x76, 0x6f,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x80, 0x03, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x48,
	0x0a, 0x0f, 0x4e, 0x75, 0x65, 0x76, 0x6f, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x50, 0x79, 0x6d,
	0x65, 0x12, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50,
	0x65, 0x64, 0x69, 0x64, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x4e, 0x75, 0x65, 0x76,
	0x6f, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64,
	0x6f, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x65,
	0x64, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0d, 0x53, 0x65, 0x67, 0x75, 0x69, 0x72,
	0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x65,
	0x64, 0x69, 0x64, 0x6f, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x65, 0x64, 0x69, 0x64,
	0x6f, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x08, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x12,
	0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x41, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x17, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4e, 0x75, 0x65, 0x76, 0x6f,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x65, 0x64,
	0x69, 0x64, 0x6f, 0x22, 0x00, 0x42, 0x5a, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x54, 0x61, 0x72, 0x65, 0x61, 0x32, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_helloworld_proto_rawDescData = file_helloworld_helloworld_proto_rawDesc
)

func file_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_helloworld_proto_rawDescData)
	})
	return file_helloworld_helloworld_proto_rawDescData
}

var file_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_helloworld_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),    // 0: helloworld.HelloRequest
	(*PedidoPyme)(nil),      // 1: helloworld.PedidoPyme
	(*PedidoRetail)(nil),    // 2: helloworld.PedidoRetail
	(*PaqueteCamion)(nil),   // 3: helloworld.PaqueteCamion
	(*RespuestaPedido)(nil), // 4: helloworld.RespuestaPedido
	(*NuevoEstado)(nil),     // 5: helloworld.NuevoEstado
	(*HelloReply)(nil),      // 6: helloworld.HelloReply
}
var file_helloworld_helloworld_proto_depIdxs = []int32{
	1, // 0: helloworld.Greeter.NuevoPedidoPyme:input_type -> helloworld.PedidoPyme
	2, // 1: helloworld.Greeter.NuevoPedidoRetail:input_type -> helloworld.PedidoRetail
	4, // 2: helloworld.Greeter.SeguirPaquete:input_type -> helloworld.RespuestaPedido
	3, // 3: helloworld.Greeter.Camiones:input_type -> helloworld.PaqueteCamion
	5, // 4: helloworld.Greeter.ActualizarEstado:input_type -> helloworld.NuevoEstado
	4, // 5: helloworld.Greeter.NuevoPedidoPyme:output_type -> helloworld.RespuestaPedido
	4, // 6: helloworld.Greeter.NuevoPedidoRetail:output_type -> helloworld.RespuestaPedido
	4, // 7: helloworld.Greeter.SeguirPaquete:output_type -> helloworld.RespuestaPedido
	4, // 8: helloworld.Greeter.Camiones:output_type -> helloworld.RespuestaPedido
	4, // 9: helloworld.Greeter.ActualizarEstado:output_type -> helloworld.RespuestaPedido
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_helloworld_proto_init() }
func file_helloworld_helloworld_proto_init() {
	if File_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedidoPyme); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedidoRetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaqueteCamion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaPedido); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NuevoEstado); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_helloworld_proto = out.File
	file_helloworld_helloworld_proto_rawDesc = nil
	file_helloworld_helloworld_proto_goTypes = nil
	file_helloworld_helloworld_proto_depIdxs = nil
}

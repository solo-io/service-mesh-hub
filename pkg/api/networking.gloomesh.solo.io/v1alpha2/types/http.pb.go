// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1alpha2/http.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
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

type HttpMethodValue int32

const (
	HttpMethodValue_GET     HttpMethodValue = 0
	HttpMethodValue_POST    HttpMethodValue = 1
	HttpMethodValue_PUT     HttpMethodValue = 2
	HttpMethodValue_DELETE  HttpMethodValue = 3
	HttpMethodValue_HEAD    HttpMethodValue = 4
	HttpMethodValue_CONNECT HttpMethodValue = 5
	HttpMethodValue_OPTIONS HttpMethodValue = 6
	HttpMethodValue_TRACE   HttpMethodValue = 7
	HttpMethodValue_PATCH   HttpMethodValue = 8
)

var HttpMethodValue_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "PUT",
	3: "DELETE",
	4: "HEAD",
	5: "CONNECT",
	6: "OPTIONS",
	7: "TRACE",
	8: "PATCH",
}

var HttpMethodValue_value = map[string]int32{
	"GET":     0,
	"POST":    1,
	"PUT":     2,
	"DELETE":  3,
	"HEAD":    4,
	"CONNECT": 5,
	"OPTIONS": 6,
	"TRACE":   7,
	"PATCH":   8,
}

func (x HttpMethodValue) String() string {
	return proto.EnumName(HttpMethodValue_name, int32(x))
}

func (HttpMethodValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c45bf20a0abe137b, []int{0}
}

func init() {
	proto.RegisterEnum("networking.gloomesh.solo.io.HttpMethodValue", HttpMethodValue_name, HttpMethodValue_value)
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/networking/v1alpha2/http.proto", fileDescriptor_c45bf20a0abe137b)
}

var fileDescriptor_c45bf20a0abe137b = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd5, 0xb6, 0x49, 0x5d, 0x0f, 0x0e, 0x7b, 0xf5, 0x0d, 0x84, 0xee, 0xa2, 0x3e, 0x80,
	0xc4, 0x74, 0x31, 0x82, 0x66, 0xa3, 0x1d, 0x3d, 0x78, 0x4b, 0x75, 0xc9, 0x2e, 0x4d, 0x3b, 0x4b,
	0x33, 0x55, 0x7c, 0x7b, 0x59, 0x2f, 0x05, 0x05, 0x6f, 0x33, 0xf3, 0xf3, 0x0d, 0xff, 0x27, 0xae,
	0xbb, 0xc0, 0x7e, 0xb7, 0x54, 0x6f, 0xb4, 0xd6, 0x03, 0xf5, 0x34, 0x0b, 0xa4, 0xbb, 0x9e, 0x68,
	0xb6, 0x76, 0x83, 0xd7, 0x6d, 0x0c, 0x7a, 0xe3, 0xf8, 0x93, 0xb6, 0xab, 0xb0, 0xe9, 0xf4, 0xc7,
	0x45, 0xdb, 0x47, 0xdf, 0x5e, 0x6a, 0xcf, 0x1c, 0x55, 0xdc, 0x12, 0x93, 0x3c, 0xdb, 0xe7, 0x2a,
	0x81, 0x89, 0x53, 0xe9, 0x93, 0x0a, 0x74, 0x3e, 0x88, 0xd3, 0x8a, 0x39, 0x3e, 0x38, 0xf6, 0xf4,
	0xfe, 0xd2, 0xf6, 0x3b, 0x27, 0x73, 0x31, 0xba, 0x35, 0x08, 0x07, 0x72, 0x2a, 0xc6, 0x8d, 0x5d,
	0x20, 0x1c, 0xa6, 0x53, 0xf3, 0x8c, 0x70, 0x24, 0x85, 0xc8, 0xe6, 0xe6, 0xde, 0xa0, 0x81, 0x51,
	0x8a, 0x2b, 0x53, 0xcc, 0x61, 0x2c, 0x4f, 0x44, 0x5e, 0xda, 0xba, 0x36, 0x25, 0xc2, 0x24, 0x2d,
	0xb6, 0xc1, 0x3b, 0x5b, 0x2f, 0x20, 0x93, 0xc7, 0x62, 0x82, 0x4f, 0x45, 0x69, 0x20, 0x4f, 0x63,
	0x53, 0x60, 0x59, 0xc1, 0xf4, 0xe6, 0xf1, 0xd5, 0xfe, 0x2b, 0x15, 0x57, 0xdd, 0x2f, 0xb1, 0x3f,
	0xc5, 0xf7, 0xa6, 0xfc, 0x15, 0xdd, 0xb0, 0xcc, 0x7e, 0x5c, 0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x65, 0x3e, 0x3a, 0x09, 0x2e, 0x01, 0x00, 0x00,
}

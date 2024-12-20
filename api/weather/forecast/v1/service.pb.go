// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: weather/forecast/v1/service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ForecastCharacterization int32

const (
	// It's best practice to have a 0'ith value for enums when using protocol buffers
	ForecastCharacterization_INVALID  ForecastCharacterization = 0
	ForecastCharacterization_COLD     ForecastCharacterization = 1
	ForecastCharacterization_MODERATE ForecastCharacterization = 2
	ForecastCharacterization_HOT      ForecastCharacterization = 3
)

// Enum value maps for ForecastCharacterization.
var (
	ForecastCharacterization_name = map[int32]string{
		0: "INVALID",
		1: "COLD",
		2: "MODERATE",
		3: "HOT",
	}
	ForecastCharacterization_value = map[string]int32{
		"INVALID":  0,
		"COLD":     1,
		"MODERATE": 2,
		"HOT":      3,
	}
)

func (x ForecastCharacterization) Enum() *ForecastCharacterization {
	p := new(ForecastCharacterization)
	*p = x
	return p
}

func (x ForecastCharacterization) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ForecastCharacterization) Descriptor() protoreflect.EnumDescriptor {
	return file_weather_forecast_v1_service_proto_enumTypes[0].Descriptor()
}

func (ForecastCharacterization) Type() protoreflect.EnumType {
	return &file_weather_forecast_v1_service_proto_enumTypes[0]
}

func (x ForecastCharacterization) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ForecastCharacterization.Descriptor instead.
func (ForecastCharacterization) EnumDescriptor() ([]byte, []int) {
	return file_weather_forecast_v1_service_proto_rawDescGZIP(), []int{0}
}

type ShortForecastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latitudes should be in the range +/- 90 degrees
	Latitude float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude" bun:"latitude" csv:"latitude" pg:"latitude" yaml:"latitude"`
	// Longitudes should be in the range +/- 180 degrees
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude" bun:"longitude" csv:"longitude" pg:"longitude" yaml:"longitude"`
}

func (x *ShortForecastRequest) Reset() {
	*x = ShortForecastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_forecast_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortForecastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortForecastRequest) ProtoMessage() {}

func (x *ShortForecastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_forecast_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortForecastRequest.ProtoReflect.Descriptor instead.
func (*ShortForecastRequest) Descriptor() ([]byte, []int) {
	return file_weather_forecast_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ShortForecastRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ShortForecastRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type ShortForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortForecast       string                   `protobuf:"bytes,1,opt,name=short_forecast,json=shortForecast,proto3" json:"short_forecast" bun:"short_forecast" csv:"short_forecast" pg:"short_forecast" yaml:"short_forecast"`
	TemperatureCategory ForecastCharacterization `protobuf:"varint,2,opt,name=temperature_category,json=temperatureCategory,proto3,enum=weather.forecast.v1.ForecastCharacterization" json:"temperature_category" bun:"temperature_category" csv:"temperature_category" pg:"temperature_category" yaml:"temperature_category"`
}

func (x *ShortForecastResponse) Reset() {
	*x = ShortForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_forecast_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortForecastResponse) ProtoMessage() {}

func (x *ShortForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_forecast_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortForecastResponse.ProtoReflect.Descriptor instead.
func (*ShortForecastResponse) Descriptor() ([]byte, []int) {
	return file_weather_forecast_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShortForecastResponse) GetShortForecast() string {
	if x != nil {
		return x.ShortForecast
	}
	return ""
}

func (x *ShortForecastResponse) GetTemperatureCategory() ForecastCharacterization {
	if x != nil {
		return x.TemperatureCategory
	}
	return ForecastCharacterization_INVALID
}

var File_weather_forecast_v1_service_proto protoreflect.FileDescriptor

var file_weather_forecast_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x72, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0f, 0xfa, 0x42, 0x0c,
	0x0a, 0x0a, 0x1d, 0x00, 0x00, 0xb4, 0x42, 0x2d, 0x00, 0x00, 0xb4, 0xc2, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x0a, 0x0a,
	0x1d, 0x00, 0x00, 0x34, 0x43, 0x2d, 0x00, 0x00, 0x34, 0xc3, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x15, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x66,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2a, 0x48, 0x0a, 0x18, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x4f, 0x54,
	0x10, 0x03, 0x32, 0x79, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x2e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72,
	0x65, 0x77, 0x53, 0x43, 0x32, 0x30, 0x38, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f,
	0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_weather_forecast_v1_service_proto_rawDescOnce sync.Once
	file_weather_forecast_v1_service_proto_rawDescData = file_weather_forecast_v1_service_proto_rawDesc
)

func file_weather_forecast_v1_service_proto_rawDescGZIP() []byte {
	file_weather_forecast_v1_service_proto_rawDescOnce.Do(func() {
		file_weather_forecast_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_forecast_v1_service_proto_rawDescData)
	})
	return file_weather_forecast_v1_service_proto_rawDescData
}

var file_weather_forecast_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_weather_forecast_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_weather_forecast_v1_service_proto_goTypes = []any{
	(ForecastCharacterization)(0), // 0: weather.forecast.v1.ForecastCharacterization
	(*ShortForecastRequest)(nil),  // 1: weather.forecast.v1.ShortForecastRequest
	(*ShortForecastResponse)(nil), // 2: weather.forecast.v1.ShortForecastResponse
}
var file_weather_forecast_v1_service_proto_depIdxs = []int32{
	0, // 0: weather.forecast.v1.ShortForecastResponse.temperature_category:type_name -> weather.forecast.v1.ForecastCharacterization
	1, // 1: weather.forecast.v1.ForecastService.ShortForecast:input_type -> weather.forecast.v1.ShortForecastRequest
	2, // 2: weather.forecast.v1.ForecastService.ShortForecast:output_type -> weather.forecast.v1.ShortForecastResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_weather_forecast_v1_service_proto_init() }
func file_weather_forecast_v1_service_proto_init() {
	if File_weather_forecast_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_forecast_v1_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShortForecastRequest); i {
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
		file_weather_forecast_v1_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ShortForecastResponse); i {
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
			RawDescriptor: file_weather_forecast_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_forecast_v1_service_proto_goTypes,
		DependencyIndexes: file_weather_forecast_v1_service_proto_depIdxs,
		EnumInfos:         file_weather_forecast_v1_service_proto_enumTypes,
		MessageInfos:      file_weather_forecast_v1_service_proto_msgTypes,
	}.Build()
	File_weather_forecast_v1_service_proto = out.File
	file_weather_forecast_v1_service_proto_rawDesc = nil
	file_weather_forecast_v1_service_proto_goTypes = nil
	file_weather_forecast_v1_service_proto_depIdxs = nil
}

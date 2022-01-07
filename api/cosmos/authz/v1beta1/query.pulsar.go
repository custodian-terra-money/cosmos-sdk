package authzv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_QueryGrantsRequest              protoreflect.MessageDescriptor
	fd_QueryGrantsRequest_granter      protoreflect.FieldDescriptor
	fd_QueryGrantsRequest_grantee      protoreflect.FieldDescriptor
	fd_QueryGrantsRequest_msg_type_url protoreflect.FieldDescriptor
	fd_QueryGrantsRequest_pagination   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_authz_v1beta1_query_proto_init()
	md_QueryGrantsRequest = File_cosmos_authz_v1beta1_query_proto.Messages().ByName("QueryGrantsRequest")
	fd_QueryGrantsRequest_granter = md_QueryGrantsRequest.Fields().ByName("granter")
	fd_QueryGrantsRequest_grantee = md_QueryGrantsRequest.Fields().ByName("grantee")
	fd_QueryGrantsRequest_msg_type_url = md_QueryGrantsRequest.Fields().ByName("msg_type_url")
	fd_QueryGrantsRequest_pagination = md_QueryGrantsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryGrantsRequest)(nil)

type fastReflection_QueryGrantsRequest QueryGrantsRequest

func (x *QueryGrantsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryGrantsRequest)(x)
}

func (x *QueryGrantsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryGrantsRequest_messageType fastReflection_QueryGrantsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryGrantsRequest_messageType{}

type fastReflection_QueryGrantsRequest_messageType struct{}

func (x fastReflection_QueryGrantsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryGrantsRequest)(nil)
}
func (x fastReflection_QueryGrantsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryGrantsRequest)
}
func (x fastReflection_QueryGrantsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGrantsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryGrantsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGrantsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryGrantsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryGrantsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryGrantsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryGrantsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryGrantsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryGrantsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryGrantsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_QueryGrantsRequest_granter, value) {
			return
		}
	}
	if x.Grantee != "" {
		value := protoreflect.ValueOfString(x.Grantee)
		if !f(fd_QueryGrantsRequest_grantee, value) {
			return
		}
	}
	if x.MsgTypeUrl != "" {
		value := protoreflect.ValueOfString(x.MsgTypeUrl)
		if !f(fd_QueryGrantsRequest_msg_type_url, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryGrantsRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryGrantsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsRequest.granter":
		return x.Granter != ""
	case "cosmos.authz.v1beta1.QueryGrantsRequest.grantee":
		return x.Grantee != ""
	case "cosmos.authz.v1beta1.QueryGrantsRequest.msg_type_url":
		return x.MsgTypeUrl != ""
	case "cosmos.authz.v1beta1.QueryGrantsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsRequest.granter":
		x.Granter = ""
	case "cosmos.authz.v1beta1.QueryGrantsRequest.grantee":
		x.Grantee = ""
	case "cosmos.authz.v1beta1.QueryGrantsRequest.msg_type_url":
		x.MsgTypeUrl = ""
	case "cosmos.authz.v1beta1.QueryGrantsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryGrantsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsRequest.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.authz.v1beta1.QueryGrantsRequest.grantee":
		value := x.Grantee
		return protoreflect.ValueOfString(value)
	case "cosmos.authz.v1beta1.QueryGrantsRequest.msg_type_url":
		value := x.MsgTypeUrl
		return protoreflect.ValueOfString(value)
	case "cosmos.authz.v1beta1.QueryGrantsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsRequest.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.authz.v1beta1.QueryGrantsRequest.grantee":
		x.Grantee = value.Interface().(string)
	case "cosmos.authz.v1beta1.QueryGrantsRequest.msg_type_url":
		x.MsgTypeUrl = value.Interface().(string)
	case "cosmos.authz.v1beta1.QueryGrantsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.authz.v1beta1.QueryGrantsRequest.granter":
		panic(fmt.Errorf("field granter of message cosmos.authz.v1beta1.QueryGrantsRequest is not mutable"))
	case "cosmos.authz.v1beta1.QueryGrantsRequest.grantee":
		panic(fmt.Errorf("field grantee of message cosmos.authz.v1beta1.QueryGrantsRequest is not mutable"))
	case "cosmos.authz.v1beta1.QueryGrantsRequest.msg_type_url":
		panic(fmt.Errorf("field msg_type_url of message cosmos.authz.v1beta1.QueryGrantsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryGrantsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsRequest.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.authz.v1beta1.QueryGrantsRequest.grantee":
		return protoreflect.ValueOfString("")
	case "cosmos.authz.v1beta1.QueryGrantsRequest.msg_type_url":
		return protoreflect.ValueOfString("")
	case "cosmos.authz.v1beta1.QueryGrantsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryGrantsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.authz.v1beta1.QueryGrantsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryGrantsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryGrantsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryGrantsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryGrantsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Grantee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MsgTypeUrl)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryGrantsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.MsgTypeUrl) > 0 {
			i -= len(x.MsgTypeUrl)
			copy(dAtA[i:], x.MsgTypeUrl)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MsgTypeUrl)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Grantee) > 0 {
			i -= len(x.Grantee)
			copy(dAtA[i:], x.Grantee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Grantee)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryGrantsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGrantsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGrantsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Grantee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryGrantsResponse_1_list)(nil)

type _QueryGrantsResponse_1_list struct {
	list *[]*Grant
}

func (x *_QueryGrantsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryGrantsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryGrantsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	(*x.list)[i] = concreteValue
}

func (x *_QueryGrantsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryGrantsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Grant)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryGrantsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryGrantsResponse_1_list) NewElement() protoreflect.Value {
	v := new(Grant)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryGrantsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryGrantsResponse            protoreflect.MessageDescriptor
	fd_QueryGrantsResponse_grants     protoreflect.FieldDescriptor
	fd_QueryGrantsResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_authz_v1beta1_query_proto_init()
	md_QueryGrantsResponse = File_cosmos_authz_v1beta1_query_proto.Messages().ByName("QueryGrantsResponse")
	fd_QueryGrantsResponse_grants = md_QueryGrantsResponse.Fields().ByName("grants")
	fd_QueryGrantsResponse_pagination = md_QueryGrantsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryGrantsResponse)(nil)

type fastReflection_QueryGrantsResponse QueryGrantsResponse

func (x *QueryGrantsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryGrantsResponse)(x)
}

func (x *QueryGrantsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryGrantsResponse_messageType fastReflection_QueryGrantsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryGrantsResponse_messageType{}

type fastReflection_QueryGrantsResponse_messageType struct{}

func (x fastReflection_QueryGrantsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryGrantsResponse)(nil)
}
func (x fastReflection_QueryGrantsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryGrantsResponse)
}
func (x fastReflection_QueryGrantsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGrantsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryGrantsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGrantsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryGrantsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryGrantsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryGrantsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryGrantsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryGrantsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryGrantsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryGrantsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Grants) != 0 {
		value := protoreflect.ValueOfList(&_QueryGrantsResponse_1_list{list: &x.Grants})
		if !f(fd_QueryGrantsResponse_grants, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryGrantsResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryGrantsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsResponse.grants":
		return len(x.Grants) != 0
	case "cosmos.authz.v1beta1.QueryGrantsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsResponse.grants":
		x.Grants = nil
	case "cosmos.authz.v1beta1.QueryGrantsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryGrantsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsResponse.grants":
		if len(x.Grants) == 0 {
			return protoreflect.ValueOfList(&_QueryGrantsResponse_1_list{})
		}
		listValue := &_QueryGrantsResponse_1_list{list: &x.Grants}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.authz.v1beta1.QueryGrantsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsResponse.grants":
		lv := value.List()
		clv := lv.(*_QueryGrantsResponse_1_list)
		x.Grants = *clv.list
	case "cosmos.authz.v1beta1.QueryGrantsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsResponse.grants":
		if x.Grants == nil {
			x.Grants = []*Grant{}
		}
		value := &_QueryGrantsResponse_1_list{list: &x.Grants}
		return protoreflect.ValueOfList(value)
	case "cosmos.authz.v1beta1.QueryGrantsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryGrantsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGrantsResponse.grants":
		list := []*Grant{}
		return protoreflect.ValueOfList(&_QueryGrantsResponse_1_list{list: &list})
	case "cosmos.authz.v1beta1.QueryGrantsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryGrantsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.authz.v1beta1.QueryGrantsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryGrantsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGrantsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryGrantsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryGrantsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryGrantsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Grants) > 0 {
			for _, e := range x.Grants {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryGrantsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Grants) > 0 {
			for iNdEx := len(x.Grants) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Grants[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryGrantsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGrantsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGrantsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Grants = append(x.Grants, &Grant{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Grants[len(x.Grants)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryGranterGrantsRequest            protoreflect.MessageDescriptor
	fd_QueryGranterGrantsRequest_granter    protoreflect.FieldDescriptor
	fd_QueryGranterGrantsRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_authz_v1beta1_query_proto_init()
	md_QueryGranterGrantsRequest = File_cosmos_authz_v1beta1_query_proto.Messages().ByName("QueryGranterGrantsRequest")
	fd_QueryGranterGrantsRequest_granter = md_QueryGranterGrantsRequest.Fields().ByName("granter")
	fd_QueryGranterGrantsRequest_pagination = md_QueryGranterGrantsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryGranterGrantsRequest)(nil)

type fastReflection_QueryGranterGrantsRequest QueryGranterGrantsRequest

func (x *QueryGranterGrantsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryGranterGrantsRequest)(x)
}

func (x *QueryGranterGrantsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryGranterGrantsRequest_messageType fastReflection_QueryGranterGrantsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryGranterGrantsRequest_messageType{}

type fastReflection_QueryGranterGrantsRequest_messageType struct{}

func (x fastReflection_QueryGranterGrantsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryGranterGrantsRequest)(nil)
}
func (x fastReflection_QueryGranterGrantsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryGranterGrantsRequest)
}
func (x fastReflection_QueryGranterGrantsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGranterGrantsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryGranterGrantsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGranterGrantsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryGranterGrantsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryGranterGrantsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryGranterGrantsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryGranterGrantsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryGranterGrantsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryGranterGrantsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryGranterGrantsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_QueryGranterGrantsRequest_granter, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryGranterGrantsRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryGranterGrantsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.granter":
		return x.Granter != ""
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.granter":
		x.Granter = ""
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryGranterGrantsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.granter":
		panic(fmt.Errorf("field granter of message cosmos.authz.v1beta1.QueryGranterGrantsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryGranterGrantsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsRequest"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryGranterGrantsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.authz.v1beta1.QueryGranterGrantsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryGranterGrantsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryGranterGrantsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryGranterGrantsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryGranterGrantsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryGranterGrantsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryGranterGrantsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGranterGrantsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGranterGrantsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryGranterGrantsResponse_1_list)(nil)

type _QueryGranterGrantsResponse_1_list struct {
	list *[]*Grant
}

func (x *_QueryGranterGrantsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryGranterGrantsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryGranterGrantsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	(*x.list)[i] = concreteValue
}

func (x *_QueryGranterGrantsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryGranterGrantsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Grant)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryGranterGrantsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryGranterGrantsResponse_1_list) NewElement() protoreflect.Value {
	v := new(Grant)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryGranterGrantsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryGranterGrantsResponse            protoreflect.MessageDescriptor
	fd_QueryGranterGrantsResponse_grants     protoreflect.FieldDescriptor
	fd_QueryGranterGrantsResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_authz_v1beta1_query_proto_init()
	md_QueryGranterGrantsResponse = File_cosmos_authz_v1beta1_query_proto.Messages().ByName("QueryGranterGrantsResponse")
	fd_QueryGranterGrantsResponse_grants = md_QueryGranterGrantsResponse.Fields().ByName("grants")
	fd_QueryGranterGrantsResponse_pagination = md_QueryGranterGrantsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryGranterGrantsResponse)(nil)

type fastReflection_QueryGranterGrantsResponse QueryGranterGrantsResponse

func (x *QueryGranterGrantsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryGranterGrantsResponse)(x)
}

func (x *QueryGranterGrantsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryGranterGrantsResponse_messageType fastReflection_QueryGranterGrantsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryGranterGrantsResponse_messageType{}

type fastReflection_QueryGranterGrantsResponse_messageType struct{}

func (x fastReflection_QueryGranterGrantsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryGranterGrantsResponse)(nil)
}
func (x fastReflection_QueryGranterGrantsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryGranterGrantsResponse)
}
func (x fastReflection_QueryGranterGrantsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGranterGrantsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryGranterGrantsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryGranterGrantsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryGranterGrantsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryGranterGrantsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryGranterGrantsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryGranterGrantsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryGranterGrantsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryGranterGrantsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryGranterGrantsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Grants) != 0 {
		value := protoreflect.ValueOfList(&_QueryGranterGrantsResponse_1_list{list: &x.Grants})
		if !f(fd_QueryGranterGrantsResponse_grants, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryGranterGrantsResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryGranterGrantsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants":
		return len(x.Grants) != 0
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants":
		x.Grants = nil
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryGranterGrantsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants":
		if len(x.Grants) == 0 {
			return protoreflect.ValueOfList(&_QueryGranterGrantsResponse_1_list{})
		}
		listValue := &_QueryGranterGrantsResponse_1_list{list: &x.Grants}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants":
		lv := value.List()
		clv := lv.(*_QueryGranterGrantsResponse_1_list)
		x.Grants = *clv.list
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants":
		if x.Grants == nil {
			x.Grants = []*Grant{}
		}
		value := &_QueryGranterGrantsResponse_1_list{list: &x.Grants}
		return protoreflect.ValueOfList(value)
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryGranterGrantsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants":
		list := []*Grant{}
		return protoreflect.ValueOfList(&_QueryGranterGrantsResponse_1_list{list: &list})
	case "cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.authz.v1beta1.QueryGranterGrantsResponse"))
		}
		panic(fmt.Errorf("message cosmos.authz.v1beta1.QueryGranterGrantsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryGranterGrantsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.authz.v1beta1.QueryGranterGrantsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryGranterGrantsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryGranterGrantsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryGranterGrantsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryGranterGrantsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryGranterGrantsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Grants) > 0 {
			for _, e := range x.Grants {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryGranterGrantsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Grants) > 0 {
			for iNdEx := len(x.Grants) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Grants[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryGranterGrantsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGranterGrantsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryGranterGrantsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Grants = append(x.Grants, &Grant{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Grants[len(x.Grants)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error)
	// GranterGrants returns list of `Authorization`, granted by granter.
	GranterGrants(ctx context.Context, in *QueryGranterGrantsRequest, opts ...grpc.CallOption) (*QueryGranterGrantsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error) {
	out := new(QueryGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/Grants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GranterGrants(ctx context.Context, in *QueryGranterGrantsRequest, opts ...grpc.CallOption) (*QueryGranterGrantsResponse, error) {
	out := new(QueryGranterGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/GranterGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error)
	// GranterGrants returns list of `Authorization`, granted by granter.
	GranterGrants(context.Context, *QueryGranterGrantsRequest) (*QueryGranterGrantsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grants not implemented")
}
func (UnimplementedQueryServer) GranterGrants(context.Context, *QueryGranterGrantsRequest) (*QueryGranterGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GranterGrants not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Grants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Grants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/Grants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Grants(ctx, req.(*QueryGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GranterGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGranterGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GranterGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/GranterGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GranterGrants(ctx, req.(*QueryGranterGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grants",
			Handler:    _Query_Grants_Handler,
		},
		{
			MethodName: "GranterGrants",
			Handler:    _Query_GranterGrants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/authz/v1beta1/query.proto",
}

// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/authz/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryGrantsRequest is the request type for the Query/Grants RPC method.
type QueryGrantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Optional, msg_type_url, when set, will query only grants matching given msg type.
	MsgTypeUrl string `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryGrantsRequest) Reset() {
	*x = QueryGrantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGrantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGrantsRequest) ProtoMessage() {}

// Deprecated: Use QueryGrantsRequest.ProtoReflect.Descriptor instead.
func (*QueryGrantsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryGrantsRequest) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *QueryGrantsRequest) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *QueryGrantsRequest) GetMsgTypeUrl() string {
	if x != nil {
		return x.MsgTypeUrl
	}
	return ""
}

func (x *QueryGrantsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryGrantsResponse is the response type for the Query/Authorizations RPC method.
type QueryGrantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// authorizations is a list of grants granted for grantee by granter.
	Grants []*Grant `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryGrantsResponse) Reset() {
	*x = QueryGrantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGrantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGrantsResponse) ProtoMessage() {}

// Deprecated: Use QueryGrantsResponse.ProtoReflect.Descriptor instead.
func (*QueryGrantsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryGrantsResponse) GetGrants() []*Grant {
	if x != nil {
		return x.Grants
	}
	return nil
}

func (x *QueryGrantsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryGranterGrantsRequest is the request type for the Query/GranterGrants RPC method.
type QueryGranterGrantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryGranterGrantsRequest) Reset() {
	*x = QueryGranterGrantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGranterGrantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGranterGrantsRequest) ProtoMessage() {}

// Deprecated: Use QueryGranterGrantsRequest.ProtoReflect.Descriptor instead.
func (*QueryGranterGrantsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryGranterGrantsRequest) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *QueryGranterGrantsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryGranterGrantsResponse is the response type for the Query/GranterGrants RPC method.
type QueryGranterGrantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// authorizations is a list of grants granted for grantee by granter.
	Grants []*Grant `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryGranterGrantsResponse) Reset() {
	*x = QueryGranterGrantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGranterGrantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGranterGrantsResponse) ProtoMessage() {}

// Deprecated: Use QueryGranterGrantsResponse.ProtoReflect.Descriptor instead.
func (*QueryGranterGrantsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryGranterGrantsResponse) GetGrants() []*Grant {
	if x != nil {
		return x.Grants
	}
	return nil
}

func (x *QueryGranterGrantsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_cosmos_authz_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_authz_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe6, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d,
	0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x20,
	0x0a, 0x0c, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97,
	0x01, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb2, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x83, 0x01, 0x0a, 0x06, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0xa2, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x72, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x12, 0x26, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73,
	0x2f, 0x7b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x7d, 0x42, 0xdc, 0x01, 0x0a, 0x18, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa,
	0x02, 0x14, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x14, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x20,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cosmos_authz_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_authz_v1beta1_query_proto_rawDescData = file_cosmos_authz_v1beta1_query_proto_rawDesc
)

func file_cosmos_authz_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_authz_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_authz_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_authz_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_authz_v1beta1_query_proto_rawDescData
}

var file_cosmos_authz_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_authz_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryGrantsRequest)(nil),         // 0: cosmos.authz.v1beta1.QueryGrantsRequest
	(*QueryGrantsResponse)(nil),        // 1: cosmos.authz.v1beta1.QueryGrantsResponse
	(*QueryGranterGrantsRequest)(nil),  // 2: cosmos.authz.v1beta1.QueryGranterGrantsRequest
	(*QueryGranterGrantsResponse)(nil), // 3: cosmos.authz.v1beta1.QueryGranterGrantsResponse
	(*v1beta1.PageRequest)(nil),        // 4: cosmos.base.query.v1beta1.PageRequest
	(*Grant)(nil),                      // 5: cosmos.authz.v1beta1.Grant
	(*v1beta1.PageResponse)(nil),       // 6: cosmos.base.query.v1beta1.PageResponse
}
var file_cosmos_authz_v1beta1_query_proto_depIdxs = []int32{
	4, // 0: cosmos.authz.v1beta1.QueryGrantsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	5, // 1: cosmos.authz.v1beta1.QueryGrantsResponse.grants:type_name -> cosmos.authz.v1beta1.Grant
	6, // 2: cosmos.authz.v1beta1.QueryGrantsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	4, // 3: cosmos.authz.v1beta1.QueryGranterGrantsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	5, // 4: cosmos.authz.v1beta1.QueryGranterGrantsResponse.grants:type_name -> cosmos.authz.v1beta1.Grant
	6, // 5: cosmos.authz.v1beta1.QueryGranterGrantsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	0, // 6: cosmos.authz.v1beta1.Query.Grants:input_type -> cosmos.authz.v1beta1.QueryGrantsRequest
	2, // 7: cosmos.authz.v1beta1.Query.GranterGrants:input_type -> cosmos.authz.v1beta1.QueryGranterGrantsRequest
	1, // 8: cosmos.authz.v1beta1.Query.Grants:output_type -> cosmos.authz.v1beta1.QueryGrantsResponse
	3, // 9: cosmos.authz.v1beta1.Query.GranterGrants:output_type -> cosmos.authz.v1beta1.QueryGranterGrantsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cosmos_authz_v1beta1_query_proto_init() }
func file_cosmos_authz_v1beta1_query_proto_init() {
	if File_cosmos_authz_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_authz_v1beta1_authz_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_authz_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGrantsRequest); i {
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
		file_cosmos_authz_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGrantsResponse); i {
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
		file_cosmos_authz_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGranterGrantsRequest); i {
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
		file_cosmos_authz_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGranterGrantsResponse); i {
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
			RawDescriptor: file_cosmos_authz_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_authz_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_authz_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_authz_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_authz_v1beta1_query_proto = out.File
	file_cosmos_authz_v1beta1_query_proto_rawDesc = nil
	file_cosmos_authz_v1beta1_query_proto_goTypes = nil
	file_cosmos_authz_v1beta1_query_proto_depIdxs = nil
}
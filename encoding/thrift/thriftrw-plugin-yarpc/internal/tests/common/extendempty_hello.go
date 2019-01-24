// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

package common

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// ExtendEmpty_Hello_Args represents the arguments for the ExtendEmpty.hello function.
//
// The arguments for hello are sent and received over the wire as this struct.
type ExtendEmpty_Hello_Args struct {
}

// ToWire translates a ExtendEmpty_Hello_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ExtendEmpty_Hello_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ExtendEmpty_Hello_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ExtendEmpty_Hello_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ExtendEmpty_Hello_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ExtendEmpty_Hello_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a ExtendEmpty_Hello_Args
// struct.
func (v *ExtendEmpty_Hello_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("ExtendEmpty_Hello_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ExtendEmpty_Hello_Args match the
// provided ExtendEmpty_Hello_Args.
//
// This function performs a deep comparison.
func (v *ExtendEmpty_Hello_Args) Equals(rhs *ExtendEmpty_Hello_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExtendEmpty_Hello_Args.
func (v *ExtendEmpty_Hello_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "hello" for this struct.
func (v *ExtendEmpty_Hello_Args) MethodName() string {
	return "hello"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ExtendEmpty_Hello_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ExtendEmpty_Hello_Helper provides functions that aid in handling the
// parameters and return values of the ExtendEmpty.hello
// function.
var ExtendEmpty_Hello_Helper = struct {
	// Args accepts the parameters of hello in-order and returns
	// the arguments struct for the function.
	Args func() *ExtendEmpty_Hello_Args

	// IsException returns true if the given error can be thrown
	// by hello.
	//
	// An error can be thrown by hello only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for hello
	// given the error returned by it. The provided error may
	// be nil if hello did not fail.
	//
	// This allows mapping errors returned by hello into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// hello
	//
	//   err := hello(args)
	//   result, err := ExtendEmpty_Hello_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from hello: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*ExtendEmpty_Hello_Result, error)

	// UnwrapResponse takes the result struct for hello
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if hello threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := ExtendEmpty_Hello_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ExtendEmpty_Hello_Result) error
}{}

func init() {
	ExtendEmpty_Hello_Helper.Args = func() *ExtendEmpty_Hello_Args {
		return &ExtendEmpty_Hello_Args{}
	}

	ExtendEmpty_Hello_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ExtendEmpty_Hello_Helper.WrapResponse = func(err error) (*ExtendEmpty_Hello_Result, error) {
		if err == nil {
			return &ExtendEmpty_Hello_Result{}, nil
		}

		return nil, err
	}
	ExtendEmpty_Hello_Helper.UnwrapResponse = func(result *ExtendEmpty_Hello_Result) (err error) {
		return
	}

}

// ExtendEmpty_Hello_Result represents the result of a ExtendEmpty.hello function call.
//
// The result of a hello execution is sent and received over the wire as this struct.
type ExtendEmpty_Hello_Result struct {
}

// ToWire translates a ExtendEmpty_Hello_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ExtendEmpty_Hello_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ExtendEmpty_Hello_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ExtendEmpty_Hello_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ExtendEmpty_Hello_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ExtendEmpty_Hello_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a ExtendEmpty_Hello_Result
// struct.
func (v *ExtendEmpty_Hello_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("ExtendEmpty_Hello_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ExtendEmpty_Hello_Result match the
// provided ExtendEmpty_Hello_Result.
//
// This function performs a deep comparison.
func (v *ExtendEmpty_Hello_Result) Equals(rhs *ExtendEmpty_Hello_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExtendEmpty_Hello_Result.
func (v *ExtendEmpty_Hello_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "hello" for this struct.
func (v *ExtendEmpty_Hello_Result) MethodName() string {
	return "hello"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ExtendEmpty_Hello_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

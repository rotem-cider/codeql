// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for k8s.io/apimachinery/pkg/runtime, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: k8s.io/apimachinery/pkg/runtime (exports: ProtobufMarshaller,ProtobufReverseMarshaller; functions: )

// Package runtime is a stub of k8s.io/apimachinery/pkg/runtime, generated by depstubber.
package runtime

type ProtobufMarshaller interface {
	MarshalTo(_ []byte) (int, error)
}

type ProtobufReverseMarshaller interface {
	MarshalToSizedBuffer(_ []byte) (int, error)
}

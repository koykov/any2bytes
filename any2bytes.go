package any2bytes

import "errors"

// Signature of conversion function.
type AnyToBytesFn func(dst []byte, val interface{}) ([]byte, error)

var (
	// Registry of conversion functions.
	anyToBytesFnRegistry = make([]AnyToBytesFn, 0)

	ErrUnknownType = errors.New("unknown type")
)

func init() {
	// Register conversion functions from builtin types.
	RegisterAnyToBytesFn(BytesToBytes)
	RegisterAnyToBytesFn(StrToBytes)
	RegisterAnyToBytesFn(BoolToBytes)
	RegisterAnyToBytesFn(IntToBytes)
	RegisterAnyToBytesFn(UintToBytes)
	RegisterAnyToBytesFn(FloatToBytes)
}

// Register new conversion function.
func RegisterAnyToBytesFn(fn AnyToBytesFn) {
	for _, f := range anyToBytesFnRegistry {
		if &f == &fn {
			return
		}
	}
	anyToBytesFnRegistry = append(anyToBytesFnRegistry, fn)
}

// Generic conversion function.
//
// Convert val to byte array and append result to the dst.
// Returns dst and conversion error message. Error is nil when succeed.
func AnyToBytes(dst []byte, val interface{}) ([]byte, error) {
	var err error
	if dst == nil {
		dst = make([]byte, 0)
	}
	dst = dst[:0]
	for _, fn := range anyToBytesFnRegistry {
		dst, err = fn(dst, val)
		if err == nil {
			return dst, nil
		}
		if err != ErrUnknownType {
			return dst, err
		}
	}
	return dst, ErrUnknownType
}

package types

import (
	"fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Epoch)(0)
var _ fssz.Marshaler = (*Epoch)(nil)
var _ fssz.Unmarshaler = (*Epoch)(nil)

// Epoch represents a single epoch.
type Epoch uint64

// ToEpoch returns x converted to Epoch.
func ToEpoch(x uint64) Epoch {
	return Epoch(x)
}

// Uint64 returns epoch as underlying type.
func (e Epoch) Uint64() uint64 {
	return uint64(e)
}

// Mul multiplies epoch by x.
func (e Epoch) Mul(x uint64) Epoch {
	return Epoch(uint64(e) * x)
}

// Div divides epoch by x.
func (e Epoch) Div(x uint64) Epoch {
	if x == 0 {
		panic("divbyzero")
	}
	return Epoch(uint64(e) / x)
}

// Add increases epoch by x.
func (e Epoch) Add(x uint64) Epoch {
	return Epoch(uint64(e) + x)
}

// Sub subtracts x from the epoch.
func (e Epoch) Sub(x uint64) Epoch {
	if uint64(e) < x {
		panic("underflow")
	}
	return Epoch(uint64(e) - x)
}

// Mod returns result of `epoch % x`.
func (e Epoch) Mod(x uint64) Epoch {
	return Epoch(uint64(e) % x)
}

// HashTreeRoot returns calculated hash root.
func (e Epoch) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(e)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (e Epoch) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(e.Uint64())
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the epoch object.
func (e *Epoch) UnmarshalSSZ(buf []byte) error {
	if len(buf) != e.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", e.SizeSSZ(), len(buf))
	}
	*e = Epoch(fssz.UnmarshallUint64(buf))
	return nil
}

// MarshalSSZTo marshals epoch with the provided byte slice.
func (e *Epoch) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := e.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals epoch into a serialized object.
func (e *Epoch) MarshalSSZ() ([]byte, error) {
	marshalled := fssz.MarshalUint64([]byte{}, e.Uint64())
	return marshalled, nil
}

// SizeSSZ returns the size of the serialized object.
func (e *Epoch) SizeSSZ() int {
	return 8
}

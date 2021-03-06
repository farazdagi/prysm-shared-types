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

// AddSlot increases epoch using slot value.
func (e Epoch) AddSlot(x Slot) Epoch {
	return e + Epoch(x)
}

// AddEpoch increases epoch using another epoch value.
func (e Epoch) AddEpoch(x Epoch) Epoch {
	return Epoch(uint64(e) + uint64(x))
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

// Mod returns result of `epoch % slot`.
func (e Epoch) ModSlot(x Slot) Epoch {
	return Epoch(uint64(e) % uint64(x))
}

// HashTreeRoot returns calculated hash root.
func (e Epoch) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(e)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (e Epoch) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(uint64(e))
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
	marshalled := fssz.MarshalUint64([]byte{}, uint64(*e))
	return marshalled, nil
}

// SizeSSZ returns the size of the serialized object.
func (e *Epoch) SizeSSZ() int {
	return 8
}

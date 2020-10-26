// Package types contains types shared between various parts of the system (beacon chain, validator, slasher).
package types

import (
	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Slot)(0)
var _ fssz.HashRoot = (Epoch)(0)

// Slot represents a single slot.
type Slot uint64

// Epoch represents a single epoch.
type Epoch uint64

// ToSlot returns x converted to Slot.
func ToSlot(x uint64) Slot {
	return Slot(x)
}

// Uint64 returns slot as underlying type.
func (s Slot) Uint64() uint64 {
	return uint64(s)
}

// Mul multiplies slot by x.
func (s Slot) Mul(x uint64) Slot {
	return Slot(uint64(s) * x)
}

// Div divides slot by x.
func (s Slot) Div(x uint64) Slot {
	if x == 0 {
		panic("divbyzero")
	}
	return Slot(uint64(s) / x)
}

// Add increases slot by x.
func (s Slot) Add(x uint64) Slot {
	return Slot(uint64(s) + x)
}

// Sub subtracts x from the slot.
func (s Slot) Sub(x uint64) Slot {
	if uint64(s) < x {
		panic("underflow")
	}
	return Slot(uint64(s) - x)
}

// HashTreeRoot returns calculated hash root.
func (s Slot) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(s)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (s Slot) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(s.Uint64())
	return nil
}

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

// HashTreeRoot returns calculated hash root.
func (e Epoch) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(e)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (e Epoch) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(e.Uint64())
	return nil
}

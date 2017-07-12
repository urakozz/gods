// Copyright 2016 Home24 AG. All rights reserved.
// Proprietary license.
package bloom

import (
	"math/rand"
	"math"
)

type Hasher interface {
	Hash(s []byte) uint64
}
type hash struct {
	seed uint64
}

func NewHash() Hasher {
	return &hash{seed:uint64(rand.Float32()) * 32 + 32}
}
func (h *hash) Hash(s []byte) uint64 {
	prime := uint64(1099511628211)
	var hash uint64
	for _, r := range s {
		hash = (hash * h.seed + uint64(r)) % prime
	}
	return hash
}

type bloom struct {
	size      uint64
	functions []Hasher
	bits      *bitSet
}

func NewBloom(size uint64, fn... Hasher) *bloom {
	return &bloom{size:size, functions:fn, bits:NewBitSet(size)}
}
func (b *bloom) Add(s []byte) {
	for _, f := range b.functions {
		b.bits.Set(f.Hash(s) % b.size)
	}
}
func (b *bloom) Test(s []byte) bool {
	for _, f := range b.functions {
		if !b.bits.Test(f.Hash(s) % b.size) {
			return false
		}
	}
	return true
}

func NewOptimalBloom(maxMembers uint64, errorProbability float64) *bloom {
	size := uint64(-(float64(maxMembers) * math.Log(errorProbability)) / (math.Ln2 * math.Ln2))
	count := uint64((float64(size) / float64(maxMembers)) * math.Ln2)

	fns := make([]Hasher, count)
	for i := range fns {
		fns[i] = NewHash()
	}
	return NewBloom(size, fns...)
}




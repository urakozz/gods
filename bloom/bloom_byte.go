// Copyright 2016 Home24 AG. All rights reserved.
// Proprietary license.
package bloom

import (
	"math/rand"
	"math"
)

type Hasher interface {
	Hash(s []byte) int64
}
type hash struct {
	seed int64
}
func NewHash()*hash{
	return &hash{seed:int64(rand.Float32())*32 + 32}
}
func (h *hash) Hash(s []byte) int64 {
	prime := int64(1099511628211)
	var hash int64
	for _, r := range s {
		hash = (hash*h.seed + r) % prime
	}
	return hash
}





type bloom struct {
	size uint64
	functions []Hasher
	bits *bitSet
}
func NewBloom(size uint64, fn... Hasher)*bloom{
	return &bloom{size:size, functions:fn, bits:NewBitSet(size)}
}
func(b *bloom) Add(s []byte){
	for _, f := range b.functions{
		b.bits.Set(f.Hash(s) % b.size)
	}
}
func(b *bloom) Test(s []byte) bool {
	for _, f := range b.functions{
		if !b.bits.Test(f.Hash(s) % b.size){
			return false
		}
	}
	return true
}


func NewOptimalBloom(maxMembers uint64, errorProbability float64)*bloom{
	size := -(maxMembers * math.Log(errorProbability))/(math.Ln2 * math.Ln2)
	count := (size/maxMembers)*math.Ln2

	fns := make([]Hasher, count)
	for i, _ := range fns{
		fns[i] = NewHash()
	}
	return NewBloom(size, fns...)
}




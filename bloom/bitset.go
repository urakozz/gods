// Copyright 2016 Home24 AG. All rights reserved.
// Proprietary license.
package bloom

// the bucketSize of a bit set
const bucketSize = uint(64)

// log2BucketSize is lg(bucketSize)
const log2BucketSize = uint(6)

type bitSet struct {
	set []uint64
}
func NewBitSet(size uint64)*bitSet{
	return &bitSet{set:make([]uint64, 64)}
}
func(b *bitSet) Set(i uint64){
	b.maybeExtendTo(i)
	b.set[i>>log2BucketSize] |= (1 << (i & (bucketSize - 1)))
}
func(b *bitSet) Unset(i uint64){
	if i>>log2BucketSize < len(b.set){
		return
	}
	b.set[i>>log2BucketSize] &^= (1 << (i & (bucketSize - 1)))
}
func (b *bitSet) Test(i uint) bool {
	if i>>log2BucketSize < len(b.set){
		return
	}
	return b.set[i>>log2BucketSize] & (1 << (i & (bucketSize-1))) != 0
}
func(b *bitSet) maybeExtendTo(i uint64){
	requiredSize := i>>log2BucketSize
	if requiredSize > len(b.set){
		newset := make([]uint64, requiredSize)
		copy(newset, b.set)
		b.set = newset
	}
}


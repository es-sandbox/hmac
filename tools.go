package main

import (
	"encoding/hex"
	"strings"
	"encoding/binary"
)

const (
	incorrectParamsError = "Incorrect params"
)

type index struct {
	byteIndex uint64
	bitIndex  uint64
}

func newIndexFromN(n uint64) *index {
	return &index{
		byteIndex: n / 8,
		bitIndex:  7 - (n % 8),
	}
}

type bitArray struct {
	data []byte
}

func newEmptyBitArray(groupSizeInBytes uint64) *bitArray {
	return &bitArray{
		data: make([]byte, groupSizeInBytes),
	}
}

func newEmpty8ByteBitArray() *bitArray {
	return newEmptyBitArray(8)
}

func newEmpty200ByteBitArray() *bitArray {
	return newEmptyBitArray(200)
}

func newBitArray(data []byte) *bitArray {
	return &bitArray{
		data: data,
	}
}

func new8ByteBitArray(data []byte) *bitArray {
	if len(data) != 8 {
		panic(incorrectParamsError)
	}
	return newBitArray(data)
}

func newBitArrayFromUint64(n uint64) *bitArray {
	result := newEmpty8ByteBitArray()
	binary.BigEndian.PutUint64(result.data, n)
	return result
}

func (self *bitArray) String() string {
	return strings.ToUpper(hex.EncodeToString(self.data))
}

func (self *bitArray) lenInBytes() uint64 {
	return uint64(len(self.data))
}

func (self *bitArray) lenInBits() uint64 {
	return self.lenInBytes() * 8
}

func (self *bitArray) copy() *bitArray {
	dest := make([]byte, self.lenInBytes())
	copy(dest, self.data)
	return newBitArray(dest)
}

// reverse func *changes* internal state and returns self
func (self *bitArray) reverseBytes() *bitArray {
	lenInBytes := int(self.lenInBytes())

	for i := 0; i < lenInBytes / 2; i++ {
		j := lenInBytes - 1 - i
		self.data[i], self.data[j] = self.data[j], self.data[i]
	}

	return self
}

func (self *bitArray) getBit(n uint64) bool {
	index := newIndexFromN(n)

	if self.data[index.byteIndex]&(1<<index.bitIndex) == 0 {
		return false
	}
	return true
}

func (self *bitArray) setBit(n uint64, bit bool) {
	if self.getBit(n) == bit {
		return
	}

	index := newIndexFromN(n)

	self.data[index.byteIndex] ^= (1 << index.bitIndex)
}

// rotr func does not change internal state
func (self *bitArray) rotr(n uint64) *bitArray {
	shifted := newEmptyBitArray(self.lenInBytes())

	for i := uint64(0); i < self.lenInBits(); i++ {
		shifted.setBit((i+n)%self.lenInBits(), self.getBit(i))
	}

	return shifted
}

// rotl func does not change internal state
func (self *bitArray) rotl(n uint64) *bitArray {
	shifted := newEmptyBitArray(self.lenInBytes())
	lenInBits := self.lenInBits()

	for i := uint64(0); i < lenInBits; i++ {
		shifted.setBit((lenInBits+i-n)%lenInBits, self.getBit(i))
	}

	return shifted
}

// shr func does not change internal state
func (self *bitArray) shr(n uint64) *bitArray {
	shifted := newEmptyBitArray(self.lenInBytes())

	for i := uint64(0); i+n < self.lenInBits(); i++ {
		shifted.setBit(i+n, self.getBit(i))
	}

	return shifted
}

// xor func does not change internal state
func (self *bitArray) xor(with *bitArray) *bitArray {
	result := newEmptyBitArray(self.lenInBytes())

	for i := uint64(0); i < self.lenInBits(); i++ {
		selfBit := toUint8(self.getBit(i))
		withBit := toUint8(with.getBit(i))

		result.setBit(i, toBool(selfBit^withBit))
	}

	return result
}

func (self *bitArray) and(with *bitArray) *bitArray {
	result := newEmptyBitArray(self.lenInBytes())

	for i := uint64(0); i < self.lenInBits(); i++ {
		result.setBit(i, self.getBit(i) && with.getBit(i))
	}

	return result
}

// not func does not change internal state
func (self *bitArray) not() *bitArray {
	result := newEmptyBitArray(self.lenInBytes())

	for i := uint64(0); i < self.lenInBits(); i++ {
		result.setBit(i, !self.getBit(i))
	}
	return result
}

func (self *bitArray) sum(with *bitArray) *bitArray {
	result := newEmptyBitArray(self.lenInBytes())

	div := uint8(0)
	for i := int(self.lenInBits()) - 1; i >= 0; i-- {
		a := toUint8(self.getBit(uint64(i)))
		b := toUint8(with.getBit(uint64(i)))

		mod := (a + b + div) % 2
		div = (a + b + div) / 2

		result.setBit(uint64(i), toBool(mod))
	}

	return result
}

func toUint8(bit bool) uint8 {
	if bit {
		return 1
	}
	return 0
}

func toBool(bit uint8) bool {
	switch bit {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(incorrectParamsError)
	}
}

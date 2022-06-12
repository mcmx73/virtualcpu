package virtualcpu

import "unsafe"

type ByteRegister interface {
	Set(byte)
	Get() byte
}

type Word [2]byte

func (w *Word) int16() int16 {
	return wordToInt(*w)
}

type WordRegister interface {
	Set(Word)
	Get() Word
	GetLow() byte
	GetHigh() byte
	SetLow(byte)
	SetHigh(byte)
}

type Flag interface {
	Set(bool)
	Get() bool
}

type byteRegister struct {
	value byte
}

func (b byteRegister) Set(data byte) {
	b.value = data
}

func (b byteRegister) Get() byte {
	return b.value
}

func intToWord(num int16) [2]byte {
	var arr [2]byte
	arr[0] = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(0)))
	arr[1] = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(1)))
	return arr
}

func wordToInt(word [2]byte) int16 {
	val := int16(0)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(0))) = word[0]
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(1))) = word[1]
	return val
}

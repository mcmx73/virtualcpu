package virtualcpu

type ByteRegister interface {
	Set(byte)
	Get() byte
}

type Word [2]byte

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

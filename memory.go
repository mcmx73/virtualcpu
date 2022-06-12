package virtualcpu

type MemoryInterface interface {
	Set(address int16, data byte)
	Get(address int16) (data byte)
}

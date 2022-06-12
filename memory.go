package virtualcpu

type MemoryInterface interface {
	Set(address Word, data byte)
	Get(address Word) (data byte)
}

type MemoryDefault struct {
	memoryData [0xFFFF]byte
}

func (m *MemoryDefault) Set(address Word, data byte) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDefault) Get(address Word) (data byte) {
	//TODO implement me
	panic("implement me")
}

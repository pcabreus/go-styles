package constantnames

const MaxPacketSize = 512

const (
	ExecuteBit = 1 << iota
	WriteBit
	ReadBit
)

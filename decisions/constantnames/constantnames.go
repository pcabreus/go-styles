package constantnames

//Constant names must use MixedCaps like all other names in Go. (Exported constants start with uppercase, while
//unexported constants start with lowercase.) This applies even when it breaks conventions in other languages.
//Constant names should not be a derivative of their values and should instead explain what the value denotes.

const MaxPacketSize = 512
const privateData = true

const (
	ExecuteBit = 1 << iota
	WriteBit
	ReadBit
)

// Bad:
//const MAX_PACKET_SIZE = 512
//const kMaxBufferSize = 1024
//const KMaxUsersPergroup = 500

//Name constants based on their role, not their values. If a constant does not have a role apart from its value,
//then it is unnecessary to define it as a constant.

const (
	NumbersOfMonths = 12
	UserNameColumn  = "username"
	GroupColumn     = "group"
)

// Bad:
//const Twelve = 12

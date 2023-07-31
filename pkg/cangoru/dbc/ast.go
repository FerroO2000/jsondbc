package dbc

type DBC struct {
	Version             string
	NewSymbols          *NewSymbols
	BitTiming           *BitTiming
	Nodes               *Nodes
	ValueTables         []*ValueTable
	Messages            []*Message
	MessageTransmitters []*MessageTransmitter
	EnvVars             []*EnvVar
	EnvVarDatas         []*EnvVarData
	SignalTypes         []*SignalType
	Comments            []*Comment
	Attributes          []*Attribute
	AttributeDefaults   []*AttributeDefault
	AttributeValues     []*AttributeValue
	ValueEncodings      []*ValueEncoding
	SignalTypeRefs      []*SignalTypeRef
	SignalGroups        []*SignalGroup
	SignalExtValueTypes []*SignalExtValueType
	ExtendedMuxes       []*ExtendedMux
}

type NewSymbols struct {
	Symbols []string
}

type BitTiming struct {
	Baudrate      uint32
	BitTimingReg1 uint32
	BitTimingReg2 uint32
}

type Nodes struct {
	Names []string
}

type ValueTable struct {
	Name   string
	Values []*ValueDescription
}

type ValueDescription struct {
	ID   uint32
	Name string
}

type ValueEncodingKind uint

const (
	ValueEncodingSignal ValueEncodingKind = iota
	ValueEncodingEnvVar
)

type ValueEncoding struct {
	Kind       ValueEncodingKind
	MessageID  uint32
	SignalName string
	EnvVarName string
	Values     []*ValueDescription
}

type Message struct {
	ID          uint32
	Name        string
	Size        uint32
	Transmitter string
	Signals     []*Signal
}

type SignalByteOrder uint

const (
	SignalLittleEndian SignalByteOrder = iota
	SignalBigEndian
)

type SignalValueType uint

const (
	SignalUnsigned SignalValueType = iota
	SignalSigned
)

type Signal struct {
	Name           string
	IsMultiplexor  bool
	IsMultiplexed  bool
	MuxSwitchValue uint32
	Size           uint32
	StartBit       uint32
	ByteOrder      SignalByteOrder
	ValueType      SignalValueType
	Factor         float64
	Offset         float64
	Min            float64
	Max            float64
	Unit           string
	Receivers      []string
}

type SignalExtValueTypeType uint

const (
	SignalExtValueTypeInteger SignalExtValueTypeType = iota
	SignalExtValueTypeFloat
	SignalExtValueTypeDouble
)

type SignalExtValueType struct {
	MessageID    uint32
	SignalName   string
	ExtValueType SignalExtValueTypeType
}

type MessageTransmitter struct {
	MessageID    uint32
	Transmitters []string
}

type EnvVarType uint

const (
	EnvVarInt EnvVarType = iota
	EnvVarFloat
	EnvVarString
)

type EnvVarAccessType uint

const (
	EnvVarDummyNodeVector0 EnvVarAccessType = iota
	EnvVarDummyNodeVector1
	EnvVarDummyNodeVector2
	EnvVarDummyNodeVector3
	EnvVarDummyNodeVector8000
	EnvVarDummyNodeVector8001
	EnvVarDummyNodeVector8002
	EnvVarDummyNodeVector8003
)

type EnvVar struct {
	Name         string
	Type         EnvVarType
	Min          float64
	Max          float64
	Unit         string
	InitialValue float64
	ID           uint32
	AccessType   EnvVarAccessType
	AccessNodes  []string
}

type EnvVarData struct {
	EnvVarName string
	DataSize   uint32
}

type SignalType struct {
	TypeName       string
	Size           uint32
	ByteOrder      SignalByteOrder
	ValueType      SignalValueType
	Factor         float64
	Offset         float64
	Min            float64
	Max            float64
	Unit           string
	DefaultValue   float64
	ValueTableName string
}

type SignalTypeRef struct {
	TypeName   string
	MessageID  uint32
	SignalName string
}

type SignalGroup struct {
	MessageID   uint32
	GroupName   string
	Repetitions uint32
	SignalNames []string
}

type CommentKind uint

const (
	CommentGeneral CommentKind = iota
	CommentNode
	CommentMessage
	CommentSignal
	CommentEnvVar
)

type Comment struct {
	Kind       CommentKind
	Text       string
	NodeName   string
	MessageID  uint32
	SignalName string
	EnvVarName string
}

type AttributeKind uint

const (
	AttributeGeneral AttributeKind = iota
	AttributeNode
	AttributeMessage
	AttributeSignal
	AttributeEnvVar
)

type AttributeType uint

const (
	AttributeInt AttributeType = iota
	AttributeFloat
	AttributeString
	AttributeEnum
	AttributeHex
)

type Attribute struct {
	Kind       AttributeKind
	Type       AttributeType
	Name       string
	MinInt     int
	MaxInt     int
	MinHex     int
	MaxHex     int
	MinFloat   float64
	MaxFloat   float64
	EnumValues []string
}

type AttributeDefault struct {
	AttributeName string
	ValueLiteral  string
}

type AttributeValue struct {
	AttributeKind AttributeKind
	AttributeName string
	NodeName      string
	MessageID     uint32
	SignalName    string
	EnvVarName    string
	ValueLiteral  string
}

type ExtendedMuxRange struct {
	From uint32
	To   uint32
}

type ExtendedMux struct {
	MessageID       uint32
	MultiplexorName string
	MultiplexedName string
	Ranges          []*ExtendedMuxRange
}

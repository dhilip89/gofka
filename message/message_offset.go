package message

const (
	msgOffsetSize = 8
	msgSizeSize   = 4
	// MsgOverhead defines the byte overhead for each Message within a MessageSet.
	MsgOverhead = msgSizeSize + msgOffsetSize
)

// MessageOffset is an utility type wrapping a Message and its logical and
// physical offsets within a MessageSet.
type MessageOffset struct {
	// Offset is the Message logical offset within a MessageSet.
	Offset uint64
	// Pos is the Message physical byte offset within a MessageSet.
	Pos uint32
	// MsgSize is the Message byte size.
	MsgSize uint32
	// Message is the actual Message payload.
	Message
}

// Size returns this Message's size within a MessageSet.
func (m *MessageOffset) Size() uint32 {
	return MsgOverhead + m.MsgSize
}

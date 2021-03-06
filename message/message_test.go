package message_test

import (
	"testing"

	. "github.com/tsenart/gofka/message"
	. "github.com/tsenart/gofka/testing"
)

func TestMessage_ChecksumHashValid(t *testing.T) {
	m := NewMessage([]byte("key"), []byte("value"))
	Equals(t, m.Checksum(), m.Hash())
	Assert(t, m.Valid(), "Valid: want: true, got: false")
	m.SetValue([]byte("foobar"))
	Assert(t, !m.Valid(), "Valid: want: false, got: true")
}

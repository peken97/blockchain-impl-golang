package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.Addr()], trb)
	assert.Equal(t, trb.peers[tra.Addr()], tra)

}

func TestSendMessage(t *testing.T) {
	
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	tra.SendMessage(trb.Addr(), []byte("hello"))

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.From, tra.Addr())
	assert.Equal(t, rpc.Payload, []byte("hello"))
}
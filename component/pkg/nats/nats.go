package nats

import (
	"github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/gen/wasmcloud/messaging/handler"
	"github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/gen/wasmcloud/messaging/types"
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

type (
	Conn struct {
		js JetStreamContext
	}
	JetStreamContext struct{}
	Msg              struct {
		Subject string
		Reply   string
		Data    []byte
		Header  map[string][]string
	}
)

type MsgHandler func(msg *Msg)

func NewConn() *Conn {
	return &Conn{}
}

func FromBrokerMessageToNatsMessage(bm types.BrokerMessage) *Msg {
	if bm.ReplyTo.None() {
		return &Msg{
			Data:    bm.Body.Slice(),
			Subject: bm.Subject,
			Reply:   "",
		}
	} else {
		return &Msg{
			Data:    bm.Body.Slice(),
			Subject: bm.Subject,
			Reply:   *bm.ReplyTo.Some(),
		}
	}
}

func ToBrokenMessageFromNatsMessage(nm *Msg) types.BrokerMessage {
	if nm.Reply == "" {
		return types.BrokerMessage{
			Subject: nm.Subject,
			Body:    cm.ToList(nm.Data),
			ReplyTo: cm.None[string](),
		}
	} else {
		return types.BrokerMessage{
			Subject: nm.Subject,
			Body:    cm.ToList(nm.Data),
			ReplyTo: cm.Some(nm.Subject),
		}
	}
}

func (conn *Conn) RegisterSubscription(fn func(*Msg)) {
	handler.Exports.HandleMessage = func(msg types.BrokerMessage) (result cm.Result[string, struct{}, string]) {
		natsMsg := FromBrokerMessageToNatsMessage(msg)
		fn(natsMsg)
		return cm.OK[cm.Result[string, struct{}, string]](struct{}{})
	}
}

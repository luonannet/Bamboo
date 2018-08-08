package hand

import (
	"log"

	tp "github.com/henrylee2cn/teleport"
)

//Client Client
type Client struct {
	tp.PushCtx
}

//Receive Receive
func (c *Client) Receive(arg *string) *tp.Rerror {
	log.Println("ReceiveItemData:", *arg)
	return nil
}

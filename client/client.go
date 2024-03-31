package client

import (
	"github.com/leap-fish/necs/transports"
)

type Client struct {
	Transport *transports.WsClientTransport
}

package ModPC

import (
	RaknetConnection "Eulogist/core/raknet"

	"github.com/sandertv/go-raknet"
)

type Server struct {
	listener  *raknet.Listener
	connected chan struct{}

	IP   string
	Port int
	*RaknetConnection.RaknetConnection
}

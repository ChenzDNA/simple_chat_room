package utils

import (
	"bufio"
	"net"
	"sync"
)

type Client struct {
	ClientId string
	Writer   *bufio.ReadWriter
	Conn     *net.Conn
}

var ClientMap sync.Map

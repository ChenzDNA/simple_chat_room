package utils

import (
	"bufio"
	"net"
)

type Client struct {
	ClientId string
	Writer   *bufio.ReadWriter
	Conn     *net.Conn
}

var ClientMap map[string]*Client

func init() {
	ClientMap = make(map[string]*Client)
}

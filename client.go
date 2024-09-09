package modbusGo

import (
	"fmt"
	"net"
	"time"
)

type TCPHandler struct {
	Address string
	Timeout time.Duration
	conn    net.Conn
}

func NewTCPHandler(address string,
	timeout time.Duration) *TCPHandler {
	return &TCPHandler{Address: address,
		Timeout: timeout}
}

func (h *TCPHandler) Connect() error {
	conn, err := net.DialTimeout("tcp", h.Address, h.Timeout)
	if err != nil {
		return fmt.Errorf("filed to connect: %w", err)
	}
	h.conn = conn
	return nil
}

func (h *TCPHandler) Close() error {
	if h.conn != nil {
		return h.conn.Close()
	}
	return nil
}

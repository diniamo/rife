package proto

import (
	"net"
)

type Conn struct {
	conn net.Conn
	rbuf [1024]byte
	wbuf [1024]byte
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{conn: conn}
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) ReadRaw() ([]byte, error) {
	n, err := c.conn.Read(c.rbuf[:])
	if err != nil {
		return nil, err
	}

	return c.rbuf[:n], nil
}

func (c *Conn) ReadPacket() (Packet, error) {
	raw, err := c.ReadRaw()
	if err != nil {
		return Packet{}, err
	}

	return DecodePacket(raw), nil
}

func (c *Conn) WriteRaw(packet []byte) error {
	_, err := c.conn.Write(packet)
	return err
}

func (c *Conn) WritePacket(packet *Packet) error {
	return c.WriteRaw(EncodePacket(packet, c.wbuf[:]))
}

package models

import "net"

type Payload struct {
	Ip        net.IP
	Countries []string
}

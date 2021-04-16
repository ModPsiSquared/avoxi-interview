package models

import "net"

type IpStatus struct {
	Ip       net.IP
	Approved bool
}

func NewIpStatus(ip net.IP, approved bool) IpStatus {
	return IpStatus{
		Ip:       ip,
		Approved: approved,
	}
}

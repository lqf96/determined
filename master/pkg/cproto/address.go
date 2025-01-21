package cproto

import (
	"fmt"

	"github.com/determined-ai/determined/proto/pkg/taskv1"
)

// AddrPort is a pair of IP address and port number.
type AddrPort struct {
 	IP string `json:"ip"`
	Port int `json:"port"`
}

// Address represents an exposed port on a container.
type Address struct {
	// ContainerAddrPort is the IP address and port pair from inside the container.
	ContainerAddrPort AddrPort `json:"container_addr"`
	// HostAddrPort is the IP address and port pair from outside the container,
	// when the port is forwarded to the host machine.
	HostAddrPort *AddrPort `json:"host_addr,omitempty"`
}

func (a AddrPort) String() string {
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}
	 
func (a Address) String() string {
	addrStr := a.ContainerAddrPort.String()
	if a.HostAddrPort != nil {
		addrStr = fmt.Sprintf("%s:%s", a.HostAddrPort.String(), addrStr)
	}
	return addrStr
}

func (a Address) TargetAddrPort() AddrPort {
	if a.HostAddrPort != nil {
		return *a.HostAddrPort
	}
	return a.ContainerAddrPort
}

// Proto returns the proto representation of address.
func (a *Address) Proto() *taskv1.Address {
	if a == nil {
		return nil
	}
	return &taskv1.Address{
		ContainerIp:   a.ContainerAddrPort.IP,
		ContainerPort: int32(a.ContainerAddrPort.Port),
		HostIp:        a.TargetAddrPort().IP,
		HostPort:      int32(a.TargetAddrPort().Port),
	}
}

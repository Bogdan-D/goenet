package goenet

/*
#cgo CFLAGS: -I${SRCDIR}/enet/include/
#cgo LDFLAGS: -lenet

#include "enet/enet.h"
*/
import "C"

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type ENetAddress C.ENetAddress

func (a *ENetAddress) GetHost() string {
	var addr [4]byte
	var str [4]string
	binary.LittleEndian.PutUint32(addr[:], uint32(a.host))
	for i, n := range addr {
		str[i] = strconv.Itoa(int(n))
	}

	return strings.Join(str[:], ".")
}

func (a *ENetAddress) GetPort() uint16 {
	return uint16(a.port)
}

func (a *ENetAddress) String() string {
	return fmt.Sprintf("%s:%s", a.GetHost(), a.GetPort())
}

func (a *ENetAddress) SetHost(h uint) {
	a.host = C.enet_uint32(h)
}

func (a *ENetAddress) SetPort(port uint) {
	a.port = C.enet_uint16(port)
}

func (a *ENetAddress) SetHostName(hostName string) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_set_host((*C.ENetAddress)(a), hName))
}

func (a *ENetAddress) HostIp(hostName string, nameLength int) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_get_host_ip((*C.ENetAddress)(a), hName, C.size_t(nameLength)))
}

func (a *ENetAddress) Host(hostName string, nameLength int) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_get_host((*C.ENetAddress)(a), hName, C.size_t(nameLength)))
}

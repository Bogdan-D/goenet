package goenet

/*

#cgo windows,686 LDFLAGS: -L${SRCDIR}/lib -lenet_windows_686 -lws2_32 -lwinmm
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib -lenet_windows_amd64 -lws2_32 -lwinmm

#cgo linux,686 LDFLAGS: -L${SRCDIR}/lib -lenet_linux_686
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/lib -lenet_linux_amd64
*/
import "C"

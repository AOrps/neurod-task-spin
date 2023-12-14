package lib

import (
	"strconv"
	"net"
	"fmt"
)

type Integer interface {
	int | int8 | int16 | int32 | int64
}

func CheckPort[T Integer](n T) bool {

	retBool := true

	port := strconv.Itoa(int(n))
	
	ln, err := net.Listen("tcp",":"+port)

	if err != nil {
		fmt.Printf("%s not available\n", port)
		retBool = false
	}

	if ln.Close() != nil {
 		fmt.Printf("couldn't stop listening on port %s\n", port)
	}

	return retBool
}

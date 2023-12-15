package lib

/*
   This file is responsible for adding the functionality for checking if a
   certain port is able to listen to a tcp connection (and then subsequently
   close that connection).


   v1.0  @AOrps
*/

import (
	"strconv"
	"net"
	"fmt"
)

// make a type called Integer that uses Generics to
// make all integers try to function similiarly
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// CheckPort: checks if the program is enable to
//            listen on the <n>th port 
// ex. CheckPort(7100) -> false
// ex. CheckPort(7101) -> true 
func CheckPort[T Integer](n T) bool {

	port := strconv.Itoa(int(n))
	
	// this is where the connection attempt happens
	ln, err := net.Listen("tcp",":"+port)

	// should it result in an error, then it should be
	// understood that the <n>th port is already in use
	if err != nil {
		fmt.Printf("%s not available\n", port)
		return false		
	}

	if ln.Close() != nil {
 		fmt.Printf("couldn't stop listening on port %s\n", port)
	}

	// should the program get to this stage, there are
	// no errors and the program was able to listen on
	// the <n>th port as well as close that listener
	return true
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		b := make([]byte, 102400)
		_, err = os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		con.Write(b)
	}
}

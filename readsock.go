package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create("some.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := con.Accept()
		if err != nil {
			fmt.Println(err)
		}
		for {
			b := make([]byte, 102400)
			_, err = c.Read(b)
			if err != nil {
				break
			}
			for x := 0; x < 102400; x++ {
				b[x] = 'a'
			}
			f.Write(b)
		}
	}
}

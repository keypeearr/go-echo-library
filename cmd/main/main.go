package main

import "github.com/kylerequez/go-echo-library/src/server"

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}

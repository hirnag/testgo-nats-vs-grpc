package main

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Server is starting...")
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		fmt.Println(err)
	}

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		//fmt.Println(string(m.Data[:]))
		nc.Publish(m.Reply, m.Data)
	})
	nc.Flush()
	runtime.Goexit()
}


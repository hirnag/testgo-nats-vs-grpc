package main

import (
	prt "../proto"
	"github.com/nats-io/go-nats"
	"time"
	"fmt"
	"google.golang.org/grpc"
	"context"
)
func main() {
	loop := 1000
	var a int64
	var b int64
	var c int64
	var d int64
	for i := 0; i < loop; i++ {
		n1,n2 := RequestNats()
		g1,g2 := RequestGrpc()
		a += n1
		b += n2
		c += g1
		d += g2
	}
	fmt.Println(a/int64(loop))
	fmt.Println(b/int64(loop))
	fmt.Println("-------")
	fmt.Println(c/int64(loop))
	fmt.Println(d/int64(loop))
}

func RequestNats() (int64, int64) {
	start := time.Now().UnixNano()
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	ir := calcTime(start)

	// Requests
	_, err = nc.Request("help", []byte("nats request"), 5000*time.Millisecond)
	if err != nil {
		panic(err)
	}
	res := calcTime(start)
	return ir, res - ir
}
func RequestGrpc() (int64, int64) {
	start := time.Now().UnixNano()
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ir := calcTime(start)

	cli := prt.NewGreeterClient(conn)
	_, err = cli.Reply(context.Background(), &prt.ReqMessage{Name: "grpc request"})
	if err != nil {
		panic(err)
	}
	res := calcTime(start)
	return ir, res - ir
}
func calcTime(s int64) int64 {
	now := time.Now().UnixNano()
	return now - s
}


package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"

	"github.com/calvernaz/go-sandbox/grpc-chunker/protos/chunker"
)

func main() {
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cc := chunker.NewChunkerClient(conn)
	client, err := cc.Chunker(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	start := time.Now()
	blob := make([]byte, 128*1024*1024)
	for {
		c, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(blob))
				elapsed := time.Since(start)
				log.Printf("Took %s", elapsed)
				return
			}

			panic(err)
		}

		blob = append(blob, c.Chunk...)
	}
}

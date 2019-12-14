package main

import (
	"context"
	"flag"
	_ "fmt"
	"github.com/aaronland/go-aws-s3"
	"log"
)

func main() {

	dsn := flag.String("dsn", "", "...")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := s3.NewS3ConnectionWithDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	keys := flag.Args()

	err = conn.DeleteKeysIfExists(ctx, keys...)

	if err != nil {
		log.Fatal(err)
	}

}

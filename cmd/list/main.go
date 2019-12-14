package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-aws-s3"
	"log"
)

func main() {

	dsn := flag.String("dsn", "", "...")
	timings := flag.Bool("timings", false, "")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := s3.NewS3ConnectionWithDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	list_cb := func(obj *s3.S3Object) error {
		log.Printf("%s\t%v\n", obj.Key, obj.LastModified)
		return nil
	}

	list_opts := s3.DefaultS3ListOptions()
	list_opts.Timings = *timings

	err = conn.List(ctx, list_cb, list_opts)

	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/datasize"
	"github.com/aaronland/go-aws-s3"
	"log"
)

func main() {

	dsn := flag.String("dsn", "", "...")
	timings := flag.Bool("timings", false, "")
	gt := flag.String("gt", "", "")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := s3.NewS3ConnectionWithDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	var sz datasize.ByteSize

	if *gt != "" {

		err := sz.UnmarshalText([]byte(*gt))

		if err != nil {
			log.Fatal(err)
		}
	}

	max_size := sz.Bytes()

	list_cb := func(obj *s3.S3Object) error {

		if max_size > 0 && uint64(obj.Size) <= max_size {
			return nil
		}

		fmt.Println(obj)
		return nil
	}

	list_opts := s3.DefaultS3ListOptions()
	list_opts.Timings = *timings

	err = conn.List(ctx, list_cb, list_opts)

	if err != nil {
		log.Fatal(err)
	}
}

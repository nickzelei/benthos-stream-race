package main

import (
	"context"

	_ "github.com/warpstreamlabs/bento/public/components/io"
	"github.com/warpstreamlabs/bento/public/service"

	"golang.org/x/sync/errgroup"
)

const (
	syncConfig = `
input:
  generate:
    count: 10
    interval: ""
    mapping: |
      root = {
        "id": uuid_v4(),
        "created_at": now(),
        "updated_at": now()
      }
pipeline:
  threads: -1
  processors: []
output:
  label: ""
  stdout:
    codec: lines
`
)

func main() {
	errgrp := errgroup.Group{}
	errgrp.Go(func() error {
		return sync(syncConfig)
	})
	errgrp.Go(func() error {
		return sync(syncConfig)
	})
	err := errgrp.Wait()
	if err != nil {
		panic(err)
	}
}

func sync(config string) error {
	env := service.NewEnvironment()
	streambldr := env.NewStreamBuilder()
	err := streambldr.SetYAML(config)
	if err != nil {
		return err
	}

	stream, err := streambldr.Build()
	if err != nil {
		return err
	}
	return stream.Run(context.Background())
}

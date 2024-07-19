# Benthos Stream Race

To reproduce:

```console
go run -race main.go
```

Should see something like this:

```console
WARNING: DATA RACE
Read at 0x00c0000f7808 by goroutine 17:
  github.com/warpstreamlabs/bento/internal/docs.FieldSpecs.SetDefault()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/internal/docs/field.go:655 +0xb8
  github.com/warpstreamlabs/bento/internal/docs.FieldSpecs.SetDefault()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/internal/docs/field.go:660 +0x15c
  github.com/warpstreamlabs/bento/public/service.NewStreamBuilder()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/public/service/stream_builder.go:80 +0x164
  github.com/warpstreamlabs/bento/public/service.(*Environment).NewStreamBuilder()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/public/service/environment.go:106 +0x4c
  main.sync()
      /Users/nick/code/playground/benthos-stream-race/main.go:51 +0x48
  main.main.func1()
      /Users/nick/code/playground/benthos-stream-race/main.go:38 +0x30
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/nick/go/pkg/mod/golang.org/x/sync@v0.6.0/errgroup/errgroup.go:78 +0x7c

Previous write at 0x00c0000f7808 by goroutine 18:
  github.com/warpstreamlabs/bento/internal/docs.FieldSpecs.SetDefault()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/internal/docs/field.go:662 +0x214
  github.com/warpstreamlabs/bento/internal/docs.FieldSpecs.SetDefault()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/internal/docs/field.go:660 +0x15c
  github.com/warpstreamlabs/bento/public/service.NewStreamBuilder()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/public/service/stream_builder.go:80 +0x164
  github.com/warpstreamlabs/bento/public/service.(*Environment).NewStreamBuilder()
      /Users/nick/go/pkg/mod/github.com/warpstreamlabs/bento@v1.1.0/public/service/environment.go:106 +0x4c
  main.sync()
      /Users/nick/code/playground/benthos-stream-race/main.go:51 +0x48
  main.main.func2()
      /Users/nick/code/playground/benthos-stream-race/main.go:41 +0x30
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/nick/go/pkg/mod/golang.org/x/sync@v0.6.0/errgroup/errgroup.go:78 +0x7c

Goroutine 17 (running) created at:
  golang.org/x/sync/errgroup.(*Group).Go()
      /Users/nick/go/pkg/mod/golang.org/x/sync@v0.6.0/errgroup/errgroup.go:75 +0x10c
  main.main()
      /Users/nick/code/playground/benthos-stream-race/main.go:37 +0x70

Goroutine 18 (running) created at:
  golang.org/x/sync/errgroup.(*Group).Go()
      /Users/nick/go/pkg/mod/golang.org/x/sync@v0.6.0/errgroup/errgroup.go:75 +0x10c
  main.main()
      /Users/nick/code/playground/benthos-stream-race/main.go:40 +0x80
```

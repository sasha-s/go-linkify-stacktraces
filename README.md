# go-linkify-stacktraces
Golang: linkify http://localhost:1234/debug/pprof/goroutine?debug stacktaces, linking file:lines to github.

Works best when everything is vendored.

## What to expect?
Get stacktraces like

goroutine profile: total 3

1 @ 0x2b3d38 0x2b3b13 0x2af444 0x1f9e3e 0x1fa050 0xad58a 0xdb9a3 0xad58a 0xaedfd 0xaf87e 0xac2ee 0x5e791

\#	0x2b3d38	runtime/pprof.writeRuntimeProfile+0xb8						/usr/local/Cellar/go/1.6/libexec/src/runtime/pprof/pprof.go:545

\#	0x2b3b13	runtime/pprof.writeGoroutine+0x93						/usr/local/Cellar/go/1.6/libexec/src/runtime/pprof/pprof.go:507

\#	0x2af444	runtime/pprof.(*Profile).WriteTo+0xd4						/usr/local/Cellar/go/1.6/libexec/src/runtime/pprof/pprof.go:236

\#	0x1f9e3e	net/http/pprof.handler.ServeHTTP+0x37e						/usr/local/Cellar/go/1.6/libexec/src/net/http/pprof/pprof.go:210

\#	0x1fa050	net/http/pprof.Index+0x200							/usr/local/Cellar/go/1.6/libexec/src/net/http/pprof/pprof.go:222

\#	0xad58a		net/http.HandlerFunc.ServeHTTP+0x3a						/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1618

\#	0xdb9a3		github.com/sasha-s/go-linkify-stacktraces.LinkifyingMiddleware.func1.1+0x373	/Users/sasha/go/src[/github.com/sasha-s/go-linkify-stacktraces/linkify.go:22](https://github.com/sasha-s/go-linkify-stacktraces/blob/b11a8bf9b8b57397d617e71fd7cd215cd7a2ce75/linkify.go\#L22)

\#	0xad58a		net/http.HandlerFunc.ServeHTTP+0x3a						/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1618

\#	0xaedfd		net/http.(*ServeMux).ServeHTTP+0x17d						/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1910

\#	0xaf87e		net/http.serverHandler.ServeHTTP+0x19e						/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:2081

\#	0xac2ee		net/http.(*conn).serve+0xf2e							/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1472

...

## Usage
See [example](https://github.com/sasha-s/go-linkify-stacktraces/blob/b11a8bf9b8b57397d617e71fd7cd215cd7a2ce75/cmd/linkify-example/lnk.go).

```go
http.DefaultServeMux.Handle("/debug/pprof/goroutine", linkify.PprofHandler(repo, commitHash))
log.Fatal(s.ListenAndServe())
```

`commitHash` is git commit hash. One can get it with
```sh
git rev-parse HEAD
```

`repo` is your repository, something like `sasha-s/go-linkify-stacktraces`.

## Linking git commit hash in.
See [link command](https://golang.org/cmd/link/), -X part.

build:
```sh
go build -ldflags="-X main.commitHash=`git rev-parse HEAD`" github.com/sasha-s/go-linkify-stacktraces/cmd/linkify-example
```

run:
```sh
go run -ldflags="-X main.commitHash=`git rev-parse HEAD`" cmd/linkify-example/lnk.go
```

this assumes you have
```go
var commitHash string
```
in lnk.go


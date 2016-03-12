# go-linkify-stacktraces
Golang: linkify /debug/pprof/goroutine?debug stacktaces, linking file:lines to github.

Works best when everything is vendored.

## What to expect?
Get stacktraces like

goroutine profile: total 3
...

1 @ 0x2f283 0x29c5e 0x29120 0x18234a 0x1823b6 0x18604c 0x1a202d 0xb0941 0xafbc9 0xafa16 0x23c7 0x2eea0 0x5e791

\#	0x29120		net.runtime_pollWait+0x60			/usr/local/Cellar/go/1.6/libexec/src/runtime/netpoll.go:160

\#	0x18234a	net.(*pollDesc).Wait+0x3a			/usr/local/Cellar/go/1.6/libexec/src/net/fd_poll_runtime.go:73

\#	0x1823b6	net.(*pollDesc).WaitRead+0x36			/usr/local/Cellar/go/1.6/libexec/src/net/fd_poll_runtime.go:78

\#	0x18604c	net.(*netFD).accept+0x27c			/usr/local/Cellar/go/1.6/libexec/src/net/fd_unix.go:426

\#	0x1a202d	net.(*TCPListener).AcceptTCP+0x4d		/usr/local/Cellar/go/1.6/libexec/src/net/tcpsock_posix.go:254

\#	0xb0941		net/http.tcpKeepAliveListener.Accept+0x41	/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:2427

\#	0xafbc9		net/http.(*Server).Serve+0x129			/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:2117

\#	0xafa16		net/http.(*Server).ListenAndServe+0x136		/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:2098

\#	0x23c7		main.main+0x2b7					/Users/sasha/go/src/[github.com/sasha-s/go-linkify-stacktraces/cmd/linkify-example/lnk.go:33](https://github.com/sasha-s/go-linkify-stacktraces/blob/8b4538f4b0654da8cae843508cd5b2491bc53d5e/cmd/linkify-example/lnk.go#L33)

\#	0x2eea0		runtime.main+0x2b0				/usr/local/Cellar/go/1.6/libexec/src/runtime/proc.go:188

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


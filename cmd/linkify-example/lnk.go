package main

// Build it with
// go build -ldflags="-X main.commitHash=`git rev-parse HEAD`" github.com/sasha-s/go-linkify-stacktraces/cmd/linkify-example
// or run
// go run -ldflags="-X main.commitHash=`git rev-parse HEAD`" cmd/linkify-example/lnk.go
// open http://localhost:31415/debug/pprof/goroutine?debug=1

import (
	"flag"
	"log"
	"net/http"

	"github.com/sasha-s/go-linkify-stacktraces"
)

var s http.Server
var repo string
var commitHash = "master" // This way links point to master if commit hash is not specified.

func init() {
	flag.StringVar(&s.Addr, "addr", ":31415", "TCP address to listen on")
	flag.StringVar(&repo, "repo", "sasha-s/go-linkify-stacktraces", "github repo")
}

func main() {
	flag.Parse()
	log.Println("open http://localhost:31415/debug/pprof/goroutine?debug=1")
	if commitHash == "master" {
		log.Println("maybe set the commitHash variable at link time. See https://github.com/sasha-s/go-linkify-stacktraces/blob/master/linkify.go")
	}
	http.DefaultServeMux.Handle("/debug/pprof/goroutine", linkify.PprofHandler(repo, commitHash))
	log.Fatal(s.ListenAndServe())
}

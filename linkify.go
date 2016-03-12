package linkify

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"regexp"
	"strings"
)

// LinkifyingMiddleware returns a middleware func that linkifyes /debug/pprof/goroutine?debug,
// pointing all lines from github repo `repo` to github, using commit hash `commitHash`.
// One can use branch name instead of commit hash.
func LinkifyingMiddleware(repo, commitHash string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(r.URL.Query()["raw"]) != 0 || r.URL.Query().Get("debug") == "" {
				h.ServeHTTP(w, r)
				return
			}
			fmt.Fprint(w, "<body><pre>")
			h.ServeHTTP(linkifyingWriter{w, repo, commitHash}, r)
			fmt.Fprint(w, "</pre></body>")
		})
	}
}

// PprofHandler returns an http handler that can handle /debug/pprof/goroutine?debug, linkifying
// all lines from repo.
// Sample usage (assuming the repo is https://github.com/sasha-s/go-linkify-stacktraces/)
// s := http.Server{Addr: ":31415"}
// http.DefaultServeMux.Handle("/debug/pprof/goroutine", PprofHandler("sasha-s/go-linkify-stacktraces", "master"))
// log.Fatal(s.ListenAndServe())
func PprofHandler(repo, commitHash string) http.Handler {
	return LinkifyingMiddleware(repo, commitHash)(http.HandlerFunc(pprof.Index))
}

type linkifyingWriter struct {
	http.ResponseWriter
	Repo       string
	CommitHash string
}

func (w linkifyingWriter) Write(b []byte) (int, error) {
	lines := strings.Split(string(b), "\n")
	pref := "/github.com/" + w.Repo + "/"
	for i, l := range lines {
		parts := strings.Split(l, pref)
		if len(parts) == 1 {
			continue
		}
		for n, p := range parts {
			if n == 0 {
				continue
			}
			matches := r.FindStringSubmatch(p)
			lm := len(matches)
			if lm < 4 {
				parts[n] = pref + p
				continue
			}
			f, ln, rest := matches[1], matches[lm-2], matches[lm-1]
			link := fmt.Sprintf(`https://github.com/%s/blob/%s/%s#L%s`, w.Repo, w.CommitHash, f, ln)
			parts[n] = fmt.Sprintf(`<a href="%s">/github.com/%s/%s:%s</a>%s`, link, w.Repo, f, ln, rest)
		}
		lines[i] = strings.Join(parts, "")
	}
	txt := strings.Join(lines, "\n")
	// Hacky: number of bytes written could be greater than input size.
	// Works fine with /debug/pprof/goroutine.
	return w.ResponseWriter.Write([]byte(txt))
}

var r = regexp.MustCompile(`^((\w|[ -.])+((\\|/)(\w|[ -.])+)*\.go):(\d+)(.*)$`)

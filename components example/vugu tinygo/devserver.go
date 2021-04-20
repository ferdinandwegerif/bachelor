// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/vugu/vugu/devutil"
)

func main() {
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)

	wc := devutil.MustNewTinygoCompiler().SetDir(".")
	defer wc.Close()
	wc.AddGoGet("go get -u github.com/vugu/vugu github.com/vugu/vjson")
	wc.NoDocker()
	wc.AddGoGet("go get -u -x github.com/vugu/vjson github.com/vugu/html github.com/vugu/xxhash")

	mux := devutil.NewMux()
	rsc := devutil.DefaultAutoReloadIndex.Replace(
		`<!-- styles -->`,
		`<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">`)
	rsc = rsc.Replace(
		`<title>Vugu App</title>`,
		`<title>Dynamic Component Vugu</title>`)
	mux.Match(devutil.NoFileExt, rsc)
	mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	mux.Default(devutil.NewFileServer().SetDir("."))

	log.Fatal(http.ListenAndServe(l, mux))
}

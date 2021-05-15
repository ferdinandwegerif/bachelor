// +build !wasm

package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	dev := flag.Bool("dev", false, "Enable development features")
	dir := flag.String("dir", ".", "Project directory")
	httpl := flag.String("http", "127.0.0.1:8844", "Listen for HTTP on this host:port")
	flag.Parse()
	wd, _ := filepath.Abs(*dir)
	os.Chdir(wd)
	log.Printf("Starting HTTP Server at %q", *httpl)
	h := simplehttp.New(wd, *dev)

	// Modified from Vugu's documentation
	content, err := ioutil.ReadFile(wd + "/index.html")
	if err != nil {
		log.Fatal(err)
	}
	h.PageHandler = &simplehttp.PageHandler{
		Template:         template.Must(template.New("_page_").Parse(string(content))),
		TemplateDataFunc: simplehttp.DefaultTemplateDataFunc,
	}
	// ---------------------------------

	log.Fatal(http.ListenAndServe(*httpl, h))
}

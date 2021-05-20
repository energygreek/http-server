/*=============================================
#
# Last modified: 2021-05-19 13:33
#
# Filename: darkhttpd.go
#
# Description:
#
============================================*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var indexing_path string

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	file_path := filepath.Join(indexing_path, r.URL.String())
	fi, err := os.Stat(file_path)
	if err != nil {
		// 404
		fmt.Fprintln(w, "Not found: ", file_path)
		return
	}
	if fi.IsDir() {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprintln(w, "<!DOCTYPE html> <html><body>")
		// 先写特殊目录 ..
		fmt.Fprintf(w, "<p><strong> <a href=%s/>%s</a> </strong></p>\n", "..", "..")
		c, err := ioutil.ReadDir(file_path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadDir error")
		}
		for _, entry := range c {
			if entry.IsDir() {
				// 目录必须以/结尾,切记
				fmt.Fprintf(w, "<strong> <a href=%s/>%s</a> </strong> <br>\n", entry.Name(), entry.Name())
			} else {
				fmt.Fprintf(w, "<a href=%s>%s</a> <br>\n", entry.Name(), entry.Name())
			}
		}
		fmt.Fprintln(w, "</body></html>")
		return
	}
	file, err := os.ReadFile(file_path)
	w.Write(file)
	return
}

func main() {
	rootPtr := flag.String("root", ".", "root dir to serve")
	addressPtr := flag.String("addr", "0.0.0.0", "address to listen on")
	portPtr := flag.Int("poot", 8080, "port to listen on")

	flag.Parse()

	indexing_path = *rootPtr
	var err error
	_, err = os.Stat(indexing_path)
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%d", *addressPtr, *portPtr)
	fmt.Printf("listening on %s, dir %s", address, indexing_path)
	http.HandleFunc("/", index)
	http.ListenAndServe(address, nil)
}

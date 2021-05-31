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
	"net/http"
	"os"
)

func main() {
	rootPtr := flag.String("root", ".", "root dir to serve")
	addressPtr := flag.String("addr", "0.0.0.0", "address to listen on")
	portPtr := flag.Int("poot", 8080, "port to listen on")

	flag.Parse()

	var indexing_path = *rootPtr
	var err error
	_, err = os.Stat(indexing_path)
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%d", *addressPtr, *portPtr)
	fileServer := http.FileServer(http.Dir(indexing_path))
	http.Handle("/", fileServer)
	fmt.Printf("listening on %s, dir %s", address, indexing_path)
	if err := http.ListenAndServe(address, nil); err != nil{
		panic(err)
	}
}

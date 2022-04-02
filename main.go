package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	// https://pkg.go.dev/net/http
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp) // Printed everything but the actual body.
	// Body: io.ReadCloser which is an interface.
	// that means, body is flexible, it can get any type which conforms to ReadCloser interface.

	// In Go: we can different interfaces and assemble them to a single interface.

	// Why reader interface:
	// there can be lots of different sources of data, each can implement the reader interface,
	// without writing custom functions for each of these different input types.

	// ----> Reader spits out a byte slice. <-------
	// Read func
	//	- Why are we passing a byte slice to it?: It is what we want to read. It contains the raw body of response.

	// -------------> IMPORTANT <---------------
	// We create a slice, pass it to reader func, and reader injects the response in that bytes slice,
	// and returns how many bytes were read into that slice & error, if any.

	// Create a byte slice.
	// bs := []byte{} // One-way

	// make(<type>, n # of empty spaces to initialize the slice with)
	// Why? Read function is not setup to aumtomatically grow. It will quit early otherwise.

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs) // Not everytime we have to pass in bs, there are automatic functions in GO for getting the data out from Reader interface.
	// fmt.Println(string(bs))

	// How to condense down the code? and automatically get the response out.
	// 1-liner
	io.Copy(os.Stdout, resp.Body)

	// Explanation:
	//
	// - Source of data -> Reader -> []byte
	// - []byte -> Writer -> Some form of Output ( find something that implements the writer interface)
	//
	// type Writer interface {
	// Write(p []byte) (n int, err error)
	// }
	//
	// io.Copy:
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// resp.Body -> implements the Reader interface
	// os.Stdout -> implements the Writer interface

	// ----------------------------------------
	// Implement a custom Writer interface.
	// lw := logWriter{}
	// io.Copy(lw, resp.Body)

}

// After implementing this func, the log writer is now implementing the "Writer" interface.
func (logWriter) Write(bs []byte) (int, error) {
	// return 1, nil (Possible)
	fmt.Println(string(bs))
	return len(bs), nil
}

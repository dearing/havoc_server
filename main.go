// Copyright © 2016 Jacob Dearing
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	//"bytes"
	//"crypto/rand"

	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"runtime"
	"runtime/debug"
	"strconv"

	_ "net/http/pprof"

	"github.com/dearing/havoc"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", HandleIndex)
	router.GET("/mem/:value", HandleMemory)
	router.GET("/free", HandleFreeMem)
	router.GET("/cpu/:value", HandleCPU)

	go func() {
		log.Println(http.ListenAndServe(":8081", nil))
	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}

// HANDLES

func HandleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, base64.StdEncoding.EncodeToString(havoc.Data))
}

func HandleMemory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	size, err := strconv.Atoi(ps.ByName("value"))
	if err != nil {
		fmt.Fprintf(w, "error in parse: %s", err)
		return
	}

	go func() {
		havoc.SetMemory(size)
	}()

	fmt.Fprintf(w, "Mindlessly filling an array with %d indices with random bytes. Good luck!", size)

}

func HandleFreeMem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	debug.FreeOSMemory()
	fmt.Fprintf(w, "Freed.")
}

func HandleCPU(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	cpus, err := strconv.Atoi(ps.ByName("value"))
	if err != nil {
		fmt.Fprintf(w, "error in parse: %s", err)
		return
	}

	if cpus < 0 {
		cpus = 0
	} else if cpus > runtime.NumCPU() {
		cpus = runtime.NumCPU()
	}

	fmt.Fprintf(w, "setting cpu to %d!\n", cpus)

}

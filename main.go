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
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"strconv"

	_ "net/http/pprof"

	//"code.google.com/p/go-uuid/uuid"
	"github.com/dearing/havoc"
	"github.com/julienschmidt/httprouter"
)

var NAME = GetRandomName(0)

func main() {
	router := httprouter.New()
	router.GET("/", HandleIndex)

	router.GET("/mem/:value", HandleMemory)

	router.GET("/reset", HandleReset)
	router.GET("/fill", HandleFill)
	router.GET("/kill", HandleKill)
	router.GET("/forever", HandleForever)

	go func() {
		log.Println(http.ListenAndServe(":8081", nil))
	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}

// HANDLES

func HandleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "My name is %s.\n\n", NAME)
	fmt.Fprint(w, base64.StdEncoding.EncodeToString(havoc.Data))
}

func HandleMemory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	size, err := strconv.Atoi(ps.ByName("value"))
	if err != nil {
		fmt.Fprintf(w, "%s: error in parse: %s", NAME, err)
		return
	}

	go func() {
		havoc.SetMemory(size)
	}()

	fmt.Fprintf(w, "%s: %d indices set", NAME, size)

}

func HandleFill(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	havoc.FillData()
	fmt.Fprintf(w, "%s: Burning through the random.", NAME)
}

func HandleKill(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s: I hardly knew thee.", NAME)

	os.Exit(0)
}

func HandleReset(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	havoc.ResetMemory()
	fmt.Fprintf(w, "%s: Reset.", NAME)
}

func HandleFreeMem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	havoc.FreeMemory()
	fmt.Fprintf(w, "%s: Freed.", NAME)
}

func HandleForever(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	go func() {
		havoc.Forever()
	}()
	fmt.Fprintf(w, "%s: Forever is a long time.", NAME)
}

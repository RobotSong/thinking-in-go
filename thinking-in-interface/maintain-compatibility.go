// use interface maintain compatibility
package main

import "net/http"

// Adding methods to an interface breaks backwards compatibility.
// like http.ResponseWriter
//type ResponseWriter interface {
//	Header() http.Header
//	Write([]byte) (int, error)
//	WriteHeader(int)
//}
// How could you add one more method without breaking anyone’s code?
// Step 1: add the method to your concrete type implementations
// Step 2: define an interface containing the new method
// Step 3: document it

type Pusher interface {
	Push(target string, opts *http.PushOptions) error
}

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(Pusher); ok {
		_ = p.Push("style.css", nil)
	}
}

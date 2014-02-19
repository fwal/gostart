// +build appengine
package main

import (
	"net/http"
)

func init() {
	h := httpHandler()
	http.Handle("/", h)
}

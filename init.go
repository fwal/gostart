// +build appengine
package main

import (
	"net/http"
)

func init() {
	m := Martini()
	http.Handle("/", m)
}

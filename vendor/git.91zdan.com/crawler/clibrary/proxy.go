package clibrary

import "net/http"

type Proxy interface {
	client() http.Client
}

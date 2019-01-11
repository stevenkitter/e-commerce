package server

import (
	"io"
	"net/http"
)

type JsonHandler func(http.ResponseWriter, io.Reader)

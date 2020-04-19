package handler

import (
	"fmt"
	"net/http"
)

var (
	SERVICE = "{{SERVICE}}"
	VERSION = "{{VERSION}}"
)

func Info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "service name is %s and version %s\n", SERVICE, VERSION)
}

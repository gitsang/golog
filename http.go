package log

import (
	"net/http"
	"strconv"
)

func StartLogLevelHttpHandle(port int) {
	// Todo: using get to set loglevel
	http.HandleFunc("/loglevel", atomicLevel.ServeHTTP)
	go func() {
		addr := ":" + strconv.Itoa(port)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err)
		}
	}()
}


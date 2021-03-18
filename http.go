package log

import (
	"net/http"
	"strconv"
)

func StartLogLevelHttpHandle(port int) {
	http.HandleFunc("/loglevel", atomicLevel.ServeHTTP)
	go func() {
		addr := ":" + strconv.Itoa(port)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err)
		}
	}()
}


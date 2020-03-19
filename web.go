package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func setHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		printHeader(r.Header)
		if cfg.Wait != 0 {
			time.Sleep(time.Second * time.Duration(cfg.Wait))
		}
		fmt.Fprint(w, cfg.Arg)
	})
}

func startWebServer() error {
	return http.ListenAndServe(":"+strconv.Itoa(cfg.Port), nil)
}

func printHeader(h map[string][]string) {
	fmt.Fprintln(os.Stdout, "Request Header")
	fmt.Fprintln(os.Stdout, "* "+time.Now().Format("2006-01-02 15:04:05.000"))
	for k, v := range h {
		fmt.Fprintln(os.Stdout, "< "+k+": "+strings.Join(v, ", "))
	}
	fmt.Fprintln(os.Stdout, "")
}

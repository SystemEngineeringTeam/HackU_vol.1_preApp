package main

import (
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/apifuncs"
)

func main() {
	http.HandleFunc("/", apifuncs.Test)
	http.HandleFunc("/test", apifuncs.DBTest)

	http.ListenAndServe(":80", nil)
}

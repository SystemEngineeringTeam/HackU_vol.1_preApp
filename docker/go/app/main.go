package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"set1.ie.aitech.ac.jp/HackU_vol_1/apifuncs"
)

func main() {
	http.HandleFunc("/", apifuncs.Response)
	http.HandleFunc("/test", apifuncs.DBTest)


	http.ListenAndServe(":80", nil)
}

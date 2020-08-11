package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TasksResponseMethodGET は/taskでメソッドがgetだったら返す
type TasksResponseMethodGET struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Deadline string   `json:"deadline"`
	Users    []string `json:"users"`
}

// Test はテスト用
func Test(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		response := TasksResponseMethodGET{ID: 0, Name: "test", Deadline: "dead", Users: []string{"hukuda", "toyama"}}

		//構造体をバイト配列にする
		b, _ := json.Marshal(response)

		fmt.Printf("%s\n", string(b))

		//バイトは配列stringにキャスト
		fmt.Fprintln(w, string(b))

		log.Println("Get Method")

	} else if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Post Method")

		/* b, err:=ioutil.ReadAll(r.Body)
		if err !=nil{
			fmt.Println("io error")
			return
		} */

		//jsonBytes := ([]byte)(b)

		/* data:=new(InputJsonSchema)
		if err:=json.Unmarshal(jsonBytes,data);err!=nil{
			fmt.Println("JSON Unmarshal error:",err)
			return
		} */

		fmt.Fprintf(w, "jsonを受け取りました")
		/*
			fmt.Fprintf(w,data.In)
			fmt.Fprintf(w,data.Name)
			fmt.Fprintf(w,data.Description)
			fmt.Fprintf(w,data.Required)
			fmt.Fprintf(w,data.Responses)	 */

	} else if r.Method == http.MethodPut {
		fmt.Fprintln(w, "Put Method")
	} else {
		fmt.Fprintln(w, "delete Method")
	}
}

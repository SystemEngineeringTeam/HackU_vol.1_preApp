package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

// // /tasksのPOSTメソッドに対するレスポンス用構造体
// //内部のみで使う構造体の先頭は小文字
// type taskResponseMethodPOST struct {
// 	ID int `json:"id"`
// }

//TaskResponse は/tasksに対する処理をする
func TaskResponse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		// response := dbctl.Task{ID: 0, Title: "test", Deadline: "dead", Users: []string{"fukuda", "toyama"}}

		// //構造体をバイト配列にする
		// b, _ := json.Marshal(response)

		// fmt.Printf("%s\n", string(b))

		// //バイトは配列stringにキャスト
		// fmt.Fprintln(w, string(b))

		// Task型の配列を受け取る
		// dbctl.CallTasks()
		// func CallTasks() ([]Task, error) {
		// func 関数名(引数 引数の型) (戻り値の型,戻り値の型...){
		// }

		// retOne()
		// one:=retOne()

		//セキリティ設定
		w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.

		tasks, err := dbctl.CallTasks()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
		}

		// 受け取った配列をJSONに変換
		jsonBytes, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
		}
		// []byteをstringに変換
		jsonString := string(jsonBytes)

		// httpステータスコードを返す<-New
		w.WriteHeader(http.StatusOK)
		// JSONを返す
		fmt.Fprintln(w, jsonString)

	} else if r.Method == http.MethodPost {
		log.Println("PostMethod")

		//jsonを読み込む
		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("io error")
			return
		}

		//構造体の初期化
		data := dbctl.Task{}

		//処理が終わったらjsonを構造体にする
		if err := json.Unmarshal(jsonBytes, &data); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
			return
		}

		//データベースに受けっとた情報を登録
		n, err := dbctl.RegisterNewTask(data)
		if err != nil {
			fmt.Println(err)
		}

		// IDをjsonに変換// "{\"name\":" + strconv.Itoa(n) + "}"がjsonの形式
		newTaskID := "{\"id\":" + strconv.Itoa(n) + "}"
		// fmt.Println(newTaskID)
		//{"id":1234}

		//構造体を返す
		fmt.Fprintln(w, newTaskID)

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

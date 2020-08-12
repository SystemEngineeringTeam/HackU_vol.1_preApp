package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

// // /tasksのPOSTメソッドに対するレスポンス用構造体
// //内部のみで使う構造体の先頭は小文字
// type taskResponseMethodPOST struct {
// 	ID int `json:"id"`
// }

//TaskResponse は/tasksに対する処理をする
func TaskResponse(w http.ResponseWriter, r *http.Request) {
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.

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
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("io error")
			return
		}

		//構造体の初期化
		data := dbctl.Task{}

		//処理が終わったらjsonを構造体にする
		if err := json.Unmarshal(jsonBytes, &data); err != nil {

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("JSON Unmarshal error:", err)
			return
		}

		//データベースに受けっとた情報を登録
		n, err := dbctl.RegisterNewTask(data)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println(err)

		}
		w.WriteHeader(http.StatusOK)

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
		//構造体の初期化
		task := dbctl.Task{}
		pathString := strings.ReplaceAll(r.URL.Path, "/tasks/", "")
		pathNumber, err := strconv.Atoi(pathString)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		task.ID = pathNumber

		//jsonを読み込み
		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			//io(input,output)
			fmt.Println("io error")
			return
		}

		if err := json.Unmarshal(jsonBytes, &task); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("JSON UNmarshal error", err)
			return
		}

		if err := dbctl.PutTasks(task); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)

		log.Println("Put Method")
	} else if r.Method == http.MethodDelete {
		//tasks/1 , ""←空文字
		idString := strings.ReplaceAll(r.URL.Path, "/tasks/", "")
		// <-Method:DELETE idString="",idString="1"

		idNum, err := strconv.Atoi(idString)
		//文字列を数字に変換
		if err != nil {
			//エラー処理
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println(err)
			return
		}

		if err := dbctl.DeleteTask(idNum); err != nil {

			w.WriteHeader(http.StatusServiceUnavailable)
			//エラー処理
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// UsersResponse は/usersに関する処理を行う
func UsersResponse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")    // Allow any access.
	w.Header().Add("Access-Control-Allow-Methods", "GET") // Allowed methods.

	if r.Method == http.MethodGet {
		users, err := dbctl.CallUsers()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
		}

		//jsonに変換//byte？
		jsonBytes, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		jsonstring := string(jsonBytes)

		// httpステータスコードを返す<-New
		w.WriteHeader(http.StatusOK)
		// JSONを返す
		fmt.Fprintln(w, jsonstring)

	}

}

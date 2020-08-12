package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

// DBTest はデータベース関数のテストする関数です
func DBTest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasks, err := dbctl.CallTasks()
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.Marshal(tasks)
		if err != nil {
			log.Fatal(err)
		}
		taskString := string(b)
		fmt.Fprintln(w, taskString)
		fmt.Println(taskString, "hogehoge")
	}
}

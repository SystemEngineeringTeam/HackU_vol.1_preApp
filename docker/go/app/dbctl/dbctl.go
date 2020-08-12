package dbctl

import (
	"database/sql"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

// Task はtaskを扱う
type Task struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Deadline string   `json:"deadline"`
	Users    []string `json:"users"`
}

//エラーの内容:err 関数の名前:f.Name() ファイルのパス:file runtimeが呼ばれた行数:line
const errFormat = "%v\nfunction:%v file:%v line:%v\n"

var db *sql.DB

// packageがimportされたときに呼び出される関数
func init() {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	var err error

	db, err = sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/pre_app_db")
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)

		//データベースを開けないと動作が継続できないためpanicを発生させる
		panic("Can't Open database.")
	}
}

// CallTasks はデータベースからタスク一覧を取り出す関数
func CallTasks() ([]Task, error) {
	tasks := make([]Task, 0)
	id := 0
	title := ""
	deadline := ""

	rows, err := db.Query("select id,title,deadline from tasks")
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &title, &deadline)
		users, err := callUsersFromTaskID(id)
		if err != nil {
			return nil, err
		}
		temporaryTask := Task{ID: id, Title: title, Deadline: deadline, Users: users}
		tasks = append(tasks, temporaryTask)
	}

	return tasks, nil
}

func callUsersFromTaskID(taskID int) ([]string, error) {
	rows, err := db.Query("select user_id from links_table where task_id=?", taskID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return nil, err
	}
	defer rows.Close()

	names := make([]string, 0)

	userID := 0
	for rows.Next() {
		rows.Scan(&userID)
		nextName, err := callUserNameFromUserID(userID)
		if err != nil {
			return nil, err
		}
		names = append(names, nextName)
	}

	return names, nil
}

func callUserNameFromUserID(userID int) (string, error) {
	rows, err := db.Query("select name from users where id=?", userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return "", err
	}
	defer rows.Close()

	name := ""
	for rows.Next() {
		rows.Scan(&name)

	}
	return name, nil
}

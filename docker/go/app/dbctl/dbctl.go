package dbctl

import (
	"database/sql"
	"log"
	"runtime"
)

// Tasks は/taskでメソッドがgetだったら返す
type Tasks struct {
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
func CallTasks() ([]Tasks, error) {
	tasks := make([]Tasks, 0)
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
		temporaryTask := Tasks{ID: id, Title: title, Deadline: deadline, Users: users}
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

// RegisterNewTask はデータベースにタスクを追加する関数です
func RegisterNewTask(task Tasks) (int, error) {
	_, err := db.Query("insert into tasks(title,deadline) values (?,?)", task.Title, task.Deadline)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	rows, err := db.Query("select id from tasks where title=?", task.Title)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&task.ID)
	}

	userIDs := make([]int, 0)

	for _, name := range task.Users {
		id, err := callUserIDFromName(name)
		if err != nil {
			return -1, err
		}

		userIDs = append(userIDs, id)
	}

	for _, u := range userIDs {
		err := linkTaskIDAndUserID(task.ID, u)
		if err != nil {
			return -1, err
		}
	}

	return task.ID, nil
}

func callUserIDFromName(name string) (int, error) {
	rows, err := db.Query("select id from users where name=?", name)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	id := 0
	for rows.Next() {
		rows.Scan(&id)
	}

	return id, nil
}

func linkTaskIDAndUserID(taskID, userID int) error {
	_, err := db.Query("insert into links_table(task_id,user_id)", taskID, userID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask はデータベースからタスクを削除する関数です
func DeleteTask(id int) error {
	_, err := db.Query("delete from tasks where id = ?", id)
	if err != nil {
		return err
	}
	_, err = db.Query("delete from links_table where task_id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

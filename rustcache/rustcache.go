package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os/user"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/img", imageHandler)
	log.Println("Server running on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	entity := r.URL.Query().Get("entity")
	crc := r.URL.Query().Get("crc")

	if len(entity) == 0 && len(crc) == 0 {
		http.NotFound(w, r)
		return
	}

	query := "SELECT data FROM data where entity = ? and crc = ?"
	usr, err := user.Current()
	checkErr(err)
	db, err := sql.Open("sqlite3", usr.HomeDir+"/Library/Application Support/Steam/SteamApps/common/Rust/cache/Storage.db")
	defer db.Close()
	var data []byte
	err = db.QueryRow(query, entity, crc).Scan(&data)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
	case err != nil:
		checkErr(err)
	}
	w.Header().Set("Content-Type", "image/png")
	_, err = w.Write(data)
	checkErr(err)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	last := r.URL.Query().Get("last")
	query := "SELECT * FROM data"
	if len(last) != 0 {
		query += " where lastaccess > ?"
	}
	query += " order by lastaccess desc"

	usr, err := user.Current()
	checkErr(err)
	db, err := sql.Open("sqlite3", usr.HomeDir+"/Library/Application Support/Steam/SteamApps/common/Rust/cache/Storage.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query(query, last)
	checkErr(err)
	defer rows.Close()

	type Data struct {
		Entity     int64
		Crc        int64
		Num        int64
		LastAccess int64
		Data       []byte // TODO: Consider creating a separate data endpoint.
	}

	var dataResults []Data

	for rows.Next() {
		var crc int64
		var entity int64
		var num int64
		var lastaccess int64
		var data []byte
		err = rows.Scan(&crc, &entity, &num, &lastaccess, &data)
		checkErr(err)
		d := Data{Entity: entity, Crc: crc, Num: num, LastAccess: lastaccess, Data: data}
		dataResults = append(dataResults, d)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataResults)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

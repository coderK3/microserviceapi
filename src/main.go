package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type repoSummary struct {
	ID         int
	Name       string
	Owner      string
	TotalStars int
}

type repo struct {
	Repositories []repoSummary
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "vigyaa"
	password = "vigyaa@123"
	dbname   = "vigyaa"
)

func main() {
	initDB()
	http.HandleFunc("/api/func", indexHandler)
	http.HandleFunc("/api/repo", repoHandler)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

// func main() {
// 	db = initDB()
// 	rows, err := db.Query("dt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(rows)
// }

func indexHandler(w http.ResponseWriter, r *http.Request) {
	repos := repo{}

	err := queryRepos(&repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func queryRepos(repos *repo) error {
	rows, err := db.Query(`SELECT
	id,
	repository_owner,
	repository_name,
	total_stars
	FROM repositories
	ORDER BY total_stars DESC`)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		repo := repoSummary{}
		err := rows.Scan(
			&repo.ID,
			&repo.Owner,
			&repo.Name,
			&repo.TotalStars,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func repoHandler(w http.ResponseWriter, r *http.Request) {

}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected")
	return db
}

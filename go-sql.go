package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
)

func main() {
	var marks string
	var age string

	db, err := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/lab_dbms")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from table_aggregrate")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&marks, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(marks, age)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// --------------------------------> Database

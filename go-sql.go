package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

//var (
//	age int
//	marks int
//)
// https://api.coinmarketcap.com/v1/ticker/bitcoin/
type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	id string `json:"id"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}
func main() {
	url := "https://api.coinmarketcap.com/v1/ticker/?limit=10"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()
	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Phone No. = ", record.id)
	//fmt.Println("Country   = ", record.CountryName)
	//fmt.Println("Location  = ", record.Location)
	//fmt.Println("Carrier   = ", record.Carrier)
	//fmt.Println("LineType  = ", record.LineType)


}

// --------------------------------> Database
//db, err := sql.Open("mysql",
//	"root:123@tcp(localhost:3306)/lab_dbms")
//if err != nil {
//	log.Fatal(err)
//}
//defer db.Close()
//rows, err := db.Query("select * from table_aggregrate")
//if err != nil {
//	log.Fatal(err)
//}
//defer rows.Close()
//for rows.Next() {
//	err := rows.Scan(&marks, &age)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(marks, age)
//}
//err = rows.Err()
//if err != nil {
//	log.Fatal(err)
//}
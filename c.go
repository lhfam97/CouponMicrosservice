package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)
const (
  host     = "tcc_pg"
  port     = 5432
  user     = "postgres"
  password = "docker"
  dbname   = "monolito"
)

var (
	name string
	discount float64
)
type Coupon struct {
	id        string
	name       string
	discount float64
}



type Result struct {
	Status string
	Discount float64
}


func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}


func home(w http.ResponseWriter, r *http.Request) {
	couponName := r.URL.Query().Get("coupon")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname) 	

	db, err := sql.Open("postgres", psqlInfo)

	sqlStatement := `SELECT name,discount FROM coupons WHERE name='`+couponName+`';`
	
	// var coupon Coupon
	rows, err1:= db.Query(sqlStatement)
	if err1 != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	name := "No Coupon"
	i := 0
	discount := float64(i)

	for rows.Next() {
		err := rows.Scan(&name, &discount)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, discount)
	}


	switch err1 {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(discount)
	default:
		panic(err)
	}

	if err != nil {
		
  panic(err)
	}
	defer db.Close()




result := Result{
	Status: name,
	Discount: discount,
}


	jsonResult, err := json.Marshal(result)
	if err != nil {
		
		log.Fatal("Error converting json")
	}
	fmt.Fprintf(w, string(jsonResult)) 
}
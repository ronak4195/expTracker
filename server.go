package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type entry struct {
	name string
}

func insert(db *sql.DB, e entry) {
	query := "INSERT INTO name(name) VALUES (?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, e.name)
	if err != nil {
		log.Printf("Error %s when inserting row into name table", err)
		return
	}
	fmt.Println("Inserted row successfully.")
}

func count(db *sql.DB) {
	query := "SELECT COUNT(*) FROM name"
	var cot int
	err := db.QueryRow(query).Scan(&cot)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	fmt.Printf("The count is %v \n", cot)
}

func main() {
	db, err := sql.Open("mysql", "root:Str0ng!Passw0rd@tcp(0.0.0.0:3306)/nameapi")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var cho int
	fmt.Println("Choose 1 to add name, 2 to view total names")
	fmt.Scanln(&cho)

	switch cho {
	case 1:
		cc(db)
	case 2:
		count(db)
	default:
		fmt.Print("Wrong Choice")
	}
	defer db.Close()
}

func cc(db *sql.DB) {
	var n string
	fmt.Print("What is the name to be added?")
	fmt.Scanf("%s", &n)
	x := entry{
		name: n,
	}
	insert(db, x)
}

/// main - start pplicstion
// handler - when you hit the endoint curl

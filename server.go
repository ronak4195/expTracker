package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type entry struct {
	Name string `json:"name"`
}

var db *sql.DB

// Handler to insert an entry
func insertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var e entry
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO name(name) VALUES (?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		http.Error(w, "Error preparing SQL statement", http.StatusInternalServerError)
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, e.Name)
	if err != nil {
		http.Error(w, "Error inserting entry", http.StatusInternalServerError)
		log.Printf("Error %s when inserting row into name table", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inserted row successfully."))
}

// Handler to count entries
func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	query := "SELECT COUNT(*) FROM name"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		http.Error(w, "Error counting entries", http.StatusInternalServerError)
		log.Printf("Error %s when counting rows", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:Str0ng!Passw0rd@tcp(0.0.0.0:3306)/nameapi")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/count", countHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/// main - start pplicstion
// handler - when you hit the endoint curl

/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	dbHost = "DB_HOST"
	dbUser = "DB_USER"
	dbPwd  = "DB_PASSWORD"
	dbName = "DB_DBNAME"
	port   = "8080"
)

var (
	dbClient *sql.DB
)

type musician struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Instrument string `json:"instrument,omitempty"`
}

type musicians struct {
	Musicians []musician `json:"musicians"`
}

func getDSN(usr, pwd, host string) string {
	cred := usr
	if pwd != "" {
		cred = cred + ":" + pwd
	}
	return fmt.Sprintf("%s@tcp(%s)/", cred, host)
}

func main() {
	host := os.Getenv(dbHost)
	if host == "" {
		log.Fatal(dbHost + " environment variable unspecified.")
	}
	user := os.Getenv(dbUser)
	if user == "" {
		log.Fatal(dbUser + " environment variable unspecified.")
	}
	pwd := os.Getenv(dbPwd)
	if pwd == "" {
		log.Print(dbPwd + " environment variable unspecified or ''")
	}
	databaseName := os.Getenv(dbName)
	if databaseName == "" {
		log.Fatal(dbName + " environment variable unspecified.")
	}

	var err error
	dbClient, err = sql.Open("mysql", getDSN(user, pwd, host))
	if err != nil {
		log.Fatalf("Failed to open sql client: %v", err)
	}

	if _, err := dbClient.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", databaseName)); err != nil {
		log.Fatalf("Failed to create a database: %v", err)
	}

	if _, err := dbClient.Exec(fmt.Sprintf("USE %s;", databaseName)); err != nil {
		log.Fatalf("Failed to use a database: %v", err)
	}

	if err := dbClient.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Printf("successfully connected to database: %v", dbClient)

	r := mux.NewRouter()
	r.HandleFunc("/reset", reset).Methods("POST")
	r.HandleFunc("/musicians", read).Methods("GET")
	r.HandleFunc("/musicians", write).Methods("POST")
	log.Println("Server listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func reset(w http.ResponseWriter, r *http.Request) {
	if _, err := dbClient.Exec("DROP TABLE IF EXISTS musicians"); err != nil {
		log.Printf("Error dropping musicians table: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if _, err := dbClient.Exec("CREATE TABLE musicians(id SERIAL, name VARCHAR(255), instrument VARCHAR(255))"); err != nil {
		log.Printf("Error creating musicians table: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte("{ \"status\": \"database created\" }")); err != nil {
		log.Printf("write did not complete: %v", err)
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := dbClient.Query("SELECT * FROM musicians")
	if err != nil {
		log.Printf("Error reading musicians table: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	defer rows.Close()

	e := musicians{Musicians: []musician{}}
	for rows.Next() {
		var id, name, instrument string
		err := rows.Scan(&id, &name, &instrument)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}
		log.Printf("%s: %s %s\n", id, name, instrument)
		e.Musicians = append(e.Musicians, musician{ID: id, Name: name, Instrument: instrument})
	}

	j, err := json.Marshal(e)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(j); err != nil {
		log.Printf("write did not complete: %v", err)
	}
}

func write(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		writeError(w, http.StatusBadRequest, "cannot read the request body")
		return
	}

	var e musician
	if err := json.Unmarshal(body, &e); err != nil {
		log.Printf("Error parsing request body: %v", err)
		writeError(w, http.StatusBadRequest, "cannot parse the reqeust body - invalid JSON")
		return
	}
	if e.Name == "" {
		writeError(w, http.StatusBadRequest, "missing required property: \"name\"")
		return
	}
	if e.Instrument == "" {
		writeError(w, http.StatusBadRequest, "missing required property: \"instrument\"")
		return
	}

	stmt, err := dbClient.Prepare("INSERT INTO musicians(name, instrument) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error writing to database: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	res, err := stmt.Exec(e.Name, e.Instrument)
	if err != nil {
		log.Printf("Error writing to database: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	e.ID = strconv.FormatInt(id, 10)
	j, err := json.Marshal(e)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(j); err != nil {
		log.Printf("write did not complete: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	j, err := json.Marshal(
		struct {
			Error string `json:"error"`
		}{
			Error: message,
		})
	if err != nil {
		if _, err := w.Write([]byte("{\"error\": \"internal server error\"}")); err != nil {
			log.Printf("write did not complete: %v", err)
		}
		return
	}
	if _, err := w.Write(j); err != nil {
		log.Printf("write did not complete: %v", err)
	}
}

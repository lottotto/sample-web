package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

// Employee  db columns info
type Employee struct {
	Department string `json:"department"`
	Group      string `json:"group"`
	Name       string `json:"name"`
	Position   string `json:"position"`
}

// Env is database connect info
type Env struct {
	DbHost     string `default:"localhost"`
	DbUser     string `default:"postgres"`
	DbPassword string `default:"postgres"`
	DbName     string `default:"postgres"`
}

func getAllEmployee(c *gin.Context) {

	var goenv Env
	envconfig.Process("", &goenv)
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", goenv.DbHost, goenv.DbUser, goenv.DbPassword, goenv.DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}
	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.Department, &e.Group, &e.Name, &e.Position)
		employees = append(employees, e)
	}
	c.HTML(http.StatusOK, "list.html", gin.H{
		"data": employees,
	})
}

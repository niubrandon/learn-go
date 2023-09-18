package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// use the OS environment
	// os.LookupEnv
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")

	// pq
  // https://pkg.go.dev/github.com/lib/pq
	// 
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	// connect to a database uring a URL
	// dbname
	// user
	// password
	// host
	// port
	// sslmode
	// use single quotes for values that contain whitespace "user=pqgotest password='with spaces'"
	// a backslash will escape the next character in values: "user=space\ man password='it\'s valid'"
	
	// support Bulk imports
	// support Notifications
	// support Kerberos authentication

	// Data Types: parameters pass through driver.DefaultParameterConverter before they are handled by this package
	// integer types smallint, integer, and bigint are returned as int64
	// floating-point types real and double precision are returned as float64
	// character types char, varchar, and text are returned as string
	// temporal types date, time, timetz, timestamp, and timestamptz are returned as time.Time
	// the boolean type is returned as bool
	// the bytea type is returned as []byte
	// TO-DO setup envs
	// TO-DO create a production postgres database
	connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type User struct {
		ID int
		Name string
	}

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			fmt.Printf("ID: %D, Name: %s\n", user.ID, user.Name)
		}
	}
	/*
		gin
	*/
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
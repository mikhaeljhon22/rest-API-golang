package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
    "os"
    "github.com/joho/godotenv"
)

type User struct {
    Username string `json:"Username"`
    Password string `json:"Password"`
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("error load env variable")
    }

    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    // Open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    // Close db
    defer db.Close()

    // Check db
    err = db.Ping()
    CheckError(err)

    router := gin.Default()

    router.POST("/post/user", func(c *gin.Context) {
        var user User 

        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "invalid error",
                "error": err.Error(),
            })
            return
        }

        // Insert data into the database
        insertStmt := `INSERT INTO "User" (username, password) VALUES ($1, $2)`
        _, e := db.Exec(insertStmt, user.Username, user.Password)
        CheckError(e)

        c.JSON(200, gin.H{
            "message": "success",
        })
    })

	router.GET("get/user", func(c *gin.Context){
		username:= c.Query("username")

		row:= db.QueryRow(`SELECT username, password FROM "User" WHERE username = $1`, username)

		var uname, pwd string
		err := row.Scan(&uname,&pwd)

		if err != nil{
			c.JSON(404, gin.H{
				"message": "data not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"Username": uname,
			"Password": pwd,
		})
	})

    router.Run(":8080")
}


func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

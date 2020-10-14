package main

import (
    "database/sql"
    "fmt"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func index(c echo.Context) error {
    // DB接続部分。接続文字列は(MySQLユーザー名:パスワード@tcp(ホスト名:ポート)/DB名)
    db, err := sql.Open("mysql", "test:test@tcp(db:3306)/sample")

    if err != nil {
        return c.String(http.StatusInternalServerError, "DB Error")
    }
    defer db.Close()

    stmt, err := db.Prepare("select * from users where id < ?")
    if err != nil {
        return c.String(http.StatusInternalServerError, "SQL Error")
    }
    defer stmt.Close()

    var (
        id   int
        name string
    )
    rows, err := stmt.Query(2)

    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            return c.String(http.StatusInternalServerError, "Error")
        }
        fmt.Println(id, name)
    }
    defer rows.Close()

    if err = rows.Err(); err != nil {
        return c.String(http.StatusInternalServerError, "Error")
    }
    return c.String(http.StatusOK, "Hello!!!, "+name)
}

func main() {
    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/", index)

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}



// package main

// import (
//     "net/http"

//     "github.com/labstack/echo/v4"
// )

// func main() {
//     e := echo.New()
//     e.GET("/", func(c echo.Context) error {
//         return c.String(http.StatusOK, "Hello, World!")
//     })

//     e.Logger.Fatal(e.Start(":1323"))
// }

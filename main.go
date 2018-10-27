package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    connStr = "user=postgres dbname=youtube host=192.168.1.63 port=30000 sslmode=disable"
)

func connection() *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    return db
}

func channels() []string {
    sqlStr := "SELECT serial FROM youtube.entities.channels ORDER BY RANDOM() LIMIT 50"
    db := connection()
    defer func() {
        err := db.Close()
        if err != nil {
            panic(err)
        }
    }()

    row, err := db.Query(sqlStr)
    if err != nil {
        panic(err)
    }

    serials := make([]string, 50)
    var idx uint8
    for row.Next() {
        var serial string

        err = row.Scan(&serial)
        if err != nil {
            panic(err)
        }

        serials[idx] = serial
        idx++
    }

    return serials
}

func main() {
    fmt.Println("Hello, world!")

    chans := channels()
    for i := 0; i < len(chans); i++ {
        fmt.Println(chans[i])
    }
}

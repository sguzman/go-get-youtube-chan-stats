package main

import (
    "database/sql"
    "fmt"
    "math/rand"
    "os"
    _ "github.com/lib/pq"
    "github.com/imroc/req"
    "strings"
    "time"
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

func getKey() string {
    rawKey := os.Getenv("API_KEY")
    splitKeys := strings.Split(rawKey, "|")

    return splitKeys[rand.Intn(len(splitKeys))]
}

func getData(cs []string) string {
    key := getKey()
    url := "https://www.googleapis.com/youtube/v3/channels"
    partStr := "snippet,contentDetails,brandingSettings,contentOwnerDetails,invideoPromotion,localizations,status,topicDetails"
    idStr := strings.Join(cs, ",")

    param := req.Param{
        "part":  partStr,
        "id": idStr,
        "key": key,
    }

    r, err := req.Get(url, param)
    if err != nil {
        panic(err)
    }

    body, err := r.ToString()
    if err != nil {
        panic(err)
    }

    fmt.Println(body)

    return ""
}

func main() {
    for {
        rand.Seed(time.Now().Unix())

        chans := channels()
        getData(chans)
    }
}

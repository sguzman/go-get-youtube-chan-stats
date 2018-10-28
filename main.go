package main

import (
    "database/sql"
    "fmt"
    "github.com/imroc/req"
    _ "github.com/lib/pq"
    "math/rand"
    "os"
    "runtime"
    "strings"
    "time"
)

const (
    connStr = "user=postgres dbname=youtube host=192.168.1.63 port=30000 sslmode=disable"
)

type Data struct {
    title string
    serial string
    customUrl string
    description string
    country string
    publishedAt string
}

func connection() *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    return db
}

func channels() []string {
    sqlStr := "select C.serial from youtube.entities.channels C WHERE C.serial NOT IN (select C.serial from youtube.entities.chans C) LIMIT 50"
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

func getJson(cs []string) interface{} {
    key := getKey()
    url := "https://www.googleapis.com/youtube/v3/channels"
    partStr := "snippet,topicDetails"
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

    var foo interface{}
    err = r.ToJSON(&foo)
    if err != nil {
        panic(err)
    }

    return foo
}

func getData(cs []string) []Data {
    jsonMap := getJson(cs).(map[string]interface{})
    items := jsonMap["items"].([]interface{})

    datas := make([]Data, len(cs))
    for i := range items {
        var data Data
        item := items[i].(map[string]interface{})
        {
            data.serial = item["id"].(string)
            {
                snippet := item["snippet"].(map[string]interface{})
                data.title = snippet["title"].(string)
                data.description = snippet["description"].(string)
                data.customUrl = snippet["customUrl"].(string)
                data.publishedAt = snippet["publishedAt"].(string)
                data.country = snippet["country"].(string)
            }
        }

        fmt.Println(data)
        datas[i] = data
    }

    return datas
}

func main() {
    for {
        rand.Seed(time.Now().Unix())

        chans := channels()
        getData(chans)
        runtime.GC()
    }
}

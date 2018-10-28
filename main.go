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
    defaultHost = "192.168.1.63"
    defaultPort = "30000"
)

func connStr() string {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    if len(host) == 0 || len(port) == 0 {
        return fmt.Sprintf("user=postgres dbname=youtube host=%s port=%s sslmode=disable", defaultHost, defaultPort)
    } else {
        return fmt.Sprintf("user=postgres dbname=youtube host=%s port=%s sslmode=disable", host, port)
    }
}

type Data struct {
    title     string
    serial    string
    customUrl *string
    country   *string
    joined    string
}

func nilToEmpty(str *string) string {
    if str == nil {
        return "nil"
    }

    return *str
}

func (that Data) String() string {
    return fmt.Sprintf("{%s, %s, %s, %s, %s}",
        that.title, that.serial, nilToEmpty(that.customUrl), nilToEmpty(that.country), that.joined)
}

func connection() *sql.DB {
    db, err := sql.Open("postgres", connStr())
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
                if snippet["customUrl"] == nil {
                    data.customUrl = nil
                } else {
                    str := snippet["customUrl"].(string)
                    data.customUrl = &str
                }

                data.joined = snippet["publishedAt"].(string)
                if snippet["country"] == nil {
                    data.country = nil
                } else {
                    str := snippet["country"].(string)
                    data.country = &str
                }
            }
        }

        fmt.Println(data)
        datas[i] = data
    }

    return datas
}

func insert(ds []Data) {
    db := connection()
    defer func() {
        err := db.Close()
        if err != nil {
            panic(err)
        }
    }()

    sqlInsert := "INSERT INTO youtube.entities.chans (serial, title, custom_url, country, joined) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING"

    for i := range ds {
        d := ds[i]

        _, err := db.Exec(sqlInsert, d.serial, d.title, d.customUrl, d.country, d.joined)
        if err != nil {
            panic(err)
        }
    }
}

func main() {
    rand.Seed(time.Now().Unix())
    for {
        chans := channels()
        datas := getData(chans)
        insert(datas)

        runtime.GC()
    }
}

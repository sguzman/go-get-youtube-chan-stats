package main

import (
    "./structs"
    "./postgres"
    "encoding/json"
    "fmt"
    "github.com/imroc/req"
    "math/rand"
    "os"
    "runtime"
    "strings"
    "time"
)

func getKey() string {
    rawKey := os.Getenv("API_KEY")
    splitKeys := strings.Split(rawKey, "|")

    return splitKeys[rand.Intn(len(splitKeys))]
}

func getJson(cs []string) structs.ResponseType {
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

    str, err := r.ToBytes()
    if err != nil {
        panic(err)
    }

    var data structs.ResponseType
    err = json.Unmarshal(str, &data)
    if err != nil {
        panic(err)
    }

    return data
}

func main() {
    rand.Seed(time.Now().Unix())
    for {
        chans := postgres.Channels()
        data := getJson(chans)
        fmt.Println(data)

        runtime.GC()
    }
}

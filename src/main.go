package main

import (
    "./postgres"
    "./structs"
    "./youtube"
    "fmt"
    "runtime"
)

func main() {
    for {
        chans := postgres.Channels()
        data := youtube.Get(chans)
        items := structs.Transform(data.Items)
        fmt.Println(items)

        runtime.GC()
    }
}

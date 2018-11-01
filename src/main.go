package main

import (
    "./postgres"
    "./youtube"
    "fmt"
    "runtime"
)

func main() {
    for {
        chans := postgres.Channels()
        data := youtube.Get(chans)
        fmt.Println(data)

        runtime.GC()
    }
}

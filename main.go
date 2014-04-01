
package main

import (
    "os"
    "log"
    "image/draw"
)

func main() {
    img := make(chan draw.Image)
    go loadInput(os.Stdin, img)
    for {
        i, ok := <-img
        if !ok { break }
        err := displayImage(i)
        if err != nil {
            log.Fatal(err)
        }
    }
}

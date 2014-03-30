
package main

import (
    "os"
    "log"
    "image"
    _ "image/png"
    _ "image/jpeg"
)

func main() {
    img := make(chan image.Image)
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

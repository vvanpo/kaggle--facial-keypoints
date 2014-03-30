
package main

import (
    "log"
    "image"
    _ "image/png"
    _ "image/jpeg"
)

func main() {
    img := make(chan image.Image)
    go loadTest(img)
    for {
        err := displayImage(<-img)
        if err != nil {
            log.Fatal(err)
        }
    }
}

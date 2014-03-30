
package main

import (
    "os"
    "log"
    "image"
    _ "image/png"
    _ "image/jpeg"
)

func main() {

    for _, file := range os.Args[1:] {
        f, err := os.Open(file)
        if err != nil {
            log.Print("Failed to open file ", file)
            continue
        }
        img, _, err := image.Decode(f)
        if err != nil {
            log.Print("Failed to decode image ", file)
            continue
        }
        err = displayImage(img)
        if err != nil {
            log.Print(err)
        }
    }
}

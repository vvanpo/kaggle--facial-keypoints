
package main

import (
    "io"
    "log"
    "image"
    "strings"
    "strconv"
    "encoding/csv"
)

var config map[string]string

type format int

const (
    None format = iota
    PGM format = iota
)

func testFormat(r *csv.Reader) (f format) {
    h, _ := r.Read()
    if len(h) < 2 { return }
    switch {
        case h[len(h) - 1] == "Image": f = PGM
        default: f = None
    }
    return
}

func loadInput(in io.Reader, out chan image.Image) {
    defer close(out)
    r := csv.NewReader(in)
    r.TrimLeadingSpace = true
    fmt := testFormat(r)
    if fmt == None { return }
    for {
        rec, err := r.Read()
        if err != nil {
            if err == io.EOF { break }
            log.Print(err)
            return
        }
        img, err := decodePGM(rec[len(rec) - 1])
        if err != nil {
            log.Print(err)
            return
        }
        out <- img
    }
}

func decodePGM(pgm string) (image.Image, error) {
    img := image.NewGray(image.Rect(0, 0, 96, 96))
    f := strings.Fields(pgm)
    for i, v := range f {
        n, err := strconv.Atoi(v)
        img.Pix[i] = uint8(n)
        if err != nil {
            return nil, err
        }
    }
    return img, nil
}

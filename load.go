
package main

import (
    "io"
    "log"
    "errors"
    "image"
    "image/color"
    "image/draw"
    "strings"
    "strconv"
    "encoding/csv"
)

func decodeFunc(h []string) (func(rec []string) (draw.Image, error)) {
    if h[0] == "ImageId" { return decodePGM }
    if h[0] == "left_eye_center_x" {
        return func(rec []string) (draw.Image, error) {
            src, err := decodePGM(rec)
            if err != nil { return nil, err }
            b := src.Bounds()
            img := image.NewRGBA(b)
            draw.Draw(img, b, src, b.Min, draw.Src)
            for i := 0; i < len(rec) / 2; i++ {
                x, err := strconv.ParseFloat(rec[2*i], 64)
                if err != nil {
                    log.Print(err)
                    continue
                }
                y, err := strconv.ParseFloat(rec[2*i + 1], 64)
                if err != nil {
                    log.Print(err)
                    continue
                }
                img.Set(int(x + 0.5), int(y + 0.5), color.RGBA{0xff, 0, 0, 0xff})
            }
            return img, nil
        }
    }
    return nil
}

func loadInput(in io.Reader, out chan draw.Image) {
    defer close(out)
    r := csv.NewReader(in)
    r.TrimLeadingSpace = true
    h, err := r.Read()
    if err != nil {
        log.Print(err)
        return
    }
    dec := decodeFunc(h)
    if dec == nil {
        log.Print(errors.New("Invalid input format"))
        return
    }
    for {
        rec, err := r.Read()
        if err != nil {
            if err == io.EOF { break }
            log.Print(err)
            return
        }
        img, err := dec(rec)
        if err != nil {
            log.Print(err)
            return
        }
        out <- img
    }
}

func decodePGM(rec []string) (draw.Image, error) {
    f := strings.Fields(rec[len(rec) - 1])
    img := image.NewGray(image.Rect(0, 0, 96, 96))
    for i, v := range f {
        n, err := strconv.Atoi(v)
        img.Pix[i] = uint8(n)
        if err != nil {
            return nil, err
        }
    }
    return img, nil
}


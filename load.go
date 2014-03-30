
package main

import (
    "os"
    "io"
    "io/ioutil"
    "log"
    "image"
    "strings"
    "strconv"
    "encoding/csv"
    "gopkg.in/yaml.v1"
)

var config map[string]string

func init() {
    if len(os.Args) != 2 {
        log.Fatal("Need to specify configuration file")
    }
    bytes, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal("Cannot load configuration file ", os.Args[1])
    }
    err = yaml.Unmarshal(bytes, &config)
    if err != nil {
        log.Print(err)
        log.Fatal("Cannot parse configuration file")
    }
}

func loadTest(out chan image.Image) {
    defer close(out)
    f, err := os.Open(config["test.csv"])
    if err != nil {
        log.Print(err)
        return
    }
    test := csv.NewReader(f)
    test.TrimLeadingSpace = true
    for {
        record, err := test.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Print(err)
            return
        }
        if record[1] == "Image" { continue }
        img, err := decodePGM(record[1])
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


package main

import (
    "os"
    "io/ioutil"
    "log"
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
    err = yaml.Unmarshal(bytes, config)
    if err != nil {
        log.Print(err)
        log.Fatal("Cannot parse configuration file")
    }
}

func loadTest() {
    f, err := os.Open(config["test.csv"])
    if err != nil {
        log.Fatal(err)
    }
    r := csv.NewReader(f)
    print(r)
}

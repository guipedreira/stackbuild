package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
)

type Config struct {
    Foo string
    Bar []string
    Driver string
}

func main() {
    filename := os.Args[1]
    var config Config
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Value: %#v\n", config.Bar[0])
    fmt.Printf("Driver Value: %#v\n", config.Driver)
}

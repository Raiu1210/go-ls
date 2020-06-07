package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
)


func handleInput(args []string) {
    if len(args) > 1 {
        if args[1] == "-a" {
            load_all_file_list()
        }
    } else {
        load_file_list()
    }
}


func load_file_list() {
    // get current directory
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    fileInfoList, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, fileInfo := range fileInfoList {
        if strings.HasPrefix(fileInfo.Name(), ".") {
            continue
        }

        fmt.Println(fileInfo.Name())
    }
}


func load_all_file_list() {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    fileInfoList, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, fileInfo := range fileInfoList {
        fmt.Println(fileInfo.Name())
    }
}

func main() {
    handleInput(os.Args)
}
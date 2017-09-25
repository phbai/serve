package main

import (
    // "os"
    "log"
    "io/ioutil"
)

func GetList(path string) ([]FileInfo, error) {
    files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    var lists []FileInfo;


    for _, file := range files {
        fileInfo := FileInfo{
            Name: file.Name(),
            Size: file.Size(),
            IsDir: file.IsDir(),
        };
        lists = append(lists, fileInfo)
    }

    return lists, err;
}

func GetFile(path string) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("File contents: %s", content)
}
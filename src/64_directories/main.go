package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    err := os.Mkdir("subdir", 9755)
    check(err)
    defer os.RemoveAll("subdir")

    createEmptyFile := func(name string) {
        d := []byte("")
        // os.WriteFile ^1.16
        // check(os.WriteFile(name, d, 0644))
        check(ioutil.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // os.ReadDir ^1.16
    // c, err := os.ReadDir("subdir/parent")
    c, err := ioutil.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("subdir/parent/child")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("../../..")
    check(err)

    fmt.Println("Visiting subdir")
    err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}

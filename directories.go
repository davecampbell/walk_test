package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "io"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var cnt, max int

func main() {

//    err := os.Mkdir("subdir", 0755)
//    check(err)

//    defer os.RemoveAll("subdir")

//    createEmptyFile := func(name string) {
//        d := []byte("")
//        check(ioutil.WriteFile(name, d, 0644))
//    }

//    createEmptyFile("subdir/file1")

//    err = os.MkdirAll("subdir/parent/child", 0755)
//    check(err)

//    createEmptyFile("subdir/parent/file2")
//    createEmptyFile("subdir/parent/file3")
//    createEmptyFile("subdir/parent/child/file4")

//    c, err := ioutil.ReadDir("subdir/parent")
//    check(err)

//    fmt.Println("Listing subdir/parent")
//    for _, entry := range c {
//        fmt.Println(" ", entry.Name(), entry.IsDir())
//    }

//    err = os.Chdir("subdir/parent/child")
//    check(err)

//    c, err = ioutil.ReadDir(".")
//    check(err)

//    fmt.Println("Listing subdir/parent/child")
//    for _, entry := range c {
//        fmt.Println(" ", entry.Name(), entry.IsDir())
//    }

//    err = os.Chdir("../../..")
//    check(err)

//    fmt.Println("Visiting subdir")
//    err = filepath.Walk("subdir", visit)

    _, err := ioutil.ReadDir("/home")
    check(err)

    cnt=0
    max, err = strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("something amiss with args.")
    }

    var startpath string
    startpath = os.Args[1]

    err = os.Chdir(startpath)
    check(err)

    fmt.Println("Visiting...")
    err = filepath.Walk(".", visit)
    check(err)

}

func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    if cnt > max {
        fmt.Println(" -- hit the max -- ")
        return io.EOF
    }
    cnt++
    fmt.Println(" ", p, info.IsDir())

    if _, err := os.Stat(p); os.IsNotExist(err) {
	fmt.Println("not here.")
    }

    if _, err := os.Stat(p); !os.IsNotExist(err) {
	fmt.Println("here.")
    }

    return nil
}

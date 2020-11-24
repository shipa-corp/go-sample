package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Listening...")
    err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    if err != nil {
        panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    dat, err := ioutil.ReadFile("/vault/secrets/config")
    if err != nil {
        fmt.Print(err.Error())
        fmt.Fprintln(res,string(err.Error()))
    }
    fmt.Fprintln(res,string(dat))
    fmt.Fprintln(res, "Hello from Shipa!")
    dbName := fmt.Sprintf("DB_NAME: %s", os.Getenv("DB_NAME"))
    fmt.Fprintln(res, dbName) 
}

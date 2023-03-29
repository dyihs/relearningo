package main

import (
    "fmt"
    "io"
    "net/http"
)

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "read body failed: %v", err)
        return
    }
    fmt.Fprintf(w, "read the data: %s\n", string(body))

    // 尝试再次读取，啥也读不到，但是也不报错
    body, err = io.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "read the data one more time got error: %v", err)
        return
    }
    fmt.Fprintf(w, "read the data one more time: [%s] and read data", body)
}

func main() {
    http.HandleFunc("/body/once", readBodyOnce)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("readBodyOnce failed the error: %v", err)
    }
}

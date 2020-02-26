package main

import (
    "net/http"
    "log"
    "io"
)

func main() {
	http.HandleFunc("/", RocksServer)

    port := "8080"

    log.Println("Rodando na porta " + port)

    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}

func RocksServer(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, bold("Code.education Rocks!"))
}

func bold(text string) string {
    return "<b>" + text + "</b>"
}
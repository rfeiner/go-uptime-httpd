package main

import (
    "fmt"
    "net/http"
    "log"
    "flag"
    "os/exec"
)

var bindSocket = flag.String("socket", ":80", "Define socket where to listen for the HTTP requests")

func handler(w http.ResponseWriter, r *http.Request) {
    out, err := exec.Command("/usr/bin/uptime").Output()
    if err != nil {
        fmt.Fprintf(w, "Error, unable to retrieve uptime: %s", err)
        log.Fatal(err)
    }
    fmt.Fprintf(w, "%s", out)
}

func main() {
    flag.Parse()
    fmt.Printf("Starting server at %s\n", *bindSocket)
    
    http.HandleFunc("/", handler)
    http.ListenAndServe(*bindSocket, nil)
}
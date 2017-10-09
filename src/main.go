package main
import (
     "os"
     "log"
     "fmt"
     "net/http"
)
func handler (w http.ResponseWriter, r *http.Request) {
    log.Println("Received Request: ", r.Host)
    fmt.Fprintf(w, r.Host + "\n")
}
func main () {
    http.HandleFunc("/", handler)

    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT environment variable was not set")
    }
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("Could not listen: ", err)
    }
}

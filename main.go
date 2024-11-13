package main

import ( 
    "errors"
    "fmt"
    "io"
    "net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request\n")
    io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /hello request\n")
    io.WriteString(w, "Hello world!\n")
}

func main() {
    // http.HandleFunc("/",getRoot)
    http.HandleFunc("/", func(w  http.ResponseWriter, r *http.Request) {
        fmt.Printf("Request /, from IP: %s\n",r.RemoteAddr)
        http.ServeFile(w,r,"./static/index.html")
    })
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if 
    })

    err := http.ListenAndServe(":3333",nil)
    if errors.Is(err, http.ErrServerClosed)  {
        fmt.Printf("server closed\n")
    } else if  err != nil {
        fmt.Printf("error starting server: %s\n", err)
    }

}

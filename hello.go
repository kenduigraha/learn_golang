package main

import ("fmt"
        "html/template"
        "net/http"
        "time"
        "hash/crc32"
)

type World struct {}

func (w * World) String() string {
  return "World 2"
}

// type Office int

// const (
//   Boston Office = iota
//   NewYork
// )

// func (o Office) String() string {
//   return "Google, " + officePlace[o]
// }

func main() {
    fmt.Printf("Hello, %s\n", "World 1")

    fmt.Printf("Hello, %s\n", new(World))

    // fmt.Printf("Hello, %s\n", Boston)

    day := time.Now().Weekday()
    fmt.Printf("Hello, %s (%d)\n", day, day)

    start := time.Now()

    http.Get("http://www.google.com/")

    fmt.Println(time.Since(start))

    h := crc32.NewIEEE()

    fmt.Fprintf(h, "Hello Wordld\n")
    fmt.Println("hash=%#x\n", h.Sum32())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var data = map[string]string{
            "Name":    "john wick",
            "Message": "have a nice day",
        }

        var t, err = template.ParseFiles("template.html")
        if err != nil {
            fmt.Println(err.Error())
        }

        t.Execute(w, data)
    })

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

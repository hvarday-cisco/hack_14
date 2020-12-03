package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    responseText := `
        <html>
        <head>
        <title>Your Page Title</title>
        <meta name="google-site-verification" content="<VALUE>" />
        </head>
        </html>
`
    fmt.Fprint(w, responseText)
}

func Notification_handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var parsed_body interface{}

//     log.Printf("Data for org - %s\n\n", ps.ByName("org_pk"))
//     printExtras(r)
    fmt.Printf("\n---------- ORG_PK ----- (%s)\n" , ps.ByName("org_pk") )
    fmt.Print("-------- EVENT DATA -----------------\n")

    d := json.NewDecoder(r.Body)
    err := d.Decode(&parsed_body)

    json_res, err := json.MarshalIndent(parsed_body, "", "  ")
    fmt.Println(string(json_res))

    fmt.Printf("\n.......................\n")
    fmt.Printf("\nSending payload to queuing service to further process by connector.......\n")
    fmt.Printf("\n.......................\n")

    w.Header().Set("Content-Type", "application/json")
    STATUS_OK_JSON, err := json.Marshal(map[string]string{
        "status": "success",
    })
    if err != nil{
        panic(err)
    }

    w.Write(STATUS_OK_JSON)
//     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("org_pk"))
}

func main() {

    router := httprouter.New()

    router.GET("/", Index)
    router.POST("/notifications/:org_pk", Notification_handler)
    router.GET("/notifications/:org_pk", Notification_handler)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func printExtras(r *http.Request){

    var parsed_body interface{}

    // Print Headers
    log.Println("Request Headers - \n")
    for name, headers := range r.Header {
        for _, h := range headers {
            fmt.Println(name, h)
        }
    }

    log.Println("\n##################\n\nRequest Query Params - \n")
    // Query Params
    queryValues := r.URL.Query()
    fmt.Println(queryValues)

    fmt.Println("\n##################\n\nRequest Body - \n")
    // Print Body
    d := json.NewDecoder(r.Body)
    err := d.Decode(&parsed_body)

    json_res, err := json.MarshalIndent(parsed_body, "", "  ")
//     fmt.Println(string(json_res), err)
    fmt.Println(string(json_res))

    if err != nil {
        if err != io.EOF{
            log.Println("Error " + string(err.Error()))
        }
    } else {
        fmt.Println(parsed_body)
    }
    fmt.Println("\n##################\n")

}
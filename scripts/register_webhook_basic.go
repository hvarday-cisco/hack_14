package main

import (
    "fmt"
    "hack/config"
    "encoding/json"
    "time"
    "net/http"
    "bytes"
    "strconv"
    "io/ioutil"
    "github.com/imroc/req"
)

func get_channel_id() string{
    return strconv.Itoa(int(time.Now().Unix()))
}

func main(){
    headers := map[string]string{
        "Authorization": conf.TOKEN,
        "Content-Type": "application/json",
    }
    body := map[string]string{
        "kind": "api#channel",
        "id": get_channel_id(),
        "type": "web_hook",
        "address": conf.HOST + "/notifications/",
        "token":"target=myApp-AdminActivities",
    }

    register(headers, body)

}

func register(headers map[string]string, body map[string]string){

    client := &http.Client{}
    body_json, err := json.Marshal(body)

    req, err := http.NewRequest("POST", conf.WATCH_ENDPOINT, bytes.NewReader(body_json))
    for key, value:= range headers{
        req.Header.Add(key, value)
    }

    if err != nil{
        fmt.Println("Error" + string(err.Error()))
    } else {

        // Make Request
        res, err := client.Do(req)
        if err != nil{
            fmt.Println("ERROR", err.Error())
        } else {
            body, err := ioutil.ReadAll(res.Body)
            if err != nil {
                panic(err)
            }
            fmt.Println(string(body))
        }

    }

}
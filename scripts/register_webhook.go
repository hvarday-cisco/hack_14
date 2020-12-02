package main

import (
    "fmt"
    "hack/config"
    "time"
    "strconv"
    "github.com/imroc/req"
)

func get_channel_id() string{
    return strconv.Itoa(int(time.Now().Unix()))
}

func main(){

    body := map[string]string{
        "kind": "api#channel",
        "id": get_channel_id(),
        "type": "web_hook",
        "address": conf.HOST + "/notifications/",
        "token":"target=myApp-AdminActivities",
    }

    req_headers := req.Header{
        "Authorization": conf.TOKEN,
        "Content-Type": "application/json",
    }

    res, err := req.Post(conf.WATCH_ENDPOINT, req_headers, req.BodyJSON(&body))

    if err != nil{
        fmt.Println("Error - ", err.Error())
    } else {
        fmt.Println(res.String())
    }
}
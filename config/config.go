// Config
package conf

// type configuration struct {
//     TOKEN string
//     HOST string
//     WATCH_ENDPOINT string
//     STOP_ENDPOINT string
// }
//
// var Config = configuration{
//     TOKEN : "",
//     HOST : "https://d359d121f97e.ngrok.io",
//     WATCH_ENDPOINT : "https://www.googleapis.com/admin/reports/v1/activity/users/all/applications/admin/watch",
//     STOP_ENDPOINT : "https://www.googleapis.com/admin/reports_v1/channels/stop",
// }

var TOKEN = "DUMMY_TOKEN";
var HOST = "https://d359d121f97e.ngrok.io";
// var WATCH_ENDPOINT = "https://www.googleapis.com/admin/reports/v1/activity/users/all/applications/admin/watch";
var WATCH_ENDPOINT = "http://localhost:8080";
var STOP_ENDPOINT = "https://www.googleapis.com/admin/reports_v1/channels/stop";

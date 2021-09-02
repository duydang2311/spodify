package main

import ( 
    "bytes"
    "html/template"
    "net/http"
    "fmt"
    "net/url"
    "encoding/json"
)

const spotify_auth string = "https://accounts.spotify.com/authorize?response_type=code";
const client_id string = "client_id=";
const client_id_value string = "xxx";
const redirect_uri string = "redirect_uri=";
const redirect_uri_value string = "http://localhost:8080/";
var tpl *template.Template

func main() {
    tpl = template.Must(template.ParseGlob("public/*.html"))
    http.HandleFunc("/", index)
    http.Handle("/fs/", http.StripPrefix("/fs/", http.FileServer(http.Dir("./public"))))
    http.HandleFunc("/oauth", onOAuth);
    http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "index.html", nil)
    var code = r.URL.Query().Get("code")
    if(len(code) > 0) {
        fmt.Fprintf(w, r.URL.Query().Encode())
        body, err := json.Marshal(map[string]string{
            "grant_type": "authorization_code",
            "code": code,
            "redirect_uri": url.QueryEscape(redirect_uri_value),
        })
        if err != nil {
        }
        http.Post("https://accounts.spotify.com/api/token?", "application/json", bytes.NewBuffer(body))
    }
    http.Get(spotify_auth + "?" + client_id + client_id_value + "&" + redirect_uri + redirect_uri_value)
}

func onOAuth(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "")
    http.Redirect(w, r, spotify_auth + "?" + client_id + client_id_value + "&" + redirect_uri + url.QueryEscape(redirect_uri_value), 302)
}

package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "tritium_oss/tr"
)

type Page struct {
    Title string
    Body []byte
}

func request(path string) (*Page, error){
    target  := "http://" + path
    u, _    := url.ParseRequestURI(target)
    urlStr  := fmt.Sprintf("%v", u)
    response, err  := http.Get(urlStr)
    if err != nil {
        e_msg := fmt.Sprintf("%v", err)
        return &Page{Title: urlStr, Body: []byte(e_msg)}, nil
    } else {
        body, _  := ioutil.ReadAll(response.Body)
        return &Page{Title: urlStr, Body: body}, nil
    }
}

func handler(w http.ResponseWriter, r *http.Request){
    output := ""
    if r.URL.Path[1:] == "" {
        output     = "t.Engine running. Pass any url through the t.Engine path to transform the domain in real time."
    } else {
        tscript, _ := ioutil.ReadFile("main.ts")
        html, _    := request(r.URL.Path[1:])
        output     = tritium.Transform(string(tscript), string(html.Body))
    }

    fmt.Fprintf(w, "%s", output)
}

func main(){
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":3030", nil)
    if err != nil {
        fmt.Fprintf("I've made a huge mistake...")
    }
}
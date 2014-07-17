package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "tritium_oss/driver"
)

type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error{
    filename := p.Title + ".txt"

    return ioutil.WriteFile(filename, p.Body, 0600)
}

func request(path string) (*Page, error){
    target  := "https://" + path
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
    tscript, _ := ioutil.ReadFile("main.ts")
    html, _    := request(r.URL.Path[1:])
    output     := driver.Transform(string(tscript), string(html.Body))
    fmt.Fprintf(w, "%s", output)
}

func main(){
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3030", nil)
}
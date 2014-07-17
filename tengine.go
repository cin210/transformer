package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
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

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil{
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request){
    tscript, _ := ioutil.ReadFile("main.ts")
    html, _ := loadPage("tEngine")
    output := driver.Transform(string(tscript), string(html.Body))
    // fmt.Fprintf(w, "URL received: %s", r.URL.Path[1:])
    fmt.Fprintf(w, "%s", output)
}

func main(){
  p1 := &Page{Title: "tEngine", Body: []byte("All t.engines running!")}
  p1.save()
  http.HandleFunc("/", handler)
  http.HandleFunc("/json", jsonHandler)
  http.ListenAndServe(":3030", nil)
}
package main

// This webserver will accept requests on port 3030
// Visiting any subdomain will return a transformed 
// html response using the Open Tritium driver.

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "open-tritium/tr"
)

type Page struct {
    Title string
    Body []byte
}

// A request is made to the web server in the form of HOSTNAME:3030/my.cool.site
// The request is made on behalf of the client, and the response is returned as the Body []byte

func request(path string) (*Page, error){
    // For now, pass through all subdomains as insecure http connections
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

// The handler will attempt to make a connection on behalf of the user
// tritium.Transform will apply the tscript "main.ts" to the response
// This will parse the body of the response and apply any and all transformations 
// specified. In this naive example, the src of all images found in the body
// are replaced with a random cat picture, and the 'moov' class is added to the body tag.
// See developer.moovweb.com and tritium.io for more applications of Open Tritium

func handler(w http.ResponseWriter, r *http.Request){
    output := ""
    if r.URL.Path[1:] == "" {
        output     = "Server running. Visit /your.domain.here to transform your.domain.here in real time"
    } else {
	    // load the tscript used by this program.
        tscript, _ := ioutil.ReadFile("main.ts")
        html, _    := request(r.URL.Path[1:])
		// Invoke a tritium transformation
        output     = tritium.Transform(string(tscript), string(html.Body), "")
    }

    fmt.Fprintf(w, "%s", output)
}

func main(){
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":3030", nil)
    if err != nil {
        fmt.Printf("I've made a huge mistake...")
    }
}

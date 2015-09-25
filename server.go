package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    log "github.com/Sirupsen/logrus"
    "os"
    "fmt"
)

var pageFolderLoc = "res/content/"
var portNumber = ":80"

type PageData struct {
    Header string
    Para1 string
    Para2 string
}

func main() {
    initLogging()

    http.HandleFunc("/pages/", pagesHandler)
    log.Info("Listening on port"+portNumber)
    log.Fatal(http.ListenAndServe(portNumber, nil))
}

func initLogging(){
    log.SetOutput(os.Stdout)
    log.SetLevel(log.DebugLevel)
}

func pagesHandler(w http.ResponseWriter, r *http.Request) {
    //Returning content by reading the file

    //Get the url:
    dataFileLoc := r.URL.Path[len("/pages/"):] + ".json"

    log.Info("Attempting to get file: "+pageFolderLoc + dataFileLoc)
    dat, err := ioutil.ReadFile(pageFolderLoc + dataFileLoc)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    var myPageData PageData
    if err = json.Unmarshal(dat, &myPageData); err != nil {
        log.Fatal(err.Error())
    }
    log.Debug("Header?: "+myPageData.Header)

    htmlString := makeHTML(myPageData)

    fmt.Fprintf(w, htmlString)
}

func makeHTML(myPageData PageData) string {
    return fmt.Sprintf("<html>" +
        "<head></head>"+
        "<body>"+
            "<h1>%v</h1>"+
            "<p>%v</p>"+
            "<p>%v</p>"+
        "</body>"+
        "</html>", myPageData.Header, myPageData.Para1, myPageData.Para2)
}

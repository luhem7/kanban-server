package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    log "github.com/Sirupsen/logrus"
    "os"
    "fmt"
    "html/template"
    "bytes"
)

//Location of content data
var contentFolderLoc = "res/content/"
//Location of the template file
var templateFileLoc = "res/templates/"
//Name of the template file
var templateFileName = "board.tmpl"
//Port number to start the server on
var portNumber = ":80"


type PageData struct {
    Title string
    Header string
    Description string
    Reasons []string
}

func main() {
    initLogging()

    http.HandleFunc("/content/", contentHandler)
    http.HandleFunc("/kanban-board/", kanbanHandler)
    log.Info("Listening on port"+portNumber)
    log.Fatal(http.ListenAndServe(portNumber, nil))
}

func initLogging(){
    log.SetOutput(os.Stdout)
    log.SetLevel(log.DebugLevel)
}

//This function handles calls related to manipulating the kanban board itself
func kanbanHandler(w http.ResponseWriter, r *http.Request){
    //TODO Add stuff here
}

//Returns content based on the input url.
//Format of url: /content/<contentName>
//contentName pulls from data from a file named contentName.json
func contentHandler(w http.ResponseWriter, r *http.Request) {

    //Get the filename from the url:
    dataFileLoc := r.URL.Path[len("/content/"):] + ".json"

    log.Info("Request for file: "+contentFolderLoc + dataFileLoc)
    dat, err := ioutil.ReadFile(contentFolderLoc + dataFileLoc)
    if err != nil {
        //Return a 404 for errrors. TODO Use actual HTTP responses
        http.NotFound(w, r)
        log.Error(err.Error())
        return
    }

    var myPageData PageData
    if err = json.Unmarshal(dat, &myPageData); err != nil {
        http.Error(w, "Error processing page", 500)
        log.Error(err.Error())
        return
    }

    htmlString, err := makeHTML(myPageData)
    if err != nil {
        http.Error(w, "Error processing page", 500)
        log.Error(err.Error())
        return
    }

    fmt.Fprintf(w, htmlString)
}

func makeHTML(myPageData PageData) (string, error) {
    //result := bytes.NewBufferString("")
    var result bytes.Buffer

    //Read the template file
    t, err := template.ParseFiles(templateFileLoc+templateFileName)
    /*Things I learnt while trying to debug the above line:
        1. Use template.ParseFiles, NOT NOT NOT (t *Template).ParseFiles, where t is
            some template that was defined earlier in the code.
        2. The name of the generated template will be the base name of the first
            file passed to ParseFiles method.
    */

    if err != nil {
        log.Error(err.Error())
        return "", err
    }

    err = t.Execute(&result, myPageData)
    /*Things I learnt while trying to get the above line to work:
        The parameters for the Execute Execute function are
        (wr io.Writer, data interface{}). I didnt want to pass in the http.ResponseWriter
        directly to the function in case there was an error while processing the template.
        If found out that instead of passing in things like os.Stdout or ResponseWriter,
        I can create a new bytes.Buffer buffer and use it instead. This way, if
        there are issues while processing the buffer, I can easily discard the contents of
        buffer with out worrying about clearing a stream.
        The type (*bytes.Buffer) implements io.Writer.
        Two ways of doing this:
            1. Make a new variable of type bytes.buffer: var result bytes.Buffer.
                Call Execute by passing in a pointer to this variable: &result
            2. Use the bytes.NewBufferString("") function, this will return a pointer
                to a new buffer. This pointer can be passed into the Execute function
    */

    return result.String(), nil
}

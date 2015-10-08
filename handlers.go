package main

import (
    "net/http"
    log "github.com/Sirupsen/logrus"
    "github.com/gorilla/mux"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

//This function handles calls related to accessing binary files
func binaryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    binaryFilePath := vars["binaryFilePath"]
	//binaryFilePath := r.URL.Path[len("/binary/"):]

	log.Info("Request for binary: " + binaryFilePath)

	dat, err := ioutil.ReadFile(binaryFolderLoc + binaryFilePath)
	if err != nil {
		//Return a 404 for errrors.
		http.NotFound(w, r)
        log.Error("Error serving Binary file "+r.URL.Path[len("/binary/"):])
		log.Error(err.Error())
		return
	}

	fmt.Fprintf(w, string(dat))
}

//This function handles calls related to accessing the kanban board itself
func kanbanHandler(w http.ResponseWriter, r *http.Request) {
	//If attempting to go to a child url of /kanban-board/, put them back at
	//kanban-board
	if len(r.URL.Path[len("/kanban-board/"):]) > 0 {
		http.Redirect(w, r, "/kanban-board/", 303)
	}

	//Read the kanban board html file
	dat, err := ioutil.ReadFile(binaryFolderLoc + kanbanHtmlFileName)
	if err != nil {
		http.Error(w, "Error processing page", 500)
		log.Error(err.Error())
		return
	}

	log.Info("Serving up kanban board")
	fmt.Fprintf(w, string(dat))
}

//Returns content based on the input url.
//Format of url: /content/<contentName>
//contentName pulls from data from a file named contentName.json
func contentHandler(w http.ResponseWriter, r *http.Request) {

	//Get the filename from the url:
	dataFileLoc := r.URL.Path[len("/kanban-board/content/"):] + ".json"

	log.Info("Request for file: " + contentFolderLoc + dataFileLoc)
	dat, err := ioutil.ReadFile(contentFolderLoc + dataFileLoc)
	if err != nil {
		//Return a 404 for errrors.
		http.NotFound(w, r)
		log.Error(err.Error())
		return
	}

	var myPageData PageDataModel
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

//Handles calls related to the handle function
func dataColumnsHandler(w http.ResponseWriter, r *http.Request){

}

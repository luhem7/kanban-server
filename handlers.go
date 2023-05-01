package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)

//This function handles calls related to accessing binary files
func binaryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	binaryFilePath := vars["binaryFilePath"]

	log.Info("Request for binary: " + binaryFilePath)

	dat, err := ioutil.ReadFile(binaryFolderLoc + binaryFilePath)
	if err != nil {
		//Return a 404 for errrors.
		http.NotFound(w, r)
		log.Error("Error serving Binary file " + r.URL.Path[len("/binary/"):])
		log.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", getMimeType(binaryFilePath))
	fmt.Fprintf(w, string(dat))
}

//Given a filename or a filepath, tries to find the corresponding MIME type
//Returns application/octet-stream as default if none is found
func getMimeType(filePath string) string {
	fileExt := filepath.Ext(filePath)
	defaultMime := "application/octet-stream"
	if len(fileExt) != 0 {
		mimeType := mime.TypeByExtension(fileExt)
		if len(mimeType) != 0 {
			return mimeType
		}
	}
	return defaultMime
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
func dataGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceName := vars["resourceName"]
	log.Debug("In dataHandler. Request for " + resourceName)
	dat, err := ioutil.ReadFile(dataDir + resourceName + ".json")
	if err != nil {
		http.Error(w, "Error fetching data for resource "+resourceName, 404)
		log.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(dat))
}

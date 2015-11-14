package main

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func main() {
	initLogging()

	router := mux.NewRouter().StrictSlash(true)
	//Presentation related Handlers
	router.HandleFunc("/binary/{binaryFilePath:.+}", binaryHandler).Methods("GET")
	router.HandleFunc("/kanban-board/content/{resourceName:.+}", contentHandler).Methods("GET")
	router.HandleFunc("/kanban-board/", kanbanHandler).Methods("GET")

	//Data related Handlers
	router.HandleFunc("/kanban-board/data/{resourceName:.+}", dataGetHandler).Methods("GET")

	log.Info("Listening on port" + portNumber)
	log.Fatal(http.ListenAndServe(portNumber, router))
}

func initLogging() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func makeHTML(myPageData PageDataModel) (string, error) {
	//result := bytes.NewBufferString("")
	var result bytes.Buffer

	//Read the template file
	t, err := template.ParseFiles(templateFolderLoc + myPageData.Templates[0])
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

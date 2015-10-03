package main

type Settings struct {
    ContentFolderLocation string
    TemplateFileLoc string
    TemplateFileName string
    PortNumber string
}

var DefaultSettings = Settings{
    ContentFolderLocation : "res/content",
}

//Location of content data
var contentFolderLoc = "res/content/"

//Location of the template file
var templateFileLoc = "res/templates/"

//Name of the template file
//var templateFileName = "About.tmpl"

//Port number to start the server on
var portNumber = ":80"

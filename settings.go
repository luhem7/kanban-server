package main

type Settings struct {
	ContentFolderLocation string
	TemplateFileLoc       string
	TemplateFileName      string
	PortNumber            string
}

var DefaultSettings = Settings{
	ContentFolderLocation: "res/content/",
}

//Location of content data
var contentFolderLoc = "res/content/"

//Location of binary files
var binaryFolderLoc = "res/binary/"

//Location of the template file
var templateFolderLoc = "res/templates/"

//Location of the kanban-board html file
var kanbanHtmlFileName = "/html/mainBoardScreen.html"

//Location of the data directory where board items are stored
var dataDir = "data/"

//filename of the columns json setting:
var columnsFileName = "columns.json"

//Port number to start the server on
var portNumber = ":80"

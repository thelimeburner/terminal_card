package main

import (
	"io"
	"log"
)

//vars for logging
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

//sets up the loggin functions
func logInit(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//prints error message and returns
func errorMsg(err error, msg ...interface{}) {

	//Error.Println(msg)
	log.Println(msg)
	Error.Println(err)

}

//prints info message and returns
func infoMsg(msg ...interface{}) {

	Info.Println(msg)

}

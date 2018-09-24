package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/Max2Inc/vh-cc-golang-card/ccc"
)

/* VARIABLES */
//flag for address to serve content
var addr = flag.String("addr", "localhost:8000", "http service address")
var debugPTR = flag.Bool("debug", false, "Enable debug print statements")
var debug bool
var card *ccc.Card

//get audio manager address and other variables
//var env ENV

//used for local testing enablement
var localTest = false
var localTest2 = false

//main instance
func main() {
	//parse flag commands
	flag.Parse()

	//debug mode enabled?
	debug = *debugPTR

	//setup logging (See logging.go)
	logInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	//sets the environment (See test.go)
	if localTest2 {
		//SetENV()
	}

	//get environment variables for running card (See utils.go)
	getEnviroment()

	//initialize handlers
	handlers := ccc.CardHandlers{}

	//register message handler (See handlers.go)
	handlers.Message = listener

	//register js and html handlers
	handlers.HTML = html
	handlers.Js = js
	handlers.Open = openHandler
	handlers.Close = closeHandler

	//instantiate new card by passing in the handlers
	card = ccc.New(handlers)
	card.Log(true)

	//connect to the card
	card.Connect()

	//registerImages()
	//run audio manager
	//go am.Run()

	//loop forever to wait for updates
	card.Listen()
}

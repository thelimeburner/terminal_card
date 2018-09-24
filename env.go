package main

import (
	"os"
)

// type ENV struct {
// 	AUDIO_MANAGER  string
// 	BUFFER         string
// 	AUDIO_LOCATION string
// }

//fetches environment variable
func getEnviroment() {
	//var env = ENV{}
	//var err error

	infoMsg("CC_SECRET=", os.Getenv("CC_SECRET"))
	infoMsg("CC_KEY=", os.Getenv("CC_KEY"))
	infoMsg("CC_ADDRESS=", os.Getenv("CC_ADDRESS"))
	infoMsg("CC_PORT=", os.Getenv("CC_PORT"))

	//return env
}

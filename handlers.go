/*

Sets the handlers for card functionality.

*/
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/Max2Inc/vh-cc-golang-card/ccc"
)

type payload map[string]interface{}

//listener listens for messages and handles them
func listener(msg ccc.Msg, cm ccc.CardMessage) error {

	switch cm["command"] {
	case "run_cmd":
		infoMsg("Received command request")
		r := handleCMD(cm)
		card.Reply(msg, r)
		return nil

	default:
		errorMsg(errors.New("Unrecognizable Message"), "listener")
		infoMsg(msg)
		return errors.New("Unrecognizable Message")
	}
	//card.Reply(msg, cm)
	//return nil
}

//OpenHandler  open handler used for new connections
//send clientID with response
func openHandler(cid string, uid string, msg ccc.Msg, c *ccc.Card) {
	//infoMsg("Ignoring Open Messages")
	infoMsg("New Client ==> cid:" + cid + " uid:" + uid)

	return

}

func closeHandler(cid string, uid string, msg ccc.Msg, c *ccc.Card) {
	//infoMsg("Ignoring Open Messages")
	infoMsg("Close Request for Client ==> cid:" + cid + " uid:" + uid)

	//infoMsg(c.Handlers)
	return
}

func html(cid string, uid string) []byte {
	//infoMsg("Retrieve initiated!")
	// b, err := ioutil.ReadFile("static/player_icons.html")
	b, err := ioutil.ReadFile("card/card.html")
	if err != nil {
		var empty = make([]byte, 1)
		return empty
	}
	return b
}

func js(cid string, uid string) []byte {
	// b, err := ioutil.ReadFile("static/player.js")
	b, err := ioutil.ReadFile("card/card.js")
	if err != nil {

		var empty = make([]byte, 1)
		return empty
	}
	return b
}

func handleCMD(cm ccc.CardMessage) payload {

	//unwrap params from message
	params, err := ccc.UnwrapParameters(cm)
	if err != nil {
		errorMsg(errors.New("failed to unwrap parameters from message. check message format "), "handleHistory")
		return genErrorReply(cm["command"].(string), "Failed to unwrap payload from message. Check message format")
	}
	//infoMsg("Parameters is:")
	//fmt.Println(params)

	//validate message is as expected
	valid := ccc.CheckVal(params, []string{"cmd", "args"})
	if !valid {
		errorMsg(errors.New("failed to validate parameters from message. check message format "), "handleHistory")
		return genErrorReply("cmd_results", "Failed to validate payload from message. Check message format")
	}

	cmd_msg := params["cmd"].(string)
	argInterface := params["args"].([]interface{})
	//args := make([]string, 0)
	infoMsg(params["args"])

	args := make([]string, len(argInterface))
	for i, v := range argInterface {
		args[i] = fmt.Sprint(v)
	}
	if len(args) > 1 && args[len(args)-1] == "" {
		args = args[:len(args)-1]
	}
	fmt.Println("args: ", args)
	// for _, k := range params["args"] {
	// 	args = append(args, k.(string))
	// }
	clientCmd := cmd{}
	clientCmd.cmd = exec.Command(cmd_msg, args...)
	//clientCmd.cmd.Args = args
	err = clientCmd.runCommand()
	if err != nil {
		return genErrorReply("cmd_results", "Failed to run command")

	}

	var parameters2 map[string]interface{}

	parameters2 = map[string]interface{}{
		"success": true,
		"output":  clientCmd.output,
	}
	//args := params["name"].(string)

	results := map[string]interface{}{
		"command":    "cmd_results",
		"parameters": parameters2,
	}
	return results
}

func genErrorReply(cmd string, msg string) payload {

	var result payload
	parameters2 := map[string]interface{}{
		"success": false,
		"msg":     msg,
	}
	result = map[string]interface{}{
		"command":    cmd,
		"parameters": parameters2,
	}
	return result
}

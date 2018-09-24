// First, we need some callback functions.
 
function initCardCallback(card, socket) {
    // This function will be called after the card is ready to go.
    // Result contains the jQuery selector for the card and the
    // WebSocket you can use to communicate with the app.
    $(card).on("click", ".run_cmd", function(e) {
        e.preventDefault();
        cmdString = $(card).find(".terminal_input").val();
        console.log(cmdString)
        //alert(cmd)
        parts = cmdString.split(" ");
        if( parts.length <= 0){
            alert("Need to specify one command!")
            return
        }
        cmd = parts[0]
        args = []
        if (parts.length > 1 ){
            args = parts.slice(1, parts.length)
        }
        console.log("Command:",cmd )
        console.log("Args:",args )

        socket.send("run_cmd", { "cmd": cmd,"args":args }, function(result) {
        //alert("Running command:",cmd)
            console.log("Message sent")
            console.log(result.parameters.output)
            //console.log(event.data.result)
            $(card).find(".terminal_output").text(result.parameters.output);

            
        })
    });
}
 
function eventCallback(card, event, socket) {
    // This function will be called whenever the Control Center
    // receives a CC_CARD_UPDATE package from the node.
    // Contents of the "event" object is up to the app developer.
    if (event.command == "cmd_results") {
        console.log(event)
    }
}
 
// Then, we need to attach ourselves to the Control Center.
// We do this by using the global vhcc (VeeaHubControlCenter) variable.
vhcc.cards.attach(initCardCallback, eventCallback);
 
// We are done.
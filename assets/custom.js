var input = document.getElementById("input");
var output = document.getElementById("output");
var socket = new WebSocket("ws://localhost:7100/ws");

socket.onopen = function () {
    console.log("Status: Connected");
    // output.innerHTML += "Status: Connected\n";
};

socket.onmessage = function (e) {
    console.log(e.data);
    output.innerHTML += "\nServer: " + e.data + "\n";
};

function send() {
    socket.send(input.value);
    // input.value = "";
}

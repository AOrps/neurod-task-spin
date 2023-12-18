// var input = document.querySelector('.input');
var output = document.querySelector('.output');
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
    let input = document.querySelector('.input');
    console.log("yes");
    console.log(input);
    socket.send(input.value);
    // input.value = "";
}

document.addEventListener('DOMContentLoaded', function() {
    console.log("ononon")
});

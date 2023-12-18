function theButton(socket) {
    let button = document.getElementById('daButton');

    button.addEventListener('click', function () {
        var input = document.getElementById('input');

        socket.onmessage = function (e) {
            const output = document.getElementById('output');
            output.textContent = input.value;
        };
        socket.send(input.value);        
    });
}


document.addEventListener('DOMContentLoaded', function() {
    console.log("ononon");

    let socket = new WebSocket("ws://localhost:7100/ws");

    socket.onopen = function () {
        console.log("Websocket Javascript Status: Connected");
    };
    
    theButton(socket);
    
});

// function send() {
//     let input = document.querySelector('.input');
//     console.log("yes");
//     console.log(input);
//     socket.send(input.value);
//     // input.value = "";
// }
function theButton(socket) {
    let button = document.getElementById('daButton');

    button.addEventListener('click', function () {

        // let input = document.querySelector('.input');
        // var input = document.getElementById('input');
        let output = document.querySelector('.output');

        var input = document.getElementById('input');
        
        console.log("yo: ", input.value);

        socket.onmessage = function (e) {
            console.log(e.data);
            input = e.data;
            // output.innerHTML += "\nServer: " + e.data + "\n";
        };
        
        console.log("yo2: ", input.value);


        socket.send(input.value);        
    });
}


document.addEventListener('DOMContentLoaded', function() {
    console.log("ononon");

    let socket = new WebSocket("ws://localhost:7100/ws");

    socket.onopen = function () {
        console.log("Status: Connected");
        // output.innerHTML += "Status: Connected\n";
    };
    
    theButton(socket);
    
});

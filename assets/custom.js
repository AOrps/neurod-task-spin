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
        // var output = document.querySelector('.output');

        var input = document.getElementById('input');

        
        // console.log("yo: ", input.value);

        socket.onmessage = function (e) {
             // input and e.data are the same
            console.log("yes:", e.data);
            const output = document.getElementById('output');
            // input = e.data;
            // output.innerHTML += "\nServer: " + e.data + "\n";
            output.textContent = input.value;
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

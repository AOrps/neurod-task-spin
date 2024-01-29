// wheel initialization
window.onload = () => {
    const props = {items: []};    
    const container = document.querySelector('.wheel-wrapper');
    const wheel = new spinWheel.Wheel(container, props);
    wheel.onCurrentIndexChange = (e) => {
        const paragraph = document.querySelector('.wheel-value');
        console.log("asdfasdf")
        paragraph.textContent = props["items"][e.currentIndex].label;
    }
    window.wheel = wheel;
}

// theButton: send info via the socket
function theButton(socket) {
    let button = document.getElementById('daButton');

    // when the button is clicked, send content to the golang websocket
    button.addEventListener('click', function () {
        var input = document.getElementById('input');
        socket.send(input.value);        
    });
}

// then DOMContentLoaded, connect to websocket and do magic
document.addEventListener('DOMContentLoaded', function() {
    let socket = new WebSocket("wss://localhost:7100/ws");

    socket.onopen = function () {
        console.log("Websocket Javascript Status: Connected");
    };
    
    theButton(socket);

    socket.onmessage = function (e) {
        var nData = JSON.parse(e.data);
        const output = document.getElementById('output');
        output.textContent = nData;

        const nContainer = document.querySelector('.wheel-wrapper');

        var oldCanvas = nContainer.querySelector('canvas');
        // removes old wheel
        oldCanvas.remove(); // NEEDED, otherwise the new wheel goes nuts
        
        var nProps = {items: nData};        
        var uWheel = new spinWheel.Wheel(nContainer,nProps);

        uWheel.onCurrentIndexChange = (event) => {
            const nParagraph = document.querySelector('.wheel-value');
            const word = nProps["items"][event.currentIndex].label;
            nParagraph.textContent = word;
        }
        window.wheel = uWheel;
    }
});

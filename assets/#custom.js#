window.onload = () => {
  const props = {
    items: [
      {
        label: 'one',
        backgroundColor: '#fff',
      },
      {
        label: 'two',
        backgroundColor: '#eee',
      },
      {
        label: 'three',
        backgroundColor: '#ddd',
      },
      {
        label: 'four',
        backgroundColor: '#ccc',
      },
      {
        label: 'five',
        backgroundColor: '#bbb',
      },          
    ]};
  const container = document.querySelector('.wheel-wrapper');
  const wheel = new spinWheel.Wheel(container, props);

  wheel.onCurrentIndexChange = (e) => {
      const paragraph = document.querySelector('.wheel-value');
      paragraph.textContent = props["items"][e.currentIndex].label;
  }

  window.wheel = wheel;
}

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

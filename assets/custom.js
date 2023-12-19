window.onload = () => {
  const props = {
    items: []};
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
            // console.log(e);
            console.log(e.data);
            const output = document.getElementById('output');
            // output.textContent = input.value;
            output.textContent = e.data;
            
            const props = {items: e.data};
            const container = document.querySelector('.wheel-wrapper');
            window.wheel = new spinWheel.Wheel(container, props);
            // const container = document.querySelector('.wheel-wrapper');
            // const wheel = new spinWheel.Wheel(container, props);

            wheel.onCurrentIndexChange = (e) => {
                const paragraph = document.querySelector('.wheel-value');
                paragraph.textContent = props["items"][e.currentIndex].label;
            }

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

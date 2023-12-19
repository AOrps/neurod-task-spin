window.onload = () => {
  const props = {items: [{label:"yes", backgroundColor: '#fff'},{label:"no", backgroundColor:'#eee'}]};
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
    let tryWheel;
    

    
    button.addEventListener('click', function () {
        var input = document.getElementById('input');


        socket.onmessage = function (e) {
            console.log(e.data);
            const output = document.getElementById('output');

            output.textContent = e.data;


            const newData = JSON.parse(e.data);
            console.log("newData", newData)

            // window.wheel.
            
            var updateWheel = new spinWheel.Wheel(document.querySelector('.wheel-wrapper'), { item: newData.items });
            // tryWheel = new spinWheel.Wheel(document.querySelector('.wheel-wrapper'), {items: [{label: "sey"}, {label: "on"}]});
            // window.wheel = tryWheel;
            wheel.innerHTML = '';

              updateWheel.onCurrentIndexChange = (e) => {
                  const paragraph = document.querySelector('.wheel-value');
                  paragraph.textContent = props["items"][e.currentIndex].label;
              }


                console.log("updated wheel:", updateWheel);
                window.wheel = updateWheel;

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

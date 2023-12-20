window.onload = () => {
    const props = {items: [{label:"yes", backgroundColor: '#fff'},{label:"no", backgroundColor:'#eee'}]};
    // const props = {items: []};    
    const container = document.querySelector('.wheel-wrapper');
    const wheel = new spinWheel.Wheel(container, props);

    console.log("wheel",wheel, typeof(wheel));
    console.log("props",props, typeof(props));

    wheel.onCurrentIndexChange = (e) => {
        const paragraph = document.querySelector('.wheel-value');
        console.log("asdfasdf")
        paragraph.textContent = props["items"][e.currentIndex].label;
    }

    window.wheel = wheel;
}

function theButton(socket) {
    let button = document.getElementById('daButton');
    
    button.addEventListener('click', function () {
        var input = document.getElementById('input');
        socket.send(input.value);        
    });

     // socket.onmessage = function (e) {
     //        const newData = JSON.parse(e.data)
     //        const output = document.getElementById('output');

     //        output.textContent = newData;
     //        console.log("newData", newData)

     //        var updateWheel = new spinWheel.Wheel(document.querySelector('.wheel-wrapper'), { item: newData.items });
     //        console.log("target",document.querySelector('.wheel-wrapper'))
     //            // ._items[0]._backgroundColor: "#dedede";
            

     //        console.log("updated wheel:", updateWheel);
     //        wheel = updateWheel;

     //    };

}



document.addEventListener('DOMContentLoaded', function() {
    let socket = new WebSocket("ws://localhost:7100/ws");

    socket.onopen = function () {
        console.log("Websocket Javascript Status: Connected");
    };
    
    theButton(socket);

    socket.onmessage = function (e) {
        var nData = JSON.parse(e.data);
        // console.log("socket onmessage (e.data):", nData, typeof(nData))
        const output = document.getElementById('output');
        output.textContent = nData;

        const nContainer = document.querySelector('.wheel-wrapper');
        // nContainer.removeChild();
        var oldCanvas = nContainer.querySelector('canvas');
        console.log(oldCanvas);
        oldCanvas.remove();
        
        var nProps = {items: nData};

        console.log("nProps", nProps, typeof(nProps))
        
        var uWheel = new spinWheel.Wheel(nContainer,nProps);

        uWheel.onCurrentIndexChange = (event) => {
            const nParagraph = document.querySelector('.wheel-value');
            const word = nProps["items"][event.currentIndex].label;
            nParagraph.textContent = word;
        }
        
        window.wheel = uWheel;
    }
    
});

import ws from 'k6/ws'
import {check} from 'k6'


export const options = {
    vus: 100,
    duration: '30s',

};

export default function(){
    const url = "ws://localhost:3334/cable";
    const params = {
        headers: {
            'Sec-Websocket-Protocol': 'actioncable-v1-json',
        },
    };

    const res = ws.connect(url,params,function(socket){
        socket.on('open',function(){
            console.log("connected");
            socket.send(JSON.stringify({
                command: "subscribe",
                identifier: JSON.stringify({channel: "ThermostatChannel"})
            }))
        });
        socket.on('message',function(data){
            console.log("Received",data);
        });
        socket.on('close',function(){
            console.log("Disconnected");
        });

        socket.setTimeout(function(){
            console.log("Closing Connection");
            socket.close();
        },1000);
    });
    check(res,{"Connected Succesufully": (r)=> r && r.status === 101});
}
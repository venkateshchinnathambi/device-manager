import consumer from "channels/consumer"

consumer.subscriptions.create(
  {channel: "DeviceChannel",device_id: "1234"}, {
  connected() {
    // Called when the subscription is ready for use on the server
     console.log("Connect to DeviceChannel")

  },

  disconnected() {
    // Called when the subscription has been terminated by the server
         console.log("Dis Connected to DeviceChannel")

  },

  received(data) {
    // Called when there's incoming data on the websocket for this channel
         console.log("Received from  DeviceChannel :",data)
          document.getElementById("device-updates").innerText = data.ZoneName
  }
});

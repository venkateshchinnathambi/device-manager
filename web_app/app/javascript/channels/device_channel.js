import consumer from "channels/consumer"

consumer.subscriptions.create("DeviceChannel", {
  connected() {
    // Called when the subscription is ready for use on the server
<<<<<<< HEAD:web_app/app/javascript/channels/device_channel.js
     console.log("Connect to DeviceChannel")
=======
     console.log("Connected to ThermostatChannel")
>>>>>>> 134670a8013b8ff2a4ed8d17bb538af3c78ae4ee:web_app/app/javascript/channels/thermostat_channel.js
  },

  disconnected() {
    // Called when the subscription has been terminated by the server
<<<<<<< HEAD:web_app/app/javascript/channels/device_channel.js
         console.log("Dis Connected to DeviceChannel")
=======
         console.log("Disconnected from ThermostatChannel")
>>>>>>> 134670a8013b8ff2a4ed8d17bb538af3c78ae4ee:web_app/app/javascript/channels/thermostat_channel.js
  },

  received(data) {
    // Called when there's incoming data on the websocket for this channel
         console.log("Received from  DeviceChannel :",data)
          document.getElementById("device-updates").innerText = data.ZoneName
  }
});

import consumer from "channels/consumer"

consumer.subscriptions.create("ThermostatChannel", {
  connected() {
    // Called when the subscription is ready for use on the server
     console.log("Connect to ThermostatChannel")
  },

  disconnected() {
    // Called when the subscription has been terminated by the server
         console.log("Dis Connected to ThermostatChannel")
  },

  received(data) {
    // Called when there's incoming data on the websocket for this channel
         console.log("Received from  ThermostatChannel :",data)
          document.getElementById("thermostat-updates").innerText = data.ZoneName
  }
});

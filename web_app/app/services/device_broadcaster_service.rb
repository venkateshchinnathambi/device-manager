# app/services/device_broadcaster_service.rb

<<<<<<< HEAD:web_app/app/services/device_broadcaster_service.rb
class DeviceBroadcasterService
  CHANNEL_NAME = "devices"
=======
class ThermostatBroadcasterService
  CHANNEL_NAME = "ThermostatChannel"
>>>>>>> 134670a8013b8ff2a4ed8d17bb538af3c78ae4ee:web_app/app/services/thermostat_broadcaster_service.rb

  def initialize(payload)
    @payload = payload
  end

  def call
    broadcast_to_clients
  end

  private

  def broadcast_to_clients
    ActionCable.server.broadcast(CHANNEL_NAME, @payload)
  rescue => e
    Rails.logger.error("Broadcast failed: #{e.message}")
  end
end

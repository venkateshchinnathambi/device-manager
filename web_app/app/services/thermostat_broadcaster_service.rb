# app/services/thermostat_broadcaster_service.rb

class ThermostatBroadcasterService
  CHANNEL_NAME = "ThermostatChannel"

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

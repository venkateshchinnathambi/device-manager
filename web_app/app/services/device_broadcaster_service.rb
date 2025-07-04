# app/services/device_broadcaster_service.rb
class DeviceBroadcasterService
  CHANNEL_NAME = "devices"
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

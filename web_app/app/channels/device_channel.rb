class DeviceChannel < ApplicationCable::Channel
  def subscribed
     device_id = params[:device_id]
     # device_id = "1234"
     if device_id.present?
     stream_from "devices:#{device_id}"
     else
      Rails.logger.error("DeviceChannel::subscribed - Invalide device id")
      reject
     end
    end

  def unsubscribed
    # Any cleanup needed when channel is unsubscribed
  end
end

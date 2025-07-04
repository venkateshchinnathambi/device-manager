class DeviceChannel < ApplicationCable::Channel
  def subscribed
     #device_id = params[:device_id]
     #stream_from "devices:#{device_id}" if device_id.present?
     stream_from "devices"
  end

  def unsubscribed
    # Any cleanup needed when channel is unsubscribed
  end
end

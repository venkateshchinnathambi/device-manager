class ThermostatChannel < ApplicationCable::Channel
  def subscribed
     Rails.logger.info("Subscribed to ThermostatChannel")
     stream_from "ThermostatChannel"
  end

  def unsubscribed
    Rails.logger.info("Unsubscribed from ThermostatChannel")
    # Any cleanup needed when channel is unsubscribed
  end
end

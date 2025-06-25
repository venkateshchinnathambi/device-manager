class ThermostatChannel < ApplicationCable::Channel
  def subscribed
     stream_from "thermostat"
  end

  def unsubscribed
    # Any cleanup needed when channel is unsubscribed
  end
end

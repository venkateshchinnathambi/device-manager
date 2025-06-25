require "ostruct"
class ThermostatUpdateConsumer < Racecar::Consumer 
    subscribes_to "thermostat.updates"

    def process(message)
        payload = JSON.parse(message.value)
        Rails.logger.info("message #{payload}")
        #ActionCable.server.broadcast("thermostat",payload)
        ThermostatBroadcasterService.new(payload).call
    end
end
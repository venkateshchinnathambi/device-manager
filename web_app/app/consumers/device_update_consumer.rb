require "ostruct"
class DeviceUpdateConsumer < Racecar::Consumer
    subscribes_to "devices.updates"

    def process(message)
        payload = JSON.parse(message.value)
        Rails.logger.info("message #{payload}")
        # ActionCable.server.broadcast("thermostat",payload)
        DeviceBroadcasterService.new(payload).call
    end
end

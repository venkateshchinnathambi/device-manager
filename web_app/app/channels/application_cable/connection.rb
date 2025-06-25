module ApplicationCable
  class Connection < ActionCable::Connection::Base
=begin
    identified_by :current_thermostat

    def connect
      self.current_thermostat = find_verified_thermostat
    end

    def find_verified_thermostat
      # Placeholder logic — replace with real auth
      Thermostat.find_by(id: request.params[:thermostat_id]) || reject_unauthorized_connection
    end
=end
  end
end

AnyCable.configure do |config|
  config.rpc_host = Rails.application.config_for(:any_cable).dig(Rails.env, "rpc_host")
end

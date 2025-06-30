AnyCable.configure do |config|
  env_config = Rails.application.config_for(:any_cable).fetch(Rails.env, {})

  host = env_config["rpc_host"] || "0.0.0.0"
  port = env_config["rpc_port"] || "50051"

  config.rpc_host = "#{host}:#{port}"
  config.secret = env_config["rpc_secret"] || ENV["ANYCABLE_RPC_SECRET"]
  config.log_level = env_config["log_level"] || "debug"

  puts "[AnyCable Init] rpc_host: #{config.rpc_host}, secret set? #{!config.secret.nil?}"
end

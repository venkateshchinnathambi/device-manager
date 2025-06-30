# config/initializers/anycable_patch.rb

if Rails.env.production?
  Rails.application.config.to_prepare do
    if defined?(AnyCable::GRPC::Server) && AnyCable::GRPC::Server.respond_to?(:build_server)
      class << AnyCable::GRPC::Server
        alias_method :original_build_server, :build_server unless method_defined?(:original_build_server)

        def build_server(**options)
          tls_credentials = options.delete(:tls_credentials) || {}
          host = ENV.fetch("ANYCABLE_RPC_HOST", "0.0.0.0")
          port = ENV.fetch("ANYCABLE_RPC_PORT", "50051")

          ::GRPC::RpcServer.new(**options).tap do |server|
            server.add_http2_port("#{host}:#{port}", server_credentials(**tls_credentials))
            server.handle(AnyCable::GRPC::Handler)
            server.handle(build_health_checker)
          end
        end
      end
      Rails.logger.info "[AnyCablePatch] build_server successfully patched."
    else
      Rails.logger.warn "[AnyCablePatch] build_server not defined on AnyCable::GRPC::Server. Patch skipped."
    end
  end
end

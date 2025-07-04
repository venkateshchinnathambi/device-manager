require "openssl"
require "jwt"
private_key = OpenSSL::PKey::RSA.new(File.read("./config/ssl/server.key"))

payload = {
  device_id: "1234",
  exp: 24
}

token = JWT.encode(payload, private_key, 'RS256')
p token


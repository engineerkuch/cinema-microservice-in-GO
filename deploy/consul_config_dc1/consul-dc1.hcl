data_dir = "/tmp/"
log_level = "DEBUG"

datacenter = "dc1"
primary_datacenter = "dc1"
node_name          = "Consul-Server"

server = true

bootstrap_expect = 1
ui = true

bind_addr = "0.0.0.0"
client_addr = "0.0.0.0"

ports {
  grpc = 8502
}

connect {
  enabled = true
}

advertise_addr = "10.7.0.10"
# advertise_addr_wan = "192.169.7.2"
enable_central_service_config = true
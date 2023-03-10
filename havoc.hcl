job "havoc" {
  datacenters = ["dc1"]
  type = "service"

  group "havoc" {
    count = 1

    task "havoc" {
      driver = "exec"

      config {
        command = "/usr/local/bin/havoc_server"
      }

      resources {
        cpu    = 500 # 500 MHz CPU
        memory = 256 # 256MB RAM
      }
    }
  }
}


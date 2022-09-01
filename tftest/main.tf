terraform {
  required_providers {
    megaport = {
      source = "megaport/megaport"
    }
  }
}


provider megaport {
  accept_purchase_terms = true
  environment = "production"
  delete_ports = false
  username = "svc-mp-ccs@stateless.net"
  password = "xyu1AHQ!xnm0zgb8bjm"
}

data megaport_locations potato {

}

data  megaport_ports potato {

}

output poato {
  value = data.megaport_ports.potato
}
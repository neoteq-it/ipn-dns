package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	_ "github.com/coredns/coredns/core/plugin"
	"github.com/coredns/coredns/coremain"
	_ "github.com/damomurf/coredns-tailscale"
	_ "github.com/neoteq-it/neoteqts4via6"
	_ "github.com/coredns/records"
)

func main() {
	dnsserver.Directives = append(dnsserver.Directives, "tailscale")
	dnsserver.Directives = append(dnsserver.Directives, "records")
	coremain.Run()
}

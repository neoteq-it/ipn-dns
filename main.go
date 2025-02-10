package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	_ "github.com/coredns/coredns/core/plugin"
	"github.com/coredns/coredns/coremain"
	_ "github.com/damomurf/coredns-tailscale"
	_ "github.com/neoteq-it/neoteqts4via6"
)

func main() {
	for i, directive := range dnsserver.Directives {
		if directive == "forward" {
			dnsserver.Directives = append(dnsserver.Directives[:i], dnsserver.Directives[i+1:]...)
			break
		}
	}

	// Fügen Sie Ihre benutzerdefinierten Plugins hinzu
	dnsserver.Directives = append(dnsserver.Directives, "tailscale")
	dnsserver.Directives = append(dnsserver.Directives, "neoteqts4via6")

	dnsserver.Directives = append(dnsserver.Directives, "forward")

	coremain.Run()
}

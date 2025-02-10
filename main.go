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
    // Entfernen Sie "any" aus der vordefinierten Liste, falls es bereits vorhanden ist
    for i, directive := range dnsserver.Directives {
        if directive == "any" {
            dnsserver.Directives = append(dnsserver.Directives[:i], dnsserver.Directives[i+1:]...)
            break
        }
    }

    // Fügen Sie Ihre benutzerdefinierten Plugins hinzu
    dnsserver.Directives = append(dnsserver.Directives, "tailscale")
    dnsserver.Directives = append(dnsserver.Directives, "records")
    dnsserver.Directives = append(dnsserver.Directives, "neoteqts4via6")

    // Fügen Sie "any" als letztes Plugin hinzu
    dnsserver.Directives = append(dnsserver.Directives, "any")

    coremain.Run()
}


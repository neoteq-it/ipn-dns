package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	_ "github.com/coredns/coredns/core/plugin"
	"github.com/coredns/coredns/coremain"
	_ "github.com/damomurf/coredns-tailscale"
	_ "github.com/neoteq-it/neoteqts4via6"
)

var directives = []string{
	"log",
	"errors",
	"template",
	"azure",
	"tailscale",
	"neoteqts4via6",
	"forward",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}

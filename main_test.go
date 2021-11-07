package main

import (
	"os"
	"testing"

	"github.com/ebrianne/cert-manager-webhook-duckdns/duckdns"
	"github.com/jetstack/cert-manager/test/acme/dns"
)

var (
	zone    = os.Getenv("TEST_ZONE_NAME")
	dnsname = os.Getenv("DNS_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.

	fixture := dns.NewFixture(duckdns.NewSolver(),
		dns.SetBinariesPath("__main__/hack/bin"),
		dns.SetResolvedZone(zone),
		dns.SetDNSName(dnsname),
		dns.SetAllowAmbientCredentials(false),
		dns.SetManifestPath("testdata/duckdns"),
	)

	fixture.RunConformance(t)
}

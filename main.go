package main

import (
	"os"

	"github.com/ebrianne/cert-manager-webhook-duckdns/duckdns"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
	"k8s.io/klog/v2"
)

func main() {
	GroupName := os.Getenv("GROUP_NAME")
	if GroupName == "" {
		klog.Fatal("GROUP_NAME must be specified")
	}

	// This will register our custom DNS provider with the webhook serving
	// library, making it available as an API under the provided GroupName.
	// You can register multiple DNS provider implementations with a single
	// webhook, where the Name() method will be used to disambiguate between
	// the different implementations.
	cmd.RunWebhookServer(GroupName, duckdns.NewSolver())
}

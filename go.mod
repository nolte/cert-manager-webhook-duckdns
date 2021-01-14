module github.com/ebrianne/cert-manager-webhook-duckdns

go 1.15

require (
	github.com/ebrianne/duckdns-go v1.0.1
	github.com/jetstack/cert-manager v1.1.0
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.19.3
	k8s.io/apiextensions-apiserver v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v0.19.3
	k8s.io/klog v1.0.0
)

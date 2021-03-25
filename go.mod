module github.com/ebrianne/cert-manager-webhook-duckdns

go 1.15

require (
	github.com/ebrianne/duckdns-go v1.0.2
	github.com/jetstack/cert-manager v1.2.1-0.20210324111646-720428406370
	github.com/pkg/errors v0.9.1
	k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/klog/v2 v2.8.0
)

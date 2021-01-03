# cert-manager-webhook-duckdns

[cert-manager](https://cert-manager.io/docs/) is a native Kubernetes certificate management controller. It can help with issuing certificates from a variety of sources, such as Let’s Encrypt, HashiCorp Vault, Venafi, a simple signing key pair, or self signed. It will ensure certificates are valid and up to date, and attempt to renew certificates at a configured time before expiry.

## TL;DR

```console
$ helm repo add 
$ helm install 
```

## Introduction

This chart bootstraps a duckdns [cert-manager-webhook](https://cert-manager.io/docs/configuration/acme/dns01/webhook/) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.12+
- Helm 3.0+

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm install my-release 
```

The command deploys cert-manager-webhook-duckdns on the Kubernetes cluster in the default configuration. The [Parameters](#parameters) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the cert-manager-webhook-duckdns chart and their default values.

| Parameter                          | Description                                     | Default                                                 |
|------------------------------------|-------------------------------------------------|---------------------------------------------------------|
| `groupName`                        | Group name for the webhook                      | `acme.duckdns.org`                                      |
| `logLevel`                         | Logging level                                   | `6`                                                     |
| `certManager.namespace`            | cert-manager namespace                          | `cert-manager`                                          |
| `certManager.serviceAccountName`   | cert-manager service account name               | `cert-manager`                                          |
| `duckdns.domain`                   | DuckDNS domain                                  | `""`                                                    |
| `duckdns.token`                    | DuckDNS token                                   | `""`                                                    |
| `clusterIssuer.email`              | Cluster issuer email address                    | `name@example.com`                                      |
| `clusterIssuer.staging.create`     | Create letsencrypt staging cluster issuer       | `false`                                                 |
| `clusterIssuer.production.create`  | Create letsencrypt production cluster issuer    | `false`                                                 |
| `image.repository`                 | Docker image repository                         | `ebrianne/cert-manager-webhook-duckdns`                 |
| `image.tag`                        | Docker image tag                                | `latest`                                                |
| `image.pullPolicy`                 | Docker image pull policy                        | `Always`                                                |
| `nameOverride`                     | Name override for the chart                     | `""`                                                    |
| `fullnameOverride`                 | Full name override for the chart                | `""`                                                    |
| `service.type`                     | Service type                                    | `ClusterIP`                                             |
| `service.port`                     | Service port                                    | `443`                                                   |
| `resources`                        | Pod ressources                                  | `nil`                                                   |
| `nodeSelector`                     | Node selector                                   | `nil`                                                   |
| `tolerations`                      | Node tolerations                                | `nil`                                                   |
| `affinity`                         | Node affinity                                   | `nil`                                                   |
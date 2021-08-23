
---

**This project are currently not maintained, please use the amazing frok [ebrianne/cert-manager-webhook-duckdns](https://github.com/ebrianne/cert-manager-webhook-duckdns)**

---


# ACME DuckDNS Certmanager Webhook

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: example-issuer
spec:
  acme:
    email: example@example.com
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: example-issuer-account-key
    solvers:
    - dns01:
        webhook:
          groupName: acme.duckdns.org
          solverName: duckdns
          config:
            zoneName: just-mfg
            secretName: duckdns
```

### Running the test suite

All DNS providers **must** run the DNS01 provider conformance testing suite,
else they will have undetermined behaviour when used with cert-manager.

**It is essential that you configure and run the test suite when creating a
DNS01 webhook.**

An example Go test file has been provided in [main_test.go]().

You can run the test suite with:

```bash
$ TEST_ZONE_NAME=example.com go test .
```

The example file has a number of areas you must fill in and replace with your
own options in order for tests to pass.



helm upgrade -i duckdns ./duckdns-webhook -n kube-system

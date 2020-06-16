# To (ken) Co (ntent) Ser (ver)

Serve a staticFS only if token in header matches


## Configuration

```bash
# Envvars

# Path to the directory where the amce stuff is placed
ACMEDIR=/acmedir

# Path to the dir which will be served under /content/
CONTENTDIR=/content

# Token which must be present in the "x-auth-token" header
TOKEN=1234

# FQDN for which acme is done.
DNSNAME=mydomain.org

```

If `DNSNAME` and `ACMEDIR` is specified, the server will listen on port 443 and tries to get a cert from Let's Encrypt.  
Make sure your `DNSNAME` points to the server. 
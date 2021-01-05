# ssl-checker

Check SSL Certificate expiration.

## Usage

- Build
  `go build -o ssl-checker cmd/*.go`

- Check the website
  `./ssl-checker [domain]`

- Result example
  `{ "cn":"*.adriantirusli.me", "not_after":"2021-09-28T12:00:00Z", "not_before":"2020-08-29T00:00:00Z", "dns_names":["*.adriantirusli.me","adriantirusli.me"] "signature_algorithm":"SHA256-RSA","issuer":"Amazon", "organizations":["Amazon"],"expiration":22982505.661074 }`

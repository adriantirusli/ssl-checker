package certs

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"net"
	"time"
)

//Cert are collections of field
type Cert struct {
	CommonName         string    `json:"cn"`
	NotAfter           time.Time `json:"not_after"`
	NotBefore          time.Time `json:"not_before"`
	DNSNames           []string  `json:"dns_names"`
	SignatureAlgorithm string    `json:"signature_algorithm"`
	IssuerCommonName   string    `json:"issuer"`
	Organizations      []string  `json:"organizations"`
	ExpireAfter        float64   `json:"expiration"`
}

//getVerifiedCertificateChains is to connects to the given network address using dialer
func getVerifiedCertificateChains(addr string, timeoutSecond time.Duration) ([][]*x509.Certificate, error) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: timeoutSecond * time.Second}, "tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	chains := conn.ConnectionState().VerifiedChains
	return chains, nil
}

//ParseRemoteCertificate is to return the resulting TLS connection
func ParseRemoteCertificate(addr string) (*Cert, error) {
	chains, err := getVerifiedCertificateChains(addr, time.Duration(30))
	if err != nil {
		return nil, err
	}

	var cert *Cert
	for _, chain := range chains {
		for _, crt := range chain {
			if !crt.IsCA {
				cert = &Cert{
					CommonName:         crt.Subject.CommonName,
					NotAfter:           crt.NotAfter,
					NotBefore:          crt.NotBefore,
					DNSNames:           crt.DNSNames,
					SignatureAlgorithm: crt.SignatureAlgorithm.String(),
					IssuerCommonName:   crt.Issuer.CommonName,
					Organizations:      crt.Issuer.Organization,
					ExpireAfter:        time.Until(crt.NotAfter).Seconds(),
				}
			}
		}
	}
	return cert, err
}

//Jsonify used to change the format given to JSON
func (cert *Cert) Jsonify() string {
	b, _ := json.Marshal(cert)
	return string(b)
}

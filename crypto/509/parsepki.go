package main

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	const pubPEM1 = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlRuRnThUjU8/prwYxbty
WPT9pURI3lbsKMiB6Fn/VHOKE13p4D8xgOCADpdRagdT6n4etr9atzDKUSvpMtR3
CP5noNc97WiNCggBjVWhs7szEe8ugyqF23XwpHQ6uV1LKH50m92MbOWfCtjU9p/x
qhNpQQ1AZhqNy5Gevap5k8XzRmjSldNAFZMY7Yv3Gi+nyCwGwpVtBUwhuLzgNFK/
yDtw2WcWmUU7NuC8Q6MWvPebxVtCfVp/iQU6q60yyt6aGOBkhAX0LpKAEhKidixY
nP9PNVBvxgu3XZ4P36gZV6+ummKdBVnc3NqwBLu5+CcdRdusmHPHd5pHf4/38Z3/
6qU2a/fPvWzceVTEgZ47QjFMTCTmCwNt29cvi7zZeQzjtwQgn4ipN9NibRH/Ax/q
TbIzHfrJ1xa2RteWSdFjwtxi9C20HUkjXSeI4YlzQMH0fPX6KCE7aVePTOnB69I/
a9/q96DiXZajwlpq3wFctrs1oXqBp5DVrCIj8hU2wNgB7LtQ1mCtsYz//heai0K9
PhE4X6hiE0YmeAZjR0uHl8M/5aW9xCoJ72+12kKpWAa0SFRWLy6FejNYCYpkupVJ
yecLk/4L1W0l6jQQZnWErXZYe0PNFcmwGXy1Rep83kfBRNKRy5tvocalLlwXLdUk
AIU+2GKjyT3iMuzZxxFxPFMCAwEAAQ==
-----END PUBLIC KEY-----`

        const pubPEM =`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC27Mf4tAITwQx4xIsVg/3657ghIIl4yOKyLPpUKpEBf9bYdpYacYkmpLhAhuqTDZLFCuDqvHicL7owA+AkMYfCO8eL5FYWookSiuwUUqiUjb0r5w7hdFQ0yhLABZoNfmK212DyvH5jUfq2v7gqxt74ZEzsy/dm8LHEgjcE/sqxFORQdXzCuzzNZtCdsiaCv5oT9LfTyE15Ou5Z4pQPj+fCRgyWsNfsn8gnM7I2cZqFu5dsyU/S58C0Thi35ULjGYQdmmrv4FDyvG4B8koeKBD49nrbro4IdtWM2ITANvvOH+xyne/qZsPb0eJIhdiD2uNzla5xrcPHRAglr4oZct+gW/IfGe3K10PVuqdaT9IyRmYOyi3IHuxxNIYE9tFbz0lnfquiD+7icKWIOQA4qQI5l7hRFzzttxfvMADN/7eCWe7OAz7PaOK9aMgOBfQDTFiQSJe46croE9KQQatrClHf9Noq+DRsDNCik3g4ZLco/Z+XkbyKBFXWHW3WpitRfuJ0VW5DbMV79iKFtpT2DggNtIGKB1r7/yzia0e9V6tz6/WXvuNMXKTnGZY4giCsDMvzXnfv4CMfApdZyIe7cZQUnr+WU415UUJ69TA4v3WT5Kl033BxqKSpDPEALLJ3KPRXlFySnwVXaDBA0LdbCYahH+2gTcSbIIZDtcVcZBTqxw== mengwp_2004@163.com`

	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		fmt.Println("pub is of type RSA:", pub)
	case *dsa.PublicKey:
		fmt.Println("pub is of type DSA:", pub)
	case *ecdsa.PublicKey:
		fmt.Println("pub is of type ECDSA:", pub)
	default:
		panic("unknown type of public key")
	}
}


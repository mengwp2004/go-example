package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
)

var decrypted string

func init() {
	flag.StringVar(&decrypted, "d", "", "加密过的数据")
	flag.Parse()
}

func main() {
	var data []byte
	var err error
	if decrypted != "" {
		data, err = base64.StdEncoding.DecodeString(decrypted)
		if err != nil {
			panic(err)
		}
	} else {
		data, err = RsaEncrypt([]byte("polaris@studygolang.com"))
		if err != nil {
			panic(err)
		}
		fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	}
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// 公钥和私钥可以从文件中读取
var privateKey1 = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

var publicKey1 = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

var privateKey2 = []byte(
`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDm4Q8k4EY4wel0dTPRnPKc1wHcPqbW4VnLq3xhPKFtm1iA/OJz
3FFHFZqoKtV5vHxUWdGLLhFkq/E7hwl8fa3EB0f6LwOpeMAN9o/Fm6blxjW+eWOI
lc0EuX4ls3i6EJycPgz/6Se/ZrNwYBPN5EbNWFNDxGMu/4NyRVOnIcPlGQIDAQAB
AoGBAOOe1zNf+VtTsjoioqfxuC9wumDNU+dd3Q2zT7j98ZkCIstohkn72BXg8s0B
Dd4rOdhfhPtWu8cozXgGtg+KSJiva7YTorl88aWI+OKnWh3dSh/iJuTdtGrtRC9n
OC8XyY3HW+0xD/1aM2gExY5WDGvp9xDeNp+8UlRDuxiBE8MRAkEA/JDHkCrz9Tg4
2YIq+/9yQZiq7iTvaeQSdszeH8h1gYseh3/+iJh4Mfvxqpnx7m3gPKJEXHbiw0o1
1Z8ppBCAAwJBAOoEySrhV4y11KjLkKNgHNBCHhDyZyK2d22pQfP8QIB0GAqq6MK4
ZTLHvyXiAL8EQ5Few8Ne39YGS/3J9sM+IbMCQCfPQkVtH8r1M7DFHbVezEmeoMKs
u7f7JRXosNJdrHfgz80X1az+K2PljHARl9q3IvMruI7Chne1yMqFLfEYULUCQE1d
W6wDZ1ArZKyQ77YzNhNbaFkt6g69x+nHBPjGMgTFXJVaPyzwjPQmuPirKJf9ruDG
NW3HVbSJzGGQfTW6uqUCQQCEnODdwL8fsvI+ibyO/N42IDFdh7ShnvYr5BLRatc6
z3CM2R88//1dtk9vVpRKwFWEUk1q2L1g3TL67a7AIw32
-----END RSA PRIVATE KEY-----`)


var publicKey2 =[]byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDm4Q8k4EY4wel0dTPRnPKc1wHc
PqbW4VnLq3xhPKFtm1iA/OJz3FFHFZqoKtV5vHxUWdGLLhFkq/E7hwl8fa3EB0f6
LwOpeMAN9o/Fm6blxjW+eWOIlc0EuX4ls3i6EJycPgz/6Se/ZrNwYBPN5EbNWFND
xGMu/4NyRVOnIcPlGQIDAQAB
-----END PUBLIC KEY----`)

var privateKey =[]byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCdBeRNjBrY66HW7JKtWAxMVkJCs8FboII9zyrxCbbPYRQqjT5w
IiM8fWVraCiuB2Zi6P7SVXz4V37f3CzuSs3+FSF20lkexCkpgeSeEHDdl5zlwYJ4
I9B1F2iK3UPWgGSlo023nw0ugXIp0Yi3xydoptZxs8/2h7Nnu9NULRrV3QIDAQAB
AoGBAJNsJq7P7ZzfjbDVp9hrpBA/pDIvxkgaaG1ThcOMFUCaqm5Q8eAhjOeL655t
ylOpyuzuR+B+NriKnSnb86s+PCrqe3dtoWxGnhqPA3jqKE7WIYgObMhJS7vXV52o
vCYtZS5D0G89vxl5sFnmlGxEZWnr5Bre+mA6wKQWWOFC4oXZAkEAyd/tR1a+QbYv
o8IbfNbpIfAkne3KCk6zWXB4/KkEzGxZw3kzlriowrgWe2Xw2CD73OtH8MASPy1z
SjskWqg+ywJBAMcfeq1H0TxmcIOmir28DQN3MMJNmRA+zvfcihEGu3ZlZ5aCdwJa
qqvFoqOrwZivrPsfFXqUaMSSPd+rvj1KwPcCQDPUVB0uviYs2kpW+auxmaVm/F+3
v1mqHw3lfIqR8nNxlhJDueUGf/2OOedwosc6oK91kMDU71pdFu6GBK1p6MUCQC6R
hyBFXoLQtlmjtmjfNO/tjSK0ASsSdVb0ZfnLaEYtOwexJWIQkD+x7fJ3NcSeuUbV
ozFSyhnG1d+ci9ZLyEsCQAUp40dU+yhPvEW7RtLYFvT8vWba58Z5ER/PZo40qIVV
/qUfGW6hhgiI8dzgePUB3z6GVGvOCbp0DHy7/3j+dCE=
-----END RSA PRIVATE KEY-----`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCdBeRNjBrY66HW7JKtWAxMVkJC
s8FboII9zyrxCbbPYRQqjT5wIiM8fWVraCiuB2Zi6P7SVXz4V37f3CzuSs3+FSF2
0lkexCkpgeSeEHDdl5zlwYJ4I9B1F2iK3UPWgGSlo023nw0ugXIp0Yi3xydoptZx
s8/2h7Nnu9NULRrV3QIDAQAB
-----END PUBLIC KEY-----
`)
// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

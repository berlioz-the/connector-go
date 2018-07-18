package berlioz

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/ssm"
)

// TBD
type SecretAccessor struct {
	name string
}

// TBD
func Secret(name string) SecretAccessor {
	log.Printf("[Secret] %s\n", name)
	return SecretAccessor{name: name}
}

func (x SecretAccessor) getKeyStr(ctx context.Context, kind string) (string, error) {
	log.Printf("[InterfaceSecret::GetKey]: \n")
	nr := getNativeResource(kind, x.name)
	svc := nr.SSM()
	mybool := true
	params := ssm.GetParameterInput{WithDecryption: &mybool}
	res, err := svc.GetParameter(ctx, &params)
	if err != nil {
		return "", err
	}
	if res.Parameter == nil {
		return "", fmt.Errorf("Parameter not present")
	}
	return *res.Parameter.Value, nil
}

func (x SecretAccessor) getPublicKey(ctx context.Context) (*rsa.PublicKey, error) {
	keyStr, err := x.getKeyStr(ctx, "secret_public_key")
	if err != nil {
		return nil, err
	}
	keyBlock, _ := pem.Decode([]byte(keyStr))
	pubKey, parseErr := x509.ParsePKCS1PublicKey(keyBlock.Bytes)
	if parseErr != nil {
		fmt.Printf("Load public key error: %v\n", parseErr)
		panic(parseErr)
	}
	return pubKey, nil
}

func (x SecretAccessor) getPrivateKey(ctx context.Context) (*rsa.PrivateKey, error) {
	keyStr, err := x.getKeyStr(ctx, "secret_private_key")
	if err != nil {
		return nil, err
	}
	keyBlock, _ := pem.Decode([]byte(keyStr))
	privKey, parseErr := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if parseErr != nil {
		fmt.Printf("Load public key error: %v\n", parseErr)
		panic(parseErr)
	}
	return privKey, nil
}

// TBD
func (x SecretAccessor) Encrypt(ctx context.Context, data string) (string, error) {
	key, err := x.getPublicKey(ctx)
	if err != nil {
		return "", err
	}

	hash := sha1.New()
	random := rand.Reader
	encryptedData, encryptErr := rsa.EncryptOAEP(hash, random, key, []byte(data), nil)
	if encryptErr != nil {
		return "", encryptErr
	}
	encodedData := base64.URLEncoding.EncodeToString(encryptedData)
	return encodedData, nil
}

// TBD
func (x SecretAccessor) Decrypt(ctx context.Context, data string) (string, error) {
	key, err := x.getPrivateKey(ctx)
	if err != nil {
		return "", err
	}

	hash := sha1.New()
	random := rand.Reader
	rawData, _ := base64.URLEncoding.DecodeString(data)
	decryptedData, decryptErr := rsa.DecryptOAEP(hash, random, key, rawData, nil)
	if decryptErr != nil {
		return "", decryptErr
	}
	str := string(decryptedData[:])
	return str, nil
}

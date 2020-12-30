package cryptopro

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCryptSignHash(t *testing.T) {

	capBytes := []byte("Hello world")

	store, err := CertOpenSystemStore("MY")
	defer CertCloseStore(store, 0)
	if err != nil {
		t.Fatal("Can't open MY store")
	}

	//client, err := CertFindCertificateInStore(store, "39da49123dbe70e953f394074d586eb692f3328e",
	//	CERT_FIND_SHA1_HASH)
	client, err := CertFindCertificateInStore(store, "beb12450d5c33200925b544d7bf7f8cb01463010",
		CERT_FIND_SHA1_HASH)

	if err != nil {
		t.Fatal(err)
	}

	context, err := CryptAquireCertificatePrivateKey(client)
	if err != nil {
		t.Fatal(err)
	}

	hash, err := CreateCryptHash(context, CALG_GR3411_2012_256)
	if err != nil {
		t.Fatal(err)
	}

	err = hash.CryptHashData(capBytes)
	if err != nil {
		t.Fatal(err)
	}

	_, err = hash.CryptGetHashParam()
	if err != nil {
		t.Fatal(err)
	}

	sigBytes, err := CryptSignHash(hash, 0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("signature: %s\n", hex.EncodeToString(sigBytes))

}
package crypto

import (
	"testing"

	"github.com/bocheninc/L0/components/crypto/crypter"
	_ "github.com/bocheninc/L0/components/crypto/crypter/secp256k1"
	_ "github.com/bocheninc/L0/components/crypto/crypter/sm2"
)

func Benchmark_secp256k1(b *testing.B) {
	b.StopTimer()
	crypter := crypter.MustCrypter("secp256k1")
	priv, pub, _ := crypter.GenerateKey()
	msg := []byte("hello")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sig, err := crypter.Sign(priv, msg)
		if err != nil {
			b.Errorf("Sign Error %s", err)
		}

		if !crypter.Verify(pub, msg, sig) {
			b.Errorf("Invalid signature")
		}
	}
}

func Benchmark_sm2_double256(b *testing.B) {
	b.StopTimer()
	crypter := crypter.MustCrypter("sm2_double256")
	priv, pub, _ := crypter.GenerateKey()
	msg := []byte("hello")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sig, err := crypter.Sign(priv, msg)
		if err != nil {
			b.Errorf("Sign Error %s", err)
		}

		if !crypter.Verify(pub, msg, sig) {
			b.Errorf("Invalid signature")
		}
	}
}

func Benchmark_sm2c_double256(b *testing.B) {
	b.StopTimer()
	crypter := crypter.MustCrypter("sm2c_double256")
	priv, pub, _ := crypter.GenerateKey()
	msg := []byte("hello")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sig, err := crypter.Sign(priv, msg)
		if err != nil {
			b.Errorf("Sign Error %s", err)
		}

		if !crypter.Verify(pub, msg, sig) {
			b.Errorf("Invalid signature")
		}
	}
}

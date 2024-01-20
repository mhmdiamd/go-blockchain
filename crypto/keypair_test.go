package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
  privKey := GeneratePrivateKey()
  pubKey := privKey.PublicKey()
  // address := pubKey.Address()

  msg := []byte("Hello World")
  sig, err := privKey.Sign(msg)
  assert.Nil(t, err)

  assert.True(t, sig.Verify(pubKey, msg))

  b := sig.Verify(pubKey, msg)
  assert.True(t, b)
}

func TestKeypairSignVerifySuccess(t *testing.T) {
  privKey := GeneratePrivateKey()
  pubKey := privKey.PublicKey()
  msg := []byte("Hello World")

  sig, err := privKey.Sign(msg)
  assert.Nil(t, err)

  assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeypairSignVerifyFail(t *testing.T) {
  privKey := GeneratePrivateKey()
  pubKey := privKey.PublicKey()
  msg := []byte("Hello World")

  sig, err := privKey.Sign(msg)
  assert.Nil(t, err)

  otherPrivKey := GeneratePrivateKey()
  otherPubKey := otherPrivKey.PublicKey()

  assert.False(t, sig.Verify(otherPubKey, msg))
  assert.False(t, sig.Verify(pubKey, []byte("xxxx")))
}

package aesencdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

var (
	// ErrBadKeySize ...
	ErrBadKeySize = fmt.Errorf("key must be 16 or 32 bytes to select AES-128 or AES-256")
)

// AESEncDec ...
type AESEncDec struct {
	aead cipher.AEAD
	b64  *base64.Encoding
}

// New ...
func New(key []byte) (*AESEncDec, error) {
	if len(key) != 16 && len(key) != 32 {
		return nil, ErrBadKeySize
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ed := &AESEncDec{
		aead: aead,
		b64:  base64.URLEncoding.WithPadding(base64.NoPadding),
	}

	return ed, nil
}

// Encrypt ...
func (ed *AESEncDec) Encrypt(src []byte) ([]byte, error) {
	nonce := make([]byte, ed.aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	nLen := ed.b64.EncodedLen(len(nonce))

	sealed := ed.aead.Seal(src[:0], nonce, src, nil)
	sLen := ed.b64.EncodedLen(len(sealed))

	dst := make([]byte, nLen+sLen)
	ed.b64.Encode(dst, nonce)
	ed.b64.Encode(dst[nLen:], sealed)

	return dst, nil
}

// Decrypt ...
func (ed *AESEncDec) Decrypt(src []byte) ([]byte, error) {
	sLen := ed.b64.DecodedLen(len(src))

	dbuf := make([]byte, sLen)
	n, err := ed.b64.Decode(dbuf, src)
	if err != nil {
		return nil, err
	}
	dbuf = dbuf[:n]

	nLen := ed.aead.NonceSize()

	return ed.aead.Open(dbuf[:0], dbuf[:nLen], dbuf[nLen:], nil)
}

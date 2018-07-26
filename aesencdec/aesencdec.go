package aesencdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AESEncDec ...
type AESEncDec struct {
	aead cipher.AEAD
}

// New ...
func New(key []byte) (*AESEncDec, error) {
	if len(key) != 16 && len(key) != 32 {
		return nil, errors.New("key must be 16 or 32 bytes to select AES-128 or AES-256")
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
	}

	return ed, nil
}

// Encrypt ...
func (ed *AESEncDec) Encrypt(src []byte) ([]byte, error) {
	nonce := make([]byte, ed.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	nLen := base64.URLEncoding.EncodedLen(len(nonce))

	sealed := ed.aead.Seal(src[:0], nonce, src, nil)
	sLen := base64.URLEncoding.EncodedLen(len(sealed))

	dst := make([]byte, nLen+sLen)
	base64.URLEncoding.Encode(dst, nonce)
	base64.URLEncoding.Encode(dst[nLen:], sealed)

	return dst, nil
}

// Decrypt ...
func (ed *AESEncDec) Decrypt(src []byte) ([]byte, error) {
	sLen := base64.URLEncoding.DecodedLen(len(src))

	dbuf := make([]byte, sLen)
	n, err := base64.URLEncoding.Decode(dbuf, src)
	if err != nil {
		return nil, err
	}
	dbuf = dbuf[:n]

	nLen := ed.aead.NonceSize()

	return ed.aead.Open(dbuf[:0], dbuf[:nLen], dbuf[nLen:], nil)
}

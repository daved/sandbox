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
	aead  cipher.AEAD
	nonce []byte
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

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ed := &AESEncDec{
		aead:  aead,
		nonce: nonce,
	}

	return ed, nil
}

// Encrypt ...
func (ed *AESEncDec) Encrypt(src []byte) ([]byte, error) {
	sealed := ed.aead.Seal(nil, ed.nonce, src, nil)

	dst := make([]byte, base64.URLEncoding.EncodedLen(len(sealed)))
	base64.URLEncoding.Encode(dst, sealed)

	return dst, nil
}

// Decrypt ...
func (ed *AESEncDec) Decrypt(src []byte) ([]byte, error) {
	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	n, err := base64.URLEncoding.Decode(dbuf, src)
	if err != nil {
		return nil, err
	}
	dbuf = dbuf[:n]

	return ed.aead.Open(nil, ed.nonce, dbuf, nil)
}

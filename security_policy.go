// Copyright 2020 Converter Systems LLC. All rights reserved.

package opcua

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
)

// SecurityPolicyURIs
const (
	SecurityPolicyURINone           = "http://opcfoundation.org/UA/SecurityPolicy#None"
	SecurityPolicyURIBasic128Rsa15  = "http://opcfoundation.org/UA/SecurityPolicy#Basic128Rsa15"
	SecurityPolicyURIBasic256       = "http://opcfoundation.org/UA/SecurityPolicy#Basic256"
	SecurityPolicyURIBasic256Sha256 = "http://opcfoundation.org/UA/SecurityPolicy#Basic256Sha256"
	SecurityPolicyURIBestAvailable  = ""
)

// SecurityPolicy is a mapping of PolicyURI to security settings
type SecurityPolicy interface {
	PolicyURI() string
	RSASign(priv *rsa.PrivateKey, plainText []byte) ([]byte, error)
	RSAVerify(pub *rsa.PublicKey, plainText, signature []byte) error
	RSAEncrypt(pub *rsa.PublicKey, plainText []byte) ([]byte, error)
	RSADecrypt(priv *rsa.PrivateKey, cipherText []byte) ([]byte, error)
	SymHMACFactory(key []byte) hash.Hash
	RSAPaddingSize() int
	SymSignatureSize() int
	SymSignatureKeySize() int
	SymEncryptionBlockSize() int
	SymEncryptionKeySize() int
}

// securityPolicyNone ...
type securityPolicyNone struct {
}

func newSecurityPolicyNone() *securityPolicyNone {
	return &securityPolicyNone{}
}

// PolicyURI ...
func (p *securityPolicyNone) PolicyURI() string { return SecurityPolicyURINone }

// RSASign ...
func (p *securityPolicyNone) RSASign(priv *rsa.PrivateKey, plainText []byte) ([]byte, error) {
	return nil, BadSecurityPolicyRejected
}

// RSAVerify ...
func (p *securityPolicyNone) RSAVerify(pub *rsa.PublicKey, plainText, signature []byte) error {
	return BadSecurityPolicyRejected
}

// RSAEncrypt ...
func (p *securityPolicyNone) RSAEncrypt(pub *rsa.PublicKey, plainText []byte) ([]byte, error) {
	return nil, BadSecurityPolicyRejected
}

// RSADecrypt ...
func (p *securityPolicyNone) RSADecrypt(priv *rsa.PrivateKey, cipherText []byte) ([]byte, error) {
	return nil, BadSecurityPolicyRejected
}

// SymHMACFactory ...
func (p *securityPolicyNone) SymHMACFactory(key []byte) hash.Hash {
	return nil
}

// RSAPaddingSize ...
func (p *securityPolicyNone) RSAPaddingSize() int { return 0 }

// SymSignatureSize ...
func (p *securityPolicyNone) SymSignatureSize() int { return 0 }

// SymSignatureKeySize ...
func (p *securityPolicyNone) SymSignatureKeySize() int { return 0 }

// SymEncryptionBlockSize ...
func (p *securityPolicyNone) SymEncryptionBlockSize() int { return 1 }

// SymEncryptionKeySize ...
func (p *securityPolicyNone) SymEncryptionKeySize() int { return 0 }

// securityPolicyBasic128Rsa15 ...
type securityPolicyBasic128Rsa15 struct {
}

// PolicyURI ...
func (p *securityPolicyBasic128Rsa15) PolicyURI() string { return SecurityPolicyURIBasic128Rsa15 }

// RSASign ...
func (p *securityPolicyBasic128Rsa15) RSASign(priv *rsa.PrivateKey, plainText []byte) ([]byte, error) {
	hashed := sha1.Sum(plainText)
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashed[:])
}

// RSAVerify ...
func (p *securityPolicyBasic128Rsa15) RSAVerify(pub *rsa.PublicKey, plainText, signature []byte) error {
	hashed := sha1.Sum(plainText)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed[:], signature)
}

// RSAEncrypt ...
func (p *securityPolicyBasic128Rsa15) RSAEncrypt(pub *rsa.PublicKey, plainText []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

// RSADecrypt ...
func (p *securityPolicyBasic128Rsa15) RSADecrypt(priv *rsa.PrivateKey, cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}

// SymHMACFactory ...
func (p *securityPolicyBasic128Rsa15) SymHMACFactory(key []byte) hash.Hash {
	return hmac.New(sha1.New, key)
}

// RSAPaddingSize ...
func (p *securityPolicyBasic128Rsa15) RSAPaddingSize() int { return 11 }

// SymSignatureSize ...
func (p *securityPolicyBasic128Rsa15) SymSignatureSize() int { return 20 }

// SymSignatureKeySize ...
func (p *securityPolicyBasic128Rsa15) SymSignatureKeySize() int { return 16 }

// SymEncryptionBlockSize ...
func (p *securityPolicyBasic128Rsa15) SymEncryptionBlockSize() int { return 16 }

// SymEncryptionKeySize ...
func (p *securityPolicyBasic128Rsa15) SymEncryptionKeySize() int { return 16 }

// securityPolicyBasic256 ...
type securityPolicyBasic256 struct {
}

// PolicyURI ...
func (p *securityPolicyBasic256) PolicyURI() string { return SecurityPolicyURIBasic256 }

// RSASign ...
func (p *securityPolicyBasic256) RSASign(priv *rsa.PrivateKey, plainText []byte) ([]byte, error) {
	hashed := sha1.Sum(plainText)
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashed[:])
}

// RSAVerify ...
func (p *securityPolicyBasic256) RSAVerify(pub *rsa.PublicKey, plainText, signature []byte) error {
	hashed := sha1.Sum(plainText)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed[:], signature)
}

// RSAEncrypt ...
func (p *securityPolicyBasic256) RSAEncrypt(pub *rsa.PublicKey, plainText []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, plainText, []byte{})
}

// RSADecrypt ...
func (p *securityPolicyBasic256) RSADecrypt(priv *rsa.PrivateKey, cipherText []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, cipherText, []byte{})
}

// SymHMACFactory ...
func (p *securityPolicyBasic256) SymHMACFactory(key []byte) hash.Hash {
	return hmac.New(sha1.New, key)
}

// RSAPaddingSize ...
func (p *securityPolicyBasic256) RSAPaddingSize() int { return 42 }

// SymSignatureSize ...
func (p *securityPolicyBasic256) SymSignatureSize() int { return 20 }

// SymSignatureKeySize ...
func (p *securityPolicyBasic256) SymSignatureKeySize() int { return 24 }

// SymEncryptionBlockSize ...
func (p *securityPolicyBasic256) SymEncryptionBlockSize() int { return 16 }

// SymEncryptionKeySize ...
func (p *securityPolicyBasic256) SymEncryptionKeySize() int { return 32 }

// securityPolicyBasic256Sha256 ...
type securityPolicyBasic256Sha256 struct {
}

// PolicyURI ...
func (p *securityPolicyBasic256Sha256) PolicyURI() string { return SecurityPolicyURIBasic256Sha256 }

// RSASign ...
func (p *securityPolicyBasic256Sha256) RSASign(priv *rsa.PrivateKey, plainText []byte) ([]byte, error) {
	hashed := sha256.Sum256(plainText)
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
}

// RSAVerify ...
func (p *securityPolicyBasic256Sha256) RSAVerify(pub *rsa.PublicKey, plainText, signature []byte) error {
	hashed := sha256.Sum256(plainText)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// RSAEncrypt ...
func (p *securityPolicyBasic256Sha256) RSAEncrypt(pub *rsa.PublicKey, plainText []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, plainText, []byte{})
}

// RSADecrypt ...
func (p *securityPolicyBasic256Sha256) RSADecrypt(priv *rsa.PrivateKey, cipherText []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, cipherText, []byte{})
}

// SymHMACFactory ...
func (p *securityPolicyBasic256Sha256) SymHMACFactory(key []byte) hash.Hash {
	return hmac.New(sha256.New, key)
}

// RSAPaddingSize ...
func (p *securityPolicyBasic256Sha256) RSAPaddingSize() int { return 42 }

// SymSignatureSize ...
func (p *securityPolicyBasic256Sha256) SymSignatureSize() int { return 32 }

// SymSignatureKeySize ...
func (p *securityPolicyBasic256Sha256) SymSignatureKeySize() int { return 32 }

// SymEncryptionBlockSize ...
func (p *securityPolicyBasic256Sha256) SymEncryptionBlockSize() int { return 16 }

// SymEncryptionKeySize ...
func (p *securityPolicyBasic256Sha256) SymEncryptionKeySize() int { return 32 }
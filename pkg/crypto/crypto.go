package crypto

import (
	"hash"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/dsa"
	"crypto/sha256"
    "crypto/sha512"
    "crypto"
    "errors"
	"encoding/base32"
	"crypto/rand"
	"math/big"
	

	"github.com/bzanni/2DDoc/pkg/model"
	"github.com/bzanni/2DDoc/pkg/authority"
)


func hashPayload(h model.Header, payload string) ( []byte) {
	var hashFunc hash.Hash
    switch(h.SignatureHash){
    case "SHA-256":
    	hashFunc = sha256.New()
    case "SHA-384":
    	hashFunc = sha512.New384()
    case "SHA-512":
    	hashFunc = sha512.New()
    }
    hashFunc.Write([]byte(payload))
    return hashFunc.Sum(nil)
}

func Verify(code model.Barcode2DDoc) (valid bool, err error) {
	valid = false;
	key, err := authority.GetPublicKey(code.Header.CertificationAuthorityId, code.Header.CertificatId)
	if err != nil {
        return false, err
    }
	payload := code.Header.Raw+code.Message.Raw
	payloadSignature := code.Signature.Raw
	data := hashPayload(code.Header, payload)

	decodedSig , err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(payloadSignature)
    if err != nil {
    	return false, err
    }

    
	switch key := key.(type) {
	case *rsa.PublicKey:

		verifyErr := rsa.VerifyPKCS1v15(key, crypto.SHA256, []byte(code.Message.Raw), []byte(code.Signature.Raw))
		valid = verifyErr == nil
		
	case *dsa.PublicKey:

		r := new(big.Int).SetBytes(decodedSig[:len(decodedSig)/2])
    	s := new(big.Int).SetBytes(decodedSig[len(decodedSig)/2:])

	    valid = dsa.Verify(
	        key,
	        data,
	        r, 
	        s)

		
	case *ecdsa.PublicKey: 
	 	
		r := new(big.Int).SetBytes(decodedSig[:len(decodedSig)/2])
    	s := new(big.Int).SetBytes(decodedSig[len(decodedSig)/2:])

	    valid = ecdsa.Verify(
	        key,
	        data,
	        r, 
	        s)

	default:
		return false, errors.New("Unknown pubkey type")

	}

	if !valid {
		return false, errors.New("Verification failed")
	}

	return true, nil

} 

func Sign (h model.Header, m model.Message) (s model.Signature, err error) {

	private_key, err := authority.GetPrivateKey(h.CertificationAuthorityId, h.CertificatId)
	if err != nil {
		return s, err
	} 

	payload := h.Raw + m.Raw
	data := hashPayload(h, payload)

	sigR, sigS, err := ecdsa.Sign(rand.Reader, private_key, data)
	if err != nil {
    	return s, err
    }

	sig := append(sigR.Bytes(), sigS.Bytes()...)

	s.Raw = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(sig) 

	return s, nil
}




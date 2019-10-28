package c40signature

import (
    "github.com/bzanni/2DDoc/pkg/model"
    "github.com/bzanni/2DDoc/pkg/encoding/encoding_utils"
)

func Decode(h model.Header, m model.Message, payload string) (s model.Signature, remaining string, err error){

	

	sig := encoding_utils.Substring(payload, 0, h.SignatureLength)
	remaining = encoding_utils.TrimLeftChars(payload, h.SignatureLength)

	s = model.Signature {
		Raw: sig,
	}

	return s, remaining, nil
}


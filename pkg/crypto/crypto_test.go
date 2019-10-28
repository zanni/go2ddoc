package crypto

import (
	"testing"

	"github.com/bzanni/2DDoc/pkg/encoding"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func assertNotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Fatalf("%s nil", a)
	}
}

func assertNil(t *testing.T, a interface{}) {
	if a != nil {
		t.Fatalf("%s not nil", a)
	}
}


func TestVerify(t *testing.T) {
	var dataToSign = "DC04FR000001198519D31201FR90MAITRE/SPECIMEN/NATACHA92RAISON SOCIALE DE TEST94SAISIE CONSERVATOIRE DE CREANCES962111201791MME/BERTHIER/CORINNE93RAISON SOCIALE DU TIERS CONCERNE951896547853AB0CNB2WS43TNFSXELLKOVZXI2LDMUXGM4RPGE4DSNRVGQ3TQNJTIFBA"
	var signature = "OOXND3NRRGDZKBYZ6VDMHSM7WHJ323ICLGTSEELJ74OW3E4GGFYI3GX6IXCN4HF45JWYZKHHU7GXTBMCSHOSU5GOUHJYN4PIH6VAA2Q"
	var payload = dataToSign + string(31) + signature
	barcode, err := encoding.Decode(payload)
	assertNil(t, err)
	ok, err := Verify(barcode)

	assertNil(t, err)
	assertEqual(t, true, ok)

	var falseSignature = "OOXND3NRRGDZKBYZ6VDMHSM7WHJ323ICLGTSEELJ74OW3E4GGFYI3GX6IXCN4HF45JWYZKHHU7GXTBMCSHOSU5GOUHJYN4PIH6VAA25"
	payload = dataToSign + string(31) + falseSignature
	barcode, err = encoding.Decode(payload)
	assertNil(t, err)
	ok, err = Verify(barcode)

	assertNotNil(t, err)
	assertEqual(t, false, ok)
}

func TestSign(t *testing.T) {
	var dataToSign = "DC04FR000001198519D31201FR90MAITRE/SPECIMEN/NATACHA92RAISON SOCIALE DE TEST94SAISIE CONSERVATOIRE DE CREANCES962111201791MME/BERTHIER/CORINNE93RAISON SOCIALE DU TIERS CONCERNE951896547853AB0CNB2WS43TNFSXELLKOVZXI2LDMUXGM4RPGE4DSNRVGQ3TQNJTIFBA"
	var signature = "OOXND3NRRGDZKBYZ6VDMHSM7WHJ323ICLGTSEELJ74OW3E4GGFYI3GX6IXCN4HF45JWYZKHHU7GXTBMCSHOSU5GOUHJYN4PIH6VAA2Q"
	
	var payload = dataToSign + string(31) + signature
	barcode, err := encoding.Decode(payload)
	assertNil(t, err)
	assertEqual(t, dataToSign, barcode.Header.Raw+barcode.Message.Raw)
	s, err := Sign(barcode.Header, barcode.Message)
	assertNil(t, err)
	barcode.Signature = s
	ok, err := Verify(barcode)
	assertNil(t, err)
	assertEqual(t, true, ok)
}
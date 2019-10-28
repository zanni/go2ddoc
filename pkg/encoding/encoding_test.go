package encoding

import (
	"testing"
	"strings"
	"github.com/bzanni/2DDoc/pkg/spec"
	"github.com/bzanni/2DDoc/pkg/crypto"
	"github.com/bzanni/2DDoc/pkg/encoding/encoding_utils"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func assertNil(t *testing.T, a interface{}) {
	if a != nil {
		t.Fatalf("%s not nil", a)
	}
}

func assertNotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Fatalf("%s nil", a)
	}
}

func assertDataEquals(t *testing.T, given map[string]string , expectedMandatory []map[string]string, expectedFacultative []map[string]string) {
	for d := range expectedMandatory {
		value := expectedMandatory[d]["value"]
		value = strings.Trim(strings.Split(value, " (")[0], " ")
		value = strings.Replace(value, "<vide>", "" , -1)
		assertEqual(t, strings.Trim(strings.Split(given[expectedMandatory[d]["key"]], " (")[0], " "), value)		
	}

	for d := range expectedFacultative {
		value := expectedFacultative[d]["value"]
		value = strings.Trim(strings.Split(value, " (")[0], " ")
		value = strings.Replace(value, "<vide>", "" , -1)
		assertEqual(t, strings.Trim(strings.Split(given[expectedFacultative[d]["key"]], " (")[0], " "), value)		
	}
}

func test(t *testing.T, data spec.Test) {
	var payload = data.Message
	payload = strings.Replace(payload, "<GS>", string(29), -1)
	payload = strings.Replace(payload, "<US>", string(31), -1)

	barcode, err := Decode(payload)
	assertNil(t, err)
	assertNotNil(t, barcode)
	assertNotNil(t, barcode.Header)
	assertNotNil(t, barcode.Message)
	assertNotNil(t, barcode.Signature)


	assertEqual(t, barcode.Header.DocumentTypeId, data.DocumentTypeId)
	assertEqual(t, barcode.Header.PerimeterId, data.PerimeterId)

	expectedCreationDate, err := encoding_utils.HexDate(data.CreationDateHex)
	assertNil(t, err)
	assertEqual(t, barcode.Header.DocumentCreationDate, expectedCreationDate.Unix())

	expectedSignatureDate, err := encoding_utils.HexDate(data.SignatureDateHex)
	assertNil(t, err)
	assertEqual(t, barcode.Header.SignatureCreationDate, expectedSignatureDate.Unix())

	assertDataEquals(t, barcode.Message.Data, data.MandatoryData, data.FacultativeData)

	ok, err := crypto.Verify(barcode)
	assertNil(t, err)
	assertEqual(t, true, ok)
	
}

func TestDecode(t *testing.T) {

	test_list := spec.TestList()

	for i := range test_list {		
		test(t, test_list[i])
	}

}
package authority

import (
	"testing"
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

func TestGetCertificateProperty(t *testing.T) {
	hash, length, err := GetCertificateProperty("FR00", "0001")
	assertNil(t, err)
	assertNotNil(t, hash)
	assertNotNil(t, length)

	hash, length, err = GetCertificateProperty("xxxx", "0001")
	assertNotNil(t, err)


	hash, length, err = GetCertificateProperty("FR00", "xxxx")
	assertNotNil(t, err)

	hash, length, err = GetCertificateProperty("FR01", "0002")
	assertNil(t, err)
	assertNotNil(t, hash)
	assertNotNil(t, length)

	hash, length, err = GetCertificateProperty("FR02", "0002")
	assertNil(t, err)
	assertNotNil(t, hash)
	assertNotNil(t, length)

	hash, length, err = GetCertificateProperty("FR03", "0002")
	assertNil(t, err)
	assertNotNil(t, hash)
	assertNotNil(t, length)

	hash, length, err = GetCertificateProperty("FR04", "0002")
	assertNil(t, err)
	assertNotNil(t, hash)
	assertNotNil(t, length)
}

func TestGetPrivateKey(t *testing.T) {
	private_key, err := GetPrivateKey("FR00", "0001")
	assertNil(t, err)
	assertNotNil(t, private_key)

	private_key, err = GetPrivateKey("xxxx", "0001")
	assertNotNil(t, err)


	private_key, err = GetPrivateKey("FR00", "xxxx")
	assertNotNil(t, err)
}
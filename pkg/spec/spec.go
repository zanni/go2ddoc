package spec

import (
	"log"
	"errors"
	"crypto"
	"crypto/x509"
	"crypto/tls"
	// "fmt"
	"encoding/json"
	// "io/ioutil"

	// "github.com/gobuffalo/packr"
	// "golang.org/x/mobile/asset"
)

type Test struct {
    CreationDate string
    CreationDateHex string
    DocumentTypeId string
    DocumentTypeName string
    FacultativeData []map[string]string
    MandatoryData []map[string]string
    Message string
   	Perimeter string
    PerimeterId string
    Signature string
    SignatureDate string
    SignatureDateHex string
    SignedData string
}

type DataIdentifier struct {
	Id string
	Max string
	Min string
	ShortDesc string
	Type string
}

type DocumentType struct {
	CreatorType string
	Id string
	IsCreationDateMandatory bool
	UserType string
}

var test_list []Test

var data_identifier_map map[string]DataIdentifier

var document_type_map map[string]DocumentType

var test_cert *x509.Certificate
var test_private_key crypto.PrivateKey

const certFile = `-----BEGIN CERTIFICATE-----
MIICVzCCAT8CCQCpMEvcR9M4RTANBgkqhkiG9w0BAQUFADBPMQswCQYDVQQGEwJG
UjETMBEGA1UECgwKQUMgREUgVEVTVDEcMBoGA1UECwwTMDAwMiAwMDAwMDAwMDAw
MDAwMDENMAsGA1UEAwwERlIwMDAeFw0xMjExMDExMzQ3NDZaFw0xNTExMDExMzQ3
NDZaMFcxCzAJBgNVBAYTAkZSMRswGQYDVQQKDBJDRVJUSUZJQ0FUIERFIFRFU1Qx
HDAaBgNVBAsMEzAwMDIgMDAwMDAwMDAwMDAwMDAxDTALBgNVBAMMBDAwMDEwWTAT
BgcqhkjOPQIBBggqhkjOPQMBBwNCAASpjw18zWKAiJO+xNQ2550YNKHW4AHXDxxM
3M2dni/iKfckBRTo3cDKmNDHRAycxJKEmg+9pz/DkvTaCuB/hMI8MA0GCSqGSIb3
DQEBBQUAA4IBAQA6HN+w/bzIdg0ZQF+ELrocplehP7r5JuRJNBAgmoqoER7IonCv
KSNUgUVbJ/MB4UKQ6CgzK7AOlCpiViAnBv+i6fg8Dh9evoUcHBiDvbl19+4iREaO
oyVZ8RAlkp7VJKrC3s6dJEmI8/19obLbTvdHfY+TZfduqpVl63RSxwLG0Fjl0SAQ
z9a+KJSKZnEvT9I0iUUgCSnqFt77RSppziQTZ+rkWcfd+BSorWr8BHqOkLtj7EiV
amIh+g3A8JtwV7nm+NUbBlhh2UPSI0eevsRjQRghtTiEn0wflVBX7xFP9zXpViHq
Ij+R9WiXzWGFYyKuAFK1pQ2QH8BxCbvdNdff
-----END CERTIFICATE-----`

const keyFile = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINbI/xP+yGOgp79v7qibvYs03x+cSIaiKzpOhJsScwDDoAoGCCqGSM49
AwEHoUQDQgAEqY8NfM1igIiTvsTUNuedGDSh1uAB1w8cTNzNnZ4v4in3JAUU6N3A
ypjQx0QMnMSShJoPvac/w5L02grgf4TCPA==
-----END EC PRIVATE KEY-----`

// func loadFull(name string) ([]byte, error) {
//     f, errOpen := asset.Open(name)
//     if errOpen != nil {
//         return nil, errOpen
//     }
//     defer f.Close()
//     buf, errRead := ioutil.ReadAll(f)
//     if errRead != nil {
//         return nil, errRead
//     }
//     log.Printf("loaded: %s (%d bytes)", name, len(buf))
//     return buf, nil
// }

func init() {

	// box := packr.NewBox("../../assets")

	// test cert, test key init
	// f, err := loadFull("test_cert.pem")
	// certFile := string(f) /
	// certFile, err := box.FindString("test_cert.pem")
	// check(err)
	// f, err = loadFull("test_key.pem")
	// keyFile := string(f)
	// keyFile, err := box.FindString("test_key.pem")
	// check(err)
	pair, err := tls.X509KeyPair([]byte(certFile), []byte(keyFile))
	check(err)
	certX509, err := x509.ParseCertificate(pair.Certificate[0])
	check(err)


	test_private_key = pair.PrivateKey
	test_cert = certX509


	// box = packr.NewBox("../../assets/spec")

	// test list init
	// f, err = loadFull("spec/test_list.json")
	// test_list_json := string(f)
	// test_list_json, err := box.FindString("test_list.json")
	// check(err)
	err = json.Unmarshal([]byte(test_list_json), &test_list)
	check(err)

	// data identifier list init
	// f, err = loadFull("spec/data_identifier.json")
	// data_identifier_map_json := string(f)
	// data_identifier_map_json, err := box.FindString("data_identifier.json")
	// check(err)
	err = json.Unmarshal([]byte(data_identifier_map_json), &data_identifier_map)
	check(err)

	// document type list init
	// f, err = loadFull("spec/document_type.json")
	// document_type_map_json := string(f)
	// document_type_map_json, err := box.FindString("document_type.json")
	// check(err)
	err = json.Unmarshal([]byte(document_type_map_json), &document_type_map)
	check(err)

	
}

func TestList() ([]Test){
	return test_list 
}

func DataIdentifierFromId(id string) (data_identifier DataIdentifier, err error) {
	val, ok :=  data_identifier_map[id]
	if !ok {
        return data_identifier, errors.New("Unknown dataIdentifierId: "+id)
    }
	return val, nil
}

func DocumentTypeFromId(id string) (document_type DocumentType, err error) {
	val, ok :=  document_type_map[id]
	if !ok {
        return document_type, errors.New("Unknown documentTypeId: "+id)
    }
	return val, nil
}

func GetTestCertificat() (*x509.Certificate) {
	return test_cert
}

func GetTestPrivateKey() (crypto.PrivateKey) {
	return test_private_key
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}
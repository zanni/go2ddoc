package authority

import (
	"io/ioutil"
	"net/http"
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"github.com/beevik/etree"
	"errors"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/elliptic"
	"crypto"
	"github.com/bzanni/2DDoc/pkg/spec"
)

const certFirst = "-----BEGIN CERTIFICATE-----"
const certLast = "-----END CERTIFICATE-----"

var cert map[string]map[string]*x509.Certificate = make(map[string]map[string]*x509.Certificate)
var privateKey map[string]map[string]crypto.PrivateKey = make(map[string]map[string]crypto.PrivateKey)

func init() {
	setCertificat("FR00", "0001", spec.GetTestCertificat())
	setPrivateKey("FR00", "0001", spec.GetTestPrivateKey())
	loadANTSAuthorities()
}

func GetPublicKey(authority string, id string) (key interface{}, err error) {
	cert, err := getCertificat(authority, id)
	if(err != nil) {
		return nil, errors.New(
			"Unknown certificat for certificationAuthorityId: "+authority + " and certificatId: "+id)
	}
	return cert.PublicKey, nil
}

func GetPrivateKey(authority string, id string) (key *ecdsa.PrivateKey, err error) {
	if auth, ok := privateKey[authority]; ok {
		if key, ok := auth[id]; ok {
			return key.(*ecdsa.PrivateKey), nil 
		} 
	}
	return nil, errors.New(
		"Unknown private key for certificationAuthorityId: "+authority + " and certificatId: "+id)
}

func GetCertificateProperty(authority string, id string) (hash string, length int, err error) {
	
	cert, err := getCertificat(authority, id)
	if(err != nil) {
		return "", 0, err
	}

	// fmt.Println("pubkey alg", cert.PublicKeyAlgorithm, cert.SignatureAlgorithm)
    var signatureLength int
	var signatureHash string
	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		// fmt.Println("pub is of type RSA:", pub)
		signatureLength = 103
	    signatureHash = "SHA-256"
	case *dsa.PublicKey:
		// fmt.Println("pub is of type DSA:", pub)
		signatureLength = 103
	    signatureHash = "SHA-256"
	case *ecdsa.PublicKey: 
	    if pub.Curve == elliptic.P256() {
	        signatureLength = 103
	        signatureHash = "SHA-256"
	    } else if pub.Curve == elliptic.P224() {
	        signatureLength = 154
	        signatureHash = "SHA-384"
	    } else if pub.Curve == elliptic.P384() {
	        signatureLength = 154
	        signatureHash = "SHA-384"
	    } else if pub.Curve == elliptic.P521() {
	        signatureLength = 212
	        signatureHash = "SHA-512"
	    }
	default:
		return "", 0, errors.New(
		"Unknown PublicKey type: "+authority + " and certificatId: "+id)
	}
	// fmt.Println(signatureHash, signatureLength)
	return signatureHash, signatureLength, nil
}

func loadANTSAuthorities() {
	resp, err := http.Get("https://ants.gouv.fr/content/download/517/5670/version/18/file/ANTS_2D-Doc_TSL_v12_sign.xml")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	doc := etree.NewDocument()
	err = doc.ReadFromString(string(body))
	if err != nil {
		
		panic(err)
	}

	var trustServiceStatusList TrustServiceStatusList

	xml.Unmarshal(body, &trustServiceStatusList)

	for _, provider := range trustServiceStatusList.TrustServiceProviderList.TrustServiceProvider {

		for _, service := range provider.TSPServices.TSPService {
			cert := certFirst;
			for i, c := range service.ServiceInformation.ServiceDigitalIdentity.DigitalId.X509Certificate {
				if i%76 == 0 {
					cert += "\n"
				}
				cert += string(c)
			}
			cert += "\n"
			cert += certLast
			block, _ := pem.Decode([]byte(cert))
			x509Cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
		        panic(err)
		    }
		    runes := []rune(x509Cert.Subject.OrganizationalUnit[len(x509Cert.Subject.OrganizationalUnit)-1])
		    setCertificat(x509Cert.Subject.CommonName, string(runes[0:4]), x509Cert)
		}
	}
}

func getCertificat(authority string, id string) (certificate *x509.Certificate, err error) {
	if auth, ok := cert[authority]; ok {
		if certificat, ok := auth[id]; ok {
			return certificat, nil 
		} 
	}
	return nil, errors.New(
		"Unknown certificat for certificationAuthorityId: "+authority + " and certificatId: "+id)
}

func setCertificat(authority string, id string, certificat *x509.Certificate) {
	if auth, ok := cert[authority]; !ok {
		auth = make(map[string]*x509.Certificate)
		auth[id] = certificat
		cert[authority] = auth
	} else {
		cert[authority][id] = certificat
	}	
}

func setPrivateKey(authority string, id string, key crypto.PrivateKey) {
	if auth, ok := privateKey[authority]; !ok {
		auth = make(map[string]crypto.PrivateKey)
		auth[id] = key
		privateKey[authority] = auth
	} else {
		privateKey[authority][id] = key
	}	
}



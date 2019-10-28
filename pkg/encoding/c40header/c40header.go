package c40header

import (
    "strconv"
    "errors"
    
    // "crypto/elliptic"

    "github.com/bzanni/2DDoc/pkg/model"
    "github.com/bzanni/2DDoc/pkg/authority"
    "github.com/bzanni/2DDoc/pkg/spec"
    "github.com/bzanni/2DDoc/pkg/encoding/encoding_utils"
)


const MandatoryIndentificationFlag = "DC"


func Decode(payload string) (decoded model.Header, remaining string, err error){

	err = nil
    identificationFlag := encoding_utils.Substring(payload, 0, 2)
    if(identificationFlag != MandatoryIndentificationFlag) {
    	return decoded, "", errors.New("Malformed payload: Missing mandatory indentification flag")
    }

    version, err := strconv.Atoi(encoding_utils.Substring(payload, 2, 4))
    if(err != nil || version > 4 ) {
    	return decoded, "", errors.New("Malformed payload: Unparsable version")
    }

    certificationAuthorityId := encoding_utils.Substring(payload, 4, 8)

    certificatId := encoding_utils.Substring(payload, 8, 12)

    documentCreationDate, err := encoding_utils.HexDate(encoding_utils.Substring(payload, 12, 16))
    if(err != nil) {
    	return decoded, "", errors.New("Malformed payload: Unparsable document creation date")
    }

    signatureCreationDate, err := encoding_utils.HexDate(encoding_utils.Substring(payload, 16, 20))
    if(err != nil) {
    	return decoded, "", errors.New("Malformed payload: Unparsable signature creation date")
    }


    documentTypeId := encoding_utils.Substring(payload, 20, 22)
    documentType, err := spec.DocumentTypeFromId(documentTypeId)
    if err != nil {
        return decoded, "", err
    }
    if documentType.IsCreationDateMandatory && documentCreationDate.IsZero() {
        return decoded, "", errors.New("Malformed payload: createDate is mandatory for documentTypeId: "+documentTypeId)
    }

    signatureHash, signatureLength, err := authority.GetCertificateProperty(certificationAuthorityId, certificatId)
    if(err != nil) {
        return decoded, "", err
    }

    decoded = model.Header{
		IdentificationFlag: identificationFlag,
		Version: version,
		CertificationAuthorityId: certificationAuthorityId,
		CertificatId: certificatId,
		DocumentCreationDate: documentCreationDate.Unix(),
		SignatureCreationDate: signatureCreationDate.Unix(),
		DocumentTypeId: documentTypeId,
        SignatureHash: signatureHash,
        SignatureLength: signatureLength,
	}

    var length = 0
    switch version {
    case 1:
        length = 22
    case 2:
        length = 22
    case 3:
        decoded.PerimeterId = encoding_utils.Substring(payload, 22, 24)
        length = 24
    case 4:
    	decoded.PerimeterId = encoding_utils.Substring(payload, 22, 24)
    	decoded.DocumentCountryId = encoding_utils.Substring(payload, 24, 26)
        length = 26
    }

    remaining = encoding_utils.TrimLeftChars(payload, length)
    decoded.Raw = encoding_utils.Substring(payload, 0, length)


	return decoded, remaining, err


}




package model

type Barcode2DDoc struct {
    Header Header
    Message Message
    Signature Signature
}


type Header struct {
    IdentificationFlag string
    Version  int
    CertificationAuthorityId string
    CertificatId string
    DocumentCreationDate int64
    SignatureCreationDate int64
    DocumentTypeId string
    PerimeterId string
    DocumentCountryId string
    
    SignatureHash string
    SignatureLength int
    Raw string
}

type Message struct {
    Data map[string]string    
    Raw string
}


type Signature struct {
    Raw string
}




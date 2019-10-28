package mobile2DDoc


import(
    "strings"
    "github.com/bzanni/2DDoc/pkg/model"
    "github.com/bzanni/2DDoc/pkg/encoding"
    "github.com/bzanni/2DDoc/pkg/spec"
)

type BindBarcode2DDoc struct {
    Header *model.Header
    Message *model.Message
    Signature *model.Signature
}


func DataIdentifierFromId(id string) (data_identifier *spec.DataIdentifier, err error) {
    val, err := spec.DataIdentifierFromId(id)
    if err != nil {
        return data_identifier, err
    }
    return &val, nil
}

func DocumentTypeFromId(id string) (document_type *spec.DocumentType, err error) {
    val, err := spec.DocumentTypeFromId(id)
    if err != nil {
        return document_type, err
    }
    return &val, nil
}

func DecodeMobile(payload string) ( barcode *BindBarcode2DDoc, err error) {

    b, err := encoding.Decode(payload)
    if err != nil {
        return barcode, err
    }

    bar := BindBarcode2DDoc{
        Header:&b.Header,
        Message:&b.Message,
        Signature:&b.Signature,
    }

    return &bar, err
}
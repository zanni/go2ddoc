package encoding

import(
	"github.com/bzanni/2DDoc/pkg/model"
    "github.com/bzanni/2DDoc/pkg/encoding/c40header"
    "github.com/bzanni/2DDoc/pkg/encoding/c40message"
    "github.com/bzanni/2DDoc/pkg/encoding/c40signature"
)


func Decode(payload string) ( barcode model.Barcode2DDoc, err error) {

    var remaining = payload
    h, remaining, err := c40header.Decode(remaining)
    if err != nil {
        return barcode, err
    }
    
    m, remaining, err := c40message.Decode(h, remaining)
    if err != nil {
        return barcode, err
    }

    s, remaining, err := c40signature.Decode(h, m, remaining)
    if err != nil {
        return barcode, err
    }

    b := model.Barcode2DDoc {
        Header: h,
        Message: m,
        Signature: s,
    }

    return b, err
}
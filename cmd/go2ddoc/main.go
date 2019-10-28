package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "encoding/json"
    "github.com/bzanni/2DDoc/pkg/crypto"
    "github.com/bzanni/2DDoc/pkg/encoding"
    "github.com/bzanni/2DDoc/pkg/model"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}

func decodeFile(file string) (c model.Barcode2DDoc) {
    log.Println("Decoding file: "+file)
	dat, err := ioutil.ReadFile(file)
    check(err)
    payload := string(dat)
    code, err := encoding.Decode(payload)
    check(err)
    return code
}

func prettyPrint(i interface{}) {
    s, _ := json.MarshalIndent(i, "", "\t")
    fmt.Println(string(s))
}

func prettyPrintDecode(c model.Barcode2DDoc) {
    prettyPrint(c.Header)
    prettyPrint(c.Message)
    prettyPrint(c.Signature)

    if ok, err := crypto.Verify(c); ok {
        log.Println("Barcode is valid")
    } else {
        log.Fatal(err)
    }

}

func main() {
    prettyPrintDecode(decodeFile("assets/barcode_c40_v2.txt"))
    prettyPrintDecode(decodeFile("assets/barcode_c40_mrz_v3.txt"))
    prettyPrintDecode(decodeFile("assets/barcode_c40_v3.txt"))
    prettyPrintDecode(decodeFile("assets/barcode_c40_v4.txt"))
}
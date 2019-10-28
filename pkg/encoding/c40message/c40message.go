package c40message

import (
	"strings"
	"errors"
	"strconv"

    "github.com/bzanni/2DDoc/pkg/model"
    "github.com/bzanni/2DDoc/pkg/encoding/encoding_utils"
    "github.com/bzanni/2DDoc/pkg/spec"
)


const ASCII_GROUP_SEPARATOR = 29
const ASCII_RECORD_SEPARATOR = 30
const ASCII_UNIT_SEPARATOR = 31



func parseField( payload string) (key string, value string, remaining string, err error){

	fieldId := encoding_utils.Substring(payload, 0, 2)
	content := encoding_utils.TrimLeftChars(payload, 2)


	field, err := spec.DataIdentifierFromId(fieldId)

	if(err != nil) {
		return "", "", "", err
	}


	max, errMax := strconv.Atoi(strings.Trim(field.Max, " "))

	var terminate = false
	var i = 0
	for !terminate {
			
		if( 
			(errMax == nil && i == max) || 
			(i >= len(content)) ||
			content[i] == ASCII_UNIT_SEPARATOR){
			terminate = true
		} else {
			if(content[i] == ASCII_GROUP_SEPARATOR || 
				content[i] == ASCII_RECORD_SEPARATOR) {
				terminate = true
				
			} else {
				value += string(content[i])
			}
			i++
		}
	
	}

	remaining =  encoding_utils.TrimLeftChars(content, i)

	return fieldId, value, remaining, nil
}

func Decode(header model.Header, payload string) (m model.Message, remaining string, err error){

	err = nil
	res := make(map[string]string)

	remaining = payload

	var fieldId = ""
	var value = ""
	var terminate = false
	for !terminate {
		fieldId, value, remaining, err = parseField(remaining)
		res[fieldId] = value
		if(err != nil) {
			return model.Message{}, "", err
		} else if(len(remaining) == 0) {
			return model.Message{}, "", errors.New("Malformed payload: Missing signature")
		} else if(remaining[0] == ASCII_UNIT_SEPARATOR) {
			remaining = encoding_utils.TrimLeftChars(remaining, 1)
			terminate = true
		}
	}

	m = model.Message{
        Data: res,
        Raw: encoding_utils.Substring(payload, 0, len(payload) - len(remaining) - 1 ),
    }

	return m, remaining, err
}

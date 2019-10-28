package spec

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestDataIdentifierFromId(t *testing.T) {
	dataIdentifier, err := DataIdentifierFromId("01")
	assertEqual(t, nil, err)
	assertEqual(t, "01", dataIdentifier.Id)
	assertEqual(t, "Aucune", dataIdentifier.Max)
	assertEqual(t, "0", dataIdentifier.Min)
	assertEqual(t, "Identifiant unique du document.", dataIdentifier.ShortDesc)
	assertEqual(t, "Alphanumérique", dataIdentifier.Type)
}

func TestDocumentTypeFromId(t *testing.T) {
	documentType, err := DocumentTypeFromId("00")
	assertEqual(t, nil, err)
	assertEqual(t, "Document émis spécifiquement pour servir de justificatif de domicile.", documentType.CreatorType)
	assertEqual(t, "00", documentType.Id)
	assertEqual(t, true, documentType.IsCreationDateMandatory)
	assertEqual(t, "Justificatif de domicile", documentType.UserType)
}
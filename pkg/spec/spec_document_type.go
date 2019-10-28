package spec

const document_type_map_json = `
{
    "00": {
        "creatorType": "Document émis spécifiquement pour servir de justificatif de domicile.",
        "id": "00",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de domicile"
    },
    "01": {
        "creatorType": "Factures de fournisseur d’énergie - Factures de téléphonie - Factures de fournisseur d’accès internet - Factures de fournisseur d’eau",
        "id": "01",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de domicile"
    },
    "02": {
        "creatorType": "Avis de taxe d’habitation",
        "id": "02",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de domicile"
    },
    "03": {
        "creatorType": "Relevé d’identité bancaire.",
        "id": "03",
        "isCreationDateMandatory": false,
        "userType": "Justificatif de domiciliation bancaire"
    },
    "04": {
        "creatorType": "Avis d’impôt sur le revenu",
        "id": "04",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de ressources"
    },
    "05": {
        "creatorType": "Relevé d’Identité SEPAmail",
        "id": "05",
        "isCreationDateMandatory": false,
        "userType": "Justificatif de domiciliation bancaire"
    },
    "06": {
        "creatorType": "Bulletin de salaire",
        "id": "06",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de ressources"
    },
    "07": {
        "creatorType": "Titre d’identité",
        "id": "07",
        "isCreationDateMandatory": true,
        "userType": "Justificatif d’identité"
    },
    "08": {
        "creatorType": "MRZ",
        "id": "08",
        "isCreationDateMandatory": true,
        "userType": "Justificatif d’identité"
    },
    "09": {
        "creatorType": "Facture étendue",
        "id": "09",
        "isCreationDateMandatory": true,
        "userType": "Justificatif fiscal"
    },
    "10": {
        "creatorType": "Contrat de travail",
        "id": "10",
        "isCreationDateMandatory": true,
        "userType": "Justificatif d’emploi"
    },
    "11": {
        "creatorType": "Relevé de compte",
        "id": "11",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de ressources"
    },
    "12": {
        "creatorType": "Acte d’huissier",
        "id": "12",
        "isCreationDateMandatory": true,
        "userType": "Justificatif juridique/judiciaire"
    },
    "A0": {
        "creatorType": "Certificat de qualité de l’air",
        "id": "A0",
        "isCreationDateMandatory": true,
        "userType": "Justificatif écologique de véhicule"
    },
    "A1": {
        "creatorType": "Courrier Permis à Points",
        "id": "A1",
        "isCreationDateMandatory": true,
        "userType": "Justificatif permis de conduire"
    },
    "A2": {
        "creatorType": "Carte Mobilité Inclusion (CMI)",
        "id": "A2",
        "isCreationDateMandatory": true,
        "userType": "Justificatif de santé"
    },
    "A3": {
        "creatorType": "Macaron VTC (Véhicule de Transport avec Chauffeur)",
        "id": "A3",
        "isCreationDateMandatory": false,
        "userType": "Justificatif d’activité"
    },
    "A4": {
        "creatorType": "Certificat de décès",
        "id": "A4",
        "isCreationDateMandatory": true,
        "userType": "Justificatif médical"
    },
    "A5": {
        "creatorType": "Carte T3P (Transport Public Particulier de Personnes)",
        "id": "A5",
        "isCreationDateMandatory": false,
        "userType": "Justificatif d’activité"
    },
    "A6": {
        "creatorType": "Carte Professionnelle Sapeur-Pompier",
        "id": "A6",
        "isCreationDateMandatory": false,
        "userType": "Justificatif d’activité"
    },
    "A7": {
        "creatorType": "Certificat de qualité de l’air (V2)",
        "id": "A7",
        "isCreationDateMandatory": true,
        "userType": "Justificatif écologique de véhicule"
    },
    "B0": {
        "creatorType": "Diplôme",
        "id": "B0",
        "isCreationDateMandatory": true,
        "userType": "Justificatif académique"
    },
    "B1": {
        "creatorType": "Attestation de Versement de la Contribution à la Vie Etudiante",
        "id": "B1",
        "isCreationDateMandatory": true,
        "userType": "Justificatif académique"
    }
}
`
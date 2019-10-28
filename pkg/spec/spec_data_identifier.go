package spec

const data_identifier_map_json = `
{
    "01": {
        "id": "01",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Identifiant unique du document.",
        "type": "Alphanumérique"
    },
    "02": {
        "id": "02",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Catégorie de document",
        "type": "Alphanumérique"
    },
    "03": {
        "id": "03",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Sous-catégorie de document",
        "type": "Alphanumérique"
    },
    "04": {
        "id": "04",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Application de composition",
        "type": "Alphanumérique"
    },
    "05": {
        "id": "05",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Version de l’application de composition",
        "type": "Alphanumérique"
    },
    "06": {
        "id": "06",
        "max": "4",
        "min": "4",
        "shortDesc": "Date de l’association entre le document et le code 2D-Doc.",
        "type": "Alphanumérique"
    },
    "07": {
        "id": "07",
        "max": "6",
        "min": "6",
        "shortDesc": "Heure de l’association entre le document et le code 2D-Doc.",
        "type": "Numérique"
    },
    "08": {
        "id": "08",
        "max": "4",
        "min": "4",
        "shortDesc": "Date d’expiration du document",
        "type": "Alphanumérique"
    },
    "09": {
        "id": "09",
        "max": "4",
        "min": "4",
        "shortDesc": "Nombre de pages du document",
        "type": "Numérique"
    },
    "0A": {
        "id": "0A",
        "max": "9",
        "min": "9",
        "shortDesc": "Editeur du 2D-Doc",
        "type": "Numérique"
    },
    "0B": {
        "id": "0B",
        "max": "9",
        "min": "9",
        "shortDesc": "Intégrateur du 2D-Doc",
        "type": "Numérique"
    },
    "0C": {
        "id": "0C",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "URL du document",
        "type": "Alphanumérique et symboles"
    },
    "10": {
        "id": "10",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 1 de la norme adresse postale du bénéficiaire de la prestation",
        "type": "Alphanumérique"
    },
    "11": {
        "id": "11",
        "max": "38",
        "min": "0",
        "shortDesc": "Qualité et/ou titre de la personne bénéficiaire de la prestation",
        "type": "Alphanumérique"
    },
    "12": {
        "id": "12",
        "max": "38",
        "min": "0",
        "shortDesc": "Prénom de la personne bénéficiaire de la prestation",
        "type": "Alphanumérique"
    },
    "13": {
        "id": "13",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom de la personne bénéficiaire de la prestation",
        "type": "Alphanumérique"
    },
    "14": {
        "id": "14",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 1 de la norme adresse postale du destinataire de la facture",
        "type": "Alphanumérique"
    },
    "15": {
        "id": "15",
        "max": "38",
        "min": "0",
        "shortDesc": "Qualité et/ou titre de la personne destinataire de la facture",
        "type": "Alphanumérique"
    },
    "16": {
        "id": "16",
        "max": "38",
        "min": "0",
        "shortDesc": "Prénom de la personne destinataire de la facture",
        "type": "Alphanumérique"
    },
    "17": {
        "id": "17",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom de la personne destinataire de la facture",
        "type": "Alphanumérique"
    },
    "18": {
        "id": "18",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Numéro de la facture",
        "type": "Alphanumérique"
    },
    "19": {
        "id": "19",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Numéro de client",
        "type": "Alphanumérique"
    },
    "1A": {
        "id": "1A",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Numéro du contrat",
        "type": "Alphanumérique"
    },
    "1B": {
        "id": "1B",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Identifiant du souscripteur du contrat",
        "type": "Alphanumérique"
    },
    "1C": {
        "id": "1C",
        "max": "8",
        "min": "8",
        "shortDesc": "Date d’effet du contrat",
        "type": "Numérique"
    },
    "1D": {
        "id": "1D",
        "max": "16",
        "min": "0",
        "shortDesc": "Montant TTC de la facture",
        "type": "Numérique"
    },
    "1E": {
        "id": "1E",
        "max": "30",
        "min": "0",
        "shortDesc": "Numéro de téléphone du bénéficiaire de la prestation",
        "type": "Numérique"
    },
    "1F": {
        "id": "1F",
        "max": "30",
        "min": "0",
        "shortDesc": "Numéro de téléphone du destinataire de la facture",
        "type": "Numérique"
    },
    "1G": {
        "id": "1G",
        "max": "1",
        "min": "1",
        "shortDesc": "Présence d’un co-bénéficiaire de la prestation non mentionné dans le code",
        "type": "Numérique"
    },
    "1H": {
        "id": "1H",
        "max": "1",
        "min": "1",
        "shortDesc": "Présence d’un co-destinataire de la facture non mentionné dans le code",
        "type": "Numérique"
    },
    "1I": {
        "id": "1I",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 1 de la norme adresse postale du co-bénéficiaire de la prestation.",
        "type": "Alphanumérique"
    },
    "1J": {
        "id": "1J",
        "max": "38",
        "min": "0",
        "shortDesc": "Qualité et/ou titre du co-bénéficiaire de la prestation.",
        "type": "Alphanumérique"
    },
    "1K": {
        "id": "1K",
        "max": "38",
        "min": "0",
        "shortDesc": "Prénom du co-bénéficiaire de la prestation.",
        "type": "Alphanumérique"
    },
    "1L": {
        "id": "1L",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom du co-bénéficiaire de la prestation.",
        "type": "Alphanumérique"
    },
    "1M": {
        "id": "1M",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 1 de la norme adresse postale du co-destinataire de la facture.",
        "type": "Alphanumérique"
    },
    "1N": {
        "id": "1N",
        "max": "38",
        "min": "0",
        "shortDesc": "Qualité et/ou titre du co-destinataire de la facture.",
        "type": "Alphanumérique"
    },
    "1O": {
        "id": "1O",
        "max": "38",
        "min": "0",
        "shortDesc": "Prénom du co-destinataire de la facture.",
        "type": "Alphanumérique"
    },
    "1P": {
        "id": "1P",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom du co-destinataire de la facture.",
        "type": "Alphanumérique"
    },
    "20": {
        "id": "20",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 2 de la norme adresse postale du point de service des prestations",
        "type": "Alphanumérique"
    },
    "21": {
        "id": "21",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 3 de la norme adresse postale du point de service des prestations",
        "type": "Alphanumérique"
    },
    "22": {
        "id": "22",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 4 de la norme adresse postale du point de service des prestations",
        "type": "Alphanumérique"
    },
    "23": {
        "id": "23",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 5 de la norme adresse postale du point de service des prestations",
        "type": "Alphanumérique"
    },
    "24": {
        "id": "24",
        "max": "5",
        "min": "5",
        "shortDesc": "Code postal ou code cedex du point de service des prestations",
        "type": "Numérique"
    },
    "25": {
        "id": "25",
        "max": "32",
        "min": "0",
        "shortDesc": "Localité de destination ou libellé cedex du point de service des prestations",
        "type": "Numérique"
    },
    "26": {
        "id": "26",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays de service des prestations",
        "type": "Alphanumérique"
    },
    "27": {
        "id": "27",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 2 de la norme adresse postale du destinataire de la facture",
        "type": "Alphanumérique"
    },
    "28": {
        "id": "28",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 3 de la norme adresse postale du destinataire de la facture",
        "type": "Alphanumérique"
    },
    "29": {
        "id": "29",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 4 de la norme adresse postale du destinataire de la facture",
        "type": "Alphanumérique"
    },
    "2A": {
        "id": "2A",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 5 de la norme adresse postale du destinataire de la facture",
        "type": "Alphanumérique"
    },
    "2B": {
        "id": "2B",
        "max": "5",
        "min": "5",
        "shortDesc": "Code postal ou code cedex du destinataire de la facture",
        "type": "Numérique"
    },
    "2C": {
        "id": "2C",
        "max": "32",
        "min": "0",
        "shortDesc": "Localité de destination ou libellé cedex du destinataire de la facture",
        "type": "Numérique"
    },
    "2D": {
        "id": "2D",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays du destinataire de la facture",
        "type": "Alphanumérique"
    },
    "30": {
        "id": "30",
        "max": "140",
        "min": "0",
        "shortDesc": "Qualité Nom et Prénom.",
        "type": "Alphanumérique"
    },
    "31": {
        "id": "31",
        "max": "38",
        "min": "14",
        "shortDesc": "Code IBAN",
        "type": "Alphanumérique"
    },
    "32": {
        "id": "32",
        "max": "11",
        "min": "8",
        "shortDesc": "Code BIC/SWIFT",
        "type": "Alphanumérique"
    },
    "33": {
        "id": "33",
        "max": "30",
        "min": "0",
        "shortDesc": "Code BBAN",
        "type": "Alphanumérique"
    },
    "34": {
        "id": "34",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays de localisation du compte",
        "type": "Alphanumérique"
    },
    "35": {
        "id": "35",
        "max": "34",
        "min": "14",
        "shortDesc": "Identifiant SEPAmail (QXBAN)",
        "type": "Alphanumérique"
    },
    "36": {
        "id": "36",
        "max": "4",
        "min": "4",
        "shortDesc": "Date de début de période",
        "type": "Alphanumérique"
    },
    "37": {
        "id": "37",
        "max": "4",
        "min": "4",
        "shortDesc": "Date de fin de période",
        "type": "Alphanumérique"
    },
    "38": {
        "id": "38",
        "max": "11",
        "min": "0",
        "shortDesc": "Solde compte début de période",
        "type": "Numérique"
    },
    "39": {
        "id": "39",
        "max": "11",
        "min": "0",
        "shortDesc": "Solde compte fin de période",
        "type": "Numérique"
    },
    "40": {
        "id": "40",
        "max": "13",
        "min": "13",
        "shortDesc": "Numéro fiscal",
        "type": "Numérique"
    },
    "41": {
        "id": "41",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Revenu fiscal de référence",
        "type": "Numérique"
    },
    "42": {
        "id": "42",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Situation du foyer",
        "type": "Alphanumérique"
    },
    "43": {
        "id": "43",
        "max": "Aucune",
        "min": "0",
        "shortDesc": "Nombre de parts",
        "type": "Numérique"
    },
    "44": {
        "id": "44",
        "max": "13",
        "min": "13",
        "shortDesc": "Référence d’avis d’impôt",
        "type": "Alphanumérique"
    },
    "50": {
        "id": "50",
        "max": "14",
        "min": "14",
        "shortDesc": "SIRET de l’employeur",
        "type": "Numérique"
    },
    "51": {
        "id": "51",
        "max": "6",
        "min": "6",
        "shortDesc": "Nombre d’heures travaillées",
        "type": "Numérique"
    },
    "52": {
        "id": "52",
        "max": "7",
        "min": "7",
        "shortDesc": "Cumul du nombre d’heures travaillées",
        "type": "Numérique"
    },
    "53": {
        "id": "53",
        "max": "4",
        "min": "4",
        "shortDesc": "Début de période",
        "type": "Alphanumérique"
    },
    "54": {
        "id": "54",
        "max": "4",
        "min": "4",
        "shortDesc": "Fin de période",
        "type": "Alphanumérique"
    },
    "55": {
        "id": "55",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de début de contrat",
        "type": "Numérique"
    },
    "56": {
        "id": "56",
        "max": "4",
        "min": "4",
        "shortDesc": "Date de fin de contrat",
        "type": "Alphanumérique"
    },
    "57": {
        "id": "57",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de signature du contrat",
        "type": "Numérique"
    },
    "58": {
        "id": "58",
        "max": "11",
        "min": "0",
        "shortDesc": "Salaire net imposable",
        "type": "Numérique"
    },
    "59": {
        "id": "59",
        "max": "12",
        "min": "0",
        "shortDesc": "Cumul du salaire net imposable",
        "type": "Numérique"
    },
    "5A": {
        "id": "5A",
        "max": "11",
        "min": "0",
        "shortDesc": "Salaire brut du mois",
        "type": "Numérique"
    },
    "5B": {
        "id": "5B",
        "max": "12",
        "min": "0",
        "shortDesc": "Cumul du salaire brut",
        "type": "Numérique"
    },
    "5C": {
        "id": "5C",
        "max": "11",
        "min": "0",
        "shortDesc": "Salaire net",
        "type": "Numérique"
    },
    "5D": {
        "id": "5D",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 2 de la norme adresse postale de l’employeur",
        "type": "Alphanumérique"
    },
    "5E": {
        "id": "5E",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 3 de la norme adresse postale de l’employeur",
        "type": "Alphanumérique"
    },
    "5F": {
        "id": "5F",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 4 de la norme adresse postale de l’employeur",
        "type": "Alphanumérique"
    },
    "5G": {
        "id": "5G",
        "max": "38",
        "min": "0",
        "shortDesc": "Ligne 5 de la norme adresse postale de l’employeur",
        "type": "Alphanumérique"
    },
    "5H": {
        "id": "5H",
        "max": "5",
        "min": "5",
        "shortDesc": "Code postal ou code cedex de l’employeur",
        "type": "Numérique"
    },
    "5I": {
        "id": "5I",
        "max": "32",
        "min": "0",
        "shortDesc": "Localité de destination ou libellé cedex de l’employeur",
        "type": "Alphanumérique"
    },
    "5J": {
        "id": "5J",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays de l’employeur",
        "type": "Alphanumérique"
    },
    "5K": {
        "id": "5K",
        "max": "50",
        "min": "0",
        "shortDesc": "Identifiant Cotisant Prestations Sociales",
        "type": "Alphanumérique"
    },
    "60": {
        "id": "60",
        "max": "60",
        "min": "0",
        "shortDesc": "Liste des prénoms",
        "type": "Alphanumérique"
    },
    "61": {
        "id": "61",
        "max": "20",
        "min": "0",
        "shortDesc": "Prénom",
        "type": "Alphanumérique"
    },
    "62": {
        "id": "62",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom patronymique",
        "type": "Alphanumérique"
    },
    "63": {
        "id": "63",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom d’usage",
        "type": "Alphanumérique"
    },
    "64": {
        "id": "64",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom d’épouse/époux",
        "type": "Alphanumérique"
    },
    "65": {
        "id": "65",
        "max": "2",
        "min": "2",
        "shortDesc": "Type de pièce d’identité",
        "type": "de pièce d’identitéTaille Min. 2 Taille Max. 2 Type Alphanumérique"
    },
    "66": {
        "id": "66",
        "max": "20",
        "min": "0",
        "shortDesc": "Numéro de la pièce d’identité",
        "type": "Alphanumérique"
    },
    "67": {
        "id": "67",
        "max": "2",
        "min": "2",
        "shortDesc": "Nationalité",
        "type": "Alphanumérique"
    },
    "68": {
        "id": "68",
        "max": "1",
        "min": "1",
        "shortDesc": "Genre",
        "type": "Alphanumérique"
    },
    "69": {
        "id": "69",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de naissance",
        "type": "Numérique"
    },
    "6A": {
        "id": "6A",
        "max": "32",
        "min": "0",
        "shortDesc": "Lieu de naissance",
        "type": "Alphanumérique"
    },
    "6B": {
        "id": "6B",
        "max": "3",
        "min": "3",
        "shortDesc": "Département du bureau émetteur",
        "type": "Alphanumérique"
    },
    "6C": {
        "id": "6C",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays de naissance",
        "type": "Alphanumérique"
    },
    "6D": {
        "id": "6D",
        "max": "60",
        "min": "0",
        "shortDesc": "Nom et prénom du père",
        "type": "Alphanumérique"
    },
    "6E": {
        "id": "6E",
        "max": "60",
        "min": "0",
        "shortDesc": "Nom et prénom de la mère",
        "type": "Alphanumérique"
    },
    "6F": {
        "id": "6F",
        "max": "90",
        "min": "0",
        "shortDesc": "Machine Readable Zone (Zone de Lecture Automatique, ZLA)",
        "type": "Alphanumérique"
    },
    "6G": {
        "id": "6G",
        "max": "38",
        "min": "1",
        "shortDesc": "Nom",
        "type": "Alphanumérique"
    },
    "6H": {
        "id": "6H",
        "max": "10",
        "min": "1",
        "shortDesc": "Civilité",
        "type": "Alphanumérique"
    },
    "6I": {
        "id": "6I",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays émetteur",
        "type": "Alphanumérique"
    },
    "70": {
        "id": "70",
        "max": "12",
        "min": "12",
        "shortDesc": "Date et heure du décès",
        "type": "Numérique"
    },
    "71": {
        "id": "71",
        "max": "12",
        "min": "12",
        "shortDesc": "Date et heure du constat de décès",
        "type": "Numérique"
    },
    "72": {
        "id": "72",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom du défunt",
        "type": "Alphanumérique"
    },
    "73": {
        "id": "73",
        "max": "60",
        "min": "0",
        "shortDesc": "Prénoms du défunt",
        "type": "Alphanumérique"
    },
    "74": {
        "id": "74",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom de jeune fille du défunt",
        "type": "Alphanumérique"
    },
    "75": {
        "id": "75",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de naissance du défunt",
        "type": "Numérique"
    },
    "76": {
        "id": "76",
        "max": "1",
        "min": "1",
        "shortDesc": "Genre du défunt",
        "type": "Alphanumérique"
    },
    "77": {
        "id": "77",
        "max": "45",
        "min": "0",
        "shortDesc": "Commune de décès",
        "type": "Alphanumérique"
    },
    "78": {
        "id": "78",
        "max": "5",
        "min": "5",
        "shortDesc": "Code postal de la commune de décès",
        "type": "Numérique"
    },
    "79": {
        "id": "79",
        "max": "114",
        "min": "0",
        "shortDesc": "Adresse du domicile du défunt",
        "type": "Alphanumérique"
    },
    "7A": {
        "id": "7A",
        "max": "5",
        "min": "5",
        "shortDesc": "Code postal du domicile du défunt",
        "type": "Numérique"
    },
    "7B": {
        "id": "7B",
        "max": "45",
        "min": "0",
        "shortDesc": "Commune du domicile du défunt",
        "type": "Alphanumérique"
    },
    "7C": {
        "id": "7C",
        "max": "1",
        "min": "1",
        "shortDesc": "Obstacle médico-légal",
        "type": "Numérique"
    },
    "7D": {
        "id": "7D",
        "max": "1",
        "min": "1",
        "shortDesc": "Mise en bière",
        "type": "Alphanumérique"
    },
    "7E": {
        "id": "7E",
        "max": "1",
        "min": "1",
        "shortDesc": "Obstacle aux soins de conservation",
        "type": "Numérique"
    },
    "7F": {
        "id": "7F",
        "max": "1",
        "min": "1",
        "shortDesc": "Obstacle aux dons du corps",
        "type": "Numérique"
    },
    "7G": {
        "id": "7G",
        "max": "1",
        "min": "1",
        "shortDesc": "Recherche de la cause du décès",
        "type": "Numérique"
    },
    "7H": {
        "id": "7H",
        "max": "2",
        "min": "2",
        "shortDesc": "Délai de transport du corps",
        "type": "Alphanumérique"
    },
    "7I": {
        "id": "7I",
        "max": "1",
        "min": "1",
        "shortDesc": "Prothèse avec pile",
        "type": "Numérique"
    },
    "7J": {
        "id": "7J",
        "max": "1",
        "min": "1",
        "shortDesc": "Retrait de la pile de prothèse",
        "type": "Numérique"
    },
    "7K": {
        "id": "7K",
        "max": "13",
        "min": "13",
        "shortDesc": "Code NNC",
        "type": "Alphanumérique"
    },
    "7L": {
        "id": "7L",
        "max": "9",
        "min": "9",
        "shortDesc": "Code Finess de l'organisme agréé",
        "type": "Alphanumérique"
    },
    "7M": {
        "id": "7M",
        "max": "64",
        "min": "0",
        "shortDesc": "Identification du médecin",
        "type": "Alphanumérique"
    },
    "7N": {
        "id": "7N",
        "max": "128",
        "min": "0",
        "shortDesc": "Lieu de validation du certificat de décès",
        "type": "Alphanumérique"
    },
    "7O": {
        "id": "7O",
        "max": "1",
        "min": "1",
        "shortDesc": "Certificat de décès supplémentaire",
        "type": "Numérique"
    },
    "80": {
        "id": "80",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom",
        "type": "Alphanumérique"
    },
    "81": {
        "id": "81",
        "max": "60",
        "min": "0",
        "shortDesc": "Prénoms",
        "type": "Alphanumérique"
    },
    "82": {
        "id": "82",
        "max": "20",
        "min": "0",
        "shortDesc": "Numéro de carte",
        "type": "Alphanumérique"
    },
    "83": {
        "id": "83",
        "max": "40",
        "min": "0",
        "shortDesc": "Organisme de tutelle",
        "type": "Alphanumérique"
    },
    "84": {
        "id": "84",
        "max": "40",
        "min": "0",
        "shortDesc": "Profession",
        "type": "Alphanumérique"
    },
    "90": {
        "id": "90",
        "max": "38",
        "min": "0",
        "shortDesc": "Identité de l'huissier de justice",
        "type": "Alphanumérique"
    },
    "91": {
        "id": "91",
        "max": "38",
        "min": "0",
        "shortDesc": "Identité ou raison sociale du demandeur",
        "type": "Alphanumérique"
    },
    "92": {
        "id": "92",
        "max": "38",
        "min": "0",
        "shortDesc": "Identité ou raison sociale du destinataire",
        "type": "Alphanumérique"
    },
    "93": {
        "id": "93",
        "max": "38",
        "min": "0",
        "shortDesc": "Identité ou raison sociale de tiers concerné",
        "type": "Alphanumérique"
    },
    "94": {
        "id": "94",
        "max": "38",
        "min": "0",
        "shortDesc": "Intitulé de l'acte",
        "type": "Alphanumérique"
    },
    "95": {
        "id": "95",
        "max": "18",
        "min": "0",
        "shortDesc": "Numéro de l'acte",
        "type": "Alphanumérique"
    },
    "96": {
        "id": "96",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de signature de l'acte",
        "type": "Numérique"
    },
    "A0": {
        "id": "A0",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays ayant émis l’immatriculation du véhicule.",
        "type": "Alphanumérique"
    },
    "A1": {
        "id": "A1",
        "max": "17",
        "min": "0",
        "shortDesc": "Immatriculation du véhicule",
        "type": "Alphanumérique"
    },
    "A2": {
        "id": "A2",
        "max": "17",
        "min": "0",
        "shortDesc": "Marque du véhicule.",
        "type": "Alphanumérique"
    },
    "A3": {
        "id": "A3",
        "max": "17",
        "min": "0",
        "shortDesc": "Nom commercial du véhicule.",
        "type": "Alphanumérique"
    },
    "A4": {
        "id": "A4",
        "max": "17",
        "min": "17",
        "shortDesc": "Numéro de série du véhicule (VIN).",
        "type": "Alphanumérique"
    },
    "A5": {
        "id": "A5",
        "max": "3",
        "min": "3",
        "shortDesc": "Catégorie du véhicule.",
        "type": "Alphanumérique"
    },
    "A6": {
        "id": "A6",
        "max": "2",
        "min": "2",
        "shortDesc": "Carburant.",
        "type": "Alphanumérique"
    },
    "A7": {
        "id": "A7",
        "max": "3",
        "min": "3",
        "shortDesc": "Taux d’émission de CO2 du véhicule (en g/km).",
        "type": "Alphanumérique"
    },
    "A8": {
        "id": "A8",
        "max": "12",
        "min": "0",
        "shortDesc": "Indication de la classe environnementale de réception CE.",
        "type": "Alphanumérique"
    },
    "A9": {
        "id": "A9",
        "max": "3",
        "min": "3",
        "shortDesc": "Classe d’émission polluante.",
        "type": "Alphanumérique"
    },
    "AA": {
        "id": "AA",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de première immatriculation du véhicule.",
        "type": "Numérique"
    },
    "AB": {
        "id": "AB",
        "max": "8",
        "min": "0",
        "shortDesc": "Type de lettre",
        "type": "de lettreTaille Min. 0 Taille Max. 8 Type Alphanumérique"
    },
    "AC": {
        "id": "AC",
        "max": "19",
        "min": "0",
        "shortDesc": "N° Dossier",
        "type": "Alphanumérique"
    },
    "AD": {
        "id": "AD",
        "max": "4",
        "min": "4",
        "shortDesc": "Date Infraction",
        "type": "Alphanumérique"
    },
    "AE": {
        "id": "AE",
        "max": "4",
        "min": "4",
        "shortDesc": "Heure de l’infraction",
        "type": "Numérique"
    },
    "AF": {
        "id": "AF",
        "max": "1",
        "min": "1",
        "shortDesc": "Nombre de points  retirés lors de l’infraction",
        "type": "Alphanumérique"
    },
    "AG": {
        "id": "AG",
        "max": "1",
        "min": "1",
        "shortDesc": "Solde de points",
        "type": "Alphanumérique"
    },
    "AH": {
        "id": "AH",
        "max": "30",
        "min": "0",
        "shortDesc": "Numéro de la carte",
        "type": "Alphanumérique"
    },
    "AI": {
        "id": "AI",
        "max": "8",
        "min": "8",
        "shortDesc": "Date d’expiration initiale",
        "type": "Numérique"
    },
    "AJ": {
        "id": "AJ",
        "max": "13",
        "min": "13",
        "shortDesc": "Numéro EVTC",
        "type": "Alphanumérique"
    },
    "AK": {
        "id": "AK",
        "max": "7",
        "min": "7",
        "shortDesc": "Numéro de macaron",
        "type": "Numérique"
    },
    "AL": {
        "id": "AL",
        "max": "11",
        "min": "11",
        "shortDesc": "Numéro de la carte",
        "type": "Alphanumérique"
    },
    "AM": {
        "id": "AM",
        "max": "5",
        "min": "0",
        "shortDesc": "Motif de sur-classement",
        "type": "Alphanumérique"
    },
    "B0": {
        "id": "B0",
        "max": "60",
        "min": "0",
        "shortDesc": "Liste des prénoms",
        "type": "Alphanumérique"
    },
    "B1": {
        "id": "B1",
        "max": "20",
        "min": "0",
        "shortDesc": "Prénom",
        "type": "Alphanumérique"
    },
    "B2": {
        "id": "B2",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom patronymique",
        "type": "Alphanumérique"
    },
    "B3": {
        "id": "B3",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom d'usage",
        "type": "Alphanumérique"
    },
    "B4": {
        "id": "B4",
        "max": "38",
        "min": "0",
        "shortDesc": "Nom d'épouse/époux",
        "type": "Alphanumérique"
    },
    "B5": {
        "id": "B5",
        "max": "2",
        "min": "2",
        "shortDesc": "Nationalité",
        "type": "Alphanumérique"
    },
    "B6": {
        "id": "B6",
        "max": "1",
        "min": "1",
        "shortDesc": "Genre",
        "type": "Alphanumérique"
    },
    "B7": {
        "id": "B7",
        "max": "8",
        "min": "8",
        "shortDesc": "Date de naissance",
        "type": "Numérique"
    },
    "B8": {
        "id": "B8",
        "max": "32",
        "min": "0",
        "shortDesc": "Lieu de naissance",
        "type": "Alphanumérique"
    },
    "B9": {
        "id": "B9",
        "max": "2",
        "min": "2",
        "shortDesc": "Pays de naissance",
        "type": "Alphanumérique"
    },
    "BA": {
        "id": "BA",
        "max": "1",
        "min": "1",
        "shortDesc": "Mention obtenue",
        "type": "Numérique"
    },
    "BB": {
        "id": "BB",
        "max": "50",
        "min": "0",
        "shortDesc": "Numéro ou code d'identification de l'étudiant",
        "type": "Alphanumérique"
    },
    "BC": {
        "id": "BC",
        "max": "20",
        "min": "0",
        "shortDesc": "Numéro du diplôme",
        "type": "Alphanumérique"
    },
    "BD": {
        "id": "BD",
        "max": "1",
        "min": "1",
        "shortDesc": "Niveau du diplôme selon la classification CEC",
        "type": "Numérique"
    },
    "BE": {
        "id": "BE",
        "max": "3",
        "min": "3",
        "shortDesc": "Crédits ECTS obtenus",
        "type": "Numérique"
    },
    "BF": {
        "id": "BF",
        "max": "3",
        "min": "3",
        "shortDesc": "Année universitaire",
        "type": "Numérique"
    },
    "BG": {
        "id": "BG",
        "max": "2",
        "min": "2",
        "shortDesc": "Type de diplôme",
        "type": "Alphanumérique"
    },
    "BH": {
        "id": "BH",
        "max": "30",
        "min": "0",
        "shortDesc": "Domaine",
        "type": "Alphanumérique"
    },
    "BI": {
        "id": "BI",
        "max": "30",
        "min": "0",
        "shortDesc": "Mention",
        "type": "Alphanumérique"
    },
    "BJ": {
        "id": "BJ",
        "max": "30",
        "min": "0",
        "shortDesc": "Spécialité",
        "type": "Alphanumérique"
    },
    "BK": {
        "id": "BK",
        "max": "14",
        "min": "14",
        "shortDesc": "Numéro de l'Attestation de versement de la CVE",
        "type": "Alphanumérique"
    }
}
`
package authority

type TrustServiceStatusList struct {
	Text              string   `xml:",chardata"`
	Tsl               string   `xml:"tsl,attr"`
	Ds                string   `xml:"ds,attr"`
	Ecc               string   `xml:"ecc,attr"`
	Tslx              string   `xml:"tslx,attr"`
	Xades             string   `xml:"xades,attr"`
	TSLTag            string   `xml:"TSLTag,attr"`
	SchemeInformation struct {
		Text                 string `xml:",chardata"`
		TSLVersionIdentifier string `xml:"TSLVersionIdentifier"`
		TSLSequenceNumber    string `xml:"TSLSequenceNumber"`
		TSLType              string `xml:"TSLType"`
		SchemeOperatorName   struct {
			Text string `xml:",chardata"`
			Name struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
		} `xml:"SchemeOperatorName"`
		SchemeOperatorAddress struct {
			Text            string `xml:",chardata"`
			PostalAddresses struct {
				Text          string `xml:",chardata"`
				PostalAddress []struct {
					Text            string `xml:",chardata"`
					Lang            string `xml:"lang,attr"`
					StreetAddress   string `xml:"StreetAddress"`
					Locality        string `xml:"Locality"`
					StateOrProvince string `xml:"StateOrProvince"`
					PostalCode      string `xml:"PostalCode"`
					CountryName     string `xml:"CountryName"`
				} `xml:"PostalAddress"`
			} `xml:"PostalAddresses"`
			ElectronicAddress struct {
				Text string `xml:",chardata"`
				URI  string `xml:"URI"`
			} `xml:"ElectronicAddress"`
		} `xml:"SchemeOperatorAddress"`
		SchemeName struct {
			Text string `xml:",chardata"`
			Name []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
		} `xml:"SchemeName"`
		SchemeInformationURI struct {
			Text string `xml:",chardata"`
			URI  []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"URI"`
		} `xml:"SchemeInformationURI"`
		StatusDeterminationApproach string `xml:"StatusDeterminationApproach"`
		SchemeTypeCommunityRules    struct {
			Text string   `xml:",chardata"`
			URI  []string `xml:"URI"`
		} `xml:"SchemeTypeCommunityRules"`
		SchemeTerritory     string `xml:"SchemeTerritory"`
		PolicyOrLegalNotice struct {
			Text           string `xml:",chardata"`
			TSLLegalNotice []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"TSLLegalNotice"`
		} `xml:"PolicyOrLegalNotice"`
		HistoricalInformationPeriod string `xml:"HistoricalInformationPeriod"`
		ListIssueDateTime           string `xml:"ListIssueDateTime"`
		NextUpdate                  struct {
			Text     string `xml:",chardata"`
			DateTime string `xml:"dateTime"`
		} `xml:"NextUpdate"`
		DistributionPoints struct {
			Text string `xml:",chardata"`
			URI  string `xml:"URI"`
		} `xml:"DistributionPoints"`
	} `xml:"SchemeInformation"`
	TrustServiceProviderList struct {
		Text                 string `xml:",chardata"`
		TrustServiceProvider []struct {
			Text           string `xml:",chardata"`
			TSPInformation struct {
				Text    string `xml:",chardata"`
				TSPName struct {
					Text string `xml:",chardata"`
					Name []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
				} `xml:"TSPName"`
				TSPTradeName struct {
					Text string `xml:",chardata"`
					Name []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
				} `xml:"TSPTradeName"`
				TSPAddress struct {
					Text            string `xml:",chardata"`
					PostalAddresses struct {
						Text          string `xml:",chardata"`
						PostalAddress []struct {
							Text            string `xml:",chardata"`
							Lang            string `xml:"lang,attr"`
							StreetAddress   string `xml:"StreetAddress"`
							Locality        string `xml:"Locality"`
							StateOrProvince string `xml:"StateOrProvince"`
							PostalCode      string `xml:"PostalCode"`
							CountryName     string `xml:"CountryName"`
						} `xml:"PostalAddress"`
					} `xml:"PostalAddresses"`
					ElectronicAddress struct {
						Text string `xml:",chardata"`
						URI  string `xml:"URI"`
					} `xml:"ElectronicAddress"`
				} `xml:"TSPAddress"`
				TSPInformationURI struct {
					Text string `xml:",chardata"`
					URI  []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"URI"`
				} `xml:"TSPInformationURI"`
			} `xml:"TSPInformation"`
			TSPServices struct {
				Text       string `xml:",chardata"`
				TSPService []struct {
					Text               string `xml:",chardata"`
					ServiceInformation struct {
						Text                  string `xml:",chardata"`
						ServiceTypeIdentifier string `xml:"ServiceTypeIdentifier"`
						ServiceName           struct {
							Text string `xml:",chardata"`
							Name []struct {
								Text string `xml:",chardata"`
								Lang string `xml:"lang,attr"`
							} `xml:"Name"`
						} `xml:"ServiceName"`
						ServiceDigitalIdentity struct {
							Text      string `xml:",chardata"`
							DigitalId struct {
								Text            string `xml:",chardata"`
								X509Certificate string `xml:"X509Certificate"`
							} `xml:"DigitalId"`
						} `xml:"ServiceDigitalIdentity"`
						ServiceStatus      string `xml:"ServiceStatus"`
						StatusStartingTime string `xml:"StatusStartingTime"`
					} `xml:"ServiceInformation"`
					ServiceHistory struct {
						Text                   string `xml:",chardata"`
						ServiceHistoryInstance struct {
							Text                  string `xml:",chardata"`
							ServiceTypeIdentifier string `xml:"ServiceTypeIdentifier"`
							ServiceName           struct {
								Text string `xml:",chardata"`
								Name []struct {
									Text string `xml:",chardata"`
									Lang string `xml:"lang,attr"`
								} `xml:"Name"`
							} `xml:"ServiceName"`
							ServiceDigitalIdentity struct {
								Text      string `xml:",chardata"`
								DigitalId struct {
									Text            string `xml:",chardata"`
									X509SubjectName string `xml:"X509SubjectName"`
								} `xml:"DigitalId"`
							} `xml:"ServiceDigitalIdentity"`
							ServiceStatus      string `xml:"ServiceStatus"`
							StatusStartingTime string `xml:"StatusStartingTime"`
						} `xml:"ServiceHistoryInstance"`
					} `xml:"ServiceHistory"`
				} `xml:"TSPService"`
			} `xml:"TSPServices"`
		} `xml:"TrustServiceProvider"`
	} `xml:"TrustServiceProviderList"`
	Signature struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"Id,attr"`
		Ds         string `xml:"ds,attr"`
		SignedInfo struct {
			Text                   string `xml:",chardata"`
			CanonicalizationMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"CanonicalizationMethod"`
			SignatureMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"SignatureMethod"`
			Reference []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"Id,attr"`
				URI        string `xml:"URI,attr"`
				Type       string `xml:"Type,attr"`
				Transforms struct {
					Text      string `xml:",chardata"`
					Transform []struct {
						Text      string `xml:",chardata"`
						Algorithm string `xml:"Algorithm,attr"`
						XPath     string `xml:"XPath"`
					} `xml:"Transform"`
				} `xml:"Transforms"`
				DigestMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"DigestMethod"`
				DigestValue string `xml:"DigestValue"`
			} `xml:"Reference"`
		} `xml:"SignedInfo"`
		SignatureValue struct {
			Text string `xml:",chardata"`
			ID   string `xml:"Id,attr"`
		} `xml:"SignatureValue"`
		KeyInfo struct {
			Text     string `xml:",chardata"`
			X509Data struct {
				Text            string `xml:",chardata"`
				X509Certificate string `xml:"X509Certificate"`
			} `xml:"X509Data"`
		} `xml:"KeyInfo"`
		Object struct {
			Text                 string `xml:",chardata"`
			QualifyingProperties struct {
				Text             string `xml:",chardata"`
				Target           string `xml:"Target,attr"`
				Ds               string `xml:"ds,attr"`
				Xades            string `xml:"xades,attr"`
				SignedProperties struct {
					Text                      string `xml:",chardata"`
					ID                        string `xml:"Id,attr"`
					SignedSignatureProperties struct {
						Text               string `xml:",chardata"`
						SigningTime        string `xml:"SigningTime"`
						SigningCertificate struct {
							Text string `xml:",chardata"`
							Cert struct {
								Text       string `xml:",chardata"`
								CertDigest struct {
									Text         string `xml:",chardata"`
									DigestMethod struct {
										Text      string `xml:",chardata"`
										Algorithm string `xml:"Algorithm,attr"`
									} `xml:"DigestMethod"`
									DigestValue string `xml:"DigestValue"`
								} `xml:"CertDigest"`
								IssuerSerial struct {
									Text             string `xml:",chardata"`
									X509IssuerName   string `xml:"X509IssuerName"`
									X509SerialNumber string `xml:"X509SerialNumber"`
								} `xml:"IssuerSerial"`
							} `xml:"Cert"`
						} `xml:"SigningCertificate"`
					} `xml:"SignedSignatureProperties"`
					SignedDataObjectProperties struct {
						Text             string `xml:",chardata"`
						DataObjectFormat struct {
							Text            string `xml:",chardata"`
							ObjectReference string `xml:"ObjectReference,attr"`
							MimeType        string `xml:"MimeType"`
						} `xml:"DataObjectFormat"`
					} `xml:"SignedDataObjectProperties"`
				} `xml:"SignedProperties"`
			} `xml:"QualifyingProperties"`
		} `xml:"Object"`
	} `xml:"Signature"`
} 
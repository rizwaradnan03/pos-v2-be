package enums

type DonationType string

const (
	DonationTypePENDING  DonationType = "PENDING"
	DonationTypeACCEPTED DonationType = "ACCEPTED"
	DonationTypeREJECTED DonationType = "REJECTED"
)

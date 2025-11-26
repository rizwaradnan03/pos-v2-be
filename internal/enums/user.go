package enums

type AccountType string

const (
	AccountTypeCREDENTIAL AccountType = "CREDENTIAL"
	AccountTypeOAUTH      AccountType = "OAUTH"
)

type AccountMemberTypeType string

const (
	AccountMemberTypeTypeINTERNAL AccountMemberTypeType = "INTERNAL"
	AccountMemberTypeTypeEXTERNAL AccountMemberTypeType = "EXTERNAL"
)

package enums

type CommissionBatchType string

const (
	CommissionBatchTypePENDING  CommissionBatchType = "PENDING"
	CommissionBatchTypeACCEPTED CommissionBatchType = "ACCEPTED"
	CommissionBatchTypeREJECTED CommissionBatchType = "REJECTED"
)

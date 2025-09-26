package enum

type RiskLevel int8

const (
	RiskLevelNULL   RiskLevel = 0
	RiskLevelLow    RiskLevel = 1
	RiskLevelMedium RiskLevel = 2
	RiskLevelHigh   RiskLevel = 3
)

var (
	RiskLevelMap = map[RiskLevel]string{
		RiskLevelNULL:   "-",
		RiskLevelLow:    "Low Risk",
		RiskLevelMedium: "Medium Risk",
		RiskLevelHigh:   "High Risk",
	}
)

func (s RiskLevel) String() string {
	if val, ok := RiskLevelMap[s]; ok {
		return val
	}

	return ""
}

type Status int8

const (
	StatusUnknown  Status = 9
	StatusRejected Status = -1
	StatusWaiting  Status = 0
	StatusAccepted Status = 1
)

var (
	StatusMap = map[Status]string{
		StatusUnknown:  "Unknown",
		StatusRejected: "Rejected",
		StatusWaiting:  "Waiting",
		StatusAccepted: "Accepted",
	}
	KeyStatusMap = map[string]Status{
		"Unknown":  StatusUnknown,
		"Rejected": StatusRejected,
		"Waiting":  StatusWaiting,
		"Accepted": StatusAccepted,
	}
)

func (s Status) String() string {
	if val, ok := StatusMap[s]; ok {
		return val
	}

	return ""
}

func ParseStatus(key string) Status {
	if val, ok := KeyStatusMap[key]; ok {
		return val
	}

	return StatusUnknown
}

type ApplicationStatus int8

const (
	ASUnknown ApplicationStatus = -1
	ASLegal   ApplicationStatus = 0
	ASManager ApplicationStatus = 1
)

var (
	ASMap = map[ApplicationStatus]string{
		ASUnknown: "Unknown",
		ASLegal:   "Legal Consil",
		ASManager: "Manager",
	}
	KeyASMap = map[string]ApplicationStatus{
		"Unknown": ASUnknown,
		"Legal":   ASLegal,
		"Manager": ASManager,
	}
)

func (s ApplicationStatus) String() string {
	if val, ok := ASMap[s]; ok {
		return val
	}

	return ""
}

func ParseAS(key string) ApplicationStatus {
	if val, ok := KeyASMap[key]; ok {
		return val
	}

	return ASUnknown
}

type Category int8

const (
	CategoryUnknown                  Category = -1
	CategoryProcurementContracts     Category = 0
	CategoryServiceContracts         Category = 1
	CategoryPartnershipMoU           Category = 2
	CategoryLicenseSoftwareAgreement Category = 3
	CategoryEmploymentHRContracts    Category = 4
)

var (
	CategoryMap = map[Category]string{
		CategoryUnknown:                  "unknown",
		CategoryProcurementContracts:     "Procurement Contracts",
		CategoryServiceContracts:         "Service Contracts",
		CategoryPartnershipMoU:           "Partnership/MoU",
		CategoryLicenseSoftwareAgreement: "License/Software Agreement",
		CategoryEmploymentHRContracts:    "Employment/HR Contracts",
	}
	KeyCategoryMap = map[string]Category{
		"procurement_contracts":      CategoryProcurementContracts,
		"service_contracts":          CategoryServiceContracts,
		"partnership_mou":            CategoryPartnershipMoU,
		"license_software_agreement": CategoryProcurementContracts,
		"employment_hr_contracts":    CategoryEmploymentHRContracts,
	}
)

func (s Category) String() string {
	if val, ok := CategoryMap[s]; ok {
		return val
	}

	return ""
}

func ParseCategory(key string) Category {
	if val, ok := KeyCategoryMap[key]; ok {
		return val
	}

	return CategoryUnknown
}

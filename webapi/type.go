package webapi

type ChaynsPaymentSites struct {
	AccountBalance           float64 `json:"accountBalance"`
	Dispo                    float64 `json:"dispo"`
	HasWorkingPaymentData    bool    `json:"hasWorkingPaymentData"`
	LocationID               int     `json:"locationId"`
	Name                     string  `json:"name"`
	PersonID                 string  `json:"personId"`
	SiteID                   string  `json:"siteId"`
	PaymentTarget            int     `json:"paymentTarget,omitempty"`
	PaymentTargetLocationIds []int   `json:"paymentTargetLocationIds,omitempty"`
}

type OpmChargeLeftAmount struct {
	LeftBankAmount            float64 `json:"leftBankAmount"`
	LimitedSepaChargeByPayPal bool    `json:"limitedSepaChargeByPayPal"`
	OriginalBankAmount        float64 `json:"originalBankAmount"`
}

type ChaynsUser struct {
	BirthDay             string `json:"birthDay"`
	CreationTime         string `json:"creationTime"`
	Email                string `json:"email"`
	FirstName            string `json:"firstName"`
	FullName             string `json:"fullName"`
	Gender               int    `json:"gender"`
	IDCardNumber         string `json:"idCardNumber"`
	IDNr                 string `json:"idNr"`
	Language             string `json:"language"`
	LastName             string `json:"lastName"`
	OpmType              int    `json:"opmType"`
	PersonalSite         string `json:"personalSite"`
	PersonID             string `json:"personId"`
	PhoneMobile          string `json:"phoneMobile"`
	SocialSecurityNumber string `json:"socialSecurityNumber"`
	Type                 string `json:"type"`
	VerificationState    int    `json:"verificationState"`
}

type OpmPersonDataPublicInformation struct {
	OpmType  int    `json:"opmType"`
	PersonID string `json:"personId"`
}

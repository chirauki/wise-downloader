package wise

type Profile struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Details struct {
		FirstName       string      `json:"firstName"`
		LastName        string      `json:"lastName"`
		DateOfBirth     string      `json:"dateOfBirth"`
		PhoneNumber     string      `json:"phoneNumber"`
		Avatar          string      `json:"avatar"`
		Occupation      string      `json:"occupation"`
		Occupations     interface{} `json:"occupations"`
		PrimaryAddress  interface{} `json:"primaryAddress"`
		FirstNameInKana interface{} `json:"firstNameInKana"`
		LastNameInKana  interface{} `json:"lastNameInKana"`
	} `json:"details"`
}

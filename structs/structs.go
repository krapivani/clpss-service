package structs

type User struct {
	UserID       string `json:"user_id"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	Phone        string `json:"phone"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Country      string `json:"country"`
}

type UserResponse struct {
	UserID       string `json:"user_id"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Country      string `json:"country"`
}

type HttpResponse struct {
	Profile interface{} `json:"profile"`
	Status  string      `json:"status"`
}

type Str struct {
	S string `json:"S"`
}

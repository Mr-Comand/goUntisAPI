package structs

// AuthParams represents the parameters for the authenticate method.
type AuthParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Client   string `json:"client"`
}

// AuthResponse represents the result of the authenticate method.
type AuthResponse struct {
	SessionID  string `json:"sessionId"`
	PersonType int    `json:"personType"`
	PersonID   int    `json:"personId"`
}

type ApiConfig struct {
	Server    string `json:"server"`
	User      string `json:"user"`
	Password  string `json:"password"`
	School    string `json:"school"`
	Useragent string `json:"ClientName"`
}

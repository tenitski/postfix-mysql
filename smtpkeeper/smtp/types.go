package smtp

type User struct {
	Login    string   `json:"login"`
	Password password `json:"password"`
}

type password string

// MarshalJSON ignores the field value completely.
func (password) MarshalJSON() ([]byte, error) {
	return []byte(`"**REDACTED**"`), nil
}

// String ignores the field value completely.
func (password) String() string {
	return "**REDACTED**"
}

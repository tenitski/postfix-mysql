package smtp

type User struct {
	Email    string   `json:"email"`
	Password password `json:"password"`
}

type Sender struct {
	//ID   string `json:"id"`
	User string `json:"string"`
	From string `json:"from"`
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

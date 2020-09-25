package apimedic

import (
	"encoding/json"
)

//Token is token returned from apimedic authorization service
type Token struct {
	Value        string `json:"Token,omitempty"`
	ValidThrough int    `json:"ValidThrough,omitempty"`
}

func DecodeToken(s string) (*Token, error) {
	var t Token
	if err := json.Unmarshal([]byte(s), &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *Token) String() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

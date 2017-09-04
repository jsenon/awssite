package lambda

import (
)

type Credential struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	TokenID string `json:TokenID`
}


type lambda struct {
	URL string `json:"URL"`
}




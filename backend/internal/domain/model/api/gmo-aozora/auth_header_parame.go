package model

type AuthHeader struct {
	AccessToken string
}

func NewAuthHeaderRequest(accessToken string) AuthHeader {
	return AuthHeader{
		AccessToken: accessToken,
	}
}

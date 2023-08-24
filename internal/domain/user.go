package domain

type User struct {
	Guid         string `bson:"guid" json:"guid"`
	RefreshToken string `bson:"refresh_token" json:"refreshtoken"`
}

type Token struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

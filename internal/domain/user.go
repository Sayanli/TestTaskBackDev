package domain

type User struct {
	Guid         string `bson:"guid"`
	RefreshToken string `bson:"refresh_token"`
}

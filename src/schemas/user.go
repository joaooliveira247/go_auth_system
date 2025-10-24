package schemas

type UserSchemaIn struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"-"`
}

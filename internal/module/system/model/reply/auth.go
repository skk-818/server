package reply

type LoginReply struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

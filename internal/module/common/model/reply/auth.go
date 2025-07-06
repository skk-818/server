package reply

type LoginReply struct {
	AccessToken  string `json:"accessToken"`
	Expires      int64  `json:"expires"`
	RefreshToken string `json:"refreshToken"`
}

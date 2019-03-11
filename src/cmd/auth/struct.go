package auth

type ResponseAuth struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

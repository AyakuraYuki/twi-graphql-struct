package model

type GraphUser struct {
	Data   *UserData   `json:"data"`
	Errors []*ErrorRsp `json:"errors"`
}

type UserData struct {
	User *UserResult `json:"user"`
}

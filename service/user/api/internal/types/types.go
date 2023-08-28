// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     int    `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Name string `json:"name"`
	Ok   string `json:"ok"`
}

type UserInfoResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
	Identity string `json:"identity"`
	Need     string `json:"need"`
}
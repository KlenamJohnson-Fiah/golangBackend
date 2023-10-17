package model

// type TokenDetails struct {
// 	AccessToken  string `json:"acess_token"`
// 	RefreshToken string `json:"refresh_token"`
// 	AccessUuid   string `json:"access_uuid"`
// 	RefreshUuid  string `json:"refresh_uuid"`
// 	AtExpires    int64  `json:"at_expires"`
// 	RtExpires    int64  `json:"rt_expires"`
// }

type TokenDetails struct {
	AccessTokenDetails  AccessTokenDetails
	RefreshTokenDetails RefreshTokenDetails
}

type AccessTokenDetails struct {
	AccessUuid  string `json:"access_uuid"`
	AccessToken string `json:"acess_token"`
	AtExpires   int64  `json:"at_expires"`
}

type RefreshTokenDetails struct {
	RefreshUuid  string `json:"refresh_uuid"`
	RefreshToken string `json:"refresh_token"`
	RtExpires    int64  `json:"rt_expires"`
}

type Login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginSucessful struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Could have used CustomerLogin and CustomerLoginSuccessful for this but
//was trying to avoid making the change in all packages that used them.
//Just Saving time

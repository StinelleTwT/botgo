package dto

// User 用户
type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	Bot              bool   `json:"bot"`
	UnionOpenID      string `json:"union_openid"`       // 特殊关联应用的 openid
	UserOpenId       string `json:"member_openid"`      // 特殊关联应用的 openid(群聊/私聊场景)
	UnionUserAccount string `json:"union_user_account"` // 机器人关联的用户信息，与 union_openid 关联的应用是同一个
	IsYou            bool   `json:"is_you"`             // 适用于群全量消息，用来标识是否为当前机器人
	Scope            string `json:"scope"`              // 适用于群全量消息，标识 AT 的范围
}

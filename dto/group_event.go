package dto

// GroupAddOrDelRobot 群添加/移除机器人信息
type GroupAddOrDelRobot struct {
	GroupOpenid    string `json:"group_openid"`
	OpMemberOpenid string `json:"op_member_openid"`
	Timestamp      int    `json:"timestamp"`
}

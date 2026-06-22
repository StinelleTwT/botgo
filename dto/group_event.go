package dto

// GroupAddOrDelRobot 群添加/移除机器人信息
type GroupAddOrDelRobot struct {
	GroupOpenid    string `json:"group_openid"`
	OpMemberOpenid string `json:"op_member_openid"`
	Timestamp      int    `json:"timestamp"`
}

// GroupMemberAddOrRemove 群添加/移除成员事件信息
type GroupMemberAddOrRemove struct {
	GroupOpenid  string `json:"group_openid"`
	MemberOpenid string `json:"member_openid"`
	Timestamp    int    `json:"timestamp"`
}

package dto

import "encoding/json"

// Interaction 互动行为对象
type Interaction struct {
	ID                string           `json:"id,omitempty"`                  // 互动行为唯一标识
	ApplicationID     string           `json:"application_id,omitempty"`      // 应用ID
	Type              InteractionType  `json:"type,omitempty"`                // 互动类型
	Data              *InteractionData `json:"data,omitempty"`                // 互动数据
	GuildID           string           `json:"guild_id,omitempty"`            // 频道 ID
	ChannelID         string           `json:"channel_id,omitempty"`          // 子频道 ID
	Version           uint32           `json:"version,omitempty"`             // 版本，默认为 1
	GroupOpenID       string           `json:"group_openid,omitempty"`        // 群OpenID
	ChatType          uint32           `json:"chat_type,omitempty"`           // 0: 频道, 1: 群, 2: c2c
	Scene             string           `json:"scene,omitempty"`               // 场景 c2c/group/guild
	UserOpenID        string           `json:"user_openid,omitempty"`         // 用户ID
	Timestamp         string           `json:"timestamp,omitempty"`           // 时间戳
	GroupMemberOpenID string           `json:"group_member_openid,omitempty"` // 群成员OpenID

}

// InteractionType 互动类型
type InteractionType uint32

const (
	// InteractionTypePing ping
	InteractionTypePing InteractionType = 1
	// InteractionTypeCommand 命令
	InteractionTypeCommand InteractionType = 2
	// InteractionTypeMsgButton 消息按钮
	InteractionTypeMsgButton InteractionType = 11
	// InteractionTypeC2CQuickMenu 单聊快捷菜单
	InteractionTypeC2CQuickMenu InteractionType = 12
	// InteractionTypeClawConfig Claw配置查询/更新事件
	InteractionTypeClawConfig InteractionType = 20
)

// InteractionData 互动数据
type InteractionData struct {
	Name     string              `json:"name,omitempty"`     // 标题
	Type     InteractionDataType `json:"type,omitempty"`     //	数据类型，不同数据类型对应不同的 resolved 数据
	Resolved json.RawMessage     `json:"resolved,omitempty"` // 跟不同的互动类型和数据类型有关系的数据
}

// InteractionDataType 互动数据类型
type InteractionDataType uint32

const (
	// InteractionDataTypeChatSearch 聊天框搜索
	InteractionDataTypeChatSearch InteractionDataType = 9
	// InteractionDataTypeInlineKeyboardClick 消息按钮点击
	InteractionDataTypeInlineKeyboardClick InteractionDataType = 11
	// InteractionDataTypeCallbackCommandClick C2C菜单点击
	InteractionDataTypeCallbackCommandClick InteractionDataType = 12
	// InteractionDataTypeMessageFeedbackClick 智能体消息反馈
	InteractionDataTypeMessageFeedbackClick InteractionDataType = 13
	// InteractionDataTypeClearSessionClick 清空会话按钮点击
	InteractionDataTypeClearSessionClick InteractionDataType = 14
	// InteractionDataTypeConfigQuery Claw 配置查询
	InteractionDataTypeConfigQuery InteractionDataType = 2001
	// InteractionDataTypeConfigUpdate Claw 配置更新
	InteractionDataTypeConfigUpdate InteractionDataType = 2002
)

// ResponseInteraction 回应互动消息接口的请求体
type ResponseInteraction struct {
	// Code 响应值：0 成功 1 操作失败 2 操作频繁 3 重复操作 4 没有权限 5 仅管理员操作
	Code uint32              `json:"code"`
	Data InteractionRespData `json:"data"`
}

type InteractionRespData struct {
	ClawCfg ClawConfig `json:"claw_cfg"`
}

type ClawConfig struct {
	// ChannelType 默认设置为"qqbot"
	ChannelType string `json:"channel_type"`
	ChannelVer  string `json:"channel_ver"`
	// ClawType 默认设置为"openclaw"
	ClawType string `json:"claw_type"`
	ClawVer  string `json:"claw_ver"`
	// RequireMention 群消息模式 "mention"=@机器人时激活 "always"=总是激活
	RequireMention string `json:"require_mention"`
	// GroupPolicy 群消息策略 open=全响应 | allowlist=白名单 | disabled=不响应
	GroupPolicy string `json:"group_policy"`
	// MentionPatterns @文本的名称提及BOT名，多个使用,分隔
	MentionPatterns string `json:"mention_patterns"`
	// OnlineState 在线状态 默认"online"
	OnlineState string `json:"online_state"`
}

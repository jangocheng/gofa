package entities

type GetUsersListStorage struct {
	GroupID string `json:"group_id"`
	Start   uint32 `json:"start"`
	Length  uint32 `json:"length"`
}

//{"error_code":0,"error_msg":"SUCCESS","log_id":304592829787243801,
// "timestamp":1542978724,"cached":0,"result":{"user_id_list":[]}}
type UserListResp struct {
	ErrorCode int32               `json:"error_code"`
	ErrorMsg  string              `json:"error_msg"`
	Result    map[string][]string `json:"result"`
}

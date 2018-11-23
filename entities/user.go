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

type AddFaceStorage struct {
	Image     string `json:"image"`
	ImageType string `json:"image_type"`
	GroupID   string `json:"group_id"`
	UserID    string `json:"user_id"`
	UserInfo  string `json:"user_info"`
}

type AddFaceResp struct {
	ErrorCode int32  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`

	Result struct {
		FaceToken string `json:"face_token"`
		Location  struct {
			Left     float32 `json:"left"`
			Top      float32 `json:"top"`
			Width    float32 `json:"width"`
			Height   float32 `json:"height"`
			Rotation float32 `json:"rotation"`
		} `json:"location"`
	} `json:"result"`
}

type FaceSearchStorage struct {
	Image       string `json:"image"`
	ImageType   string `json:"image_type"`
	GroupIDList string `json:"group_id_list"`
	MaxUserNum  uint32 `json:"max_user_num"`
}

type FaceSearchResp struct {
	ErrorCode int32  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`

	Result struct {
		FaceToken string `json:"face_token"`
		UserList  []struct {
			GroupID  string  `json:"group_id"`
			UserID   string  `json:"user_id"`
			UserInfo string  `json:"user_info"`
			Score    float32 `json:"score"`
		} `json:"user_list"`
	} `json:"result"`
}

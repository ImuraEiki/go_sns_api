package domain

type Following struct {
	Id         int `json:"id"`
	FollowId   int `json:"follow_id"`   // フォローする側のユーザーID
	FollowedId int `json:"followed_id"` // フォローされる側のユーザーID
}

package domain

type Following struct {
	ID         int `json:"id"`
	FollowID   int `json:"follow_id"`   // フォローする側のユーザーID
	FollowedID int `json:"followed_id"` // フォローされる側のユーザーID
}

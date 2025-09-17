package domain

type Following struct {
	Id             int `json:"id"`
	FollowUserId   int `json:"followUserId"`   // フォローする側のユーザーID
	FollowedUserId int `json:"followedUserId"` // フォローされる側のユーザーID
}

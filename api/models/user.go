package models

type User struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type UpdateUser struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"authors"`
}

package models

type Author struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthorPrimaryKey struct {
	Id string `json:"id"`
}

type CreateAuthor struct {
	Name string `json:"name"`
}

type UpdateAuthor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetListAuthorRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListAuthorResponse struct {
	Count   int       `json:"count"`
	Authors []*Author `json:"books"`
}

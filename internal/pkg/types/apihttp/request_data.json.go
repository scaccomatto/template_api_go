package apihttp

type CreateDataRequest struct {
	Name  string `json:"name" example:"mario"`
	Value int    `json:"value,omitempty" example:"1"`
}

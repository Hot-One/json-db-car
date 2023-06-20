package models

type CarPrimaryKey struct{
	Id string `json:"id"`
}

type CarCreate struct{
	Model string `json:"model"`
	Speed int `json:"speed"`
}

type Car struct{
	Id string `json:"id"`
	Model string `json:"model"`
	Speed int `json:"speed"`
}

type CarUpdate struct{
	Id string `json:"id"`
	Model string `json:"model"`
	Speed int `json:"speed"`
}

type CarGetListRequest struct {
	Offset int
	Limit  int
}

type CarGetListResponse struct {
	Count int
	Cars []*Car
}

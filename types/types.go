package types

type Todo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

type Response map[string]interface{}

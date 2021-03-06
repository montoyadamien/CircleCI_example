package entities

type Tree struct {
	Id		int		`json:"id" pg:",pk"`
	Name	string	`json:"name"`
}

package entities

type Plant struct {
	Id		int		`json:"id" pg:",pk"`
	Name	string	`json:"name"`
}

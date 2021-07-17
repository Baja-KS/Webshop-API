package database

type GroupOut struct {
	ID uint `json:"id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
}

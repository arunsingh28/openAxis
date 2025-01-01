package types

type Project struct {
	ID          int    `json:"id"`
	Name        string `validate:"required" json:"name"`
	Description string `json:"description" validate:"required"`
}

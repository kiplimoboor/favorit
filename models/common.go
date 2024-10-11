package models

type UpdateRequest struct {
	Field    string `json:"field"`
	NewValue string `json:"newVal"`
}

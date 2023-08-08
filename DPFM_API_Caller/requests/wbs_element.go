package requests

type WBSElement struct {
	Project             int   `json:"Project"`
	WBSElement          int   `json:"WBSElement"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

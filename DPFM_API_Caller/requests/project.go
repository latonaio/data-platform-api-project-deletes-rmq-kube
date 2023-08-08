package requests

type Project struct {
	Project             int   `json:"Project"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

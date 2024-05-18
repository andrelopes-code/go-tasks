package handler

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *CreateTaskRequest) validate() error {
	if r.Title == "" && r.Description == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}

func (r *UpdateTaskRequest) validate() error {
	if r.Title != "" || r.Description != "" || r.Done != nil {
		return nil
	}

	// if none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}

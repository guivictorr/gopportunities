package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Remote   *bool  `json:"remote"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Salary <= 0 && r.Remote == nil {
		return fmt.Errorf("request body is empty")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	return nil
}

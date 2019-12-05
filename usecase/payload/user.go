package payload

type CreateUserPayload struct {
	Name string `json:"name"`
}

type UpdateUserPayload struct {
	Name string `json:"name"`
}

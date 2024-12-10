package modle

type Requets struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (r *Requets) GetUsername() string {
	return r.Username
}

func (r *Requets) GetPassword() string {
	return r.Password
}

func (r *Response) GetMessage() string {
	return r.Message
}

func (r *Response) GetStatus() string {
	return r.Status
}

func (r *Response) SetMessage(message string) {
	r.Message = message
}

func (r *Response) SetStatus(status string) {
	r.Status = status
}

func NewRequest(username, password string) *Requets {
	return &Requets{
		Username: username,
		Password: password,
	}
}

func NewResponse(message, status string) *Response {
	return &Response{
		Message: message,
		Status:  status,
	}
}

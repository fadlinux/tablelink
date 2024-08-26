package user

type RequestLogin struct {
	Email    string `json:"status"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	RoleId   string `json:"role_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RequestUpdateUser struct {
	Name string `json:"name"`
}

type ResponseDefault struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
	Id         int    `json:"id"`
	RoleId     string `json:"role_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	LastAccess string `json:"last_access"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type ResponseLogin struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    ResponseLoginData `json:"data"`
}

type ResponseLoginData struct {
	RoleId      string `json:"role_id"`
	RoleName    string `json:"role_name"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	LastAccess  string `json:"last_access"`
	AccessToken string `json:"access_token"`
}

package handlers

//User Struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Response Generic
type Response struct {
	Message string `json:"message"`
}

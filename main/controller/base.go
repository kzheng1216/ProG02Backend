package controller

type (
	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	JwtResponse struct {
		Token string `json:"token"`
	}
)

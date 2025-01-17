package service

type (
	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	JwtResponse struct {
		Token string `json:"token"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

const (
	constUser123 = "user123"
	constPass123 = "pass123"
)

var secretKey = []byte("abcd1234") // 密钥，用于签名 JWT

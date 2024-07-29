package authentication

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/MatasGedziunas/Checkers.git/services/databaseFunctionality"
	"github.com/MatasGedziunas/Checkers.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "SuperSecretKeyFromEnvVariable"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Credentials) Validate() (string, bool) {
	return "", true
}

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	credentials, err := utils.ReadBody[Credentials](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRows, err := db.Query("SELECT * FROM users WHERE username = $1 AND password = $2", credentials.Username, utils.HashString(credentials.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem inserting record into database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	var users []databaseFunctionality.User
	for userRows.Next() {
		var user databaseFunctionality.User
		userRows.Scan(&user.UserId, &user.Username, &user.Password)
		users = append(users, user)
	}
	if len(users) != 1 {
		http.Error(w, "duplicate users for these credentials found", http.StatusInternalServerError)
		return
	}
	jwtToken, err := GetJwtToken(users[0].UserId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem getting jwt token: %s", err.Error()), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    jwtToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Successfully logged in"))
}

func GetJwtToken(user_id int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": user_id,
		"iss": "checkersBackend",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

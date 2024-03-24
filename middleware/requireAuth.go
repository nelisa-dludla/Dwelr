package middleware

import (
	"context"
	"dwelr/models"
	"dwelr/setup"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get cookie from request
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Decode/validate it
		tokenStr := cookie.Value
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Validate alg
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			// Check exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// Find user with token
			var user models.User
			setup.DB.Find(&user, claims["sub"])

			if user.ID == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// Attach to request
			r = r.WithContext(context.WithValue(r.Context(), "user", user))

			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}

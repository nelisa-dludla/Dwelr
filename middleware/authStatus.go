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

func AuthStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get cookie from request
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			fmt.Println("Failed to find cookie")
			w.Header().Set("X-Auth-Status", "Unauthenticated")
			next.ServeHTTP(w, r)
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
				fmt.Println("Token is expired")
				w.Header().Set("X-Auth-Status", "Unauthenticated")
				next.ServeHTTP(w, r)
				return
			}
			// Find user with token
			var user models.User
			setup.DB.Find(&user, claims["sub"])

			if user.ID == 0 {
				fmt.Println("Token doesn't match any user")
				w.Header().Set("X-Auth-Status", "Unauthenticated")
			
				fmt.Println("No user found #1")

				next.ServeHTTP(w, r)
				return
			} else {
				fmt.Println("User found")
				// Attach to request
				w.Header().Set("X-Auth-Status", "Authenticated")
				r = r.WithContext(context.WithValue(r.Context(), "user", user))
				next.ServeHTTP(w, r)
			}
		} else {
			fmt.Println("Token invalid")
			w.Header().Set("X-Auth-Status", "Unauthenticated")

			fmt.Println("No user found #2")

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			return
		}
	})
}

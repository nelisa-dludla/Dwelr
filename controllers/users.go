package controllers

import (
	"dwelr/models"
	"dwelr/setup"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	// Check if passwords match
	if password != confirmPassword {
		fmt.Fprintln(w, "Passwords do not match.")
		//http.Error(w, "Passwords do not match.", http.StatusBadRequest)
		return
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		http.Error(w, "Failed to hash password.", http.StatusInternalServerError)
		return
	}
	fmt.Println("Password was hashed")
	// Create user
	user := models.User{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: string(hashedPassword),
	}
	fmt.Println("User Model:")
	fmt.Println(user)
	result := setup.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Failed to create user.", http.StatusInternalServerError)
		return
	}
	fmt.Println("User was created")
	// Respond
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Retrieve email and password
	email := r.FormValue("email")
	password := r.FormValue("password")
	// Find user
	var user models.User
	setup.DB.Find(&user, "email = ?", email)

	if user.ID == 0 {
		fmt.Fprintf(w, "Incorrect email or password")
		//http.Error(w, "Incorrect email or password", http.StatusBadRequest)
		return
	}
	// Compare stored password with recieved password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Fprintf(w, "Incorrect email or password")
		//http.Error(w, "Incorrect email or password", http.StatusBadRequest)
		return
	}
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to create JWT token.", http.StatusInternalServerError)
		return
	}
	// Create a cookie
	cookie := &http.Cookie{
		Name: "Authorization",
		Value: tokenStr,
		Path: "/",
	}

	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)
	cookie.MaxAge = 3600 * 24 * 30
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode

	http.SetCookie(w, cookie)
	// Send to frontend
	//w.Header().Add("token", tokenStr)
	//fmt.Fprintf(w, "Logged in successfully")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// Create a cookie
	cookie := &http.Cookie{
		Name: "Authorization",
		Value: "",
		Path: "/",
	}

	cookie.MaxAge = -1
	cookie.Secure = false
	cookie.HttpOnly = true

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I'm logged in")
}

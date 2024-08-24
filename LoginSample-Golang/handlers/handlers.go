package handlers

import (
	"log"
	"net/http"

	"github.com/antonioSilvaBentoRodrigues/Login-GO/models"
	"github.com/antonioSilvaBentoRodrigues/Login-GO/utils"
	"github.com/thedevsaddam/renderer"
	"golang.org/x/crypto/bcrypt"
)

var render *renderer.Render

func init() {
	options := renderer.Options{
		ParseGlobPattern: "./static/*.html",
	}
	render = renderer.New(options)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		for _, user := range models.AllUsers {
			if email == user.Email {
				err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
				if err != nil {
					log.Println("Error logging")
					return
				} else {
					log.Println("Logged")

				}
			}
		}

	}

	render.HTML(w, http.StatusOK, "login", nil)

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		confirmEmail := r.FormValue("confirm-email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm-password")

		okEmail := utils.CheckValues(email, confirmEmail)
		okPassword := utils.CheckValues(password, confirmPassword)
		if okEmail && okPassword {
			hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
			newUser := models.User{
				Email:    email,
				Password: hash,
			}
			models.AllUsers = append(models.AllUsers, newUser)
			log.Println("Users:", models.AllUsers)
		} else {
			log.Println("Error, account not created")
		}

	}
	render.HTML(w, http.StatusOK, "signup", nil)
}

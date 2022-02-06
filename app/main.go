/*
first try of golang
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     uint16
	Money   int16
	Hobbies []string
}

func (user *User) getAllInfo() string {
	return fmt.Sprintf("User name: %s. Age: %d. Money: %d.", user.Name, user.Age, user.Money)
}

func landing_page(w http.ResponseWriter, r *http.Request) {
	u := User{Name: "Max", Age: 32, Money: 1000, Hobbies: []string{"Not die", "Sadly looking at window", "Crying"}}
	tmpl, err := template.ParseFiles("../templates/landing_page.html")

	//"comma ok" pattern
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, u)
}

func user_page(w http.ResponseWriter, r *http.Request) {
	u := User{Name: "Max", Age: 32, Money: 1000, Hobbies: []string{"Not die", "Sadly looking at window", "Crying"}}
	fmt.Fprint(w, u.getAllInfo())
}

func handleRequest() {
	http.HandleFunc("/landing", landing_page)
	http.HandleFunc("/user_list", user_page)
	http.ListenAndServe(":8081", nil)
}

func main() {
	handleRequest()
}

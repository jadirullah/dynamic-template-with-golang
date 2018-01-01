package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//menampilkan halaman home
func home(w http.ResponseWriter, r *http.Request) {

	//deklarasi variabel nilai title dan pesan untuk di tampilkan di file html nya
	info := make(map[string]string)
	info["title"] = "Ini Halaman Home"
	info["message"] = "Selamat datang di halaman home"

	//fungsi penggabungan file html header, footer, serta file body nya
	var templates = template.Must(template.ParseFiles("layout/header.html", "layout/footer.html", "views/index.html"))
	//fungsi untuk mengeksekusi dan pengirim variabel info ke data html
	templates.ExecuteTemplate(w, "content", info)
}

//menampilkan halaman about
func about(w http.ResponseWriter, r *http.Request) {
	//deklarasi variabel nilai title dan pesan untuk di tampilkan di file html nya
	info := make(map[string]string)
	info["title"] = "Ini Halaman About"
	info["message"] = "Selamat datang di halaman About Us"

	//fungsi penggabungan file html header, footer, serta file body nya
	var templates = template.Must(template.ParseFiles("layout/header.html", "layout/footer.html", "views/about.html"))

	//fungsi untuk mengeksekusi dan pengirim variabel info ke data html
	templates.ExecuteTemplate(w, "content", info)
}

//menampilkan halaman home
func contact(w http.ResponseWriter, r *http.Request) {
	//deklarasi variabel nilai title dan pesan untuk di tampilkan di file html nya
	info := make(map[string]string)
	info["title"] = "Ini Halaman Contact"
	info["message"] = "Selamat datang di halaman Contact"

	//fungsi penggabungan file html header, footer, serta file body nya
	var templates = template.Must(template.ParseFiles("layout/header.html", "layout/footer.html", "views/contact.html"))

	//fungsi untuk mengeksekusi dan pengirim variabel info ke data html
	templates.ExecuteTemplate(w, "content", info)
}

//first file eksekusi
func main() {
	//routers
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	//file static on assets

	//static assets direktory for css, js, image or more
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//run server active
	fmt.Println("Port Acive is :8080")
	http.ListenAndServe(":8080", nil)
}

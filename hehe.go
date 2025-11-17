package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		showForm(w)
		return
	}

	processForm(w, r)
}

func showForm(w http.ResponseWriter) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	<body>
		<h2>Simple Form</h2>
		<form method="POST">
			Name: <br>
			<input name="name" required><br><br>

			Email: <br>
			<input name="email" type="email" required><br><br>

			Message: <br>
			<textarea name="message" rows="4" cols="40"></textarea><br><br>

			<input type="reset">
			<input type="submit">
		</form>
	</body>
	</html>
	`)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w, `
	<html>
	<body>
		<h2 style="color: green;">Form Submitted!</h2>
		<p><b>Name:</b> %s</p>
		<p><b>Email:</b> %s</p>
		<p><b>Message:</b> %s</p>
		<br>
		<a href="/">Back</a>
	</body>
	</html>
	`,
		r.FormValue("name"),
		r.FormValue("email"),
		r.FormValue("message"),
	)
}

package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're logged in!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlStr)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/main", mainHandler)
	http.ListenAndServe(":8080", nil)
}

const htmlStr = `
<html>
<body>
<form action="https://mitreid.org/authorize" method="GET" target="_blank" accept-charset="UTF-8"
enctype="application/x-www-form-urlencoded" autocomplete="off" novalidate>
<input type="text" name="response_type" value="code">
<input type="text" name="scope" value="openid">
<input type="text" name="client_id" value="client">
<input type="text" name="redirect_uri" value="http://localhost:8080/main">
<input type="submit" value="Submit">
</form>
</body>
</html>
`

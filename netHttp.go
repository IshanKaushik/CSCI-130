package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<body>
	Testing Go capability For Http
	</body>
</html>`,
	)
}

func main() {
	http.HandleFunc("/playMe", hello)
	http.ListenAndServe(":8080", nil)
}

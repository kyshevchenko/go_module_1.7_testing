// Демо-сервер калькулятора: /calc + Swagger UI на /swagger/index.html.
//
//	@title		stringutil Calc API
//	@version	1.0
//	@description	HTTP-калькулятор учебного пакета stringutil.
//	@BasePath	/
package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"stringutil"
	_ "stringutil/docs" // сгенерированная swagger-документация
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /calc", stringutil.Handler)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	addr := ":8081"
	log.Printf("calc server listening on %s", addr)
	log.Printf("swagger UI: http://localhost%s/swagger/index.html", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

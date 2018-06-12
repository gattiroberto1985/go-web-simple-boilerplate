package webapi

import (
	"net/http"
	// application packages
	"it/bob/go/boilerplate/beans"
)

func SimpleBeanHandler(writer http.ResponseWriter, r *http.Request) {

	bean := beans.SimpleBean{}

	bean.SayHello()
}

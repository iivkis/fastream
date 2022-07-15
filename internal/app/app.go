package app

import httpapi "github.com/iivkis/fastream/internal/api/http"

func Launch() {
	APIHttp := httpapi.NewAPIHttp()

	if err := APIHttp.Run(":8080"); err != nil {
		panic(err)
	}
}

package api

import "github.com/AliKhedmati/routate-backend/src/api/routes"

func Init() {
	r := routes.NewRouter()

	if err := r.Run(); err != nil {
		defer func() {
			panic(err)
		}()
	}
}

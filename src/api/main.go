package api

import (
	"github.com/AliKhedmati/routate-backend/src/api/router"
)

func Init() (err error) {
	if err = router.NewRouter().Run(); err != nil {
		defer func() {
			panic(err)
		}()
	}

	return
}

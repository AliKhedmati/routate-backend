package api

func Init() {
	r := router.NewRouter()

	if err := r.Run(); err != nil {
		defer func() {
			panic(err)
		}()
	}
}

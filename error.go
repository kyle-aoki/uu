package uu

func MustExec(err error) {
	if err != nil {
		panic(err)
	}
}

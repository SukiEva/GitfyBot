package utils

func DropErr(e error) {
	if e != nil {
		panic(e)
	}
}

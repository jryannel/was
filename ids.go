package was

func MakeId() func() int {
	var id int
	return func() int {
		id++
		return id
	}
}

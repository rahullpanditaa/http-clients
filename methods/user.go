package methods

type User struct {
	Role       string
	ID         string
	Experience int
	Remote     bool
	User       struct {
		Name     string
		Location string
		Age      int
	}
}

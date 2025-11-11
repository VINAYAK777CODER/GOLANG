package main
import("fmt")
type Timestamps struct {
	createdAt string
	updatedAt string
}

type User struct {
	name  string
	email string
}

type Admin struct {
	User
	Timestamps
	role string
}

func main() {
	admin := Admin{
		User: User{name: "Amit", email: "amit@example.com"},
		Timestamps: Timestamps{
			createdAt: "2025-11-07",
			updatedAt: "2025-11-08",
		},
		role: "SuperAdmin",
	}

	// Access all promoted fields
	fmt.Println(admin.User)
	fmt.Println(admin.name)
	fmt.Println(admin.email)
	fmt.Println(admin.createdAt)
	fmt.Println(admin.updatedAt)
	fmt.Println(admin.role)
}

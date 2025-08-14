package databse

import (
	"github.com/gowaliullah/blog-rest-api/models"
)

var UserList []models.User

func Users() {
	user2 := models.User{ID: 1, Name: "Alice Smith", Email: "alice.s@example.com", Password: "hashed_password_1"}
	user3 := models.User{ID: 2, Name: "Bob Johnson", Email: "bob.j@example.com", Password: "hashed_password_2"}
	user4 := models.User{ID: 3, Name: "Charlie Brown", Email: "charlie.b@example.com", Password: "hashed_password_3"}
	user5 := models.User{ID: 4, Name: "Diana Prince", Email: "diana.p@example.com", Password: "hashed_password_4"}
	user6 := models.User{ID: 5, Name: "Edward Clark", Email: "edward.c@example.com", Password: "hashed_password_5"}

	UserList = append(UserList, user2, user3, user4, user5, user5, user6)
}

package users

import "net/http"

type userKey string

const UserCtx userKey = "user"

func GetUserFromContext(r *http.Request) *User {
	user, _ := r.Context().Value(UserCtx).(*User)

	return user
}

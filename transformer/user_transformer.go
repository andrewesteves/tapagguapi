package transformer

import "github.com/andrewesteves/tapagguapi/model"

// UserTransformer struct
type UserTransformer struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// TransformOne user specified JSON
func (ut UserTransformer) TransformOne(user model.User) UserTransformer {
	var newUser UserTransformer
	newUser.ID = user.ID
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Token = user.Token
	return newUser
}

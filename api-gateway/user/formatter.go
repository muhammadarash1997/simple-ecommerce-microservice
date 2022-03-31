package user

type UserRegisteredFormatter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UserLoggedFormatter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Token   string `json:"token"`
}

func FormatUserRegistered(user User) UserRegisteredFormatter {
	userFormatted := UserRegisteredFormatter{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	return userFormatted
}

func FormatUserLogged(user User, tokenGenerated string) UserLoggedFormatter {
	userFormatted := UserLoggedFormatter{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Token:   tokenGenerated,
	}

	return userFormatted
}

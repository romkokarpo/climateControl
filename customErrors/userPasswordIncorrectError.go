package customErrors

type UserPasswordIncorrectError struct {
	Message string
}

func (err UserPasswordIncorrectError) New() *UserPasswordIncorrectError {
	err.Message = "User password is incorrect"

	return &err
}

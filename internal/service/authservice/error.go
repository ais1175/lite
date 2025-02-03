package authservice

type ErrSessionExpired struct{}

func (ErrSessionExpired) Error() string {
	return "auth session expired"
}

func (ErrSessionExpired) Is(target error) bool {
	_, ok := target.(ErrSessionExpired)
	return ok
}

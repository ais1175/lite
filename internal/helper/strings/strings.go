package strings

func Null(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

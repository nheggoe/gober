package util

func MapError[A, B any](as []A, f func(A) (B, error)) ([]B, error) {
	errs := new(Errors)
	bs := make([]B, len(as))
	for i, a := range as {
		val, err := f(a)
		if err != nil {
			errs.Append(err)
		} else {
			bs[i] = val
		}
	}
	if !errs.IsEmpty() {
		return nil, errs
	}
	return bs, nil
}

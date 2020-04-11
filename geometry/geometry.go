package geometry

import "errors"

//Cube area
func Cube(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("Zore is not allowed")
	} else {
		return n * n * n, nil
	}

}

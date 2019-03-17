package rps

func UserIsValid(uName, pwd string) bool {
	_uName, _pwd, _isValid := "admin", "test", false
	if uName == _uName && pwd == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}
	return _isValid
}

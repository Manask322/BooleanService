package models


type BooleanError struct {
	Code int
	Status int
}


func GetErrorMessage(code int) string{
	switch code {
	case 1:
		return "User Not Found !! Enter Valid Token."
	case 2:
		return "Value is required Field and is Boolean."
	case 3:
		return "Requested Record Not Found for the User."
	case 4:
		return "Please Enter Valid Fields!! Value => Boolean and Key => String."
	case 5:
		return "Sorry !! There is a error from our Side Please Try Again Later."
	case 6:
		return "Key should be String."
	}
	return ""
}

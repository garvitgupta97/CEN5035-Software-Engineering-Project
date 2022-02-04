package func_judge

import (
	usr "backend/data"
)

//Judge whether there is a user
func IsExist(user string) bool {
	//If the length is 0, there is no user registration
	if len(usr.Slice) == 0 {
		return false
	} else {
		//Traversing slice
		for _, v := range usr.Slice {
			// return v.Name == user / / at this time, it can only be compared with the first one, so it is all false after the first one
			if v.Name == user {
				return true
			}
		}
	}
	return false
}

//Determine whether the password is correct
func IsRight(user string, passwd string) bool {
	for _, v := range usr.Slice {
		if v.Name == user {
			//Confirm that the names are the same first, and return true if the passwords are the same
			return v.Passwd == passwd
		}
	}
	return false
}

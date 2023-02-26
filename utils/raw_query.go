package utils

const (
	//USER
	SELECT_ALL_USER  = "SELECT * FROM users"
	INSERT_NEW_USER  = "INSERT INTO users (name, email, phonenumber, password, address, birthDate) VALUES ($1, $2, $3, $4, $5, $6)"
	UPDATE_USER_BYID = "UPDATE users SET name=$1, email=$2, phonenumber=$3, password=$4, address=$5, birthdate=$6"
)

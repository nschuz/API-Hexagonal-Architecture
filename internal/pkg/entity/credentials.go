package entity

//entity para el tipo de crendiciales es decir el login.
//si la infocion crece vemos la necesida de expandir creamos

// DefaultCredential represnets a user login
type DefaultCredential struct {
	Email    string
	Password string
}

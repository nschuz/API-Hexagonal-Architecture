package entity

import "time"

//se puede definir un tipo custom
// se utiliza ne godocs siempr ecomentar en typos , constante sy variables exportdas de un pauqte
//siempre temina conpunto todo lo ue se porta se docuenra

// User represent an aplication
type User struct {
	ID        int // ES MEJOR UTLIZAR EL UUID
	Name      string
	Lastname  string
	Email     string `gorm:"unique"`
	Password  string
	CratedAt  *time.Time //TO Know eh a user created his account
	DeletedAt *time.Time // TO KNOW eh a user deleted his account
}

// defineimo el schema de postgres
func (User) TableName() string {
	return "app.users"
}

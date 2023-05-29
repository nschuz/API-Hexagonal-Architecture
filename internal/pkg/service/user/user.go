package user

import (
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/ports"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	//tiene como atirbuto nuestro puerto
	repo ports.UserRepository
}

func NewService(repo ports.UserRepository) *service {
	return &service{
		repo: repo,
	} //crea una instancoa de servicio
}

// strign elregresa le token de la sesion de service
func (s *service) Login(credentials *entity.DefaultCredential) (string, error) {

	user := &entity.User{}

	//1 . obetenmos el usuario poremail
	//si existe se alacena en user
	if err := s.repo.First(user, "email = ?", credentials.Email); err != nil {
		return "", err
	}

	//2. Try match with passwords
	err := tryMatchPasswords(user.Password, credentials.Password)
	if err != nil {
		return "", err
	}

	//geenerate token for sesion
	return createToken(user)

}

func (s *service) Create(user *entity.User) error {
	//1. Valid email address
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	//Hashing el password de usuario
	//un prcoeso de hasing es de una via es sidento a los prcoeso de cifrados
	// los proceso hasing son de un via tengo un input a  un ouput
	//Salt temino psudoaleatorio agrega info pseudalteorio al hash porque hay ataques de diccionarios
	user.Password = hashAndSalt(user.Password)

	//3. save user
	return s.repo.Create(user)

}

// 2 a la potencia del costo
func hashAndSalt(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}
	return string(hash)
}

func tryMatchPasswords(userPassword, credentialPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentialPassword))
}

func createToken(user *entity.User) (string, error) {
	clains := jwt.MapClaims{}
	clains["userID"] = user.ID
	clains["userName"] = user.Name
	clains["user"] = user
	clains["exp"] = time.Hour * 24

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	//una lalve que solo el servir va teber
	strToken, err := jwtToken.SignedString([]byte("CLAVESUPERSEGURA"))

	if err != nil {
		return "", err
	}
	return strToken, nil
}

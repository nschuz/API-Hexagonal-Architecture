package firestore

import (
	"context"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type client struct {
	cli *firestore.Client

	//syncOnece que solo se ejecut solavez unicamemte para todos lco lciente
}

var (
	onceClientLoaded sync.Once
	firestoreCli     *client
)

// NewClient creates a new instance of the Firestore
func NewClient() *client {
	onceClientLoaded.Do(func() {
		//lo que queremos se ejcuta una solo vez es la conexion
		ctx := context.Background()
		opt := option.WithCredentialsJSON([]byte(
			`{
				"type": "service_account",
				"project_id": "chrome-sublime-387003",
				"private_key_id": "0b0f63bbb0a6253f197978bef7820d0742c708a5",
				"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCgmT44yc2kZzvu\nx/2u3UDDYHkAiXGweBLXSd91O3COloUI3TRCiVDz9elOYZOWX26d2AXYaaEmESj2\nty0wcEjdnXbYhMfJByUM3PcioGP1076GMdzm16mUvNTDjiuBnHoYBwBvvB1l4WQS\nxUhLEL8c9CI//5KIQMFRZMi652F86LlNXxLxsaXN9mMO5N8W1DpEED2sag9wy6dA\neHzryiTY2avN8NKaQZK9pWhzNbXV6RKP5Obc9+gncHIAs42iHi3tJIw3mQ4VjU+n\nB4y4QsyH4SN7YuBxMaMeyJXnTlIyzlJm+n5Qgt58a1uatYoFswnNYD5UMwCSfiBB\nvndmLSMpAgMBAAECggEAJY0ve1/UccihvueMbcLWTdjU8kcxCi5JWYWwdlz7qJuR\nZ80xVw18nE63ytgOBm8yEj5cv4mTYK9KgMW8D0NyBj+3dlE11c+R3jwcnUAbn3Nv\nbTljiaaLxHeS2Fb53UiVU+DvPTEjPptBJxJyIPs9dKuxBn3f6XQav8ltFibcKLRz\nKj6Q4z5AFekMKxyGPdxXqrrgTvUl/EahPKT4fHVQb3KcHIjf2FBvmHdBtvw59jrT\nZT4VrkoiyXsn1LIzan8m5dgJXpyLT62H+b0qY3klsvLSI7NNF+pVdswj5cm4R5TM\nVFu1C8nRQbQJEzSXg/cTU7zMdieAuUbgs1kTGNsAMQKBgQDN06EcXEmlRl0ds7Lr\n+df1HffLk+l/A5+hu0Zzv1CCF+rV5NQ0051aX3gBXn9K9zTgAxjZEITnz6qjg6Cy\nsMR/f6pKs3M+r/C67nz8GAjPZy7YKhs2pp9vdeX0NSS2mWCsXS1rZ5YqWXggqgD0\nkmZ77kbAr+fM/cZcTWqM8ympXwKBgQDHvzWXZaRaziUqBdmnS7fzgZnU5Y77dBx+\nSDc6qnlJZe+C3REu5XE8S56httw5Pvur1UaPGG4sDTt14iSx9+xEslLGyM9gIefU\namfhq/7DhkJY1ytzV+lzMJOZ+qOGmFdUwwjbdHh7Arg+5UAy8Y1APh7EOYLAGOak\ncYZktCGYdwKBgB1UigHbmNcHdt9zqwx7du3EDnqhIkXqQ0YtxLVbzuIq/FjsmdzJ\nRwQI1LYFPEh3f75oUpMkCnxqGFvitvSfwfCVRbQNbF/DfJ92urLFzgOYJZHCNkyM\nY+3jNBifthKbOq51PLKweTKhuz+UWjx/3EDOzKBKsNKCW1Dt60AYqpgrAoGAC7y9\nnm7XkIpqfqSnXIBDh+iGrI7srvPstLLzo9vekqSNxWfOGa5b0Ao1gRUorah4y2kV\n41SxLJ9+bffi5h5GYOcFnC6ymNiFMeMqxYUAzGZ4QZYrNHTm/+DL25FxyHftMcrg\nG10lszy+rNt1wUiWXz8HOcqTV0xfSSxulD5NdxMCgYBlZW/EcfkAUdhiLYnrmn46\nfCeTbrcBIUd8huU0MCZoF0bila6ehuxk2GW5TV8bdpBog9qG3xR2gEYJ7d4N8+M2\ndjc7QfWvyUbxFt91FU8+rOUMd0bZtFZ0S8Ay5HdxU3b43AHDhCzyfUFErOmhjPG5\nA3Ag4+78LyNNjvoWiCgZqA==\n-----END PRIVATE KEY-----\n",
				"client_email": "gonuapp@chrome-sublime-387003.iam.gserviceaccount.com",
				"client_id": "108059740633577699542",
				"auth_uri": "https://accounts.google.com/o/oauth2/auth",
				"token_uri": "https://oauth2.googleapis.com/token",
				"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
				"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/gonuapp%40chrome-sublime-387003.iam.gserviceaccount.com",
				"universe_domain": "googleapis.com"
			  }`,
		))
		c, err := firestore.NewClient(ctx, "chrome-sublime-387003", opt)
		if err != nil {
			panic("Failed to connect firestore client: " + err.Error())
		}
		firestoreCli = &client{
			cli: c,
		}
		log.Info("Connected to Firestore")
	})
	return firestoreCli
}

// AddDocument add a new Doument to given collection
func (c *client) AddDocument(collection string, data interface{}) error {
	_, _, err := c.cli.Collection(collection).Add(context.Background(), data)
	return err
}

// otra forinserat a una coleccion eespecifica el de arriba es genrica
func (c *client) AddExpense(expense *entity.Expense) error {
	_, _, err := c.cli.Collection("expenses").Add(context.Background(), expense)
	return err
}

func (c *client) GetUserExpenses(userID int, startDate time.Time, endDate time.Time) ([]*entity.Expense, error) {
	var expenses []*entity.Expense

	expensesCollection := c.cli.Collection("expenses")

	query := expensesCollection.Where("userID", "==", userID).Where("date", ">=", startDate).Where("date", "<=", endDate)
	iterador := query.Documents(context.Background())
	defer iterador.Stop()
	for {
		doc, err := iterador.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		expense := &entity.Expense{}
		doc.DataTo(expense) //deserializmaos cada uno de los docuemntos de firesto
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (c *client) AddDeposit(deposit *entity.Deposit) error {
	_, _, err := c.cli.Collection("deposits").Add(context.Background(), deposit)
	return err
}

func (c *client) GetUserDeposits(userID int, startDate time.Time, endDate time.Time) ([]*entity.Deposit, error) {

	var deposits []*entity.Deposit

	depositsCollection := c.cli.Collection("deposits")
	query := depositsCollection.Where("userID", "==", userID).Where("date", ">=", startDate).Where("date", "<=", endDate)

	iterador := query.Documents(context.Background())
	defer iterador.Stop()
	for {
		doc, err := iterador.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		deposit := &entity.Deposit{}
		doc.DataTo(deposit) //deserializmaos cada uno de los docuemntos de firesto
		deposits = append(deposits, deposit)
	}

	return deposits, nil
}

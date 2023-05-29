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
			`{CREDENTIALS_KEY}`,
		))
		c, err := firestore.NewClient(ctx, "GOOGLE_PROYECT_CLIENT", opt)
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

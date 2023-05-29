package entity

import "time"

type Deposit struct {
	UserID      int       `firestore:"userID"`
	Category    string    `firestore:"category"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Amount      float32   `firestore:"amount"`
	Date        time.Time `firestore:"date"`
}

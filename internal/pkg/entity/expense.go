package entity

import "time"

// Expense represent a single  movement at type 'expense'.
type Expense struct {
	UserID      int       `firestore:"userID"`
	Categoria   string    `firestore:"category"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Amount      float32   `firestore:"amount"`
	Date        time.Time `firestore:"date"`
}

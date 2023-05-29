package main

// si le ponemos _ le decimos que no lo vamos a ocuap
import (
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/api"
	_ "github.com/nschuz/go-arquitectura-hexagonal/internal/infra/repositories/postgres"
)

func main() {
	// c := firestore.NewClient()

	// startDate, _ := time.Parse("2006-01-02", "2023-05-01")
	// endDate, _ := time.Parse("2006-01-02", "2023-06-01")

	// c.AddDeposit(&entity.Deposit{
	// 	UserID: 11,
	// 	Date:   time.Now(),
	// })

	// deposit, err := c.GetUserDeposits(1, startDate, endDate)
	// fmt.Println(err)
	// fmt.Println(deposit)

	// userRepository := postgres.NewClient()
	// userService := user.NewService(userRepository)

	// err := userService.Create(&entity.User{
	// 	Name:     "Chuz",
	// 	Lastname: "Regis",
	// 	Email:    "chuz@gmail.com",
	// 	Password: "MyPassword",
	// })

	// token, err := userService.Login(&entity.DefaultCredential{
	// 	Email:    "chuz@gmail.com",
	// 	Password: "MyPassword",
	// })

	// fmt.Println("Token: ", token)
	// fmt.Println("My err: ", err)

	// repoFirestone := firestore.NewClient()
	// service := deposits.NewService(repoFirestone)
	// deposits, err := service.GetUserDeposits(11, "2023-05-10", "2023-05-19")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(deposits)

	// expenseRepository := firestore.NewClient()
	// serviceExpense := expense.NewService(expenseRepository)

	// err := serviceExpense.Add(&entity.Expense{
	// 	UserID:    21,
	// 	Name:      "Test",
	// 	Categoria: "TEST",
	// 	Date:      time.Now(),
	// })

	// if err != nil {
	// 	fmt.Println("ERROR:", err)
	// }

	// expenses, err := serviceExpense.GetUserExpenses(21, "2023-05-19", "2023-05-19")
	// fmt.Println("Error:", err)
	// fmt.Println(expenses)

	api.RunServer()
}

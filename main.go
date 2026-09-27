package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Account struct {
	ID       uint64
	Login    string
	Password string
	Balance  float64
}

const (
	DB_USERNAME = "postgres"
	DB_PASSWORD = "1234"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_NAME     = "testdb"
)

/* Process user input */
func sanitize_string(input string) string {
	/* 1. Check if input is a valid string (prepend \ to special characters to
	 * avoid SQLi, length to MAX of the smallest fields db can store)
	 */
	return ""
}

/* Try to insert new entry in users table */
func create_account(conn *pgx.Conn, login string, password string) error {
	if login == "" || password == "" {
		fmt.Println("Login or password is empty")
		os.Exit(1)
	}

	var existingLogin string
	err := conn.QueryRow(
		context.Background(),
		"SELECT login FROM users WHERE login = $1",
		login,
	).Scan(&existingLogin)

	if err == nil {
		fmt.Println("Такое имя занято")
		os.Exit(1)
	}

	if errors.Is(err, pgx.ErrNoRows) {

		_, err = conn.Exec(
			context.Background(),
			"INSERT INTO users (login, password, balance) VALUES ($1, $2, $3)",
			login,
			password,
			0.0,
		)

		if err != nil {
			fmt.Println("Erorr!")
			os.Exit(1)
		}
	}

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("Another err")
		os.Exit(1)
	}
	/* 1. Sanitize login and password
		 * 1. Check if login exists already; if true, return error
		 * 2. Insert new item, set balance to 0
	     Вроде все выполнено, но:
	     Я ебал пришлось сидеть читать и Posgres и pgx и разбираться с этим контекстком
	     и conn оказывается по функции передавать надод.
	     Короче:
	     1) Нужно решить что вместо os.Exit(1) писать - какие ошибки, чтобы обрабтывать
	     Мб создать структуру ошибок или типо такого?
	     2) Я не знаю правильно ли я все написал тут код создания поэтому перепроверь
	     3) Нужно ли проверять как в С++ длину логина и пароля? типо чтобы не обрезался, я хз
	     4) напиши в тг если хуйня
	*/
	return nil
}

/* Try to remove item by its id */
func remove_account(acc Account) error {
	/* 1. Query for item with id acc.ID
	 * 2. If found, remove, else return error
	 */
	return nil
}

/* Try to set login */
func set_login(acc Account, login string) error {
	/* 1. Check if item with acc.ID exists; if it doesn't, return error
	 * 3. Check if current login and provided login differ 4. If they differ,
	 * update; else return error
	 */
	return nil
}

/* Try to set password */
func set_password(acc Account, password string) error {
	/* 1. Check if item with acc.ID exists; if it doesn't, return error
	 * 2. Sanitize password
	 * 3. Update password
	 */
	return nil
}

/* Try to set balance */
func set_balance(ID uint64, value float64) error {
	/* 1. Check if item with acc.ID exists; if it doesn't, return error
	 * 2. Set balance
	 */
	return nil
}

/* Check if credentials match */
func validate_login(login string, password string) error {
	/* 1. Check if login exists and password matches; if one of these fail,
	 * return error
	 * 2. On success, return nil
	 */
	return nil
}

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := "postgres://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed connecting to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var login string
	var balance float64
	sql_query := "select login, balance from users"

	// Gets rows and stores them in rows var
	rows, _ := conn.Query(context.Background(), sql_query)
	// For each row in rows var, scan results into login and balance, then run func()
	_, err = pgx.ForEachRow(rows, []any{&login, &balance}, func() error {
		fmt.Fprintf(os.Stdout, "%s %f\n", login, balance)
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "ForEachRow failed: %v\n", err)
		os.Exit(1)
	}
}

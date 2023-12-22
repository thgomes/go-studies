package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/learn_go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	printCommands()

	var command string
	for {
		fmt.Scanln(&command)
		switch command {
		case "create-product":
			createProduct(db)
		case "list-products":
			listAllProducts(db)
		case "list-commands":
			printCommands()
		default:
			println("Comando inexistente.")
		}
	}
}

func printCommands() {
	commands := []string{"create-product", "delete-product", "view-product", "list-products", "list-commands"}
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("  - %s\n", cmd)
	}
	fmt.Println()
}

func createProduct(db *sql.DB) {
	var product Product

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Product Name: ")
	product.Name, _ = reader.ReadString('\n')

	fmt.Printf("Product Description: ")
	product.Description, _ = reader.ReadString('\n')

	fmt.Printf("Product Price: ")
	fmt.Scanf("%f\n", &product.Price)

	product.ID = uuid.New().String()

	stmt, err := db.Prepare("INSERT INTO products(id, name, description, price) VALUES(?, ?, ?, ?);")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Description, product.Price)
	if err != nil {
		panic(err)
	}
	println("Produto criado com sucesso!\n")
}

func listAllProducts(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM products;")
	if err != nil {
		panic(err)
	}

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Printf("ID: %s\nNome: %s\nDescricao: %s\n Preco: %f\n", p.ID, p.Name, p.Description, p.Price)
	}
}

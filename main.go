package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"git.enigmacamp.com/enigma-20/agnes-maria-anggelina/challenge-godb/entity"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"bufio"
	"os/exec"
	"runtime"
)

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "123"
	dbname = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database!")
	}

	return db
}

func consoleProgram() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(".....:::::Welcome to Enigma Laundry:::::.....")
	fmt.Println("\nPilih tabel yang ingin digunakan: ")
	fmt.Println("1. Customer")
	fmt.Println("2. Service")
	fmt.Println("3. Transaction")
	fmt.Println("[Ketik 1 untuk ke tabel Customer, 2 ke tabel Service, dst ..]")
	scanner.Scan()
	tabel_pilihan := scanner.Text()
	clearScreen()

	switch tabel_pilihan {
	case "1":
		processCustomerTable()
	case "2":
		processServiceTable()
	case "3":
		processTransactionTable()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func main() {	
	consoleProgram()
}

// TRANSACTION

func makeTransaction(transaction entity.Transaction) {
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	insertTransaction(transaction, tx)
	jumlahTransaksi, totalTransaksi := getSumANDTotalTransaction(transaction.CustomerId, tx)
	updateTransaction(jumlahTransaksi, totalTransaksi, transaction.CustomerId, tx)

	// Commit the transaction if everything is successful
	err = tx.Commit()
	if err != nil {
		// Rollback the transaction if there's an error during commit
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback!")
	} else {
		fmt.Println("TRANSACTION SUCCESS!")
	}
}

func insertTransaction(transaction entity.Transaction, tx *sql.Tx) {
	insertTransactionData := "INSERT INTO trx_transaction (customer_id, service_id, quantity, transaction_date) VALUES ($1, $2, $3, $4)"

	_, err := tx.Exec(insertTransactionData, transaction.CustomerId, transaction.ServiceId, transaction.Quantity, transaction.TransactionDate)
	validate(err, "Insert", tx)
}

func getSumANDTotalTransaction(customerID int, tx *sql.Tx) (int, int) {
	sumTransactionQuery := "SELECT COUNT(*) FROM trx_transaction WHERE customer_id = $1"
	totalTransactionQuery := "UPDATE trx_transaction AS t SET total_transaksi = m.price * t.quantity FROM mst_service AS m WHERE t.service_id = m.id AND t.customer_id = $1;"

	var sumTransaction, totalTransaction int
	err := tx.QueryRow(sumTransactionQuery, customerID).Scan(&sumTransaction)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case when no rows are returned
			sumTransaction = 0
		} else {
			// Handle other errors
			panic(err)
		}
	}

	err = tx.QueryRow(totalTransactionQuery, customerID).Scan(&totalTransaction)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case when no rows are returned
			totalTransaction = 0
		} else {
			// Handle other errors
			panic(err)
		}
	}

	return sumTransaction, totalTransaction
}


func updateTransaction(jumlahTransaksi, totalTransaksi, transactionID int, tx *sql.Tx) {
	updateSumTransaction := "UPDATE mst_customer SET jumlah_transaksi = $1 WHERE id = $2"
	updateTotalTransaction := "UPDATE trx_transaction SET total_transaksi = $1 WHERE id = $2"

	_, err := tx.Exec(updateSumTransaction, jumlahTransaksi, transactionID)
	validate(err, "Update", tx)

	_, err = tx.Exec(updateTotalTransaction, totalTransaksi, transactionID)
	validate(err, "Update", tx)
}


func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		panic(err)
	}
}


// VIEW ALL DATA WITH TABLE 

func getAllData(table string, headers []string, scanDest ...interface{}) {
	db := connectDb()
	defer db.Close()

	sqlStatement := fmt.Sprintf("SELECT * FROM %s;", table)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tableWriter := tablewriter.NewWriter(os.Stdout)
	tableWriter.SetHeader(headers)

	for rows.Next() {
		err := rows.Scan(scanDest...)
		if err != nil {
			panic(err)
		}

		rowValues := make([]string, len(scanDest))
		for i, dest := range scanDest {
			switch v := dest.(type) {
			case *int:
				rowValues[i] = strconv.Itoa(*v)
			case *string:
				rowValues[i] = *v
			default:
				rowValues[i] = fmt.Sprintf("%v", v)
			}
		}

		tableWriter.Append(rowValues)
	}

	tableWriter.Render()
}

func addData(table string, data interface{}) {
	db := connectDb()
	defer db.Close()
	var err error

	switch v := data.(type) {
	case entity.Customer:
		sqlStatement := fmt.Sprintf("INSERT INTO %s (customer_name, phone_number, jumlah_transaksi) VALUES ($1, $2, 0);", table)
		_, err = db.Exec(sqlStatement, v.CustomerName, v.PhoneNumber)
	case entity.Service:
		sqlStatement := fmt.Sprintf("INSERT INTO %s (id, service_name, price) VALUES ($1, $2, $3);", table)
		_, err = db.Exec(sqlStatement, v.Id, v.ServiceName, v.Price)
	default:
		panic("Unsupported data type")
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully inserted data!")
	}
}

func editData(table string, data interface{}){
	db := connectDb()
	defer db.Close()
	var err error

	switch v := data.(type) {	
		case entity.Customer:
			sqlStatement := fmt.Sprintf("UPDATE %s SET customer_name = $2, phone_number = $3 WHERE id = $1;", table)
			_, err = db.Exec(sqlStatement, v.Id, v.CustomerName, v.PhoneNumber)
		case entity.Service:
			sqlStatement := fmt.Sprintf("UPDATE %s SET service_name = $2, price = $3 WHERE id = $1;", table)
			_, err = db.Exec(sqlStatement, v.Id, v.ServiceName, v.Price)
		default:
			panic("Unsupported data type")
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully edit data!")
	}
}

func deleteData(table string, data interface{}) {
	db := connectDb()
	defer db.Close()
	var err error

	switch v := data.(type) {	
		case entity.Customer:
			sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", table)
			_, err = db.Exec(sqlStatement, v.Id)
		case entity.Service:
			sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", table)
			_, err = db.Exec(sqlStatement, v.Id)
		default:
			panic("Unsupported data type")
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully delete data!")
	}
}

func deleteCustomer(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan id customer yang ingin didelete: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid id. Please enter a valid number.")
		return
	}

	customer := entity.Customer{
		Id: id,
	}

	deleteData("mst_customer", customer)
}

func deleteService() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan id service yang ingin diupdate: ")
	scanner.Scan()
	id := scanner.Text()
	
	service := entity.Service{
		Id: id,
	}

	deleteData("mst_service", service)
}

func addServiceData() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Service Id:")
	scanner.Scan()
	id := scanner.Text()


	fmt.Println("Enter Service Name:")
	scanner.Scan()
	serviceName := scanner.Text()

	fmt.Println("Enter Price:")
	scanner.Scan()
	priceStr := scanner.Text()
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		fmt.Println("Invalid price. Please enter a valid number.")
		return
	}

	service := entity.Service{
		Id:          id, // Generate a unique ID using UUID
		ServiceName: serviceName,
		Price:       price,
	}

	addData("mst_service", service)
}

func addCustomerData() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Customer Name:")
	scanner.Scan()
	customer_name := scanner.Text()

	fmt.Println("Enter Phone Number:")
	scanner.Scan()
	phone_number := scanner.Text()

	customer := entity.Customer{
		CustomerName: customer_name,
		PhoneNumber: phone_number,
	}

	addData("mst_customer", customer)
}

func editCustomerData(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan id customer yang ingin diupdate: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid id. Please enter a valid number.")
		return
	}

	fmt.Println("Enter Customer Name:")
	scanner.Scan()
	customer_name := scanner.Text()

	fmt.Println("Enter Phone Number:")
	scanner.Scan()
	phone_number := scanner.Text()

	customer := entity.Customer{
		Id: id,
		CustomerName: customer_name,
		PhoneNumber: phone_number,
	}

	editData("mst_customer", customer)
}

func editServiceData(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan id service yang ingin diupdate: ")
	scanner.Scan()
	id := scanner.Text()
	
	fmt.Println("Enter Service Name:")
	scanner.Scan()
	service_name := scanner.Text()

	fmt.Println("Enter Price:")
	scanner.Scan()
	priceStr := scanner.Text()
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		fmt.Println("Invalid price. Please enter a valid number.")
		return
	}
	service := entity.Service{
		Id: id,
		ServiceName: service_name,
		Price: price,
	}

	editData("mst_service", service)
}


func processCustomerTable() {
	dml := masterTable()
	switch dml {
	case "1":
		tabel := "mst_customer"
		headers := []string{"Id", "Customer Name", "Phone Number", "Jumlah Transaksi"}
		scanDest := []interface{}{new(int), new(string), new(string), new(int)}
		getAllData(tabel, headers, scanDest...)
	case "2":
		addCustomerData()
	case "3":
		editCustomerData()
	case "4":
		deleteCustomer()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func processServiceTable() {
	dml := masterTable()

	switch dml {
	case "1":
		tabel := "mst_service"
		headers := []string{"Id", "Service Name", "Price"}
		scanDest := []interface{}{new(string), new(string), new(int)}
		getAllData(tabel, headers, scanDest...)
	case "2":
		addServiceData()
	case "3":
		editServiceData()
	case "4":
		deleteService()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func addTransaction(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan id customer: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid id. Please enter a valid number.")
		return
	}

	fmt.Println("Masukkan Service Id:")
	scanner.Scan()
	service_id := scanner.Text()

	fmt.Println("Masukkan Quantity:")
	scanner.Scan()
	quantityStr := scanner.Text()
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		fmt.Println("Invalid quantity. Please enter a valid number.")
		return
	}

	fmt.Println("Masukkan Tanggal Transaksi (yyyy-mm-dd):")
	scanner.Scan()
	tgl_transaksi := scanner.Text()

	transaction := entity.Transaction{CustomerId: id,ServiceId: service_id,Quantity: quantity,TransactionDate: tgl_transaksi}
	makeTransaction(transaction)
}

func processTransactionTable() {
	dml := transactionTable()
	switch dml {
	case "1":
		tabel := "trx_transaction"
		headers := []string{"Id", "Customer Id", "Service Id", "Quantity", "Tanggal Transaksi", "Total Transaksi"}
		scanDest := []interface{}{new(int), new(int), new(string), new(int), new(string), new(int)}
		getAllData(tabel, headers, scanDest...)
	case "2":
		addTransaction()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func masterTable() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pilih DML yang ingin digunakan: ")
	fmt.Println("1. VIEW")
	fmt.Println("2. INSERT")
	fmt.Println("3. UPDATE")
	fmt.Println("4. DELETE")
	scanner.Scan()
	dml := scanner.Text()
	clearScreen()
	return dml
}

func transactionTable() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pilih DML yang ingin digunakan: ")
	fmt.Println("1. VIEW")
	fmt.Println("2. INSERT")
	scanner.Scan()
	dml := scanner.Text()
	clearScreen()
	return dml
}





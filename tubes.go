package main

import (
	"fmt"
	"os"
)

const maxAccounts = 100
const maxTransactions = 100

type Transaction struct {
	Type   string
	Amount float64
}

type Account struct {
	Name         string
	Number       string
	Balance      float64
	CardNumber   string
	PIN          string
	Transactions [maxTransactions]Transaction
	TransCount   int
}

var accounts [maxAccounts]Account
var accountCount int

func main() {
	for {
		showMenu()
		choice := getUserChoice()
		processChoice(choice)
	}
}

func showMenu() {
	fmt.Println("=== ATM Menu ===")
	fmt.Println("1. Pendataan Nasabah")
	fmt.Println("2. Cari Data Nasabah (Binary Search)")
	fmt.Println("3. Cari Data Nasabah (Sequential Search)")
	fmt.Println("4. Lihat Riwayat Transaksi")
	fmt.Println("5. Transaksi")
	fmt.Println("6. Edit Data Nasabah")
	fmt.Println("7. Hapus Data Nasabah")
	fmt.Println("8. Tampilkan Data Nasabah (Selection Sort)")
	fmt.Println("9. Tampilkan Data Nasabah (Insertion Sort)")
	fmt.Println("10. Exit")
	fmt.Print("Pilih opsi: ")
}

func getUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func processChoice(choice int) {
	switch choice {
	case 1:
		pendaftaranNasabah()
	case 2:
		cariDataNasabahBinary()
	case 3:
		cariDataNasabahSequential()
	case 4:
		lihatRiwayatTransaksi()
	case 5:
		transaksi()
	case 6:
		editDataNasabah()
	case 7:
		hapusDataNasabah()
	case 8:
		tampilkanDataNasabah("selection")
	case 9:
		tampilkanDataNasabah("insertion")
	case 10:
		exitATM()
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func pendaftaranNasabah() {
	if accountCount >= maxAccounts {
		fmt.Println("Jumlah maksimum akun telah tercapai.")
		return
	}

	var name, number, cardNumber, pin string
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan nomor rekening: ")
	fmt.Scan(&number)
	fmt.Print("Masukkan nomor kartu: ")
	fmt.Scan(&cardNumber)
	fmt.Print("Masukkan PIN: ")
	fmt.Scan(&pin)

	for i := 0; i < accountCount; i++ {
		if accounts[i].Number == number {
			fmt.Println("Nomor rekening sudah ada.")
			return
		}
		if accounts[i].CardNumber == cardNumber {
			fmt.Println("Nomor kartu sudah ada.")
			return
		}
	}

	accounts[accountCount] = Account{Name: name, Number: number, Balance: 0, CardNumber: cardNumber, PIN: pin, TransCount: 0}
	accountCount++
	fmt.Println("Akun berhasil dibuat.")
	sortAccounts() // Sort the accounts after adding a new one
}

func sortAccounts() {
	for i := 0; i < accountCount-1; i++ {
		minIdx := i
		for j := i + 1; j < accountCount; j++ {
			if accounts[j].Number < accounts[minIdx].Number {
				minIdx = j
			}
		}
		accounts[i], accounts[minIdx] = accounts[minIdx], accounts[i]
	}
}

func cariDataNasabahBinary() {
	var input string
	fmt.Print("Masukkan nomor rekening atau nomor kartu: ")
	fmt.Scan(&input)

	account := binarySearchAccount(input)
	if account != nil {
		fmt.Println("=== Informasi Akun ===")
		fmt.Printf("Nama: %s\n", account.Name)
		fmt.Printf("Nomor Rekening: %s\n", account.Number)
		fmt.Printf("Nomor Kartu: %s\n", account.CardNumber)
		fmt.Printf("Saldo: $%.2f\n", account.Balance)
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func cariDataNasabahSequential() {
	var input string
	fmt.Print("Masukkan nomor rekening atau nomor kartu: ")
	fmt.Scan(&input)

	account := sequentialSearchAccount(input)
	if account != nil {
		fmt.Println("=== Informasi Akun ===")
		fmt.Printf("Nama: %s\n", account.Name)
		fmt.Printf("Nomor Rekening: %s\n", account.Number)
		fmt.Printf("Nomor Kartu: %s\n", account.CardNumber)
		fmt.Printf("Saldo: $%.2f\n", account.Balance)
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func sequentialSearchAccount(key string) *Account {
	for i := 0; i < accountCount; i++ {
		if accounts[i].Number == key || accounts[i].CardNumber == key {
			return &accounts[i]
		}
	}
	return nil
}

func getAccount() *Account {
	var number, pin string
	for i := 0; i < 3; i++ {
		fmt.Print("Masukkan nomor rekening: ")
		fmt.Scan(&number)
		fmt.Print("Masukkan PIN: ")
		fmt.Scan(&pin)

		account := binarySearchAccount(number)
		if account != nil {
			if account.PIN != pin {
				fmt.Println("PIN salah.")
				continue
			}
			return account
		} else {
			fmt.Println("Akun tidak ditemukan.")
		}
	}
	fmt.Println("PIN salah tiga kali. Kembali ke menu utama.")
	return nil
}

func lihatRiwayatTransaksi() {
	account := getAccount()
	if account == nil {
		return
	}

	fmt.Println("=== Riwayat Transaksi ===")
	for i := 0; i < account.TransCount; i++ {
		fmt.Printf("%s: $%.2f\n", account.Transactions[i].Type, account.Transactions[i].Amount)
	}
}

func transaksi() {
	account := getAccount()
	if account == nil {
		return
	}

	for {
		fmt.Println("=== Menu Transaksi ===")
		fmt.Println("1. Tarik Tunai")
		fmt.Println("2. Setor Tunai")
		fmt.Println("3. Transfer")
		fmt.Println("4. Pembayaran")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			withdraw(account)
		case 2:
			deposit(account)
		case 3:
			transfer(account)
		case 4:
			payment(account)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func withdraw(account *Account) {
	var amount float64
	fmt.Print("Masukkan jumlah yang akan ditarik: ")
	fmt.Scan(&amount)

	if amount > account.Balance {
		fmt.Println("Saldo tidak mencukupi.")
	} else if amount <= 0 {
		fmt.Println("Jumlah harus lebih dari nol.")
	} else if account.TransCount >= maxTransactions {
		fmt.Println("Jumlah maksimum transaksi telah tercapai.")
	} else {
		account.Balance -= amount
		account.Transactions[account.TransCount] = Transaction{Type: "Withdraw", Amount: amount}
		account.TransCount++
		fmt.Printf("Anda telah menarik: $%.2f\n", amount)
		fmt.Printf("Saldo baru: $%.2f\n", account.Balance)
	}
}

func deposit(account *Account) {
	var amount float64
	fmt.Print("Masukkan jumlah yang akan disetor: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah harus lebih dari nol.")
	} else if account.TransCount >= maxTransactions {
		fmt.Println("Jumlah maksimum transaksi telah tercapai.")
	} else {
		account.Balance += amount
		account.Transactions[account.TransCount] = Transaction{Type: "Deposit", Amount: amount}
		account.TransCount++
		fmt.Printf("Anda telah menyetor: $%.2f\n", amount)
		fmt.Printf("Saldo baru: $%.2f\n", account.Balance)
	}
}

func transfer(account *Account) {
	var targetNumber string
	var amount float64
	fmt.Print("Masukkan nomor rekening tujuan: ")
	fmt.Scan(&targetNumber)
	fmt.Print("Masukkan jumlah yang akan ditransfer: ")
	fmt.Scan(&amount)

	targetAccount := binarySearchAccount(targetNumber)
	if targetAccount == nil {
		fmt.Println("Rekening tujuan tidak ditemukan.")
		return
	}

	if amount > account.Balance {
		fmt.Println("Saldo tidak mencukupi.")
	} else if amount <= 0 {
		fmt.Println("Jumlah harus lebih dari nol.")
	} else if account.TransCount >= maxTransactions || targetAccount.TransCount >= maxTransactions {
		fmt.Println("Jumlah maksimum transaksi telah tercapai.")
	} else {
		account.Balance -= amount
		targetAccount.Balance += amount
		account.Transactions[account.TransCount] = Transaction{Type: "Transfer Out", Amount: amount}
		account.TransCount++
		targetAccount.Transactions[targetAccount.TransCount] = Transaction{Type: "Transfer In", Amount: amount}
		targetAccount.TransCount++
		fmt.Printf("Anda telah mentransfer: $%.2f ke rekening %s\n", amount, targetNumber)
		fmt.Printf("Saldo baru: $%.2f\n", account.Balance)
	}
}

func payment(account *Account) {
	var amount float64
	fmt.Print("Masukkan jumlah pembayaran: ")
	fmt.Scan(&amount)

	if amount > account.Balance {
		fmt.Println("Saldo tidak mencukupi.")
	} else if amount <= 0 {
		fmt.Println("Jumlah harus lebih dari nol.")
	} else if account.TransCount >= maxTransactions {
		fmt.Println("Jumlah maksimum transaksi telah tercapai.")
	} else {
		account.Balance -= amount
		account.Transactions[account.TransCount] = Transaction{Type: "Payment", Amount: amount}
		account.TransCount++
		fmt.Printf("Anda telah membayar: $%.2f\n", amount)
		fmt.Printf("Saldo baru: $%.2f\n", account.Balance)
	}
}

func editDataNasabah() {
	account := getAccount()
	if account == nil {
		return
	}

	var choice int
	fmt.Println("=== Edit Data Nasabah ===")
	fmt.Println("1. Nama")
	fmt.Println("2. Nomor Kartu")
	fmt.Println("3. PIN")
	fmt.Print("Pilih data yang ingin diubah: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var newName string
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&newName)
		account.Name = newName
		fmt.Println("Nama berhasil diubah.")
	case 2:
		var newCardNumber string
		fmt.Print("Masukkan nomor kartu baru: ")
		fmt.Scan(&newCardNumber)
		account.CardNumber = newCardNumber
		fmt.Println("Nomor kartu berhasil diubah.")
	case 3:
		var newPIN string
		fmt.Print("Masukkan PIN baru: ")
		fmt.Scan(&newPIN)
		account.PIN = newPIN
		fmt.Println("PIN berhasil diubah.")
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func hapusDataNasabah() {
	var input string
	fmt.Print("Masukkan nomor rekening atau nomor kartu: ")
	fmt.Scan(&input)

	account := binarySearchAccount(input)
	if account == nil {
		fmt.Println("Akun tidak ditemukan.")
		return
	}

	for i := 0; i < accountCount; i++ {
		if &accounts[i] == account {
			// Menggeser semua akun setelah akun yang akan dihapus ke kiri
			for j := i; j < accountCount-1; j++ {
				accounts[j] = accounts[j+1]
			}
			accountCount--
			fmt.Println("Akun berhasil dihapus.")
			return
		}
	}
}

func tampilkanDataNasabah(sortType string) {
	if accountCount == 0 {
		fmt.Println("Tidak ada nasabah yang terdaftar.")
		return
	}

	if sortType == "selection" {
		selectionSortByName()
	} else if sortType == "insertion" {
		insertionSortByName()
	}

	fmt.Println("=== Data Nasabah ===")
	for i := 0; i < accountCount; i++ {
		account := accounts[i]
		fmt.Printf("Nama: %s, Nomor Rekening: %s, Nomor Kartu: %s, Saldo: $%.2f\n", account.Name, account.Number, account.CardNumber, account.Balance)
	}
}

func selectionSortByName() {
	for i := 0; i < accountCount-1; i++ {
		minIdx := i
		for j := i + 1; j < accountCount; j++ {
			if accounts[j].Name < accounts[minIdx].Name {
				minIdx = j
			}
		}
		accounts[i], accounts[minIdx] = accounts[minIdx], accounts[i]
	}
}

func insertionSortByName() {
	for i := 1; i < accountCount; i++ {
		key := accounts[i]
		j := i - 1
		for j >= 0 && accounts[j].Name > key.Name {
			accounts[j+1] = accounts[j]
			j = j - 1
		}
		accounts[j+1] = key
	}
}

func binarySearchAccount(key string) *Account {
	left, right := 0, accountCount-1
	for left <= right {
		mid := left + (right-left)/2
		if accounts[mid].Number == key || accounts[mid].CardNumber == key {
			return &accounts[mid]
		} else if accounts[mid].Number < key && accounts[mid].CardNumber < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nil
}

func exitATM() {
	fmt.Println("Terima kasih telah menggunakan ATM. Sampai jumpa!")
	os.Exit(0)
}
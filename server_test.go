package main

import (
	"database/sql"
	"os"
	"testing"
	// "fmt"
	// "log"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateAccountsDB(t *testing.T) {
	dbPath := "/usr/share/ats-htmx-echo/ats.db"
	defer os.Remove(dbPath)

	createAccountsDB(dbPath)
	
	var err error // Declare the err variable

	// Check if the database file exists
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		t.Errorf("Accounts database file does not exist")
	}

	// Open the database and check if the table exists
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Errorf("Failed to open accounts database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("SELECT * FROM accounts")
	if err != nil {
		t.Errorf("Failed to create accounts table: %v", err)
	}
}

func TestCreateCommentsDB(t *testing.T) {
	dbPath := "/usr/share/ats-htmx-echo/ats.db"
	defer os.Remove(dbPath)

	createCommentsDB(dbPath)
	var err error // Declare the err variable

	// Check if the database file exists
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		t.Errorf("Comments database file does not exist")
	}

	// Open the database and check if the table exists
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Errorf("Failed to open comments database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("SELECT * FROM comments")
	if err != nil {
		t.Errorf("Failed to create comments table: %v", err)
	}
}

//write a test for the createEstimatesDB function
func TestCreateEstimatesDB(t *testing.T) {
	dbPath := "/usr/share/ats-htmx-echo/ats.db"
	defer os.Remove(dbPath)

	createEstimatesDB(dbPath)
	var err error // Declare the err variable

	// Check if the database file exists
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		t.Errorf("Estimates database file does not exist")
	}

	// Open the database and check if the table exists
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Errorf("Failed to open estimates database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("SELECT * FROM estimates")
	if err != nil {
		t.Errorf("Failed to create estimates table: %v", err)
	}
}

//write a test for the atsUUID function
func TestAtsUUID(t *testing.T) {
	uuid := atsUUID()
	if len(uuid) != 36 {
		t.Errorf("Invalid UUID length: %s", uuid)
	}
}
//write a test for the todasDate function
func TestTodaysDate(t *testing.T) {
	date := todaysDate()
	if len(date) != 10 {
		t.Errorf("Invalid date length: %s", date)
	}
}

//write a test for the nameCheck function
func TestNameCheck(t *testing.T) {
	name := "John Doe"
	if !nameCheck(name) {
		t.Errorf("Invalid name: %s", name)
	}
	name = "John_Doe"
	if nameCheck(name) {
		t.Errorf("Invalid name: %s", name)
	}
}

//write a test for the checkEmailParts function
func TestCheckEmailParts(t *testing.T) {
	email := "boo@gmail.com"
	if !checkEmailParts(email) {
		t.Errorf("Invalid email: %s", email)
	}
}

//write a test for the emailCheck function
func TestEmailCheck(t *testing.T) {
	email := "boo@gmail.com"
	if !emailCheck(email) {
		t.Errorf("Invalid email: %s", email)
	}
	email = "boo@gmail"
	if emailCheck(email) {
		t.Errorf("Invalid email: %s", email)
	}
	email = "boogmail.com"
	if emailCheck(email) {
		t.Errorf("Invalid email: %s", email)
	}
}
//write a test for the ratingCheck function
func TestRatingCheck(t *testing.T) {
	rating := "5"
	if !ratingCheck(rating) {
		t.Errorf("Invalid rating: %s", rating)
	}
	rating = "0"
	if ratingCheck(rating) {
		t.Errorf("Invalid rating: %s", rating)
	}
	rating = "6"
	if ratingCheck(rating) {
		t.Errorf("Invalid rating: %s", rating)
	}
}
//write a test for the commentCheck function
func TestCommentCheck(t *testing.T) {
	comment := "This is a comment"
	if commentCheck(comment) {
		t.Errorf("Invalid comment: %s", comment)
	}
	comment = "This is a fucked up comment"
	if !commentCheck(comment) {
		t.Errorf("Invalid Comment: %s", comment)
	}
}

//write a test for the addressCheck function
func TestAddressCheck(t *testing.T) {
	address := "123 Main St"
	if addressCheck(address) {
		t.Errorf("Invalid address: %s", address)
	}
	address = "Apt 456"
	if addressCheck(address) {
		t.Errorf("Invalid address: %s", address)
	}
	address = "Suite 789"
	if addressCheck(address) {
		t.Errorf("Invalid address: %s", address)
	}
	address = "P.O. Box 123"
	if addressCheck(address) {
		t.Errorf("Invalid address: %s", address)
	}
	address = ""
	if !addressCheck(address) {
		t.Errorf("Invalid address: %s", address)
	}
}
//write a test for the phoneCheck function
func TestPhoneCheck(t *testing.T) {
	phone := "123 456 7899"
	if !phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
	phone = "1234567890"
	if phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
	phone = "123-456-7890"
	if phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
	phone = "(123) 456-7890"
	if phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
	phone = "123-456"
	if phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
	phone = "12345678901234567890"
	if phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
	phone = ""
	if phoneCheck(phone) {
		t.Errorf("Invalid phone number: %s", phone)
	}
}

func TestServDateCheck(t *testing.T) {
	servdate := "12 31 2022"
	if !servDateCheck(servdate) {
		t.Errorf("Invalid service date: %s", servdate)
	}

	servdate = "20221231"
	if servDateCheck(servdate) {
		t.Errorf("Invalid service date: %s", servdate)
	}

	servdate = "2022-02-30"
	if servDateCheck(servdate) {
		t.Errorf("Invalid service date: %s", servdate)
	}

	servdate = ""
	if servDateCheck(servdate) {
		t.Errorf("Invalid service date: %s", servdate)
	}

	servdate = "2022-12-32"
	if servDateCheck(servdate) {
		t.Errorf("Invalid service date: %s", servdate)
	}
}

func TestAccountCheck(t *testing.T) {
	email := "test@example.com"
	hasAccount := accountCheck(email)
	if hasAccount {
		t.Errorf("Expected accountCheck to return false for email: %s", email)
	}

	email = "existing@example.com"
	hasAccount = accountCheck(email)
	if !hasAccount {
		t.Errorf("Expected accountCheck to return true for email: %s", email)
	}
}









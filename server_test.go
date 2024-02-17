package main

import (
	"testing"
	"os"
	"mime/multipart"
	"fmt"

)

func TestAtsUUID(t *testing.T) {
	uuid := atsUUID()
	if len(uuid) != 36 {
		t.Errorf("atsUUID() returned an invalid UUID")
	}
}

func TestNameCheck(t *testing.T) {
	// Test case 1: Valid name
	if !nameCheck("John Doe") {
		t.Errorf("nameCheck(\"John Doe\") returned false, expected true")
	}

	// Test case 2: Empty name
	if nameCheck("") {
		t.Errorf("nameCheck(\"\") returned true, expected false")
	}

	// Test case 3: Name with special characters
	if nameCheck("John@Doe") {
		t.Errorf("nameCheck(\"John@Doe\") returned true, expected false")
	}
}

func TestEmailCheck(t *testing.T) {
	// Test case 1: Valid email
	if !emailCheck("john.doe@example.com") {
		t.Errorf("emailCheck(\"john.doe@example.com\") returned false, expected true")
	}

	// Test case 2: Invalid email
	if emailCheck("john.doe@example") {
		t.Errorf("emailCheck(\"john.doe@example\") returned true, expected false")
	}

	// Test case 3: Empty email
	if emailCheck("") {
		t.Errorf("emailCheck(\"\") returned true, expected false")
	}
}

func TestRatingCheck(t *testing.T) {
	// Test case 1: Valid rating
	if !ratingCheck("5") {
		t.Errorf("ratingCheck(\"5\") returned false, expected true")
	}

	// Test case 2: Invalid rating
	if ratingCheck("10") {
		t.Errorf("ratingCheck(\"10\") returned true, expected false")
	}

	// Test case 3: Empty rating
	if ratingCheck("") {
		t.Errorf("ratingCheck(\"\") returned true, expected false")
	}
}

func TestCommentCheck(t *testing.T) {
	// Test case 1: Valid comment
	if commentCheck("This is a valid comment.") {
		t.Errorf("commentCheck(\"This is a valid comment.\") returned false, expected true")
	}

	// Test case 2: Empty comment
	if commentCheck("") {
		t.Errorf("commentCheck(\"\") returned true, expected false")
	}

	// Test case 3: Comment with special characters
	if commentCheck("This is a comment with @#$%%^&* special characters.") {
		t.Errorf("commentCheck(\"This is a comment with @#$%%^&* special characters.\") returned true, expected false")
	}
}

func TestAddressCheck(t *testing.T) {
	// Test case 1: Valid address
	if !addressCheck("123 Main St") {
		t.Errorf("addressCheck(\"123 Main St\") returned false, expected true")
	}

	// Test case 2: Empty address
	if addressCheck("") {
		t.Errorf("addressCheck(\"\") returned true, expected false")
	}

	// Test case 3: Address with special characters
	if addressCheck("123 @#$%^&* St") {
		t.Errorf("addressCheck(\"123 @#$%%^&* St\") returned true, expected false")
	}
}

func TestPhoneCheck(t *testing.T) {
	// Test case 1: Valid phone number
	if !phoneCheck("123 456 7890") {
		t.Errorf("phoneCheck(\"123 456 7890\") returned false, expected true")
	}

	// Test case 2: Empty phone number
	if phoneCheck("") {
		t.Errorf("phoneCheck(\"\") returned true, expected false")
	}

	// Test case 3: Phone number with special characters
	if phoneCheck("123-456-7890") {
		t.Errorf("phoneCheck(\"123-456-7890\") returned true, expected false")
	}
}

func TestServDateCheck(t *testing.T) {
	// Test case 1: Valid service date
	if !servDateCheck("12 31 2022") {
		t.Errorf("servDateCheck(\"12 31 2022\") returned false, expected true")
	}

	// Test case 2: Empty service date
	if servDateCheck("") {
		t.Errorf("servDateCheck(\"\") returned true, expected false")
	}

	// Test case 3: Invalid service date format
	if servDateCheck("12-31-2022") {
		t.Errorf("servDateCheck(\"31-12-2022\") returned true, expected false")
	}
}

func TestCheckEstInputs(t *testing.T) {
	// Test case 1: Valid inputs
	name := "John Doe"
	address := "123 Main St"
	city := "New York"
	phone := "1234567890"
	email := "johndoe@example.com"
	servdate := "12 31 2022"
	comment := "This is a test comment"
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned false, expected true")
	}

	// Test case 2: Empty name
	name = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}

	// Test case 3: Empty address
	name = "John Doe"
	address = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}

	// Test case 4: Empty city
	address = "123 Main St"
	city = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}

	// Test case 5: Empty phone
	city = "New York"
	phone = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}

	// Test case 6: Empty email
	phone = "1234567890"
	email = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}

	// Test case 7: Empty servdate
	email = "johndoe@example.com"
	servdate = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}

	// Test case 8: Empty comment
	servdate = "12 31 2022"
	comment = ""
	if checkEstInputs(name, address, city, phone, email, servdate, comment) {
		t.Errorf("checkEstInputs returned true, expected false")
	}
}

func TestCheckComInputs(t *testing.T) {
	// Test case 1: Valid inputs
	if checkComInputs("John Doe", "john@example.com", "5", "Great service!") {
		t.Error("Expected checkComInputs to return true for valid inputs")
	}

	// Test case 2: Empty name
	if checkComInputs("", "john@example.com", "5", "Great service!") {
		t.Error("Expected checkComInputs to return false for empty name")
	}

	// Test case 3: Invalid email
	if checkComInputs("John Doe", "invalid_email", "5", "Great service!") {
		t.Error("Expected checkComInputs to return false for invalid email")
	}

	// Test case 4: Invalid rating
	if checkComInputs("John Doe", "john@example.com", "10", "Great service!") {
		t.Error("Expected checkComInputs to return false for invalid rating")
	}

	// Test case 5: Empty comment
	if checkComInputs("John Doe", "john@example.com", "5", "") {
		t.Error("Expected checkComInputs to return false for empty comment")
	}
}

func TestSaveFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile := "/usr/share/ats-htmx-echo/testfile.webp"
	
	defer os.Remove(tempFile)

	// Create a mock multipart.FileHeader
	fileHeader := &multipart.FileHeader{
		Filename: tempFile,
		Size:     219016,
	}

	// Call the save_file function
	comid := "12345"
	filePath, err := save_file(comid, fileHeader)
	if err != nil {
		t.Errorf("save_file returned an error: %v", err)
	}

	// Verify that the file was saved to the correct path
	expectedPath := fmt.Sprintf("/usr/share/ats-htmx-echo/UpLoads/%s", comid)
	if filePath != expectedPath {
		t.Errorf("save_file returned an incorrect file path. Expected: %s, Got: %s", expectedPath, filePath)
	}

	// Verify that the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("save_file did not save the file to the expected path: %s", filePath)
	}
}




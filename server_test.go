package main

import "testing"

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
	if !commentCheck("This is a valid comment.") {
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








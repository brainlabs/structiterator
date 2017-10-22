package validator

import (
	"testing"
	"time"
)

func TestValidation_Email(t *testing.T) {
	validation := Validation{
		PhoneFormat: `^(62|0)([0-9]*)$`,
		EmailFormat: `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`,
	}

	t.Log("\nTesting Validation Email:")
	{
		err := validation.Email("samalekom", "email_address","")
		t.Log("The error", err)

		err = validation.Email("kubilisme@example.com", "email_address2", "")
		t.Log("The error2", err)

		err = validation.Email(12345, "email_address", "")
		t.Log("The error3", err)
	}

	t.Log("\nTesting Validation Phone")
	{
		err := validation.Phone("79178877", "phone_number", "")
		t.Log("The error", err)

		err = validation.Phone("021791788777", "phone_number", "")
		t.Log("The error2", err)

		err = validation.Phone("62817180018", "cell_phone", "")
		t.Log("The error3", err)
	}
}

func TestValidation_AfterDate(t *testing.T) {
	validtn := Validation{
		PhoneFormat: `^(62|0)([0-9]*)$`,
		EmailFormat: `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`,
		DateFormat: "02/01/2006",
	}

	t.Log("\nTesting validation after date:")
	{
		app := Application{}
		app.Id = 1
		app.AppliedTime = "20/09/2017"
		app.ApprovedTime = "19/09/2017"
		err := validtn.AfterDate(app, "approved_time", "applied_time", "")
		t.Log("Error", err)
	}
}

func TestDate(t *testing.T) {
	t.Log("\nTesting parsing date:")
	{
		tim, err := time.Parse("01/02/2006", "09/20/2017")

		t.Log("The time", tim, "The error", err)

		value  := "05/19/11"
		// Writing down the way the standard time would look like formatted our way
		layout := "01/02/06"
		tim2, _ := time.Parse(layout, value)
		t.Log(tim2)
	}
}
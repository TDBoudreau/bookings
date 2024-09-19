package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")
	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/testing", nil)
	form := New(r.PostForm)
	has := form.Has("c")
	if has {
		t.Error("shows form has field, when it should not")
	}

	data := url.Values{}
	data.Add("a", "a")
	form = New(data)
	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field, when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/testing", nil)
	form := New(r.PostForm)
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("shows min length for unknown field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	data := url.Values{}
	data.Add("some_field", "some value")
	form = New(data)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 met when data is shorter")
	}

	data = url.Values{}
	data.Add("another_field", "abc123")
	form = New(data)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should have an error, but did not get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/testing", nil)
	form := New(r.PostForm)
	form.IsEmail("unknown")
	if form.Valid() {
		t.Error("expected form error, but got none")
	}

	data := url.Values{}
	data.Add("bad_email", "test@testing")
	form = New(data)
	form.IsEmail("bad_email")
	if form.Valid() {
		t.Error("expected form error, but got none")
	}

	data = url.Values{}
	data.Add("good_email", "me@here.com")
	form = New(data)
	form.IsEmail("good_email")
	if !form.Valid() {
		t.Error("form is invalid, when it should be valid")
	}
}

package auth

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

// //ONLY RUN TO CREATE NEW USER
// func TestNewUser(t *testing.T) {
// 	u := User{
// 		Username: "sample",
// 		Password: "pw1234",
// 	}
// 	err := NewUser(u)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
// func TestEndpoint3(t *testing.T) {
// 	e := echo.New()
// 	validUser := `{"username":"sample","password":"pw1234"}`
// 	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(validUser))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	if err := CreateToken(c); err != nil {
// 		t.Error(err)
// 	}
// 	e.Use(Validate)
// 	tokenString := rec.Body.String()
// 	req = httptest.NewRequest(echo.POST, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("authentication", tokenString)
// 	rec = httptest.NewRecorder()
// 	c = e.NewContext(req, rec)
// 	err := SampleEndpoint(c)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

func Test3(t *testing.T) {
	e := echo.New()
	validUser := `{"username":"sample","password":"pw1234"}`
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(validUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := CreateToken(c); err != nil {
		t.Error(err)
	}
	// e.Use(Validate)
	tokenString := rec.Body.String()
	req = httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", tokenString)
	// fmt.Println("ehader is ", req.Header.Get("Authorization"))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h := Validate(nil)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	err := h(c)
	if err != nil {
		t.Error(err)
	}
}

func TestEndpoint2(t *testing.T) {
	// STEP 1: GET TOKEN
	e := echo.New()
	validUser := `{"username":"sample","password":"pw1234"}`
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(validUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := CreateToken(c); err != nil {
		t.Error(err)
	}
	// STEP 2: TEST TOKEN IS VALID BY GETTING USERNAME+PASSWORD FROM TOKEN
	tokenString := rec.Body.String()
	req = httptest.NewRequest(echo.GET, "/", nil)
	q := req.URL.Query()
	q.Add("token", tokenString)
	req.URL.RawQuery = q.Encode()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err := ProtectedEndpoint(c)
	if err != nil {
		t.Error(err)
	}

}

func TestEndpoint(t *testing.T) {
	// SHOULD FAIL BECAUSE WHATEVER WHATEVER
	e := echo.New()
	invalidUser := `{"username":"Jon Snow","password":"jon@labstack.com"}`
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(invalidUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreateToken(c)
	if err == nil {
		t.Error(errors.New("should have erred with invalid user"))
	}
	fmt.Println(rec.Body.String())

	validUser := `{"username":"sample","password":"pw1234"}`
	req = httptest.NewRequest(echo.POST, "/", strings.NewReader(validUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = CreateToken(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rec.Body.String())

}

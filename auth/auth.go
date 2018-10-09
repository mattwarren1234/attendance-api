package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type JwtToken struct {
	Token string `json:"token"`
}

func NewUser(u User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into auth(username, password) values ($1, $2)", u.Username, hash)
	return err

}

func getPWHash(u User) ([]byte, error) {
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		return nil, err
	}
	var hash []byte
	err = db.QueryRow("select password from auth where username=$1", u.Username).Scan(&hash)
	if err != nil {
		fmt.Println("getpwhash: err is", err, "for", u.Username)
	}
	return hash, err

}

func checkPassword(u User) error {
	pwHash, err := getPWHash(u)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(u.Password)); err != nil {
		return err
	}
	return nil
}

func CreateToken(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return err
	}
	if err := checkPassword(u); err != nil {
		return err
	}
	// TODO: check db for validated stuff n things!
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
	})
	tokenString, error := token.SignedString([]byte(secret))
	if error != nil {
		fmt.Println(error)
	}
	return c.JSON(http.StatusOK, JwtToken{Token: tokenString})
}

var secret = "someSecret"

func ProtectedEndpoint(c echo.Context) error {
	unparsedToken := c.QueryParam("token")
	var j JwtToken
	err := json.Unmarshal([]byte(unparsedToken), &j)
	if err != nil {
		fmt.Println("PARS err is", err)
	}
	token, err := jwt.Parse(j.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		return c.JSON(http.StatusOK, user)
	} else {
		return errors.New("Invalid auth token")
	}
}

func SampleEndpoint(c echo.Context) error {
	decoded := c.Get("decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	return c.JSON(http.StatusOK, user)
}

func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		authorizationHeader := c.Request().Header.Get("authentication")
		if authorizationHeader == "" {
			// to do : return 401
			return errors.New("Auth Header empty")
		}
		var j JwtToken
		err := json.Unmarshal([]byte(authorizationHeader), &j)
		if err != nil {
			fmt.Println("parse err is", err)
			return err
		}
		fmt.Println("here HERE HERE")
		token, err := jwt.Parse(j.Token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(secret), nil
		})
		fmt.Println("token.cal", token.Claims, "err is", err)
		if err != nil {
			return err
		}
		if token.Valid {
			c.Set("decoded", token.Claims)
		} else {
			return errors.New("Invalid authorization token")
		}
		return nil

	}
}

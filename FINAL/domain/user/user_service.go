package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var sampleSecretKey = []byte("GoLinuxCloudKey")

type repository interface {
	writeUserRepo
}

type writeUserRepo interface {
	createUser(ctx context.Context, param User) (err error, user User)
	getUser(ctx context.Context, username string, password string) (err error, user User)
}

type userService struct {
	repo repository
	r    userRepo
}

func newUserService(repo repository, r userRepo) userService {
	return userService{
		repo: repo,
		r:    r,
	}
}

func (u userService) createUser(ctx context.Context, req UserReq) (err error, resp UserResp) {
	var user User

	req.Password, err = HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return
	}
	err, user = u.repo.createUser(ctx, req.UserReqIntoUser())
	if err != nil {
		log.Println(err)
		return
	}

	if user.CreatedAt == "" {
		return
	} else {
		resp = user.UserIntoUserResp()
	}

	return err, resp
}

func (u userService) userLogin(ctx context.Context, req UserReq) (err error, resp LoginResp) {
	if err != nil {
		log.Println(err)
		return
	}

	err, user := u.repo.getUser(ctx, req.Username, req.Password)
	if user.UserId == 0 {
		return err, resp
	}

	if CheckPasswordHash(req.Password, user.Password) {
		err, resp.AccessToken = generateJWT(user.UserId)
		if err != nil {
			return err, resp
		}
	}

	return err, resp
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(userId int) (err error, accessToken string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	accessToken, err = token.SignedString(sampleSecretKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return err, accessToken
	}

	return err, accessToken
}

func VerifyJWT(accessToken string) (err error, userId float64) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("couldn't parse claims")
			return nil, errors.New("There was an error in parsing")
		}
		return sampleSecretKey, nil
	})
	if err != nil {
		log.Println(err)
		return err, userId
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("couldn't parse claims")
		return errors.New("Token error"), userId
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		log.Println("couldn't parse claims")
		return errors.New("Token error"), userId
	}

	return err, claims["userId"].(float64)
}

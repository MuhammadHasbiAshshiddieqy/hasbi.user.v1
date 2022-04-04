package auth

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// )

// type Token struct{}

// func NewToken() *Token {
// 	return &Token{}
// }

// type TokenInterface interface {
// 	CreateToken(userid string) (*TokenDetails, error)
// 	ExtractTokenMetadata(*fiber.Ctx) (*AccessDetails, error)
// }

// //Token implements the TokenInterface
// var _ TokenInterface = &Token{}

// func (t *Token) CreateToken(userid string) (*TokenDetails, error) {
// 	td := &TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
// 	td.TokenUuid = uuid.Must(uuid.NewRandom()).String()

// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = td.TokenUuid + "++" + userid

// 	var err error
// 	//Creating Access Token
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["access_uuid"] = td.TokenUuid
// 	atClaims["user_id"] = userid
// 	atClaims["exp"] = td.AtExpires
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	//Creating Refresh Token
// 	rtClaims := jwt.MapClaims{}
// 	rtClaims["refresh_uuid"] = td.RefreshUuid
// 	rtClaims["user_id"] = userid
// 	rtClaims["exp"] = td.RtExpires
// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return td, nil
// }

// func TokenValid(c *fiber.Ctx) error {
// 	token, err := VerifyToken(c)
// 	if err != nil {
// 		return err
// 	}
// 	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
// 		return err
// 	}
// 	return nil
// }

// func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
// 	tokenString := ExtractToken(c)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("ACCESS_SECRET")), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return token, nil
// }

// //get the token from the request body
// func ExtractToken(c *fiber.Ctx) string {
// 	bearToken := c.Get("Authorization")
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}
// 	return ""
// }

// func (t *Token) ExtractTokenMetadata(c *fiber.Ctx) (*AccessDetails, error) {
// 	fmt.Println("WE ENTERED METADATA")
// 	token, err := VerifyToken(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		accessUuid, ok := claims["access_uuid"].(string)
// 		if !ok {
// 			return nil, err
// 		}
// 		userId, ok := claims["user_id"].(string)
// 		if !ok {
// 			return nil, err
// 		}
// 		return &AccessDetails{
// 			TokenUuid: accessUuid,
// 			UserId:    userId,
// 		}, nil
// 	}
// 	return nil, err
// }
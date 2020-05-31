package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	//"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"physiobot/config"
	"physiobot/modules/common/errors"
	//"golang.org/x/crypto/bcrypt"
)

// AuthResponse struct with auth response
type AuthResponse struct {
	authorized bool
	partnerID  int
	clientID   string
	userID     string
}

// configuration global
var configuration = config.GetConfig()

// CheckCredentials basic auth method to check credentials, first check in redis if not found make a request to a secure endpooint to get credentials
func CheckCredentials(clientID string, clientSecret string, c echo.Context) *AuthResponse {
	/*
	cache := db.Cache()

	AuthResponse := new(AuthResponse)
	AuthResponse.authorized = false
	partnerID := ""

	hash, redisErr := cache.Get(clientID).Result()

	if redisErr == redis.Nil || redisErr != nil { //Secure Auth

		configuration.Jwt.Alg = os.ExpandEnv(configuration.Jwt.Alg)
		configuration.Jwt.Key = os.ExpandEnv(configuration.Jwt.Key)
		configuration.FcbDomains.Secure = os.ExpandEnv(configuration.FcbDomains.Secure)

		secretKey := configuration.Jwt.Key

		data := jwt.MapClaims{
			"jti":      uuid.New(),
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Unix() + configuration.Jwt.Exp,
			"app_code": "TS",
			"user_id":  1,
			"username": clientID,
		}

		decodedSecretKey, _ := base64.StdEncoding.DecodeString(secretKey)
		dataJwt := jwt.MapClaims{
			"data": data,
			"key":  decodedSecretKey,
			"alg":  configuration.Jwt.Alg,
			"kid":  "k1",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS512, dataJwt)
		tokenString, _ := token.SignedString(decodedSecretKey)
		url := configuration.FcbDomains.Secure + "/oauth/credentials"
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Set("Authorization", tokenString)

		response, err := http.DefaultClient.Do(request)

		if err != nil {
			return AuthResponse
		}

		defer response.Body.Close()
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return AuthResponse
		}

		responseArray := make(map[string]interface{})

		err = json.Unmarshal([]byte(responseData), &responseArray)

		if err != nil {
			return AuthResponse
		}

		partnerIDPrec, okPartner := responseArray["partner_id"]
		hashPrec, okHash := responseArray["client_secret"]

		if !okPartner || !okHash {
			return AuthResponse
		}

		partnerID = partnerIDPrec.(string)
		hash = hashPrec.(string)

		if redisErr == redis.Nil {
			clientParams := partnerID + " " + hash
			err = cache.Set(clientID, clientParams, time.Duration(configuration.Redis.ExpirationTime)*time.Second).Err()
		}
	} else {
		clientParams := strings.Fields(hash)
		partnerID = clientParams[0]
		hash = clientParams[1]
	}

	hashString := hash
	err := bcrypt.CompareHashAndPassword([]byte(hashString), []byte(clientSecret))

	if err != nil {
		return AuthResponse
	}

	AuthResponse.partnerID, _ = strconv.Atoi(partnerID)
	AuthResponse.authorized = true

	*/
	//return AuthResponse
	return nil;
}

//IsAuthBasic function to check if a user is atuthenticated
func IsAuthBasic(clientID string, clientSecret string, c echo.Context) (bool, error) {
	fmt.Println("AUTH BASIC")
	/*
	authorized := CheckCredentials(clientID, clientSecret, c)
	if !authorized.authorized {
		return authorized.authorized, echo.NewHTTPError(errors.INVALID_API_CREDENTIALS)
	}
	c.Set("partnerID", authorized.partnerID)
	c.Set("clientID", clientID)
	c.Set("clientSecret", clientSecret)
	return authorized.authorized, nil
	*/
	return false, nil
}

//IsAuthBearer function to check if a user is atuthenticated
func IsAuthBearer(key string, c echo.Context) (bool, error) {
	fmt.Println("AUTH BEARER")
	authorized := IsJWT(key, c)
	if !authorized.authorized {
		authorized = IsAuth0(key, c)
	}
	if !authorized.authorized {
		return authorized.authorized, echo.NewHTTPError(errors.INVALID_API_CREDENTIALS)
	}
	c.Set("partnerID", authorized.partnerID)
	c.Set("clientID", authorized.clientID)
	c.Set("clientSecret", key)
	return authorized.authorized, nil
}

// IsJWT check if the tokan is valid
func IsJWT(tokenString string, c echo.Context) *AuthResponse {
	AuthResponse := new(AuthResponse)
	AuthResponse.authorized = false

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(configuration.Jwt.Key), nil
	})

	if err != nil {
		return AuthResponse
	} else if token.Valid && token.Header["alg"] == configuration.Jwt.Alg {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["app_code"] == "MP" || claims["app_code"] == "TS" {
				AuthResponse.authorized = true
				AuthResponse.partnerID, _ = strconv.Atoi(claims["partner_id"].(string))
				AuthResponse.clientID = claims["username"].(string)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	return AuthResponse
}

// IsAuth0 verify if is a valida auth0 token
func IsAuth0(token string, c echo.Context) *AuthResponse {
	AuthResponse := new(AuthResponse)
	AuthResponse.authorized = false
/*
	hasher := md5.New()
	hasher.Write([]byte(token))
	key := "oauth-token-" + hex.EncodeToString(hasher.Sum(nil))

	cache := db.Cache()

	tokenInfo, redisErr := cache.Get(key).Result()

	if redisErr == redis.Nil || redisErr != nil { //Secure Auth
		url := configuration.FcbDomains.Secure + "/oauth/resource"
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return AuthResponse
		}

		defer response.Body.Close()
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return AuthResponse
		}

		ParseAuthResponse(AuthResponse, responseData)

		err = cache.Set(key, responseData, time.Duration(configuration.Redis.ExpirationTime)*time.Second).Err()

	} else {
		ParseAuthResponse(AuthResponse, []byte(tokenInfo))
	}
 */
	return AuthResponse
}

// ParseAuthResponse parse []byte to authresponse
func ParseAuthResponse(auth *AuthResponse, data []byte) *AuthResponse {
	responseArray := make(map[string]interface{})

	err := json.Unmarshal([]byte(data), &responseArray)

	if err != nil {
		return auth
	}

	if responseArray["success"] == true {
		auth.authorized = true
		auth.partnerID, _ = strconv.Atoi(responseArray["partner_id"].(string))
		auth.clientID = responseArray["client_id"].(string)
		if responseArray["user_id"] != nil {
			auth.userID = responseArray["user_id"].(string)
		}
	}
	return auth
}

// SkipBearer - Bearer Authentication will not be processed if this method returns true
func SkipBearer(c echo.Context) bool {
	if !strings.HasPrefix(c.Request().Header.Get("Authorization"), "Bearer") {
		return true
	}
	return false
}

// SkipBasic - Basic Authentication will not be processed if this method returns true
func SkipBasic(c echo.Context) bool {
	if !strings.HasPrefix(c.Request().Header.Get("Authorization"), "Basic") {
		return true
	}
	return false
}

// AuthHeader verify authorization header
func AuthHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := c.Request().RequestURI
		if uri == "/v3/system/status" { // strings.HasPrefix(uri, "/v3/customs/edt"
			return next(c)
		}
		if SkipBearer(c) && SkipBasic(c) {
			return echo.NewHTTPError(errors.INVALID_API_CREDENTIALS)
		}
		return next(c)
	}
}

// BuildPartnerJWT build jwt token
func BuildPartnerJWT(partnerID int) string {
	data := jwt.MapClaims{
		"app_code":   "TS",
		"user_id":    1,
		"username":   "admin",
		"partner_id": partnerID,
	}

	decodedSecretKey, _ := base64.StdEncoding.DecodeString(configuration.Jwt.Key)

	dataJwt := jwt.MapClaims{
		"jti":  "",//uuid.New(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Unix() + configuration.Jwt.Exp,
		"data": data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, dataJwt)
	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": configuration.Jwt.Alg,
		"kid": "k1",
	}
	tokenString, _ := token.SignedString(decodedSecretKey)
	return tokenString
}

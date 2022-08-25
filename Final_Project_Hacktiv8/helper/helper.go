package helper

import (
	"encoding/json"
	"errors"
	"golang-crud-sql/model"

	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/crypto/bcrypt"
)

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetConfig() (*model.Config, error) {
	var cfg model.Config

	err := cleanenv.ReadConfig("config/.env", &cfg)

	if err != nil {
		return nil, errors.New("failed to read configuration file")
	}

	return &cfg, nil
}

func CreateErrorResponse(message string) (byteMessage []byte) {
	var errMessage model.ErrorHandler

	errMessage.Error = message
	json, _ := json.Marshal(errMessage)
	return json
}

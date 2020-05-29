package util

import (
	"encoding/json"
	"experiment/pkg/constvar"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
	"io"
	"io/ioutil"
	"os"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func getDataScoreFileName() string {
	name, _ := GenShortId()
	return fmt.Sprintf("%s/%s.sql", viper.GetString("data_scour"), name)
}

func StoreFile(score io.Reader) (string, []byte, error) {
	fileName := getDataScoreFileName()
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal("open file fail", err)
		return "", nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(score)
	if err != nil {
		log.Fatal("read data fail", err)
		return "", nil, err
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal("write data fail", err)
		return "", nil, err
	}
	return fileName, data, nil
}

func GetUserId(c *gin.Context) int {
	userId, exists := c.Get("userId")
	if !exists {
		return constvar.EMPTY
	}
	return int(userId.(float64))
}

func MsgEncode(data interface{}) ([]byte, error) {
	realMsg, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return realMsg, nil
}

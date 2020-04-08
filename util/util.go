package util

import (
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

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func getDataScoreFileName() string {
	name, _ := GenShortId()
	return fmt.Sprintf("%s/%s.sql", viper.GetString("data_scour"), name)
}

func StoreFile(score io.Reader) (string, error) {
	fileName := getDataScoreFileName()
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open file fail", err)
		return "", err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(score)
	if err != nil {
		log.Fatal("read data fail", err)
		return "", err
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal("write data fail", err)
		return "", err
	}
	return fileName, nil
}

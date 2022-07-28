package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	MAX_FILE_SIZE = 30 * 1024 * 1024
	UPLOAD_INPUT  = "file"
)

type UploadFile struct {
	Url string
}

type FormData struct {
	StoreId  string `form:"storeId"`
	FileName string `form:"fileName"`
}

type UploadFileReq struct{}

func UploadProcess(c *gin.Context) ([]*UploadFile, error) {
	// formInfo := &FormData{}
	// c.Bind(formInfo)

	err := c.Request.ParseMultipartForm(MAX_FILE_SIZE)
	if err != nil {
		return nil, err
	}

	uploadFiles := []*UploadFile{}

	formData := c.Request.MultipartForm
	files := formData.File[UPLOAD_INPUT]

	for _, f := range files {
		file, err := f.Open()
		defer file.Close()
		if err != nil {
			return nil, err
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, file)
		if err != nil {
			return nil, err
		}

		originalFileName := f.Filename
		ext := path.Ext(originalFileName)
		localPath := "/img"
		filename := fmt.Sprintf("%v%v", time.Now().UnixNano(), ext)
		// filepath := fmt.Sprintf("/%v", formInfo.StoreId)
		filepathName := fmt.Sprintf("/%v", filename)
		err = os.MkdirAll("."+localPath, os.ModePerm)
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile("."+localPath+filepathName, buf.Bytes(), 0644)
		if err != nil {
			return nil, err
		}

		upfile := &UploadFile{
			Url: localPath + filepathName,
		}

		uploadFiles = append(uploadFiles, upfile)
	}
	return uploadFiles, nil
}

package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	// "bitbucket.org/bora777/bora-server/base/configure"
	// "bitbucket.org/bora777/bora-server/base/random2"
	// "bitbucket.org/bora777/bora-server/ec"
	// "bitbucket.org/bora777/bora-server/params"
	"github.com/gin-gonic/gin"
)

const (
	MAX_FILE_SIZE = 30 * 1024 * 1024
	UPLOAD_INPUT  = "file"
)

type UploadFile struct {
	// OriginalFileName string
	// NewFileName      string
	Url string
}

type FormData struct {
	StoreId  string `form:"storeId"`
	FileName string `form:"fileName"`
}

type UploadFileReq struct{}

// func validateUploadReq(req interface{}) *ec.Err {
// 	return nil
// }

func UploadProcess(c *gin.Context) ([]*UploadFile, error) {
	// Pwd, _ := os.Getwd()

	// Log.Infof("entering ... c = %p", c)

	// s := c.MustGet(proto.CONTEXT_KEY_SPROTO_CONTEXT).(*sproto.SProtoContext)
	// sess := s.GetSession()
	formInfo := &FormData{}
	c.Bind(formInfo)

	err := c.Request.ParseMultipartForm(MAX_FILE_SIZE)
	if err != nil {
		// Log.Error(err)
		// s.ClientError(ec.ERR_UPLOAD_FAILED)
		return nil, err
	}

	uploadFiles := []*UploadFile{}

	formdata := c.Request.MultipartForm
	files := formdata.File[UPLOAD_INPUT]

	for _, f := range files {
		file, err := f.Open()
		defer file.Close()
		if err != nil {
			// Log.Error(err)
			// s.ClientError(ec.ERR_UPLOAD_FAILED)
			return nil, err
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, file)
		if err != nil {
			// Log.Error(err)
			// s.ClientError(ec.ERR_UPLOAD_FAILED)
			return nil, err
		}

		originalFileName := f.Filename
		ext := path.Ext(originalFileName)
		localPath := "/img"
		filename := fmt.Sprintf("%v%v", formInfo.FileName, ext)
		filepath := fmt.Sprintf("/%v", formInfo.StoreId)
		filepathName := fmt.Sprintf("%v/%v", filepath, filename)
		// rstr := random2.RandomAlphabetic(4)
		// filename := fmt.Sprintf("%v%v%v", rstr, time.Now().UnixNano(), ext)
		// filepath := fmt.Sprintf("/%v/%v", sess.UserType, sess.LoginID)
		// filepathName := fmt.Sprintf("/%v/%v/%v", sess.UserType, sess.LoginID, filename)

		// localPath := configure.GetItem(params.CFG_UPLOAD_LOCAL_PATH).String()
		// Log.Debugf("localPath = %v, filepath = %v", localPath, filepath)
		err = os.MkdirAll("."+localPath+filepath, os.ModePerm)
		if err != nil {
			// Log.Error(err)
			// s.ClientError(ec.ERR_UPLOAD_FAILED)
			return nil, err
		}

		err = ioutil.WriteFile("."+localPath+filepathName, buf.Bytes(), 0644)
		if err != nil {
			// Log.Error(err)
			// s.ClientError(ec.ERR_UPLOAD_FAILED)
			return nil, err
		}

		// urlPrefix := configure.GetItem(params.CFG_UPLOAD_HTTP_PREFIX).String()
		upfile := &UploadFile{
			// OriginalFileName: originalFileName,
			// NewFileName:      filename,
			Url: "/v1" + localPath + filepathName,
		}

		uploadFiles = append(uploadFiles, upfile)
	}
	return uploadFiles, nil
	// s.Success(uploadFiles)
}

package handler

import (
	"fmt"
	"github.com/google/uuid"
	"imagego-go-api/database"
	"imagego-go-api/util"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseMultipartForm(100 << 20) // 100MB 제한
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err, responseCode := createImageInfo(req)
	if err != nil {
		http.Error(res, err.Error(), responseCode)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func createImageInfo(req *http.Request) (error, int) {
	file, header, err := req.FormFile("image")
	if err != nil {
		return err, http.StatusBadRequest
	}
	defer file.Close()

	userId := req.FormValue("userId")
	title := req.FormValue("title")
	description := req.FormValue("description")

	filename := fmt.Sprintf("%s%s", uuid.New(), filepath.Ext(header.Filename))
	err = imageCopy(file, filename)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = imageDBInsert(userId, title, description, filename)
	if err != nil {
		// db insert 실패하면 생성된 파일을 지운다.
		os.Remove(fmt.Sprintf(util.GetServerConfig().ImageDir+"/%s", filename))
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func imageCopy(targetFile multipart.File, destFilename string) error {
	imageDir := util.GetServerConfig().ImageDir

	// 현재 하위 디렉토리에 image 폴더를 생성한다.
	err := os.MkdirAll(imageDir, os.ModePerm)
	dst, err := os.Create(fmt.Sprintf(imageDir+"/%s", destFilename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, targetFile)
	if err != nil {
		return err
	}

	return nil
}

func imageDBInsert(userId, title, description, filename string) error {
	image := database.NewImage()

	image.UserID = userId
	image.Title = title
	image.Description = description
	image.ImageName = filename

	err := image.Create()
	if err != nil {
		return err
	}

	return nil
}

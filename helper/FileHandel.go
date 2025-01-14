package helper

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/crypto/scrypt"
)

func SaveFileToDestination(sub, sem, uni string, r *http.Request) (string, error) {
	fmt.Println("sub ", sub, " sem ", sem, " uni ", uni)

	_, file, err := r.FormFile("bookfile")

	if err != nil {
		fmt.Println("Error in getting file", err)
		return "", err
	}

	src, err := file.Open()

	if err != nil {
		fmt.Println("Error in opening file", err)
		return "", err
	}
	defer src.Close()

	name, _ := os.Getwd()
	name += "/static/bookinfo/pdf"
	newname, fileextension := createFileHash(file.Filename)
	newfilename := filepath.Base(sub + sem + uni + newname + fileextension)
	newfilepath := filepath.Join(name, newfilename)

	dst, err := os.Create(newfilepath)
	if err != nil {
		fmt.Println("Error in creating file", err)
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("Error in copying file", err)
		return "", err
	}

	if err != nil {
		fmt.Println("Error in getting bookfile", err)
		return "", err
	}

	return newfilename, nil
}
func SaveImgToDestination(sub, sem, uni string, r *http.Request) (string, error) {
	_, file, err := r.FormFile("bannerimage")
	if err != nil {
		fmt.Println("Error in getting file", err)
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("Error in opening file", err)
		return "", err
	}
	defer src.Close()

	name, _ := os.Getwd()
	name += "/static/bookinfo/img"
	newname, fileextension := createFileHash(file.Filename)
	newfilename := filepath.Base(sub + sem + uni + newname + fileextension)
	newfilepath := filepath.Join(name, newfilename)
	dst, err := os.Create(newfilepath)

	if err != nil {
		fmt.Println("Error in creating file", err)
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("Error in copying file", err)
		return "", err
	}

	if err != nil {
		fmt.Println("Error in getting bookfile", err)
		return "", err
	}

	return newfilename, nil
}

func createFileHash(filename string) (filenames, extesion string) {
	salt := []byte(os.Getenv("FILESALT"))

	const (
		MEMORYCOST = 16384
		THREADS    = 8
		KEYLENGTH  = 32
	)

	hashpwd, err := scrypt.Key([]byte(filename), salt, MEMORYCOST, THREADS, 1, KEYLENGTH)

	extesion = filepath.Ext(filename)
	filenames = hex.EncodeToString(hashpwd)

	if err != nil {
		log.Println(err)
	}

	return filenames, extesion
}

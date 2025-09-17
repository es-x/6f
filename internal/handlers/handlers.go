package handlers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/es-x/6f/internal/service"
)

func MainHandler(res http.ResponseWriter, req *http.Request) {

	// получаем новую текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	s := http.FileServer(http.Dir(curDir))
	s.ServeHTTP(res, req)

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	// получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	// закрываем файл
	defer file.Close()

	root, err := os.OpenRoot("../")
	if err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	scanner := bufio.NewScanner(file)
	var res string
	for scanner.Scan() {

		res, err = service.Converter(scanner.Text())
		if err != nil {
			fmt.Printf("error convert: %s\n", err.Error())
			return
		}
		io.WriteString(w, res)
	}
	nameFile := fmt.Sprintf("%s.log", time.Now().UTC().String())
	resFile, err := os.OpenFile(nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		http.Error(w, "error create or open file", http.StatusInternalServerError)
	}
	// отложенное закрытие файла
	defer resFile.Close()

	_, err = resFile.WriteString(res)
	if err != nil {
		http.Error(w, "error write data", http.StatusInternalServerError)
	}
}

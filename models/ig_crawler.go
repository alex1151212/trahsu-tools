package models

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"trahsu-tools/utils"
)

// https://instaloader.github.io/

/*
	1.商業帳號 -- 無須登入
	2.私人帳號 -- 需先登入
*/

// 私人帳號 -- 需先登入

/* LOGIN ERROR */

// func GetPrivatePost(userId string, username string, password string) {
// 	loginCmdStr := "--login=" + username
// 	passwordCmdStr := "--password=" + password
// 	cmd := exec.Command("instaloader", loginCmdStr, passwordCmdStr, userId)
// 	output, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Command output: %s", output)
// }

// 商業帳號 -- 無須登入
func GetPublicPost(userId string) ([]string, int) {

	// 切換至目標位置
	// utils.ChDir("C:/xampp/htdocs")

	cmd := exec.Command("python", ".\\utils\\crawler\\ig.py", userId)
	// _, err := cmd.CombinedOutput()
	// if err != nil {
	// 	panic(err)
	// }

	// 執行 Python 腳本
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("無法獲取輸出流。")
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("無法啟動腳本。")
		panic(err)
	}

	imgSrcChan := make(chan string)

	// 讀取輸出流
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		imgSrcChan <- scanner.Text()

	}

	// 檢查是否有錯誤
	if err := cmd.Wait(); err != nil {
		fmt.Println("腳本執行錯誤。")
		panic(err)
	}

	return PostFileName(userId)
}

func PostFileName(userId string) ([]string, int) {
	amount := 0
	path := "C:/xampp/htdocs/" + userId
	exts := []string{".jpg"}
	amount, filesName := utils.FindDirFileName(path, exts)
	var imgSrc []string = make([]string, amount)
	for i, fileName := range filesName {
		imgSrc[i] = "http://localhost:80/" + userId + "/" + fileName
		amount += 1
	}
	return imgSrc, amount
}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

func GetPublicPostHandler(writer http.ResponseWriter, request *http.Request) []string {
	query := request.URL.Query()
	userId := query.Get("userId")

	cmd := exec.Command("python", ".\\utils\\crawler\\ig.py", userId)

	// 執行 Python 腳本
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("無法獲取輸出流。")
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("無法啟動腳本。")
		panic(err)
	}

	imgSrcList := make([]string, 0)
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		imgSrcList = append(imgSrcList, scanner.Text())
		fmt.Println(scanner.Text())
	}

	// 檢查是否有錯誤
	if err := cmd.Wait(); err != nil {
		fmt.Println("腳本執行錯誤。")
		panic(err)
	}

	return imgSrcList
}

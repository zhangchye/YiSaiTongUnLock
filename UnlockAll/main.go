package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	path, _ := os.Executable()
	_, selfName := filepath.Split(path)
	str, _ := os.Getwd()
	allFile, _ := getAllFileIncludeSubFolder(str)
	for _, path := range allFile {
		if strutil.AfterLast(path, "Unlock.exe") == "" {
			continue
		}
		if strutil.AfterLast(path, selfName) == "" {
			continue
		}
		dstFilePath := path + ".temp"
		copyFile(path, dstFilePath)
		err := os.Remove(path)
		if err != nil {
			log.Printf("文件%v未执行成功", path)
		}
		renameFile(dstFilePath, path)
	}
	log.Println("解密完成，按回车键退出")
	fmt.Scanln()
}

func renameFile(sourcePath, dstFilePath string) {
	str, _ := os.Getwd()
	unlockPath := filepath.Join(str, "Unlock.exe")
	arg := fmt.Sprintf(` -sourcePath="%v" -destPath="%v"`, sourcePath, dstFilePath)
	cmd := exec.Command(unlockPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c" + arg}
	output, err := cmd.Output()
	if err != nil {
		log.Println("Failed to run command:", err)
	} else {
		info := string(output)
		if info != "" {
			log.Println(string(output))
		}
	}
}

func copyFile(sourcePath, dstFilePath string) (err error) {
	source, _ := os.Open(sourcePath)
	destination, _ := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer source.Close()
	defer destination.Close()
	buf := make([]byte, 1024)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

//
// getAllFileIncludeSubFolder
//  @Description: 获取目录下所有文件（包含子目录）
//  @param folder
//  @return []string
//  @return error
//
func getAllFileIncludeSubFolder(folder string) ([]string, error) {
	var result []string
	filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}
		if !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result, nil
}

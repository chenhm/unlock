// Copyright (C) 2013 chenhm. All rights reserved.

package main

import (
	//"bytes"
	"code.google.com/p/mahonia"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/*
亿赛通解锁小工具
	亿赛通解密规则是通过判断程序名和扩展名实现的，所以只需要修改程序名符合规则即可正确读出数据，
	例如本程序中RTX.exe，你也可以根据自己公司的规则修改程序名，比如用winword.exe打开.doc。
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:\r\n\t" + os.Args[0] + " file1 file2")
		return
	}
	my := os.Args[0]
	files := os.Args[1:]

	fin, err := os.Open(my)
	if err != nil && os.IsNotExist(err) {
		my = my + ".exe"
		fin, err = os.Open(my)
		if err != nil {
			fmt.Println(my, err)
			return
		}
	}
	defer fin.Close()
	//可执行文件就是UnxUtils中的cat命令，build时被append到了本程序之后，参见make.sh
	//cat.exe下载 http://sourceforge.net/projects/unxutils/
	exeFile := os.TempDir() + string(os.PathSeparator) + "7z.exe"

	fout, err := os.Create(exeFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(exeFile, err)
		return
	}
	buf := make([]byte, 1024)
	fin.Seek(-22016, os.SEEK_END)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		fout.Write(buf[:n])
	}
	fout.Close()
	fin.Close()

	encoder := mahonia.NewEncoder("gb18030")

	for i := 0; i < len(files); i++ {
		if err != nil {
			fmt.Println(err)
		}

		//cmd需要接受GBK编码的命令
		fileName := encoder.ConvertString(files[i])
		cmd := exec.Command("cmd", "/Q")
		cmdstr := exeFile + " " + fileName + " > " + fileName + ".1l\r\n"

		cmd.Stdin = strings.NewReader(cmdstr)
		cmd.Stderr = os.Stdout
		cmd.Stdout = new(DevNull) //忽略输出流
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		t := time.Now().Unix()
		//生成一个带有时间戳的备份文件
		err = os.Rename(files[i], files[i]+"."+strconv.FormatInt(t, 16))
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(files[i]+".1l", files[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

type DevNull struct{}

func (DevNull) Write(p []byte) (int, error) {
	return len(p), nil
}

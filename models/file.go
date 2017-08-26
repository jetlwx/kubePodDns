package models

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Recoder struct {
	Ip   string //IP
	Name string //自定义域名
}

const (
	hostsfile = "/etc/hosts"
)

var HOSTS []Recoder
var Myfile *os.File

func ReadFile() (d []Recoder, ff *os.File) {
	f, err := os.OpenFile(hostsfile, os.O_APPEND|os.O_RDWR, 0666) //打开文件
	//defer f.Close()              //打开文件出错处理
	if nil == err {
		buff := bufio.NewReader(f) //读入缓存
		for {
			line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break //结束读取
			}

			if !strings.HasPrefix(line, "#") || !strings.HasPrefix(line, "127.0.0.1") || !strings.HasPrefix(line, "::1") {
				//	fmt.Print(line) //可以对一行进行处理
				str := strings.Fields(line)
				//fmt.Print("len= ", len(str))
				if len(str) == 2 {
					if len(str[0]) > 0 && len(str[1]) > 0 {
						r := Recoder{}
						r.Ip = str[0]
						r.Name = str[1]
						d = append(d, r)
					}
					//fmt.Print(r)
				}

			}

		}

	}
	return d, f
}

func UpdateFile(OldText, NewText string) error {
	path := hostsfile

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(buf)

	//替换
	newContent := strings.Replace(content, OldText, NewText, -1)

	//重新写入
	err1 := ioutil.WriteFile(path, []byte(newContent), 0)
	if err1 != nil {
		log.Println("更新时出错：", err1)
	}
	log.Println(OldText, "更新为：", NewText)
	return nil
}

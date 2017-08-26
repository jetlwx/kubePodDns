package main

import (
	"fmt"
	"log"

	"flag"
	"strings"

	"github.com/Diggs/go-http-stream-reader"
	"github.com/bitly/go-simplejson"
	"github.com/jetlwx/kubePodDns/models"
)

var (
	domain  string
	etcdurl string
	kubeapi string
)

func Init() {
	flag.StringVar(&domain, "domain", "leyoujia.online", "doamin name of you set")
	flag.StringVar(&etcdurl, "etcdurl", "http://127.0.0.1:2379", "etcd url eg: http://127.0.0.1:2379,http://192.168.0.1:2379,http://192.168.0.2:2379")
	flag.StringVar(&kubeapi, "kubeapi", "http://127.0.0.1:8080", "the kube master api ,just support http now!!,sorry,default:http://127.0.0.1:8080")
}
func main() {
	fmt.Println("Version:20170822")
	fmt.Println("Attention: the etcd API version must be 2")
	Init()
	flag.Parse()
	url := kubeapi + "/api/v1/watch/pods"
	s := stream.NewStream(url)
	//s.Headers = map[string]string{"Authorization":"foobar"}
	s.Connect()
	for {
		select {
		case data := <-s.Data:
			//log.Println("Data: %s", string(data))
			Show2(data, etcdurl, domain)

		case err := <-s.Error:
			log.Println("Stream Error: %v", err)
		case <-s.Exit:
			log.Println("Stream closed.")
			return
		}
	}

}

func Show2(d []byte, etcdurl, domain string) {
	//	var cpodName string
	js, er := simplejson.NewJson(d)
	if er != nil {
		fmt.Println("json error:", er)
		return
	}
	//fmt.Println("js=", js)
	//get opertion kind ,mybe added ,modyfied ,etc.
	kind, err := js.Get("type").String()
	if err != nil {
		log.Println("get opertion type error:", err)
		return
	}
	//get containers's hostname
	hostname, err := js.GetPath("object", "metadata", "name").String()
	if err != nil {
		log.Println("get hostname error:", err)
		return
	}

	//get containers name ,may be mutil numbers
	containerNames, _ := js.GetPath("object", "status", "containerStatuses").Array()
	name := []string{}
	for _, v := range containerNames {
		//fmt.Println(v.(map[string]interface{})["name"])

		if c, ok := v.(map[string]interface{})["name"].(string); ok {
			name = append(name, c)

		}
	}

	//get container ip
	ip, err := js.GetPath("object", "status", "podIP").String()
	if err != nil {
		log.Println("get host ip error:", err)
		return
	}

	if len(kind) == 0 || len(hostname) == 0 || len(ip) == 0 {
		return
	}

	log.Println("[ kube API said ] kind=", kind, "containerNames=", name, "hostname=", hostname, "ip=", ip)
	etcd := gentetcd(etcdurl)
	models.Do(etcd, kind, name, domain, hostname, ip)
	fmt.Println("")
}

func gentetcd(etcdurl string) (s []string) {
	str := strings.Split(etcdurl, ",")
	for _, v := range str {
		s = append(s, v)
	}
	return s
}

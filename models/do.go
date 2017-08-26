package models

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/coreos/etcd/client"
)

type CNAME struct {
	Name string
	Host string
}
type HOST struct {
	Host string
}

/*
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
	Error    EventType = "ERROR"
etcdctl set /skydns/local/jjshome/dubbo/ask-and-answer-1357188408-0nzd3 '{"Name":"dubbo","Host":"ask-and-answer-1357188408-0nzd3.jjshome.local"}'
//etcdctl set /skydns/local/jjshome/ask-and-answer-1357188408-0nzd3 '{"Host":"192.168.1.46"}'
*/
func Do(etcdurl []string, kind string, containerNames []string, domain, hostname, ip string) {
	kapi := NewKeyAPI(NewCli(etcdurl))
	for _, v := range containerNames {
		cnamekey := "/skydns/" + ReversalDomain(domain) + "/" + v + "/" + hostname
		hostkey := "/skydns/" + ReversalDomain(domain) + "/" + hostname
		//	log.Println("etcdkey=", cnamekey)

		switch kind {
		case "MODIFIED", "ADDED":
			//setcname
			err := SetCname(kapi, cnamekey, v, hostname, domain)
			if err != nil {
				continue
			}
			SetHostname(kapi, cnamekey, hostkey, hostname, ip)

		case "DELETED":
			_, err := DelKey(kapi, cnamekey)
			if err != nil {
				log.Println("error at del cnanme key", err)
				continue
			}
			log.Println("[DEL1]", cnamekey)

			_, err3 := DelKey(kapi, hostkey)
			if err3 != nil {
				log.Println("error at del hostkey", err3)
				continue
			}
			log.Println("[DEL2]", hostkey)
			//	log.Println("respkey=", respk)

		}
	}
}

//set etcd key,CNAME
/*
etcdctl set  \
/skydns/local/jjshome/dubbo/ask-and-answer-1357188408-0nzd3 \
 '{"Name":"dubbo","Host":"ask-and-answer-1357188408-0nzd3.jjshome.local"}'

*/
func SetCname(kapi client.KeysAPI, cnamekey, vkey, vvalue, domain string) error {
	c := CNAME{}
	c.Name = vkey
	c.Host = vvalue + "." + domain
	js, err := json.Marshal(c)
	if err != nil {
		log.Println("error at SetCname:", err)
		return err
	}

	_, err1 := SetKey(kapi, cnamekey, string(js))
	if err1 != nil {
		log.Println("error at set key:", err1)
		return err
	}
	//log.Println("resp=", resp)
	log.Println("[SET]", cnamekey, "-->", string(js))
	return nil
}

//etcdctl set /skydns/local/jjshome/dubbo3 '{"Host":"192.168.1.46"}'
func SetHostname(kapi client.KeysAPI, cnamekey, hostkey, hostname, ip string) {
	h := HOST{}
	h.Host = ip
	js, err := json.Marshal(h)
	if err != nil {
		log.Println("error at set etcdkey", err)
		// if err ,will delete cname relation hostname ,then return
		//delete
		_, err1 := DelKey(kapi, cnamekey)
		if err1 != nil {
			log.Println("error at delete cname key", err1)
			return
		}
		log.Println("[DEL3]", cnamekey)
		return
	}

	_, err1 := SetKey(kapi, hostkey, string(js))
	if err1 != nil {
		log.Println("error at Set host name:", err1)
		return
	}
	//	log.Println("resp=", resp, "err=", err)
	log.Println("[SET]", hostkey, "-->", ip)
	return
}

//gent the cname
func gentCname(name, host string) (string, error) {
	c := CNAME{}
	c.Host = host
	c.Name = name
	res, err := json.Marshal(c)
	return string(res), err
}

//gent the host json
func gentHost(value string) (string, error) {
	h := HOST{}
	h.Host = value
	res, err := json.Marshal(h)
	return string(res), err
}

//jjshome.com --> com/jjshome
func ReversalDomain(d string) (d2 string) {
	s := strings.Split(d, ".")
	s1 := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		s1 = append(s1, s[i])
	}

	return strings.Join(s1, "/")
}

//set the cname
/*
etcdctl set /skydns/local/jjshome/dubbo/1 '{"Name":"dubbo","Host":"dubbo1.jjshome.local"}'
etcdctl set /skydns/local/jjshome/dubbo/2 '{"Name":"dubbo","Host":"dubbo2.jjshome.local"}'
etcdctl set /skydns/local/jjshome/dubbo/3 '{"Name":"dubbo","Host":"dubbo3.jjshome.local"}'
*/
// func OPTCname(dir, podName, domain string, ip []string, kapi client.KeysAPI) map[string]string {
// 	cname := make(map[string]string)
// 	for k, v := range ip {
// 		if len(v) == 0 {
// 			continue
// 		}

// 		kstr := strconv.Itoa(k)
// 		h := podName + kstr + "." + domain
// 		value, err := gentCname(podName, h)

// 		if err != nil {
// 			log.Println("gentCname error:", err)
// 			continue
// 		}

// 		key := dir + "/" + kstr

// 		if len(key) == 0 || len(value) == 0 {
// 			continue
// 		}

// 		//set key
// 		resp, err := SetKey(kapi, key, value)
// 		if err != nil {
// 			log.Println("error at set key:", err)
// 			continue
// 		}
// 		k := dir + "/" + podName + kstr
// 		cname[k] = v
// 		log.Println("resp=", resp)
// 	}
// 	return cname
// }

// func CheckIPExist(ip string, domain string, hosts []Recoder) bool {
// 	for _, h := range hosts {
// 		if ip == h.Ip && h.Name == domain {
// 			return true
// 		}
// 	}

// 	return false
// }

// func Findsuffdomain(suff string, hosts []Recoder) (h []Recoder) {
// 	for _, v := range hosts {
// 		if suff == v.Name {
// 			h2 := Recoder{}
// 			h2.Ip = v.Ip
// 			h2.Name = v.Name
// 			h = append(h, h2)
// 		}
// 	}

// 	return h
// }

// func DeleteInfo(f *os.File, s string) {
// 	buff := bufio.NewReader(f) //读入缓存
// 	for {
// 		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
// 		if err != nil || io.EOF == err {
// 			break //结束读取
// 		}
// 		if strings.HasSuffix(line, s) {
// 			log.Println("line=", line)
// 		}
// 	}
// 	return
// }

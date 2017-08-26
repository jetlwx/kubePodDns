package models

type Info struct {
	Type   string
	Domain string
	Ip     string
}

//从接口获取IP信息
func IP(subs []interface{}) (ip []string) {
	for _, v := range subs {
		if a, ok := v.(map[string]interface{})["addresses"]; ok {
			//fmt.Println("a=", a)
			if c, ok := a.([]interface{}); ok {
				//fmt.Println("----------------c=--------", c)
				for _, c1 := range c {
					if c2 := c1.(map[string]interface{})["ip"]; ok {
						//fmt.Println("c2",c2)
						if c2 != nil {
							if value, ok := c2.(string); ok {
								ip = append(ip, value)
							}
						}
					}
				}
			}
		}

	}
	return ip
}

//生成域名IP对应表
func GentList(typ string, domain string, ip []string) (R []Info) {
	for _, v := range ip {
		r := Info{}
		r.Domain = domain
		r.Type = typ
		r.Ip = v
		R = append(R, r)
	}
	return R
}

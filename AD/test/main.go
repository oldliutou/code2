package main

import (
    "fmt"
    "github.com/go-ldap/ldap/v3"
	"log"
)

func main() {
	// 设置LDAP连接信息
	l, err := ldap.Dial("tcp", "192.168.10.15:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// 进行身份验证
	// CN=刘节威\\Jiewei Liu,OU=IT部,OU=科技公司,DC=labnvr,DC=cn
	err = l.Bind("jiewei.liu@labnvr.cn", "123.com")
	if err != nil {
		log.Fatal(err)
	}

	// 搜索示例
	searchRequest := ldap.NewSearchRequest(
		"OU=科技公司,DC=labnvr,DC=cn", // 搜索根
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=organizationalPerson)", // 过滤条件
		[]string{"dn", "cn", "mail"},        // 要检索的属性
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	fmt.Println("Search Results:")
	for _, entry := range sr.Entries {
		fmt.Printf("DN: %s\n", entry.DN)
		for _, attr := range entry.Attributes {
			fmt.Printf("%s: %v\n", attr.Name, attr.Values)
		}
	}
}
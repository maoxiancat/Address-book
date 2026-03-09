package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Contact 结构体表示一个联系人
type Contact struct {
	Name    string
	Phone   string
	Address string
}

var contacts []Contact

// 添加联系人
func addContact() {
	var name, phone, address string
	// 获取信息
	fmt.Print("请输入姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入电话号码: ")
	fmt.Scanln(&phone)
	fmt.Print("请输入住址: ")
	fmt.Scanln(&address)

	// 添加信息
	contact := Contact{Name: name, Phone: phone, Address: address}
	key := rankContactsByName(name)
	if key != 0 {
		contactsss := contacts[key:]
		var contactssss []Contact
		contactssss = append(contactssss, contact)
		contactsss = append(contactssss, contactsss...)
		contactss := contacts[:key]
		contacts = append(contactss, contactsss...)
	} else if key == len(contacts) {
		contacts = append(contacts, contact)
	} else {
		var contactss []Contact
		contactss = append(contactss, contact)
		contacts = append(contactss, contacts...)
	}
	fmt.Println("成功添加联系人")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

// 使添加联系人根据姓名拼音进行排序保存
func rankContactsByName(name string) (key int) {
	for key, contact := range contacts {
		if strings.ToLower(contact.Name)[0] > strings.ToLower(name)[0] {
			return key
		}
	}
	return len(contacts)
}

// 显示所有联系人
func displayContacts() {
	if len(contacts) == 0 {
		fmt.Println("无任何联系人")
		fmt.Println()
		fmt.Println()
		fmt.Println()
		return
	}
	for key, contact := range contacts {
		fmt.Printf("%d 姓名:%s 电话号码：%s 住址：%s\n", key+1, contact.Name, contact.Phone, contact.Address)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

// 通过姓名查找联系人(精确)
func findContactByName1() {
	var name string
	fmt.Print("请输入名字来查询: ")
	fmt.Scanln(&name)

	for _, contact := range contacts {
		//转小写再比较
		if strings.ToLower(contact.Name) == strings.ToLower(name) {
			fmt.Printf("查询结果: 姓名：%s 电话号码：%s 住址：%s\n", contact.Name, contact.Phone, contact.Address)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			return
		}
	}
	fmt.Println("无查找到相关联系人")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

// 通过电话号码查找联系人(精确)
func findContactByPhone1() {
	var phone string
	fmt.Print("请输入电话号码来查询: ")
	fmt.Scanln(&phone)

	for _, contact := range contacts {
		if contact.Phone == phone {
			fmt.Printf("查询结果: 姓名：%s 电话号码：%s 住址：%s\n", contact.Name, contact.Phone, contact.Address)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			return
		}
	}
	fmt.Println("无查找到相关联系人")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

// 通过住址查找联系人(精确)
func findContactByAddress1() {
	var address string
	fmt.Print("请输入住址来查询: ")
	fmt.Scanln(&address)

	for _, contact := range contacts {
		if contact.Address == address {
			fmt.Printf("查询结果: 姓名：%s 电话号码：%s 住址：%s\n", contact.Name, contact.Phone, contact.Address)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			return
		}
	}
	fmt.Println("无查找到相关联系人")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

// 通过姓名查找联系人(模糊)
func findContactByName2() {
	var name string
	fmt.Print("请输入部分名字来查询: ")
	fmt.Scanln(&name)

	var n int = 0
	for _, contact := range contacts {
		if strings.Contains(strings.ToLower(contact.Name), strings.ToLower(name)) {
			fmt.Printf("查询结果: 姓名：%s 电话号码：%s 住址：%s\n", contact.Name, contact.Phone, contact.Address)
			n++
		}
	}
	if n == 0 {
		fmt.Println("无查找到相关联系人")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return
}

// 通过电话号码查找联系人(模糊)
func findContactByPhone2() {
	var phone string
	fmt.Print("请输入部分电话号码来查询: ")
	fmt.Scanln(&phone)

	var n int = 0
	for _, contact := range contacts {
		if strings.Contains(strings.ToLower(contact.Phone), strings.ToLower(phone)) {
			fmt.Printf("查询结果: 姓名：%s 电话号码：%s 住址：%s\n", contact.Name, contact.Phone, contact.Address)
			n++
		}
	}
	if n == 0 {
		fmt.Println("无查找到相关联系人")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return
}

// 通过住址查找联系人(模糊)
func findContactByAddress2() {
	var address string
	fmt.Print("请输入部分住址来查询: ")
	fmt.Scanln(&address)

	var n int = 0
	for _, contact := range contacts {
		if strings.Contains(strings.ToLower(contact.Address), strings.ToLower(address)) {
			fmt.Printf("查询结果: 姓名：%s 电话号码：%s 住址：%s\n", contact.Name, contact.Phone, contact.Address)
			n++
		}
	}
	if n == 0 {
		fmt.Println("无查找到相关联系人")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return
}

// 修改联系人
func postContact() {
	var name, newName, newPhone, newAddress string
	fmt.Print("请输入姓名来进行对应的修改: ")
	fmt.Scanln(&name)

	for key, contact := range contacts {
		if strings.ToLower(contact.Name) == strings.ToLower(name) {
			fmt.Print("请输入新的姓名: ")
			fmt.Scanln(&newName)
			fmt.Print("请输入新的电话号码: ")
			fmt.Scanln(&newPhone)
			fmt.Print("请输入新的住址: ")
			fmt.Scanln(&newAddress)

			contacts[key] = Contact{Name: newName, Phone: newPhone, Address: newAddress}
			fmt.Println("成功修改联系人")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			return
		}
	}
	fmt.Println("无查找到相关联系人")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

// 删除联系人
func deleteContact() {
	var name string
	fmt.Print("请输入姓名来进行对应的删除: ")
	fmt.Scanln(&name)

	for key, contact := range contacts {
		if strings.ToLower(contact.Name) == strings.ToLower(name) {
			contacts = append(contacts[:key], contacts[key+1:]...)
			fmt.Println("成功删除联系人")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			return
		}
	}
	fmt.Println("无查找到相关联系人")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func main() {
	for {
		fmt.Println("--- 通讯录 ---")
		fmt.Println("输入“添加”来添加联系人")
		fmt.Println("输入“查询全部”来查询所有联系人")
		fmt.Println("输入“查找”来查找某个联系人")
		fmt.Println("输入“修改”来修改联系人")
		fmt.Println("输入“删除”来删除联系人")
		fmt.Println("输入“退出”来退出程序")
		fmt.Print("请输入要执行的内容：")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "添加":
			addContact()
		case "查询全部":
			displayContacts()
		case "查找":
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("请选择查找方式")
			fmt.Println("输入“精确”来进行精确查找")
			fmt.Println("输入“模糊”来进行模糊查找")
			fmt.Print("请输入要执行的内容：")
			var way string
			fmt.Scanln(&way)
			switch way {
			case "精确":
				fmt.Println("输入“姓名”来通过姓名进行查找")
				fmt.Println("输入“号码”来通过电话号码进行查找")
				fmt.Println("输入“住址”来通过住址进行查找")
				fmt.Print("请输入要执行的内容：")
				var method string
				fmt.Scanln(&method)

				switch method {
				case "姓名":
					findContactByName1()
				case "电话号码":
					findContactByPhone1()
				case "住址":
					findContactByAddress1()
				default:
					fmt.Println("无效的查找方式")
					fmt.Println()
					fmt.Println()
					fmt.Println()
				}
			case "模糊":
				fmt.Println("输入“姓名”来通过姓名进行查找")
				fmt.Println("输入“号码”来通过电话号码进行查找")
				fmt.Println("输入“住址”来通过住址进行查找")
				fmt.Print("请输入要执行的内容：")
				var method string
				fmt.Scanln(&method)

				switch method {
				case "姓名":
					findContactByName2()
				case "电话号码":
					findContactByPhone2()
				case "住址":
					findContactByAddress2()
				default:
					fmt.Println("无效的查找方式")
					fmt.Println()
					fmt.Println()
					fmt.Println()
				}
			default:
				fmt.Println("无效的查找方式")
				fmt.Println()
				fmt.Println()
				fmt.Println()
			}
		case "修改":
			postContact()
		case "删除":
			deleteContact()
		case "退出":
			fmt.Println("退出中...")
			time.Sleep(time.Second * 3 / 2)
			os.Exit(0)
			return
		default:
			fmt.Println("无效输入")
		}
	}
}

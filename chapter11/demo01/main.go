package main

import "fmt"

//定义一个账户的结构体的字段和行为
func main() {

	account := Account{
		AccountNo: "gs111111",
		Pwd:       "6666",
		Balance:   1000,
	}

	account.Query("6666")
	account.Deposite(200, "6666")
}

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//1 存款的方法
func (account *Account) Deposite(money float64, pwd string) {

	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功~~~")
}

//2 取款的方法
func (account *Account) WithDraw(money float64, pwd string) {

	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功~~~")
}

//3 查询余额的方法
func (account *Account) Query(pwd string) {

	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n", account.AccountNo, account.Balance)
}

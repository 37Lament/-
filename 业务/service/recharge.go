package main

func main() {
	func ChangeRMB(usename string,newrmb int)error{
		err := dao.UpdatePassword(usename,newrmb)
		return err
	}
}

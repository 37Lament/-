package main

func main() {
	func UpdatePassword(username, newPassword string) error {
		_, err := dB.Exec("UPDATE user SET rmb = ? WHERE username = ?", newrmb, username)
		return err
	}

}

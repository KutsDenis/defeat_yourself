package main

import "defeat_yourself/internal/update"

func main() {
	var u update.Updater

	u = update.Update{}
	u.Update()
}

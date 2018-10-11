package main

import (
	"fmt"
	keychain "github.com/FreekingDean/go-keychain"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService("fixops")
	item.SetAccount(user.Username)
	item.SetMatchLimit(keychain.MatchLimitOne)
	list, err := keychain.GetACLList(item)
	fmt.Println(err)
	for _, acl := range list {
		fmt.Println(
		fmt.Printf("%+v\n", acl)
	}
}

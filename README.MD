# Go MAIB ECommerce Bridge

    Copyright (C) 2023 Nejintev Nicolai

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

This project provides tools to interact with MAIB ECommerce in a type safe way.

This library is an **UNOFFICIAL** implementation, inspired by **public** Integration Guide and
[implementations](https://github.com/maibank) shared by MAIB.

### Simple example
```go
package main

import (
	"fmt"
	maib "github.com/NikSays/go-maib-ecomm"
)

func main() {
	client, err := maib.NewClient(maib.Config{
		PFXPath:                 "cert.pfx",
		Passphrase:              "Passphrase",
		MerchantHandlerEndpoint: "https://example.org:443/handler",
	})
	res1, err := client.RegisterTransaction(maib.SMS, maib.RegisterPayload{
		Amount:          1,
		Currency:        maib.CurrencyMDL,
		ClientIPAddress: "127.0.0.1",
		Description:     "No description",
		Language:        maib.LanguageEnglish,
	})
	if err != nil {
		fmt.Printf("ERR: %s", err.Error())
		return
	} else {
		fmt.Printf("Ok: %s", res1.TransactionID)
	}
	res2, err := client.ReverseTransaction(maib.ReversePayload{
		Amount:        1,
		TransactionID: res1.TransactionID,
	})
	if err != nil {
		fmt.Printf("ERR: %s", err.Error())
		return
	} else {
		fmt.Printf("Ok: %s", res2.Result)

	}
	res3, err := client.CloseDay()
	if err != nil {
		fmt.Printf("ERR: %s", err.Error())
	} else {
		fmt.Printf("Ok: %s", res3.Result)
	}
}
```

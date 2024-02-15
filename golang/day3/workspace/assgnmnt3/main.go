package main

import (
	"fmt"
)

func main() {
	dao := CustomerDAO{}
	dao.LoadFromJsonFile()
	for {
		choice := Menu()
		if choice == 0 {
			fmt.Println("Thank you, have a nice day.")
			break
		}

		switch choice {
		case 1:
			dao.LoadFromJsonFile()
		case 7:
			customers := dao.GetAllCustomers()

			fmt.Printf("+%5s +%-20s +%-20s +%-35s +\n",
				Chars('-', 5),
				Chars('-', 20),
				Chars('-', 20),
				Chars('-', 35))

			fmt.Printf("|%5s |%-20s |%-20s |%-35s |\n",
				"ID",
				"Name",
				"City",
				"Email")

			fmt.Printf("+%5s +%-20s +%-20s +%-35s +\n",
				Chars('-', 5),
				Chars('-', 20),
				Chars('-', 20),
				Chars('-', 35))

			
			for _, c := range customers {
				fmt.Printf("|%5d |%-20s |%-20s |%-35s |\n",
					c.Id, c.Name, c.City, c.Email)
			}

			fmt.Printf("+%5s +%-20s +%-20s +%-35s +\n",
				Chars('-', 5),
				Chars('-', 20),
				Chars('-', 20),
				Chars('-', 35))
		}
	}
}

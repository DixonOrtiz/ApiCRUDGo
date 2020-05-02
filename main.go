package main

import "fmt"

func main() {
	r1 := Raper{
		Name:    "Dixon",
		Country: "Chile",
		Age:     23,
	}

	i, err := CreateRaper(r1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Filas afectadas:", i)

}

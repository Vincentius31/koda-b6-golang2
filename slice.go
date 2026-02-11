package main

import "fmt"

// Nomor 1
type the struct {
	best string
}

type are struct {
	the the
}

type weType struct {
	are are
}

// Nomor 2
type helloType struct {
	world string
}

// Nomor 3
type tech struct {
	academy string
}

type man struct {
	tech tech
}

type stritem struct {
	man []man
}

type objtype struct {
	str [][][]stritem
}

// Nomor 4
type fruit struct {
	is string
}

type favourite struct {
	fruit fruit
}

type mytype struct {
	favourite []favourite
}

// Nomor 5
type numtype struct {
	first  []int
	second []int
}

func main() {
	//Hasil Nomor 1
	var we weType = weType{
		are: are{
			the: the{
				best: "Koda",
			},
		},
	}

	fmt.Println(we.are.the.best)

	//Hasil Nomor 2
	var hello helloType = helloType{
		world: "Hello World",
	}

	fmt.Println(hello.world)

	//Hasil Nomor 3
	var obj objtype = objtype{
		[][][]stritem{
			{},
			{},
			{},
			{
				{},
				{
					{},
					{},
					{
						man: []man{
							{
								tech: tech{
									academy: "Tech Academy",
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

	//Hasil Nomor 4
	my := []mytype{
		{
			favourite: []favourite{
				{},
				{},
				{},
				{
					fruit: fruit{
						is: "Apple",
					},
				},
			},
		},
	}

	fmt.Println(my[0].favourite[3].fruit.is)

	//Hasil Nomor 5

	var num numtype = numtype{
		first: []int{
			10,
			12,
		},
		second: []int{
			5,
			10,
			20,
		},
	}

	fmt.Println(num.first[1] + num.second[2])
}

package main

import "fmt"

func main() {
	fmt.Println("map - struct")

	// Map has key: string, value: int
	userAges := map[string]int{
		"Kiet":  12,
		"Hoa":   13,
		"Hoang": 15,
	}

	// *** CAN NOT USE SLICE AS KEY

	fmt.Println(userAges)

	// Get value from map
	fmt.Println(userAges["Kiet"])

	// Remove key, value from array
	delete(userAges, "Hoang")

	fmt.Println("After removing Hoang: ", userAges)

	// We can get value of removed key
	removedHoang := userAges["Hoang"]

	fmt.Println(removedHoang)

	// To avoid this we need to check exist

	hoangAge, isExist := userAges["Hoang"]

	fmt.Println(hoangAge, isExist)

	// Len of map is same with array
	fmt.Println("Map length is: ", len(userAges))

	// Copy map

	userAgesCopy := userAges

	// Try to change value of copied array
	userAgesCopy["Kiet"] = 800

	// After copy
	fmt.Println(userAges)
	fmt.Println(userAgesCopy)

	/**
		STRUCT
	**/
	student1 := Student{
		id:       "ST1",
		name:     "Kiet",
		subjects: []string{"Math", "English", "Database"},
	}

	fmt.Println(student1)

	// Update value
	student1.name = "Kiet Tan Le"
	fmt.Println(student1)

	// Struct in Struct
	student2 := NewStudent{
		id:       "ST2",
		name:     "Yen",
		subjects: []Subject{{id: 1, name: "Math"}},
	}

	fmt.Println(student2)

	// Copy struct will clone new object to new memory
	student2Copy := student2

	student2Copy.name = "Yen Fake"

	fmt.Println(student2)
	fmt.Println(student2Copy)

	// Use & to copy as same memory address
	student1Copy := &student1

	student1Copy.name = "Kiet Le"

	fmt.Println(student1)
	fmt.Println(student1Copy)

	// Embedded
	author := Author{
		// People{id: 1, name: "Kiet"},
		books: []string{"Harry Potter"},
	}
	author.name = "Kiet Le"
	author.id = 1

	fmt.Println(author)
}

type Student struct {
	id       string
	name     string
	subjects []string
}

type People struct {
	id   int
	name string
}

type NewStudent struct {
	id       string
	name     string
	subjects []Subject
}

type Subject struct {
	id   int
	name string
}

type Author struct {
	People
	books []string
}

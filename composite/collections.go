package composite

import "fmt"

func DemoCollections() {
	//Array
	var primesArray [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primesArray)

	//Slice
	// primesSlice := []int{2, 3, 5, 7, 11}
	primesSlice := make([]int, 6, 20)
	primesSlice = append(primesSlice, 7) //hängt eine 7 an
	fmt.Println(len(primesSlice))        //länge eines Slice
	fmt.Println(cap(primesSlice))        //kapazität ds Slice max größe

	for index, value := range primesSlice {
		fmt.Println(index, value)
	}

	for _, value := range primesSlice {
		fmt.Println(value) //_ nehmen wenn man variable brauch aber der wwert nicht interessiert
	}

	//Maps
	points := map[string]string{
		"A": "Maik",
		"B": "Peters",
	}
	fmt.Println(points)

	some, ok := points["A"]
	fmt.Println(some, ok)
}

package main

import "fmt"

const (
	PI = 3.14159
)

func main() {
	var (
		radius float64
		option int
	)

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanln(&radius)

	fmt.Println("what do you need to calculate ? ")
	fmt.Println("1: Area of the circle")
	fmt.Println("2  Perimeter of the circle")
	fmt.Println("3 - Diameter of the circle")
	fmt.Scanln(&option)
   
	func printResult (radius float64, getUserInput(option)) {

    }
}   

func getUserInput( option int ) float64 {

	func convert_option_to_func  map[int]func(float64) float64{
		1: circleArea,
		2: circlePerimeter,
		3: circleDiameter,
	}
	return convert_option_to_func[option](radius)

	// switch option {
	// case 1:
	// 	return circleArea(radius)
	// case 2:
	// 	return circlePerimeter(radius)
	// case 3:
	// 	return circleDiameter(radius)
	// default:
	// 	return radius
	// }
}

func circleArea(radius float64) float64 {
	fmt.Println("area of the circle :")
	return PI * radius * radius
}

//perimeter of a circle function
func circlePerimeter(radius float64) float64 {
	fmt.Println("perimeter of the circle:")
	return PI * radius * 2
}

// diameter of a circle function
func circleDiameter(radius float64) float64 {
	fmt.Println("diameter of the circle:")
	return 2 * radius
}

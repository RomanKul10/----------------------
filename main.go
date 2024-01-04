package main

import "fmt"

type Subscriber struct { //підписник
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address // створюємо параметр зі своєю структурою
}

type Address struct {
	City     string
	Street   string
	Building string
}

func printInfo(s Subscriber) { //в дужках (s- змінна, subscriber - тип змінної)
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Atcive: ", s.Active)
	fmt.Println("Address: ", s.HomeAddress)
}

func defaultSubscriber() Subscriber { //default- за замовчуванням (name - змінна, str- тип)
	// і повертає значення subscriber
	var s Subscriber // створюємо змінну - s
	s.Name = "Alex"
	s.Rate = 11.999
	s.Active = true
	s.HomeAddress.City = "Kyiv"
	return s // ось тут ми і повертаємо значення subscriber
}
func defaultSubscriber2(name string, rate float64, active bool, city string) Subscriber { // (змінні зі своїми типами)
	// і повертає значення subscriber
	var s Subscriber // створюємо змінну - s
	s.Name = name
	s.Rate = rate
	s.Active = active
	s.HomeAddress.City = city
	return s // ось тут ми і повертаємо значення subscriber

}

func applyDiscount(s *Subscriber) { // apply Discount - застосувати знижку
	// створюємо зміну s і передаємо вказівник *subscriber
	// бо якщо просто subscriber, ми його копіюємо і змінюємо, а так передаємо йому дані)
	s.Rate = 4.99
}

func main() {
	address1 := Address{
		City:     "Kyiv",
		Street:   "Zavas",
		Building: "233",
	}
	// var subscriber1 defaultSubscriber
	// або записуємо скорочено:
	subscriber1 := defaultSubscriber()
	subscriber1.HomeAddress = address1
	subscriber2 := defaultSubscriber2("Roma", 12.0005, true, "Lviv")

	applyDiscount(&subscriber1) //тут так само передаємо вказівник на 1
	printInfo(subscriber1)
	printInfo(subscriber2)
}

/*func defaultSubscriber(name string, rate float64, active bool) subscriber { // (name - змінна, str- тип)
	// і повертає значення subscriber
	var s subscriber // створюємо змінну - s
	s.name = name
	s.rate = rate
	s.active = active
	return s // ось тут ми і повертаємо значення subscriber

}

func main() {
	// var subscriber1 defaultSubscriber
	// або записуємо скорочено:
	subscriber1 := defaultSubscriber("Alex", 12.0005, true)
	printInfo(subscriber1)
}
*/

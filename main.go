package main

import "fmt"

type Computer struct {
	motherBoard []motheSpec
	graphicCard []graphSpec
	processor   []procSpec
	numOfFans   []Fans
}

type motheSpec struct {
	formFactor string
	socket     string
	chipset    string
	memory     string
}

type graphSpec struct {
	brand            string
	graphCoprocessor string
	graphRAMSize     int
	memoryClock      int
}

type procSpec struct {
	socket       string
	frequency    string
	numOfThreads int
}

//Number of funs in side of computer
type Fans struct {
	graphicCardFans int
	processorFans   int
	computerCase    int
}

func main() {

	ms := motheSpec{
		formFactor: "ATX",
		socket:     "AMD AM4",
		chipset:    "AMD B450",
		memory:     "Max RAM: 64 GB",
	}
	fmt.Println("motherBoard: "+("formFactor: "+ms.formFactor), ("socket: " + ms.socket))

}

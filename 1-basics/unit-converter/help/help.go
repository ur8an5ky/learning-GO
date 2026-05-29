package help

import "fmt"

func Help() {
	fmt.Println("Usage:")
	fmt.Println("\t<value> <unit1> to <unit2>     	- convert <value> from <unit1> to <unit2>")
	fmt.Println("\thelp                           	- show this message")
	fmt.Println("\thelp <category>                	- show available units for a category")
	fmt.Println("\tquit, exit                 		- exit the program")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("\t25 c to f")
	fmt.Println("\t10 km to mi")
	fmt.Println("\t500 g to lb")
	fmt.Println()
	fmt.Println("Available categories:")
	fmt.Println("\ttemperature     - Celsius, Fahrenheit, Kelvin")
	fmt.Println("\tlength          - metric and imperial length units")
	fmt.Println("\tweight          - metric and imperial weight units")
	fmt.Println()
	fmt.Println("Type 'help <category>' to see the unit codes for that category.")
}

func HelpTemperature() {
	fmt.Println("Temperature units:")
	fmt.Println("\tc    - Celsius")
	fmt.Println("\tf    - Fahrenheit")
	fmt.Println("\tk    - Kelvin")
	fmt.Println()
	fmt.Println("Example: 25 c to f")
}

func HelpLength() {
	fmt.Println("Length units:")
	fmt.Println("\tm     - metre  (base unit)")
	fmt.Println("\tkm    - kilometre")
	fmt.Println("\tcm    - centimetre")
	fmt.Println("\tmi    - mile")
	fmt.Println("\tft    - foot")
	fmt.Println()
	fmt.Println("Example: 10 km to mi")
}

func HelpWeight() {
	fmt.Println("Weight units:")
	fmt.Println("\tg     - gram  (base unit)")
	fmt.Println("\tkg    - kilogram")
	fmt.Println("\tt     - tonne")
	fmt.Println("\tmg    - milligram")
	fmt.Println("\tlb    - pound")
	fmt.Println("\toz    - ounce")
	fmt.Println()
	fmt.Println("Example: 500 g to lb")
}

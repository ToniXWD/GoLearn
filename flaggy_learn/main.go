package main

import "github.com/integrii/flaggy"

// Super Simple Example
func main1() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"

	// Add a flag
	flaggy.String(&stringFlag, "f", "flag", "A test string flag")

	// Parse the flag
	flaggy.Parse()

	// Use the flag
	print(stringFlag)
}

// Example with Subcommand
func main2() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"

	// Create the subcommand
	subcommand := flaggy.NewSubcommand("subcommandExample")

	// Add a flag to the subcommand
	subcommand.String(&stringFlag, "f", "flag", "A test string flag")

	// Add the subcommand to the parser at position 1
	flaggy.AttachSubcommand(subcommand, 1)

	// Parse the subcommand and all flags
	flaggy.Parse()

	// Use the flag
	print(stringFlag)
}

// Example with Nested Subcommands, Various Flags and Trailing Arguments
func main3() {
	// Declare variables and their defaults
	var stringFlagF = "defaultValueF"
	var intFlagT = 3
	var boolFlagB bool

	// Create the subcommands
	subcommandExample := flaggy.NewSubcommand("subcommandExample")
	nestedSubcommand := flaggy.NewSubcommand("nestedSubcommand")

	// Add a flag to both subcommands
	subcommandExample.String(&stringFlagF, "t", "testFlag", "A test string flag")
	nestedSubcommand.Int(&intFlagT, "f", "flag", "A test int flag")

	// add a global bool flag for fun
	flaggy.Bool(&boolFlagB, "y", "yes", "A sample boolean flag")

	// attach the nested subcommand to the parent subcommand at position 1
	subcommandExample.AttachSubcommand(nestedSubcommand, 1)
	// attach the base subcommand to the parser at position 1
	flaggy.AttachSubcommand(subcommandExample, 1)

	// Parse everything, then use the flags and trailing arguments
	flaggy.Parse()
	print(stringFlagF)
	print(intFlagT)
	print(boolFlagB)
	print(flaggy.TrailingArguments[0])
}

func main() {
	// main1()
	// main2()
	main3()
}

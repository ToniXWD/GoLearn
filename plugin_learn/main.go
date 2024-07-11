package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("adder.so")
	if err != nil {
		panic(err)
	}

	f, _ := p.Lookup("Add")
	add, ok := f.(func(int, int) int)
	if !ok {
		panic("could not assert to correct type")
	}

	fmt.Printf("Add function result: %d\n", add(3, 4))

	v, _ := p.Lookup("PluginVar")
	pluginVar, ok := v.(*int)
	if !ok {
		panic("could not assert to correct type")
	}

	fmt.Printf("PluginVar value: %d\n", *pluginVar)
}

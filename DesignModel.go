package main

type Color struct {
	Name string
	Value string
}

type Item struct {
	Name string
	Value string
}

type Style struct {
	Name string
	Parent string
	Items []Item
}
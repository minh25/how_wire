package main

func main() {
	e, err := InitializeEvent("help me")
	if err != nil {
		panic(err)
	}
	e.Start()
}

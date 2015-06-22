package main

type Logging interface{
	Log(message string) error
}

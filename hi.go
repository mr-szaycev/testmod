package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("Hi!, %s", name)
}

// By returns a bed girl
func ByBy(name string) string {
   return fmt.Sprintf("By fack !, #{name}")
}
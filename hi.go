package testmod

import (
	"fmt"
	"errors"
)

// Hi returns a friendly greeting
func Hi(name string, lang string) (string, error) {
   switch lang {
   case "en":
      return fmt.Sprintf("Hi!, %s", name), nil
   case "pt":
      return fmt.Sprintf("Oi, %s!", name), nil
   case "es":
      return fmt.Sprintf("¡Hola, %s!", name), nil
   case "fr":
      return fmt.Sprintf("Bonjour, %s!", name), nil
   case "ru":
      return fmt.Sprintf("Привет !!!!, %s!", name), nil
   default:
      return "", errors.New("unknown language")
   }
}

// By returns a bed girl
func ByBy(name string) string {
   return fmt.Sprintf("By fack !!!, %s", name)
}
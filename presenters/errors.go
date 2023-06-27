package presenters

import (
	"log"
)

func PresentErrors(err error) {
	log.Panic(err)
}

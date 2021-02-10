package uuid

import (
	"fmt"
	"strings"
	"github.com/pborman/uuid"
)

func main() {
	PrintRandom()
}

func PrintRandom() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}

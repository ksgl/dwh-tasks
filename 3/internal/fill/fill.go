package fill

import (
	"os"

	"github.com/jackc/fake"
)

func Populate() {
	f, _ := os.Create("./src.txt")
	defer f.Close()

	for i := 0; i < 50; i++ {
		// [DWH] name@domain||FirstName||LastName||Age||PhoneNumber [\DWH]
		f.WriteString("[DWH] " + fake.Word() + "@" + fake.DomainName() + "||" + fake.FirstName() + "||" + fake.LastName() + "||" + fake.DigitsN(2) + "||" + fake.Phone() + " [\\DWH]\n")
	}

	f.Sync()
}

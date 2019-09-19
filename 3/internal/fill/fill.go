package fill

import (
	"os"

	"github.com/jackc/fake"
)

func Populate(count int) {
	f, _ := os.Create("./src.txt")

	for i := 0; i < count; i++ {
		// [DWH] name@domain||FirstName||LastName||Age||PhoneNumber [\DWH]
		f.WriteString("[DWH] " + fake.Word() + "@" + fake.DomainName() + "||" + fake.FirstName() + "||" + fake.LastName() + "||" + fake.DigitsN(2) + "||" + fake.Phone() + " [\\DWH]\n")
	}

	f.Sync()

	f.Close()
}

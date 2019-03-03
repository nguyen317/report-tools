package modules

import "fmt"
import "encoding/json"
import "strings"

type Test struct {
	Name  string
	Array []uint8
}

func (t *Test) MarshalJSON() ([]byte, error) {
	var array string
	if t.Array == nil {
		array = "null"
	} else {
		array = strings.Join(strings.Fields(fmt.Sprintf("%d", t.Array)), ",")
	}
	jsonResult := fmt.Sprintf(`{"Name":%q,"Array":%s}`, t.Name, array)
	return []byte(jsonResult), nil
}

func main() {
	t := &Test{"Go", []uint8{'h', 'e', 'l', 'l', 'o'}}

	m, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", m) // {"Name":"Go","Array":[104,101,108,108,111]}
}

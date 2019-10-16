package spell

import "fmt"

func ExampleParse() {
	s := `bash -c 'for i in 0 1 2 3; do echo $i && sleep 1; done' && echo "foo baa"`
	tokens := Parse(s)
	fmt.Println(tokens)
	fmt.Println(len(tokens))
	// Output:
	// [bash -c 'for i in 0 1 2 3; do echo $i && sleep 1; done' && echo "foo baa"]
	// 6
}

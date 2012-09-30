// Simple substitute for readline. Using it from emacs does not really
// use readline at all anyway.
package readline

import (
	"fmt"
	"bufio"
	"sync"
	"os"
)

var rdr *bufio.Reader
var initrdr = &sync.Once{}

func ReadLine(prompt *string) (*string, error) {
	initrdr.Do(func() {
		rdr = bufio.NewReader(os.Stdin)
	})
	fmt.Print(*prompt)
	line, err := rdr.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return &line, nil
}

func AddHistory(s string) {

}

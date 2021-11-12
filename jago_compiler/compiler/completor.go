package compiler

import (
	"fmt"
)

func complete(code string, cat int) string {
	ret := ""
	switch cat {
	case FULL:
		ret = code
	case MAIN_METHOD:
		ret = fmt.Sprintf(`
		class Main {
			%s
		}
		`, code)
	default: // singleton
		ret = fmt.Sprintf(`
	class Main {
		public static void main(String[] args){		
			%s
		}
	}`, code)
	}
	return ret
}

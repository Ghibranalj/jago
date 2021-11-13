package compiler

import (
	"os"
	"regexp"
	"testing"

	"github.com/Ghibranalj/jago/jago_compiler/utils"
)

type _code struct {
	code string
	cat  int
}

var codes = []_code{
	{
		code: `System.out.println("hello world");`,
		cat:  SINGLETON,
	},
	{
		code: `
		public static void main (String[] args) {
			System.out.println("hello world");
		}`,
		cat: MAIN_METHOD,
	},
	{
		code: `
		class program {

			public static void main (String[] args){
				System.out.println(\"hello world\");
			}
		
		}`,
		cat: FULL,
	},
	{
		code: `
		public static void main (String[] args) {
			System.out.println("hello world");
		}`,
		cat: MAIN_METHOD,
	},
	{
		code: `
		public static void main (String[] args)
		{
			System.out.println("hello world");
		}`,
		cat: MAIN_METHOD,
	},
	{
		code: `
		class program 
		{

			public static void main (String[] args){
				System.out.println("hello world");
			}
		
		}`,
		cat: FULL,
	},
}

func TestCategorize(t *testing.T) {

	for _, code := range codes {
		if code.cat != categorize(code.code) {
			t.Error("wrong category")
		}
	}

}

var _completeOutput = `class program {
	public static void main(String[] args){		
		System.out.println("hello World");
	}
}`
var codeIn = []string{
	`class program {
		public static void main(String[] args){		
			System.out.println("hello World");
		}
	}`,
	`public static void main(String[] args){		
		System.out.println("hello World");
	}`,
	"System.out.println(\"hello World\");",
}

func TestComplete(t *testing.T) {
	// match tabs or new line
	m := regexp.MustCompile(`\t|\n`)

	trim := func(val string) string {
		return m.ReplaceAllString(val, "")
	}
	o := trim(_completeOutput)

	for i, v := range codeIn {
		comp := trim(complete(v, i))

		if o != comp {
			t.Errorf("%s \n should be: %s \n", comp, o)
		}
	}
}

func TestRunJavac(t *testing.T) {
	path := "./" + progName + ".java"
	utils.WriteToFile(path,
		`class program {
		public static void main(String[] argss){		
			System.out.println("hello World");
	}	
	}`)

	err := javac(path)
	os.Remove(path)
	os.Remove("./" + progName + ".class")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCompile(t *testing.T) {
	for _, code := range codes {
		if _, err := Compile(code.code); err != nil {
			t.Errorf("%s", err.Error())
		}
	}
}

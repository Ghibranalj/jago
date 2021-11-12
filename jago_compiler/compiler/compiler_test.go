package compiler

import (
	"regexp"
	"testing"
)

type _code struct {
	code string
	cat  int
}

var codes = []_code{
	{
		code: `System.out.Println("hello world");`,
		cat:  SINGLETON,
	},
	{
		code: `
		public static void main (String[] args) {
			System.out.Println("hello world");
		}`,
		cat: MAIN_METHOD,
	},
	{
		code: `
		class TestingClass {

			public static void main (String[] args){
				System.out.Println("hello world");
			}
		
		}`,
		cat: FULL,
	},
	{
		code: `
		public static void main (String[] args) {
			System.out.Println("hello world");
		}`,
		cat: MAIN_METHOD,
	},
	{
		code: `
		public static void main (String[] args)
		{
			System.out.Println("hello world");
		}`,
		cat: MAIN_METHOD,
	},
	{
		code: `
		class TestingClass 
		{

			public static void main (String[] args){
				System.out.Println("hello world");
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

var _completeOutput = `class Main {
	public static void main(String[] args){		
		System.out.println("hello World");
	}
}`
var codeIn = []string{
	`class Main {
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
			t.Errorf("OWW : %s \n should be: %s \n", comp, o)
		}
	}
}

func TestRunJavac(t *testing.T) {
	ret := javac("../" + progName + ".java")

	t.Log("LOG:", ret)
}

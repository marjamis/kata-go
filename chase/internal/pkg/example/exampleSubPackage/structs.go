package exampleSubPackage

type testStruct struct {
	Public  int
	private int
}

type Drink struct {
	Name string
	Ice  bool
}

func External() testStruct {
	return testStruct{
		Public:  1,
		private: 2,
	}
}

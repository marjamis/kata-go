package advent2019

// Day2 function
func Day2(v ...int) []int {
	position := 0
	for true {
		switch v[position] {
		case 1:
			//Addition
			v[v[position+3]] = v[v[position+1]] + v[v[position+2]]
			position = position + 4
		case 2:
			//Multiplication
			v[v[position+3]] = v[v[position+1]] * v[v[position+2]]
			position = position + 4
		case 99:
			//End of app
			return v
		default:
			return nil
		}
	}

	return nil
}

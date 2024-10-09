package main

// run this file with:
//
// $ go run .

func main() {
	segments := NewIntensitySegments()
	segments.Add(10, 30, 1)
	segments.ToString() // 输出: [[10,1],[30,0]]

	segments.Add(20, 40, 1)
	segments.ToString() // 输出: [[10,1],[20,2],[30,1],[40,0]]

	segments.Add(10, 40, -2)
	segments.ToString() // 输出: [[10,-1],[20,0],[30,-1],[40,0]]

	segments = NewIntensitySegments()
	segments.ToString() // Should be "[]"
	segments.Add(10, 30, 1)
	segments.ToString() // Should be "[[10,1],[30,0]]"
	segments.Add(20, 40, 1)
	segments.ToString() // Should be "[[10,1],[20,2],[30,1],[40,0]]"
	segments.Add(10, 40, -1)
	segments.ToString() // Should be "[[20,1],[30,0]]"
	segments.Add(10, 40, -1)
	segments.ToString() // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"

}

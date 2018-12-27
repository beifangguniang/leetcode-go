package Problem0048

func rotate(matrix [][]int) {
	//
	//	[1,2,3],  [0,0 0,1 0,2]
	//  [4,5,6],  [1,0 1,1 1,2]
	//  [7,8,9]   [2,0 2,1 2,2]
	//

	//	[7,4,1],   [2,0 1,0 0,0]
	//  [8,5,2],   [2,1 1,1 0,1]
	//  [9,6,3]    [2,2 1,2 0,2]
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[j][i] = matrix[j][i]
		}
	}
}

func swap() {

}

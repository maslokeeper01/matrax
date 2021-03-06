package utils

func multiply(m1 Matrix, m2 Matrix) Matrix {
	var numOfRows int = m1.Rows
	var numOfCols int = m2.Columns

	var newMatrix = *NewMatrix(numOfRows, numOfCols)

	for it := 1; it <= numOfRows; it++ {

		for i := 0; i < numOfRows; i++ {
			for j := 0; j < numOfCols; j++ {
				acc := 0
				for t := 0; t < m1.Columns; t++ {
					acc += m1.Elements[i][t] * m2.Elements[t][j]
				}

				newMatrix.Elements[i][j] = acc
			}
		}
	}

	return newMatrix
}

func Multiply(m ...Matrix) Matrix {
	return Reduce(multiply, m...)
}

// highly optimized multiplication of square matrices
func multiplySquare(m1 Matrix, m2 Matrix) Matrix {
	n := m1.Rows

	var newMatrix = *NewMatrix(n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				newMatrix.Elements[i][j] += m1.Elements[i][k] * m2.Elements[k][j]
			}
		}
	}

	return newMatrix
}

func MultiplySquare(m ...Matrix) Matrix {
	return Reduce(multiplySquare, m...)
}

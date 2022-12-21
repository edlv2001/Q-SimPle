package main

import (
	"strconv"
)

type matrix [][]complex128

//type vector [][1]complex128

func multiply(m1 matrix, m2 matrix) matrix {
	if m1 == nil || m2 == nil {
		return nil
	}
	m, p := m1.dimension()
	n, q := m2.dimension()
	if n != p {
		return nil
	}
	result := newMatrix(m, q)
	total := complex128(0)
	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			for k := 0; k < p; k++ {
				total = total + m1[i][k]*m2[k][j]
			}
			result[i][j] = total
			total = 0
		}
	}
	return result
}

// 1a dim: numero de filas (num elem en la vertical)
// 2a dim: numero de columnas (num elem en la horizontal)
func (m matrix) dimension() (int, int) {
	return len(m), len(m[0])
}

func newMatrix(rows int, columns int) matrix {
	mat := make([][]complex128, rows)
	for i := range mat {
		mat[i] = make([]complex128, columns)
	}
	return mat
}

func newMatrixValue(a ...[]complex128) matrix {
	vec := make([][]complex128, len(a))
	for i := range vec {
		vec[i] = a[i]
	}
	return vec
}

func newVectorVert(a ...complex128) matrix {
	vec := make([][]complex128, len(a))
	for i := range vec {
		vec[i] = make([]complex128, 1)
		vec[i][0] = a[i]
	}
	return vec
}

func newVectorVertSize(size int) matrix {
	vec := make([][]complex128, size)
	for i := range vec {
		vec[i] = make([]complex128, 1)
		vec[i][0] = 0
	}
	return vec
}

func newVector1234(size int) matrix {
	vec := make([][]complex128, size)
	j := complex128(0)
	for i := 0; i < size; i++ {
		vec[i] = make([]complex128, 1)
		vec[i][0] = j
		j++
	}
	return vec
}

func newVectorHor(a ...complex128) matrix {
	vec := make([][]complex128, 1)
	vec[0] = make([]complex128, len(a))
	for i := range vec[0] {
		vec[0][i] = a[i]
	}
	return vec
}

/*
func (m matrix) matrixToVector() vector {
	_, p := m.dimension()
	res := make(vector, p, 1)
	for i := 0; i < p; i++ {
		res[i][0] = m[i][0]
	}
	return res
}

func (v vector) vectorToMatrix() matrix {
	res := make(matrix, len(v))
	for i := 0; i < len(v); i++ {
		res[i] = make([]complex128, 1)
		res[i][0] = v[i][0]
	}
	return res
}
*/

func (v matrix) tensorVectorProduct(w matrix) matrix {
	res := newVectorVertSize(len(v) * len(w))
	k := 0
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(w); j++ {
			res[k][0] = v[i][0] * w[j][0]
			k++
		}
	}
	return res
}

func (m matrix) tensorProduct(n matrix) matrix {
	p, q := m.dimension()
	a, b := n.dimension()

	out := make(matrix, 0, p*a)
	for i := 0; i < p; i++ {
		for k := 0; k < a; k++ {
			r := make([]complex128, 0, q*b)
			for j := 0; j < q; j++ {
				for l := 0; l < b; l++ {
					r = append(r, m[i][j]*n[k][l])
				}
			}

			out = append(out, r)
		}
	}
	return out

}

func (m matrix) toString() string {
	res := ""
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			res += strconv.FormatComplex(m[i][j], 'f', -1, 64) + "   "
		}
		res += "\n"
	}
	return res
}

/*func newVector(a ...complex128) vector {
	res := make(vector, len(a))
	for i := 0; i < len(a); i++ {
		res[i][0] = a[i]
	}
	return res
}*/

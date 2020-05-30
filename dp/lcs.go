package dp

import "math"

//LCS : https://www.geeksforgeeks.org/tag/lcs/
func LCS(x []rune, y []rune) [][]int {
	//L(X[0..m-1], Y[0..n-1]
	// if last character of two string match
	// i.e. X[m-1] == Y[m-1] then L(X[0..m-1], Y[0..n-1] = 1 + L(X[0..m-2], Y[0..n-2]
	//else L(X[0..m-1], Y[0..n-1] = max(L(X[0..m-1], Y[0..n-2], L(X[0..m-2], Y[0..n-1])
	//consider X -> “AGGTAB” and Y -> “GXTXAYB”
	// X[5] = B = Y[6] then L(X(0,5),Y(0,6)) = 1 + L(X(0,4),Y(0,5))
	// but X(0,4),Y(0,5) i.e. AGGTA , GXTXAY => L(X(0,5),Y(0,6)) = 1 + max(L(X(0,3),Y(0,5)),L(X(0,4),Y(0,4)))
	//                                 L(X(0,5),Y(0,6))
	//						                   |
	//					              1 + L(X(0,4),   Y(0,5))
	// 						  			 (AGGTA ,   GXTXAY)
	// 							  			 /              \
	// 						Max(	 L(X(0,3),Y(0,5)) ,	L(X(0,4),Y(0,4)) )
	// 					   			(AGGT , GXTXAY)      (AGGTA , GXTXA)
	//                        			 /                    \
	// 		Max(L(X(0,2),Y(0,5)) , L(X(0,3),Y(0,4))      1 + L(X(0,3),Y(0,3)) )
	//             (AGG,GXTXAY)    (AGGT , GXTXA)             (AGGT , GXTX)
	//              /												|
	//  Max(L(X(0,1),Y(0,5)),L(X(0,3),Y(0,3))       <<<<<<<<<<<<<<<<

	// bottom up => lcs(X(0,0),Y(0,0)) = 0
	// lcs(X(0,1),Y(0,1)) = Max(lcs(X(0,0),Y(0,1)),lcs(X(0,1),Y(0,0))) => Max(lcs(A,GX),lcs(AG,G))

	lcs := make([][]int, len(x)+1)
	for row := range lcs {
		lcs[row] = make([]int, len(y)+1)
	}

	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if x[i-1] == y[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = int(math.Max(float64(lcs[i-1][j]), float64(lcs[i][j-1])))
			}
		}
	}
	return lcs
}

//PrintLCS : print the Longets common subsequence
func PrintLCS(x []rune, y []rune) string {
	lcs := LCS(x, y)
	index := lcs[len(x)][len(y)]

	result := make([]rune, index+1)
	i, j := len(x), len(y)
	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			result[index-1] = x[i-1]
			i--
			j--
			index--
		} else if lcs[i-1][j] > lcs[i][j-1] {
			i--
		} else {
			j--
		}

	}
	return string(result)

}

package sudoku_data

const N int = 9


type Board struct {
	data [N][N]int,
	nspace int,
}


func New (data *[N][N]int) Board {
	nspace := int (0)
	for col:=int (0); col<N; col++ {
		for row:=int (0); row<N; row++ {
			if 0>=data[col][row] || data[col][row]>9 {
				nspace ++;
			}
		}
	}

	return Board { data: &data, fill: fill }
}

func (b Board) ToString () string {
	s := ""
	for col:=int (0); col<N; col++ {
		if col%3==0 {
			s += "-------------\n"
		}
		for row:=int (0); row<N; row++ {
			if row%3==0 {
				s += "|"
			}
			if b.data[col][row]!=-1 {
				s += b.data[col][row];
			}
		}
		s += "|\n"
	}
	s += "-------------\n"

	return s
}


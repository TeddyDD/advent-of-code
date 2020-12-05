package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func id(row, column int) int {
	return (row * 8) + column
}

func searchRow(spec string) int {
	return search(0, 127, spec[0:7])
}

func searchColumn(spec string) int {
	return search(0, 7, spec[7:])
}

func search(low, up int, spec string) int {
	s := spec[0 : len(spec)-1]
	for _, l := range s {
		switch l {
		case 'F', 'L':
			up = (low + up) / 2
		case 'B', 'R':
			low = (up+low)/2 + 1
		}
	}
	if spec[len(spec)-1] == 'F' ||
		spec[len(spec)-1] == 'L' {
		return min(low, up)
	}
	return max(low, up)
}

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a:=make([][]uint8,dy)
	for i:=0;i<dy;i++{
		a[i]=make([]uint8,dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j % 10 == 0:
				a[i][j] = uint8(i^j)
			case j % 7 == 0:
				a[i][j] = uint8(i*j)
			case j % 3 == 0:
				a[i][j] = uint8(j^2)
			default:
				a[i][j] = uint8(i^j)
			}
		}
	}	
	return a
}

func main() {
	pic.Show(Pic)
}
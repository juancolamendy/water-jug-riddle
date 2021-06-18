package mathutils

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Gcd(x int,y int) int {
	var gcdnum int
	for i := 1; i <=x && i <=y ; i++ {
		if(x%i==0 && y%i==0) {
			gcdnum=i
		} 
	}
	return gcdnum
}

func IsMultiple(x, y int) bool {
	return (x % y) == 0
}
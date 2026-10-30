package main

func sa() {
	if len(A) >= 2 {
		A[0], A[1] = A[1], A[0]
	}
	noter("sa")
}

func sb() {
	if len(B) >= 2 {
		B[0], B[1] = B[1], B[0]
	}
	noter("sb")
}

func ss() {
	if len(A) >= 2 && len(B) >= 2 {
		sb()
		sa()
	}
	noter("ss")
}


func pb() {
	if len(A) > 0 {
		haut := A[0]
		A = A[1:]
		B = append([]int{haut}, B...)
	}
	noter("pb")
}

func pa() {
	if len(B) > 0 {
		haut := B[0]
		B = B[1:]
		A = append([]int{haut}, A...)
	}
	noter("pa")
}

func ra() {
	if len(A) >= 2 {
		A = append(A[1:], A[0])
	}
	noter("ra")
}

func rb() {
	if len(B) >= 2 {
		B = append(B[1:], B[0])
	}
	noter("rb")
}

func rr() {
	if len(A) >= 2 && len(B) >= 2 {
		rb()
		ra()
	}
	noter("rr")
}

func rra() {
	if len(A) >= 2 {
		fond := A[len(A)-1]
		A = append([]int{fond}, A[:len(A)-1]...)
	}
	noter("rra")
}

func rrb() {
	if len(B) >= 2 {
		fond := B[len(B)-1]
		B = append([]int{fond}, B[:len(B)-1]...)
	}
	noter("rrb")
}

func rrr() {
	if len(A) >= 2 && len(B) >= 2 {
		rrb()
		rra()
	}
	noter("rrr")
}


func noter(geste string) {
	if len(Gestes) >= 100000 {
		panic("plus de 100 000 gestes, une boucle ne s'arrête sans doute jamais")
	}
	Gestes = append(Gestes, geste)
}
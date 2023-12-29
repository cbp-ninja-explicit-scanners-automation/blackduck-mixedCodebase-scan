package data

import "fmt"

type Stack []rune

func (s *Stack) length() int {
	return len(*s)
}

func (s *Stack) push(b rune) {
	*s = append(*s, b)
}

func (s *Stack) pop() rune {
	st := *s
	pop := st[len(st)-1]
	return pop
}

func (s *Stack) remove() []rune {
	st := *s
	st = st[:len(st)-1]
	return st
}

func isLenEven(s string) bool {
	if len(s)%2 == 0 {
		return true
	} else {
		return false
	}
}

func isValid(s string) bool {
	var result bool
	bracket := map[rune]rune{40: 41, 91: 93, 123: 125}
	var st Stack
	for _, val := range s {
		_, ok := bracket[val]
		if ok {
			st.push(val)
		} else {
			// if st.length() == 0 {
			//     continue
			// } else {
			popped := st.pop()
			if bracket[popped] == val {
				st = st.remove()
			} else {
				st.push(val)
			}
			// }
		}
		fmt.Println(st)
	}
	if st.length() == 0 && isLenEven(s) {
		result = true
	} else {
		result = false
	}
	return result
}

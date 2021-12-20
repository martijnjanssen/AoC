package day_18

import (
	"testing"
)

func TestExplode(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}
	for _, tst := range tests {
		ast := createAST(tst.in, 0, nil)
		traverse(ast)
		if getString(ast) != tst.out {
			t.Errorf("\n%s\nWas:\n%s\nShould be:\n%s", tst.in, getString(ast), tst.out)
		}
	}
}

func TestTraverse(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}
	for _, tst := range tests {
		ast := createAST(tst.in, 0, nil)
		for traverse(ast) {
		}
		if getString(ast) != tst.out {
			t.Errorf("\n%s\nWas:\n%s\nShould be:\n%s", tst.in, getString(ast), tst.out)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}, "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}, "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}, "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
	}
	for _, tst := range tests {
		acc := ""
		for _, n := range tst.in {
			if acc == "" {
				acc = n
				continue
			}

			combAST := createAST("["+acc+","+n+"]", 0, nil)
			for traverse(combAST) {
			}
			acc = getString(combAST)
		}
		if acc != tst.out {
			t.Errorf("\n%s\nWas:\n%s\nShould be:\n%s", tst.in, acc, tst.out)
		}
	}
}

func TestAddBig(t *testing.T) {
	in := []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]"}
	out := []string{
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		"[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}

	acc := in[0]
	for i := range out {

		combAST := createAST("["+acc+","+in[i+1]+"]", 0, nil)
		for traverse(combAST) {
		}
		acc = getString(combAST)
		if acc != out[i] {
			t.Errorf("\nWas:\n%s\nShould be:\n%s", acc, out[i])
			t.FailNow()
		}
	}
}

func TestAddBig2(t *testing.T) {
	in := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}
	out := "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"

	acc := in[0]
	for i := 1; i < len(in); i++ {
		combAST := createAST("["+acc+","+in[i]+"]", 0, nil)
		for traverse(combAST) {
		}
		acc = getString(combAST)
	}
	if acc != out {
		t.Errorf("\nWas:\n%s\nShould be:\n%s", acc, out)
		t.FailNow()
	}
	m := magnitude(createAST(acc, 0, nil))
	if m != 4140 {
		t.Errorf("\nMagnitude was: %d\nShould be: %d\n", m, 4140)
		t.FailNow()
	}
}

func TestMagnitude(t *testing.T) {
	in := []string{
		"[[1,2],[[3,4],5]]",
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}
	out := []int{
		143,
		1384,
		445,
		791,
		1137,
		3488,
	}

	for i := 0; i < len(in); i++ {
		m := magnitude(createAST(in[i], 0, nil))
		if m != out[i] {
			t.Errorf("\n%s\nMagnitude was: %d\nShould be: %d\n", in[i], m, out[i])
			t.FailNow()
		}
	}

}
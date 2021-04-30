package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHalt(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{})
	var want uint64 = 1
	got := g.P
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64 {
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	var want uint64 = 2
	got := g.P
	if want != got {
		t.Errorf("want: %d, got %d", want, got)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{gmachine.OpINCA})
	var want uint64 = 1
	got := g.A
	if want != got {
		t.Errorf("want %d, got %d", want ,got)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{gmachine.OpSETA, 2, gmachine.OpDECA})
	var want uint64 = 1
	got := g.A
	if want != got {
		t.Errorf("want %d, got %d", want ,got)
	}
}

func TestProgramSubtract2From3(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{gmachine.OpSETA, 3, gmachine.OpDECA, gmachine.OpDECA})
	var want uint64 = 1
	got := g.A
	if want != got {
		t.Errorf("want %d, got %d", want ,got)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64 {gmachine.OpSETA, 5})
	var want uint64 = 5
	got := g.A
	if want != got {
		t.Errorf("want %d, got %d", want ,got)
	}
}

func TestSubtract2(t *testing.T) {
	t.Parallel()
	g := gmachine.New()

	g.RunProgram([]uint64{gmachine.OpSETA, 2, gmachine.OpDECA, gmachine.OpDECA})
	var want2 uint64 = 0
	got2 := g.A
	if want2 != got2 {
		t.Errorf("want: %d got: %d", want2, got2)
	}

	g.P = 0
	g.RunProgram([]uint64{gmachine.OpSETA, 10, gmachine.OpDECA, gmachine.OpDECA})
	var want10 uint64 = 8
	got10 := g.A
	if want10 != got10 {
		t.Errorf("want: %d got: %d", want10, got10)
	}

	g.P = 0
	g.RunProgram([]uint64{gmachine.OpSETA, 5, gmachine.OpDECA, gmachine.OpDECA})
	var want5 uint64 = 3
	got5 := g.A
	if want5 != got5 {
		t.Errorf("want: %d got: %d", want5, got5)
	}


}
package bf

import "testing"

func TestHelloWorld(t *testing.T) {
	code := `++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`
	expected := "Hello World!\n"

	interpreter := New(code)
	output, err := interpreter.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}

func TestLoop(t *testing.T) {
	code := "++++[>++++<-]>."
	expected := string([]byte{16})

	interpreter := New(code)
	output, err := interpreter.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}

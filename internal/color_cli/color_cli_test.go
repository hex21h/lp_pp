package color_cli

import (
	"testing"
)

func TestResetColor(t *testing.T) {
	expectedOutput := "\033[0m"
	output := ResetColor()

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_red(t *testing.T) {
	input := "red"
	expectedOutput := "\033[31m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_green(t *testing.T) {
	input := "green"
	expectedOutput := "\033[32m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_yellow(t *testing.T) {
	input := "yellow"
	expectedOutput := "\033[33m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_blue(t *testing.T) {
	input := "blue"
	expectedOutput := "\033[34m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_purple(t *testing.T) {
	input := "purple"
	expectedOutput := "\033[35m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_cyan(t *testing.T) {
	input := "cyan"
	expectedOutput := "\033[36m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_white(t *testing.T) {
	input := "white"
	expectedOutput := "\033[37m"
	output, _ := GetColor(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %s", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestGetColor_err(t *testing.T) {
	input := "asd"
	_, err := GetColor(input)

	if err == nil {
		t.Errorf("Failed ! err got nil ")
	} else {
		t.Logf("Success !")
	}
}

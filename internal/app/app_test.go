package app

import (
	"bufio"
	"errors"
	"github.com/AFHH999/ToDo/internal/models"
	"strings"
	"testing"
)

func TestGetInput(t *testing.T) {
	input := "Hello, World!\n"
	reader := bufio.NewReader(strings.NewReader(input))
	result := GetInput("Prompt: ", reader)
	if result != "Hello, World!" {
		t.Errorf("Expected: 'Hello, World!', got '%s'", result)
	}
}

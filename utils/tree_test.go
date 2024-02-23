package utils_test

import (
	"os"
	"testing"

	"github.com/xsymphony/telegram-gemini-bot/utils"
)

func TestPrintTree(t *testing.T) {
	dir, _ := os.Getwd()
	utils.PrintTree(dir, "")
}

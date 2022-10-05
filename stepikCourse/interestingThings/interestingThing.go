package main

import (
	"os"
	"github.com/you/proj/cli"
)

/*
Такой подход связан с тем что из package main нельзя ничего импортировать
Желательно использовать, использовать минимум кода в этом паркетете
*/
func main() {
	os.Exit(cli.Run())
}

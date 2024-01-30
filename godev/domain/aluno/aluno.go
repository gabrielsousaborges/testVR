package aluno

import (
	"fmt"
	"regexp"
	"strings"
)

func NormalizaNome(nome string) []string {
	re := regexp.MustCompile("[[:punct:]]")
	nome = re.ReplaceAllString(nome, "")

	nomes := strings.Split(nome, " ")

	for i, nome := range nomes {
		nomes[i] = fmt.Sprintf("%%%s%%", strings.ToLower(nome))
	}
	
	return nomes
}
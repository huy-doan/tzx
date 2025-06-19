package bankaccount

import (
	"regexp"

	"github.com/test-tzs/nomraeite/internal/pkg/utils"
	"github.com/test-tzs/nomraeite/internal/pkg/utils/static"
)

type AccountHolderKana string

const (
	AozoraAccountNameRegex = `^[ｦ-ﾟa-zA-Z0-9()-./, ]+$`
)

func (ahk AccountHolderKana) Value() string {
	return string(ahk)
}

func (ahk AccountHolderKana) IsValid() bool {
	return regexp.MustCompile(AozoraAccountNameRegex).MatchString(utils.ConvertKana(string(ahk), []string{
		static.CONV_HALFWIDTHKATA_FROM_FULLWIDTHKATA,
	}))
}

func FromStringToAccountHolderKana(s string) AccountHolderKana {
	return AccountHolderKana(s)
}

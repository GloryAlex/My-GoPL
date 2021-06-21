package math2

import (
	"fmt"
	"strings"
)

func Max(values ...int) (ret int, err error) {
	if len(values) <= 0 {
		return 0, fmt.Errorf("max: args must more than one")
	}
	ret = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > ret {
			ret = values[i]
		}
	}
	return ret, nil
}

func Join(strs ...string) string {
	ret := strings.Builder{}
	for _, s := range strs {
		ret.WriteString(s)
	}
	return ret.String()
}

package percentage

import (
	"fmt"
	"math"
)

func Display(value float64) string {
	per := value * 100
	per = math.Round(per*100) / 100

	return fmt.Sprintf("%.1f", per) + "%"
}

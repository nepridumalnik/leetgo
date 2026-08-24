package findnuniqueintegerssumuptozero_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/find_n_unique_integers_sum_up_to_zero"

	"github.com/stretchr/testify/require"
)

func Test_FindNUniqueIntegersSumUpToZero(t *testing.T) {
	for i := 4; i < 1000; i++ {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			result := findnuniqueintegerssumuptozero.SumZero(i)
			sum := 0

			for _, n := range result {
				sum += n
			}

			require.Zero(t, sum)
		})
	}

}

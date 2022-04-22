package myAtoi

import (
	"math"
	"strconv"
	"testing"
)

func TestAToi(t *testing.T) {
	type T struct {
		expectNum    int64 // expected number
		expectFailed bool  // is expected fail ?
	}

	inputs := map[string]T{
		"123":                  {123, false},
		"+123":                 {123, false},
		"-123":                 {-123, false},
		" -123":                {-123, false},
		" -123 ":               {-123, false},
		"0":                    {0, false},
		"-0":                   {0, false},
		"-1":                   {-1, false},
		"9223372036854775807":  {math.MaxInt64, false},
		"-9223372036854775808": {math.MinInt64, false},

		"9223372036854775808":     {0, true},
		"92233720368547758081111": {0, true},
		"-9223372036854775809":    {0, true},
		"":                        {0, true},
		"   ":                     {0, true},
	}

	for str, expected := range inputs {
		n, err := AToi(str)

		// 预期外的错误
		if (err != nil) != expected.expectFailed {
			t.Fatal(err)
		}

		// 返回预期意外的
		if n != expected.expectNum {
			t.Fatalf("expected: %d, get: %d", expected.expectNum, n)
		}
	}
}

func BenchmarkAToiV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AToi(strconv.FormatInt(int64(i), 10))
	}
}

func BenchmarkStandardAToi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Atoi(strconv.FormatInt(int64(i), 10))
	}
}

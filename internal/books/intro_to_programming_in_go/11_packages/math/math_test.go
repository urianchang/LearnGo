package math

import "testing"

type testpair struct {
    values []float64
    average float64
    min float64
    max float64
}

var tests = []testpair{
    { []float64{1,2}, 1.5, 1, 2 },
    { []float64{1,1,1,1,1,1}, 1, 1, 1 },
    { []float64{-1,1}, 0, -1, 1 },
}

func TestMathFuncs(t *testing.T) {
    for _, pair := range tests {
        avg := Average(pair.values)
        if avg != pair.average {
            t.Error("For", pair.values, "expected", pair.average, "got", avg,)
        }
        min := Min(pair.values)
        if min != pair.min {
            t.Error("For", pair.values, "expected", pair.min, "got", min,)
        }
        max := Max(pair.values)
        if max != pair.max {
            t.Error("For", pair.values, "expected", pair.max, "got", max,)
        }
    }
}


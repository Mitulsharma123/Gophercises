package calci

import "testing"

type minTestPair struct{
  values []float64
  min float64
}

type maxTestPair struct{
  values []float64
  max float64
}

var minTests = []minTestPair {
  { []float64{1,5},1},
  { []float64{6,2,5},2},
  { []float64{9,22,3},3},
}

var maxTests = []maxTestPair {
  { []float64{1,5},5},
  { []float64{6,2,5},6},
  { []float64{9,22,3},22},
}

func TestMax(t *testing.T) {
  for _, pair := range maxTests{
	  v := Max(pair.values)
	  if v != pair.max{
		t.Error(
			"For", pair.values,
			"expected", pair.max,
			"got", v,
		)
		}
	}
}

func TestMin(t *testing.T) {
  for _, pair := range minTests{
          v := Min(pair.values)
          if v != pair.min{
                t.Error(
                        "For", pair.values,
                        "expected", pair.min,
                        "got", v,
                )
                }
        }
}

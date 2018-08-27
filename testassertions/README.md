Using the following code based on [this example](http://goinbigdata.com/testing-go-code-with-testify):

```
func TestAssertSubtract(t *testing.T) {
    ds := []struct {
        name           string
        min, sub, diff int
    }{
        {"zeroes", 0, 0, 0},
        {"tens", 10, 10, 0},
        {"larger min", 10, 1, 9},
        {"larger sub", 1, 10, -9},
        {"fail", 10, 10, -999},
    }

    calc := NewCalculator(nil)

    for _, d := range ds {
        t.Run(d.name, func(t *testing.T) {
            assert.Equal(t, d.diff, calc.Subtract(d.min, d.sub))
        })
    }
}

func TestStandardSubtract(t *testing.T) {
    ds := []struct {
        name           string
        min, sub, want int
    }{
        {"zeroes", 0, 0, 0},
        {"tens", 10, 10, 0},
        {"larger min", 10, 1, 9},
        {"larger sub", 1, 10, -9},
        {"fail", 10, 10, -999},
    }

    calc := NewCalculator(nil)

    for _, d := range ds {
        t.Run(d.name, func(t *testing.T) {
            got := calc.Subtract(d.min, d.sub)
            if got != d.want {
                t.Errorf("got %v, want %v", got, d.want)
            }
        })
    }
}
```

It's output is ugly and needlessly verbose:

```
=== RUN   TestAssertSubtract
=== RUN   TestAssertSubtract/zeroes
=== RUN   TestAssertSubtract/tens
=== RUN   TestAssertSubtract/larger_min
=== RUN   TestAssertSubtract/larger_sub
=== RUN   TestAssertSubtract/fail
--- FAIL: TestAssertSubtract (0.00s)
    --- PASS: TestAssertSubtract/zeroes (0.00s)
    --- PASS: TestAssertSubtract/tens (0.00s)
    --- PASS: TestAssertSubtract/larger_min (0.00s)
    --- PASS: TestAssertSubtract/larger_sub (0.00s)
    --- FAIL: TestAssertSubtract/fail (0.00s)
        calculator_test.go:25: 
                        Error Trace:    calculator_test.go:25
                        Error:          Not equal: 
                                        expected: -999
                                        actual  : 0
                        Test:           TestAssertSubtract/fail
=== RUN   TestStandardSubtract
=== RUN   TestStandardSubtract/zeroes
=== RUN   TestStandardSubtract/tens
=== RUN   TestStandardSubtract/larger_min
=== RUN   TestStandardSubtract/larger_sub
=== RUN   TestStandardSubtract/fail
--- FAIL: TestStandardSubtract (0.00s)
    --- PASS: TestStandardSubtract/zeroes (0.00s)
    --- PASS: TestStandardSubtract/tens (0.00s)
    --- PASS: TestStandardSubtract/larger_min (0.00s)
    --- PASS: TestStandardSubtract/larger_sub (0.00s)
    --- FAIL: TestStandardSubtract/fail (0.00s)
        calculator_test.go:48: got 0, want -999
```

Dropping the subtest decreases assert's friction with the standard library, but then context is lost.

It's usage, to attain it's purpose of dropping a few lines of code, is also ugly and difficult to parse with a glance:

```
assert.Equal(t, d.diff, calc.Subtract(d.min, d.sub))

got := calc.Subtract(d.min, d.sub)
if got != d.want {
    t.Errorf("got %v, want %v", got, d.want)
}
```

More so, in my opinion (and if for only this), reducing a handful of lines of code at the cost of having to learn a proprietary set of equivalency rules makes using an assertion library astoundingly unappealing.

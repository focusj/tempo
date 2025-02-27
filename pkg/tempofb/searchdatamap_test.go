package tempofb

import (
	"fmt"
	"testing"
)

func BenchmarkSearchDataMapAdd(b *testing.B) {
	intfs := []struct {
		name string
		f    func() SearchDataMap
	}{
		{"SearchDataMapSmall", func() SearchDataMap { return make(SearchDataMapSmall, 10) }},
		{"SearchDataMapLarge", func() SearchDataMap { return make(SearchDataMapLarge, 10) }},
	}

	testCases := []struct {
		name    string
		values  int
		repeats int
	}{
		{"inserts", 1, 0},
		{"inserts", 10, 0},
		{"inserts", 100, 0},
		{"repeats", 10, 100},
		{"repeats", 10, 1000},
		{"repeats", 100, 100},
		{"repeats", 100, 1000},
	}

	for _, tc := range testCases {
		for _, intf := range intfs {
			b.Run(fmt.Sprint(tc.name, "/", tc.values, "x value/", tc.repeats, "x repeat", "/", intf.name), func(b *testing.B) {

				var k []string
				for i := 0; i < b.N; i++ {
					k = append(k, fmt.Sprintf("key%d", i))
				}

				var v []string
				for i := 0; i < tc.values; i++ {
					v = append(v, fmt.Sprintf("value%d", i))
				}

				s := intf.f()
				insert := func() {
					for i := 0; i < len(k); i++ {
						for j := 0; j < len(v); j++ {
							s.Add(k[i], v[j])
						}
					}
				}

				// insert
				b.ResetTimer()
				insert()

				// reinsert?
				if tc.repeats > 0 {
					b.ResetTimer()
					for i := 0; i < tc.repeats; i++ {
						insert()
					}
				}
			})
		}
	}

}

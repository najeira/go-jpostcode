package jpostcode

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_convertJSONToAddress(t *testing.T) {
	tests := map[string]struct {
		input   interface{}
		want    *Address
		wantErr bool
	}{
		"ok with valid input": {
			input: map[string]interface{}{
				"postcode":         "1638001",
				"prefecture":       "東京都",
				"prefecture_kana":  "トウキョウト",
				"prefecture_code":  13,
				"city":             "新宿区",
				"city_kana":        "シンジュクク",
				"town":             "西新宿",
				"town_kana":        "ニシシンジュク",
				"street":           "２丁目８−１",
				"office_name":      "東京都庁",
				"office_name_kana": "トウキヨウトチヨウ",
			},
			want: &Address{
				PostCode:       "1638001",
				Prefecture:     "東京都",
				PrefectureCode: 13,
				PrefectureKana: "トウキョウト",
				City:           "新宿区",
				CityKana:       "シンジュクク",
				Town:           "西新宿",
				TownKana:       "ニシシンジュク",
				Street:         "２丁目８−１",
				OfficeName:     "東京都庁",
				OfficeNameKana: "トウキヨウトチヨウ",
			},
			wantErr: false,
		},
		"ng with invalid input": {
			input: map[string]interface{}{
				"post_code":       999999,
				"prefecture":      999999,
				"prefecture_kana": 999999,
				"prefecture_code": "abcde",
				"city":            []int{1, 2, 3},
				"city_kana":       []float64{1, 2, 3},
			},
			wantErr: true,
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			got, err := convertJSONToAddress(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertJSONToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if d := cmp.Diff(got, tt.want); d != "" {
				t.Error(d)
			}
		})
	}
}

func Benchmark_newBadgerAdapter(b *testing.B) {
	ms := startMemStats()
	if _, err := newBadgerAdapter(); err != nil {
		b.Error(err)
	}
	ms.Stop("newBadgerAdapter")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := newBadgerAdapter()
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_newMapAdapter(b *testing.B) {
	ms := startMemStats()
	if _, err := newMapAdapter(); err != nil {
		b.Error(err)
	}
	ms.Stop("newMapAdapter")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := newMapAdapter()
		if err != nil {
			b.Error(err)
		}
	}
}

type memStats struct {
	s runtime.MemStats
}

func startMemStats() *memStats {
	s := new(memStats)
	runtime.GC()
	runtime.ReadMemStats(&s.s)
	return s
}

func (s *memStats) Stop(name string) {
	var rs runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&rs)

	netHeap := rs.HeapAlloc - s.s.HeapAlloc
	netAllocs := rs.Mallocs - s.s.Mallocs
	netBytes := rs.TotalAlloc - s.s.TotalAlloc
	fmt.Printf(
		"%s: %d B/op %d allocs/op %d B/inc\n",
		name, netBytes, netAllocs, netHeap)
	s.s = rs
}

package test

import (
	"sort"
	"testing"
)

func TestOrderedSlice(t *testing.T) {
	// 정수형 슬라이스 테스트
	intSlice := OrderedSlice[int]{5, 2, 8, 1, 9}
	sort.Sort(intSlice)
	expected := OrderedSlice[int]{1, 2, 5, 8, 9}
	for i := range intSlice {
		if intSlice[i] != expected[i] {
			t.Errorf("정수 정렬 실패: 인덱스 %d에서 %d를 기대했으나 %d를 얻음", i, expected[i], intSlice[i])
		}
	}

	// 문자열 슬라이스 테스트
	strSlice := OrderedSlice[string]{"banana", "apple", "cherry"}
	sort.Sort(strSlice)
	expectedStr := OrderedSlice[string]{"apple", "banana", "cherry"}
	for i := range strSlice {
		if strSlice[i] != expectedStr[i] {
			t.Errorf("문자열 정렬 실패: 인덱스 %d에서 %s를 기대했으나 %s를 얻음", i, expectedStr[i], strSlice[i])
		}
	}

	// 실수형 슬라이스 테스트
	floatSlice := OrderedSlice[float64]{3.14, 1.41, 2.71}
	sort.Sort(floatSlice)
	expectedFloat := OrderedSlice[float64]{1.41, 2.71, 3.14}
	for i := range floatSlice {
		if floatSlice[i] != expectedFloat[i] {
			t.Errorf("실수 정렬 실패: 인덱스 %d에서 %f를 기대했으나 %f를 얻음", i, expectedFloat[i], floatSlice[i])
		}
	}
}

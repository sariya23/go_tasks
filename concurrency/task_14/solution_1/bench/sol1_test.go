// solution_test - пакет с бенчмарками.
//
// Exec - go test -race -bench=. -benchmem . -count=20 | tee stats.txt
//
// Check stat - benchstat stats.txt
package solution_test

import (
	"testing"

	solution "github.com/sariya23/go_tasks/concurrency/task_14/solution_1/bench"
)

var global int64

func BenchmarkInitFunc(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = solution.InitFunc()
	}
	global = res
}

func BenchmarkAtomicFunc(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = solution.AtomicFunc()
	}
	global = res
}

func BenchmarkSemaFunc(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = solution.SemaFunc()
	}
	global = res
}

func BenchmarkCriticalSchemaFunc(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = solution.CriticalSectionFunc()
	}
	global = res
}

package stately_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/StatelyCloud/go-sdk/stately"
)

func TestFilterItemList(t *testing.T) {
	testItems := []stately.Item{
		&BarItem{Bar: "1"},
		&FooItem{},
		&BazItem{},
		&BarItem{Bar: "2"},
		&BazItem{Baz: "1"},
		&BarItem{Bar: "3"},
	}

	tests := []struct {
		name string
		run  func(t testing.TB, items ...stately.Item)
	}{
		{
			name: "filter *testschema.Employee",
			run: func(t testing.TB, items ...stately.Item) {
				givenItems := stately.ReduceTo[[]*BarItem](items...)
				assert.Equalf(t, []*BarItem{
					{Bar: "1"},
					{Bar: "2"},
					{Bar: "3"},
				}, givenItems, "ReduceTo(*testschema.Employee)")
			},
		},
		{
			name: "filter *testschema.VersionedFoo which doesn't exist",
			run: func(t testing.TB, items ...stately.Item) {
				givenItems := stately.ReduceTo[[]*BuzItem](items...)
				assert.Nil(t, givenItems, "ReduceTo(*testschema.VersionedFoo)")
			},
		},
		{
			name: "filter *testschema.NativeBarz",
			run: func(t testing.TB, items ...stately.Item) {
				givenItems := stately.ReduceTo[[]*BazItem](items...)
				assert.Equalf(t, []*BazItem{
					{},
					{Baz: "1"},
				}, givenItems, "ReduceTo(*testschema.NativeBarz)")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.run(t, testItems...)
		})
	}
}

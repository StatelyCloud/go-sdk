package stately_test

import (
	"github.com/StatelyCloud/go-sdk/pb/db"
)

type FooItem struct {
	Foo string
}

func (f *FooItem) StatelyItemType() string {
	return "FooItem"
}

func (f *FooItem) UnmarshalStately(_ *db.Item) error {
	return nil
}

func (f *FooItem) MarshalStately() (*db.Item, error) {
	return nil, nil
}

func (f *FooItem) KeyPath() string {
	return ""
}

type BarItem struct {
	Bar string
}

func (f *BarItem) StatelyItemType() string {
	return "BarItem"
}

func (f *BarItem) UnmarshalStately(_ *db.Item) error {
	return nil
}

func (f *BarItem) MarshalStately() (*db.Item, error) {
	return nil, nil
}

func (f *BarItem) KeyPath() string {
	return ""
}

type BazItem struct {
	Baz string
}

func (f *BazItem) StatelyItemType() string {
	return "BazItem"
}

func (f *BazItem) UnmarshalStately(_ *db.Item) error {
	return nil
}

func (f *BazItem) MarshalStately() (*db.Item, error) {
	return nil, nil
}

func (f *BazItem) KeyPath() string {
	return ""
}

type BuzItem struct {
	Buz string
}

func (f *BuzItem) StatelyItemType() string {
	return "BuzItem"
}

func (f *BuzItem) UnmarshalStately(_ *db.Item) error {
	return nil
}

func (f *BuzItem) MarshalStately() (*db.Item, error) {
	return nil, nil
}

func (f *BuzItem) KeyPath() string {
	return ""
}

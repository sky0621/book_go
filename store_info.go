package main

// アルファベット大文字で始まると定数扱い
const DefaultStoreName = "store.json"

func NewStoreInfo(storeName string) StoreInfo {
	if storeName == "" {
		return &storeInfo{storeName: DefaultStoreName}
	}
	return &storeInfo{storeName: storeName}
}

type StoreInfo interface {
	GetName() string
}

type storeInfo struct {
	storeName string
}

func (s *storeInfo) GetName() string {
	return s.storeName
}

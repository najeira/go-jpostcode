package jpostcode

import (
	"sync"
)

var (
	UseMapAdapter bool

	bAdapter     Adapter
	bAdapterOnce sync.Once

	mAdapter     Adapter
	mAdapterOnce sync.Once
)

func Find(postCode string) (*Address, error) {
	addresses, err := Search(postCode)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, ErrNotFound
	}
	return addresses[0], nil
}

func Search(postCode string) ([]*Address, error) {
	if UseMapAdapter {
		return getMapAdapter().SearchAddressesFromPostCode(postCode)
	}
	return getBadgerAdapter().SearchAddressesFromPostCode(postCode)
}

func getBadgerAdapter() Adapter {
	bAdapterOnce.Do(func() {
		// set default adapter
		var err error
		bAdapter, err = newBadgerAdapter()
		if err != nil {
			panic(err)
		}
		// closing DB is not needed because badger adapter is using in-memory DB
		// adapter.db.Close()
	})
	return bAdapter
}

func getMapAdapter() Adapter {
	mAdapterOnce.Do(func() {
		// set default adapter
		var err error
		mAdapter, err = newMapAdapter()
		if err != nil {
			panic(err)
		}
	})
	return mAdapter
}

package util

import (
	"github.com/libp2p/go-libp2p-core/peer"
	maddr "github.com/multiformats/go-multiaddr"
)

func StringsToAddrs(addrStrings ...string) (maddrs []maddr.Multiaddr, err error) {
	for _, addrString := range addrStrings {
		addr, err := maddr.NewMultiaddr(addrString)
		if err != nil {
			return maddrs, err
		}
		maddrs = append(maddrs, addr)
	}
	return
}

func StringsToPeerInfos(addrStrings ...string) ([]peer.AddrInfo, error) {
	maddrs, err := StringsToAddrs(addrStrings...)
	if err != nil {
		return nil, err
	}
	return peer.AddrInfosFromP2pAddrs(maddrs...)
}

package nft_service

import (
	uuid "github.com/google/uuid"
)

const defaultUri string = ""

type Nft struct {
	index string
	owner string
	uri   string
}

func Mint(owner_ string) *Nft {
	nft := &Nft{
		index: uuid.New().String(),
		owner: owner_,
		uri:   defaultUri,
	}
	nft.uri += "/" + nft.index
	return nft
}

func (nft *Nft) SetUri(uri_ string) {
	nft.uri = defaultUri + "/" + uri_
}

func (nft *Nft) Transfer(newOwner_ string) {
	nft.owner = newOwner_
}

func (nft *Nft) GetIndex() string {
	return nft.index
}

func (nft *Nft) GetOwner() string {
	return nft.owner
}

func (nft *Nft) GetUri() string {
	return nft.uri
}

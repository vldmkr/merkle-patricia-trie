package mpt

type (
	PersistNodeBase struct {
		Full  *PersistNodeFull  `cbor:"10,keyasint,omitempty"`
		Short *PersistNodeShort `cbor:"11,keyasint,omitempty"`
		Value *PersistNodeValue `cbor:"12,keyasint,omitempty"`
	}
	PersistNodeFull struct {
		_        struct{} `cbor:",toarray"`
		Children [][]byte
	}
	PersistNodeShort struct {
		_     struct{} `cbor:",toarray"`
		Key   []byte
		Value []byte
	}
	PersistNodeValue []byte

	PersistTrie struct {
		_     struct{} `cbor:",toarray"`
		Pairs []*PersistTriePair
	}
	PersistTriePair struct {
		_     struct{} `cbor:",toarray"`
		Key   []byte
		Value []byte
	}
)

package p2p

import (
	pubsubpb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/n42blockchain/N42/common/hash"
	"github.com/n42blockchain/N42/common/types"
)

// MsgID is a content addressable ID function.
// `SHA256(message.data)[:20]`.
func MsgID(genesisHash types.Hash, pmsg *pubsubpb.Message) string {
	h := hash.Hash(pmsg.Data)
	return string(h[:20])
}

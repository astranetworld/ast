package p2p

import (
	"github.com/astranetworld/ast/common/hash"
	"github.com/astranetworld/ast/common/types"
	pubsubpb "github.com/libp2p/go-libp2p-pubsub/pb"
)

// MsgID is a content addressable ID function.
// `SHA256(message.data)[:20]`.
func MsgID(genesisHash types.Hash, pmsg *pubsubpb.Message) string {
	h := hash.Hash(pmsg.Data)
	return string(h[:20])
}

package sender

import (
	cutils "db-service/common/utils"
	"db-service/g"

	rings "github.com/toolkits/consistent/rings"
)

//KafkaNodeRing consistent hash ring used to manage server nodes
var KafkaNodeRing *rings.ConsistentHashNodeRing

func initNodeRings() {
	cfg := g.Config()

	KafkaNodeRing = rings.NewConsistentHashNodesRing(int32(cfg.Kafka.Replicas), cutils.KeysOfMap(cfg.Kafka.Cluster))
}

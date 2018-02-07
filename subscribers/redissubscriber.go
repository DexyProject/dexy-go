package subscribers

import (
	"gopkg.in/redis.v2"
)

type RedisSubscriber struct {
	pubsub  *redis.PubSub
	channel string
}

func (rs *RedisSubscriber) Subscribe() error {
	err := rs.pubsub.Subscribe(rs.channel)
	if err != nil {
		return err
	}

	return nil
}

func (rs *RedisSubscriber) Listen() (string, error) {
	panic("implement me")
}

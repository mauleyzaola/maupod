package queues

import (
	"gopkg.in/redis.v4"
)

type QueueManager struct {
	client *redis.Client
}

func NewManager(client *redis.Client) *QueueManager {
	return &QueueManager{
		client: client,
	}
}

func (m *QueueManager) LPush(key string, data []byte) error {
	return m.client.LPush(key, data).Err()
}

func (m *QueueManager) RPush(key string, data []byte) error {
	return m.client.RPush(key, data).Err()
}

func (m *QueueManager) LPop(key string) ([]byte, error) {
	return m.client.LPop(key).Bytes()
}

func (m *QueueManager) RPop(key string) ([]byte, error) {
	return m.client.RPop(key).Bytes()
}

func (m *QueueManager) List(key string) ([][]byte, error) {
	var result [][]byte
	cmd := m.client.LRange(key, 0, -1)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	for _, v := range cmd.Val() {
		result = append(result, []byte(v))
	}

	return result, nil
}

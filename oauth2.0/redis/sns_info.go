package redis

import (
	"SNSLoginWithGolang/oauth2.0/domain"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

const SnsKey = "sns.info"

type SnsInfo struct {
	domain.SnsInfo
}

func (s *SnsInfo) Read(ctx context.Context, rds *redis.Client, snsID string) error {
	result, err := rds.HGet(ctx, SnsKey, snsID).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(result), s)
}

func (s *SnsInfo) ReadAll() {

}

func (s *SnsInfo) Create(ctx context.Context, rds *redis.Client) error {
	return rds.HSet(ctx, SnsKey, s.SnsID, s).Err()
}

func (s *SnsInfo) Delete(ctx context.Context, rds *redis.Client, snsID string) error {
	_, err := rds.HDel(ctx, SnsKey, snsID).Result()
	return err
}

func (s *SnsInfo) Update() {

}

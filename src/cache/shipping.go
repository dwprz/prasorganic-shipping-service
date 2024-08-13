package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type ShippingImpl struct {
	redis *redis.ClusterClient
}

func NewShipping(r *redis.ClusterClient) cache.Shipping {
	return &ShippingImpl{
		redis: r,
	}
}

func (s *ShippingImpl) CacheProvinces(ctx context.Context, p *dto.ShipperRes[[]*entity.Province]) {
	b, err := json.Marshal(p)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheProvinces", "section": "json.Marshal"}).Error(err)
		return
	}

	key := "provinces"
	if _, err := s.redis.SetEx(ctx, key, string(b), 24*time.Hour).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheProvinces", "section": "redis.SetEx"}).Error(err)
	}
}

func (s *ShippingImpl) CacheCitiesByProvinceId(ctx context.Context, provinceId int, p *dto.ShipperRes[[]*entity.City]) {
	b, err := json.Marshal(p)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheProvinces", "section": "json.Marshal"}).Error(err)
		return
	}

	key := fmt.Sprintf("province_id:%d:cities", provinceId)
	if _, err := s.redis.SetEx(ctx, key, string(b), 24*time.Hour).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheCitiesByProvinceId", "section": "redis.SetEx"}).Error(err)
	}
}

func (s *ShippingImpl) FindProvinces(ctx context.Context) *dto.ShipperRes[[]*entity.Province] {
	res, err := s.redis.Get(ctx, "provinces").Result()
	if err != nil {
		if err != redis.Nil {
			log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindProvinces", "section": "redis.Get"}).Error(err)
		}

		return nil
	}

	provinces := new(dto.ShipperRes[[]*entity.Province])
	if err := json.Unmarshal([]byte(res), provinces); err != nil {

		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindProvinces", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return provinces
}

func (s *ShippingImpl) FindCitiesByProvinceId(ctx context.Context, provinceId int) *dto.ShipperRes[[]*entity.City] {
	key := fmt.Sprintf("province_id:%d:cities", provinceId)

	res, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindCitiesByProvinceId", "section": "redis.Get"}).Error(err)
		}

		return nil
	}

	provinces := new(dto.ShipperRes[[]*entity.City])
	if err := json.Unmarshal([]byte(res), provinces); err != nil {

		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindCitiesByProvinceId", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return provinces
}

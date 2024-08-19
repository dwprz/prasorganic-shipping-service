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

func (s *ShippingImpl) CacheTrackingByShippingId(ctx context.Context, shippingId string, t *entity.Tracking) {
	b, err := json.Marshal(t)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheTrackingByShippingId", "section": "json.Marshal"}).Error(err)
		return
	}

	key := fmt.Sprintf("shipping_id:%s:trackings", shippingId)

	if _, err := s.redis.SetEx(ctx, key, string(b), 2*time.Hour).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheTrackingByShippingId", "section": "redis.SetEx"}).Error(err)
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
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheCitiesByProvinceId", "section": "json.Marshal"}).Error(err)
		return
	}

	key := fmt.Sprintf("province_id:%d:cities", provinceId)
	if _, err := s.redis.SetEx(ctx, key, string(b), 24*time.Hour).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheCitiesByProvinceId", "section": "redis.SetEx"}).Error(err)
	}
}

func (s *ShippingImpl) CacheSuburbsByCityId(ctx context.Context, cityId int, p *dto.ShipperRes[[]*entity.Suburb]) {
	b, err := json.Marshal(p)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheSuburbsByCityId", "section": "json.Marshal"}).Error(err)
		return
	}

	key := fmt.Sprintf("city_id:%d:suburbs", cityId)
	if _, err := s.redis.SetEx(ctx, key, string(b), 24*time.Hour).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheSuburbsByCityId", "section": "redis.SetEx"}).Error(err)
	}
}

func (s *ShippingImpl) CacheAreasBySuburbId(ctx context.Context, suburbId int, p *dto.ShipperRes[[]*entity.Area]) {
	b, err := json.Marshal(p)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheAreasBySuburbId", "section": "json.Marshal"}).Error(err)
		return
	}

	key := fmt.Sprintf("suburb_id:%d:areas", suburbId)
	if _, err := s.redis.SetEx(ctx, key, string(b), 24*time.Hour).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/CacheAreasBySuburbId", "section": "redis.SetEx"}).Error(err)
	}
}

func (s *ShippingImpl) FindTrackingByShippingId(ctx context.Context, shippingId string) *entity.Tracking {
	key := fmt.Sprintf("shipping_id:%s:trackings", shippingId)

	res, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindTrackingByShippingId", "section": "redis.Get"}).Error(err)
		}

		return nil
	}

	tracking := new(entity.Tracking)
	if err := json.Unmarshal([]byte(res), tracking); err != nil {

		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindTrackingByShippingId", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return tracking
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

	cities := new(dto.ShipperRes[[]*entity.City])
	if err := json.Unmarshal([]byte(res), cities); err != nil {

		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindCitiesByProvinceId", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return cities
}

func (s *ShippingImpl) FindSuburbsByCityId(ctx context.Context, cityId int) *dto.ShipperRes[[]*entity.Suburb] {
	key := fmt.Sprintf("city_id:%d:suburbs", cityId)

	res, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindSuburbsByCityId", "section": "redis.Get"}).Error(err)
		}

		return nil
	}

	suburbs := new(dto.ShipperRes[[]*entity.Suburb])
	if err := json.Unmarshal([]byte(res), suburbs); err != nil {

		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindSuburbsByCityId", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return suburbs
}

func (s *ShippingImpl) FindAreasBySuburbId(ctx context.Context, suburbId int) *dto.ShipperRes[[]*entity.Area] {
	key := fmt.Sprintf("suburb_id:%d:areas", suburbId)

	res, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindAreasBySuburbId", "section": "redis.Get"}).Error(err)
		}

		return nil
	}

	areas := new(dto.ShipperRes[[]*entity.Area])
	if err := json.Unmarshal([]byte(res), areas); err != nil {

		log.Logger.WithFields(logrus.Fields{"location": "cache.ShippingImpl/FindAreasBySuburbId", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return areas
}

func (s *ShippingImpl) UpdateTracking(ctx context.Context, shippingId string, data *entity.TrackingData) {
	if res := s.FindTrackingByShippingId(ctx, shippingId); res != nil {
		res.Trackings = append(res.Trackings, data)

		return
	}

	s.CacheTrackingByShippingId(ctx, shippingId, &entity.Tracking{
		Trackings: []*entity.TrackingData{data},
	})
}

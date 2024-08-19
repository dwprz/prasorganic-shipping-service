package helper

import "github.com/dwprz/prasorganic-shipping-service/src/model/entity"

func FormatTrackingData(data *entity.Shipper) *entity.TrackingData {
	return &entity.TrackingData{
		ShipperStatus: entity.Status{
			Code:        data.ExternalStatus.Code,
			Name:        data.ExternalStatus.Name,
			Description: data.ExternalStatus.Description,
		},
		LogisticStatus: entity.Status{
			Code:        data.External.Id,
			Name:        data.External.Name,
			Description: data.External.Description,
		},
		CreatedDate: data.StatusDate,
	}
}

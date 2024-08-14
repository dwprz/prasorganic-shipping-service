package entity

type Logistic struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LogoURL      string `json:"logo_url"`
	Code         string `json:"code"`
	CompanyName  string `json:"company_name"`
	CodFee       int    `json:"cod_fee"`
	CodMinAmount int    `json:"cod_min_amount"`
	CodMaxAmount int    `json:"cod_max_amount"`
}

type Rate struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Description     string `json:"description"`
	FullDescription string `json:"full_description"`
	IsHubless       bool   `json:"is_hubless"`
	IsDropOff       bool   `json:"is_drop_off"`
	IsMultikoli     bool   `json:"is_multikoli"`
}

type Pricing struct {
	Origin      Location `json:"origin"`
	Destination Location `json:"destination"`
	Pricings    []struct {
		Logistic                       Logistic `json:"logistic"`
		Rate                           Rate     `json:"rate"`
		Weight                         float32  `json:"weight"`
		Volume                         int      `json:"volume"`
		VolumeWeight                   float32  `json:"volume_weight"`
		FinalWeight                    int  `json:"final_weight"`
		MinDay                         int      `json:"min_day"`
		MaxDay                         int      `json:"max_day"`
		UnitPrice                      int      `json:"unit_price"`
		TotalPrice                     int      `json:"total_price"`
		Discount                       float32  `json:"discount"`
		DiscountValue                  int      `json:"discount_value"`
		DiscountedPrice                int      `json:"discounted_price"`
		InsuranceFee                   int      `json:"insurance_fee"`
		MustUseInsurance               bool     `json:"must_use_insurance"`
		LiabilityValue                 int      `json:"liability_value"`
		FinalPrice                     int      `json:"final_price"`
		Currency                       string   `json:"currency"`
		InsuranceApplied               bool     `json:"insurance_applied"`
		BasePrice                      int      `json:"base_price"`
		SurchargeFee                   int      `json:"surcharge_fee"`
		HeavyDutySurchargeFee          int      `json:"heavy_duty_surcharge_fee"`
		FuelSurchargeFee               int      `json:"fuel_surcharge_fee"`
		EmergencySituationSurchargeFee int      `json:"emergency_situation_surcharge_fee"`
		CodFeeAmount                   int      `json:"cod_fee_amount"`
		PlatformFee                    int      `json:"platform_fee"`
		PlatformFeeAmount              int      `json:"platform_fee_amount"`
	}
}

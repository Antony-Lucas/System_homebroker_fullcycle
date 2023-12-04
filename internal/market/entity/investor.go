package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetposition := range i.AssetPosition {
		if assetposition.AssetId == assetId {
			return assetposition
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetId string
	Shares  int
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetId: assetId,
		Shares:  shares,
	}
}
package service

import (
	adapter "github.com/test-tzs/nomraeite/internal/domain/adapter/makeshop"
	apiModel "github.com/test-tzs/nomraeite/internal/domain/model/api/makeshop"
)

type ShopDomainService interface {
	GetShopByChunkId(shopIDs []string, chunk int) (result []*apiModel.Shop, errorShopIDs []string)
}

type shopDomainServiceImpl struct {
	makeshopShopAdapter adapter.MakeshopShopAdapter
}

func NewShopDomainService(
	makeshopShopAdapter adapter.MakeshopShopAdapter,
) ShopDomainService {
	return &shopDomainServiceImpl{
		makeshopShopAdapter: makeshopShopAdapter,
	}
}

func (s *shopDomainServiceImpl) GetShopByChunkId(shopIDs []string, chunk int) (result []*apiModel.Shop, errorShopIDs []string) {

	makeshopShops := []*apiModel.Shop{}
	shopChunkIDs := []string{}

	// Implementation to fetch trunk shop by ID
	for i, shopID := range shopIDs {
		// chunk ShopID to get shop information
		shopChunkIDs = append(shopChunkIDs, shopID)
		if len(shopChunkIDs) == chunk || i == len(shopIDs)-1 {
			shops, err := s.makeshopShopAdapter.GetShopByIDs(shopChunkIDs)
			if err != nil {
				errorShopIDs = append(errorShopIDs, shopChunkIDs...)
			} else {
				// Append the fetched shops to makeshopShops
				makeshopShops = append(makeshopShops, shops...)
			}
			shopChunkIDs = []string{} // Reset shopIDs for the next chunk
		}
	}
	return makeshopShops, errorShopIDs

}

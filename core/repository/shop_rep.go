/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 17:16
 * description :
 * history :
 */

package repository

import (
	"fmt"
	"github.com/jsix/gof/db"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/merchant/shop"
	"go2o/core/infrastructure"
	"go2o/core/infrastructure/log"
	"go2o/core/variable"
)

var _ shop.IShopRep = new(shopRep)

type shopRep struct {
	db.Connector
}

func NewShopRep(c db.Connector) shop.IShopRep {
	return &shopRep{
		Connector: c,
	}
}

// 获取线上商店
func (this *shopRep) GetOnlineShop(shopId int) *shop.OnlineShop {
	e := shop.OnlineShop{}
	if this.GetOrm().Get(shopId, &e) != nil {
		return nil
	}
	return &e
}

// 保存线上商店
func (this *shopRep) SaveOnlineShop(v *shop.OnlineShop) error {
	_, _, err := this.GetOrm().Save(v.ShopId, v)
	if err != nil {
		_, _, err = this.GetOrm().Save(nil, v)
	}
	return err
}

// 获取线下商店
func (this *shopRep) GetOfflineShop(shopId int) *shop.OfflineShop {
	e := shop.OfflineShop{}
	if this.GetOrm().Get(shopId, &e) != nil {
		return nil
	}
	return &e
}

// 保存线下商店
func (this *shopRep) SaveOfflineShop(v *shop.OfflineShop) error {
	_, _, err := this.GetOrm().Save(v.ShopId, v)
	if err != nil {
		_, _, err = this.GetOrm().Save(nil, v)
	}
	return err
}

// 获取站点配置
func (this *shopRep) GetSiteConf(merchantId int) *shop.ShopSiteConf {
	var siteConf shop.ShopSiteConf
	if err := this.Connector.GetOrm().Get(merchantId, &siteConf); err == nil {
		if len(siteConf.Host) == 0 {
			var usr string
			this.Connector.ExecScalar(
				`SELECT usr FROM mch_merchant WHERE id=?`,
				&usr, merchantId)
			siteConf.Host = fmt.Sprintf("%s.%s", usr,
				infrastructure.GetApp().Config().
					GetString(variable.ServerDomain))
		}
		return &siteConf
	}
	return nil
}

func (this *shopRep) SaveSiteConf(merchantId int, v *shop.ShopSiteConf) error {
	var err error
	if v.MerchantId > 0 {
		_, _, err = this.Connector.GetOrm().Save(v.MerchantId, v)
	} else {
		v.MerchantId = merchantId
		_, _, err = this.Connector.GetOrm().Save(nil, v)
	}
	return err
}

// 保存API信息
func (this *shopRep) SaveApiInfo(v *merchant.ApiInfo) error {
	var err error
	orm := this.GetOrm()
	if v.MerchantId <= 0 {
		_, _, err = orm.Save(nil, v)
	} else {
		_, _, err = orm.Save(v.MerchantId, v)
	}
	return err
}

// 获取API信息
func (this *shopRep) GetApiInfo(merchantId int) *merchant.ApiInfo {
	var d *merchant.ApiInfo = new(merchant.ApiInfo)
	if err := this.GetOrm().Get(merchantId, d); err == nil {
		return d
	}
	return nil
}

func (this *shopRep) SaveShop(v *shop.Shop) (int, error) {
	orm := this.Connector.GetOrm()
	if v.Id > 0 {
		_, _, err := orm.Save(v.Id, v)
		return v.Id, err
	} else {
		_, _, err := orm.Save(nil, v)

		//todo: return id
		return 0, err
	}
}

func (this *shopRep) GetValueShop(merchantId, shopId int) *shop.Shop {
	var v *shop.Shop = new(shop.Shop)
	err := this.Connector.GetOrm().Get(shopId, v)
	if err == nil &&
		v.MerchantId == merchantId {
		return v
	} else {
		log.Error(err)
	}
	return nil
}

func (this *shopRep) GetShopsOfMerchant(merchantId int) []*shop.Shop {
	shops := []*shop.Shop{}
	err := this.Connector.GetOrm().SelectByQuery(&shops,
		"SELECT * FROM mch_shop WHERE merchant_id=?", merchantId)

	if err != nil {
		log.Error(err)
		return nil
	}

	return shops
}

func (this *shopRep) DeleteShop(merchantId, shopId int) error {
	_, err := this.Connector.GetOrm().Delete(shop.Shop{},
		"merchant_id=? AND id=?", merchantId, shopId)
	return err
}

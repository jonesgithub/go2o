/**
 * Copyright 2014 @ ops.
 * name :
 * author : jarryliu
 * date : 2013-11-11 19:51
 * description :
 * history :
 */

package partner

type SaleConf struct {
	//合作商编号
	PartnerId int `db:"pt_id" auto:"no" pk:"yes"`

	//反现比例,0则不返现
	CashBackPercent float32 `db:"cb_percent"`
	//一级比例
	CashBackTg1Percent float32 `db:"cb_tg1_percent"`
	//二级比例
	CashBackTg2Percent float32 `db:"cb_tg2_percent"`
	//会员比例
	CashBackMemberPercent float32 `db:"cb_member_percent"`
	//每一元返多少积分
	IntegralBackNum int `db:"ib_num"`
	//每单额外赠送
	IntegralBackExtra int `db:"ib_extra"`

	// 自动设置订单
	AutoSetupOrder int `db:"auto_setup_order"`
}

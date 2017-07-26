package models

import "flume-client/components/setting"

type RankListModel struct {
	ProductModel
	IPTimeModel
	AccountModel

	Type     string `json:"type" ini:"TYPE"`
	Value    int32 `json:"value" ini:"VALUE"`
	Comments string `json:"comments" ini:"COMMENTS"`
	Rank     int32 `json:"rank" ini:"RANK"`
}

var RankList RankListModel

func (RankListModel) Init() error {
	RankList = RankListModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.rank_list").MapTo(&RankList)
	if err != nil {
		return err
	}
	return nil
}

func (RankListModel) GetType() string {
	return "rank-list"
}

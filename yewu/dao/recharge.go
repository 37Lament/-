package dao

import (
	"yewu/model"
)

func Updatermb(username string, newrmb int) error {
		_, err := dB.Exec("UPDATE user SET rmb = ? WHERE username = ?", newrmb, username)
		return err
	}
func InserttradingRecord(tradingRecord model.TradingRecord) error {
	_, err := dB.Exec("INSERT INTO tradingrecord(username, forusername,time,transerRMB,accountBalanceNow,txt) "+"values(?, ?,?,?,?,?);", tradingRecord.Username, tradingRecord.Forusername,tradingRecord.Time,tradingRecord.TranserRMB,tradingRecord.AccountBalanceNow,tradingRecord.Txt)
	return err
}
// 查询多条数据示例
func QueryTxt(Qusername,Qtxt string)([]model.TradingRecord,error) {
	var tradingRecords []model.TradingRecord

	s:="%"+Qtxt+"%"
	rows, err := dB.Query("SELECT id,username, forusername, time ,transerRMB,accountBalanceNow,txt FROM tradingrecord WHERE txt LIKE ? and username = ? or forusername = ? ",s,Qusername,Qusername)
	if err != nil {
		return nil, err
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var TradingRecord model.TradingRecord
		err := rows.Scan(&TradingRecord.Id, &TradingRecord.Username, &TradingRecord.Forusername,&TradingRecord.Time,&TradingRecord.TranserRMB ,&TradingRecord.AccountBalanceNow,&TradingRecord.Txt)
		if err != nil {
			return nil, err
		}
		tradingRecords = append(tradingRecords, TradingRecord)
	}
	return tradingRecords, nil
}


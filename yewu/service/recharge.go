package service

import (
	"yewu/dao"
	"yewu/model"
)

func ChangeRmb(username string,newrmb int)error{
		err := dao.Updatermb(username,newrmb)
		return err
	}
func Userselect(username string) (model.User, error) {
	user, err := dao.SelectUserByUsername(username)
	return  user,err
}

func Registertrade(tradingRecord model.TradingRecord) error {
	err := dao.InserttradingRecord(tradingRecord)
	return err
}
func SearchByTxt(Qusername,Qtxt string)([]model.TradingRecord,error){
	return dao.QueryTxt(Qusername,Qtxt)

}
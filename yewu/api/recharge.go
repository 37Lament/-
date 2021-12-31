package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"yewu/model"
	"yewu/service"
	"yewu/tool"
)

func changeRMB(ctx *gin.Context) {
	username := ctx.PostForm("username")
	rechargeRMB:=ctx.PostForm("recharge_rmb")
	user, err := service.Userselect(username)
	if err != nil {
		fmt.Println("oldRMB err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	i,err := strconv.Atoi(rechargeRMB)
	if err != nil {
		fmt.Println("recharge rmb err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	newrmb := i+user.Rmb
	err = service.ChangeRmb(username, newrmb)
	if err != nil {
		fmt.Println("recharge rmb err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func transferAccounts(ctx *gin.Context){
	Payee:= ctx.PostForm("Payee")
	transferRMB:=ctx.PostForm("RMB")
	Tips:=ctx.PostForm("txt")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	user, err := service.Userselect(username)
	if err != nil {
		fmt.Println("oldRMB err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	i,err := strconv.Atoi(transferRMB)
	if err != nil {
		fmt.Println("recharge rmb err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	if user.Rmb<i {
		fmt.Println(" rmb buzhu: ", user.Rmb,"<",i)
		tool.RespczError(ctx)
		return
	}
	newrmb := user.Rmb-i
	err = service.ChangeRmb(username, newrmb)
	if err != nil {
		fmt.Println("oldRMB err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	foruser, err := service.Userselect(Payee)
	if err != nil {
		fmt.Println("oldRMB err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	if err != nil {
		fmt.Println("recharge rmb err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	newrmb2:=foruser.Rmb+i
	err = service.ChangeRmb(Payee, newrmb2)
	if err != nil {
		fmt.Println("recharge rmb err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	trading:=model.TradingRecord{
		Username: username,
		Forusername: Payee,
		Time: timenow,
		TranserRMB: i,
		AccountBalanceNow: newrmb,
		Txt: Tips,
	}
	err=service.Registertrade(trading)
	if err != nil {
		fmt.Println("Registertrade err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

func searchByTxt(ctx *gin.Context)  {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	Qusername:= username
	Qtxt:=ctx.PostForm("txt")
	tradingRecords,err:=service.SearchByTxt(Qusername,Qtxt)
	if err != nil {
		fmt.Println("searchByTxt err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithimformation(ctx,tradingRecords)

}
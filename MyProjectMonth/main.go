package main

import (
	"time"
	"./copy"
	"./delete"
	"./detectmonth"
)

func main(){
	
	b := time.Now().Format("20060102")
	b2 := detectmonth.DetectMonth(time.Now().Month().String())
	b3 := time.Now().Format("2006")
	
	//SS7 Month
	srcPath :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/MONTH_REPORT_SS7/SS7_MONTH_REPORT_"+b+".xls"
	destPath :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/статистика/для подсчета ПП/SS7_MONTH_REPORT "+b2+" "+b3+".xls"
    copy.File(srcPath, destPath)
	delete.DeleteFile(srcPath)
	}
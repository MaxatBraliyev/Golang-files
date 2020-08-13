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
	
	//SMPP1 Month Report
	srcPath1 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/MONTH_REPORT_SMPP/Месячная статистика по немаркированным SMS с SMPP за_"+b+".xls"
	destPath1 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежемесячные отчеты/Месячная статистика по немаркированным SMS с SMPP за "+b2+" "+b3+".xls"
    copy.File(srcPath1, destPath1)
	delete.DeleteFile(srcPath1)
	
	//SMPP2 Month Report
	srcPath2 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/MONTH_REPORT_SMPP/Месячная статистика по маркированным SMS с SMPP за_"+b+".xls"
	destPath2 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежемесячные отчеты/Месячная статистика по маркированным SMS с SMPP за "+b2+" "+b3+".xls"
    copy.File(srcPath2, destPath2)
	delete.DeleteFile(srcPath2)
	
	//SMPP3 Month Report
	srcPath3 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/MONTH_REPORT_SMPP/Статистика по правилам за_"+b+".xls"
	destPath3 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежемесячные отчеты/Статистика по правилам за "+b2+" "+b3+".xls"
    copy.File(srcPath3, destPath3)
	delete.DeleteFile(srcPath3)
	
	}
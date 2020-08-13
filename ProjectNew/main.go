package main

import (
	"time"
	"./mkrdir"
	"./copy"
	"./delete"
)

func main(){
	
	b := time.Now().Format("20060102")
	t := time.Now().Format("02.01.2006")
	today := time.Now()
	oneDay := time.Hour * -24
	t2 :=today.Add(oneDay).Format("02012006")
	
	mkrdir.MkrDir("//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/")
	
	//1 SMPP
	srcPathSmpp1 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SMPP/Детальные данные по немаркированным SMS локал A2P за_"+b+".csv"
	destPathSmpp1 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/Детальные данные по немаркированным SMS локал A2P за_"+t2+".csv"
    copy.File(srcPathSmpp1, destPathSmpp1)
	delete.DeleteFile(srcPathSmpp1)
	
	//2 SMPP
	srcPathSmpp2 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SMPP/Детальные данные по немаркированным SMS межнар A2P за_"+b+".csv"
	destPathSmpp2 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/Детальные данные по немаркированным SMS межнар A2P за_"+t2+".csv"
    copy.File(srcPathSmpp2, destPathSmpp2)
	delete.DeleteFile(srcPathSmpp2)
	
	//3 SMPP
	srcPathSmpp3 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SMPP/Детальные данные по отбитым sms SMPP за_"+b+".xls"
	destPathSmpp3 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/Детальные данные по отбитым sms SMPP за_"+t2+".xls"
    copy.File(srcPathSmpp3, destPathSmpp3)
	delete.DeleteFile(srcPathSmpp3)
	
	//4 SMPP
	srcPathSmpp4 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SMPP/Статистика по маркированным SMS с SMPP за_"+b+".xls"
	destPathSmpp4 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/Статистика по маркированным SMS с SMPP за_"+t2+".xls"
    copy.File(srcPathSmpp4, destPathSmpp4)
	delete.DeleteFile(srcPathSmpp4)
	
	//5 SMPP
	srcPathSmpp5 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SMPP/Статистика по немаркированным SMS с SMPP за_"+b+".xls"
	destPathSmpp5 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/Статистика по немаркированным SMS с SMPP за_"+t2+".xls"
    copy.File(srcPathSmpp5, destPathSmpp5)
	delete.DeleteFile(srcPathSmpp5)
	
	//6 SMPP RULES
	srcPathSmpp6 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/RULES/RULES_SMPP_"+b+".xls"
	destPathSmpp6 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/RULES_SMPP_"+t+".xls"
    copy.File(srcPathSmpp6, destPathSmpp6)
	delete.DeleteFile(srcPathSmpp6)
	
	//7 SMPP DUBLICATED
	srcPathSmpp7 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SMPP/Повторяющийся текст с SMPP за_"+b+".xls"
	destPathSmpp7 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SMPP/Ежедневные отчеты/"+t+"/Повторяющийся текст с SMPP за_"+t2+".xls"
    copy.File(srcPathSmpp7, destPathSmpp7)
	delete.DeleteFile(srcPathSmpp7)
	
}
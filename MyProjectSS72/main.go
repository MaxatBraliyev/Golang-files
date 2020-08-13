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
	firstDay := time.Hour * -24
	lastDay := time.Hour * -96
	t3 :=today.Add(firstDay).Format("02.01.2006")
	t2 :=today.Add(lastDay).Format("02.01.2006")
	
	mkrdir.MkrDir("//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Агрегированные и хэшированные/")
	mkrdir.MkrDir("//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/")
	
	//1 SS7 agg
	srcPathSs71 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/agg_ss7_"+b+".xls"
	destPathSs71 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Агрегированные и хэшированные/"+t+"/agg_ss7 "+t2+" - "+t3+".xls"
    copy.File(srcPathSs71, destPathSs71)
	delete.DeleteFile(srcPathSs71)
	
	//2 SS7 hash
	srcPathSs72 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/tz_ss7_hash_"+b+".csv"
	destPathSs72 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Агрегированные и хэшированные/"+t+"/tz_ss7_hash "+t2+" - "+t3+".csv"
    copy.File(srcPathSs72, destPathSs72)
	delete.DeleteFile(srcPathSs72)
	
	//3 SS7 rejected
	srcPathSs73 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/rejected_sms_ss7_"+b+".xls"
	destPathSs73 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/"+t+"/rejected_sms_ss7 "+t2+" - "+t3+".xls"
    copy.File(srcPathSs73, destPathSs73)
	delete.DeleteFile(srcPathSs73)
	
	//4 SS7 details1
	srcPathSs74 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/details_ss7_part1_"+b+".csv"
	destPathSs74 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/"+t+"/details_ss7_part1 "+t2+" - "+t3+".csv"
    copy.File(srcPathSs74, destPathSs74)
	delete.DeleteFile(srcPathSs74)
	
	//5 SS7 details2
	srcPathSs75 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/details_ss7_part2_"+b+".csv"
	destPathSs75 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/"+t+"/details_ss7_part2 "+t2+" - "+t3+".csv"
    copy.File(srcPathSs75, destPathSs75)
	delete.DeleteFile(srcPathSs75)
	
	//6 SS7 details3
	srcPathSs76 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/details_ss7_part3_"+b+".csv"
	destPathSs76 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/"+t+"/details_ss7_part3 "+t2+" - "+t3+".csv"
    copy.File(srcPathSs76, destPathSs76)
	delete.DeleteFile(srcPathSs76)
	
	//7 SS7 details4
	srcPathSs77 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/details_ss7_part4_"+b+".csv"
	destPathSs77 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/"+t+"/details_ss7_part4 "+t2+" - "+t3+".csv"
    copy.File(srcPathSs77, destPathSs77)
	delete.DeleteFile(srcPathSs77)
	
	//8 SS7 details5
	srcPathSs78 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/Максат/OutFiles/SS7/details_ss7_part5_"+b+".csv"
	destPathSs78 :="//jv.mt-s.kz/fs/Internal Control/RA and Fraud/A2P Проект/Выгрузки/details/SS7/Детальные отчеты/"+t+"/details_ss7_part5 "+t2+" - "+t3+".csv"
    copy.File(srcPathSs78, destPathSs78)
	delete.DeleteFile(srcPathSs78)
	}
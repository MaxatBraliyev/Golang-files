package detectmonth

func DetectMonth(v string) string{
	if v=="January"{v="Декабрь"}
	if v=="February"{v="Январь"}
	if v=="March"{v="Февраль"}
	if v=="April"{v="Март"}
	if v=="May"{v="Апрель"}
	if v=="June"{v="Май"}
	if v=="July"{v="Июнь"}
	if v=="August"{v="Июль"}
	if v=="September"{v="Август"}
	if v=="October"{v="Сентябрь"}
	if v=="November"{v="Октябрь"}
	if v=="December"{v="Ноябрь"}
	return v
}
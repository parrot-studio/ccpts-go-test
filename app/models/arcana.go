package models

type Arcana struct {
    Id  uint64 `gorm:"primary_key" json:"id"`
		Name string `sql:"size:100" json:"name"`
		Title string `sql:"size:100" json:"title"`
		ArcanaType string `sql:"size:20" json:"arcana_type"`
		Rarity uint32 `json:"rarity"`
		Cost uint32 `json:"cost"`
		ChainCost uint32 `json:"chain_cost"`
		WeaponType string `sql:"size:10" json:"weapon_type"`
		JobType string `sql:"size:10" json:"job_type"`
		JobIndex uint32 `json:"job_index"`
		JobCode string `sql:"size:10" json:"job_code"`
		JobDetail string `sql:"size:50" json:"job_detail"`
		SourceCategory string `sql:"size:50" json:"source_category"`
		Source string `sql:"size:50" json:"source"`
		Union string `sql:"size:20" json:"union"`
		PersonCode string `sql:"size:10" json:"person_code"`
		LinkCode string `sql:"size:10" json:"link_code"`
		MaxAtk uint32 `json:"max_atk"`
		MaxHp uint32 `json:"max_Hp"`
		LimitAtk uint32 `json:"limit_atk"`
		LimitHp uint32 `json:"limit_hp"`
}

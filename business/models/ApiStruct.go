package models

import (
	"time"


)
var _ = time.Time{}
var m time.Month


//home page
type Contact struct {
	ID     			    int        `gorm:"column:id" json:"id"`
	PHONEONE 			string		   `gorm:"column:phone_one" json:"phone_one"`
	PHONETWO		    string 	   `gorm:"column:phone_two" json:"phone_two"`
	FAX                 string     `gorm:"column:fax" json:"fax"`
	FACEBOOK            string     `gorm:"column:facebook" json:"facebook"`
	TWITER              string     `gorm:"column:twiter" json:"twiter"`
	MAIL                string     `gorm:"column:mail" json:"mail"`
	ADDRESS                string     `gorm:"column:address" json:"address"`
	BAIRSHIL                string     `gorm:"column:bairshil" json:"bairshil"`
	YOUTUBE                string     `gorm:"column:youtube" json:"youtube"`


}
type Comments struct {
	ID				int	   		    `gorm:"column:id"  json:"id"`
	NER				string			`gorm:"column:ner"  json:"ner"`
	COMMENTS		string			`gorm:"column:comments"  json:"comments"`
	MAIL            string          `gorm:"column:mail" json:"mail"`
}
type Buteegdehuun struct {
	AimagID           int        `gorm:"column:aimag_id" json:"aimag_id"`
	BagID             int        `gorm:"column:bag_id" json:"bag_id"`
	Buteegdehuun      string     `gorm:"column:buteegdehuun" json:"buteegdehuun"`
	CompanyID         int        `gorm:"column:company_id" json:"company_id"`
	DedSalbarID       int        `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	DedTorolID        int        `gorm:"column:ded_torol_id" json:"ded_torol_id"`
	ID                int        `gorm:"column:id;primary_key" json:"id"`
	Ontsgoi           int        `gorm:"column:ontsgoi" json:"ontsgoi"`
	SalbarID          int        `gorm:"column:salbar_id" json:"salbar_id"`
	SumID             int        `gorm:"column:sum_id" json:"sum_id"`
	Tailbar           string     `gorm:"column:tailbar" json:"tailbar"`
	TorolID           int        `gorm:"column:torol_id" json:"torol_id"`
	Tseg              string     `gorm:"column:tseg" json:"tseg"`
	UildverlesenOgnoo *time.Time `gorm:"column:uildverlesen_ognoo" json:"uildverlesen_ognoo"`
	Zurag             string     `gorm:"column:zurag" json:"zurag"`
	Zuragnuud         string     `gorm:"column:zuragnuud" json:"zuragnuud"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
type HamtragchBaiguullaga struct {
	ID   int    `gorm:"column:id;primary_key" json:"id"`
	Logo string `gorm:"column:logo" json:"logo"`
	Ner  string `gorm:"column:ner" json:"ner"`
	URL  string `gorm:"column:url" json:"url"`
}
type Uilchilgee struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	AimagID      int        `gorm:"column:aimag_id" json:"aimag_id"`
	BagID        int        `gorm:"column:bag_id" json:"bag_id"`
	CompanyID    int        `gorm:"column:company_id" json:"company_id"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	DedTorolID   int        `gorm:"column:ded_torol_id" json:"ded_torol_id"`
	Fax          string     `gorm:"column:fax" json:"fax"`
	GarUtas      int        `gorm:"column:gar_utas" json:"gar_utas"`
	Hayag        string     `gorm:"column:hayag" json:"hayag"`
	HuchinChadal string     `gorm:"column:huchin_chadal" json:"huchin_chadal"`
	Mail         string     `gorm:"column:mail" json:"mail"`
	Name         string     `gorm:"column:name" json:"name"`
	Social       string     `gorm:"column:social" json:"social"`
	SumID        int        `gorm:"column:sum_id" json:"sum_id"`
	Tailbar      string     `gorm:"column:tailbar" json:"tailbar"`
	TorolID      int        `gorm:"column:torol_id" json:"torol_id"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas         string        `gorm:"column:utas" json:"utas"`
	Zurag        string     `gorm:"column:zurag" json:"zurag"`
	Zuragnuud         string     `gorm:"column:zuragnuud" json:"zuragnuud"`
}
type CompanyRegister struct {
	AimagID         int        `gorm:"column:aimag_id" json:"aimag_id"`
	AjAhiuNegjID    int        `gorm:"column:aj_ahiu_negj_id" json:"aj_ahiu_negj_id"`
	AjiltanTooID    int        `gorm:"column:ajiltan_too_id" json:"ajiltan_too_id"`
	BagID           int        `gorm:"column:bag_id" json:"bag_id"`
	//BaiguulsanOgnoo *time.Time `gorm:"column:baiguulsan_ognoo" json:"baiguulsan_ognoo"`
	//CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	SalbarID        int        `gorm:"column:salbar_id" json:"salbar_id"`
	DedSalbarID     int        `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	//DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Description     string     `gorm:"column:description" json:"description"`
	Director        string     `gorm:"column:director" json:"director"`
	Fax             string     `gorm:"column:fax" json:"fax"`
	ID              int        `gorm:"column:id;primary_key" json:"id"`
	Mail            string     `gorm:"column:mail" json:"mail"`
	Name            string     `gorm:"column:name" json:"name"`
	PhoneOne        string        `gorm:"column:phone_one" json:"phone_one"`
	PhoneTwo        string        `gorm:"column:phone_two" json:"phone_two"`
	Photo           string     `gorm:"column:photo" json:"photo"`
	Photos          string     `gorm:"column:photos" json:"photos"`
	Plus            string     `gorm:"column:plus" json:"plus"`
	Register        string     `gorm:"column:register" json:"register"`
	Social          string     `gorm:"column:social" json:"social"`
	SumID           int        `gorm:"column:sum_id" json:"sum_id"`
	UilChiglelID    int        `gorm:"column:uil_chiglel_id" json:"uil_chiglel_id"`
	//UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	VideoLink       string     `gorm:"column:video_link" json:"video_link"`
}
type NuurSlide struct {
	ID    int    `gorm:"column:id;primary_key" json:"id"`
	Zurag string `gorm:"column:zurag" json:"zurag"`
}
type Aimag struct {
	Aimagcode    string `gorm:"column:aimagcode" json:"aimagcode"`
	Aimagname    string `gorm:"column:aimagname" json:"aimagname"`
	AimagnameMgl string `gorm:"column:aimagname_mgl" json:"aimagname_mgl"`
	Aimagzipcode string `gorm:"column:aimagzipcode" json:"aimagzipcode"`
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	Location     string `gorm:"column:location" json:"location"`
	Zoneid       int    `gorm:"column:zoneid" json:"zoneid"`
	Zurag       string    `gorm:"column:zurag" json:"zurag"`
}


// count all by Paginate
type SalbarCountAll struct {
	AanToo      int    `gorm:"column:aan_too" json:"aan_too"`
	ButeegdehuunToo      int    `gorm:"column:buteegdehuun_too" json:"buteegdehuun_too"`
	UilchilgeeToo      int    `gorm:"column:uilchilgee_too" json:"uilchilgee_too"`
	TpSectionId      int  `gorm:"column:tp_section_id" json:"tp_section_id"`
	TPSection string `gorm:"column:tp_section" json:"tp_section"`
	TPZurag string `gorm:"column:zurag" json:"zurag"`
	Icon string `gorm:"column:icon" json:"icon"`

}
type AimagCountAll struct {
	AanToo      int    `gorm:"column:aan_too" json:"aan_too"`
	ButeegdehuunToo      int    `gorm:"column:buteegdehuun_too" json:"buteegdehuun_too"`
	UilchilgeeToo       int    `gorm:"column:uilchilgee_too" json:"uilchilgee_too"`
	SalbarToo           int    `gorm:"column:salbar_too" json:"salbar_too"`
	Zurag              string    `gorm:"column:zurag" json:"zurag"`
	TpSectionId        int  `gorm:"column:tp_section_id" json:"tp_section_id"`
	TPSection          string `gorm:"column:tp_section" json:"tp_section"`
	AimagId            int  `gorm:"column:aimag_id" json:"aimag_id"`
	AimagName          string  `gorm:"column:aimagname" json:"aimagname"`

}



 // count example : *muhadai
type AnnAimagCount struct {
	Too      int    `gorm:"column:too" json:"too"`
	AimagID      int  `gorm:"column:aimag_id" json:"aimag_id"`
	AimagName string `gorm:"column:aimagname" json:"aimagname"`
}
type AanSalbarCount struct {
	Too      int    `gorm:"column:too" json:"too"`
	SalbarId      int  `gorm:"column:salbar_id" json:"salbar_id"`
	TPSection string `gorm:"column:tp_section" json:"tp_section"`
}
type AanAimagSalbaraar struct {
	AimagId string `gorm:"column:id" json:"id"`
	Aimagname string `gorm:"column:aimagname" json:"aimagname"`
	Salbars []AanSalbarCount
}

type DSButeegdehuun struct {
	AimagID           int        `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname         string     `gorm:"column:aimagname" json:"aimagname"`
	AjAhiuNegjID      int        `gorm:"column:aj_ahiu_negj_id" json:"aj_ahiu_negj_id"`
	BagID             int        `gorm:"column:bag_id" json:"bag_id"`
	Bagname           string     `gorm:"column:bagname" json:"bagname"`
	BaiguullagaID     int        `gorm:"column:baiguullaga_id" json:"baiguullaga_id"`
	BaiguullagaIds    int        `gorm:"column:baiguullaga_ids" json:"baiguullaga_ids"`
	Buteegdehuun      string     `gorm:"column:buteegdehuun" json:"buteegdehuun"`
	CompanyID         int        `gorm:"column:company_id" json:"company_id"`
	CompanyName       string     `gorm:"column:company_name" json:"company_name"`
	PhoneOne          string     `gorm:"column:phone_one" json:"phone_one"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	DedSalbarID       int        `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	DedTorol          string     `gorm:"column:ded_torol" json:"ded_torol"`
	DedTorolID        int        `gorm:"column:ded_torol_id" json:"ded_torol_id"`
	ID                int        `gorm:"column:id" json:"id"`
	Ontsgoi           int        `gorm:"column:ontsgoi" json:"ontsgoi"`
	SalbarID          int        `gorm:"column:salbar_id" json:"salbar_id"`
	SumID             int        `gorm:"column:sum_id" json:"sum_id"`
	Sumname           string     `gorm:"column:sumname" json:"sumname"`
	Tailbar           string     `gorm:"column:tailbar" json:"tailbar"`
	Test              *time.Time `gorm:"column:test" json:"test"`
	Torol             string     `gorm:"column:torol" json:"torol"`
	TorolID           int        `gorm:"column:torol_id" json:"torol_id"`
	TpSection         string     `gorm:"column:tp_section" json:"tp_section"`
	Tseg              string     `gorm:"column:tseg" json:"tseg"`
	UilAChID          int        `gorm:"column:uil_a_ch_id" json:"uil_a_ch_id"`
	UilchilgeeChID    int        `gorm:"column:uilchilgee_ch_id" json:"uilchilgee_ch_id"`
	UildverlesenOgnoo *time.Time `gorm:"column:uildverlesen_ognoo" json:"uildverlesen_ognoo"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Zurag             string     `gorm:"column:zurag" json:"zurag"`
	Zuragnuud         string     `gorm:"column:zuragnuud" json:"zuragnuud"`
	JuulchlalID       int        `gorm:"column:juulchlal_id" json:"juulchlal_id"`
	NerJ              string     `gorm:"column:ner_j" json:"ner_j"`



}



type DSUilchilgee struct {
	AimagID        int        `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname      string     `gorm:"column:aimagname" json:"aimagname"`
	BagID          int        `gorm:"column:bag_id" json:"bag_id"`
	Bagname        string     `gorm:"column:bagname" json:"bagname"`
	CompanyID      int        `gorm:"column:company_id" json:"company_id"`
	Companyname    string     `gorm:"column:companyname" json:"companyname"`
	DedSalbarID    int        `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	DedTorol       string     `gorm:"column:ded_torol" json:"ded_torol"`
	DedTorolID     int        `gorm:"column:ded_torol_id" json:"ded_torol_id"`
	Fax            string     `gorm:"column:fax" json:"fax"`
	GarUtas        string     `gorm:"column:gar_utas" json:"gar_utas"`
	Hayag          string     `gorm:"column:hayag" json:"hayag"`
	HuchinChadal   string     `gorm:"column:huchin_chadal" json:"huchin_chadal"`
	ID             int        `gorm:"column:id" json:"id"`
	Mail           string     `gorm:"column:mail" json:"mail"`
	Name           string     `gorm:"column:name" json:"name"`
	SalbarID       int        `gorm:"column:salbar_id" json:"salbar_id"`
	Social         string     `gorm:"column:social" json:"social"`
	SumID          int        `gorm:"column:sum_id" json:"sum_id"`
	Sumname        string     `gorm:"column:sumname" json:"sumname"`
	Tailbar        string     `gorm:"column:tailbar" json:"tailbar"`
	Test           *time.Time `gorm:"column:test" json:"test"`
	Torol          string     `gorm:"column:torol" json:"torol"`
	TorolID        int        `gorm:"column:torol_id" json:"torol_id"`
	TpSection      string     `gorm:"column:tp_section" json:"tp_section"`
	UBaiguullagaID int        `gorm:"column:u_baiguullaga_id" json:"u_baiguullaga_id"`
	UilAChID       int        `gorm:"column:uil_a_ch_id" json:"uil_a_ch_id"`
	UilchilgeeChID int        `gorm:"column:uilchilgee_ch_id" json:"uilchilgee_ch_id"`
	Utas           string     `gorm:"column:utas" json:"utas"`
	Zurag          string     `gorm:"column:zurag" json:"zurag"`
	Zuragnuud      string     `gorm:"column:zuragnuud" json:"zuragnuud"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	JuulchlalID       int        `gorm:"column:juulchlal_id" json:"juulchlal_id"`
	NerJ              string     `gorm:"column:ner_j" json:"ner_j"`


}




type AanAimagCompanys struct {
	AimagId string `gorm:"column:id" json:"id"`
	Aimagname string `gorm:"column:aimagname" json:"aimagname"`
	Aan []AanAimagAan
}
type AanAimagAan struct {
	CompanyID              int        `gorm:"column:id;primary_key" json:"id"`
	Name            string     `gorm:"column:name" json:"name"`

}
type AanAimagCounts struct {
	Too      int    `gorm:"column:too" json:"too"`
	AimagId      int  `gorm:"column:aimag_id" json:"aimag_id"`
	AimagName string `gorm:"column:aimagname" json:"aimagname"`
}

//buteegdehuun
type BTorol struct {
	ID            int    `gorm:"column:id;primary_key" json:"id"`
	Too            int    `gorm:"column:too;" json:"too"`
	LastMonthToo            int    `gorm:"column:last_month_too;" json:"last_month_too"`
	Torol         string `gorm:"column:torol" json:"torol"`
}

func (b *BTorol) TableName() string {
	return "b_torol"
}

// uilchilgee


type UTorol struct {
	ID             int    `gorm:"column:id;primary_key" json:"id"`
	Torol          string `gorm:"column:torol" json:"torol"`
	Too            int    `gorm:"column:too;" json:"too"`
	LastMonthToo            int    `gorm:"column:last_month_too;" json:"last_month_too"`

}

func (u *UTorol) TableName() string {
	return "u_torol"
}






type ButeegdehuunDedTorol struct {
	Total     			 int   			 `gorm:"column:total" json:"total"`
	DedTorolID  	     string        `gorm:"column:ded_torol_id" json:"ded_torol_id"`
	DedTorol    	     string        `gorm:"column:ded_torol" json:"ded_torol"`


}
type ButeegdehuunTorol struct {
	Too      int    `gorm:"column:too" json:"too"`
	TorolID     		 int  			`gorm:"column:torol_id" json:"torol_id"`
	Torol				 string 		`gorm:"column:torol" json:"torol"`
}
type ButeegdehuunAimag struct {
	Too      int    `gorm:"column:too" json:"too"`
	AimagID           int        `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname         string     `gorm:"column:aimagname" json:"aimagname"`
}
type ButeegdehuunSalbar struct {
	Too      int    `gorm:"column:too" json:"too"`
	SalbarId      int  `gorm:"column:salbar_id" json:"salbar_id"`
	TPSection string `gorm:"column:tp_section" json:"tp_section"`
}

//company
type ProductsByCompany struct {
	Too               int        `gorm:"column:too" json:"too"`
	CompanyID         int        `gorm:"column:company_id" json:"company_id"`
	Buteegdehuun      string     `gorm:"column:buteegdehuun" json:"buteegdehuun"`
}

//baiguullaga
type CompanySub struct {
	ID             int `gorm:"column:id;primary_key" json:"id"`
	CompanyID      int `gorm:"column:company_id" json:"company_id"`
	DedSalbarID    int `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	SalbarID       int `gorm:"column:salbar_id" json:"salbar_id"`
	UilchilgeeChID int `gorm:"column:uilchilgee_ch_id" json:"uilchilgee_ch_id"`
}

func (c *CompanySub) TableName() string {
	return "company_sub"
}
type DSCompanyRegister struct {
	AimagID         int        `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname       string     `gorm:"column:aimagname" json:"aimagname"`
	AjAhiuNegjID    int        `gorm:"column:aj_ahiu_negj_id" json:"aj_ahiu_negj_id"`
	AjiltanTooID    int        `gorm:"column:ajiltan_too_id" json:"ajiltan_too_id"`
	BagID           int        `gorm:"column:bag_id" json:"bag_id"`
	BaiguulsanOgnoo *time.Time `gorm:"column:baiguulsan_ognoo" json:"baiguulsan_ognoo"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Description     string     `gorm:"column:description" json:"description"`
	Director        string     `gorm:"column:director" json:"director"`
	Fax             string     `gorm:"column:fax" json:"fax"`
	ID              int        `gorm:"column:id" json:"id"`
	Mail            string     `gorm:"column:mail" json:"mail"`
	Name            string     `gorm:"column:name" json:"name"`
	Negj            string     `gorm:"column:negj" json:"negj"`
	PhoneOne        string     `gorm:"column:phone_one" json:"phone_one"`
	PhoneTwo        string     `gorm:"column:phone_two" json:"phone_two"`
	Photo           string     `gorm:"column:photo" json:"photo"`
	Photos          string     `gorm:"column:photos" json:"photos"`
	Plus            string     `gorm:"column:plus" json:"plus"`
	Register        string     `gorm:"column:register" json:"register"`
	Social          string     `gorm:"column:social" json:"social"`
	SumID           int        `gorm:"column:sum_id" json:"sum_id"`
	Sumname         string     `gorm:"column:sumname" json:"sumname"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	VideoLink       string     `gorm:"column:video_link" json:"video_link"`
}

func (d *DSCompanyRegister) TableName() string {
	return "ds_company_register"
}

type DSCompanyRegisterDetail struct {
	AimagID         int        `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname       string     `gorm:"column:aimagname" json:"aimagname"`
	AjAhiuNegjID    int        `gorm:"column:aj_ahiu_negj_id" json:"aj_ahiu_negj_id"`
	AjiltanTooID    int        `gorm:"column:ajiltan_too_id" json:"ajiltan_too_id"`
	BagID           int        `gorm:"column:bag_id" json:"bag_id"`
	Bagname         string     `gorm:"column:bagname" json:"bagname"`
	BaiguulsanOgnoo *time.Time `gorm:"column:baiguulsan_ognoo" json:"baiguulsan_ognoo"`
	Chiglel         string     `gorm:"column:chiglel" json:"chiglel"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	DedSalbarID     int        `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Description     string     `gorm:"column:description" json:"description"`
	Director        string     `gorm:"column:director" json:"director"`
	Fax             string     `gorm:"column:fax" json:"fax"`
	ID              int        `gorm:"column:id" json:"id"`
	Mail            string     `gorm:"column:mail" json:"mail"`
	Name            string     `gorm:"column:name" json:"name"`
	PhoneOne        string        `gorm:"column:phone_one" json:"phone_one"`
	PhoneTwo        string        `gorm:"column:phone_two" json:"phone_two"`
	Photo           string     `gorm:"column:photo" json:"photo"`
	Photos          string     `gorm:"column:photos" json:"photos"`
	Plus            string     `gorm:"column:plus" json:"plus"`
	Register        string     `gorm:"column:register" json:"register"`
	SalbarID        int        `gorm:"column:salbar_id" json:"salbar_id"`
	Social          string     `gorm:"column:social" json:"social"`
	SumID           int        `gorm:"column:sum_id" json:"sum_id"`
	Sumname         string     `gorm:"column:sumname" json:"sumname"`
	TpSection       string     `gorm:"column:tp_section" json:"tp_section"`
	TpSubSection    string     `gorm:"column:tp_sub_section" json:"tp_sub_section"`
	UilchilgeeChID  int        `gorm:"column:uilchilgee_ch_id" json:"uilchilgee_ch_id"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	VideoLink       string     `gorm:"column:video_link" json:"video_link"`
	Zurag           string     `gorm:"column:zurag" json:"zurag"`
	UilAjilgaaChigleluud []DSCompanySub `gorm:"column:uil_ajilgaa_chigleluud" json:"uil_ajilgaa_chigleluud"`
	Buteegdehuuns []DSButeegdehuun `gorm:"column:buteegdehuuns" json:"buteegdehuuns"`
	Uilchilgeenuud []DSUilchilgee `gorm:"column:uilchilgeenuud" json:"uilchilgeenuud"`
}
func (d *DSCompanyRegisterDetail) TableName() string {
	return "ds_company_register"
}
type DSCompanySub struct {
	ID             int    `gorm:"column:id" json:"id"`
	Chiglel        string `gorm:"column:chiglel" json:"chiglel"`
	CompanyID      int    `gorm:"column:company_id" json:"company_id"`
	DedSalbarID    int    `gorm:"column:ded_salbar_id" json:"ded_salbar_id"`
	SalbarID       int    `gorm:"column:salbar_id" json:"salbar_id"`
	TpSection      string `gorm:"column:tp_section" json:"tp_section"`
	TpSubSection   string `gorm:"column:tp_sub_section" json:"tp_sub_section"`
	UilchilgeeChID int    `gorm:"column:uilchilgee_ch_id" json:"uilchilgee_ch_id"`
}
func (d *DSCompanySub) TableName() string {
	return "ds_company_sub"
}


//nemelt




//uilchilgee
type ServiceAngilalCount struct {
	Too      int    `gorm:"column:too" json:"too"`
	TorolId      int  `gorm:"column:torol_id" json:"torol_id"`
	Torol string `gorm:"column:torol" json:"torol"`
}
type ServiceSalbarCount struct {
	Too      int    `gorm:"column:too" json:"too"`
	SalbarId      int  `gorm:"column:salbar_id" json:"salbar_id"`
	TPSection string `gorm:"column:tp_section" json:"tp_section"`
}
type ServiceCompany struct {
	CompanyID            int        `gorm:"column:id" json:"id"`
	Companyname          string `gorm:"column:companyname" json:"companyname"`
	Total     			 int   			 `gorm:"column:total" json:"total"`

}
type ServiceTorol struct {
	TorolID  	      string        `gorm:"column:id" json:"id"`
	Torol  	          string        `gorm:"column:torol" json:"torol"`
	Angilaluud []ServiceCompany

}
type ServiceAimagCount struct {
	Too      int    `gorm:"column:too" json:"too"`
	AimagId      int  `gorm:"column:aimag_id" json:"aimag_id"`
	AimagName string `gorm:"column:aimagname" json:"aimagname"`
}

//REQUEST
type RequestCompany struct {
	AiamagIDS  []int `json:"aimag_ids" form:"aimag_ids" query:"aimag_ids"`
	SalbarIDs []int `json:"salbar_ids" form:"salbar_ids" query:"salbar_ids"`
	Ner string `json:"ner" form:"ner" query:"ner"`
	OrderType string `json:"orderType" form:"orderType" query:"orderType"`
	Order string `json:"order" form:"order" query:"order"`
}
type RequestProduct struct {
	AiamagIDS       []int `json:"aimag_ids" form:"aimag_ids" query:"aimag_ids"`
	SalbarIDs []int `json:"salbar_ids" form:"salbar_ids" query:"salbar_ids"`
	TorolIDs        []int `json:"torol_ids" form:"torol_ids" query:"torol_ids"`
	//DedTorolIDs      []int    `gorm:"column:ded_torol_ids" json:"ded_torol_ids" query:"ded_torol_ids"`
	Buteegdehuun     string `json:"buteegdehuun" form:"buteegdehuun" query:"buteegdehuun"`
	OrderType string `json:"orderType" form:"orderType" query:"orderType"`
	Order string `json:"order" form:"order" query:"order"`
}
type RequestService struct {
	AiamagIDS  []int `json:"aimag_ids" form:"aimag_ids" query:"aimag_ids"`
	TorolIDs []int `json:"torol_ids" form:"torol_ids" query:"torol_ids"`
	SalbarIDs []int `json:"salbar_ids" form:"salbar_ids" query:"salbar_ids"`
	Name string `json:"name" form:"name" query:"name"`
	OrderType string `json:"orderType" form:"orderType" query:"orderType"`
	Order string `json:"order" form:"order" query:"order"`
}
type RequestTravel struct {
	AiamagIDS       []int `json:"aimag_ids" form:"aimag_ids" query:"aimag_ids"`
	Ner 			string `json:"ner" form:"ner" query:"ner"`
	Taniltsuulga string `gorm:"column:taniltsuulga" json:"taniltsuulga"`
	Zurag        string `gorm:"column:zurag" json:"zurag"`

}

//Home api

func (a *Contact) TableName() string {
	return "contact"
}
func (a *Comments) TableName() string {
	return "comments"
}
func (b *Buteegdehuun) TableName() string {
	return "buteegdehuun"
}
func (h *HamtragchBaiguullaga) TableName() string {
	return "hamtragch_baiguullaga"
}
func (u *Uilchilgee) TableName() string {
	return "uilchilgee"
}
func (c *CompanyRegister) TableName() string {
	return "company_register"
}
func (n *NuurSlide) TableName() string {
	return "nuur_slide"
}
func (a *Aimag) TableName() string {
	return "aimag"
}


//Other Api

func (t *Taniltsuulga) TableName() string {
	return "taniltsuulga"
}
func (d *DSButeegdehuun) TableName() string {
	return "ds_buteegdehuun"
}
func (d *DSUilchilgee) TableName() string {
	return "ds_uilchilgee"
}

//other api

type Taniltsuulga struct {
	ID      int    `gorm:"column:id;primary_key" json:"id"`
	Tailbar string `gorm:"column:tailbar" json:"tailbar"`
}


type Zone struct {
	ID   int    `gorm:"column:id;primary_key" json:"id"`
	Zone string `gorm:"column:zone" json:"zone"`
}

func (z *Zone) TableName() string {
	return "zone"
}


type BudgBurtgel struct {
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	Ner          string `gorm:"column:ner" json:"ner"`
	Taniltsuulga string `gorm:"column:taniltsuulga" json:"taniltsuulga"`
	Zurag        string `gorm:"column:zurag" json:"zurag"`
	Zuragnuud        string `gorm:"column:zuragnuud" json:"zuragnuud"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`


}

func (b *BudgBurtgel) TableName() string {
	return "budg_burtgel"
}





type DSJuulchlal struct {
	AimagID        int        `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname      string     `gorm:"column:aimagname" json:"aimagname"`
	Buteegdehuun   string     `gorm:"column:buteegdehuun" json:"buteegdehuun"`
	ButeegdehuunID int        `gorm:"column:buteegdehuun_id" json:"buteegdehuun_id"`
	CompanyID      int        `gorm:"column:company_id" json:"company_id"`
	Utas          string     `gorm:"column:utas" json:"utas"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	ID             int        `gorm:"column:id" json:"id"`
	Namecom        string     `gorm:"column:namecom" json:"namecom"`
	Nameuilchilgee string     `gorm:"column:nameuilchilgee" json:"nameuilchilgee"`
	Ner            string     `gorm:"column:ner" json:"ner"`
	Taniltsuulga   string     `gorm:"column:taniltsuulga" json:"taniltsuulga"`
	UilchilgeeID   int        `gorm:"column:uilchilgee_id" json:"uilchilgee_id"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Zurag          string     `gorm:"column:zurag" json:"zurag"`
	Zuragnuud      string     `gorm:"column:zuragnuud" json:"zuragnuud"`
	TravelProduct []DSJuulchinButeegdehuun `gorm:"column:product" json:"product"`
	TravelSer []DSJuulchinUilchilgee `gorm:"column:uilchilgee" json:"uilchilgee"`
}

func (d *DSJuulchlal) TableName() string {
	return "ds_juulchlal"
}

type DSJuulchinButeegdehuun struct {
	ButNer     string `gorm:"column:but_ner" json:"but_ner"`
	ID         int    `gorm:"column:id" json:"id"`
	JiilchinID int    `gorm:"column:jiilchin_id" json:"jiilchin_id"`
	Medeelel   string `gorm:"column:medeelel" json:"medeelel"`
	Ner        string `gorm:"column:ner" json:"ner"`
	Utas       string `gorm:"column:utas" json:"utas"`
	Zurag          string     `gorm:"column:zurag" json:"zurag"`

}

func (d *DSJuulchinButeegdehuun) TableName() string {
	return "ds_juulchin_buteegdehuun"
}
type DSJuulchinUilchilgee struct {
	ID         int    `gorm:"column:id" json:"id"`
	JuNer      string `gorm:"column:ju_ner" json:"ju_ner"`
	JuulchinID int    `gorm:"column:juulchin_id" json:"juulchin_id"`
	Ner        string `gorm:"column:ner" json:"ner"`
	Tailbar    string `gorm:"column:tailbar" json:"tailbar"`
	Utas       string `gorm:"column:utas" json:"utas"`
	Zurag          string     `gorm:"column:zurag" json:"zurag"`

}

func (d *DSJuulchinUilchilgee) TableName() string {
	return "ds_juulchin_uilchilgee"
}


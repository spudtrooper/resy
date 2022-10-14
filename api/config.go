package api

import (
	"log"
	"time"

	"github.com/spudtrooper/goutil/request"
)

type configInfoPayload struct {
	AuthMobileCountryCodes struct {
		Voice []int `json:"voice"`
	} `json:"auth_mobile_country_codes"`
	DefaultReservationColor struct {
		Background string `json:"background"`
		Font       string `json:"font"`
	} `json:"default_reservation_color"`
	CountryCodes []struct {
		Iso3166     string `json:"iso3166"`
		Captcha     bool   `json:"captcha"`
		Code        string `json:"code"`
		CountryName string `json:"country_name"`
	} `json:"country_codes"`
	FinderDefaultLimit      int         `json:"finder_default_limit"`
	IphoneMinVersionMessage string      `json:"iphone_min_version_message"`
	LeadTimeInDays          int         `json:"lead_time_in_days"`
	RegToken                interface{} `json:"reg_token"`
	SelectAccentColor       string      `json:"select_accent_color"`
	FinderFilters           []struct {
		Name  string      `json:"name"`
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"finder_filters"`
	ServiceTypes []struct {
		ID      int    `json:"id"`
		Key     string `json:"key"`
		Value   string `json:"value"`
		Filters []struct {
			Key   string      `json:"key"`
			Value string      `json:"value"`
			Start interface{} `json:"start"`
			End   string      `json:"end"`
		} `json:"filters,omitempty"`
	} `json:"service_types"`
	Urls struct {
		Base  string `json:"base"`
		Paths struct {
			Filter string `json:"filter"`
		} `json:"paths"`
	} `json:"urls"`
	Venue struct {
		Config struct {
			AllowBypassPaymentMethod int `json:"allow_bypass_payment_method"`
			AllowMultipleResys       int `json:"allow_multiple_resys"`
			EnableInvite             int `json:"enable_invite"`
			EnableResypay            int `json:"enable_resypay"`
			HospitalityIncluded      int `json:"hospitality_included"`
		} `json:"config"`
		Contact struct {
			PhoneNumber string `json:"phone_number"`
		} `json:"contact"`
		Events       []interface{} `json:"events"`
		MaxPartySize int           `json:"max_party_size"`
		MinPartySize int           `json:"min_party_size"`
		Name         string        `json:"name"`
	} `json:"venue"`
	Whitelabel        interface{} `json:"whitelabel"`
	AndroidMinVersion string      `json:"android_min_version"`
	AndroidUpdateDate time.Time   `json:"android_update_date"`
	// Calendar          struct {
	// 	NumDays struct {
	// 		Num0     int `json:"0"`
	// 		Num222   int `json:"222"`
	// 		Num223   int `json:"223"`
	// 		Num282   int `json:"282"`
	// 		Num306   int `json:"306"`
	// 		Num339   int `json:"339"`
	// 		Num399   int `json:"399"`
	// 		Num409   int `json:"409"`
	// 		Num860   int `json:"860"`
	// 		Num861   int `json:"861"`
	// 		Num904   int `json:"904"`
	// 		Num55018 int `json:"55018"`
	// 		Num59047 int `json:"59047"`
	// 		Num05091 int `json:"05091"`
	// 		Two6B    int `json:"26b"`
	// 		Six5W    int `json:"65w"`
	// 		Abd      int `json:"abd"`
	// 		Abdh     int `json:"abdh"`
	// 		Abl      int `json:"abl"`
	// 		Abn      int `json:"abn"`
	// 		Abq      int `json:"abq"`
	// 		Abrd     int `json:"abrd"`
	// 		Abrn     int `json:"abrn"`
	// 		Abrnh    int `json:"abrnh"`
	// 		Aby      int `json:"aby"`
	// 		Acc      int `json:"acc"`
	// 		Ack      int `json:"ack"`
	// 		Acw      int `json:"acw"`
	// 		Ad5      int `json:"ad5"`
	// 		Aft      int `json:"aft"`
	// 		Ags      int `json:"ags"`
	// 		Agw      int `json:"agw"`
	// 		Ahm      int `json:"ahm"`
	// 		Ahn      int `json:"ahn"`
	// 		Aht      int `json:"aht"`
	// 		Akn      int `json:"akn"`
	// 		Alb      int `json:"alb"`
	// 		Ald      int `json:"ald"`
	// 		All      int `json:"all"`
	// 		Alm      int `json:"alm"`
	// 		Aln      int `json:"aln"`
	// 		Alp      int `json:"alp"`
	// 		Alph     int `json:"alph"`
	// 		Als      int `json:"als"`
	// 		Alt      int `json:"alt"`
	// 		Altwn    int `json:"altwn"`
	// 		Alx      int `json:"alx"`
	// 		Amb      int `json:"amb"`
	// 		Ame      int `json:"ame"`
	// 		Amg      int `json:"amg"`
	// 		Amm      int `json:"amm"`
	// 		Amst     int `json:"amst"`
	// 		Ana      int `json:"ana"`
	// 		Anc      int `json:"anc"`
	// 		Ancn     int `json:"ancn"`
	// 		And      int `json:"and"`
	// 		Ands     int `json:"ands"`
	// 		Andvr    int `json:"andvr"`
	// 		Andvrm   int `json:"andvrm"`
	// 		Ank      int `json:"ank"`
	// 		Anp      int `json:"anp"`
	// 		Antw     int `json:"antw"`
	// 		App      int `json:"app"`
	// 		Aqb      int `json:"aqb"`
	// 		Arb      int `json:"arb"`
	// 		Arc      int `json:"arc"`
	// 		Ard      int `json:"ard"`
	// 		Ardm     int `json:"ardm"`
	// 		Ardmr    int `json:"ardmr"`
	// 		Arl      int `json:"arl"`
	// 		Arlb     int `json:"arlb"`
	// 		Arln     int `json:"arln"`
	// 		Arm      int `json:"arm"`
	// 		Arn      int `json:"arn"`
	// 		Arq      int `json:"arq"`
	// 		Arr      int `json:"arr"`
	// 		Arv      int `json:"arv"`
	// 		Asc      int `json:"asc"`
	// 		Ase      int `json:"ase"`
	// 		Ash      int `json:"ash"`
	// 		Ashb     int `json:"ashb"`
	// 		Ashbr    int `json:"ashbr"`
	// 		Aspk     int `json:"aspk"`
	// 		Ast      int `json:"ast"`
	// 		Atc      int `json:"atc"`
	// 		Ath      int `json:"ath"`
	// 		Atl      int `json:"atl"`
	// 		Atln     int `json:"atln"`
	// 		Att      int `json:"att"`
	// 		Atx      int `json:"atx"`
	// 		Auo      int `json:"auo"`
	// 		Avl      int `json:"avl"`
	// 		Avla     int `json:"avla"`
	// 		Avn      int `json:"avn"`
	// 		Avnc     int `json:"avnc"`
	// 		Avr      int `json:"avr"`
	// 		Avt      int `json:"avt"`
	// 		Ayc      int `json:"ayc"`
	// 		Ayl      int `json:"ayl"`
	// 		Bal      int `json:"bal"`
	// 		Bala     int `json:"bala"`
	// 		Ban      int `json:"ban"`
	// 		Bar      int `json:"bar"`
	// 		Bbr      int `json:"bbr"`
	// 		Bch      int `json:"bch"`
	// 		Bcn      int `json:"bcn"`
	// 		Bcr      int `json:"bcr"`
	// 		Bdd      int `json:"bdd"`
	// 		Bdf      int `json:"bdf"`
	// 		Bdfr     int `json:"bdfr"`
	// 		Bdj      int `json:"bdj"`
	// 		Bdl      int `json:"bdl"`
	// 		Bdr      int `json:"bdr"`
	// 		Bel      int `json:"bel"`
	// 		Bfd      int `json:"bfd"`
	// 		Bff      int `json:"bff"`
	// 		Bfg      int `json:"bfg"`
	// 		Bfr      int `json:"bfr"`
	// 		Bgm      int `json:"bgm"`
	// 		Bgn      int `json:"bgn"`
	// 		Bgs      int `json:"bgs"`
	// 		Bgt      int `json:"bgt"`
	// 		Bgtc     int `json:"bgtc"`
	// 		Bhb      int `json:"bhb"`
	// 		Bhm      int `json:"bhm"`
	// 		Bhr      int `json:"bhr"`
	// 		Bid      int `json:"bid"`
	// 		Bil      int `json:"bil"`
	// 		Bkd      int `json:"bkd"`
	// 		Bkk      int `json:"bkk"`
	// 		Blc      int `json:"blc"`
	// 		Blck     int `json:"blck"`
	// 		Bld      int `json:"bld"`
	// 		Bldn     int `json:"bldn"`
	// 		Bldw     int `json:"bldw"`
	// 		Blf      int `json:"blf"`
	// 		Blg      int `json:"blg"`
	// 		Blgn     int `json:"blgn"`
	// 		Blh      int `json:"blh"`
	// 		Bll      int `json:"bll"`
	// 		Bllm     int `json:"bllm"`
	// 		Blln     int `json:"blln"`
	// 		Bllns    int `json:"bllns"`
	// 		Bllrt    int `json:"bllrt"`
	// 		Bllv     int `json:"bllv"`
	// 		Bllw     int `json:"bllw"`
	// 		Blm      int `json:"blm"`
	// 		Blmf     int `json:"blmf"`
	// 		Blmn     int `json:"blmn"`
	// 		Blr      int `json:"blr"`
	// 		Blrd     int `json:"blrd"`
	// 		Blrm     int `json:"blrm"`
	// 		Blrmd    int `json:"blrmd"`
	// 		Blrs     int `json:"blrs"`
	// 		Bmh      int `json:"bmh"`
	// 		Bna      int `json:"bna"`
	// 		Bnd      int `json:"bnd"`
	// 		Bndb     int `json:"bndb"`
	// 		Bng      int `json:"bng"`
	// 		Bngl     int `json:"bngl"`
	// 		Bnglw    int `json:"bnglw"`
	// 		Bngr     int `json:"bngr"`
	// 		Bnn      int `json:"bnn"`
	// 		Bnt      int `json:"bnt"`
	// 		Boi      int `json:"boi"`
	// 		Bol      int `json:"bol"`
	// 		Bos      int `json:"bos"`
	// 		Boz      int `json:"boz"`
	// 		Brb      int `json:"brb"`
	// 		Brd      int `json:"brd"`
	// 		Brdf     int `json:"brdf"`
	// 		Brdg     int `json:"brdg"`
	// 		Brdl     int `json:"brdl"`
	// 		Brdn     int `json:"brdn"`
	// 		Brdp     int `json:"brdp"`
	// 		Brds     int `json:"brds"`
	// 		Bre      int `json:"bre"`
	// 		Brg      int `json:"brg"`
	// 		Brgn     int `json:"brgn"`
	// 		Bri      int `json:"bri"`
	// 		Brig     int `json:"brig"`
	// 		Brk      int `json:"brk"`
	// 		Brkh     int `json:"brkh"`
	// 		Brkl     int `json:"brkl"`
	// 		Brl      int `json:"brl"`
	// 		Brln     int `json:"brln"`
	// 		Brlng    int `json:"brlng"`
	// 		Brm      int `json:"brm"`
	// 		Brml     int `json:"brml"`
	// 		Brn      int `json:"brn"`
	// 		Brnf     int `json:"brnf"`
	// 		Brnr     int `json:"brnr"`
	// 		Brns     int `json:"brns"`
	// 		Brnsw    int `json:"brnsw"`
	// 		Brnt     int `json:"brnt"`
	// 		Brr      int `json:"brr"`
	// 		Brrn     int `json:"brrn"`
	// 		Brrr     int `json:"brrr"`
	// 		Brs      int `json:"brs"`
	// 		Brsb     int `json:"brsb"`
	// 		Brsl     int `json:"brsl"`
	// 		Brss     int `json:"brss"`
	// 		Brst     int `json:"brst"`
	// 		Brstl    int `json:"brstl"`
	// 		Brstlr   int `json:"brstlr"`
	// 		Brstlv   int `json:"brstlv"`
	// 		Brt      int `json:"brt"`
	// 		Brtn     int `json:"brtn"`
	// 		Brtt     int `json:"brtt"`
	// 		Brv      int `json:"brv"`
	// 		Brw      int `json:"brw"`
	// 		Brwn     int `json:"brwn"`
	// 		Brwy     int `json:"brwy"`
	// 		Bry      int `json:"bry"`
	// 		Bsh      int `json:"bsh"`
	// 		Bsm      int `json:"bsm"`
	// 		Bss      int `json:"bss"`
	// 		Bst      int `json:"bst"`
	// 		Bth      int `json:"bth"`
	// 		Bthb     int `json:"bthb"`
	// 		Bthl     int `json:"bthl"`
	// 		Bthn     int `json:"bthn"`
	// 		Bths     int `json:"bths"`
	// 		Btr      int `json:"btr"`
	// 		Btt      int `json:"btt"`
	// 		Buf      int `json:"buf"`
	// 		Bur      int `json:"bur"`
	// 		Bvn      int `json:"bvn"`
	// 		Bvr      int `json:"bvr"`
	// 		Bvrc     int `json:"bvrc"`
	// 		Bvrl     int `json:"bvrl"`
	// 		Bvt      int `json:"bvt"`
	// 		Bwm      int `json:"bwm"`
	// 		Bwn      int `json:"bwn"`
	// 		Byc      int `json:"byc"`
	// 		Byhr     int `json:"byhr"`
	// 		Byn      int `json:"byn"`
	// 		Byr      int `json:"byr"`
	// 		Byrn     int `json:"byrn"`
	// 		Bys      int `json:"bys"`
	// 		Byv      int `json:"byv"`
	// 		Cadz     int `json:"cadz"`
	// 		Cal      int `json:"cal"`
	// 		Cba      int `json:"cba"`
	// 		Cbh      int `json:"cbh"`
	// 		Cbr      int `json:"cbr"`
	// 		Ccb      int `json:"ccb"`
	// 		Ccr      int `json:"ccr"`
	// 		Cda      int `json:"cda"`
	// 		Cdd      int `json:"cdd"`
	// 		Cdr      int `json:"cdr"`
	// 		Cdrf     int `json:"cdrf"`
	// 		Cdrh     int `json:"cdrh"`
	// 		Cdz      int `json:"cdz"`
	// 		Cgksdq   int `json:"cgksdq"`
	// 		Chc      int `json:"chc"`
	// 		Chcc     int `json:"chcc"`
	// 		Chcl     int `json:"chcl"`
	// 		Chd      int `json:"chd"`
	// 		Che      int `json:"che"`
	// 		Chey     int `json:"chey"`
	// 		Chg      int `json:"chg"`
	// 		Chh      int `json:"chh"`
	// 		Chi      int `json:"chi"`
	// 		Chk      int `json:"chk"`
	// 		Chln     int `json:"chln"`
	// 		Chls     int `json:"chls"`
	// 		Chlt     int `json:"chlt"`
	// 		Chm      int `json:"chm"`
	// 		Chn      int `json:"chn"`
	// 		Chnd     int `json:"chnd"`
	// 		Chnt     int `json:"chnt"`
	// 		Cho      int `json:"cho"`
	// 		Chrc     int `json:"chrc"`
	// 		Chrl     int `json:"chrl"`
	// 		Chrls    int `json:"chrls"`
	// 		Chrlt    int `json:"chrlt"`
	// 		Chrr     int `json:"chrr"`
	// 		Chrs     int `json:"chrs"`
	// 		Chs      int `json:"chs"`
	// 		Chst     int `json:"chst"`
	// 		Chstr    int `json:"chstr"`
	// 		Chstrn   int `json:"chstrn"`
	// 		Chstrp   int `json:"chstrp"`
	// 		Cht      int `json:"cht"`
	// 		Chth     int `json:"chth"`
	// 		Chv      int `json:"chv"`
	// 		Cid      int `json:"cid"`
	// 		Cin      int `json:"cin"`
	// 		Ciri     int `json:"ciri"`
	// 		Ckh      int `json:"ckh"`
	// 		Clb      int `json:"clb"`
	// 		Clc      int `json:"clc"`
	// 		Clds     int `json:"clds"`
	// 		Cldst    int `json:"cldst"`
	// 		Cldstr   int `json:"cldstr"`
	// 		Cldw     int `json:"cldw"`
	// 		Cle      int `json:"cle"`
	// 		Clf      int `json:"clf"`
	// 		Clft     int `json:"clft"`
	// 		Clgn     int `json:"clgn"`
	// 		Cllc     int `json:"cllc"`
	// 		Cllcn    int `json:"cllcn"`
	// 		Cllg     int `json:"cllg"`
	// 		Cllgp    int `json:"cllgp"`
	// 		Clln     int `json:"clln"`
	// 		Clm      int `json:"clm"`
	// 		Clmb     int `json:"clmb"`
	// 		Cln      int `json:"cln"`
	// 		Clr      int `json:"clr"`
	// 		Clrm     int `json:"clrm"`
	// 		Clrmn    int `json:"clrmn"`
	// 		Clrmnt   int `json:"clrmnt"`
	// 		Clrn     int `json:"clrn"`
	// 		Clrnc    int `json:"clrnc"`
	// 		Cls      int `json:"cls"`
	// 		Clt      int `json:"clt"`
	// 		Clv      int `json:"clv"`
	// 		Clvl     int `json:"clvl"`
	// 		Clvr     int `json:"clvr"`
	// 		Clvrs    int `json:"clvrs"`
	// 		Clvs     int `json:"clvs"`
	// 		Clw      int `json:"clw"`
	// 		Clws     int `json:"clws"`
	// 		Cly      int `json:"cly"`
	// 		Clyt     int `json:"clyt"`
	// 		Cmb      int `json:"cmb"`
	// 		Cmbrd    int `json:"cmbrd"`
	// 		Cmd      int `json:"cmd"`
	// 		Cmf      int `json:"cmf"`
	// 		Cmp      int `json:"cmp"`
	// 		Cmpb     int `json:"cmpb"`
	// 		Cmph     int `json:"cmph"`
	// 		Cnc      int `json:"cnc"`
	// 		Cncr     int `json:"cncr"`
	// 		Cncrd    int `json:"cncrd"`
	// 		Cnh      int `json:"cnh"`
	// 		Cns      int `json:"cns"`
	// 		Cnt      int `json:"cnt"`
	// 		Cntn     int `json:"cntn"`
	// 		Cntnm    int `json:"cntnm"`
	// 		Cntnms   int `json:"cntnms"`
	// 		Cntr     int `json:"cntr"`
	// 		Cntrb    int `json:"cntrb"`
	// 		Cntrbr   int `json:"cntrbr"`
	// 		Cntrc    int `json:"cntrc"`
	// 		Cntrl    int `json:"cntrl"`
	// 		Cntrlm   int `json:"cntrlm"`
	// 		Cntrlv   int `json:"cntrlv"`
	// 		Cntrv    int `json:"cntrv"`
	// 		Cnv      int `json:"cnv"`
	// 		Cny      int `json:"cny"`
	// 		Cnyr     int `json:"cnyr"`
	// 		Col      int `json:"col"`
	// 		Com      int `json:"com"`
	// 		Con      int `json:"con"`
	// 		Cos      int `json:"cos"`
	// 		Cpc      int `json:"cpc"`
	// 		Cpm      int `json:"cpm"`
	// 		Cpn      int `json:"cpn"`
	// 		Cpr      int `json:"cpr"`
	// 		Crb      int `json:"crb"`
	// 		Crck     int `json:"crck"`
	// 		Crd      int `json:"crd"`
	// 		Crk      int `json:"crk"`
	// 		Crl      int `json:"crl"`
	// 		Crlt     int `json:"crlt"`
	// 		Crm      int `json:"crm"`
	// 		Crn      int `json:"crn"`
	// 		Crnc     int `json:"crnc"`
	// 		Crnf     int `json:"crnf"`
	// 		Crnl     int `json:"crnl"`
	// 		Crnls    int `json:"crnls"`
	// 		Crns     int `json:"crns"`
	// 		Crnw     int `json:"crnw"`
	// 		Crp      int `json:"crp"`
	// 		Crpn     int `json:"crpn"`
	// 		Crs      int `json:"crs"`
	// 		Crt      int `json:"crt"`
	// 		Crtg     int `json:"crtg"`
	// 		Crw      int `json:"crw"`
	// 		Crwf     int `json:"crwf"`
	// 		Cry      int `json:"cry"`
	// 		Cryd     int `json:"cryd"`
	// 		Crys     int `json:"crys"`
	// 		Crz      int `json:"crz"`
	// 		Csc      int `json:"csc"`
	// 		Csh      int `json:"csh"`
	// 		Csp      int `json:"csp"`
	// 		Cst      int `json:"cst"`
	// 		Cstl     int `json:"cstl"`
	// 		Cstlh    int `json:"cstlh"`
	// 		Cstm     int `json:"cstm"`
	// 		Cstr     int `json:"cstr"`
	// 		Cth      int `json:"cth"`
	// 		Cts      int `json:"cts"`
	// 		Ctt      int `json:"ctt"`
	// 		Ctw      int `json:"ctw"`
	// 		Cul      int `json:"cul"`
	// 		Cvn      int `json:"cvn"`
	// 		Cvnt     int `json:"cvnt"`
	// 		Cwd      int `json:"cwd"`
	// 		Cyd      int `json:"cyd"`
	// 		Cyp      int `json:"cyp"`
	// 		Czn      int `json:"czn"`
	// 		Dav      int `json:"dav"`
	// 		Dbl      int `json:"dbl"`
	// 		Dbn      int `json:"dbn"`
	// 		Dbnt     int `json:"dbnt"`
	// 		Dbntd    int `json:"dbntd"`
	// 		Dc       int `json:"dc"`
	// 		Dca      int `json:"dca"`
	// 		Ddg      int `json:"ddg"`
	// 		Dem      int `json:"dem"`
	// 		Den      int `json:"den"`
	// 		Det      int `json:"det"`
	// 		Dfw      int `json:"dfw"`
	// 		Dga      int `json:"dga"`
	// 		Dgl      int `json:"dgl"`
	// 		Dhq      int `json:"dhq"`
	// 		Dis      int `json:"dis"`
	// 		Dkofnh   int `json:"dkofnh"`
	// 		Dlc      int `json:"dlc"`
	// 		Dlh      int `json:"dlh"`
	// 		Dlm      int `json:"dlm"`
	// 		Dlr      int `json:"dlr"`
	// 		Dlt      int `json:"dlt"`
	// 		Dmn      int `json:"dmn"`
	// 		Dnb      int `json:"dnb"`
	// 		Dnd      int `json:"dnd"`
	// 		Dndr     int `json:"dndr"`
	// 		Dnk      int `json:"dnk"`
	// 		Dnp      int `json:"dnp"`
	// 		Dnv      int `json:"dnv"`
	// 		Dov      int `json:"dov"`
	// 		Drby     int `json:"drby"`
	// 		Drg      int `json:"drg"`
	// 		Drl      int `json:"drl"`
	// 		Drn      int `json:"drn"`
	// 		Dro      int `json:"dro"`
	// 		Dsm      int `json:"dsm"`
	// 		Dsr      int `json:"dsr"`
	// 		Dss      int `json:"dss"`
	// 		Dst      int `json:"dst"`
	// 		Dtn      int `json:"dtn"`
	// 		Dub      int `json:"dub"`
	// 		Dvl      int `json:"dvl"`
	// 		Dvn      int `json:"dvn"`
	// 		Dvr      int `json:"dvr"`
	// 		Dwn      int `json:"dwn"`
	// 		Dwnn     int `json:"dwnn"`
	// 		Dwnr     int `json:"dwnr"`
	// 		Dxb      int `json:"dxb"`
	// 		Dxr      int `json:"dxr"`
	// 		Dyt      int `json:"dyt"`
	// 		Dytn     int `json:"dytn"`
	// 		Ecl      int `json:"ecl"`
	// 		Ecp      int `json:"ecp"`
	// 		Edg      int `json:"edg"`
	// 		Edgc     int `json:"edgc"`
	// 		Edgr     int `json:"edgr"`
	// 		Edm      int `json:"edm"`
	// 		Edmn     int `json:"edmn"`
	// 		Edmnd    int `json:"edmnd"`
	// 		Edn      int `json:"edn"`
	// 		Ednb     int `json:"ednb"`
	// 		Eds      int `json:"eds"`
	// 		Egl      int `json:"egl"`
	// 		Egn      int `json:"egn"`
	// 		Ell      int `json:"ell"`
	// 		Ellc     int `json:"ellc"`
	// 		Elm      int `json:"elm"`
	// 		Elmw     int `json:"elmw"`
	// 		Elp      int `json:"elp"`
	// 		Els      int `json:"els"`
	// 		Emr      int `json:"emr"`
	// 		Enw      int `json:"enw"`
	// 		Erc      int `json:"erc"`
	// 		Ess      int `json:"ess"`
	// 		Est      int `json:"est"`
	// 		Estb     int `json:"estb"`
	// 		Estbt    int `json:"estbt"`
	// 		Estf     int `json:"estf"`
	// 		Estfr    int `json:"estfr"`
	// 		Estg     int `json:"estg"`
	// 		Estm     int `json:"estm"`
	// 		Estn     int `json:"estn"`
	// 		Estr     int `json:"estr"`
	// 		Ests     int `json:"ests"`
	// 		Etp      int `json:"etp"`
	// 		Evl      int `json:"evl"`
	// 		Evlg     int `json:"evlg"`
	// 		Evns     int `json:"evns"`
	// 		Evr      int `json:"evr"`
	// 		Evrt     int `json:"evrt"`
	// 		Evs      int `json:"evs"`
	// 		Ewr      int `json:"ewr"`
	// 		Exc      int `json:"exc"`
	// 		Ext      int `json:"ext"`
	// 		Extn     int `json:"extn"`
	// 		Eyw      int `json:"eyw"`
	// 		Fax      int `json:"fax"`
	// 		Fcs      int `json:"fcs"`
	// 		Fdb      int `json:"fdb"`
	// 		Fisffg   int `json:"fisffg"`
	// 		Fjrd     int `json:"fjrd"`
	// 		Fkk      int `json:"fkk"`
	// 		Fkn      int `json:"fkn"`
	// 		Flf      int `json:"flf"`
	// 		Flg      int `json:"flg"`
	// 		Fll      int `json:"fll"`
	// 		Flm      int `json:"flm"`
	// 		Flmn     int `json:"flmn"`
	// 		Flms     int `json:"flms"`
	// 		Flmt     int `json:"flmt"`
	// 		Fln      int `json:"fln"`
	// 		Flnd     int `json:"flnd"`
	// 		Flr      int `json:"flr"`
	// 		Flrn     int `json:"flrn"`
	// 		Fls      int `json:"fls"`
	// 		Flt      int `json:"flt"`
	// 		Flw      int `json:"flw"`
	// 		Flwd     int `json:"flwd"`
	// 		Fmth     int `json:"fmth"`
	// 		Fnc      int `json:"fnc"`
	// 		Fng      int `json:"fng"`
	// 		Fnn      int `json:"fnn"`
	// 		Fra      int `json:"fra"`
	// 		Frb      int `json:"frb"`
	// 		Frbr     int `json:"frbr"`
	// 		Frd      int `json:"frd"`
	// 		Frdb     int `json:"frdb"`
	// 		Fre      int `json:"fre"`
	// 		Frf      int `json:"frf"`
	// 		Frfld    int `json:"frfld"`
	// 		Frfldn   int `json:"frfldn"`
	// 		Frn      int `json:"frn"`
	// 		Frnd     int `json:"frnd"`
	// 		Frnh     int `json:"frnh"`
	// 		Frnk     int `json:"frnk"`
	// 		Frnkl    int `json:"frnkl"`
	// 		Frnt     int `json:"frnt"`
	// 		Frp      int `json:"frp"`
	// 		Frprt    int `json:"frprt"`
	// 		Frs      int `json:"frs"`
	// 		Frsc     int `json:"frsc"`
	// 		Frst     int `json:"frst"`
	// 		Frstr    int `json:"frstr"`
	// 		Frsy     int `json:"frsy"`
	// 		Frt      int `json:"frt"`
	// 		Frtl     int `json:"frtl"`
	// 		Frts     int `json:"frts"`
	// 		Frtw     int `json:"frtw"`
	// 		Frv      int `json:"frv"`
	// 		Fsh      int `json:"fsh"`
	// 		Fth      int `json:"fth"`
	// 		Ftm      int `json:"ftm"`
	// 		Ftv      int `json:"ftv"`
	// 		Fxb      int `json:"fxb"`
	// 		Fyt      int `json:"fyt"`
	// 		Fytt     int `json:"fytt"`
	// 		Gbs      int `json:"gbs"`
	// 		Gby      int `json:"gby"`
	// 		Gcm      int `json:"gcm"`
	// 		Gdi      int `json:"gdi"`
	// 		Gdl      int `json:"gdl"`
	// 		Gdlj     int `json:"gdlj"`
	// 		Ggh      int `json:"ggh"`
	// 		Ghn      int `json:"ghn"`
	// 		Ghz      int `json:"ghz"`
	// 		Glb      int `json:"glb"`
	// 		Glc      int `json:"glc"`
	// 		Glcs     int `json:"glcs"`
	// 		Gld      int `json:"gld"`
	// 		Gldc     int `json:"gldc"`
	// 		Gldn     int `json:"gldn"`
	// 		Glf      int `json:"glf"`
	// 		Gllt     int `json:"gllt"`
	// 		Gllw     int `json:"gllw"`
	// 		Gln      int `json:"gln"`
	// 		Glnc     int `json:"glnc"`
	// 		Glndn    int `json:"glndn"`
	// 		Glnl     int `json:"glnl"`
	// 		Glnll    int `json:"glnll"`
	// 		Glnm     int `json:"glnm"`
	// 		Glns     int `json:"glns"`
	// 		Glo      int `json:"glo"`
	// 		Glr      int `json:"glr"`
	// 		Gls      int `json:"gls"`
	// 		Glsg     int `json:"glsg"`
	// 		Glst     int `json:"glst"`
	// 		Glt      int `json:"glt"`
	// 		Gnb      int `json:"gnb"`
	// 		Gnd      int `json:"gnd"`
	// 		Gns      int `json:"gns"`
	// 		Gnv      int `json:"gnv"`
	// 		Gnvn     int `json:"gnvn"`
	// 		Goo      int `json:"goo"`
	// 		Gpz      int `json:"gpz"`
	// 		Gr       int `json:"gr"`
	// 		Grd      int `json:"grd"`
	// 		Grdn     int `json:"grdn"`
	// 		Grg      int `json:"grg"`
	// 		Grgr     int `json:"grgr"`
	// 		Grl      int `json:"grl"`
	// 		Grm      int `json:"grm"`
	// 		Grmn     int `json:"grmn"`
	// 		Grmnt    int `json:"grmnt"`
	// 		Grn      int `json:"grn"`
	// 		Grna     int `json:"grna"`
	// 		Grnd     int `json:"grnd"`
	// 		Grndh    int `json:"grndh"`
	// 		Grndl    int `json:"grndl"`
	// 		Grndp    int `json:"grndp"`
	// 		Grng     int `json:"grng"`
	// 		Grnl     int `json:"grnl"`
	// 		Grnp     int `json:"grnp"`
	// 		Grnvl    int `json:"grnvl"`
	// 		Grnvll   int `json:"grnvll"`
	// 		Grnw     int `json:"grnw"`
	// 		Grp      int `json:"grp"`
	// 		Grr      int `json:"grr"`
	// 		Grrt     int `json:"grrt"`
	// 		Grs      int `json:"grs"`
	// 		Grt      int `json:"grt"`
	// 		Grz      int `json:"grz"`
	// 		Gsh      int `json:"gsh"`
	// 		Gso      int `json:"gso"`
	// 		Gst      int `json:"gst"`
	// 		Gtb      int `json:"gtb"`
	// 		Gvl      int `json:"gvl"`
	// 		Gwh      int `json:"gwh"`
	// 		Gwy      int `json:"gwy"`
	// 		Gys      int `json:"gys"`
	// 		Har      int `json:"har"`
	// 		Hbb      int `json:"hbb"`
	// 		Hck      int `json:"hck"`
	// 		Hdd      int `json:"hdd"`
	// 		Hds      int `json:"hds"`
	// 		Hdsn     int `json:"hdsn"`
	// 		Hdsnm    int `json:"hdsnm"`
	// 		Hgh      int `json:"hgh"`
	// 		Hghln    int `json:"hghln"`
	// 		Hghw     int `json:"hghw"`
	// 		Hgr      int `json:"hgr"`
	// 		Hgw      int `json:"hgw"`
	// 		Hhs      int `json:"hhs"`
	// 		Hil      int `json:"hil"`
	// 		Hkg      int `json:"hkg"`
	// 		Hlb      int `json:"hlb"`
	// 		Hld      int `json:"hld"`
	// 		Hlh      int `json:"hlh"`
	// 		Hll      int `json:"hll"`
	// 		Hlln     int `json:"hlln"`
	// 		Hlls     int `json:"hlls"`
	// 		Hlly     int `json:"hlly"`
	// 		Hlm      int `json:"hlm"`
	// 		Hlms     int `json:"hlms"`
	// 		Hlv      int `json:"hlv"`
	// 		Hlw      int `json:"hlw"`
	// 		Hmb      int `json:"hmb"`
	// 		Hmbr     int `json:"hmbr"`
	// 		Hmca     int `json:"hmca"`
	// 		Hmd      int `json:"hmd"`
	// 		Hml      int `json:"hml"`
	// 		Hmlt     int `json:"hmlt"`
	// 		Hmp      int `json:"hmp"`
	// 		Hmpt     int `json:"hmpt"`
	// 		Hnb      int `json:"hnb"`
	// 		Hnbr     int `json:"hnbr"`
	// 		Hnc      int `json:"hnc"`
	// 		Hnd      int `json:"hnd"`
	// 		Hndr     int `json:"hndr"`
	// 		Hndrs    int `json:"hndrs"`
	// 		Hng      int `json:"hng"`
	// 		Hnl      int `json:"hnl"`
	// 		Hnly     int `json:"hnly"`
	// 		Hns      int `json:"hns"`
	// 		Hnt      int `json:"hnt"`
	// 		Hntb     int `json:"hntb"`
	// 		Hntn     int `json:"hntn"`
	// 		Hntng    int `json:"hntng"`
	// 		Hntr     int `json:"hntr"`
	// 		Hnts     int `json:"hnts"`
	// 		Hnv      int `json:"hnv"`
	// 		Hob      int `json:"hob"`
	// 		Hol      int `json:"hol"`
	// 		Hou      int `json:"hou"`
	// 		Hpk      int `json:"hpk"`
	// 		Hpn      int `json:"hpn"`
	// 		Hpw      int `json:"hpw"`
	// 		Hrbr     int `json:"hrbr"`
	// 		Hrl      int `json:"hrl"`
	// 		Hrp      int `json:"hrp"`
	// 		Hrr      int `json:"hrr"`
	// 		Hrrs     int `json:"hrrs"`
	// 		Hrrsn    int `json:"hrrsn"`
	// 		Hrs      int `json:"hrs"`
	// 		Hrsh     int `json:"hrsh"`
	// 		Hrts     int `json:"hrts"`
	// 		Hrv      int `json:"hrv"`
	// 		Hrw      int `json:"hrw"`
	// 		Hsc      int `json:"hsc"`
	// 		Hst      int `json:"hst"`
	// 		Hsv      int `json:"hsv"`
	// 		Hsy      int `json:"hsy"`
	// 		Hts      int `json:"hts"`
	// 		Htsp     int `json:"htsp"`
	// 		Hud      int `json:"hud"`
	// 		Hvk      int `json:"hvk"`
	// 		Hvn      int `json:"hvn"`
	// 		Hwd      int `json:"hwd"`
	// 		Hwk      int `json:"hwk"`
	// 		Hwl      int `json:"hwl"`
	// 		Hwll     int `json:"hwll"`
	// 		Hyd      int `json:"hyd"`
	// 		Ibz      int `json:"ibz"`
	// 		Ict      int `json:"ict"`
	// 		Idh      int `json:"idh"`
	// 		Ikv      int `json:"ikv"`
	// 		Inc      int `json:"inc"`
	// 		Ind      int `json:"ind"`
	// 		Ing      int `json:"ing"`
	// 		Inn      int `json:"inn"`
	// 		Inns     int `json:"inns"`
	// 		Invr     int `json:"invr"`
	// 		Irv      int `json:"irv"`
	// 		Irvn     int `json:"irvn"`
	// 		Irvng    int `json:"irvng"`
	// 		Irvngt   int `json:"irvngt"`
	// 		Isl      int `json:"isl"`
	// 		Ist      int `json:"ist"`
	// 		Ith      int `json:"ith"`
	// 		Iwc      int `json:"iwc"`
	// 		Izm      int `json:"izm"`
	// 		Jan      int `json:"jan"`
	// 		Jax      int `json:"jax"`
	// 		Jck      int `json:"jck"`
	// 		Jcks     int `json:"jcks"`
	// 		Jdd      int `json:"jdd"`
	// 		Jddh     int `json:"jddh"`
	// 		Jff      int `json:"jff"`
	// 		Jhn      int `json:"jhn"`
	// 		Jhns     int `json:"jhns"`
	// 		Jhnss    int `json:"jhnss"`
	// 		Jky      int `json:"jky"`
	// 		Jn       int `json:"jn"`
	// 		Jnn      int `json:"jnn"`
	// 		Jns      int `json:"jns"`
	// 		Jpt      int `json:"jpt"`
	// 		Jsc      int `json:"jsc"`
	// 		Kal      int `json:"kal"`
	// 		Kc       int `json:"kc"`
	// 		Khh      int `json:"khh"`
	// 		Khl      int `json:"khl"`
	// 		Kingst   int `json:"kingst"`
	// 		Kld      int `json:"kld"`
	// 		Klg      int `json:"klg"`
	// 		Klh      int `json:"klh"`
	// 		Klk      int `json:"klk"`
	// 		Kll      int `json:"kll"`
	// 		Kllmc    int `json:"kllmc"`
	// 		Klln     int `json:"klln"`
	// 		Kln      int `json:"kln"`
	// 		Klnm     int `json:"klnm"`
	// 		Kls      int `json:"kls"`
	// 		Kmb      int `json:"kmb"`
	// 		Knd      int `json:"knd"`
	// 		Kndr     int `json:"kndr"`
	// 		Kng      int `json:"kng"`
	// 		Kngf     int `json:"kngf"`
	// 		Kngs     int `json:"kngs"`
	// 		Knn      int `json:"knn"`
	// 		Knnbn    int `json:"knnbn"`
	// 		Knnp     int `json:"knnp"`
	// 		Knnt     int `json:"knnt"`
	// 		Kns      int `json:"kns"`
	// 		Knth     int `json:"knth"`
	// 		Knx      int `json:"knx"`
	// 		Kpl      int `json:"kpl"`
	// 		Krk      int `json:"krk"`
	// 		Krn      int `json:"krn"`
	// 		Ksw      int `json:"ksw"`
	// 		Ktc      int `json:"ktc"`
	// 		Ktt      int `json:"ktt"`
	// 		Kty      int `json:"kty"`
	// 		Kwh      int `json:"kwh"`
	// 		Kyp      int `json:"kyp"`
	// 		Kyt      int `json:"kyt"`
	// 		LR       int `json:"l'r"`
	// 		La       int `json:"la"`
	// 		Lag      int `json:"lag"`
	// 		Lan      int `json:"lan"`
	// 		Lar      int `json:"lar"`
	// 		Las      int `json:"las"`
	// 		Lax      int `json:"lax"`
	// 		Lbc      int `json:"lbc"`
	// 		Lbk      int `json:"lbk"`
	// 		Lbn      int `json:"lbn"`
	// 		Lbnn     int `json:"lbnn"`
	// 		Lca      int `json:"lca"`
	// 		Lcb      int `json:"lcb"`
	// 		Lck      int `json:"lck"`
	// 		Lckw     int `json:"lckw"`
	// 		Lcn      int `json:"lcn"`
	// 		Lcnn     int `json:"lcnn"`
	// 		Lcnt     int `json:"lcnt"`
	// 		Ldh      int `json:"ldh"`
	// 		Ldl      int `json:"ldl"`
	// 		Ldlw     int `json:"ldlw"`
	// 		Ldn      int `json:"ldn"`
	// 		Lds      int `json:"lds"`
	// 		Lem      int `json:"lem"`
	// 		Leo      int `json:"leo"`
	// 		Lex      int `json:"lex"`
	// 		Lft      int `json:"lft"`
	// 		Lg       int `json:"lg"`
	// 		Lgb      int `json:"lgb"`
	// 		Lgbl     int `json:"lgbl"`
	// 		Lgh      int `json:"lgh"`
	// 		Lghn     int `json:"lghn"`
	// 		Lgn      int `json:"lgn"`
	// 		Lgnn     int `json:"lgnn"`
	// 		Lgr      int `json:"lgr"`
	// 		Lgrn     int `json:"lgrn"`
	// 		Lhn      int `json:"lhn"`
	// 		Li       int `json:"li"`
	// 		Lis      int `json:"lis"`
	// 		Lja      int `json:"lja"`
	// 		Lkb      int `json:"lkb"`
	// 		Lkf      int `json:"lkf"`
	// 		Lkh      int `json:"lkh"`
	// 		Lkl      int `json:"lkl"`
	// 		Lkp      int `json:"lkp"`
	// 		Lks      int `json:"lks"`
	// 		Lkt      int `json:"lkt"`
	// 		Lkv      int `json:"lkv"`
	// 		Lkvl     int `json:"lkvl"`
	// 		Lkw      int `json:"lkw"`
	// 		Lkwd     int `json:"lkwd"`
	// 		Lkwr     int `json:"lkwr"`
	// 		Lkz      int `json:"lkz"`
	// 		Lll      int `json:"lll"`
	// 		Lln      int `json:"lln"`
	// 		Lm       int `json:"lm"`
	// 		Lmb      int `json:"lmb"`
	// 		Lmbrgb   int `json:"lmbrgb"`
	// 		Lmbrt    int `json:"lmbrt"`
	// 		Lmd      int `json:"lmd"`
	// 		Lmp      int `json:"lmp"`
	// 		Lmr      int `json:"lmr"`
	// 		Lms      int `json:"lms"`
	// 		Lmw      int `json:"lmw"`
	// 		Lncv     int `json:"lncv"`
	// 		Lnd      int `json:"lnd"`
	// 		Lndn     int `json:"lndn"`
	// 		Lnds     int `json:"lnds"`
	// 		Lng      int `json:"lng"`
	// 		Lngb     int `json:"lngb"`
	// 		Lngbch   int `json:"lngbch"`
	// 		Lngbr    int `json:"lngbr"`
	// 		Lngbt    int `json:"lngbt"`
	// 		Lngmn    int `json:"lngmn"`
	// 		Lngp     int `json:"lngp"`
	// 		Lngs     int `json:"lngs"`
	// 		Lnn      int `json:"lnn"`
	// 		Lns      int `json:"lns"`
	// 		Lnsd     int `json:"lnsd"`
	// 		Lnt      int `json:"lnt"`
	// 		Lnx      int `json:"lnx"`
	// 		Lnz      int `json:"lnz"`
	// 		Los      int `json:"los"`
	// 		Lou      int `json:"lou"`
	// 		Low      int `json:"low"`
	// 		Lpl      int `json:"lpl"`
	// 		Lrd      int `json:"lrd"`
	// 		Lrj      int `json:"lrj"`
	// 		Lrm      int `json:"lrm"`
	// 		Lrt      int `json:"lrt"`
	// 		Lrv      int `json:"lrv"`
	// 		Lsb      int `json:"lsb"`
	// 		Lsbn     int `json:"lsbn"`
	// 		Lsbr     int `json:"lsbr"`
	// 		Lsg      int `json:"lsg"`
	// 		Lsl      int `json:"lsl"`
	// 		Lsln     int `json:"lsln"`
	// 		Lsp      int `json:"lsp"`
	// 		Lsv      int `json:"lsv"`
	// 		Lth      int `json:"lth"`
	// 		Ltm      int `json:"ltm"`
	// 		Ltt      int `json:"ltt"`
	// 		Lttlt    int `json:"lttlt"`
	// 		Lttz     int `json:"lttz"`
	// 		Lv       int `json:"lv"`
	// 		Lva      int `json:"lva"`
	// 		Lvd      int `json:"lvd"`
	// 		Lvl      int `json:"lvl"`
	// 		Lvm      int `json:"lvm"`
	// 		Lvn      int `json:"lvn"`
	// 		Lvng     int `json:"lvng"`
	// 		Lvr      int `json:"lvr"`
	// 		Lwd      int `json:"lwd"`
	// 		Lwr      int `json:"lwr"`
	// 		Lws      int `json:"lws"`
	// 		Lwsk     int `json:"lwsk"`
	// 		Lwst     int `json:"lwst"`
	// 		Lwstw    int `json:"lwstw"`
	// 		Lwsv     int `json:"lwsv"`
	// 		Lxm      int `json:"lxm"`
	// 		Lxmb     int `json:"lxmb"`
	// 		Lxn      int `json:"lxn"`
	// 		Lxng     int `json:"lxng"`
	// 		Lyh      int `json:"lyh"`
	// 		Lym      int `json:"lym"`
	// 		Lymr     int `json:"lymr"`
	// 		Lync     int `json:"lync"`
	// 		Lyt      int `json:"lyt"`
	// 		Lytn     int `json:"lytn"`
	// 		Mad      int `json:"mad"`
	// 		Mair     int `json:"mair"`
	// 		Malnj    int `json:"malnj"`
	// 		Manc     int `json:"manc"`
	// 		Mas      int `json:"mas"`
	// 		Mau      int `json:"mau"`
	// 		Mav      int `json:"mav"`
	// 		Mbk      int `json:"mbk"`
	// 		Mbl      int `json:"mbl"`
	// 		Mbsc     int `json:"mbsc"`
	// 		Mccl     int `json:"mccl"`
	// 		Mch      int `json:"mch"`
	// 		Mchg     int `json:"mchg"`
	// 		Mchn     int `json:"mchn"`
	// 		Mci      int `json:"mci"`
	// 		Mcl      int `json:"mcl"`
	// 		Mcm      int `json:"mcm"`
	// 		Mcn      int `json:"mcn"`
	// 		Mcr      int `json:"mcr"`
	// 		Mdc      int `json:"mdc"`
	// 		Mdd      int `json:"mdd"`
	// 		Mddl     int `json:"mddl"`
	// 		Mddlb    int `json:"mddlb"`
	// 		Mddlbr   int `json:"mddlbr"`
	// 		Mddls    int `json:"mddls"`
	// 		Mdf      int `json:"mdf"`
	// 		Mdfr     int `json:"mdfr"`
	// 		Mdfrd    int `json:"mdfrd"`
	// 		Mdl      int `json:"mdl"`
	// 		Mdln     int `json:"mdln"`
	// 		Mdlt     int `json:"mdlt"`
	// 		Mdnb     int `json:"mdnb"`
	// 		Mdnh     int `json:"mdnh"`
	// 		Mdp      int `json:"mdp"`
	// 		Mdr      int `json:"mdr"`
	// 		Mdsn     int `json:"mdsn"`
	// 		Mdst     int `json:"mdst"`
	// 		Mdstne   int `json:"mdstne"`
	// 		Mdw      int `json:"mdw"`
	// 		Mel      int `json:"mel"`
	// 		Mem      int `json:"mem"`
	// 		Men      int `json:"men"`
	// 		Mfd      int `json:"mfd"`
	// 		Mfm      int `json:"mfm"`
	// 		Mhw      int `json:"mhw"`
	// 		Mia      int `json:"mia"`
	// 		Mid      int `json:"mid"`
	// 		Mjx      int `json:"mjx"`
	// 		Mke      int `json:"mke"`
	// 		Mlb      int `json:"mlb"`
	// 		Mld      int `json:"mld"`
	// 		Mlf      int `json:"mlf"`
	// 		Mlfr     int `json:"mlfr"`
	// 		Mlfrd    int `json:"mlfrd"`
	// 		Mlfrdn   int `json:"mlfrdn"`
	// 		Mlg      int `json:"mlg"`
	// 		Mllb     int `json:"mllb"`
	// 		Mllc     int `json:"mllc"`
	// 		Mllm     int `json:"mllm"`
	// 		Mllr     int `json:"mllr"`
	// 		Mllv     int `json:"mllv"`
	// 		Mln      int `json:"mln"`
	// 		Mlnn     int `json:"mlnn"`
	// 		Mlr      int `json:"mlr"`
	// 		Mlt      int `json:"mlt"`
	// 		Mlu      int `json:"mlu"`
	// 		Mlv      int `json:"mlv"`
	// 		Mlvl     int `json:"mlvl"`
	// 		Mmb      int `json:"mmb"`
	// 		Mmr      int `json:"mmr"`
	// 		Mnc      int `json:"mnc"`
	// 		Mnch     int `json:"mnch"`
	// 		Mnd      int `json:"mnd"`
	// 		Mnh      int `json:"mnh"`
	// 		Mnht     int `json:"mnht"`
	// 		Mnj      int `json:"mnj"`
	// 		Mnl      int `json:"mnl"`
	// 		Mnm      int `json:"mnm"`
	// 		Mnmt     int `json:"mnmt"`
	// 		Mnr      int `json:"mnr"`
	// 		Mnrg     int `json:"mnrg"`
	// 		Mnrv     int `json:"mnrv"`
	// 		Mnt      int `json:"mnt"`
	// 		Mntc     int `json:"mntc"`
	// 		Mntg     int `json:"mntg"`
	// 		Mntn     int `json:"mntn"`
	// 		Mntnb    int `json:"mntnb"`
	// 		Mntr     int `json:"mntr"`
	// 		Mntrry   int `json:"mntrry"`
	// 		Mntrv    int `json:"mntrv"`
	// 		Mnz      int `json:"mnz"`
	// 		Mon      int `json:"mon"`
	// 		Mot      int `json:"mot"`
	// 		Moy      int `json:"moy"`
	// 		Mplg     int `json:"mplg"`
	// 		Mplw     int `json:"mplw"`
	// 		Mpt      int `json:"mpt"`
	// 		Mrb      int `json:"mrb"`
	// 		Mrc      int `json:"mrc"`
	// 		Mrcd     int `json:"mrcd"`
	// 		Mrch     int `json:"mrch"`
	// 		Mrcr     int `json:"mrcr"`
	// 		Mrcs     int `json:"mrcs"`
	// 		Mrd      int `json:"mrd"`
	// 		Mrg      int `json:"mrg"`
	// 		Mrlb     int `json:"mrlb"`
	// 		Mrm      int `json:"mrm"`
	// 		Mrmr     int `json:"mrmr"`
	// 		Mrn      int `json:"mrn"`
	// 		Mrrl     int `json:"mrrl"`
	// 		Mrrr     int `json:"mrrr"`
	// 		Mrrw     int `json:"mrrw"`
	// 		Mrs      int `json:"mrs"`
	// 		Mrsh     int `json:"mrsh"`
	// 		Mrshl    int `json:"mrshl"`
	// 		Mrsv     int `json:"mrsv"`
	// 		Mrt      int `json:"mrt"`
	// 		Mrth     int `json:"mrth"`
	// 		Mrtt     int `json:"mrtt"`
	// 		Mrw      int `json:"mrw"`
	// 		Mry      int `json:"mry"`
	// 		Mryn     int `json:"mryn"`
	// 		Msc      int `json:"msc"`
	// 		Mshp     int `json:"mshp"`
	// 		Msk      int `json:"msk"`
	// 		Msl      int `json:"msl"`
	// 		Msm      int `json:"msm"`
	// 		Msn      int `json:"msn"`
	// 		Msp      int `json:"msp"`
	// 		Msqp     int `json:"msqp"`
	// 		Mss      int `json:"mss"`
	// 		Mssn     int `json:"mssn"`
	// 		Mtb      int `json:"mtb"`
	// 		Mtc      int `json:"mtc"`
	// 		Mtch     int `json:"mtch"`
	// 		Mtf      int `json:"mtf"`
	// 		Mtg      int `json:"mtg"`
	// 		Mtj      int `json:"mtj"`
	// 		Mtk      int `json:"mtk"`
	// 		Mtl      int `json:"mtl"`
	// 		Mtln     int `json:"mtln"`
	// 		Mtr      int `json:"mtr"`
	// 		Mtrl     int `json:"mtrl"`
	// 		Mtt      int `json:"mtt"`
	// 		Mtv      int `json:"mtv"`
	// 		Mty      int `json:"mty"`
	// 		Mvj      int `json:"mvj"`
	// 		Mwc      int `json:"mwc"`
	// 		Mwh      int `json:"mwh"`
	// 		Mxc      int `json:"mxc"`
	// 		Mys      int `json:"mys"`
	// 		Nar      int `json:"nar"`
	// 		Nat      int `json:"nat"`
	// 		Nb       int `json:"nb"`
	// 		Nbm      int `json:"nbm"`
	// 		Nbp      int `json:"nbp"`
	// 		Nct      int `json:"nct"`
	// 		Ndh      int `json:"ndh"`
	// 		Ndt      int `json:"ndt"`
	// 		NearMe   int `json:"near_me"`
	// 		Newb     int `json:"newb"`
	// 		Nfk      int `json:"nfk"`
	// 		Nfl      int `json:"nfl"`
	// 		Ngy      int `json:"ngy"`
	// 		Nj       int `json:"nj"`
	// 		Nmk      int `json:"nmk"`
	// 		Nmr      int `json:"nmr"`
	// 		Nnh      int `json:"nnh"`
	// 		Nnk      int `json:"nnk"`
	// 		No       int `json:"no"`
	// 		Npa      int `json:"npa"`
	// 		Npk      int `json:"npk"`
	// 		Npm      int `json:"npm"`
	// 		Nps      int `json:"nps"`
	// 		Npt      int `json:"npt"`
	// 		Npv      int `json:"npv"`
	// 		Nrf      int `json:"nrf"`
	// 		Nrm      int `json:"nrm"`
	// 		Nrmb     int `json:"nrmb"`
	// 		Nrr      int `json:"nrr"`
	// 		Nrrw     int `json:"nrrw"`
	// 		Nrt      int `json:"nrt"`
	// 		Nrth     int `json:"nrth"`
	// 		Nrthb    int `json:"nrthb"`
	// 		Nrthbr   int `json:"nrthbr"`
	// 		Nrthc    int `json:"nrthc"`
	// 		Nrthh    int `json:"nrthh"`
	// 		Nrthhl   int `json:"nrthhl"`
	// 		Nrthk    int `json:"nrthk"`
	// 		Nrthl    int `json:"nrthl"`
	// 		Nrthp    int `json:"nrthp"`
	// 		Nrths    int `json:"nrths"`
	// 		Nrthsl   int `json:"nrthsl"`
	// 		Nrtht    int `json:"nrtht"`
	// 		Nrthv    int `json:"nrthv"`
	// 		Nrw      int `json:"nrw"`
	// 		Nsh      int `json:"nsh"`
	// 		Ntc      int `json:"ntc"`
	// 		Nth      int `json:"nth"`
	// 		Nthmn    int `json:"nthmn"`
	// 		Ntr      int `json:"ntr"`
	// 		Ntt      int `json:"ntt"`
	// 		Nvd      int `json:"nvd"`
	// 		Nvmm     int `json:"nvmm"`
	// 		Nvr      int `json:"nvr"`
	// 		Nvt      int `json:"nvt"`
	// 		Nwbd     int `json:"nwbd"`
	// 		Nwbr     int `json:"nwbr"`
	// 		Nwbrn    int `json:"nwbrn"`
	// 		Nwbry    int `json:"nwbry"`
	// 		Nwc      int `json:"nwc"`
	// 		Nwcs     int `json:"nwcs"`
	// 		Nwd      int `json:"nwd"`
	// 		Nwh      int `json:"nwh"`
	// 		Nwhm     int `json:"nwhm"`
	// 		Nwhv     int `json:"nwhv"`
	// 		Nwhy     int `json:"nwhy"`
	// 		Nwk      int `json:"nwk"`
	// 		Nwkn     int `json:"nwkn"`
	// 		Nwl      int `json:"nwl"`
	// 		Nwm      int `json:"nwm"`
	// 		Nwp      int `json:"nwp"`
	// 		Nwpl     int `json:"nwpl"`
	// 		Nwpr     int `json:"nwpr"`
	// 		Nwprt    int `json:"nwprt"`
	// 		Nwprtt   int `json:"nwprtt"`
	// 		Nwq      int `json:"nwq"`
	// 		Nwr      int `json:"nwr"`
	// 		Nwrk     int `json:"nwrk"`
	// 		Nwrkd    int `json:"nwrkd"`
	// 		Nwrkh    int `json:"nwrkh"`
	// 		Nwt      int `json:"nwt"`
	// 		Nwtn     int `json:"nwtn"`
	// 		Nwtnb    int `json:"nwtnb"`
	// 		Nwtnn    int `json:"nwtnn"`
	// 		Nwtw     int `json:"nwtw"`
	// 		Nwtwn    int `json:"nwtwn"`
	// 		Nww      int `json:"nww"`
	// 		Nxm      int `json:"nxm"`
	// 		Ny       int `json:"ny"`
	// 		Nya      int `json:"nya"`
	// 		Nzr      int `json:"nzr"`
	// 		Oak      int `json:"oak"`
	// 		Obr      int `json:"obr"`
	// 		Obx      int `json:"obx"`
	// 		Ocm      int `json:"ocm"`
	// 		Ocnc     int `json:"ocnc"`
	// 		Oco      int `json:"oco"`
	// 		Odp      int `json:"odp"`
	// 		Ods      int `json:"ods"`
	// 		Ogm      int `json:"ogm"`
	// 		Ogqt     int `json:"ogqt"`
	// 		Ojc      int `json:"ojc"`
	// 		Okb      int `json:"okb"`
	// 		Okc      int `json:"okc"`
	// 		Okl      int `json:"okl"`
	// 		Okm      int `json:"okm"`
	// 		Okv      int `json:"okv"`
	// 		Old      int `json:"old"`
	// 		Oldr     int `json:"oldr"`
	// 		Oldw     int `json:"oldw"`
	// 		Oln      int `json:"oln"`
	// 		Olp      int `json:"olp"`
	// 		Olt      int `json:"olt"`
	// 		Oma      int `json:"oma"`
	// 		Omk      int `json:"omk"`
	// 		Onn      int `json:"onn"`
	// 		Ont      int `json:"ont"`
	// 		Onz      int `json:"onz"`
	// 		Opl      int `json:"opl"`
	// 		Orc      int `json:"orc"`
	// 		Ore      int `json:"ore"`
	// 		Org      int `json:"org"`
	// 		Orgn     int `json:"orgn"`
	// 		Orh      int `json:"orh"`
	// 		Ork      int `json:"ork"`
	// 		Orl      int `json:"orl"`
	// 		Orln     int `json:"orln"`
	// 		Orm      int `json:"orm"`
	// 		Orn      int `json:"orn"`
	// 		Orngp    int `json:"orngp"`
	// 		Osd      int `json:"osd"`
	// 		Osk      int `json:"osk"`
	// 		Osl      int `json:"osl"`
	// 		Ost      int `json:"ost"`
	// 		Ott      int `json:"ott"`
	// 		Ovd      int `json:"ovd"`
	// 		Ovr      int `json:"ovr"`
	// 		Oxf      int `json:"oxf"`
	// 		Oxn      int `json:"oxn"`
	// 		Pal      int `json:"pal"`
	// 		Par      int `json:"par"`
	// 		Pb       int `json:"pb"`
	// 		Pbl      int `json:"pbl"`
	// 		Pcb      int `json:"pcb"`
	// 		Pch      int `json:"pch"`
	// 		Pcht     int `json:"pcht"`
	// 		Pct      int `json:"pct"`
	// 		Pcu      int `json:"pcu"`
	// 		Pdx      int `json:"pdx"`
	// 		Pec      int `json:"pec"`
	// 		Per      int `json:"per"`
	// 		Pft      int `json:"pft"`
	// 		Pgd      int `json:"pgd"`
	// 		Ph9      int `json:"ph9"`
	// 		Pha      int `json:"pha"`
	// 		Phl      int `json:"phl"`
	// 		Phn      int `json:"phn"`
	// 		Phsm     int `json:"phsm"`
	// 		Phx      int `json:"phx"`
	// 		Pier     int `json:"pier"`
	// 		Pit      int `json:"pit"`
	// 		Pkr      int `json:"pkr"`
	// 		Plk      int `json:"plk"`
	// 		Plm      int `json:"plm"`
	// 		Plmc     int `json:"plmc"`
	// 		Plmd     int `json:"plmd"`
	// 		Plmh     int `json:"plmh"`
	// 		Plml     int `json:"plml"`
	// 		Plmy     int `json:"plmy"`
	// 		Pln      int `json:"pln"`
	// 		Plnt     int `json:"plnt"`
	// 		Plp      int `json:"plp"`
	// 		Plr      int `json:"plr"`
	// 		Plrm     int `json:"plrm"`
	// 		Pls      int `json:"pls"`
	// 		Plsh     int `json:"plsh"`
	// 		Plsn     int `json:"plsn"`
	// 		Plsntv   int `json:"plsntv"`
	// 		Pltn     int `json:"pltn"`
	// 		Ply      int `json:"ply"`
	// 		Plydl    int `json:"plydl"`
	// 		Plym     int `json:"plym"`
	// 		Plymt    int `json:"plymt"`
	// 		Plymth   int `json:"plymth"`
	// 		Pmb      int `json:"pmb"`
	// 		Pmbr     int `json:"pmbr"`
	// 		Pmbrk    int `json:"pmbrk"`
	// 		Pmp      int `json:"pmp"`
	// 		Pnd      int `json:"pnd"`
	// 		Pndr     int `json:"pndr"`
	// 		Pnf      int `json:"pnf"`
	// 		Png      int `json:"png"`
	// 		Pnm      int `json:"pnm"`
	// 		Pnn      int `json:"pnn"`
	// 		Pnnn     int `json:"pnnn"`
	// 		Pnpl     int `json:"pnpl"`
	// 		Pnr      int `json:"pnr"`
	// 		Pns      int `json:"pns"`
	// 		Pnt      int `json:"pnt"`
	// 		Pntp     int `json:"pntp"`
	// 		Pntpl    int `json:"pntpl"`
	// 		Pntv     int `json:"pntv"`
	// 		Pnv      int `json:"pnv"`
	// 		Pou      int `json:"pou"`
	// 		Ppc      int `json:"ppc"`
	// 		Prai     int `json:"prai"`
	// 		Prc      int `json:"prc"`
	// 		Prct     int `json:"prct"`
	// 		Prk      int `json:"prk"`
	// 		Prkr     int `json:"prkr"`
	// 		Prks     int `json:"prks"`
	// 		Prl      int `json:"prl"`
	// 		Prms     int `json:"prms"`
	// 		Prr      int `json:"prr"`
	// 		Prrm     int `json:"prrm"`
	// 		Prrr     int `json:"prrr"`
	// 		Prrv     int `json:"prrv"`
	// 		Prry     int `json:"prry"`
	// 		Prs      int `json:"prs"`
	// 		Prt      int `json:"prt"`
	// 		Prtc     int `json:"prtc"`
	// 		Prtch    int `json:"prtch"`
	// 		Prtl     int `json:"prtl"`
	// 		Prtp     int `json:"prtp"`
	// 		Prtr     int `json:"prtr"`
	// 		Prts     int `json:"prts"`
	// 		Prtst    int `json:"prtst"`
	// 		Prtw     int `json:"prtw"`
	// 		Prv      int `json:"prv"`
	// 		Psd      int `json:"psd"`
	// 		Psm      int `json:"psm"`
	// 		Psp      int `json:"psp"`
	// 		Psr      int `json:"psr"`
	// 		Ptc      int `json:"ptc"`
	// 		Ptl      int `json:"ptl"`
	// 		Ptmn     int `json:"ptmn"`
	// 		Ptt      int `json:"ptt"`
	// 		Ptts     int `json:"ptts"`
	// 		Pttsb    int `json:"pttsb"`
	// 		Pttsbr   int `json:"pttsbr"`
	// 		Pttsf    int `json:"pttsf"`
	// 		Pva      int `json:"pva"`
	// 		Pvd      int `json:"pvd"`
	// 		Pvo      int `json:"pvo"`
	// 		Pwl      int `json:"pwl"`
	// 		Pwll     int `json:"pwll"`
	// 		Pwly     int `json:"pwly"`
	// 		Pwm      int `json:"pwm"`
	// 		Pwt      int `json:"pwt"`
	// 		Qec      int `json:"qec"`
	// 		Qnc      int `json:"qnc"`
	// 		Qtm      int `json:"qtm"`
	// 		Rbb      int `json:"rbb"`
	// 		Rbk      int `json:"rbk"`
	// 		Rca      int `json:"rca"`
	// 		Rcf      int `json:"rcf"`
	// 		Rch      int `json:"rch"`
	// 		Rchh     int `json:"rchh"`
	// 		Rchl     int `json:"rchl"`
	// 		Rchmn    int `json:"rchmn"`
	// 		Rck      int `json:"rck"`
	// 		Rckl     int `json:"rckl"`
	// 		Rckv     int `json:"rckv"`
	// 		Rckw     int `json:"rckw"`
	// 		Rcn      int `json:"rcn"`
	// 		Rdg      int `json:"rdg"`
	// 		Rdgf     int `json:"rdgf"`
	// 		Rdgfl    int `json:"rdgfl"`
	// 		Rdh      int `json:"rdh"`
	// 		Rdj      int `json:"rdj"`
	// 		Rdl      int `json:"rdl"`
	// 		Rdln     int `json:"rdln"`
	// 		Rdm      int `json:"rdm"`
	// 		Rdn      int `json:"rdn"`
	// 		Rdu      int `json:"rdu"`
	// 		Rdw      int `json:"rdw"`
	// 		Red      int `json:"red"`
	// 		Rgr      int `json:"rgr"`
	// 		Rgrs     int `json:"rgrs"`
	// 		Rgrsr    int `json:"rgrsr"`
	// 		Rhb      int `json:"rhb"`
	// 		Ri       int `json:"ri"`
	// 		Ric      int `json:"ric"`
	// 		Rivo     int `json:"rivo"`
	// 		Rkv      int `json:"rkv"`
	// 		Rmg      int `json:"rmg"`
	// 		Rmi      int `json:"rmi"`
	// 		Rmm      int `json:"rmm"`
	// 		Rmn      int `json:"rmn"`
	// 		Rms      int `json:"rms"`
	// 		Rnd      int `json:"rnd"`
	// 		Rndl     int `json:"rndl"`
	// 		Rndt     int `json:"rndt"`
	// 		Rnm      int `json:"rnm"`
	// 		Rno      int `json:"rno"`
	// 		Rns      int `json:"rns"`
	// 		Roa      int `json:"roa"`
	// 		Roc      int `json:"roc"`
	// 		Roch     int `json:"roch"`
	// 		Roy      int `json:"roy"`
	// 		Rsh      int `json:"rsh"`
	// 		Rsll     int `json:"rsll"`
	// 		Rsln     int `json:"rsln"`
	// 		Rsly     int `json:"rsly"`
	// 		Rsm      int `json:"rsm"`
	// 		Rst      int `json:"rst"`
	// 		Rsv      int `json:"rsv"`
	// 		Rsw      int `json:"rsw"`
	// 		Rswl     int `json:"rswl"`
	// 		Rut      int `json:"rut"`
	// 		Rvr      int `json:"rvr"`
	// 		Rvrd     int `json:"rvrd"`
	// 		Rvrf     int `json:"rvrf"`
	// 		Rvrm     int `json:"rvrm"`
	// 		Rvrs     int `json:"rvrs"`
	// 		Rwl      int `json:"rwl"`
	// 		Rydh     int `json:"rydh"`
	// 		Ryl      int `json:"ryl"`
	// 		Rylk     int `json:"rylk"`
	// 		Ryn      int `json:"ryn"`
	// 		Ryns     int `json:"ryns"`
	// 		Ryt      int `json:"ryt"`
	// 		Sa       int `json:"sa"`
	// 		Saal     int `json:"saal"`
	// 		Sac      int `json:"sac"`
	// 		Sar      int `json:"sar"`
	// 		Sau      int `json:"sau"`
	// 		Sav      int `json:"sav"`
	// 		Sba      int `json:"sba"`
	// 		Sbn      int `json:"sbn"`
	// 		Sbr      int `json:"sbr"`
	// 		Sbs      int `json:"sbs"`
	// 		Sch      int `json:"sch"`
	// 		Schm     int `json:"schm"`
	// 		Schn     int `json:"schn"`
	// 		Scl      int `json:"scl"`
	// 		Sclf     int `json:"sclf"`
	// 		Scm      int `json:"scm"`
	// 		Scr      int `json:"scr"`
	// 		Sct      int `json:"sct"`
	// 		Sctt     int `json:"sctt"`
	// 		Scz      int `json:"scz"`
	// 		Sd       int `json:"sd"`
	// 		Sdb      int `json:"sdb"`
	// 		Sdd      int `json:"sdd"`
	// 		Sdr      int `json:"sdr"`
	// 		Sea      int `json:"sea"`
	// 		Sed      int `json:"sed"`
	// 		See      int `json:"see"`
	// 		Sf       int `json:"sf"`
	// 		Sfd      int `json:"sfd"`
	// 		Sff      int `json:"sff"`
	// 		Sfn      int `json:"sfn"`
	// 		Sfs      int `json:"sfs"`
	// 		Sft      int `json:"sft"`
	// 		Sgn      int `json:"sgn"`
	// 		Sgr      int `json:"sgr"`
	// 		Sgrg     int `json:"sgrg"`
	// 		Sgrt     int `json:"sgrt"`
	// 		Sgv      int `json:"sgv"`
	// 		Shb      int `json:"shb"`
	// 		Shf      int `json:"shf"`
	// 		Shl      int `json:"shl"`
	// 		Shlt     int `json:"shlt"`
	// 		Shn      int `json:"shn"`
	// 		Shrd     int `json:"shrd"`
	// 		Shrm     int `json:"shrm"`
	// 		Shrv     int `json:"shrv"`
	// 		Shrw     int `json:"shrw"`
	// 		Sj       int `json:"sj"`
	// 		Sjc      int `json:"sjc"`
	// 		Sjn      int `json:"sjn"`
	// 		Sjw      int `json:"sjw"`
	// 		Ski      int `json:"ski"`
	// 		Skk      int `json:"skk"`
	// 		Slb      int `json:"slb"`
	// 		Slc      int `json:"slc"`
	// 		Slh      int `json:"slh"`
	// 		Sll      int `json:"sll"`
	// 		Slm      int `json:"slm"`
	// 		Sln      int `json:"sln"`
	// 		Slnb     int `json:"slnb"`
	// 		Slo      int `json:"slo"`
	// 		Sls      int `json:"sls"`
	// 		Slv      int `json:"slv"`
	// 		Slvr     int `json:"slvr"`
	// 		Slvrs    int `json:"slvrs"`
	// 		Slz      int `json:"slz"`
	// 		Smm      int `json:"smm"`
	// 		Smr      int `json:"smr"`
	// 		Smrs     int `json:"smrs"`
	// 		Smrv     int `json:"smrv"`
	// 		Sms      int `json:"sms"`
	// 		Smt      int `json:"smt"`
	// 		Smth     int `json:"smth"`
	// 		Smtr     int `json:"smtr"`
	// 		Sna      int `json:"sna"`
	// 		Snc      int `json:"snc"`
	// 		Sncr     int `json:"sncr"`
	// 		Sndh     int `json:"sndh"`
	// 		Sndrl    int `json:"sndrl"`
	// 		Sndw     int `json:"sndw"`
	// 		Sndy     int `json:"sndy"`
	// 		Snf      int `json:"snf"`
	// 		Snj      int `json:"snj"`
	// 		Snl      int `json:"snl"`
	// 		Snln     int `json:"snln"`
	// 		Snm      int `json:"snm"`
	// 		Snmt     int `json:"snmt"`
	// 		Snn      int `json:"snn"`
	// 		Snng     int `json:"snng"`
	// 		Snp      int `json:"snp"`
	// 		Snr      int `json:"snr"`
	// 		Snrf     int `json:"snrf"`
	// 		Snrm     int `json:"snrm"`
	// 		Snt      int `json:"snt"`
	// 		Sntc     int `json:"sntc"`
	// 		Sntg     int `json:"sntg"`
	// 		Sntn     int `json:"sntn"`
	// 		Spd      int `json:"spd"`
	// 		Spk      int `json:"spk"`
	// 		Spl      int `json:"spl"`
	// 		Spp      int `json:"spp"`
	// 		Spr      int `json:"spr"`
	// 		Sprc     int `json:"sprc"`
	// 		Sprn     int `json:"sprn"`
	// 		Sprng    int `json:"sprng"`
	// 		Sprngc   int `json:"sprngc"`
	// 		Sprngf   int `json:"sprngf"`
	// 		Sprngl   int `json:"sprngl"`
	// 		Sprngt   int `json:"sprngt"`
	// 		Sra      int `json:"sra"`
	// 		Srd      int `json:"srd"`
	// 		Srf      int `json:"srf"`
	// 		Srp      int `json:"srp"`
	// 		Srs      int `json:"srs"`
	// 		Ssd      int `json:"ssd"`
	// 		Sst      int `json:"sst"`
	// 		Stag     int `json:"stag"`
	// 		Stc      int `json:"stc"`
	// 		Stfr     int `json:"stfr"`
	// 		Stg      int `json:"stg"`
	// 		Sth      int `json:"sth"`
	// 		Sthj     int `json:"sthj"`
	// 		Sthk     int `json:"sthk"`
	// 		Sthl     int `json:"sthl"`
	// 		Sthld    int `json:"sthld"`
	// 		Sthn     int `json:"sthn"`
	// 		Sthp     int `json:"sthp"`
	// 		Sthv     int `json:"sthv"`
	// 		Stj      int `json:"stj"`
	// 		Stl      int `json:"stl"`
	// 		Stll     int `json:"stll"`
	// 		Stm      int `json:"stm"`
	// 		Stmb     int `json:"stmb"`
	// 		Stmc     int `json:"stmc"`
	// 		Stn      int `json:"stn"`
	// 		Stnh     int `json:"stnh"`
	// 		Stnn     int `json:"stnn"`
	// 		Stny     int `json:"stny"`
	// 		Sto      int `json:"sto"`
	// 		Stpt     int `json:"stpt"`
	// 		Stptb    int `json:"stptb"`
	// 		Stptr    int `json:"stptr"`
	// 		Str      int `json:"str"`
	// 		Strl     int `json:"strl"`
	// 		Strln    int `json:"strln"`
	// 		Strlng   int `json:"strlng"`
	// 		Strt     int `json:"strt"`
	// 		Strtf    int `json:"strtf"`
	// 		Strth    int `json:"strth"`
	// 		Strtm    int `json:"strtm"`
	// 		Strts    int `json:"strts"`
	// 		Sts      int `json:"sts"`
	// 		Stsm     int `json:"stsm"`
	// 		Stt      int `json:"stt"`
	// 		Sttc     int `json:"sttc"`
	// 		Stts     int `json:"stts"`
	// 		Stv      int `json:"stv"`
	// 		Stw      int `json:"stw"`
	// 		Stwn     int `json:"stwn"`
	// 		Stwv     int `json:"stwv"`
	// 		Stz      int `json:"stz"`
	// 		Svd      int `json:"svd"`
	// 		Svg      int `json:"svg"`
	// 		Svgm     int `json:"svgm"`
	// 		Svgmd    int `json:"svgmd"`
	// 		Svl      int `json:"svl"`
	// 		Svll     int `json:"svll"`
	// 		Svr      int `json:"svr"`
	// 		Swa      int `json:"swa"`
	// 		Swm      int `json:"swm"`
	// 		Swn      int `json:"swn"`
	// 		Sxp      int `json:"sxp"`
	// 		Syd      int `json:"syd"`
	// 		Syk      int `json:"syk"`
	// 		Syl      int `json:"syl"`
	// 		Syr      int `json:"syr"`
	// 		Tac      int `json:"tac"`
	// 		Tck      int `json:"tck"`
	// 		Tckh     int `json:"tckh"`
	// 		Tcs      int `json:"tcs"`
	// 		Terl     int `json:"terl"`
	// 		Test     int `json:"test"`
	// 		Thb      int `json:"thb"`
	// 		Thc      int `json:"thc"`
	// 		Ths      int `json:"ths"`
	// 		Thsn     int `json:"thsn"`
	// 		Thv      int `json:"thv"`
	// 		Thw      int `json:"thw"`
	// 		Tkm      int `json:"tkm"`
	// 		Tkw      int `json:"tkw"`
	// 		Tky      int `json:"tky"`
	// 		Tld      int `json:"tld"`
	// 		Tlds     int `json:"tlds"`
	// 		Tlg      int `json:"tlg"`
	// 		Tll      int `json:"tll"`
	// 		Tllm     int `json:"tllm"`
	// 		Tlt      int `json:"tlt"`
	// 		Tmc      int `json:"tmc"`
	// 		Tml      int `json:"tml"`
	// 		Tmn      int `json:"tmn"`
	// 		Tmp      int `json:"tmp"`
	// 		Tmpl     int `json:"tmpl"`
	// 		Tmr      int `json:"tmr"`
	// 		Tmrc     int `json:"tmrc"`
	// 		Tms      int `json:"tms"`
	// 		Tnc      int `json:"tnc"`
	// 		Tnl      int `json:"tnl"`
	// 		Tnn      int `json:"tnn"`
	// 		Tnt      int `json:"tnt"`
	// 		Tnv      int `json:"tnv"`
	// 		Toa      int `json:"toa"`
	// 		Tor      int `json:"tor"`
	// 		Tpa      int `json:"tpa"`
	// 		Trc      int `json:"trc"`
	// 		Trf      int `json:"trf"`
	// 		Trl      int `json:"trl"`
	// 		Trn      int `json:"trn"`
	// 		Trno     int `json:"trno"`
	// 		Trp      int `json:"trp"`
	// 		Trr      int `json:"trr"`
	// 		Trrm     int `json:"trrm"`
	// 		Trv      int `json:"trv"`
	// 		Try      int `json:"try"`
	// 		Tsm      int `json:"tsm"`
	// 		Tsn      int `json:"tsn"`
	// 		Tsnm     int `json:"tsnm"`
	// 		Ttn      int `json:"ttn"`
	// 		Ttnv     int `json:"ttnv"`
	// 		Ttw      int `json:"ttw"`
	// 		Tul      int `json:"tul"`
	// 		Tus      int `json:"tus"`
	// 		Twd      int `json:"twd"`
	// 		Twn      int `json:"twn"`
	// 		Tws      int `json:"tws"`
	// 		Tyb      int `json:"tyb"`
	// 		Tyl      int `json:"tyl"`
	// 		Unc      int `json:"unc"`
	// 		Unnc     int `json:"unnc"`
	// 		Uok      int `json:"uok"`
	// 		Urb      int `json:"urb"`
	// 		Utc      int `json:"utc"`
	// 		Uxwbny   int `json:"uxwbny"`
	// 		Vanc     int `json:"vanc"`
	// 		Vans     int `json:"vans"`
	// 		Vct      int `json:"vct"`
	// 		Vctrn    int `json:"vctrn"`
	// 		Vgb      int `json:"vgb"`
	// 		Vir      int `json:"vir"`
	// 		Vl       int `json:"vl"`
	// 		Vlkn     int `json:"vlkn"`
	// 		Vll      int `json:"vll"`
	// 		Vllg     int `json:"vllg"`
	// 		Vln      int `json:"vln"`
	// 		Vlnc     int `json:"vlnc"`
	// 		Vlp      int `json:"vlp"`
	// 		Vlpr     int `json:"vlpr"`
	// 		Vnc      int `json:"vnc"`
	// 		Vnce     int `json:"vnce"`
	// 		Vncf     int `json:"vncf"`
	// 		Vnl      int `json:"vnl"`
	// 		Vnns     int `json:"vnns"`
	// 		Vntn     int `json:"vntn"`
	// 		Vrh      int `json:"vrh"`
	// 		Vrn      int `json:"vrn"`
	// 		Vrnn     int `json:"vrnn"`
	// 		Vrs      int `json:"vrs"`
	// 		Vsl      int `json:"vsl"`
	// 		Vst      int `json:"vst"`
	// 		Vzc      int `json:"vzc"`
	// 		Wal      int `json:"wal"`
	// 		Washva   int `json:"washva"`
	// 		Wch      int `json:"wch"`
	// 		Wchp     int `json:"wchp"`
	// 		Wco      int `json:"wco"`
	// 		Wdb      int `json:"wdb"`
	// 		Wdf      int `json:"wdf"`
	// 		Wdl      int `json:"wdl"`
	// 		Wdr      int `json:"wdr"`
	// 		Wds      int `json:"wds"`
	// 		Wdsd     int `json:"wdsd"`
	// 		Whr      int `json:"whr"`
	// 		Wht      int `json:"wht"`
	// 		Whtb     int `json:"whtb"`
	// 		Whtby    int `json:"whtby"`
	// 		Whtl     int `json:"whtl"`
	// 		Whtn     int `json:"whtn"`
	// 		Whtp     int `json:"whtp"`
	// 		Whw      int `json:"whw"`
	// 		Wic      int `json:"wic"`
	// 		Wil      int `json:"wil"`
	// 		Wkf      int `json:"wkf"`
	// 		Wkv      int `json:"wkv"`
	// 		Wlc      int `json:"wlc"`
	// 		Wld      int `json:"wld"`
	// 		Wldw     int `json:"wldw"`
	// 		Wlf      int `json:"wlf"`
	// 		Wlk      int `json:"wlk"`
	// 		Wll      int `json:"wll"`
	// 		Wllg     int `json:"wllg"`
	// 		Wllm     int `json:"wllm"`
	// 		Wllms    int `json:"wllms"`
	// 		Wlln     int `json:"wlln"`
	// 		Wlls     int `json:"wlls"`
	// 		Wlmt     int `json:"wlmt"`
	// 		Wln      int `json:"wln"`
	// 		Wls      int `json:"wls"`
	// 		Wlt      int `json:"wlt"`
	// 		Wma      int `json:"wma"`
	// 		Wmb      int `json:"wmb"`
	// 		Wmt      int `json:"wmt"`
	// 		Wnc      int `json:"wnc"`
	// 		Wnd      int `json:"wnd"`
	// 		Wndh     int `json:"wndh"`
	// 		Wnds     int `json:"wnds"`
	// 		Wnf      int `json:"wnf"`
	// 		Wnn      int `json:"wnn"`
	// 		Wnnt     int `json:"wnnt"`
	// 		Wnp      int `json:"wnp"`
	// 		Wns      int `json:"wns"`
	// 		Wnt      int `json:"wnt"`
	// 		Wpk      int `json:"wpk"`
	// 		Wpp      int `json:"wpp"`
	// 		Wrm      int `json:"wrm"`
	// 		Wrn      int `json:"wrn"`
	// 		Wrr      int `json:"wrr"`
	// 		Wrrn     int `json:"wrrn"`
	// 		Wrrnv    int `json:"wrrnv"`
	// 		Wrt      int `json:"wrt"`
	// 		Wrv      int `json:"wrv"`
	// 		Wrw      int `json:"wrw"`
	// 		Wrwc     int `json:"wrwc"`
	// 		Wsh      int `json:"wsh"`
	// 		Wsk      int `json:"wsk"`
	// 		Wsm      int `json:"wsm"`
	// 		Wst      int `json:"wst"`
	// 		Wstb     int `json:"wstb"`
	// 		Wstby    int `json:"wstby"`
	// 		Wstc     int `json:"wstc"`
	// 		Wstcr    int `json:"wstcr"`
	// 		Wstf     int `json:"wstf"`
	// 		Wstfl    int `json:"wstfl"`
	// 		Wsth     int `json:"wsth"`
	// 		Wstl     int `json:"wstl"`
	// 		Wstln    int `json:"wstln"`
	// 		Wstn     int `json:"wstn"`
	// 		Wstnv    int `json:"wstnv"`
	// 		Wstp     int `json:"wstp"`
	// 		Wstpr    int `json:"wstpr"`
	// 		Wstr     int `json:"wstr"`
	// 		Wstt     int `json:"wstt"`
	// 		Wstw     int `json:"wstw"`
	// 		Wstwd    int `json:"wstwd"`
	// 		Wsw      int `json:"wsw"`
	// 		Wtf      int `json:"wtf"`
	// 		Wtl      int `json:"wtl"`
	// 		Wtr      int `json:"wtr"`
	// 		Wtrb     int `json:"wtrb"`
	// 		Wtrs     int `json:"wtrs"`
	// 		Wvr      int `json:"wvr"`
	// 		Wwd      int `json:"wwd"`
	// 		Wxf      int `json:"wxf"`
	// 		Wyl      int `json:"wyl"`
	// 		Wym      int `json:"wym"`
	// 		Wyn      int `json:"wyn"`
	// 		Wynd     int `json:"wynd"`
	// 		Wynn     int `json:"wynn"`
	// 		Wynp     int `json:"wynp"`
	// 		Wyz      int `json:"wyz"`
	// 		Yak      int `json:"yak"`
	// 		Yhz      int `json:"yhz"`
	// 		Ykh      int `json:"ykh"`
	// 		Ykhm     int `json:"ykhm"`
	// 		Ykn      int `json:"ykn"`
	// 		Yng      int `json:"yng"`
	// 		Yngh     int `json:"yngh"`
	// 		Ynt      int `json:"ynt"`
	// 		Yovohg   int `json:"yovohg"`
	// 		Yrb      int `json:"yrb"`
	// 		Yrd      int `json:"yrd"`
	// 		Yrkk     int `json:"yrkk"`
	// 		Yrm      int `json:"yrm"`
	// 		Yyz      int `json:"yyz"`
	// 		Zmr      int `json:"zmr"`
	// 		Zrg      int `json:"zrg"`
	// 	} `json:"num_days"`
	// } `json:"calendar"`
	EnableButtonTracking     int  `json:"enable_button_tracking"`
	EnableDigits             int  `json:"enable_digits"`
	EnableGlobalDiningAccess bool `json:"enable_global_dining_access"`
	EventGaMaxPartySize      int  `json:"event_ga_max_party_size"`
	Events                   struct {
	} `json:"events"`
	IphoneMinVersion string   `json:"iphone_min_version"`
	MaxBookHour      int      `json:"max_book_hour"`
	MaxPartySize     int      `json:"max_party_size"`
	MinBookHour      int      `json:"min_book_hour"`
	MinPartySize     int      `json:"min_party_size"`
	Occasions        []string `json:"occasions"`
	Profile          struct {
		Allergies []struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"allergies"`
		DietRestrictions []struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"diet_restrictions"`
		SpecialDates []struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"special_dates"`
		Wines []struct {
			ID       int    `json:"id"`
			Value    string `json:"value"`
			WineType string `json:"wine_type"`
		} `json:"wines"`
	} `json:"profile"`
	Version          int    `json:"version"`
	CalendarDateFrom string `json:"calendar_date_from"`
	CalendarDateTo   string `json:"calendar_date_to"`
	LargeParty       struct {
		Message interface{} `json:"message"`
	} `json:"large_party"`
	SearchHints struct {
		Neighborhood []string `json:"neighborhood"`
		Cuisine      []string `json:"cuisine"`
		City         []string `json:"city"`
		Restaurant   []string `json:"restaurant"`
	} `json:"search_hints"`
	UseStripeAsDefaultPaymentProcessor bool   `json:"use_stripe_as_default_payment_processor"`
	StripePublishableKey               string `json:"stripe_publishable_key"`
}

type ConfigInfoSearchHints struct {
	Neighborhood []string
	Cuisine      []string
	City         []string
	Restaurant   []string
}

type ConfigInfoVenueConfig struct {
	AllowBypassPaymentMethod int
	AllowMultipleResys       int
	EnableInvite             int
	EnableResypay            int
	HospitalityIncluded      int
}

type ConfigInfoVenueContact struct {
	PhoneNumber string
}

type ConfigInfoVenue struct {
	Config       ConfigInfoVenueConfig
	Contact      ConfigInfoVenueContact
	MaxPartySize int
	MinPartySize int
	Name         string
}

type ConfigInfo struct {
	MaxBookHour  int
	MaxPartySize int
	MinBookHour  int
	MinPartySize int
	Occasions    []string
	SearchHints  ConfigInfoSearchHints
	Venue        ConfigInfoVenue
}

func convertConfigInfoPayload(c configInfoPayload) *ConfigInfo {
	return &ConfigInfo{
		MaxBookHour:  c.MaxBookHour,
		MaxPartySize: c.MaxPartySize,
		MinBookHour:  c.MinBookHour,
		MinPartySize: c.MinPartySize,
		Occasions:    c.Occasions,
		SearchHints: ConfigInfoSearchHints{
			Neighborhood: c.SearchHints.Neighborhood,
			Cuisine:      c.SearchHints.Cuisine,
			City:         c.SearchHints.City,
			Restaurant:   c.SearchHints.Restaurant,
		},
		Venue: ConfigInfoVenue{
			Config: ConfigInfoVenueConfig{
				AllowBypassPaymentMethod: c.Venue.Config.AllowBypassPaymentMethod,
				AllowMultipleResys:       c.Venue.Config.AllowMultipleResys,
				EnableInvite:             c.Venue.Config.EnableInvite,
				EnableResypay:            c.Venue.Config.EnableResypay,
				HospitalityIncluded:      c.Venue.Config.HospitalityIncluded,
			},
			Contact: ConfigInfoVenueContact{
				PhoneNumber: c.Venue.Contact.PhoneNumber,
			},
			MaxPartySize: c.Venue.MaxPartySize,
			MinPartySize: c.Venue.MinPartySize,
			Name:         c.Venue.Name,
		},
	}
}

//go:generate genopts --function Config --params --required "venueID int" --extends Base
func (c *Client) Config(venueID int, optss ...ConfigOption) (*ConfigInfo, error) {
	opts := MakeConfigOptions(optss...)

	uri := request.MakeURL("https://api.resy.com/2/config",
		request.Param{"venue_id", venueID},
	)
	headers := c.makeHeaders(false, opts.ToBaseOptions()...)

	var payload configInfoPayload

	if _, err := request.Get(uri, &payload, request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}

	if opts.DebugPayload() {
		log.Printf("payload: %s", payload)
	}

	res := convertConfigInfoPayload(payload)

	return res, nil
	return nil, nil
}

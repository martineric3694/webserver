package main

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	Response InquiryNoTagihanResponse `xml:"inquiryNoTagihanResponse"`
}

type InquiryNoTagihanResponse struct {
	Return Return `xml:"return"`
}

type Return struct {
	AddValues      string `xml:"addValues"`
	Kode           string `xml:"kode"`
	Msg            string `xml:"msg"`
	Ret            string `xml:"ret"`
	BlnTagihan     string `xml:"blnTagihan"`
	IuranJHT       string `xml:"iuranJHT"`
	IuranJKK       string `xml:"iuranJKK"`
	IuranJKM       string `xml:"iuranJKM"`
	IuranJPK       string `xml:"iuranJPK"`
	IuranJPN       string `xml:"iuranJPN"`
	KodeDivisi     string `xml:"kodeDivisi"`
	NamaPerusahaan string `xml:"namaPerusahaan"`
	NoTagihan      string `xml:"noTagihan"`
	Npp            string `xml:"npp"`
	TotalBPJSK     string `xml:"totalBPJSK"`
	TotalBPJSTK    string `xml:"totalBPJSTK"`
	TotalIuran     string `xml:"totalIuran"`
}

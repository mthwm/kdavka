package davka16

import "time"

// Header represents the "D" line (batch header).
type Header struct {
	Char  string  // CHAR (e.g., "P" for original)
	Dtyp  string  // DTYP = "16"
	Ico   string  // DICO (IČZ, 8 chars)
	Pob   string  // DPOB (4 chars)
	Rok   int     // DROK (YYYY)
	Mes   int     // DMES (MM)
	Cid   int     // DCID (6 digits)
	Poc   int     // DPOC (num docs, 3 digits; auto-set if 0)
	Body  int     // DBODY (points, optional -> 0 if unset)
	Fin   float64 // DFIN (total CZK, 18.2, optional -> 0.00)
	Dpp   string  // DDPP (e.g., "1")
	Dvdr1 string  // DVDR1 (version, e.g., "16:6.2.XXX")
	Dvdr2 string  // DVDR2 (empty for simple davka)
	Ddtyp string  // DDTYP (empty)
}

// Document represents a single settlement document, consisting of one "L" line,
// at least one Naklad ("U"), and optional Sdeleni ("S").
type Document struct {
	L        DocumentL  // Main "L" line
	Naklads  []NakladU  // "U" lines (min 1, max 50)
	Sdelenis []SdeleniS // "S" lines (max 20)
}

// DocumentL is the "L" line structure.
type DocumentL struct {
	Icll    int       // ICLL (3 digits)
	Cdok    int       // CDOK (7 digits)
	Ind1    string    // IND_1 (9 chars)
	Cop     string    // COP (4 chars)
	TypLp   string    // TYP_LP (K/P/D)
	Jmeno   string    // JMENO (30 chars)
	Cp      string    // CP (10 chars)
	JmenoPr string    // JMENO_PR (30 chars)
	CpPr    string    // CP_PR (10 chars)
	Dnast   time.Time // DNAST (date)
	Dukon   time.Time // DUKON (date)
	Dodj    time.Time // DODJ (date)
	Jmevyst string    // JMEVYST (30 chars)
	Dvyst   time.Time // DVYST (date)
	Prod    int       // PROD (3 digits)
	KodUko  string    // KOD_UKO (1 char)
	CenaPob float64   // CENA_POB (10.2, optional)
}

// NakladU is a "U" line (cost item).
type NakladU struct {
	Datod     time.Time // DATOD (date)
	KodNak    string    // KOD_NAK (1 char)
	KodNak1   string    // KOD_NAK1 (1 char, reserve empty)
	Doba      int       // DOBA (3 digits)
	Sazba     float64   // SAZBA (7.2)
	Cena      float64   // CENA (9.2, optional)
	Luzko     int       // LUZKO (1 digit, for accom.)
	Kateg     string    // KATEG (3 chars, for accom.)
	KodPrerus string    // KOD_PRERUS (1 char)
}

// SdeleniS is an "S" line (note).
type SdeleniS struct {
	CisR int    // CIS_R (2 digits)
	Text string // TEXT (80 chars)
}

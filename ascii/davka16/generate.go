package davka16

import (
	"strings"
)

// generateHeader builds the "D" line (length 88).
func generateHeader(h Header) string {
	var sb strings.Builder
	sb.WriteString("D")                       // TYP
	sb.WriteString(formatC(h.Char, 1))        // CHAR
	sb.WriteString(formatC(h.Dtyp, 2))        // DTYP
	sb.WriteString(formatC(h.Ico, 8))         // DICO
	sb.WriteString(formatC(h.Pob, 4))         // DPOB
	sb.WriteString(formatN(h.Rok, 4))         // DROK
	sb.WriteString(formatN(h.Mes, 2))         // DMES
	sb.WriteString(formatN(h.Cid, 6))         // DCID
	sb.WriteString(formatN(h.Poc, 3))         // DPOC
	sb.WriteString(formatN(h.Body, 11))       // DBODY
	sb.WriteString(formatMoney(h.Fin, 18, 2)) // DFIN
	sb.WriteString(formatC(h.Dpp, 1))         // DDPP
	sb.WriteString(formatC(h.Dvdr1, 13))      // DVDR1
	sb.WriteString(formatC(h.Dvdr2, 13))      // DVDR2
	sb.WriteString(formatC(h.Ddtyp, 1))       // DDTYP
	return sb.String()
}

// generateL builds the "L" line (length 182).
func generateL(l DocumentL) string {
	var sb strings.Builder
	sb.WriteString("L")                           // TYP
	sb.WriteString(" ")                           // DTYP reserve
	sb.WriteString(formatN(l.Icll, 3))            // ICLL
	sb.WriteString(formatN(l.Cdok, 7))            // CDOK
	sb.WriteString(formatC(l.Ind1, 9))            // IND_1
	sb.WriteString(formatC(l.Cop, 4))             // COP
	sb.WriteString(formatC(l.TypLp, 1))           // TYP_LP
	sb.WriteString(formatC(l.Jmeno, 30))          // JMENO
	sb.WriteString(formatC(l.Cp, 10))             // CP
	sb.WriteString(formatC(l.JmenoPr, 30))        // JMENO_PR
	sb.WriteString(formatC(l.CpPr, 10))           // CP_PR
	sb.WriteString(formatD(l.Dnast))              // DNAST
	sb.WriteString(formatD(l.Dukon))              // DUKON
	sb.WriteString(formatD(l.Dodj))               // DODJ
	sb.WriteString(formatC(l.Jmevyst, 30))        // JMEVYST
	sb.WriteString(formatD(l.Dvyst))              // DVYST
	sb.WriteString(formatN(l.Prod, 3))            // PROD
	sb.WriteString(formatC(l.KodUko, 1))          // KOD_UKO
	sb.WriteString(formatMoney(l.CenaPob, 10, 2)) // CENA_POB
	return sb.String()
}

// generateU builds a "U" line (length 36).
func generateU(u NakladU) string {
	var sb strings.Builder
	sb.WriteString("U")                        // TYP
	sb.WriteString(" ")                        // DTYP reserve
	sb.WriteString(formatD(u.Datod))           // DATOD
	sb.WriteString(formatC(u.KodNak, 1))       // KOD_NAK
	sb.WriteString(formatC(u.KodNak1, 1))      // KOD_NAK1
	sb.WriteString(formatN(u.Doba, 3))         // DOBA
	sb.WriteString(formatMoney(u.Sazba, 7, 2)) // SAZBA
	sb.WriteString(formatMoney(u.Cena, 9, 2))  // CENA
	sb.WriteString(formatN(u.Luzko, 1))        // LUZKO
	sb.WriteString(formatC(u.Kateg, 3))        // KATEG
	sb.WriteString(formatC(u.KodPrerus, 1))    // KOD_PRERUS
	return sb.String()
}

// generateS builds an "S" line (length 84).
func generateS(s SdeleniS) string {
	var sb strings.Builder
	sb.WriteString("S")                 // TYP
	sb.WriteString(" ")                 // STYP reserve
	sb.WriteString(formatN(s.CisR, 2))  // CIS_R
	sb.WriteString(formatC(s.Text, 80)) // TEXT
	return sb.String()
}

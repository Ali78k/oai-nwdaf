/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type NfTypeAnyOf string

// List of NfTypeAnyOf
const (
	NFTYPEANYOF_NRF        NfTypeAnyOf = "NRF"
	NFTYPEANYOF_UDM        NfTypeAnyOf = "UDM"
	NFTYPEANYOF_AMF        NfTypeAnyOf = "AMF"
	NFTYPEANYOF_SMF        NfTypeAnyOf = "SMF"
	NFTYPEANYOF_AUSF       NfTypeAnyOf = "AUSF"
	NFTYPEANYOF_NEF        NfTypeAnyOf = "NEF"
	NFTYPEANYOF_PCF        NfTypeAnyOf = "PCF"
	NFTYPEANYOF_SMSF       NfTypeAnyOf = "SMSF"
	NFTYPEANYOF_NSSF       NfTypeAnyOf = "NSSF"
	NFTYPEANYOF_UDR        NfTypeAnyOf = "UDR"
	NFTYPEANYOF_LMF        NfTypeAnyOf = "LMF"
	NFTYPEANYOF_GMLC       NfTypeAnyOf = "GMLC"
	NFTYPEANYOF__5_G_EIR   NfTypeAnyOf = "5G_EIR"
	NFTYPEANYOF_SEPP       NfTypeAnyOf = "SEPP"
	NFTYPEANYOF_UPF        NfTypeAnyOf = "UPF"
	NFTYPEANYOF_N3_IWF     NfTypeAnyOf = "N3IWF"
	NFTYPEANYOF_AF         NfTypeAnyOf = "AF"
	NFTYPEANYOF_UDSF       NfTypeAnyOf = "UDSF"
	NFTYPEANYOF_BSF        NfTypeAnyOf = "BSF"
	NFTYPEANYOF_CHF        NfTypeAnyOf = "CHF"
	NFTYPEANYOF_NWDAF      NfTypeAnyOf = "NWDAF"
	NFTYPEANYOF_PCSCF      NfTypeAnyOf = "PCSCF"
	NFTYPEANYOF_CBCF       NfTypeAnyOf = "CBCF"
	NFTYPEANYOF_HSS        NfTypeAnyOf = "HSS"
	NFTYPEANYOF_UCMF       NfTypeAnyOf = "UCMF"
	NFTYPEANYOF_SOR_AF     NfTypeAnyOf = "SOR_AF"
	NFTYPEANYOF_SPAF       NfTypeAnyOf = "SPAF"
	NFTYPEANYOF_MME        NfTypeAnyOf = "MME"
	NFTYPEANYOF_SCSAS      NfTypeAnyOf = "SCSAS"
	NFTYPEANYOF_SCEF       NfTypeAnyOf = "SCEF"
	NFTYPEANYOF_SCP        NfTypeAnyOf = "SCP"
	NFTYPEANYOF_NSSAAF     NfTypeAnyOf = "NSSAAF"
	NFTYPEANYOF_ICSCF      NfTypeAnyOf = "ICSCF"
	NFTYPEANYOF_SCSCF      NfTypeAnyOf = "SCSCF"
	NFTYPEANYOF_DRA        NfTypeAnyOf = "DRA"
	NFTYPEANYOF_IMS_AS     NfTypeAnyOf = "IMS_AS"
	NFTYPEANYOF_AANF       NfTypeAnyOf = "AANF"
	NFTYPEANYOF__5_G_DDNMF NfTypeAnyOf = "5G_DDNMF"
	NFTYPEANYOF_NSACF      NfTypeAnyOf = "NSACF"
	NFTYPEANYOF_MFAF       NfTypeAnyOf = "MFAF"
	NFTYPEANYOF_EASDF      NfTypeAnyOf = "EASDF"
	NFTYPEANYOF_DCCF       NfTypeAnyOf = "DCCF"
	NFTYPEANYOF_MB_SMF     NfTypeAnyOf = "MB-SMF"
	NFTYPEANYOF_TSCTSF     NfTypeAnyOf = "TSCTSF"
	NFTYPEANYOF_ADRF       NfTypeAnyOf = "ADRF"
	NFTYPEANYOF_GBA_BSF    NfTypeAnyOf = "GBA_BSF"
	NFTYPEANYOF_CEF        NfTypeAnyOf = "CEF"
	NFTYPEANYOF_MB_UPF     NfTypeAnyOf = "MB-UPF"
	NFTYPEANYOF_NSWOF      NfTypeAnyOf = "NSWOF"
)

// AssertNfTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertNfTypeAnyOfRequired(obj NfTypeAnyOf) error {
	return nil
}

// AssertRecurseNfTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NfTypeAnyOf (e.g. [][]NfTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNfTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNfTypeAnyOf, ok := obj.(NfTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNfTypeAnyOfRequired(aNfTypeAnyOf)
	})
}

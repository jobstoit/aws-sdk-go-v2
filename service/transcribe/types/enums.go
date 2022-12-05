// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BaseModelName string

// Enum values for BaseModelName
const (
	BaseModelNameNarrowBand BaseModelName = "NarrowBand"
	BaseModelNameWideBand   BaseModelName = "WideBand"
)

// Values returns all known values for BaseModelName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BaseModelName) Values() []BaseModelName {
	return []BaseModelName{
		"NarrowBand",
		"WideBand",
	}
}

type CallAnalyticsJobStatus string

// Enum values for CallAnalyticsJobStatus
const (
	CallAnalyticsJobStatusQueued     CallAnalyticsJobStatus = "QUEUED"
	CallAnalyticsJobStatusInProgress CallAnalyticsJobStatus = "IN_PROGRESS"
	CallAnalyticsJobStatusFailed     CallAnalyticsJobStatus = "FAILED"
	CallAnalyticsJobStatusCompleted  CallAnalyticsJobStatus = "COMPLETED"
)

// Values returns all known values for CallAnalyticsJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CallAnalyticsJobStatus) Values() []CallAnalyticsJobStatus {
	return []CallAnalyticsJobStatus{
		"QUEUED",
		"IN_PROGRESS",
		"FAILED",
		"COMPLETED",
	}
}

type CLMLanguageCode string

// Enum values for CLMLanguageCode
const (
	CLMLanguageCodeEnUs CLMLanguageCode = "en-US"
	CLMLanguageCodeHiIn CLMLanguageCode = "hi-IN"
	CLMLanguageCodeEsUs CLMLanguageCode = "es-US"
	CLMLanguageCodeEnGb CLMLanguageCode = "en-GB"
	CLMLanguageCodeEnAu CLMLanguageCode = "en-AU"
	CLMLanguageCodeDeDe CLMLanguageCode = "de-DE"
	CLMLanguageCodeJaJp CLMLanguageCode = "ja-JP"
)

// Values returns all known values for CLMLanguageCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CLMLanguageCode) Values() []CLMLanguageCode {
	return []CLMLanguageCode{
		"en-US",
		"hi-IN",
		"es-US",
		"en-GB",
		"en-AU",
		"de-DE",
		"ja-JP",
	}
}

type InputType string

// Enum values for InputType
const (
	InputTypeRealTime InputType = "REAL_TIME"
	InputTypePostCall InputType = "POST_CALL"
)

// Values returns all known values for InputType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (InputType) Values() []InputType {
	return []InputType{
		"REAL_TIME",
		"POST_CALL",
	}
}

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeAfZa LanguageCode = "af-ZA"
	LanguageCodeArAe LanguageCode = "ar-AE"
	LanguageCodeArSa LanguageCode = "ar-SA"
	LanguageCodeDaDk LanguageCode = "da-DK"
	LanguageCodeDeCh LanguageCode = "de-CH"
	LanguageCodeDeDe LanguageCode = "de-DE"
	LanguageCodeEnAb LanguageCode = "en-AB"
	LanguageCodeEnAu LanguageCode = "en-AU"
	LanguageCodeEnGb LanguageCode = "en-GB"
	LanguageCodeEnIe LanguageCode = "en-IE"
	LanguageCodeEnIn LanguageCode = "en-IN"
	LanguageCodeEnUs LanguageCode = "en-US"
	LanguageCodeEnWl LanguageCode = "en-WL"
	LanguageCodeEsEs LanguageCode = "es-ES"
	LanguageCodeEsUs LanguageCode = "es-US"
	LanguageCodeFaIr LanguageCode = "fa-IR"
	LanguageCodeFrCa LanguageCode = "fr-CA"
	LanguageCodeFrFr LanguageCode = "fr-FR"
	LanguageCodeHeIl LanguageCode = "he-IL"
	LanguageCodeHiIn LanguageCode = "hi-IN"
	LanguageCodeIdId LanguageCode = "id-ID"
	LanguageCodeItIt LanguageCode = "it-IT"
	LanguageCodeJaJp LanguageCode = "ja-JP"
	LanguageCodeKoKr LanguageCode = "ko-KR"
	LanguageCodeMsMy LanguageCode = "ms-MY"
	LanguageCodeNlNl LanguageCode = "nl-NL"
	LanguageCodePtBr LanguageCode = "pt-BR"
	LanguageCodePtPt LanguageCode = "pt-PT"
	LanguageCodeRuRu LanguageCode = "ru-RU"
	LanguageCodeTaIn LanguageCode = "ta-IN"
	LanguageCodeTeIn LanguageCode = "te-IN"
	LanguageCodeTrTr LanguageCode = "tr-TR"
	LanguageCodeZhCn LanguageCode = "zh-CN"
	LanguageCodeZhTw LanguageCode = "zh-TW"
	LanguageCodeThTh LanguageCode = "th-TH"
	LanguageCodeEnZa LanguageCode = "en-ZA"
	LanguageCodeEnNz LanguageCode = "en-NZ"
)

// Values returns all known values for LanguageCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LanguageCode) Values() []LanguageCode {
	return []LanguageCode{
		"af-ZA",
		"ar-AE",
		"ar-SA",
		"da-DK",
		"de-CH",
		"de-DE",
		"en-AB",
		"en-AU",
		"en-GB",
		"en-IE",
		"en-IN",
		"en-US",
		"en-WL",
		"es-ES",
		"es-US",
		"fa-IR",
		"fr-CA",
		"fr-FR",
		"he-IL",
		"hi-IN",
		"id-ID",
		"it-IT",
		"ja-JP",
		"ko-KR",
		"ms-MY",
		"nl-NL",
		"pt-BR",
		"pt-PT",
		"ru-RU",
		"ta-IN",
		"te-IN",
		"tr-TR",
		"zh-CN",
		"zh-TW",
		"th-TH",
		"en-ZA",
		"en-NZ",
	}
}

type MediaFormat string

// Enum values for MediaFormat
const (
	MediaFormatMp3  MediaFormat = "mp3"
	MediaFormatMp4  MediaFormat = "mp4"
	MediaFormatWav  MediaFormat = "wav"
	MediaFormatFlac MediaFormat = "flac"
	MediaFormatOgg  MediaFormat = "ogg"
	MediaFormatAmr  MediaFormat = "amr"
	MediaFormatWebm MediaFormat = "webm"
)

// Values returns all known values for MediaFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MediaFormat) Values() []MediaFormat {
	return []MediaFormat{
		"mp3",
		"mp4",
		"wav",
		"flac",
		"ogg",
		"amr",
		"webm",
	}
}

type MedicalContentIdentificationType string

// Enum values for MedicalContentIdentificationType
const (
	MedicalContentIdentificationTypePhi MedicalContentIdentificationType = "PHI"
)

// Values returns all known values for MedicalContentIdentificationType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (MedicalContentIdentificationType) Values() []MedicalContentIdentificationType {
	return []MedicalContentIdentificationType{
		"PHI",
	}
}

type ModelStatus string

// Enum values for ModelStatus
const (
	ModelStatusInProgress ModelStatus = "IN_PROGRESS"
	ModelStatusFailed     ModelStatus = "FAILED"
	ModelStatusCompleted  ModelStatus = "COMPLETED"
)

// Values returns all known values for ModelStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ModelStatus) Values() []ModelStatus {
	return []ModelStatus{
		"IN_PROGRESS",
		"FAILED",
		"COMPLETED",
	}
}

type OutputLocationType string

// Enum values for OutputLocationType
const (
	OutputLocationTypeCustomerBucket OutputLocationType = "CUSTOMER_BUCKET"
	OutputLocationTypeServiceBucket  OutputLocationType = "SERVICE_BUCKET"
)

// Values returns all known values for OutputLocationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OutputLocationType) Values() []OutputLocationType {
	return []OutputLocationType{
		"CUSTOMER_BUCKET",
		"SERVICE_BUCKET",
	}
}

type ParticipantRole string

// Enum values for ParticipantRole
const (
	ParticipantRoleAgent    ParticipantRole = "AGENT"
	ParticipantRoleCustomer ParticipantRole = "CUSTOMER"
)

// Values returns all known values for ParticipantRole. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParticipantRole) Values() []ParticipantRole {
	return []ParticipantRole{
		"AGENT",
		"CUSTOMER",
	}
}

type PiiEntityType string

// Enum values for PiiEntityType
const (
	PiiEntityTypeBankAccountNumber PiiEntityType = "BANK_ACCOUNT_NUMBER"
	PiiEntityTypeBankRouting       PiiEntityType = "BANK_ROUTING"
	PiiEntityTypeCreditDebitNumber PiiEntityType = "CREDIT_DEBIT_NUMBER"
	PiiEntityTypeCreditDebitCvv    PiiEntityType = "CREDIT_DEBIT_CVV"
	PiiEntityTypeCreditDebitExpiry PiiEntityType = "CREDIT_DEBIT_EXPIRY"
	PiiEntityTypePin               PiiEntityType = "PIN"
	PiiEntityTypeEmail             PiiEntityType = "EMAIL"
	PiiEntityTypeAddress           PiiEntityType = "ADDRESS"
	PiiEntityTypeName              PiiEntityType = "NAME"
	PiiEntityTypePhone             PiiEntityType = "PHONE"
	PiiEntityTypeSsn               PiiEntityType = "SSN"
	PiiEntityTypeAll               PiiEntityType = "ALL"
)

// Values returns all known values for PiiEntityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PiiEntityType) Values() []PiiEntityType {
	return []PiiEntityType{
		"BANK_ACCOUNT_NUMBER",
		"BANK_ROUTING",
		"CREDIT_DEBIT_NUMBER",
		"CREDIT_DEBIT_CVV",
		"CREDIT_DEBIT_EXPIRY",
		"PIN",
		"EMAIL",
		"ADDRESS",
		"NAME",
		"PHONE",
		"SSN",
		"ALL",
	}
}

type RedactionOutput string

// Enum values for RedactionOutput
const (
	RedactionOutputRedacted              RedactionOutput = "redacted"
	RedactionOutputRedactedAndUnredacted RedactionOutput = "redacted_and_unredacted"
)

// Values returns all known values for RedactionOutput. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RedactionOutput) Values() []RedactionOutput {
	return []RedactionOutput{
		"redacted",
		"redacted_and_unredacted",
	}
}

type RedactionType string

// Enum values for RedactionType
const (
	RedactionTypePii RedactionType = "PII"
)

// Values returns all known values for RedactionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RedactionType) Values() []RedactionType {
	return []RedactionType{
		"PII",
	}
}

type SentimentValue string

// Enum values for SentimentValue
const (
	SentimentValuePositive SentimentValue = "POSITIVE"
	SentimentValueNegative SentimentValue = "NEGATIVE"
	SentimentValueNeutral  SentimentValue = "NEUTRAL"
	SentimentValueMixed    SentimentValue = "MIXED"
)

// Values returns all known values for SentimentValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SentimentValue) Values() []SentimentValue {
	return []SentimentValue{
		"POSITIVE",
		"NEGATIVE",
		"NEUTRAL",
		"MIXED",
	}
}

type Specialty string

// Enum values for Specialty
const (
	SpecialtyPrimarycare Specialty = "PRIMARYCARE"
)

// Values returns all known values for Specialty. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Specialty) Values() []Specialty {
	return []Specialty{
		"PRIMARYCARE",
	}
}

type SubtitleFormat string

// Enum values for SubtitleFormat
const (
	SubtitleFormatVtt SubtitleFormat = "vtt"
	SubtitleFormatSrt SubtitleFormat = "srt"
)

// Values returns all known values for SubtitleFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubtitleFormat) Values() []SubtitleFormat {
	return []SubtitleFormat{
		"vtt",
		"srt",
	}
}

type TranscriptFilterType string

// Enum values for TranscriptFilterType
const (
	TranscriptFilterTypeExact TranscriptFilterType = "EXACT"
)

// Values returns all known values for TranscriptFilterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscriptFilterType) Values() []TranscriptFilterType {
	return []TranscriptFilterType{
		"EXACT",
	}
}

type TranscriptionJobStatus string

// Enum values for TranscriptionJobStatus
const (
	TranscriptionJobStatusQueued     TranscriptionJobStatus = "QUEUED"
	TranscriptionJobStatusInProgress TranscriptionJobStatus = "IN_PROGRESS"
	TranscriptionJobStatusFailed     TranscriptionJobStatus = "FAILED"
	TranscriptionJobStatusCompleted  TranscriptionJobStatus = "COMPLETED"
)

// Values returns all known values for TranscriptionJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscriptionJobStatus) Values() []TranscriptionJobStatus {
	return []TranscriptionJobStatus{
		"QUEUED",
		"IN_PROGRESS",
		"FAILED",
		"COMPLETED",
	}
}

type Type string

// Enum values for Type
const (
	TypeConversation Type = "CONVERSATION"
	TypeDictation    Type = "DICTATION"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"CONVERSATION",
		"DICTATION",
	}
}

type VocabularyFilterMethod string

// Enum values for VocabularyFilterMethod
const (
	VocabularyFilterMethodRemove VocabularyFilterMethod = "remove"
	VocabularyFilterMethodMask   VocabularyFilterMethod = "mask"
	VocabularyFilterMethodTag    VocabularyFilterMethod = "tag"
)

// Values returns all known values for VocabularyFilterMethod. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VocabularyFilterMethod) Values() []VocabularyFilterMethod {
	return []VocabularyFilterMethod{
		"remove",
		"mask",
		"tag",
	}
}

type VocabularyState string

// Enum values for VocabularyState
const (
	VocabularyStatePending VocabularyState = "PENDING"
	VocabularyStateReady   VocabularyState = "READY"
	VocabularyStateFailed  VocabularyState = "FAILED"
)

// Values returns all known values for VocabularyState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VocabularyState) Values() []VocabularyState {
	return []VocabularyState{
		"PENDING",
		"READY",
		"FAILED",
	}
}

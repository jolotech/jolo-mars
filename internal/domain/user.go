package domain


import (
    "time"
    "gorm.io/datatypes"
)

type User struct {
    ID                 uint           `json:"id" gorm:"primaryKey"`
    PartnerName        string         `json:"partner_name"`
    Email              string         `json:"email" gorm:"unique;not null"`
    Credentials        datatypes.JSON `json:"credentials" gorm:"type:json"`       
    WebhookURL         string         `json:"webhook_url"`
    Balance            float64        `json:"balance" gorm:"default:0"`
    WhitelistedIPs     datatypes.JSON `json:"whitelisted_ips" gorm:"type:json"`
    PasswordHash       string         `json:"-" gorm:"not null"`
    AutoDebitCard      string         `json:"autoDebitCard"`    // Paystack authorization code
    AutoTopUpAmount    float64        `json:"autoTopUpAmount"`
    StoreID            uint           `json:"storeId" gorm:"default:0"`
    WebhookEnabled     bool           `json:"webhook_enabled" gorum:"default:'false'"`
    AutoDebitEnabled   bool           `json:"autoDebitEnabled" gorum:"default:'false'"`
    AutoDebitLimit     float64        `json:"autoDebitLimit"` // threshold for triggering auto-topup
    Mode               string         `json:"mode" gorm:"default:'test'"`
    MustChangePassword bool           `json:"must_change_password" gorm:"default:true"`
    CreatedAt          time.Time      `json:"created_at"`
    UpdatedAt          time.Time      `json:"updated_at"`
}

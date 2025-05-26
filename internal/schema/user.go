package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	COLLECTION_USERS = "users"
)

type USER_STATUS string
type DELETION_REASON string

const (
	// User statuses
	USER_STATUS_ACTIVE     USER_STATUS = "active"     // Normal active account
	USER_STATUS_INACTIVE   USER_STATUS = "inactive"   // Temporarily inactive
	USER_STATUS_PENDING    USER_STATUS = "pending"    // Pending verification
	USER_STATUS_SUSPENDED  USER_STATUS = "suspended"  // Suspended account
	USER_STATUS_DELETED    USER_STATUS = "deleted"    // Account deleted
	USER_STATUS_ANONYMIZED USER_STATUS = "anonymized" // Account data anonymized

	// Deletion reasons
	DELETION_REASON_USER_REQUEST      DELETION_REASON = "user_request"      // User requested deletion
	DELETION_REASON_ADMIN_ACTION      DELETION_REASON = "admin_action"      // Deleted by admin
	DELETION_REASON_SYSTEM_ACTION     DELETION_REASON = "system_action"     // System-initiated deletion
	DELETION_REASON_LEGAL_REQUIREMENT DELETION_REASON = "legal_requirement" // Legal requirement
)

// User model for authentication and authorization
type User struct {
	ID        bson.ObjectID  `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at" json:"updated_at"`
	CreatedBy *bson.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"` // Who created this account (if admin-created)

	// Basic user information
	Username    string `bson:"username,omitempty" json:"username,omitempty"`         // Unique username
	DisplayName string `bson:"display_name" json:"display_name"`                     // Public display name
	Email       string `bson:"email" json:"email"`                                   // Primary email
	PhoneNumber string `bson:"phone_number,omitempty" json:"phone_number,omitempty"` // Phone number
	Password    string `bson:"password" json:"-"`                                    // Password hash
	AvatarURL   string `bson:"avatar_url,omitempty" json:"avatar_url,omitempty"`     // Profile picture URL
	Locale      string `bson:"locale" json:"locale"`                                 // User's locale (e.g., en-US, fr-CA)
	TimeZone    string `bson:"timezone,omitempty" json:"timezone,omitempty"`         // User's timezone (e.g., America/New_York, Europe/Paris)
	AccountType string `bson:"account_type" json:"account_type"`                     // Type of account (e.g., "free", "premium", "business", "enterprise", etc.)

	// Authentication and verification information
	AuthInfo AuthInfo `bson:"auth_info" json:"auth_info"`

	// Activity tracking
	LastLoginAt     *time.Time `bson:"last_login_at,omitempty" json:"last_login_at,omitempty"`         // Last successful login
	LastActivityAt  *time.Time `bson:"last_activity_at,omitempty" json:"last_activity_at,omitempty"`   // Last user activity
	LastEmailChange *time.Time `bson:"last_email_change,omitempty" json:"last_email_change,omitempty"` // Last email change
	LastPhoneChange *time.Time `bson:"last_phone_change,omitempty" json:"last_phone_change,omitempty"` // Last phone change

	// User status
	Status USER_STATUS `bson:"status" json:"status"` // Current account status

	// Suspension information
	Suspension SuspensionInfo `bson:"suspension" json:"suspension"` // Account suspension status

	// GDPR Deletion tracking
	DeletionInfo *DeletionInfo `bson:"deletion_info,omitempty" json:"deletion_info,omitempty"` // Account deletion information
}

// DeletionInfo tracks GDPR-compliant account deletion
type DeletionInfo struct {
	DeletedAt      time.Time       `bson:"deleted_at" json:"deleted_at"`                               // When the account was marked for deletion
	AnonymizedAt   *time.Time      `bson:"anonymized_at,omitempty" json:"anonymized_at,omitempty"`     // When the data was anonymized
	Reason         DELETION_REASON `bson:"reason" json:"reason"`                                       // Why the account was deleted
	RequestedBy    string          `bson:"requested_by" json:"requested_by"`                           // Who requested the deletion (user_id or system)
	AnonymizedBy   string          `bson:"anonymized_by,omitempty" json:"anonymized_by,omitempty"`     // Who performed the anonymization
	RetentionUntil *time.Time      `bson:"retention_until,omitempty" json:"retention_until,omitempty"` // Legal retention period if applicable
}

// AuthInfo groups fields for password, token, and OTP management
type AuthInfo struct {
	// Password management
	PasswordResetToken  string     `bson:"password_reset_token,omitempty" json:"-"`
	PasswordResetSentAt *time.Time `bson:"password_reset_sent_at,omitempty" json:"-"`
	LastPasswordChange  time.Time  `bson:"last_password_change" json:"-"`
	LastPasswordReset   *time.Time `bson:"last_password_reset,omitempty" json:"-"`

	// Verification status
	EmailVerified           bool       `bson:"email_verified" json:"email_verified"`
	PhoneVerified           bool       `bson:"phone_verified" json:"phone_verified"`
	EmailVerificationToken  string     `bson:"email_verification_token,omitempty" json:"-"`
	EmailVerificationSentAt *time.Time `bson:"email_verification_sent_at,omitempty" json:"-"`
	PhoneVerificationToken  string     `bson:"phone_verification_token,omitempty" json:"-"`
	PhoneVerificationSentAt *time.Time `bson:"phone_verification_sent_at,omitempty" json:"-"`

	// OTP/TOTP details for two-factor authentication
	TOTPSecret      string     `bson:"totp_secret,omitempty" json:"-"`
	Is2FAEnabled    bool       `bson:"is_2fa_enabled" json:"is_2fa_enabled"`
	OTPBackupCodes  []string   `bson:"otp_backup_codes,omitempty" json:"-"`
	Last2FAVerified *time.Time `bson:"last_2fa_verified,omitempty" json:"-"`

	// OAuth providers
	OAuthProviders map[string]OAuthProvider `bson:"oauth_providers,omitempty" json:"oauth_providers,omitempty"`
}

// OAuthProvider represents a connected OAuth account
type OAuthProvider struct {
	ProviderID       string    `bson:"provider_id" json:"provider_id"`                                 // ID from the OAuth provider
	ProviderEmail    string    `bson:"provider_email" json:"provider_email"`                           // Email from the OAuth provider
	ProviderUsername string    `bson:"provider_username,omitempty" json:"provider_username,omitempty"` // Username from provider
	ProviderAvatar   string    `bson:"provider_avatar,omitempty" json:"provider_avatar,omitempty"`     // Avatar from provider
	ConnectedAt      time.Time `bson:"connected_at" json:"connected_at"`                               // When the account was connected
	LastUsedAt       time.Time `bson:"last_used_at" json:"last_used_at"`                               // When the account was last used
}

// SuspensionInfo handles account suspension status
type SuspensionInfo struct {
	IsSuspended     bool       `bson:"is_suspended" json:"is_suspended"`
	SuspendedAt     *time.Time `bson:"suspended_at,omitempty" json:"suspended_at,omitempty"`
	SuspendedUntil  *time.Time `bson:"suspended_until,omitempty" json:"suspended_until,omitempty"`
	SuspendedReason string     `bson:"suspended_reason,omitempty" json:"suspended_reason,omitempty"`
	SuspendedBy     string     `bson:"suspended_by,omitempty" json:"suspended_by,omitempty"`
}

package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	COLLECTION_LOGIN_HISTORY    = "login_history"
	COLLECTION_EMAIL_HISTORY    = "email_history"
	COLLECTION_ACCOUNT_HISTORY  = "account_history"
	COLLECTION_SECURITY_HISTORY = "security_history"
	COLLECTION_ADMIN_HISTORY    = "admin_history"

	// TTL values for MongoDB indexes (in seconds)
	TTL_LOGIN_HISTORY    = 90 * 24 * 60 * 60  // 90 days
	TTL_EMAIL_HISTORY    = 30 * 24 * 60 * 60  // 30 days
	TTL_ACCOUNT_HISTORY  = 365 * 24 * 60 * 60 // 1 year
	TTL_SECURITY_HISTORY = 365 * 24 * 60 * 60 // 1 year
	TTL_ADMIN_HISTORY    = 730 * 24 * 60 * 60 // 2 years
)

// Event type definitions
type (
	LoginEventType    string
	EmailEventType    string
	AccountEventType  string
	SecurityEventType string
	AdminEventType    string
)

// Login event types
const (
	LOGIN_EVENT_SUCCESS LoginEventType = "success" // Successful login
	LOGIN_EVENT_FAILED  LoginEventType = "failed"  // Failed login attempt
	LOGIN_EVENT_LOGOUT  LoginEventType = "logout"  // User logout
	LOGIN_EVENT_EXPIRED LoginEventType = "expired" // Session expired
	LOGIN_EVENT_REVOKED LoginEventType = "revoked" // Session revoked
)

// Account event types
const (
	ACCOUNT_EVENT_EMAIL_CHANGE    AccountEventType = "email_change"    // Email address change (e.g. old: "user@old.com" -> new: "user@new.com")
	ACCOUNT_EVENT_PHONE_CHANGE    AccountEventType = "phone_change"    // Phone number change (e.g. old: "+1234567890" -> new: "+0987654321")
	ACCOUNT_EVENT_PASSWORD_CHANGE AccountEventType = "password_change" // Password change (e.g. field: "password", old: "[REDACTED]", new: "[REDACTED]")
	ACCOUNT_EVENT_PROFILE_UPDATE  AccountEventType = "profile_update"  // Profile information update (e.g. field: "name", old: "John" -> new: "John Doe")
	ACCOUNT_EVENT_ACCOUNT_TYPE    AccountEventType = "account_type"    // Account type change (e.g. old: "free" -> new: "premium")
)

// Security event types
const (
	SECURITY_EVENT_2FA_ENABLE     SecurityEventType = "2fa_enable"     // 2FA enabled
	SECURITY_EVENT_2FA_DISABLE    SecurityEventType = "2fa_disable"    // 2FA disabled
	SECURITY_EVENT_OAUTH_LINK     SecurityEventType = "oauth_link"     // OAuth account linked
	SECURITY_EVENT_OAUTH_UNLINK   SecurityEventType = "oauth_unlink"   // OAuth account unlinked
	SECURITY_EVENT_PASSWORD_RESET SecurityEventType = "password_reset" // Password reset
)

// Admin event types
const (
	ADMIN_EVENT_SUSPEND     AdminEventType = "suspend"     // Account suspension
	ADMIN_EVENT_UNSUSPEND   AdminEventType = "unsuspend"   // Account unsuspension
	ADMIN_EVENT_DELETE      AdminEventType = "delete"      // Account deletion
	ADMIN_EVENT_ANONYMIZE   AdminEventType = "anonymize"   // Account anonymization
	ADMIN_EVENT_ROLE_CHANGE AdminEventType = "role_change" // Role/permission change
)

// LoginHistory model to track user login activity
type LoginHistory struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"` // Used for TTL index

	UserID    bson.ObjectID  `bson:"user_id" json:"user_id"`                 // Reference to User model
	EventType LoginEventType `bson:"event_type" json:"event_type"`           // Type of login event
	IPAddress string         `bson:"ip_address" json:"ip_address"`           // IP address of the user
	Country   string         `bson:"country" json:"country"`                 // Country code (e.g. "US", "GB")
	UserAgent string         `bson:"user_agent" json:"user_agent"`           // User agent string
	Success   bool           `bson:"success" json:"success"`                 // Whether login was successful
	Error     string         `bson:"error,omitempty" json:"error,omitempty"` // Error message if failed
}

// EmailHistory model to track email sending attempts
type EmailHistory struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"` // Used for TTL index

	UserID    bson.ObjectID  `bson:"user_id,omitempty" json:"user_id,omitempty"` // Reference to User model (if applicable)
	EmailType EmailEventType `bson:"email_type" json:"email_type"`               // Type of email (verification, reset, etc.)
	To        string         `bson:"to" json:"to"`                               // Recipient email
	Subject   string         `bson:"subject" json:"subject"`                     // Email subject
	Success   bool           `bson:"success" json:"success"`                     // Whether email was sent successfully
	Error     string         `bson:"error,omitempty" json:"error,omitempty"`     // Error message if failed
}

// AccountHistory model to track account changes
type AccountHistory struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"` // Used for TTL index

	UserID    bson.ObjectID    `bson:"user_id" json:"user_id"`                           // Reference to User model
	EventType AccountEventType `bson:"event_type" json:"event_type"`                     // Type of account event
	Field     string           `bson:"field" json:"field"`                               // Field that was changed
	OldValue  string           `bson:"old_value,omitempty" json:"old_value,omitempty"`   // Previous value
	NewValue  string           `bson:"new_value,omitempty" json:"new_value,omitempty"`   // New value
	ChangedBy string           `bson:"changed_by" json:"changed_by"`                     // Who made the change (user_id or system)
	IPAddress string           `bson:"ip_address,omitempty" json:"ip_address,omitempty"` // IP address of the change
	Country   string           `bson:"country,omitempty" json:"country,omitempty"`       // Country code (e.g. "US", "GB")
	UserAgent string           `bson:"user_agent,omitempty" json:"user_agent,omitempty"` // User agent of the change
}

// SecurityHistory model to track security-related events
type SecurityHistory struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"` // Used for TTL index

	UserID    bson.ObjectID     `bson:"user_id" json:"user_id"`                       // Reference to User model
	EventType SecurityEventType `bson:"event_type" json:"event_type"`                 // Type of security event
	Provider  string            `bson:"provider,omitempty" json:"provider,omitempty"` // OAuth provider (if applicable)
	IPAddress string            `bson:"ip_address" json:"ip_address"`                 // IP address of the event
	Country   string            `bson:"country" json:"country"`                       // Country code (e.g. "US", "GB")
	UserAgent string            `bson:"user_agent" json:"user_agent"`                 // User agent string
	Success   bool              `bson:"success" json:"success"`                       // Whether the action was successful
	Error     string            `bson:"error,omitempty" json:"error,omitempty"`       // Error message if failed
}

// AdminHistory model to track administrative actions
type AdminHistory struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"` // Used for TTL index

	AdminID   bson.ObjectID  `bson:"admin_id" json:"admin_id"`                         // Reference to admin User model
	UserID    bson.ObjectID  `bson:"user_id" json:"user_id"`                           // Reference to affected User model
	EventType AdminEventType `bson:"event_type" json:"event_type"`                     // Type of admin event
	Action    string         `bson:"action" json:"action"`                             // Action taken
	Reason    string         `bson:"reason,omitempty" json:"reason,omitempty"`         // Reason for the action
	Details   string         `bson:"details,omitempty" json:"details,omitempty"`       // Additional details
	IPAddress string         `bson:"ip_address" json:"ip_address"`                     // IP address of the admin
	Country   string         `bson:"country" json:"country"`                           // Country code (e.g. "US", "GB")
	UserAgent string         `bson:"user_agent" json:"user_agent"`                     // User agent string
	ExpiresAt *time.Time     `bson:"expires_at,omitempty" json:"expires_at,omitempty"` // When the action expires (if applicable)
}

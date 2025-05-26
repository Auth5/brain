# Auth5 User Schema Documentation

This document outlines the schema design for Auth5's authentication system.

## User Model Overview

The User model is the core schema for authentication and user management. It's designed to be:

- Secure: Sensitive data is properly handled
- Flexible: Supports multiple authentication methods
- Compliant: Follows GDPR and other regulatory requirements
- Maintainable: Clear separation of concerns

## Core Components

### Basic User Information

```go
type User struct {
    ID          bson.ObjectID  // Unique identifier
    CreatedAt   time.Time      // Account creation timestamp
    UpdatedAt   time.Time      // Last update timestamp
    CreatedBy   *bson.ObjectID // Admin who created the account (if applicable)

    // Identity Information
    Username    string       // Unique username
    DisplayName string       // Public display name
    Email       string       // Primary email
    PhoneNumber string       // Phone number
    Password    string       // Password hash (never exposed in JSON)
    AvatarURL   string       // Profile picture URL
    Locale      string       // User's locale (e.g., en-US)
    TimeZone    string       // User's timezone
    AccountType string       // Account type/tier (e.g., "free", "premium", "business", "enterprise", etc.)
}
```

### Account Types

Account types are stored as strings to provide maximum flexibility. Common values include:

- `"free"` - Basic free account
- `"premium"` - Premium individual account
- `"business"` - Business account
- `"enterprise"` - Enterprise account

However, the system supports any custom account type string, allowing for:

- Custom tiers (e.g., `"pro"`, `"ultimate"`, `"starter"`)
- Industry-specific types (e.g., `"healthcare"`, `"education"`)
- Regional variations (e.g., `"enterprise-us"`, `"enterprise-eu"`)
- Time-limited types (e.g., `"trial"`, `"beta"`)

This flexibility enables:

1. Easy addition of new account types without code changes
2. Support for different pricing tiers
3. Custom account types for specific use cases
4. Future-proofing for unknown requirements

## Authentication Components

### AuthInfo

Groups all authentication-related fields:

```go
type AuthInfo struct {
    // Password Management
    PasswordResetToken  string     // Token for password reset
    PasswordResetSentAt *time.Time // When reset was requested
    LastPasswordChange  time.Time  // Last password update
    LastPasswordReset   *time.Time // Last password reset

    // Verification Status
    EmailVerified           bool       // Email verification status
    PhoneVerified           bool       // Phone verification status
    EmailVerificationToken  string     // Email verification token
    EmailVerificationSentAt *time.Time // When verification was sent
    PhoneVerificationToken  string     // Phone verification token
    PhoneVerificationSentAt *time.Time // When verification was sent

    // Two-Factor Authentication
    TOTPSecret      string     // TOTP secret key
    Is2FAEnabled    bool       // 2FA status
    OTPBackupCodes  []string   // Backup codes for 2FA
    Last2FAVerified *time.Time // Last 2FA verification

    // OAuth Providers
    OAuthProviders map[string]OAuthProvider // Connected OAuth accounts
}
```

### OAuth Provider

Represents a connected third-party account:

```go
type OAuthProvider struct {
    ProviderID       string    // ID from the OAuth provider
    ProviderEmail    string    // Email from the provider
    ProviderUsername string    // Username from provider
    ProviderAvatar   string    // Avatar from provider
    ConnectedAt      time.Time // When account was connected
    LastUsedAt       time.Time // When account was last used
}
```

## Account Status Management

### User Status

```go
const (
    USER_STATUS_ACTIVE     = "active"     // Normal active account
    USER_STATUS_INACTIVE   = "inactive"   // Temporarily inactive
    USER_STATUS_PENDING    = "pending"    // Pending verification
    USER_STATUS_SUSPENDED  = "suspended"  // Suspended account
    USER_STATUS_DELETED    = "deleted"    // Account deleted
    USER_STATUS_ANONYMIZED = "anonymized" // Account data anonymized
)
```

### Suspension Management

```go
type SuspensionInfo struct {
    IsSuspended     bool       // Suspension status
    SuspendedAt     *time.Time // When suspended
    SuspendedUntil  *time.Time // Suspension end date
    SuspendedReason string     // Reason for suspension
    SuspendedBy     string     // Who suspended the account
}
```

## Activity Tracking

The system tracks various user activities:

```go
// Activity Fields in User struct
LastLoginAt     *time.Time // Last successful login
LastActivityAt  *time.Time // Last user activity
LastEmailChange *time.Time // Last email change
LastPhoneChange *time.Time // Last phone change
```

## GDPR Compliance

See [GDPR Deletion Documentation](gdpr_deletion.md) for detailed information about:

- Account deletion process
- Data anonymization
- Retention periods
- Legal compliance

## Security Considerations

1. **Sensitive Data**

   - Passwords are never stored in plain text
   - Tokens are never exposed in JSON responses
   - OAuth tokens are managed separately
   - 2FA secrets are properly secured

2. **Data Access**

   - Clear separation of public and private data
   - Proper JSON field tags for data exposure
   - Sensitive fields marked with `json:"-"`

3. **Audit Trail**
   - All important actions are timestamped
   - Changes are tracked with proper metadata
   - Suspension and deletion actions are logged

## Best Practices

1. **Field Usage**

   - Use pointers (`*time.Time`) for optional timestamps
   - Use `omitempty` for optional fields
   - Keep sensitive data out of JSON responses

2. **Validation**

   - Email addresses should be validated
   - Phone numbers should follow E.164 format
   - Usernames should be unique
   - Passwords should meet security requirements

3. **Updates**
   - Always update `UpdatedAt` on changes
   - Track who made changes when applicable
   - Maintain proper audit trails

## Future Considerations

1. **Potential Additions**

   - Device tracking
   - Login history
   - IP address tracking
   - Security preferences

2. **Scalability**

   - Schema supports horizontal scaling
   - Indexes should be added for frequent queries
   - Consider sharding for large deployments

3. **Integration**
   - Schema designed for easy integration
   - Supports multiple authentication methods
   - Flexible for future additions

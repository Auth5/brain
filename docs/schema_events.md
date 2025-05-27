# Events Schema Documentation

This document describes the event logging schema used for tracking various system events.

## Collections

All event collections have TTL (Time To Live) indexes on the `created_at` field:

| Collection         | TTL Period | Description                      |
| ------------------ | ---------- | -------------------------------- |
| `login_history`    | 90 days    | User login attempts and sessions |
| `email_history`    | 30 days    | Email sending attempts           |
| `account_history`  | 1 year     | Account changes and updates      |
| `security_history` | 1 year     | Security-related events          |
| `admin_history`    | 2 years    | Administrative actions           |

## Common Fields

All event records include:

- `_id`: MongoDB ObjectID
- `created_at`: Timestamp of the event (used for TTL)

## Event Types

### Login Events

```go
type LoginEventType string

const (
    LOGIN_EVENT_SUCCESS = "success" // Successful login
    LOGIN_EVENT_FAILED  = "failed"  // Failed login attempt
    LOGIN_EVENT_LOGOUT  = "logout"  // User logout
    LOGIN_EVENT_EXPIRED = "expired" // Session expired
    LOGIN_EVENT_REVOKED = "revoked" // Session revoked
)
```

### Account Events

```go
type AccountEventType string

const (
    ACCOUNT_EVENT_EMAIL_CHANGE    = "email_change"    // Email address change (e.g. old: "user@old.com" -> new: "user@new.com")
    ACCOUNT_EVENT_PHONE_CHANGE    = "phone_change"    // Phone number change (e.g. old: "+1234567890" -> new: "+0987654321")
    ACCOUNT_EVENT_PASSWORD_CHANGE = "password_change" // Password change (e.g. field: "password", old: "[REDACTED]", new: "[REDACTED]")
    ACCOUNT_EVENT_PROFILE_UPDATE  = "profile_update"  // Profile information update (e.g. field: "name", old: "John" -> new: "John Doe")
    ACCOUNT_EVENT_ACCOUNT_TYPE    = "account_type"    // Account type change (e.g. old: "free" -> new: "premium")
)
```

### Security Events

```go
type SecurityEventType string

const (
    SECURITY_EVENT_2FA_ENABLE     = "2fa_enable"     // 2FA enabled
    SECURITY_EVENT_2FA_DISABLE    = "2fa_disable"    // 2FA disabled
    SECURITY_EVENT_OAUTH_LINK     = "oauth_link"     // OAuth account linked
    SECURITY_EVENT_OAUTH_UNLINK   = "oauth_unlink"   // OAuth account unlinked
    SECURITY_EVENT_PASSWORD_RESET = "password_reset" // Password reset
)
```

### Admin Events

```go
type AdminEventType string

const (
    ADMIN_EVENT_SUSPEND     = "suspend"     // Account suspension
    ADMIN_EVENT_UNSUSPEND   = "unsuspend"   // Account unsuspension
    ADMIN_EVENT_DELETE      = "delete"      // Account deletion
    ADMIN_EVENT_ANONYMIZE   = "anonymize"   // Account anonymization
    ADMIN_EVENT_ROLE_CHANGE = "role_change" // Role/permission change
)
```

## Event Models

### LoginHistory

Tracks user login activity.

| Field        | Type           | Required | Description                    |
| ------------ | -------------- | -------- | ------------------------------ |
| `user_id`    | ObjectID       | Yes      | Reference to User model        |
| `event_type` | LoginEventType | Yes      | Type of login event            |
| `ip_address` | string         | Yes      | IP address of the user         |
| `country`    | string         | Yes      | Country code (e.g. "US", "GB") |
| `user_agent` | string         | Yes      | User agent string              |
| `success`    | bool           | Yes      | Whether login was successful   |
| `error`      | string         | No       | Error message if failed        |

### EmailHistory

Tracks email sending attempts.

| Field        | Type           | Required | Description                               |
| ------------ | -------------- | -------- | ----------------------------------------- |
| `user_id`    | ObjectID       | No       | Reference to User model (if applicable)   |
| `email_type` | EmailEventType | Yes      | Type of email (verification, reset, etc.) |
| `to`         | string         | Yes      | Recipient email                           |
| `subject`    | string         | Yes      | Email subject                             |
| `success`    | bool           | Yes      | Whether email was sent successfully       |
| `error`      | string         | No       | Error message if failed                   |

### AccountHistory

Tracks account changes.

| Field        | Type             | Required | Description                             |
| ------------ | ---------------- | -------- | --------------------------------------- |
| `user_id`    | ObjectID         | Yes      | Reference to User model                 |
| `event_type` | AccountEventType | Yes      | Type of account event                   |
| `field`      | string           | Yes      | Field that was changed                  |
| `old_value`  | string           | No       | Previous value                          |
| `new_value`  | string           | No       | New value                               |
| `changed_by` | string           | Yes      | Who made the change (user_id or system) |
| `ip_address` | string           | No       | IP address of the change                |
| `country`    | string           | No       | Country code (e.g. "US", "GB")          |
| `user_agent` | string           | No       | User agent of the change                |

### SecurityHistory

Tracks security-related events.

| Field        | Type              | Required | Description                       |
| ------------ | ----------------- | -------- | --------------------------------- |
| `user_id`    | ObjectID          | Yes      | Reference to User model           |
| `event_type` | SecurityEventType | Yes      | Type of security event            |
| `provider`   | string            | No       | OAuth provider (if applicable)    |
| `ip_address` | string            | Yes      | IP address of the event           |
| `country`    | string            | Yes      | Country code (e.g. "US", "GB")    |
| `user_agent` | string            | Yes      | User agent string                 |
| `success`    | bool              | Yes      | Whether the action was successful |
| `error`      | string            | No       | Error message if failed           |

### AdminHistory

Tracks administrative actions.

| Field        | Type           | Required | Description                             |
| ------------ | -------------- | -------- | --------------------------------------- |
| `admin_id`   | ObjectID       | Yes      | Reference to admin User model           |
| `user_id`    | ObjectID       | Yes      | Reference to affected User model        |
| `event_type` | AdminEventType | Yes      | Type of admin event                     |
| `action`     | string         | Yes      | Action taken                            |
| `reason`     | string         | No       | Reason for the action                   |
| `details`    | string         | No       | Additional details                      |
| `ip_address` | string         | Yes      | IP address of the admin                 |
| `country`    | string         | Yes      | Country code (e.g. "US", "GB")          |
| `user_agent` | string         | Yes      | User agent string                       |
| `expires_at` | time.Time      | No       | When the action expires (if applicable) |

## Best Practices

1. **Data Retention**

   - Events are automatically deleted based on their TTL
   - Admin history is kept longer (2 years) for compliance
   - Login and email history have shorter retention (90/30 days)

2. **Security**

   - Sensitive data (like passwords) should be redacted in old/new values
   - IP addresses and country codes are required for security events
   - User agent strings help track suspicious activity

3. **Event Types**

   - Use the predefined event type constants
   - Each event type has a specific purpose and data requirements
   - Follow the examples in comments for proper usage

4. **Required Fields**
   - All events must have `created_at` for TTL
   - Security-related fields (IP, country) are required for security events
   - User identification is required for all events except some email events

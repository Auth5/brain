# GDPR-Compliant Account Deletion

This document outlines the process and implementation of GDPR-compliant account deletion in Auth5.

## Account States

### Deleted Status (`USER_STATUS_DELETED`)

- Initial state when a user requests account deletion
- Account is marked for deletion but data still exists
- User cannot log in
- Data is preserved for legal/compliance purposes
- Triggered when user clicks "Delete Account" button

### Anonymized Status (`USER_STATUS_ANONYMIZED`)

- Personal data is irreversibly scrambled
- Direct identifiers are removed or replaced with random values
- Data exists but cannot be linked to the original user
- Example: Email "john@example.com" becomes "user_123@deleted.com"

## Deletion Process

### 1. Account Deletion Request

When a user requests account deletion, the system:

- Sets account status to `USER_STATUS_DELETED`
- Records deletion timestamp in `DeletionInfo.DeletedAt`
- Sets deletion reason to `DELETION_REASON_USER_REQUEST`
- Records requester in `DeletionInfo.RequestedBy`
- Revokes all active sessions
- Prevents new login attempts

### 2. Anonymization Period

The system maintains a grace period before anonymization:

- If `RetentionUntil` is set (legal requirement):
  - Data is preserved until the specified date
  - Anonymization proceeds after retention period
- If no retention period:
  - System waits for grace period (default: 30 days)
  - Anonymization proceeds after grace period

### 3. Data Anonymization

During anonymization, the system:

- Anonymizes personal identifiers:
  - Email: "john@example.com" → "user_123@deleted.com"
  - Username: "johndoe" → "user_456"
  - DisplayName: "John Doe" → "Deleted User"
  - Phone: "+1234567890" → "DELETED"
  - Password: [hashed value] → [new random hash]
  - OAuth data: Removed or anonymized
- Updates account status to `USER_STATUS_ANONYMIZED`
- Records anonymization timestamp in `DeletionInfo.AnonymizedAt`
- Records anonymization actor in `DeletionInfo.AnonymizedBy`

## Timeline Example

```
Day 1: User requests deletion
├── Account marked as DELETED
├── User login disabled
└── Data preserved for 30-day grace period

Day 30: Grace period ends
├── System anonymizes the data
├── Status changes to ANONYMIZED
└── Data becomes unlinkable to original user

Day 365: Legal retention period ends (if applicable)
├── System can fully delete the record
└── Or maintain anonymized data for analytics
```

## Special Cases

### Legal Hold

When an account is under legal investigation:

- System sets `RetentionUntil` to required date
- Original data is preserved until that date
- Anonymization proceeds after legal hold is lifted

### Admin Deletion

When an admin initiates account deletion:

- Process similar to user-initiated deletion
- Reason set to `DELETION_REASON_ADMIN_ACTION`
- May follow different retention rules based on policy

### System-Initiated Deletion

For automated deletion cases:

- Triggered for inactive accounts
- Triggered for policy violations
- Process is automated but follows same steps
- Reason set to `DELETION_REASON_SYSTEM_ACTION`

## Implementation Notes

### Data Fields

The system tracks deletion through the `DeletionInfo` struct:

```go
type DeletionInfo struct {
    DeletedAt      time.Time       // When account was marked for deletion
    AnonymizedAt   *time.Time      // When data was anonymized
    Reason         DELETION_REASON // Why account was deleted
    RequestedBy    string          // Who requested deletion
    AnonymizedBy   string          // Who performed anonymization
    RetentionUntil *time.Time      // Legal retention period
}
```

### Deletion Reasons

```go
const (
    DELETION_REASON_USER_REQUEST    = "user_request"    // User requested deletion
    DELETION_REASON_ADMIN_ACTION    = "admin_action"    // Deleted by admin
    DELETION_REASON_SYSTEM_ACTION   = "system_action"   // System-initiated
    DELETION_REASON_LEGAL_REQUIREMENT = "legal_requirement" // Legal requirement
)
```

## Benefits

This implementation ensures:

- Users cannot access their data after deletion
- Legal requirements are met
- Data is properly anonymized
- Clear audit trail is maintained
- Grace period for account recovery
- Compliance with GDPR and other regulations

## Security Considerations

1. **Data Access**

   - Deleted accounts are immediately inaccessible to users
   - Admin access is logged and audited
   - Legal access requires proper authorization

2. **Data Retention**

   - Clear retention periods are enforced
   - Legal holds override standard deletion process
   - All retention decisions are logged

3. **Audit Trail**
   - All deletion-related actions are logged
   - System maintains clear record of who did what and when
   - Logs are preserved according to legal requirements

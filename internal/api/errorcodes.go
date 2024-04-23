package api

type ErrorCode = string

const (
	// ErrorCodeUnknown should not be used directly, it only indicates a failure in the error handling system in such a way that an error code was not assigned properly.
	ErrorCodeUnknown ErrorCode = "unknown"

	// ErrorCodeUnexpectedFailure signals an unexpected failure such as a 500 Internal Server Error.
	ErrorCodeUnexpectedFailure ErrorCode = "unexpected_failure"

	ErrorCodeValidationFailed                  ErrorCode = "validation_failed"
	ErrorCodeBadJSON                           ErrorCode = "bad_json"
	ErrorCodeEmailExists                       ErrorCode = "email_exists"
	ErrorCodePhoneExists                       ErrorCode = "phone_exists"
	ErrorCodeBadJWT                            ErrorCode = "bad_jwt"
	ErrorCodeNotAdmin                          ErrorCode = "not_admin"
	ErrorCodeNoAuthorization                   ErrorCode = "no_authorization"
	ErrorCodeUserNotFound                      ErrorCode = "user_not_found"
	ErrorCodeSessionNotFound                   ErrorCode = "session_not_found"
	ErrorCodeFlowStateNotFound                 ErrorCode = "flow_state_not_found"
	ErrorCodeFlowStateExpired                  ErrorCode = "flow_state_expired"
	ErrorCodeSignupDisabled                    ErrorCode = "signup_disabled"
	ErrorCodeUserBanned                        ErrorCode = "user_banned"
	ErrorCodeProviderEmailNeedsVerification    ErrorCode = "provider_email_needs_verification"
	ErrorCodeInviteNotFound                    ErrorCode = "invite_not_found"
	ErrorCodeBadOAuthState                     ErrorCode = "bad_oauth_state"
	ErrorCodeBadOAuthCallback                  ErrorCode = "bad_oauth_callback"
	ErrorCodeOAuthProviderNotSupported         ErrorCode = "oauth_provider_not_supported"
	ErrorCodeUnexpectedAudience                ErrorCode = "unexpected_audience"
	ErrorCodeSingleIdentityNotDeletable        ErrorCode = "single_identity_not_deletable"
	ErrorCodeEmailConflictIdentityNotDeletable ErrorCode = "email_conflict_identity_not_deletable"
	ErrorCodeIdentityAlreadyExists             ErrorCode = "identity_already_exists"
	ErrorCodeEmailProviderDisabled             ErrorCode = "email_provider_disabled"
	ErrorCodePhoneProviderDisabled             ErrorCode = "phone_provider_disabled"
	ErrorCodeTooManyEnrolledMFAFactors         ErrorCode = "too_many_enrolled_mfa_factors"
	ErrorCodeMFAFactorNameConflict             ErrorCode = "mfa_factor_name_conflict"
	ErrorCodeMFAFactorNotFound                 ErrorCode = "mfa_factor_not_found"
	ErrorCodeMFAIPAddressMismatch              ErrorCode = "mfa_ip_address_mismatch"
	ErrorCodeMFAChallengeExpired               ErrorCode = "mfa_challenge_expired"
	ErrorCodeMFAVerificationFailed             ErrorCode = "mfa_verification_failed"
	ErrorCodeMFAVerificationRejected           ErrorCode = "mfa_verification_rejected"
	ErrorCodeInsufficientAAL                   ErrorCode = "insufficient_aal"
	ErrorCodeCaptchaFailed                     ErrorCode = "captcha_failed"
	ErrorCodeSAMLProviderDisabled              ErrorCode = "saml_provider_disabled"
	ErrorCodeManualLinkingDisabled             ErrorCode = "manual_linking_disabled"
	ErrorCodeSMSSendFailed                     ErrorCode = "sms_send_failed"
	ErrorCodeEmailNotConfirmed                 ErrorCode = "email_not_confirmed"
	ErrorCodePhoneNotConfirmed                 ErrorCode = "phone_not_confirmed"
	ErrorCodeSAMLRelayStateNotFound            ErrorCode = "saml_relay_state_not_found"
	ErrorCodeSAMLRelayStateExpired             ErrorCode = "saml_relay_state_expired"
	ErrorCodeSAMLIdPNotFound                   ErrorCode = "saml_idp_not_found"
	ErrorCodeSAMLAssertionNoUserID             ErrorCode = "saml_assertion_no_user_id"
	ErrorCodeSAMLAssertionNoEmail              ErrorCode = "saml_assertion_no_email"
	ErrorCodeUserAlreadyExists                 ErrorCode = "user_already_exists"
	ErrorCodeSSOProviderNotFound               ErrorCode = "sso_provider_not_found"
	ErrorCodeSAMLMetadataFetchFailed           ErrorCode = "saml_metadata_fetch_failed"
	ErrorCodeSAMLIdPAlreadyExists              ErrorCode = "saml_idp_already_exists"
	ErrorCodeSSODomainAlreadyExists            ErrorCode = "sso_domain_already_exists"
	ErrorCodeSAMLEntityIDMismatch              ErrorCode = "saml_entity_id_mismatch"
	ErrorCodeConflict                          ErrorCode = "conflict"
	ErrorCodeProviderDisabled                  ErrorCode = "provider_disabled"
	ErrorCodeUserSSOManaged                    ErrorCode = "user_sso_managed"
	ErrorCodeReauthenticationNeeded            ErrorCode = "reauthentication_needed"
	ErrorCodeSamePassword                      ErrorCode = "same_password"
	ErrorCodeReauthenticationNotValid          ErrorCode = "reauthentication_not_valid"
	ErrorCodeOTPExpired                        ErrorCode = "otp_expired"
	ErrorCodeOTPDisabled                       ErrorCode = "otp_disabled"
	ErrorCodeIdentityNotFound                  ErrorCode = "identity_not_found"
	ErrorCodeWeakPassword                      ErrorCode = "weak_password"
	ErrorCodeOverRequestRateLimit              ErrorCode = "over_request_rate_limit"
	ErrorCodeOverEmailSendRateLimit            ErrorCode = "over_email_send_rate_limit"
	ErrorCodeOverSMSSendRateLimit              ErrorCode = "over_sms_send_rate_limit"
	ErrorBadCodeVerifier                       ErrorCode = "bad_code_verifier"
	ErrorCodeAnonymousProviderDisabled         ErrorCode = "anonymous_provider_disabled"
	ErrorCodeHookTimeout                       ErrorCode = "hook_timeout"
	ErrorCodeHookTimeoutAfterRetry             ErrorCode = "hook_timeout_after_retry"
	ErrorCodeHookPayloadOverSizeLimit          ErrorCode = "hook_payload_over_size_limit"
	ErrorCodeHookPayloadUnknownSize            ErrorCode = "hook_payload_unknown_size"
)

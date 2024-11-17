package utils 

import (
	"regexp"
)

var (
  uppercasePattern   = regexp.MustCompile(`[A-Z]`)
  lowercasePattern   = regexp.MustCompile(`[a-z]`)
  digitPattern       = regexp.MustCompile(`[0-9]`)
  specialCharPattern = regexp.MustCompile(`[^\w\s]`)
  spacePattern       = regexp.MustCompile(`\s`)
)

type FieldValidationRequest struct {
  FieldName string `json:"fieldName"`
  Value     string `json:"value"`
}

func ValidateFields(fields []FieldValidationRequest) map[string]string {
  errors := make(map[string]string)

  for _, field := range fields {
    switch field.FieldName {
    case "displayName":
      errors[field.FieldName] = validateDisplayName(field.Value)
    case "email":
      errors[field.FieldName] = validateEmail(field.Value)
    case "username":
      errors[field.FieldName] = validateUsername(field.Value)
    case "password":
      errors[field.FieldName] = validatePassword(field.Value)
    default:
      errors[field.FieldName] = "Unknown field"
    }
  }

  return errors
}

func validateDisplayName(value string) string {
	if value == "" {
		return "Display name is required"
	}
	if len(value) > 50 {
		return "Display name cannot exceed 50 characters"
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9 _-]+$`)
	if !re.MatchString(value) {
		return "Display name can only contain letters, numbers, spaces, underscores, and hypnens"
	}
	return ""
}

func validateEmail(value string) string {
	if value == "" {
		return "Email is required"
	}
	// Basic email regex pattern
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(value) {
		return "Invalid email format"
	}
	return ""
}

func validateUsername(value string) string {
	if value == "" {
		return "Username is required"
	}
	if len(value) > 30 {
		return "Username cannot exceed 30 characters"
	}
	// Allow only alphanumeric characters (no spaces or special characters)
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(value) {
		return "Username can only contain letters and numbers"
	}
	return ""
}

func validatePassword(value string) string {
  if value == "" {
    return "Password is required"
  }
  if len(value) < 8 {
    return "Password must be at least 8 characters."
  }
  if !uppercasePattern.MatchString(value) {
    return "Password must contain at least one uppercase letter"
  }
  if !lowercasePattern.MatchString(value) {
    return "Password must contain at least one lowercase letter"
  }
  if !digitPattern.MatchString(value) {
    return "Password must contain at least one digit"
  }
  if !specialCharPattern.MatchString(value) {
    return "Password must contain at least one special character"
    }
  if spacePattern.MatchString(value) {
    return "Password cannot contain spaces"
  }
  return ""
}
package beacon

import (
	"fmt"
	"strings"
	"testing"
)

func TestDetailTypeValidate(t *testing.T) {
	cases := []struct {
		name           string
		detailType     DetailType
		expectedErr    bool
		expectedErrMsg string
	}{
		{
			"knowngood",
			DetailType("test"),
			false,
			"",
		},
		{
			"empty",
			DetailType(""),
			true,
			"Detail type cannot be empty",
		},
		{
			"at maximum",
			DetailType(strings.Repeat("a", DETAIL_TYPE_MAX_LENGTH)),
			false,
			"",
		},
		{
			"too long",
			DetailType(strings.Repeat("a", DETAIL_TYPE_MAX_LENGTH+1)),
			true,
			fmt.Sprintf("Detail type length of %d bytes exceeds %d bytes", DETAIL_TYPE_MAX_LENGTH+1, DETAIL_TYPE_MAX_LENGTH),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.detailType.Validate()

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
					t.Errorf("Expected error message %q but got %q", tt.expectedErrMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestInstanceIDValidate(t *testing.T) {
	cases := []struct {
		name           string
		instanceID     InstanceID
		expectedErr    bool
		expectedErrMsg string
	}{
		{
			"knowngood",
			InstanceID("test"),
			false,
			"",
		},
		{
			"empty",
			InstanceID(""),
			true,
			"Instance ID cannot be empty",
		},
		{
			"at maximum",
			InstanceID(strings.Repeat("a", RESOURCE_ARN_MAX_LENGTH)),
			false,
			"",
		},
		{
			"too long",
			InstanceID(strings.Repeat("a", RESOURCE_ARN_MAX_LENGTH+1)),
			true,
			fmt.Sprintf("Instance ID length of %d bytes exceeds %d bytes", RESOURCE_ARN_MAX_LENGTH+1, RESOURCE_ARN_MAX_LENGTH),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.instanceID.Validate()

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
					t.Errorf("Expected error message %q but got %q", tt.expectedErrMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

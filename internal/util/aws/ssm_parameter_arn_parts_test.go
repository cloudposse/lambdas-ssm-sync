package awsutil

import (
	"reflect"
	"testing"
)

func TestParseArn(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    SSMParameterARNParts
		wantErr bool
	}{
		{
			name:  "Valid ARN",
			input: "arn:aws:ssm:us-east-1:333333333333:parameter/terraform/shared-state/111111111111/us-east-2/foo",
			want: SSMParameterARNParts{
				SourceRegion:  "us-east-1",
				SourceAccount: "333333333333",
				PathAccount:   "111111111111",
				PathRegion:    "us-east-2",
			},
			wantErr: false,
		},
		{
			name:  "Valid ARN",
			input: "arn:aws:ssm:us-east-1:000000000000:parameter/terraform/shared-state/999999999999/us-east-2/foo/bar/baz",
			want: SSMParameterARNParts{
				SourceRegion:  "us-east-1",
				SourceAccount: "000000000000",
				PathAccount:   "999999999999",
				PathRegion:    "us-east-2",
			},
			wantErr: false,
		},
		{
			name:    "Invalid ARN - Wrong Format",
			input:   "invalid:format",
			want:    SSMParameterARNParts{},
			wantErr: true,
		},
		{
			name:    "Invalid ARN - Wrong Service",
			input:   "arn:aws:ec2:us-east-1:999999999999:parameter/terraform/shared-state/111111111111/us-east-2/foo",
			want:    SSMParameterARNParts{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSSMParameterARNParts(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArn() = %v, want %v", got, tt.want)
			}
		})
	}
}

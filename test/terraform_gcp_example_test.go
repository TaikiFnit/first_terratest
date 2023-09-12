package test
import (
  "fmt"
  "strings"
  "testing"
  "time"

  "github.com/gruntwork-io/terratest/modules/gcp"
  "github.com/gruntwork-io/terratest/modules/random"
  "github.com/gruntwork-io/terratest/modules/retry"
  "github.com/gruntwork-io/terratest/modules/ssh"
  "github.com/gruntwork-io/terratest/modules/terraform"
  test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
  "github.com/stretchr/testify/assert"
)

func TestTerraformGcpExample(t *testing.T) {
  t.Parallel()

  exampleDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/terraform-gcp-example")
  projectId := gcp.GetGoogleProjectIDFromEnvVar(t)
  zone := "us-east1-b"

  expectedBucketName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId())))
  expectedInstanceName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: exampleDir,
    Vars: map[string]interface{}{
      "gcp_project_id": projectId,
      "zone": zone,
      "instance_name": expectedInstanceName,
      "bucket_name": expectedBucketName,
    },
  })

  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t terraformOptions)

  bucketURL := terraform.Output(t, terraformOptions, "bucket_url")
  instanceName := terraform.Output(t, terraformOptions, "instance_name")

  expectedURL := fmt.Sprintf("gs://%s", expectedBucketName)
  assert.Equal(t, expectedURL, bucketURL)

  gcp.AssertStorageBucketExists(t, expectedBucketName)

  instance := gcp.FetchInstance(t, projectId, instanceName)
  instance.SetLabels(t, map[string]string{"testing": "testing-tag-value2"})

  maxRetries := 12
  timeBetweenRetries := 5 * time.Second
  expectedText := "testing-tag-value2"

  retry.DoWithRetry(t, fmt.Sprintf("Checking Instance %s for labels", instanceName), maxRetries, timeBetweenRetries, func() (string, error) {
    instance := gcp.FetchInstance(t, projectId, instanceName)
    instanceLabels := instance.GetLabels(t)

    testingTag, containsTestingTag := instanceLabels["testing"]
    actualText := strings.TrimSpace(testingTag)
    if !containsTestingTag {
      return "", fmt.Errorf("Expected the tag 'testing' to exist")
    }

    if actualText != expectedText {
      return "", fmt.Errorf("Expected GetLabelsForComputeInstanceE to return '%s' but got '%s'", expectedTest, actualText)
    }

    return "", nil
  })
}


func TestSshAccessToComputeInstance(t *testing.T) {
	t.Parallel()

	exampleDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/terraform-gcp-example")

	// Setup values for our Terraform apply
	projectID := gcp.GetGoogleProjectIDFromEnvVar(t)
	randomValidGcpName := gcp.RandomValidGcpName()
	zone := gcp.GetRandomZone(t, projectID, ZonesThatSupportF1Micro, nil, nil)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: exampleDir,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"gcp_project_id": projectID,
			"instance_name":  randomValidGcpName,
			"bucket_name":    randomValidGcpName,
			"zone":           zone,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	publicIp := terraform.Output(t, terraformOptions, "public_ip")

	// Attempt to SSH and execute the command
	instance := gcp.FetchInstance(t, projectID, randomValidGcpName)

	sampleText := "Hello World"
	sshUsername := "terratest"

	keyPair := ssh.GenerateRSAKeyPair(t, 2048)
	instance.AddSshKey(t, sshUsername, keyPair.PublicKey)

	host := ssh.Host{
		Hostname:    publicIp,
		SshKeyPair:  keyPair,
		SshUserName: sshUsername,
	}

	maxRetries := 20
	sleepBetweenRetries := 3 * time.Second

	retry.DoWithRetry(t, "Attempting to SSH", maxRetries, sleepBetweenRetries, func() (string, error) {
		output, err := ssh.CheckSshCommandE(t, host, fmt.Sprintf("echo '%s'", sampleText))
		if err != nil {
			return "", err
		}

		if strings.TrimSpace(sampleText) != strings.TrimSpace(output) {
			return "", fmt.Errorf("Expected: %s. Got: %s\n", sampleText, output)
		}

		return "", nil
	})
}

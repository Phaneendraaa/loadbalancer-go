package serverHandler

import (
	"fmt"

	"os/exec"
	"strings"
)

func createInstance() (string, error) {

	cmd := exec.Command("aws", "ec2", "run-instances",
		"--image-id", "ami-0ba6571c361794862",
		"--count", "1",
		"--instance-type", "t2.micro",
		"--key-name", "health-keypair",
		"--security-group-ids", "sg-03f188c58cdb64667",
		"--subnet-id", "subnet-08888897fdfd5a5d2",
		"--associate-public-ip-address",                                                                   
		"--tag-specifications", "ResourceType=instance,Tags=[{Key=Name,Value=health-web-dev-from-awscli}]",
		"--region", "us-east-1", 
		"--query", "Instances[0].InstanceId", 
		"--output", "text", 
	)

	instanceIDBytes, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("unable to create instance, %v", err)
	}
	instanceID := strings.TrimSpace(string(instanceIDBytes))
	fmt.Printf("Successfully launched instance with ID: %s\n", instanceID)

	cmd = exec.Command("aws", "ec2", "describe-instances",
		"--instance-ids", instanceID,
		"--query", "Reservations[0].Instances[0].PublicIpAddress",
		"--output", "text",
	)
	publicIPBytes, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("unable to describe instance, %v", err)
	}
	publicIP := strings.TrimSpace(string(publicIPBytes))

	return publicIP, nil
}

func deleteInstance(publicIP string) error {
	describeCmd := exec.Command("aws", "ec2", "describe-instances",
		"--filters", "Name=ip-address,Values="+publicIP,
		"--query", "Reservations[0].Instances[0].InstanceId",
		"--output", "text",
	)

	instanceIDBytes, err := describeCmd.Output()
	if err != nil {
		return fmt.Errorf("unable to describe instance with IP %s, %v", publicIP, err)
	}

	instanceID := strings.TrimSpace(string(instanceIDBytes))
	if instanceID == "" {
		return fmt.Errorf("no instance found with public IP %s", publicIP)
	}

	fmt.Printf("Found instance ID: %s for IP: %s\n", instanceID, publicIP)

	terminateCmd := exec.Command("aws", "ec2", "terminate-instances",
		"--instance-ids", instanceID,
	)

	terminateResult, err := terminateCmd.Output()
	if err != nil {
		return fmt.Errorf("unable to terminate instance with ID %s, %v", instanceID, err)
	}

	fmt.Printf("Termination result: %s\n", string(terminateResult))

	return nil
}

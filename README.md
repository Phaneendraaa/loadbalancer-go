git add .
git commit -m "commit"
git push -u origin main

# creating instance using aws-cli
aws ec2 run-instances \
    --image-id ami-0ba6571c361794862 \
    --count 1 \
    --instance-type t2.micro \
    --key-name health-keypair \
    --security-group-ids sg-03f188c58cdb64667 \
    --subnet-id subnet-08888897fdfd5a5d2 \
    --associate-public-ip-address \
    --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=health-web-dev-from-awscli}]' \
    --region us-east-1


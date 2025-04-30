# Interview Prep Fortinet 

real end game

# Study Guide for Fortinet Staff Software Development QA Engineer Position

## Technical Skills to Study

• Python programming fundamentals and advanced concepts
  - Object-oriented programming principles in Python
  - Python test frameworks and libraries
  - Error handling and debugging techniques

• Pytest framework for test automation
  - Test fixtures and dependency injection
  - Parameterized testing
  - Mocking and patching
  - Test reporting features
  - Plugins and extensions
    - Py Test can run tests in parallel
    - if not explicitly mentioned has own way of detecting tests
    - can run subset of tests doesnt have to be all 
    - running tests without mentioning it will run all \*\_test.py and all test_\*\.py 
    - explicitely mention the file to run that 
    - functions HAVE to start with test
    - can run tests that have a string in the names by doing this `pytest -k <substring>`
    - fixtures run before the test, usually before the thing is being ran to feed data
    - `pytest-xdist` for multithreadign, `pytest -n <number of workers>`

• Cloud platforms (AWS, Azure, GCP) and their Python client libraries
  - AWS: EC2, S3, Lambda, IAM, CloudFormation basics
    - EC2: Basic VM youve used this before, scale as you go
    - Lambda, serverless compute, so it runs it based off of events, when your shit is used then it spawns it and spawns more of them as needed
    - S3 easy just storage
    - IAM identity access management, securely control AWS services
    - Cloudformation, infrastrcuture as code, configure it with code so you can have a copy of its deployment basically
  - Azure: Virtual Machines, Storage, Functions, Active Directory
    - Virtual Mahcines, EC2 basically 
    - Storage, basically s3, blob storage, queue storage for reliable messaging
    - Functions, serverless compute 
    - Active Directory, IAM basically allows you to manage its stuff
  - GCP: Compute Engine, Cloud Storage, Cloud Functions
    - Compute Engine, EC2 and VM from azure basically 
    - Cloud Storage, similar to s3 and storage from azure 
    - Cloud Functions basically serverless compute like lambda and functions
  - Authentication and security best practices
  - Resource provisioning and management

• Terraform for infrastructure as code
  - HCL syntax and structure
  - Resource and provider configuration
  - State management
  - Modules and reusable components
  - Provisioning cloud resources

• Docker and containerization concepts
  - Dockerfile creation and best practices
  - Container lifecycle management
  - Multi-container applications
  - Docker networking
  - Security considerations

• Kubernetes architecture and operations
  - Pod management and deployment strategies
  - Services and networking
  - ConfigMaps and Secrets
  - StatefulSets and persistent storage
  - RBAC and security controls

• Virtualization (VM) technologies
  - Hypervisor types and differences
  - VM resource management
  - VM networking concepts
  - Snapshot and backup strategies

• Database concepts (including Snowflake)
  - SQL fundamentals
  - Database design principles
  - Snowflake architecture and features
  - Data warehousing concepts
  - Query optimization

• TCP/IP protocols and networking fundamentals
  - OSI model and protocol stack
  - Common protocols (HTTP, HTTPS, DNS, SSH)
  - Network security concepts
  - Subnetting and routing basics
  - Troubleshooting network issues

• Selenium for web automation
  - Web element selection techniques
  - Waiting and synchronization strategies
  - Cross-browser testing
  - Page object model design pattern
  - Test reporting and screenshot capture

## Cloud Security Concepts

• Cloud security architecture principles
  - Shared responsibility model
  - Defense in depth strategies
  - Identity and access management
  - Network security in cloud environments
  - Data protection and encryption

• Common cloud security threats and mitigations
  - Misconfiguration risks
  - Account hijacking prevention
  - DDoS protection
  - Data breach prevention
  - Insider threat controls

• Security testing for cloud applications
  - Vulnerability assessment techniques
  - Penetration testing methodologies
  - Security scanning tools
  - API security testing
  - Compliance validation

## Quality Assurance Methodologies

• Test planning and documentation
  - Test strategy development
  - Risk-based testing approaches
  - Test plan components and structure
  - Traceability matrices
  - Documentation best practices

• User story creation and analysis
  - Acceptance criteria definition
  - Identifying testable requirements
  - Breaking down complex features
  - Edge case identification
  - User-centered testing approaches

• Test case development techniques
  - Positive and negative test scenarios
  - Boundary value analysis
  - Equivalence partitioning
  - Decision table testing
  - State transition testing

• Automation framework design
  - Page object patterns
  - Data-driven testing
  - Keyword-driven frameworks
  - Hybrid framework approaches
  - CI/CD integration patterns

## Development Lifecycle

• Agile/Scrum processes
  - Sprint planning and execution
  - User story estimation
  - Daily stand-ups and retrospectives
  - Definition of done criteria
  - Incremental delivery techniques

• CI/CD pipelines
  - Pipeline configuration and management
  - Automated testing integration
  - Deployment strategies
  - Infrastructure provisioning
  - Quality gates and approval processes

• DevOps practices in cloud environments
  - Infrastructure as code implementation
  - Monitoring and observability
  - Incident response automation
  - Configuration management
  - Compliance as code


## AWS Notes

### Compute
- EC2: K
  - elastic compute cloud
  - VM with configurability, auto scaling 
- Lambda K
  - Serverless, runs functions based off of events. auto scaling 
- ECS
  - elastic container service, fully managed so u dont do anything auto scaling
- EKS K
  - elastic kubernates service, managed kubernates service to run containerized applications
- FarGate
  - Similar to Lambda but thing to remember is Lambda has max 15 min run time 
- Batch
  - fully managed batch processing for efficiently running batch computing workloads has auto scalign 
- Light Sail 
  - simplified VPS, with easy deployment, simplified ec2 basically 
- AWS App Runner 
  - Fully Managed service for building deploying and scaling containerized web apps 
  
### Storage
- S3 K
  - object storage, scalability, avaialbaility, and security auto scales
- EBS K
  - elastic block store, high performance block storage volumes, no auto scaling, can be resized manually 
- EFS
  - Fully managed file system for EC2 and serverless applications
- FSX
  - fully manaed fily system for windows, and others, auto scales throuput storage has to be manually
- S3 Glacier K
  - low cost archived storage for rarely accessed but require long term retention
- Storage Gateway 
  - hybrid storage, connects on premise storage to cloud, no auto scaling
- AWS back up
  - centralized back ups for aws and on premise workloads, no auto scalign for service
- Snow Family
  - to move petabytes of data in and out of aws

### Database
- RDS K
  - Relational database service, supports multiple db languages, autoscaling for storage only 
- DynamoDB 
  - fully managed nosql database with a fully managed db service
- ElastiCache
  - in memory caching service for highly connected db
- Neptune
  - fully managed graph db service
- Redshift 
  - fully managed pt scaled warehouse for analytics
- Document DB 
  - mongodb compatable db service for storing querying and index json data autoscaling for storage only
- KeySpaces
  - fully managed apache cassandra compatbile service 
- QLDB
  - quantum ledger database, fully managed ledger database with immutable cryptograophy 
- TimeStream 
  - fully managed IOT and operational applications scales storage and cmpute
- MemoryDB for redis 
  - redis compatibe durable in memory db service no autoscaling but can add more nodes

### Networking and Content Delivery 
- VPC K
  - virtual private cloud isolated virtual network cloud to launch AWS resources
- Cloud front
  - global cdn to deliver data, videos with low latency 
- Route 53 K
  - scalable domain name web service
  - managed DNS 
  - load balancing 
- API Gateway K
  - managed service for creating, publishing, and securing APIs at any scale
- Direct Connect
  - dedicated connection to from on premise AWS
- Global Accelerator
  - service that improves the availability and performance using the AWS global network 
- App Mesh
  - service mesh that provides application level networking for microservices, no auto scaling 
- Transit Gateway
  - Network hub to connect VPCs
- VPC lattice
  - connects, secures, and monitros accounts across VPCs
- Elastic Load balancing K
  - distrubutes load accross multiple targets, scales with traffic

### Application Integration 
- SQS K
  - simple queue service, fully managed queuing service for decoupling services
- SNS 
  - simple notification service, pubsub messaging for microservices
- Event Bridge
  - serverless event buss that connects apps with real time data
- Step Functions
  - visuals for building distributed services
- App Flow
  - 
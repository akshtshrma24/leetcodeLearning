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
  - Azure: Virtual Machines, Storage, Functions, Active Directory
  - GCP: Compute Engine, Cloud Storage, Cloud Functions
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
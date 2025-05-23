# Interview Prep Lambda

## What does lambda do

Lambda labs is the number one gpu cloud for ML/AI teams for training fine tuning and inferencing AI Models 
engineers can build test and deploy their models 

products = on premis gpus, publically and privately hosted gpus for AI researchers and ai companies world wide

AI developer cloud 

delivers the fastest network for distributed training
3200 gb/s nvidia quantum 2 infini-band networking 

only cloud specifically designed for Generative AI training and inference
complete data privacy and isolation 

established for 1 reason and 1 reason to tackle the challenge of teaching machines how to see and learn 


* 2012 Lambda launches ML facial recognition API 
* 2013 - 2016 Lambda establishes an internal GPU cloud for clients to support their AI image editing 
* 2017 Lamdba launches quad deep learning gpu workstation and lambda blade gpu server, worlds first plug play gpu super comp under 20k 
* 2018 lambda launches GPU cloud and lambda stack a comprehensive AI software repo
* 2021 gets funding from googles AI centric fund, shifts focus to GPU cloud 
* 2022 gets funding for first serious large scale GPU deployment
* 2023 more funding 
* 2024 lambda launches 1 click clusters, 
* 2025 more investment 
* 2026 hopes to go IPO and continue 

## are there competitors 

AWS, Google Cloud


## who are the founders
Michael Balaban and Stephen Balaban

## how is lambda different and similar than competitors 

both offer GPUs in the cloud for AI workflows but lambdalabs has unparalleled performance 

## Evan questions 

* How often do you guys interact with Nvidia 
* do you receive software to test the gpus from Nvidia ,
* how do you guys interpret test results from the GPUs
* how do you determine when a GPU is unhleathy, when GPUs are unhealthy waht is the rpeaier process do you pull it out does it stay in the rack 
* How do you use NVLINK, do you guys rent out space by the rack
* how do repairs work? 
* Do you guys have tools to track repairs? 
* do you guys keep the history of the server say im monitoring a server and taken out of server, is there a tool to track of the server was it bad before has it needed repair before 
* what languages do you guys use besides Go and python, what tools do you guys use anectodadley 


* Nvidia SMI, tells you what gpus plugged in and what not

## Evans questions elaborated

* How often do you guys interact with Nvidia and their team 
  * given their strong relationship with Nvidia it is important to know how often they interact with their team 
* Do you guys receive software to test the gpus from Nvidia, 
  * Evan gets software to test Nvidia GPUs these tests consist of FIND OUT HERE <_____>
* How do you guys interpret the test results from the GPUs
  * will make more sense when find out what tests consist of 
* Nvidia SMI - nvidia system managment interface
  * command line utility based on Nvidia Manamgnet library (NVML)
  * designed to help monitor and manage nvidia gpus
  * allows you to query the GPU device state 
  * allows the change of the GPU device state
  * you can get cool metrics of the GPU like fan speed, and tx and rx througput
* how do you guys monitor the fleet right now what tools do you guys use 


## Bare Metal Configuration 

* DB Control 
  * I had 3 pis each running docker containers that would take /dev/file that would correspond to the usb port 
  * send bits through pymodbus to the machines got metrics through that 
  * I had 1 ansible system that would manage all of these 

## 2+ years of configuration 

* I have used Ansible extensively in my past, through self projects and professionaly
* Most notable at PayPal, where I was given the oppurtunity to work along side another team to setup a local setup of mock projects we had in PayPal. 
  * It was 6 servers and the reason that I was given this oppurtunity is because of a self project I did where I had 6 raspberry pis i deployed wiht ansible 
  * and basicallly what I did there was provision the 6 servers 
  * wrote various playbooks the main ones were creating ssh keys updating the apt 
  * playbooks to deploy python scripts across all of them for metrics, had to implement top features in python such as cpu usage memory and swap memory usage 
  * deployed docker containers across the ones that we were testing 

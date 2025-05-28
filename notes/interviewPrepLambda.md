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


## 2+ years of golang and python 

## Questions for the Culture fit 

* STAR Specific Example

* Tell me about a time 
  * when you had a team member that wasnt pulling their weight 
    * who the team member was what the problem was the result of it
      * natalya and i were given this project, i was responsible for this piece, 1 week into the project natalya does nothing, schedule the time with natalya to teach her, 
  * when you had a disagreement with the team, what were the steps you took 
  * when you had to make a difficult decision under pressure.
  * when you handled the problem outside the scope of work, try to find out information 
  * when you had to pivot on a project that was mostly complete 
    * YC 
    * have backup 
  * when you convinced the team to agree with you but outcome was negative 
    * take a look at the videos
    * low ego
  * when you pushed back against the project or aspect of the project you disagreed with 
  * when you had a project worked on that would have gone better without your involvment
* what are the 

## Dad Linux study what to 

```bash
fsck 
/etc/fstab
systemctl
journalctl 
dmesg 
oom 
ps 
htop 
top
vmstat
iostat
/proc
ulimit # how many file handdles what is a file handle how many files user can write to read to per user and system wide
# file descriptor
inode # file 
du # difference between these 2
df 
# df shows file system usage, when delete a file doesnt release space, 
# if file is open and delete it doesnt release the space
lsof # list of open files 
lsusb # list of us 
lspci

sed 
awk 

mount -v # check what file system you have 
# mount -v # check what file system you have alternative V
cat /etc/fstab

# making file system 
#   raw device
#   mkfs

# start at boot 
# create service file, enable it, start it with systemctl 


# Memory amangement ============
# paging in linux memory allocation, memory happens with linux in paging
# swapping, 
# ========

# se linux
# Secure Linux
# 


# file descriptor

# static ip with netplan 
# /etc/netplan/ yml

# nfs
# how to export a file system 

# stty sane

# kill
#   sighup 
#   different -x 

# stickybit 
#

# how can users change the password but not edit the file 
# file attributes

```


## Command study 

- `fsck` : file system check, checks for inconsistencies in the file system and cleans them, disk has unreadable sectors, inodes hard link counts are different
- `/etc/fstab` : mounts file system at boot
- `systemctl` : used to control the systemd services, start, stop, enable/disable
- `journalctl` : used to view logs, stores them in binary, allows you to search, view, and filter them `journalctl -u service` allows you to view the logs from that service used when debugging services
- `dmesg` : different freom journalctl because it can only view kernel logs, non persistent cleared on reboot, used to debug hardware or kernel
- 
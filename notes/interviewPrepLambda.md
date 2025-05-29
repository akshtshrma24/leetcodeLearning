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
- `ps` : used to see processes 
- `htop top` : used to see top processes htop more human readable scroll mouse 
- `vmstat` : virtual memory statistics, mem usage, cpu activity swap usage, io performance, system processes, gives real time snapshot of how your cpu and memory is performing, `vmstat 2` refreshes every 2 seconds
- `iostat` : similar to vmstat because it shows you hardware statistics how is the hardware performign
- `/proc` : used to virtual file system not stored on disk, mounted on boot, reflects current state of kernel and system, hardware details, kernel settings, process information, memory cpu device stats, 
- `ulimit` : how many files I can write to, if you set it too low you wont be able to boot
- `inode` : these are a set amount, you can only write this many of files in that directory
- `du` : disk usage, how much disk usage is being used by the directory mentioned 
- `df` : disk fyle system shows the disk free space on the available mounted file systems
- `lsof` : ls of open files
- `lsusb` : shows the information about the available usb ports
- `lspci` : shows information about the pci buses and the devices connected to them 
- `sed` : stream editor, use like this `sed 'expression' file`
- `awk` : reads a file line by line and splits them allows you to do operations on them


## Culture Fit questions with answers 
* Tell me about a time 
  * when you had a team member that wasnt pulling their weight 
    * S
      * During my time at DB Control, I was tasked with developing a remote monitoring stack, I chose integrating Prometheus and Grafana with low level pymodbus code to get metrics and build pipelines to the stack, I was on a short deadline and I needed to get going fast. We had about 15+ machines to test with, so we needed to move.
    * T
      * As the project progressed I noticed a lot of the key dashboards were missing or misconfigured and the prometheus metrics werent being visualized. This is a bottle neck to the rest of the team as they need to view the data. 
    * A
      * I reached out to the engineer who was tasked to do this, turns out they were new to how PromQL worked and how grafana and prometheus worked. I offered to set up a meeting with him just a 1:1 and we went over how promql worked how grafana and prometheus worked why we chose this stack, and even made a dashboard with him in kind of like a pair programming type of way. I showed him how to create dashboards, make custom promql statements, how to dynamically label the points on the charts, and about as much as I could show him. 
    * R
      * Now that the engineer had a clear understanding of how to use prometheus, why we used it, and how to make the grafana dashboards, We were able to deliver relevant dashboards for all of the machines in the time that we had it. The dashboards became critical, allowing the test engineers to not have to be present on the actual site, combined with the other software I wrote, we sped up the amount of tests that could be done and the reliability of our data. 
  * when you had a disagreement with the team, what were the steps you took 
    * S
      * At Db control, my team and I were discussing what monitoring stack we should go with. Some of the team mates wanted to use a commercial solution like Datadog as it promised greater up time, quicker setup, and out the box dashboards, but I was advocating for a different approach. I wanted to go with a self hosted approach more specifically Grafana and Prometheus, becuase we had sensitive government data, and for our scale a self hosted approach would fit our needs best. 
    * T
      * I needed to convince my team that Grafana and Prometheus was what we should go with, but I also wanted to hear them out because me being wrong is always a possibility. I wanted to show them that Grafana and prometheus would be cheaper, integrate seemlessly in the current infrastructure, 
    * A 
      * I prepared a detailed comparison that I would go over with the team during our weekly meetings. The main things I covered were Cost,, integration, and made a small poc of it. I showed them that the costs of running Datadog were much higher than Grafana and prometheus becuase well it was free. I showed them how easy it was to integrate prometheus into the current stack I had deployed which were raspberry pis. I also implemented it with mock data on a server I had at home and showed them through a VPN I made. There were concerns but I talked through them by proposing documentation and knowledge sharing sessions I could have.
    * R
      * I convinced the team to go with Grafana and prometheus, we adopted prometheus and Grafana which allowed us to remotely monitor our machines, and I even onboarded the rest of the machines. This process also fostered a much better work culture there as engineers didnt have to sit next to the machines writing things down on paper as their form of collecting data. We also didnt have to worry abotu the costs of Data Dog. 
  * when you had to make a difficult decision under pressure.
    * S
      * While working at PayPal I worked on the MRE team, which was basically Merchant Reliability Engineers and I worked on the development side of things. We had just began integrating several key merchants including Shopify, and all of their merchants, and Apple Pay. Throughout this process we got a lot of false Alerts because of the new integration of merchants ie Shopify and Apple Pay.
    * T
      * I had to quickly decide what to do and how to manage these alerts because I was getting messages from the SRE team that these alerts are mostly false and theres nothing to worry about. They were getting overwhelmed with the amount of alerts, and I had to quickly decide how can I fix these.
    * A
      * I was in fact under a lot of pressure at this time, but I acted quick. I analyzed Datadog dashboards and their corresponding alert patterns and noticed that the thresholds were a lot lower than what they should be due to not taking into account the fact that we are onboarding many more merchants. I suppressed the ones that I deemed problematic and quickly realised that I could make what are called dynamic thresholds using the rolling standard deviation of the metric. I then did that for the problematic datadog dashboards. Got approval from my manager to roll it into prod. During this time of suppressed alerts, I had to tell them to manually look at the alerts but it will only be for a couple of hours.
    * R
      * This approach quickly returned our number of alerts to what they were before we started the integration of new merchants. It actually improved the accuracy of the alerts than what we had before because for some reason we had hard coded values, but after changing that. Our false positives went down by a lot. 
  * when you handled the problem outside the scope of work
    * S
      * During PayPal I was a Software Engineer on the MRE team, I was not on the SRE team so I was not involved much if at all in the provisioning of the Control Room servers, what would go on it, or even if I could see how these were deployed, there was communication between us as we were all in the same team but as far as what containers were deployed, where they were hosted, I was not involved in. But one day in the slack channel they asked for help if anyone knew how to automate some of their deployments with ansible as provisioning 6 servers with all of the packages/repos/deployments they needed would take time if done manually.
    * T
      * Although I wasnt on the SRE team and the provisioning of those servers wasnt a part of my job, I noticed they needed help and I believe that my knowledge of ansible would help them. They were also under a time crunch and I know they could use my help.
    * A
      * I replied offering some help and that I believe ansible could help with this. After being granted temporary access to the repos they had and where they hosted the images, I worked side by side with the SRE team to provision these servers. I wrote ansible playbooks and had a knowledge session with them about why I thougth ansible would be a good pick for this and about what it was. I also documented everything I did in a github repo I owned and shared with them just so I never lost it and I could always share/edit it if they had more questions even after my access was revoked. 
     * R
       * With my help the SRE team deployed the full cluster a bit ahead of schedule, their main goal was to be on time. The environments were setup and made so adding new server would be easy with the playbooks and documentation I wrote. Even though this wasnt in my job scope, I saw an adjacent area that I could step in and be of value, and help get my coworkers on pace and even ahead. 
  * when you had to pivot on a project that was mostly complete 
    * S
      * During College I was given an oppurtunity to work for a Stealth Fintech Startup that was in such early stages they were looking for a talented college student they could pay with zelle. No HR no payroll just straight code this Ill give you zelle. I was ecstatic I mean what an oppurtunity to work full time in college. It was a startup trying to build a venmo like platform that operated on top of crypto, meaning you could send money internationally because it would be sending digital currency. I had nearly finished the back end of this project, I wrote the signup process, the process to creat wallets, the process to send money, and to convert real money to crypto. 
    * T
      * Right before on boarding some early test users, the founders realised that this could be a huge legal issue, the movement of crypto across countries by our platform could be seen as a huge legal issue. We had to pivot quickly.
    * A
      * I rewrote the back end so that the users couldnt actually convert money into crypto, it was all digital kind of spoof money that we worked with. Instead of generating new wallets for users on our platform, I set it up so people could just link their wallets. So it wasnt on us to send the money. This reduced complexity and made it so we could start early testing. 
    * R
      * Within a bout a month of talking to users to see how they were liking our "app", we were able to do a soft launch with limited beta users that we knew could send money across countries. The new pivot which was delinking banks and crypto, where users would just sync their wallet. The startup ultimately didnt take off, and we had to sunset the idea, the pivot was well received and a step in the right direction. 
  * when you convinced the team to agree with you but outcome was negative 
    * take a look at the videos
    * low ego
  * when you pushed back against the project or aspect of the project you disagreed with 
  * when you had a project worked on that would have gone better without your involvment
* what are the 
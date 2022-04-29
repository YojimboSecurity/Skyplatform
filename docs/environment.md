# Environment

This document describes the environment command, its purpose, and how to use it.
In Skytap an environment can be a single virtual machine or a collection of
virtual machines. For example, you can create an environment that contains a
Windows Server VM and several Windows client VMs. Each of these VMs can be run
in isolation or in a cluster.

## Purpose

The environment command is used to get information about a Skytap environment.
Some Skytap environments have a lot of information about them and are made up of
several virtual machines. Each virtual machine has a lot of information
describing it. Some of it useful, some of it not so much. The environment
command does not filter the information it displays.

## How To

To use any command you must have a Skytap account. You can create a Skytap by
going to your supervisor and requesting one. Once you have an account must have
the proper evnironment variables set.

Bash and Zsh users can set the environment variables by adding the following:

```bash
export SKYTAP_USER="username"
export SKYTAP_TOKEN="token"
```

Fish users can set the environment variables by adding the following:

```bash
set -x SKYTAP_USER username
set -x SKYTAP_TOKEN token
```

To get the environment information you can run the following command:

```bash
skyplatform environment --configuration-id <configuration_id>
```


## To Do

There is a lot of information that can be displayed some of it is more useful
than others. The environment command should be able to filter the information.
For example, if you only want to see the information about the virtual machines.
Or if you only want to see the information about the virtual machines networks.
Or even more granularly if you want to see the information about the virtual
machines networks interfaces and what IP addresses, public IP addresses, and
published services are assigned to them. One may just want to see the
environment state. One may also want to start, stop, suspend, or kill one or all
of the virtual machines in the environment.

Environment example:

```bash
Key                        Value
----                       ------
storage                    61440
network_count              2
error_details              []
svms                       6
sequencing_enabled         false
vpn_count                  0
all_vms_support_suspend    true
shutdown_at_time           <nil>
suspend_on_idle            <nil>
disable_internet           false
can_lock                   true
runstate                   suspended
vm_count                   3
tag_list                   
auto_shutdown_description  After 1 hour of idle time
licensed_vm_count          0
vms                         3
                           Key                       Value
                           ---                       -----
                           local_mouse_cursor        true
                           can_change_object_state   true
                           configuration_url         https://cloud.skytap.com/v2/configurations/122243460
                           hardware                  map[architecture:x86 boot_option:bios copy_paste_enabled:true cpus:1 cpus_per_socket:1 disks:[map[controller:0 id:disk-60926776-93682335-scsi-0-0 lun:0 size:20480 type:SCSI]] guestOS:otherlinux-64 max_cpus:24 max_ram:524288 min_ram:256 nested_virtualization:false ram:2048 rtc_start_time:<nil> storage:20480 supports_multicore:true svms:2 time_sync_enabled:true upgradable:false uuid:<nil> vnc_keymap:us]
                           error                     false
                           error_details             false
                           notes                     []
                           labels                    []
                           region_backend            skytap
                           containers                <nil>
                           licenses                  []
                           rate_limited              false
                           hardware_version          11
                           interfaces                [map[hostname:host-1 id:nic-60926776-93682335-0 ip:10.0.1.1 mac:00:50:56:2C:BD:65 network_id:67927308 network_name:custom network_subnet:10.0.1.0/24 network_type:automatic network_url:https://cloud.skytap.com/v2/configurations/122243460/networks/67927308 nic_type:e1000 public_ip_attachments:[] public_ips:[] public_ips_count:0 secondary_ips:[] services:[map[external_ip:services-uscentral.skytap.com external_port:18211 id:22 internal_port:22] map[external_ip:services-uscentral.skytap.com external_port:14266 id:443 internal_port:443]] services_count:2 status:Powered off vm_id:90540245 vm_name:Rodimus] map[hostname:host-1 id:nic-60926776-93682335-1 ip:10.0.0.1 mac:00:50:56:3A:E5:0C network_id:67927309 network_name:nat network_subnet:10.0.0.0/24 network_type:automatic network_url:https://cloud.skytap.com/v2/configurations/122243460/networks/67927309 nic_type:e1000 public_ip_attachments:[] public_ips:[] public_ips_count:0 secondary_ips:[] services:[] services_count:0 status:Powered off vm_id:90540245 vm_name:Rodimus]]
                           license_types             []
                           desktop_resizable         true
                           environment_locked        false
                           created_at                2022/03/16 10:25:01 -0400
                           supports_suspend          true
                           id                        90540245
                           name                      Rodimus
                           asset_id                  <nil>
                           max_hardware_version      11
                           credentials               []
                           runstate                  stopped
                           maintenance_lock_engaged  false
                           
                           id                        90540259
                           hardware_version          11
                           max_hardware_version      11
                           notes                     []
                           license_types             []
                           containers                <nil>
                           hardware                  map[architecture:x86 boot_option:bios copy_paste_enabled:true cpus:1 cpus_per_socket:1 disks:[map[controller:0 id:disk-60926776-93682349-scsi-0-0 lun:0 size:20480 type:SCSI]] guestOS:otherlinux-64 max_cpus:24 max_ram:524288 min_ram:256 nested_virtualization:false ram:2048 rtc_start_time:<nil> storage:20480 supports_multicore:true svms:2 time_sync_enabled:true upgradable:false uuid:<nil> vnc_keymap:us]
                           error_details             false
                           asset_id                  <nil>
                           interfaces                [map[hostname:host-2 id:nic-60926776-93682349-0 ip:10.0.1.2 mac:00:50:56:02:7C:50 network_id:67927308 network_name:custom network_subnet:10.0.1.0/24 network_type:automatic network_url:https://cloud.skytap.com/v2/configurations/122243460/networks/67927308 nic_type:e1000 public_ip_attachments:[map[address:184.170.232.119 connect_type:static dns_name:<nil> hostname:<nil> id:pip_attach-35808395-bc92dea0b6604fd8b5021401d55e674c public_ip_attachment_key:pip_attach-35808395-bc92dea0b6604fd8b5021401d55e674c public_ip_key:184.170.232.119]] public_ips:[map[address:184.170.232.119 id:184.170.232.119]] public_ips_count:1 secondary_ips:[] services:[map[external_ip:services-uscentral.skytap.com external_port:20431 id:22 internal_port:22]] services_count:1 status:Suspended vm_id:90540259 vm_name:Primus] map[hostname:host-2 id:nic-60926776-93682349-1 ip:10.0.0.2 mac:00:50:56:07:55:E3 network_id:67927309 network_name:nat network_subnet:10.0.0.0/24 network_type:automatic network_url:https://cloud.skytap.com/v2/configurations/122243460/networks/67927309 nic_type:e1000 public_ip_attachments:[] public_ips:[] public_ips_count:0 secondary_ips:[] services:[] services_count:0 status:Suspended vm_id:90540259 vm_name:Primus]]
                           maintenance_lock_engaged  false
                           created_at                2022/03/16 10:25:44 -0400
                           configuration_url         https://cloud.skytap.com/v2/configurations/122243460
                           runstate                  suspended
                           rate_limited              false
                           error                     false
                           desktop_resizable         true
                           can_change_object_state   true
                           licenses                  []
                           environment_locked        false
                           name                      Primus
                           labels                    []
                           credentials               []
                           local_mouse_cursor        true
                           region_backend            skytap
                           supports_suspend          true
                           
                           hardware                  map[architecture:x86 boot_option:bios copy_paste_enabled:true cpus:1 cpus_per_socket:1 disks:[map[controller:0 id:disk-60926776-93682385-scsi-0-0 lun:0 size:20480 type:SCSI]] guestOS:otherlinux-64 max_cpus:24 max_ram:524288 min_ram:256 nested_virtualization:false ram:2048 rtc_start_time:<nil> storage:20480 supports_multicore:true svms:2 time_sync_enabled:true upgradable:false uuid:<nil> vnc_keymap:us]
                           error                     false
                           interfaces                [map[hostname:host-3 id:nic-60926776-93682385-0 ip:10.0.1.3 mac:00:50:56:26:94:7F network_id:67927308 network_name:custom network_subnet:10.0.1.0/24 network_type:automatic network_url:https://cloud.skytap.com/v2/configurations/122243460/networks/67927308 nic_type:e1000 public_ip_attachments:[] public_ips:[] public_ips_count:0 secondary_ips:[] services:[] services_count:0 status:Powered off vm_id:90540295 vm_name:Megatron] map[hostname:host-3 id:nic-60926776-93682385-1 ip:10.0.0.3 mac:00:50:56:31:00:04 network_id:67927309 network_name:nat network_subnet:10.0.0.0/24 network_type:automatic network_url:https://cloud.skytap.com/v2/configurations/122243460/networks/67927309 nic_type:e1000 public_ip_attachments:[] public_ips:[] public_ips_count:0 secondary_ips:[] services:[] services_count:0 status:Powered off vm_id:90540295 vm_name:Megatron]]
                           can_change_object_state   true
                           containers                <nil>
                           desktop_resizable         true
                           supports_suspend          true
                           environment_locked        false
                           id                        90540295
                           name                      Megatron
                           runstate                  stopped
                           asset_id                  <nil>
                           hardware_version          11
                           max_hardware_version      11
                           credentials               []
                           license_types             []
                           maintenance_lock_engaged  false
                           region_backend            skytap
                           rate_limited              false
                           error_details             false
                           notes                     []
                           labels                    []
                           local_mouse_cursor        true
                           created_at                2022/03/16 10:27:10 -0400
                           licenses                  []
                           configuration_url         https://cloud.skytap.com/v2/configurations/122243460
                           
environment_locked         false
suspend_at_time            <nil>
project_count              0
schedule_count             0
prefer_local_routing       false
Networks                    2
                           Key                   Value
                           ---                   -----
                           primary_nameserver    <nil>
                           id                    67927308
                           subnet                10.0.1.0/24
                           region                US-Central
                           gateway               10.0.1.254
                           secondary_nameserver  <nil>
                           domain                example1.com
                           vpn_attachments       []
                           url                   https://cloud.skytap.com/v2/configurations/122243460/networks/67927308
                           name                  custom
                           network_type          automatic
                           subnet_size           24
                           tunnelable            false
                           tunnels               []
                           subnet_addr           10.0.1.0
                           
                           gateway               10.0.0.254
                           url                   https://cloud.skytap.com/v2/configurations/122243460/networks/67927309
                           name                  nat
                           primary_nameserver    <nil>
                           secondary_nameserver  <nil>
                           domain                example.com
                           tunnelable            false
                           id                    67927309
                           network_type          automatic
                           subnet_size           24
                           vpn_attachments       []
                           tunnels               []
                           subnet                10.0.0.0/24
                           subnet_addr           10.0.0.0
                           region                US-Central
                           
container_hosts_count      0
can_edit                   true
public_ip_count            1
can_tag                    true
svms_by_architecture       map[power:0 x86:6]
environment_lock           <nil>
id                         122243460
publish_set_count          0
region                     US-Central
tags                       []
project_count_for_user     0
routable                   false
rate_limited               false
created_at                 2022/03/16 10:24:42 -0400
label_count                0
label_category_count       0
alerts                     []
staged_execution           <nil>
containers_count           0
region_backend             skytap
can_save_as_template       true
platform_errors            []
owner_id                   544314
note_count                 0
outbound_traffic           false
shutdown_on_idle           3600
description                
owner_url                  https://cloud.skytap.com/v2/users/544314
owner_name                 David Johnson
can_copy                   true
can_change_state           true
can_share                  true
can_change_owner           false
stages                     <nil>
url                        https://cloud.skytap.com/v2/configurations/122243460
name                       IntegrationTesterSuite
errors                     []
last_run                   2022/04/19 13:12:27 -0400
can_delete                 true
published_service_count    3
auto_suspend_description   <nil>
```

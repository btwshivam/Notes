# Week 2 — Linux Fundamentals (Module 1–3)
---

## Table of Contents

### Module 1 — Linux Fundamentals
1. [Introduction to Linux](#1-introduction-to-linux)
2. [Components of Linux](#2-components-of-linux)
3. [Linux Kernel Overview](#3-linux-kernel-overview)
4. [Linux Architecture / Structure](#4-linux-architecture--structure)
5. [Linux Distributions](#5-linux-distributions)
6. [Environment Setup](#6-environment-setup)
7. [Terminal Basics](#7-terminal-basics)
8. [First Linux Commands](#8-first-linux-commands)
9. [Package Manager](#9-package-manager)

### Module 2 — File System & Directory Structure
10. [Linux File System Hierarchy](#10-linux-file-system-hierarchy)
11. [Important Directories](#11-important-directories)
12. [Paths in Linux](#12-paths-in-linux)
13. [Navigation Concepts](#13-navigation-concepts)
14. [Understanding File Types](#14-understanding-file-types)

### Module 3 — File Management (CRUD)
15. [Create Operations](#15-create-operations)
16. [Read Operations](#16-read-operations)
17. [Update Operations](#17-update-operations)
18. [Delete Operations](#18-delete-operations)
19. [Copy & Move](#19-copy--move)
20. [Wildcards & Patterns](#20-wildcards--patterns)
21. [Editors — Nano & Vim](#21-editors--nano--vim)

---

# MODULE 1 — Linux Fundamentals

---

## 1. Introduction to Linux

### 1.1 What is Linux?

**Definition**: Linux is a free, open-source, Unix-like operating system kernel first created by **Linus Torvalds** in **1991**. When combined with GNU tools and other software, it forms a complete operating system commonly referred to as **GNU/Linux**.


**Key Facts:**
| Property | Detail |
|---|---|
| Creator | Linus Torvalds |
| Initial Release | September 17, 1991 |
| License | GNU General Public License v2 (GPLv2) |
| Written In | C, Assembly |
| Type | Monolithic Kernel |
| Mascot | Tux (the penguin) |

### 1.2 Kernel vs Operating System

Understanding the difference between a **kernel** and an **operating system** is fundamental.

```
┌─────────────────────────────────────────┐
│            APPLICATION LAYER            │
│    (Browser, Editor, Terminal, etc.)    │
├─────────────────────────────────────────┤
│           OPERATING SYSTEM              │
│  (System Libraries, Shell, Utilities)   │
├─────────────────────────────────────────┤
│              KERNEL                     │
│  (Process, Memory, Device Management)   │
├─────────────────────────────────────────┤
│             HARDWARE                    │
│    (CPU, RAM, Disk, Network, I/O)       │
└─────────────────────────────────────────┘
```
![linux_os](https://github.com/user-attachments/assets/093262a6-66c4-4aa1-a947-a4f0ce7a242b)

| Aspect | Kernel | Operating System |
|---|---|---|
| **Definition** | Core component that manages hardware resources | Complete system including kernel + utilities + libraries |
| **Function** | Low-level hardware abstraction | User-facing environment |
| **Scope** | Process scheduling, memory, drivers | File management, GUI, networking, applications |
| **Example** | Linux kernel (5.x, 6.x) | Ubuntu, Fedora, Debian |
| **User Interaction** | No direct user interaction | Users interact through shell/GUI |
| **Size** | Few MB | Several GB |

**Analogy**: The kernel is the **engine** of a car. The operating system is the **entire car** — engine + body + dashboard + seats + everything else.

### 1.3 Open-Source Concept

**Open Source** means the source code is freely available for anyone to:
- **View** — Inspect how the software works internally
- **Modify** — Change or improve the code
- **Distribute** — Share original or modified versions
- **Use** — Run for any purpose without restrictions

**Why Linux Being Open-Source Matters:**
1. **Security** — Thousands of developers audit the code; vulnerabilities are found and patched quickly
2. **Customization** — Anyone can tailor Linux to their exact needs
3. **No Vendor Lock-in** — Freedom to switch, modify, or host anywhere
4. **Cost** — Free to use, even for enterprise production environments
5. **Community** — Massive global community contributing patches, drivers, and features

**Popular Open-Source Licenses:**
| License | Key Feature |
|---|---|
| GPL v2 | Must share derivative work (Linux Kernel uses this) |
| MIT | Very permissive, minimal restrictions |
| Apache 2.0 | Permissive + patent protection |
| BSD | Permissive, similar to MIT |

### 1.4 Use Cases of Linux

Linux dominates several critical technology domains:

```
Use Cases of Linux
├── Servers & Cloud
│   ├── 96.3% of top 1 million web servers run Linux
│   ├── AWS, Azure, GCP default to Linux instances
│   └── Apache, Nginx, databases all run on Linux
├── DevOps & CI/CD
│   ├── Docker containers are Linux-based
│   ├── Kubernetes orchestrates Linux containers
│   └── Jenkins, GitLab CI run on Linux
├── Cybersecurity
│   ├── Kali Linux — penetration testing
│   ├── Parrot OS — security auditing
│   └── Most security tools are Linux-native
├── Embedded Systems & IoT
│   ├── Routers, smart TVs, cameras
│   ├── Raspberry Pi runs Linux
│   └── Automotive infotainment systems
├── Android
│   ├── Built on Linux kernel
│   └── 3+ billion devices worldwide
├── Supercomputers
│   └── 100% of TOP500 supercomputers run Linux
└── Desktop
    ├── Developer workstations
    ├── Education environments
    └── Privacy-focused users
```

**Industry Adoption:**
| Company | Linux Usage |
|---|---|
| Google | Entire infrastructure, Android, ChromeOS |
| Amazon | AWS runs on Linux |
| Facebook/Meta | All servers run Linux |
| Netflix | FreeBSD/Linux for content delivery |
| Tesla | Autopilot runs on Linux |
| NASA | Mars rovers run Linux |

---

## 2. Components of Linux

Linux is built from five major components working together as a cohesive system.

### 2.1 Architecture Overview

```
┌──────────────────────────────────────────────────┐
│                 APPLICATIONS                      │
│  (Firefox, VS Code, Docker, Apache, MySQL, etc.) │
├──────────────────────────────────────────────────┤
│                  UTILITIES                        │
│  (grep, awk, sed, find, tar, curl, wget, etc.)   │
├──────────────────────────────────────────────────┤
│                   SHELL                           │
│  (Bash, Zsh, Fish — Command interpreter)         │
├──────────────────────────────────────────────────┤
│               FILE SYSTEM                         │
│  (ext4, XFS, Btrfs, ZFS — Data organization)     │
├──────────────────────────────────────────────────┤
│                  KERNEL                           │
│  (Process, Memory, Device, Network management)    │
├──────────────────────────────────────────────────┤
│                 HARDWARE                          │
│  (CPU, RAM, Storage, NIC, GPU, Peripherals)       │
└──────────────────────────────────────────────────┘
```
![linux_architecture](https://github.com/user-attachments/assets/25fe4614-25f8-4c67-bd6d-d4282119fb81)

### 2.2 Kernel

The **kernel** is the heart of Linux — it is the first program loaded when the system boots.

**Responsibilities:**
- **Process Management** — Creating, scheduling, and terminating processes
- **Memory Management** — Allocating/deallocating RAM, virtual memory, paging
- **Device Drivers** — Communicating with hardware (disk, network, USB, GPU)
- **System Calls** — Providing API for user-space programs to request kernel services
- **Security** — Enforcing access control, user permissions, SELinux/AppArmor

**Type**: Linux uses a **monolithic kernel** — all core services run in kernel space for performance, but it supports **loadable kernel modules (LKMs)** for flexibility.

### 2.3 Shell

The **shell** is the command-line interpreter — the bridge between the user and the kernel.

**How It Works:**
```
User types command → Shell interprets it → Kernel executes it → Output returned
```

**Popular Shells:**
| Shell | Full Name | Key Feature |
|---|---|---|
| `bash` | Bourne Again Shell | Default on most distros, scripting powerhouse |
| `zsh` | Z Shell | Advanced autocompletion, themes (Oh My Zsh) |
| `fish` | Friendly Interactive Shell | User-friendly, syntax highlighting |
| `sh` | Bourne Shell | Original Unix shell, POSIX compliant |
| `dash` | Debian Almquist Shell | Lightweight, fast for scripts |

**Check your current shell:**
```bash
echo $SHELL
# Output: /bin/bash
```

### 2.4 File System

The file system organizes and stores data on disk. In Linux, **everything is a file** — even hardware devices.

**Common Linux File Systems:**
| File System | Use Case | Max File Size | Max Volume Size |
|---|---|---|---|
| **ext4** | Default for most distros | 16 TB | 1 EB |
| **XFS** | High-performance servers (RHEL default) | 8 EB | 8 EB |
| **Btrfs** | Snapshots, RAID, modern features | 16 EB | 16 EB |
| **ZFS** | Enterprise storage, data integrity | 16 EB | 256 ZB |
| **tmpfs** | Temporary files in RAM | RAM-limited | RAM-limited |

### 2.5 Utilities

Linux provides hundreds of small, focused command-line tools following the **Unix Philosophy**:

> *"Do one thing and do it well."* — Doug McIlroy

**Categories:**
| Category | Examples |
|---|---|
| File Operations | `cp`, `mv`, `rm`, `mkdir`, `touch` |
| Text Processing | `grep`, `awk`, `sed`, `cut`, `sort` |
| System Info | `uname`, `uptime`, `df`, `free`, `top` |
| Networking | `ping`, `curl`, `wget`, `ssh`, `netstat` |
| Compression | `tar`, `gzip`, `zip`, `bzip2` |
| User Management | `useradd`, `passwd`, `chmod`, `chown` |

### 2.6 Applications

User-level software that runs on top of the Linux ecosystem:
- **Web Servers** — Apache, Nginx
- **Databases** — MySQL, PostgreSQL, MongoDB
- **Containers** — Docker, Podman
- **Editors** — Vim, Nano, VS Code
- **Browsers** — Firefox, Chromium

---

## 3. Linux Kernel Overview

### 3.1 Role of the Kernel

The kernel is the **lowest-level software** that interfaces directly with hardware. It acts as a **resource manager** and **abstraction layer**.

<img width="620" height="658" alt="ditaa-b9ffae65be16d30be11b5eca188a7a143b1b8227 (1)" src="https://github.com/user-attachments/assets/5589de7e-94c4-42ce-959b-8bdadf6c9427" />


### 3.2 Process Management

A **process** is a running instance of a program. The kernel manages the entire lifecycle.

**Process Lifecycle:**
```
Created → Ready → Running → Waiting/Blocked → Terminated
                    ↑            │
                    └────────────┘
```

**Key Concepts:**
| Concept | Description |
|---|---|
| **PID** | Process ID — unique identifier for each process |
| **PPID** | Parent Process ID — the process that spawned it |
| **Scheduler** | Decides which process gets CPU time (CFS — Completely Fair Scheduler) |
| **Context Switch** | Saving/restoring process state when switching between processes |
| **Fork** | Creating a child process (copy of parent) |
| **Exec** | Replacing current process image with a new program |
| **Zombie** | Process that has finished but parent hasn't read its exit status |
| **Orphan** | Process whose parent has terminated (adopted by init/systemd) |

**Process States:**
| State | Symbol | Meaning |
|---|---|---|
| Running | `R` | Actively using CPU |
| Sleeping | `S` | Waiting for an event (interruptible) |
| Disk Sleep | `D` | Waiting for I/O (uninterruptible) |
| Stopped | `T` | Paused (e.g., by `Ctrl+Z`) |
| Zombie | `Z` | Finished, awaiting parent cleanup |

### 3.3 Memory Management

The kernel manages both **physical RAM** and **virtual memory**.

**Key Mechanisms:**
| Mechanism | Description |
|---|---|
| **Virtual Memory** | Each process gets its own virtual address space |
| **Paging** | Memory divided into fixed-size pages (typically 4 KB) |
| **Swap Space** | Disk space used when RAM is full (slower but extends capacity) |
| **Page Cache** | Frequently accessed disk data cached in RAM for speed |
| **OOM Killer** | Out-of-Memory killer terminates processes when RAM is critically low |
| **Memory Mapping** | `mmap()` — maps files directly into memory for fast I/O |

**Memory Layout of a Process:**
```
High Address ┌──────────────────┐
             │      Stack       │  ← Local variables, function calls
             │        ↓         │
             │                  │
             │        ↑         │
             │       Heap       │  ← Dynamic allocation (malloc)
             ├──────────────────┤
             │       BSS        │  ← Uninitialized global data
             ├──────────────────┤
             │       Data       │  ← Initialized global data
             ├──────────────────┤
             │       Text       │  ← Program code (read-only)
Low Address  └──────────────────┘
```

### 3.4 Device Drivers

Device drivers allow the kernel to communicate with hardware without knowing specific hardware details.

**Types:**
| Type | Description | Example |
|---|---|---|
| **Character Device** | Stream of bytes, no buffering | Keyboard, mouse, serial ports |
| **Block Device** | Fixed-size blocks, buffered I/O | Hard drives, SSDs, USB storage |
| **Network Device** | Packet-based communication | Ethernet, WiFi adapters |

**Device Files in `/dev`:**
```
/dev/sda      → First SATA/SCSI disk
/dev/sda1     → First partition of first disk
/dev/nvme0n1  → First NVMe SSD
/dev/tty      → Terminal device
/dev/null     → Discards all data (black hole)
/dev/zero     → Infinite stream of null bytes
/dev/random   → Random number generator
```

### 3.5 File Systems (Kernel Perspective)

The kernel uses the **Virtual File System (VFS)** as an abstraction layer — applications use the same system calls (`open`, `read`, `write`, `close`) regardless of the underlying file system.

```
Application
    ↓
System Call (open, read, write)
    ↓
Virtual File System (VFS)
    ↓
┌────────┬────────┬────────┬────────┐
│  ext4  │  XFS   │ Btrfs  │  NFS   │
└────────┴────────┴────────┴────────┘
    ↓
Block Device Layer
    ↓
Physical Disk
```

### 3.6 Kernel Versions

Linux kernel follows a **major.minor.patch** versioning scheme.

```
Example: 6.5.3
         │ │ │
         │ │ └── Patch (bug fixes)
         │ └──── Minor (new features, improvements)
         └────── Major (significant changes)
```

**Check your kernel version:**
```bash
uname -r
# Example output: 6.5.0-44-generic
```

**Notable Kernel Milestones:**
| Version | Year | Significance |
|---|---|---|
| 0.01 | 1991 | Linus Torvalds' first release |
| 1.0 | 1994 | First production-ready release |
| 2.6 | 2003 | Major performance improvements, scalability |
| 3.0 | 2011 | Numbering change (20th anniversary) |
| 4.0 | 2015 | Live patching support |
| 5.0 | 2019 | AMD FreeSync, Adiantum encryption |
| 6.0 | 2022 | Performance optimizations, Rust support |

---

## 4. Linux Architecture / Structure

### 4.1 Four-Layer Architecture

Linux follows a **layered architecture** where each layer has a specific role and communicates only with adjacent layers.

<img width="392" height="328" alt="muITn" src="https://github.com/user-attachments/assets/39335d59-ce40-43bd-ad51-e835f480750b" />


### 4.2 Hardware Layer

The **physical components** that the system runs on:

| Component | Role |
|---|---|
| **CPU** | Executes instructions, runs processes |
| **RAM** | Stores data for running processes (volatile) |
| **Storage** | HDD/SSD — persistent data storage |
| **NIC** | Network Interface Card — network connectivity |
| **GPU** | Graphics processing, machine learning workloads |
| **Peripherals** | Keyboard, mouse, USB devices |

### 4.3 Kernel Layer

The kernel sits directly above hardware and provides:
- **Hardware Abstraction** — Programs don't need to know hardware specifics
- **Resource Allocation** — Fair distribution of CPU, memory, I/O
- **Security Enforcement** — User permissions, process isolation
- **System Calls** — ~400+ system calls (e.g., `open()`, `read()`, `write()`, `fork()`, `exec()`)

### 4.4 Shell Layer

The shell is the **user interface to the kernel**:
- Reads user commands
- Parses and interprets them
- Makes appropriate system calls
- Returns results to the user

**Shell can operate in two modes:**
| Mode | Description |
|---|---|
| **Interactive** | User types commands one by one in terminal |
| **Non-Interactive** | Shell executes a script file (`.sh`) |

### 4.5 User Space

Everything that runs **outside the kernel** — applications, libraries, and user processes.

**Key Characteristics:**
- Processes run with limited privileges (cannot directly access hardware)
- Must use **system calls** to request kernel services
- Crashes in user space don't crash the entire system
- Each process has isolated virtual memory

---

## 5. Linux Distributions

A **Linux distribution (distro)** = Linux Kernel + GNU Tools + Package Manager + Desktop Environment + Additional Software.
<img width="758" height="405" alt="58111436-f8e29380-7be0-11e9-9eb6-a554429eeb21" src="https://github.com/user-attachments/assets/cf805646-0592-44ce-8a10-095c8ecbe6a7" />
<img width="512" height="512" alt="Logo_Collage_Linux_Distro" src="https://github.com/user-attachments/assets/14c499e7-d231-4843-8038-94d877fe5c8c" />

### 5.1 Distribution Family Tree

```
         ┌──────────────────────────────────────────────┐
         │              Linux Kernel                    │
         └──────────────────┬───────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
   ┌────▼────┐        ┌────▼────┐        ┌─────▼────┐
   │  Debian  │        │ Red Hat │        │   Arch   │
   │  (.deb)  │        │ (.rpm)  │        │(rolling) │
   └────┬────┘        └────┬────┘        └─────┬────┘
        │                   │                   │
   ┌────▼────┐        ┌────▼────┐        ┌─────▼──────┐
   │ Ubuntu  │        │ Fedora  │        │  Manjaro   │
   └────┬────┘        └────┬────┘        └────────────┘
        │                   │
   ┌────▼────┐        ┌────▼──────┐
   │  Mint   │        │  CentOS   │
   └─────────┘        │  Rocky    │
                      │  AlmaLinux│
                      └───────────┘
```

### 5.2 Major Distributions Compared


o | Base | Package Manager | Best For | Release Model |
|---|---|---|---|---|
| **Ubuntu** | Debian | `apt` | Beginners, servers, cloud | Fixed (LTS every 2 years) |
| **Debian** | Independent | `apt` | Stability, servers | Fixed (slow, stable) |
| **Fedora** | Red Hat | `dnf` | Developers, cutting-edge | Fixed (~6 months) |
| **Arch** | Independent | `pacman` | Advanced users, customization | Rolling release |
| **Kali** | Debian | `apt` | Penetration testing, security | Rolling release |
| **CentOS/Rocky** | Red Hat | `yum`/`dnf` | Enterprise servers | Fixed (long-term) |

### 5.3 Deep Dive — Ubuntu

- **Most popular** desktop and server Linux distro
- **LTS (Long Term Support)** releases every 2 years (supported for 5 years)
- Default desktop: **GNOME**
- Excellent **community and documentation**
- Example LTS versions: 20.04, 22.04, 24.04

### 5.4 Deep Dive — Debian

- **"The Universal Operating System"**
- Known for extreme **stability and reliability**
- Ubuntu is based on Debian
- Three release channels: **Stable**, **Testing**, **Unstable (Sid)**
- Preferred for production servers

### 5.5 Deep Dive — Fedora

- **Sponsored by Red Hat** (IBM)
- Bleeding-edge — latest kernel, GNOME, and technologies
- Proving ground for **RHEL (Red Hat Enterprise Linux)**
- Great for developers wanting the latest software

### 5.6 Deep Dive — Arch Linux

- **DIY philosophy** — "You build it from scratch"
- **Rolling release** — always up to date, no version upgrades
- **Arch Wiki** — the best documentation in the Linux world
- **`pacman`** package manager + AUR (Arch User Repository)
- Not recommended for beginners

### 5.7 Deep Dive — Kali Linux

- Built specifically for **cybersecurity professionals**
- Pre-installed with **600+ security tools** (Nmap, Wireshark, Metasploit, Burp Suite)
- Based on **Debian Testing**
- Should NOT be used as a daily driver OS

### 5.8 Deep Dive — CentOS / Rocky / AlmaLinux

- **CentOS** was a free RHEL clone (discontinued in 2021 → CentOS Stream)
- **Rocky Linux** and **AlmaLinux** are community-driven RHEL replacements
- Used heavily in **enterprise environments**
- 10-year support lifecycle

---

## 6. Environment Setup

### 6.1 VirtualBox Installation

**VirtualBox** is a free, open-source hypervisor by Oracle for running virtual machines.

**Setup Steps:**
```
1. Download VirtualBox from https://www.virtualbox.org/
2. Install on your host OS (Windows/macOS/Linux)
3. Download Ubuntu ISO from https://ubuntu.com/download
4. Create a New Virtual Machine:
   ├── Name: Ubuntu-Lab
   ├── Type: Linux
   ├── Version: Ubuntu (64-bit)
   ├── RAM: 2048 MB (minimum) / 4096 MB (recommended)
   ├── Disk: 25 GB (dynamically allocated)
   └── Network: NAT (default) or Bridged
5. Mount the ISO → Start VM → Follow installation wizard
6. Install Guest Additions for better performance
```

**Recommended VM Settings:**
| Setting | Minimum | Recommended |
|---|---|---|
| RAM | 2 GB | 4 GB |
| CPU Cores | 1 | 2 |
| Disk | 20 GB | 40 GB |
| Video Memory | 16 MB | 128 MB |

### 6.2 WSL Installation (Windows Subsystem for Linux)

**WSL** allows running Linux natively on Windows without a VM.

**WSL2 Installation (Windows 10/11):**
```powershell
# Open PowerShell as Administrator

# Install WSL with Ubuntu (default)
wsl --install

# Or install a specific distro
wsl --install -d Ubuntu-22.04

# List available distros
wsl --list --online

# Check WSL version
wsl --version

# Set default version to WSL2
wsl --set-default-version 2
```

**WSL1 vs WSL2:**
| Feature | WSL1 | WSL2 |
|---|---|---|
| Architecture | Translation layer | Real Linux kernel |
| Performance (filesystem) | Faster on Windows files | Faster on Linux files |
| System Call Compatibility | Limited | Full |
| Docker Support | No | Yes |
| Memory Usage | Lower | Higher |
| Networking | Shares host network | Virtual network |

### 6.3 AWS EC2 Setup

**Amazon EC2** provides virtual Linux servers in the cloud.

**Launch Steps:**
```
1. Sign in to AWS Console → EC2 Dashboard
2. Click "Launch Instance"
3. Configure:
   ├── Name: Linux-Lab
   ├── AMI: Ubuntu 22.04 LTS (Free Tier eligible)
   ├── Instance Type: t2.micro (Free Tier)
   ├── Key Pair: Create new → Download .pem file
   ├── Network: Default VPC
   ├── Security Group:
   │   ├── SSH (port 22) — Your IP only
   │   └── HTTP (port 80) — optional
   └── Storage: 8 GB gp3 (default)
4. Launch Instance
5. Wait for "Running" status
```

### 6.4 SSH Access

**SSH (Secure Shell)** provides encrypted remote access to Linux servers.

**Connect to EC2:**
```bash
# Set correct permissions on key file
chmod 400 my-key.pem

# Connect via SSH
ssh -i my-key.pem ubuntu@<public-ip>

# Example
ssh -i linux-lab.pem ubuntu@54.123.45.67
```

**SSH Key Components:**
| Component | Description |
|---|---|
| **Private Key** (`.pem`) | Stays on your local machine — NEVER share |
| **Public Key** | Stored on the server in `~/.ssh/authorized_keys` |
| **Port** | Default: 22 |
| **Protocol** | Encrypted communication channel |

**SSH Config for Convenience** (`~/.ssh/config`):
```
Host my-server
    HostName 54.123.45.67
    User ubuntu
    IdentityFile ~/.ssh/linux-lab.pem
    Port 22
```
After config, simply connect with:
```bash
ssh my-server
```

---

## 7. Terminal Basics

### 7.1 What is a Terminal?

The **terminal** (terminal emulator) is the application that provides a text-based interface to interact with the shell.

```
┌─────────────────────────────────────────────┐
│ Terminal Emulator                           │
│ ┌─────────────────────────────────────────┐ │
│ │ Shell (Bash)                            │ │
│ │                                         │ │
│ │ shivam@ubuntu:~$                        │ │
│ │                                         │ │
│ └─────────────────────────────────────────┘ │
└─────────────────────────────────────────────┘
```

**Terminal vs Shell vs Console:**
| Term | Definition |
|---|---|
| **Terminal** | GUI application that hosts the shell (e.g., GNOME Terminal, iTerm2, Windows Terminal) |
| **Shell** | Command interpreter program (e.g., Bash, Zsh) |
| **Console** | Physical or virtual text-mode interface (e.g., tty1-tty6) |

### 7.2 Command Syntax

Every Linux command follows a standard structure:

```
command [options] [arguments]
   │       │         │
   │       │         └── What to act on (files, directories, etc.)
   │       └──────────── Modify behavior (flags: -r, -l, --verbose)
   └──────────────────── The program to execute
```

**Examples:**
```bash
ls                    # command only
ls -l                 # command + option
ls -la /home          # command + options + argument
ls --all --human-readable /var/log   # long-form options
```

### 7.3 Flags and Arguments

**Short Flags** (single dash + letter): `-l`, `-a`, `-h`
**Long Flags** (double dash + word): `--list`, `--all`, `--human-readable`

**Flag Behavior:**
```bash
# Flags can be combined
ls -l -a -h    # separate
ls -lah        # combined (same result)

# Flags with values
grep -n "error" file.txt       # -n enables line numbers
tail -n 20 file.txt            # -n takes a value (20 lines)

# Boolean flags (toggle on/off)
rm -r directory/               # -r enables recursive mode
cp -v source dest              # -v enables verbose output
```

**Arguments:**
```bash
# Single argument
cat file.txt

# Multiple arguments
cp source.txt destination.txt
mkdir dir1 dir2 dir3

# Standard input as argument
echo "hello world" | grep "hello"
```

---

## 8. First Linux Commands

### 8.1 Understanding the Prompt

The command prompt provides contextual information:

```
shivam@ubuntu:~$
  │       │    │ │
  │       │    │ └── $ = regular user  |  # = root user
  │       │    └──── ~ = current directory (home)
  │       └───────── hostname (machine name)
  └───────────────── username (logged-in user)
```

**Root vs Regular User:**
| Indicator | User Type | Power Level |
|---|---|---|
| `$` | Regular user | Limited permissions |
| `#` | Root (superuser) | Unrestricted access — **use with caution** |

### 8.2 Basic Command Execution

**How a command is executed:**
```
User types command
        ↓
Shell searches for command in $PATH
        ↓
Found? → Kernel loads and executes the program
        ↓
Output is sent back to terminal (stdout)
Errors are sent to terminal (stderr)
        ↓
Shell displays the prompt again, ready for next command
```

**Essential First Commands:**
```bash
# Who am I?
whoami                    # Output: shivam

# Where am I?
pwd                       # Output: /home/shivam

# What's in this directory?
ls                        # List files and directories

# What system am I on?
uname -a                  # Full system information
hostname                  # Machine name

# What's the date/time?
date                      # Current date and time
cal                       # Calendar of current month
uptime                    # How long the system has been running

# Clear the screen
clear                     # Or press Ctrl + L

# Command history
history                   # Show previously executed commands

# Exit terminal
exit                      # Close the terminal session
```

### 8.3 Manual Pages (man)

The **`man`** command is the built-in documentation system — the ultimate reference for every command.

**Usage:**
```bash
man <command>

# Examples
man ls          # Full documentation for 'ls'
man grep        # Full documentation for 'grep'
man ssh         # Full documentation for 'ssh'
```

**Man Page Structure:**
| Section | Content |
|---|---|
| **NAME** | Command name and brief description |
| **SYNOPSIS** | Usage syntax |
| **DESCRIPTION** | Detailed explanation |
| **OPTIONS** | All available flags and options |
| **EXAMPLES** | Usage examples |
| **SEE ALSO** | Related commands |
| **EXIT STATUS** | Return codes |

**Navigation Inside Man Pages:**
| Key | Action |
|---|---|
| `Space` / `f` | Forward one page |
| `b` | Back one page |
| `q` | Quit |
| `/pattern` | Search forward |
| `n` | Next search result |
| `N` | Previous search result |
| `g` | Go to beginning |
| `G` | Go to end |

**Quick Help Alternative:**
```bash
# Short help for any command
ls --help
grep --help

# One-line description
whatis ls        # Output: ls (1) - list directory contents

# Search man pages by keyword
apropos "copy files"
```

---

## 9. Package Manager

### 9.1 What is a Package Manager?

A **package manager** is a tool that automates installing, updating, configuring, and removing software packages.

**Without Package Manager:**
```
Download source → Install dependencies → Compile → Configure → Install → Update manually
```

**With Package Manager:**
```bash
sudo apt install nginx    # One command does everything
```

**Package Manager Responsibilities:**
- **Dependency Resolution** — Automatically installs required libraries
- **Version Management** — Tracks installed versions, handles upgrades
- **Repository Management** — Downloads from trusted sources
- **Integrity Verification** — Checks package signatures (GPG)
- **Clean Removal** — Removes software and its dependencies

### 9.2 APT (Advanced Package Tool) — Debian/Ubuntu

**`apt`** is the standard package manager for Debian-based distributions.

**Package Lifecycle Commands:**
```bash
# Update package index (fetch latest package lists)
sudo apt update

# Upgrade all installed packages
sudo apt upgrade

# Update + Upgrade (full system update)
sudo apt update && sudo apt upgrade -y

# Install a package
sudo apt install <package-name>
sudo apt install nginx
sudo apt install git curl wget     # Multiple packages

# Remove a package (keeps config files)
sudo apt remove <package-name>

# Remove completely (including config files)
sudo apt purge <package-name>

# Remove unused dependencies
sudo apt autoremove

# Search for packages
apt search <keyword>
apt search "web server"

# Show package details
apt show nginx

# List installed packages
apt list --installed

# List upgradable packages
apt list --upgradable
```

**Understanding `apt` vs `apt-get`:**
| Feature | `apt` | `apt-get` |
|---|---|---|
| Purpose | End-user friendly | Scripting / backward compatible |
| Progress Bar | Yes | No |
| Color Output | Yes | No |
| Recommendation | **Use `apt`** for interactive use | Use `apt-get` in scripts |

### 9.3 Snap

**Snap** is a universal package format by Canonical that bundles dependencies.

```bash
# Install a snap
sudo snap install <package>
sudo snap install code --classic       # VS Code
sudo snap install discord

# List installed snaps
snap list

# Update all snaps
sudo snap refresh

# Remove a snap
sudo snap remove <package>

# Find snaps
snap find <keyword>
```

**APT vs Snap:**
| Feature | APT | Snap |
|---|---|---|
| Packages | `.deb` files | Self-contained bundles |
| Dependencies | Shared system libraries | Bundled inside snap |
| Updates | Manual (`apt upgrade`) | Automatic |
| Sandboxing | No | Yes (confined) |
| Size | Smaller | Larger (bundles deps) |
| Speed | Faster startup | Slower startup |

### 9.4 YUM / DNF Overview — Red Hat/Fedora

**`yum`** (Yellowdog Updater Modified) and its successor **`dnf`** (Dandified YUM) are for RPM-based distributions.

```bash
# DNF Commands (Fedora, RHEL 8+, Rocky, AlmaLinux)
sudo dnf update                     # Update all packages
sudo dnf install <package>          # Install
sudo dnf remove <package>           # Remove
sudo dnf search <keyword>           # Search
dnf info <package>                  # Package info
dnf list installed                  # List installed

# YUM Commands (CentOS 7, older RHEL) — same syntax
sudo yum update
sudo yum install <package>
sudo yum remove <package>
```

**Package Manager Summary by Distro:**
| Distro | Package Format | Package Manager | Command |
|---|---|---|---|
| Ubuntu/Debian | `.deb` | APT | `apt install` |
| Fedora | `.rpm` | DNF | `dnf install` |
| CentOS 7 | `.rpm` | YUM | `yum install` |
| Arch | `.pkg.tar.zst` | Pacman | `pacman -S` |
| openSUSE | `.rpm` | Zypper | `zypper install` |

---

# MODULE 2 — File System & Directory Structure

---

## 10. Linux File System Hierarchy

### 10.1 Root Directory `/`

In Linux, **everything starts from root (`/`)**. Unlike Windows which has drive letters (C:\, D:\), Linux has a **single unified tree** structure.

![standard-unix-filesystem-hierarchy-1 (1)](https://github.com/user-attachments/assets/abe1d251-8144-4862-aa00-46028ac2e9a6)

```
/                          ← Root — the top of the entire file system
├── bin/                   ← Essential user binaries
├── boot/                  ← Boot loader files (kernel, GRUB)
├── dev/                   ← Device files
├── etc/                   ← System configuration files
├── home/                  ← User home directories
├── lib/                   ← Shared libraries
├── media/                 ← Mount points for removable media
├── mnt/                   ← Temporary mount points
├── opt/                   ← Optional/third-party software
├── proc/                  ← Virtual filesystem — process info
├── root/                  ← Root user's home directory
├── run/                   ← Runtime data
├── sbin/                  ← System binaries (admin commands)
├── srv/                   ← Service data (web, FTP)
├── sys/                   ← Virtual filesystem — kernel/hardware info
├── tmp/                   ← Temporary files (cleared on reboot)
├── usr/                   ← User programs and data
└── var/                   ← Variable data (logs, caches, mail)
```

### 10.2 FHS (Filesystem Hierarchy Standard)

The **Filesystem Hierarchy Standard** defines the directory structure and directory contents in Unix-like operating systems. All major Linux distros follow FHS.

**Key Principle**: Linux separates files by **purpose**, not by application.

| Windows Approach | Linux Approach |
|---|---|
| `C:\Program Files\App\app.exe` | `/usr/bin/app` (binary) |
| `C:\Program Files\App\config.ini` | `/etc/app/config` (config) |
| `C:\Program Files\App\logs\` | `/var/log/app/` (logs) |

### 10.3 Hierarchical Design

```
                    /  (root)
                    │
        ┌───────────┼───────────┐
        │           │           │
      /home       /etc        /var
        │           │           │
    ┌───┼───┐     ┌─┼─┐      ┌─┼──┐
  shivam  raj   nginx ssh   log cache
    │
  ┌─┼──┐
docs  projects
```
<img width="728" height="590" alt="0_bFnHaO8eYpW3dSuz" src="https://github.com/user-attachments/assets/83a59f80-46c8-483d-8273-ce88864c5172" />


**Key Concepts:**
- **Single root** — One starting point (`/`), unlike Windows drive letters
- **Tree structure** — Hierarchical parent-child relationships
- **Mount points** — Different physical disks can be mounted anywhere in the tree
- **Everything is a file** — Devices, processes, pipes — all represented as files

---

## 11. Important Directories

### 11.1 `/home` — User Home Directories

```
/home/
├── shivam/           ← Your personal space
│   ├── Documents/
│   ├── Downloads/
│   ├── Desktop/
│   ├── .bashrc       ← Shell configuration (hidden)
│   ├── .ssh/         ← SSH keys (hidden)
│   └── .config/      ← App configs (hidden)
├── raj/
└── priya/
```

| Property | Detail |
|---|---|
| **Purpose** | Personal files and configurations for each user |
| **Shortcut** | `~` represents your home directory |
| **Full Path** | `/home/<username>` |
| **Permissions** | Each user can only access their own home (by default) |
| **Root's Home** | Root user's home is `/root` (not `/home/root`) |

### 11.2 `/etc` — Configuration Files

**`/etc`** = "Editable Text Configuration" — stores **system-wide configuration files**.

```
/etc/
├── hostname          ← System hostname
├── hosts             ← Static hostname-to-IP mapping
├── passwd            ← User account information
├── shadow            ← Encrypted passwords
├── group             ← Group definitions
├── fstab             ← Filesystem mount table
├── crontab           ← Scheduled tasks
├── sudoers           ← Sudo permissions
├── ssh/
│   └── sshd_config   ← SSH server configuration
├── nginx/
│   └── nginx.conf    ← Nginx web server config
├── apt/
│   └── sources.list  ← APT repository sources
└── systemd/
    └── system/       ← Systemd service files
```

> **Rule of Thumb**: If you need to configure a service, look in `/etc/`.

### 11.3 `/var` — Variable Data

**`/var`** stores files that **change frequently** during system operation.

```
/var/
├── log/              ← System and application logs
│   ├── syslog        ← General system log
│   ├── auth.log      ← Authentication log
│   ├── kern.log      ← Kernel log
│   └── nginx/
│       ├── access.log
│       └── error.log
├── cache/            ← Application cache
│   └── apt/          ← Downloaded .deb packages
├── lib/              ← State information
│   ├── docker/       ← Docker data
│   └── mysql/        ← MySQL databases
├── mail/             ← User mailboxes
├── spool/            ← Print and mail queues
└── tmp/              ← Temporary files (preserved across reboots)
```

### 11.4 `/usr` — User Programs

**`/usr`** = "Unix System Resources" — contains the majority of **user-installed programs and data**.

```
/usr/
├── bin/              ← User command binaries (ls, grep, git, python)
├── sbin/             ← System admin binaries
├── lib/              ← Libraries for /usr/bin and /usr/sbin
├── local/            ← Locally compiled software
│   ├── bin/
│   ├── lib/
│   └── etc/
├── share/            ← Architecture-independent data
│   ├── man/          ← Man pages
│   ├── doc/          ← Documentation
│   └── icons/        ← Icon files
└── include/          ← C/C++ header files
```

### 11.5 `/bin` — Essential Binaries

Contains **essential commands** needed for single-user mode and system recovery:
```
/bin/
├── ls        ← List directory contents
├── cp        ← Copy files
├── mv        ← Move/rename files
├── rm        ← Remove files
├── cat       ← Concatenate and display files
├── echo      ← Print text
├── bash      ← Bourne Again Shell
├── grep      ← Search text patterns
├── mkdir     ← Create directories
└── chmod     ← Change file permissions
```

> **Note**: In modern Linux (Ubuntu 20.04+), `/bin` is a **symlink** to `/usr/bin`.

### 11.6 `/tmp` — Temporary Files

| Property | Detail |
|---|---|
| **Purpose** | Temporary file storage |
| **Permissions** | World-writable (anyone can create files) |
| **Persistence** | Cleared on reboot (by default) |
| **Sticky Bit** | Set — users can only delete their own files |
| **Use Cases** | Build artifacts, temporary downloads, session files |

---

## 12. Paths in Linux
![absolute-relative-path-linux](https://github.com/user-attachments/assets/849f9e2e-c765-4927-9326-bb153d2c5019)


### 12.1 Absolute Path

An **absolute path** starts from the root `/` and specifies the **complete location** of a file.

```
/home/shivam/Documents/notes.txt
│                              │
└── Always starts with /       └── Full path from root to file
```

**Characteristics:**
- Always begins with `/`
- Unambiguous — works from **any** current directory
- Complete path from root to target

**Examples:**
```bash
/home/shivam/projects/app.py
/etc/nginx/nginx.conf
/var/log/syslog
/usr/bin/python3
```

### 12.2 Relative Path

A **relative path** is specified **relative to the current working directory**.

```bash
# If current directory is /home/shivam/
Documents/notes.txt          # relative to current directory
../raj/file.txt              # go up one level, then into raj/
./script.sh                  # current directory explicitly
```

**Comparison:**
| Feature | Absolute Path | Relative Path |
|---|---|---|
| Starts with | `/` | Directory name, `./`, or `../` |
| Depends on CWD | No | Yes |
| Length | Longer | Shorter |
| Use Case | Scripts, configs | Quick navigation |
| Example | `/home/shivam/file.txt` | `./file.txt` |

### 12.3 Special Path Symbols

| Symbol | Meaning | Example |
|---|---|---|
| `/` | Root directory | `cd /` |
| `~` | Home directory of current user | `cd ~` → `/home/shivam` |
| `.` | Current directory | `./run.sh` |
| `..` | Parent directory (one level up) | `cd ..` |
| `-` | Previous directory | `cd -` |
| `~username` | Home directory of specific user | `cd ~raj` → `/home/raj` |

---

## 13. Navigation Concepts

### 13.1 Moving Between Directories
![1_KbTO72sPS_K2f9HN1HSX0w](https://github.com/user-attachments/assets/18d8546c-2a8b-46d6-9d4f-1ec86c3a3222)


```bash
# Go to home directory
cd ~
cd              # Same as cd ~

# Go to root
cd /

# Go to specific directory (absolute)
cd /var/log

# Go to subdirectory (relative)
cd Documents

# Go up one level
cd ..

# Go up two levels
cd ../..

# Go to previous directory
cd -

# Print current directory
pwd
```

### 13.2 Listing Directory Contents

```bash
# Basic listing
ls

# Long format (permissions, owner, size, date)
ls -l

# Show hidden files
ls -a

# Combine: long format + hidden files
ls -la

# Human-readable file sizes
ls -lh

# Sort by modification time (newest first)
ls -lt

# Sort by size (largest first)
ls -lS

# Recursive listing (include subdirectories)
ls -R

# List specific directory
ls -la /etc
```

**Understanding `ls -l` Output:**
![d3e74a87f423bbb62e39d9de30e6399d](https://github.com/user-attachments/assets/b774427f-6477-4c31-8517-061035141eb2)

### 13.3 Hidden Files

Files and directories starting with `.` (dot) are **hidden** by default.

```bash
# Hidden files are NOT shown with regular ls
ls            # shows: Documents  Downloads  file.txt

# Show hidden files
ls -a         # shows: .  ..  .bashrc  .ssh  Documents  Downloads  file.txt
ls -A         # shows hidden but excludes . and ..
```

**Common Hidden Files:**
| File | Purpose |
|---|---|
| `.bashrc` | Bash shell configuration |
| `.bash_history` | Command history |
| `.profile` | Login shell configuration |
| `.ssh/` | SSH keys and configuration |
| `.gitconfig` | Git global configuration |
| `.config/` | Application configurations |
| `.vimrc` | Vim editor configuration |

### 13.4 Directory Tree Visualization

```bash
# Install tree command
sudo apt install tree

# Show directory tree
tree

# Limit depth
tree -L 2

# Show only directories
tree -d

# Show hidden files
tree -a

# Show with file sizes
tree -sh
```

**Example Output:**
![find-hidden-files-in-linux-using-tree-command-1](https://github.com/user-attachments/assets/84edd301-6b49-478d-a185-1a183a989a0a)

```
/home/shivam
├── Documents
│   ├── notes.txt
│   └── report.pdf
├── Downloads
│   └── installer.deb
├── projects
│   ├── app
│   │   ├── main.py
│   │   └── utils.py
│   └── README.md
├── .bashrc
└── .ssh
    ├── id_rsa
    └── id_rsa.pub
```

---

## 14. Understanding File Types
![file_type_codes_table](https://github.com/user-attachments/assets/f52d2c1a-905f-49a9-bbf9-de2b2a31d6e3)

### 14.1 Regular Files

Standard files containing data — text, binary, images, etc.

**Identified by:** `-` in `ls -l` output
```bash
-rw-r--r--  1  shivam  staff  4096  Feb 10  notes.txt    # regular file
```

**Subtypes:**
| Type | Description | Example |
|---|---|---|
| Text files | Human-readable content | `.txt`, `.conf`, `.sh`, `.py` |
| Binary files | Compiled programs | `/usr/bin/ls`, executables |
| Data files | Application data | `.db`, `.dat` |
| Media files | Images, audio, video | `.jpg`, `.mp3`, `.mp4` |

**Identify file type:**
```bash
file notes.txt          # Output: ASCII text
file /usr/bin/ls        # Output: ELF 64-bit LSB executable
file image.png          # Output: PNG image data, 1920 x 1080
```

### 14.2 Directories

Directories are special files that contain **references to other files and directories**.

**Identified by:** `d` in `ls -l` output
```bash
drwxr-xr-x  2  shivam  staff  4096  Feb 10  Documents    # directory
```

### 14.3 Hidden Files

Any file or directory with a name starting with `.` (dot):
```bash
.bashrc           # Hidden config file
.ssh/             # Hidden directory
.gitignore        # Hidden git file
.env              # Hidden environment file
```

**Why Hidden?**
- Keeps the directory listing clean
- Configuration files don't clutter the view
- Protects system files from accidental modification

### 14.4 Configuration Files

Config files control the behavior of system services and applications.

**Characteristics:**
- Usually plain text (can be edited with any text editor)
- System-wide configs in `/etc/`
- User-specific configs in `~/` (hidden dot files)
- Common formats: key-value, INI, YAML, JSON, XML

**Important Config Files:**
| File | Purpose | Location |
|---|---|---|
| `bashrc` | Shell settings, aliases | `~/.bashrc` |
| `sshd_config` | SSH server configuration | `/etc/ssh/sshd_config` |
| `nginx.conf` | Nginx web server config | `/etc/nginx/nginx.conf` |
| `fstab` | Filesystem mount table | `/etc/fstab` |
| `hosts` | Static DNS entries | `/etc/hosts` |
| `crontab` | Scheduled tasks | `/etc/crontab` |
| `resolv.conf` | DNS resolver config | `/etc/resolv.conf` |

### 14.5 All File Types in Linux

Linux recognizes **7 file types**:

| Symbol | Type | Description | Example |
|---|---|---|---|
| `-` | Regular file | Normal data files | `notes.txt` |
| `d` | Directory | Contains other files | `Documents/` |
| `l` | Symbolic link | Shortcut/pointer to another file | `python → python3` |
| `c` | Character device | Byte-stream devices | `/dev/tty` |
| `b` | Block device | Block-level devices | `/dev/sda` |
| `p` | Named pipe (FIFO) | Inter-process communication | Created with `mkfifo` |
| `s` | Socket | Network/IPC communication | `/var/run/docker.sock` |

---


# MODULE 3 — File Management (CRUD)

---

## 15. Create Operations

### 15.1 Creating Files

```bash
# Method 1: touch — Create an empty file (or update timestamp)
touch file.txt
touch file1.txt file2.txt file3.txt        # Multiple files

# Method 2: Redirect operator — Create with content
echo "Hello World" > file.txt              # Create/overwrite with content
echo "New line" >> file.txt                # Append to file

# Method 3: cat with heredoc — Multi-line content
cat > notes.txt << EOF
Line 1: Introduction
Line 2: Main content
Line 3: Conclusion
EOF

# Method 4: printf — Formatted content
printf "Name: %s\nAge: %d\n" "Shivam" 25 > profile.txt

# Method 5: Using an editor
nano newfile.txt                           # Opens Nano editor
vim newfile.txt                            # Opens Vim editor
```

**`touch` Deep Dive:**
| Usage | Command |
|---|---|
| Create empty file | `touch file.txt` |
| Create multiple files | `touch a.txt b.txt c.txt` |
| Update access time | `touch -a file.txt` |
| Update modification time | `touch -m file.txt` |
| Set specific timestamp | `touch -t 202601151200 file.txt` |

### 15.2 Creating Directories

```bash
# Create a single directory
mkdir projects

# Create multiple directories
mkdir dir1 dir2 dir3

# Create nested directories (with parents)
mkdir -p projects/backend/src/controllers
#        └── -p creates all intermediate directories

# Create with specific permissions
mkdir -m 755 secure-dir

# Verbose output (confirm creation)
mkdir -v new-folder
# Output: mkdir: created directory 'new-folder'
```

### 15.3 Nested Directory Creation

```bash
# Without -p flag (FAILS if parent doesn't exist)
mkdir projects/backend/src
# Error: mkdir: cannot create directory 'projects/backend/src': No such file or directory

# With -p flag (creates entire path)
mkdir -p projects/backend/src/controllers
mkdir -p projects/frontend/src/components

# Result:
# projects/
# ├── backend/
# │   └── src/
# │       └── controllers/
# └── frontend/
#     └── src/
#         └── components/
```

**Create a complete project structure in one go:**
```bash
mkdir -p myapp/{src/{controllers,models,views},config,tests,docs}

# Result:
# myapp/
# ├── config/
# ├── docs/
# ├── src/
# │   ├── controllers/
# │   ├── models/
# │   └── views/
# └── tests/
```

---

## 16. Read Operations

### 16.1 Viewing File Contents

```bash
# cat — Display entire file content
cat file.txt

# cat with line numbers
cat -n file.txt

# cat multiple files
cat file1.txt file2.txt

# Display non-printing characters
cat -A file.txt
```

### 16.2 Viewing Large Files

```bash
# less — Paginated viewer (most recommended for large files)
less /var/log/syslog
# Navigation: Space (next page), b (back), q (quit), /search

# more — Simple paginated viewer
more largefile.txt

# head — Show first N lines (default: 10)
head file.txt             # First 10 lines
head -n 20 file.txt       # First 20 lines
head -n 5 file.txt        # First 5 lines

# tail — Show last N lines (default: 10)
tail file.txt             # Last 10 lines
tail -n 20 file.txt       # Last 20 lines
tail -f /var/log/syslog   # Follow (live updates) — crucial for log monitoring

# Watch log files in real-time
tail -f /var/log/syslog    # Continuously shows new log entries (Ctrl+C to stop)
```

### 16.3 File Preview & Information

```bash
# wc — Word count (lines, words, characters)
wc file.txt               # Output: 10  50  350  file.txt
wc -l file.txt            # Line count only
wc -w file.txt            # Word count only
wc -c file.txt            # Byte count only

# file — Determine file type
file image.jpg             # Output: JPEG image data
file script.sh             # Output: Bourne-Again shell script

# stat — Detailed file information
stat file.txt
# Output: Size, blocks, permissions, timestamps, inode

# du — Disk usage of file/directory
du -sh file.txt            # Human-readable size
du -sh Documents/          # Total size of directory

# nl — Number lines (more control than cat -n)
nl file.txt
```

---

## 17. Update Operations

### 17.1 Editing Files

```bash
# Using Nano (beginner-friendly)
nano file.txt

# Using Vim (powerful, professional)
vim file.txt

# Using sed (stream editor — non-interactive)
sed -i 's/old-text/new-text/g' file.txt    # Find and replace in-place

# Using echo with redirect (overwrite entire file)
echo "New content" > file.txt
```

### 17.2 Appending Content

```bash
# Append single line
echo "New entry" >> file.txt

# Append multiple lines
cat >> file.txt << EOF
Line 1
Line 2
Line 3
EOF

# Append output of a command
date >> log.txt                    # Append current date
ls -la >> directory-listing.txt    # Append directory listing

# Append one file to another
cat extra.txt >> main.txt
```

**Redirect Operators:**
| Operator | Action |
|---|---|
| `>` | **Overwrite** — Creates file or replaces all content |
| `>>` | **Append** — Adds content to end of file |
| `2>` | Redirect **stderr** to file |
| `2>>` | Append **stderr** to file |
| `&>` | Redirect **both** stdout and stderr |
| `<` | Read input from file |

### 17.3 Renaming Files

```bash
# mv command is used for renaming (move to same directory with new name)
mv oldname.txt newname.txt

# Rename a directory
mv old-folder new-folder

# Rename with verbose output
mv -v report.txt final-report.txt
# Output: renamed 'report.txt' -> 'final-report.txt'
```

---

## 18. Delete Operations

### 18.1 Removing Files

```bash
# Remove a single file
rm file.txt

# Remove multiple files
rm file1.txt file2.txt file3.txt

# Remove with confirmation prompt
rm -i file.txt
# Output: rm: remove regular file 'file.txt'? y

# Force remove (no confirmation, no errors if missing)
rm -f file.txt

# Verbose remove
rm -v file.txt
# Output: removed 'file.txt'
```

### 18.2 Removing Directories

```bash
# Remove empty directory only
rmdir empty-folder

# Remove non-empty directory (FAILS)
rmdir full-folder
# Error: rmdir: failed to remove 'full-folder': Directory not empty
```

### 18.3 Recursive Deletion

```bash
# Remove directory and ALL its contents
rm -r folder-name

# Force recursive remove (no prompts) — ⚠ DANGEROUS
rm -rf folder-name

# Interactive recursive remove (asks for each file)
rm -ri folder-name

# Verbose recursive remove
rm -rv folder-name
```

> **⚠ CRITICAL WARNING**
> ```bash
> # NEVER run these commands:
> rm -rf /              # Deletes ENTIRE filesystem
> rm -rf /*             # Same effect
> rm -rf ~              # Deletes your entire home directory
> 
> # The --no-preserve-root flag exists as a safety mechanism
> # Modern Linux refuses to run rm -rf / without it
> ```

**Safe Deletion Practices:**
1. Always **double-check** the path before pressing Enter
2. Use `rm -ri` (interactive) when unsure which files will be deleted
3. Use `ls` first to preview what will be affected
4. Consider using `trash-cli` instead of `rm` for recoverable deletion

---

## 19. Copy & Move

### 19.1 Copying Files

```bash
# Copy a file
cp source.txt destination.txt

# Copy to a directory
cp file.txt /home/shivam/Documents/

# Copy multiple files to a directory
cp file1.txt file2.txt file3.txt /backup/

# Copy with verbose output
cp -v source.txt dest.txt
# Output: 'source.txt' -> 'dest.txt'

# Copy and preserve attributes (timestamps, permissions)
cp -p source.txt dest.txt

# Interactive (prompt before overwrite)
cp -i source.txt dest.txt

# Copy directory recursively
cp -r source-dir/ destination-dir/

# Copy directory preserving everything
cp -a source-dir/ destination-dir/       # archive mode (-r + preserve all)
```

### 19.2 Moving Files

```bash
# Move a file to another directory
mv file.txt /home/shivam/Documents/

# Move multiple files
mv file1.txt file2.txt /backup/

# Move a directory
mv project-dir/ /home/shivam/archives/

# Move with verbose output
mv -v file.txt /backup/
# Output: renamed 'file.txt' -> '/backup/file.txt'

# Interactive (prompt before overwrite)
mv -i file.txt /backup/

# Force move (overwrite without prompt)
mv -f file.txt /backup/

# Do not overwrite existing files
mv -n file.txt /backup/
```

### 19.3 Renaming Files (Using `mv`)

```bash
# Rename a file
mv old-name.txt new-name.txt

# Rename a directory
mv old-folder/ new-folder/

# Rename with backup of destination if it exists
mv --backup=numbered file.txt existing-file.txt
```

**`cp` vs `mv` Summary:**
| Feature | `cp` (Copy) | `mv` (Move/Rename) |
|---|---|---|
| Original file | **Kept** | **Removed** from source |
| Data duplication | Yes (uses disk space) | No (just updates path) |
| Speed (same disk) | Slower (copies data) | Instant (updates metadata) |
| Speed (cross disk) | Same | Same |
| Recursive flag | `-r` required for dirs | Not needed |

---

## 20. Wildcards & Patterns

### 20.1 Globbing (File Matching)

Wildcards allow you to match multiple files using **patterns** instead of typing each filename.

| Wildcard | Meaning | Example | Matches |
|---|---|---|---|
| `*` | Any number of any characters | `*.txt` | `file.txt`, `notes.txt`, `a.txt` |
| `?` | Exactly one character | `file?.txt` | `file1.txt`, `fileA.txt` (not `file10.txt`) |
| `[...]` | One character from the set | `file[123].txt` | `file1.txt`, `file2.txt`, `file3.txt` |
| `[!...]` or `[^...]` | One character NOT in the set | `file[!0-9].txt` | `fileA.txt`, `fileB.txt` |
| `{a,b,c}` | Brace expansion | `file.{txt,md,py}` | `file.txt`, `file.md`, `file.py` |

### 20.2 Practical Examples

```bash
# List all .txt files
ls *.txt

# List all files starting with "log"
ls log*

# List files with single character extension
ls *.?

# List files with 3-character names
ls ???.*

# List all .jpg and .png files
ls *.{jpg,png}

# Copy all Python files to backup
cp *.py /backup/

# Remove all .tmp files
rm *.tmp

# Move all log files from 2025
mv *2025*.log /archive/

# List files starting with a, b, or c
ls [abc]*

# List files NOT starting with a vowel
ls [!aeiou]*
```

### 20.3 Bulk Operations

```bash
# Rename all .txt to .md (requires rename utility)
# Install: sudo apt install rename
rename 's/\.txt$/\.md/' *.txt

# Create 10 numbered files
touch file{1..10}.txt
# Creates: file1.txt, file2.txt, ... file10.txt

# Create files with padding
touch file{01..10}.txt
# Creates: file01.txt, file02.txt, ... file10.txt

# Create files with step
touch file{0..20..5}.txt
# Creates: file0.txt, file5.txt, file10.txt, file15.txt, file20.txt

# Batch copy to multiple destinations
for dir in /backup1 /backup2 /backup3; do
    cp important.txt "$dir/"
done

# Find and delete all .log files older than 7 days
find /var/log -name "*.log" -mtime +7 -delete
```

---

## 21. Editors — Nano & Vim

### 21.1 Nano

**Nano** is a simple, beginner-friendly terminal text editor.

**Opening Files:**
```bash
nano file.txt              # Open existing file or create new
nano +10 file.txt          # Open at line 10
sudo nano /etc/hosts       # Edit system files (requires sudo)
```

**Nano Interface:**
```
  GNU nano 6.2                     file.txt

This is the content of the file.
You can type and edit normally.
The cursor moves with arrow keys.




^G Help    ^O Write Out  ^W Where Is  ^K Cut     ^U Paste
^X Exit    ^R Read File  ^\ Replace   ^J Justify ^T Spell
```

**Essential Nano Shortcuts:**
| Shortcut | Action |
|---|---|
| `Ctrl + O` | **Save** (Write Out) — then press Enter |
| `Ctrl + X` | **Exit** (prompts to save if modified) |
| `Ctrl + K` | **Cut** current line |
| `Ctrl + U` | **Paste** cut line |
| `Ctrl + W` | **Search** for text |
| `Ctrl + \` | **Find and Replace** |
| `Ctrl + G` | **Help** menu |
| `Ctrl + _` | **Go to** specific line number |
| `Ctrl + C` | Show **cursor position** (line, column) |
| `Alt + U` | **Undo** |
| `Alt + E` | **Redo** |
| `Ctrl + Shift + 6` | **Select** text (then use Ctrl+K to cut) |

### 21.2 Vim

**Vim** (Vi Improved) is the most powerful terminal text editor, favored by professionals and DevOps engineers.

**Opening Files:**
```bash
vim file.txt               # Open file
vim +20 file.txt           # Open at line 20
vim +/pattern file.txt     # Open at first occurrence of pattern
vim -O file1.txt file2.txt # Open side by side (vertical split)
```

**Vim's Modal Design:**
Vim operates in **modes** — this is its core concept.

```
┌──────────────────────────────────────────────────────────┐
│                                                          │
│   ┌──────────────────┐         ┌──────────────────┐     │
│   │   NORMAL MODE    │◄────────│  COMMAND MODE    │     │
│   │  (Navigation,    │  Esc    │  (Save, Quit,    │     │
│   │   Operations)    │────────►│   Search)        │     │
│   └────────┬─────────┘   :    └──────────────────┘     │
│            │                                            │
│        i,a,o (Enter Insert)                             │
│            │                                            │
│   ┌────────▼─────────┐         ┌──────────────────┐    │
│   │   INSERT MODE    │         │  VISUAL MODE     │    │
│   │  (Type/Edit      │         │  (Select text)   │    │
│   │   text)          │         │                  │    │
│   └────────┬─────────┘         └──────────────────┘    │
│            │                          ▲                  │
│         Esc (Back to Normal)          │                  │
│                              v, V, Ctrl+V               │
│                         (from Normal mode)               │
└──────────────────────────────────────────────────────────┘
```

**Vim Modes Explained:**
| Mode | Purpose | Enter With | Exit With |
|---|---|---|---|
| **Normal** | Navigation, delete, copy, paste | `Esc` | — |
| **Insert** | Type and edit text | `i`, `a`, `o` | `Esc` |
| **Visual** | Select text | `v`, `V`, `Ctrl+V` | `Esc` |
| **Command** | Save, quit, search, replace | `:` | `Enter` or `Esc` |

**Essential Vim Commands:**

**Entering Insert Mode:**
| Key | Action |
|---|---|
| `i` | Insert before cursor |
| `I` | Insert at beginning of line |
| `a` | Append after cursor |
| `A` | Append at end of line |
| `o` | Open new line below |
| `O` | Open new line above |

**Navigation (Normal Mode):**
| Key | Action |
|---|---|
| `h` | Move left |
| `j` | Move down |
| `k` | Move up |
| `l` | Move right |
| `w` | Jump to next word |
| `b` | Jump to previous word |
| `0` | Jump to beginning of line |
| `$` | Jump to end of line |
| `gg` | Go to first line of file |
| `G` | Go to last line of file |
| `10G` or `:10` | Go to line 10 |
| `Ctrl+f` | Page forward |
| `Ctrl+b` | Page backward |

**Editing (Normal Mode):**
| Key | Action |
|---|---|
| `dd` | Delete (cut) entire line |
| `5dd` | Delete 5 lines |
| `dw` | Delete word |
| `d$` or `D` | Delete to end of line |
| `yy` | Yank (copy) entire line |
| `5yy` | Yank 5 lines |
| `p` | Paste after cursor |
| `P` | Paste before cursor |
| `u` | Undo |
| `Ctrl+r` | Redo |
| `x` | Delete character under cursor |
| `r` | Replace single character |
| `.` | Repeat last command |

**Saving & Quitting (Command Mode — press `:` first):**
| Command | Action |
|---|---|
| `:w` | Save (write) |
| `:q` | Quit |
| `:wq` or `:x` | Save and quit |
| `:q!` | Quit without saving (force) |
| `:wq!` | Save and quit (force) |
| `:w filename` | Save as new filename |

**Search & Replace (Command Mode):**
| Command | Action |
|---|---|
| `/pattern` | Search forward |
| `?pattern` | Search backward |
| `n` | Next match |
| `N` | Previous match |
| `:%s/old/new/g` | Replace all occurrences in file |
| `:%s/old/new/gc` | Replace with confirmation |
| `:5,10s/old/new/g` | Replace in lines 5-10 |

### 21.3 Nano vs Vim Comparison

| Feature | Nano | Vim |
|---|---|---|
| **Learning Curve** | Low (minutes) | High (days to weeks) |
| **Efficiency** | Moderate | Very high (once learned) |
| **Customization** | Limited | Extremely customizable |
| **Plugins** | None | Extensive ecosystem |
| **Availability** | Most distros | Virtually every Unix system |
| **Best For** | Quick edits, beginners | Power users, daily editing |
| **Config Editing** | Great for simple edits | Professional-grade editing |

### 21.4 Practical Vim Workflow

```
1. Open file:            vim config.yaml
2. Navigate to line 25:  25G
3. Enter insert mode:    i
4. Make changes:         (type your changes)
5. Exit insert mode:     Esc
6. Search for "port":    /port
7. Save and quit:        :wq
```

---

## Quick Reference — Command Cheat Sheet

### File System Navigation
| Command | Description |
|---|---|
| `pwd` | Print current directory |
| `cd <dir>` | Change directory |
| `cd ~` | Go to home |
| `cd ..` | Go up one level |
| `cd -` | Go to previous directory |
| `ls -la` | List all files (detailed) |
| `tree -L 2` | Directory tree (2 levels) |

### File Operations
| Command | Description |
|---|---|
| `touch file` | Create empty file |
| `mkdir -p a/b/c` | Create nested directories |
| `cat file` | View file content |
| `less file` | Paginated viewer |
| `head -n 20 file` | First 20 lines |
| `tail -f file` | Follow file (live) |
| `cp -r src/ dest/` | Copy directory |
| `mv old new` | Move / rename |
| `rm -r dir/` | Delete directory |
| `rm -i file` | Delete with confirmation |

### Package Management
| Command | Description |
|---|---|
| `sudo apt update` | Refresh package index |
| `sudo apt upgrade` | Upgrade packages |
| `sudo apt install pkg` | Install package |
| `sudo apt remove pkg` | Remove package |
| `apt search keyword` | Search packages |

### System Information
| Command | Description |
|---|---|
| `uname -a` | System info |
| `uname -r` | Kernel version |
| `whoami` | Current user |
| `hostname` | Machine name |
| `uptime` | System uptime |
| `df -h` | Disk usage |
| `free -h` | Memory usage |

---


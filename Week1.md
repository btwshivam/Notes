# Week 1 — Open Source Foundations + SDLC + Git/GitHub
## Complete Guide (Basic to Advanced)

---

## Table of Contents
1. [Open Source Software (OSS) Fundamentals](#1-open-source-software-oss-fundamentals)
2. [LFX Mentorship & Selection Process](#2-lfx-mentorship--selection-process)
3. [Software Development Life Cycle (SDLC)](#3-software-development-life-cycle-sdlc)
4. [Agile & Scrum Methodology](#4-agile--scrum-methodology)
5. [Git Fundamentals](#5-git-fundamentals)
6. [GitHub Essentials](#6-github-essentials)
7. [Branching Strategies](#7-branching-strategies)
8. [Professional Commit Messages](#8-professional-commit-messages)
9. [Writing Effective Pull Requests](#9-writing-effective-pull-requests)
10. [Advanced Git Topics](#10-advanced-git-topics)
11. [Best Practices & Real-world Workflows](#11-best-practices--real-world-workflows)

---

## 1. Open Source Software (OSS) Fundamentals

### 1.1 What is Open Source?

**Definition**: Open Source Software is software with source code that anyone can inspect, modify, and enhance.

**Key Principles:**
- **Freedom to use** - Run the program for any purpose
- **Freedom to study** - Access and understand the source code
- **Freedom to modify** - Make changes and improvements
- **Freedom to distribute** - Share copies with others

### 1.2 Why Open Source Matters

**Benefits:**
1. **Transparency** - Code is visible and auditable
2. **Community-driven** - Collective intelligence and collaboration
3. **Quality** - Peer review leads to better code
4. **Innovation** - Rapid development and experimentation
5. **Career Growth** - Build portfolio and reputation
6. **Learning** - Learn from experienced developers

**Real-world Examples:**
- Linux (Operating System)
- Apache (Web Server)
- Firefox (Browser)
- Kubernetes (Container Orchestration)
- TensorFlow (Machine Learning)
- VS Code (Code Editor)

### 1.3 Open Source Licenses

**Permissive Licenses:**
- **MIT License** - Simple and permissive, allows commercial use
- **Apache 2.0** - Includes patent protection
- **BSD** - Very permissive, minimal restrictions

**Copyleft Licenses:**
- **GPL (GNU General Public License)** - Derivatives must be open source
- **LGPL** - Less restrictive version of GPL
- **AGPL** - Covers network use cases

**Creative Commons:**
- Used for documentation and non-code assets
- Various combinations (BY, SA, NC, ND)

### 1.4 OSS Culture & Etiquette

**Core Values:**
1. **Respect** - Treat everyone with dignity
2. **Collaboration** - Work together, not against
3. **Meritocracy** - Good ideas and code win
4. **Transparency** - Open communication and decision-making
5. **Inclusivity** - Welcome all contributors

**Code of Conduct:**
- Be welcoming and friendly
- Be patient with newcomers
- Accept constructive criticism
- Focus on what's best for the community
- Show empathy towards others

**Communication Best Practices:**
- Be clear and concise
- Provide context in issues/PRs
- Search before asking
- Be patient for responses
- Thank maintainers for their time

### 1.5 Types of Contributions

**Code Contributions:**
- Bug fixes
- New features
- Performance improvements
- Refactoring

**Non-code Contributions:**
- Documentation improvements
- Translation
- Design work
- Testing and bug reporting
- Community support
- Event organization

### 1.6 Finding Projects to Contribute

**Platforms:**
- GitHub Explore
- GitLab
- Good First Issue
- Up For Grabs
- CodeTriage
- First Timers Only

**Selection Criteria:**
- Active maintenance
- Clear documentation
- Welcoming community
- Aligned with your interests
- Appropriate skill level

---

## 2. LFX Mentorship & Selection Process

### 2.1 What is LFX (Linux Foundation X)?

**Overview:**
LFX Mentorship is a platform by The Linux Foundation that connects mentors with mentees to work on open source projects.

**Program Types:**
1. **Full-time (12 weeks)** - $3,000-6,000 stipend
2. **Part-time (24 weeks)** - Similar total compensation
3. **Seasonal** - Spring, Summer, Fall terms

**Participating Projects:**
- Kubernetes
- Cloud Native Computing Foundation (CNCF)
- Hyperledger
- Linux Kernel
- Jenkins
- And 100+ other projects

### 2.2 Application Process

**Step 1: Preparation (2-3 months before)**
- Build foundational skills
- Explore project repositories
- Make small contributions
- Understand project architecture

**Step 2: Profile Creation**
- Create LFX account
- Complete profile thoroughly
- Highlight relevant skills
- Upload resume/CV

**Step 3: Project Selection**
- Browse available projects
- Read project descriptions carefully
- Check required skills
- Review mentor information

**Step 4: Application Submission**
- Write compelling cover letter
- Showcase relevant experience
- Explain motivation clearly
- Propose project approach

**Step 5: Interview Process**
- Technical assessment may be required
- Video interview with mentors
- Code review or task completion
- Community interaction

### 2.3 Selection Criteria

**Technical Skills:**
- Programming proficiency
- Understanding of project technology
- Problem-solving ability
- Code quality

**Community Engagement:**
- Prior contributions (huge plus)
- Active in project channels
- Helpful and collaborative
- Understanding of project goals

**Communication:**
- Clear written communication
- Responsiveness
- Professionalism
- English proficiency

**Commitment:**
- Time availability
- Long-term interest
- Realistic goals
- Dedication to learning

### 2.4 Tips for Success

**Before Applying:**
1. Start contributing 2-3 months early
2. Join project communication channels (Slack, Discord, Mailing lists)
3. Read all documentation thoroughly
4. Build small proof-of-concept projects
5. Engage with the community

**During Application:**
1. Be specific about your interests
2. Show, don't just tell
3. Tailor each application
4. Proofread everything
5. Follow instructions exactly

**After Selection:**
1. Set clear milestones
2. Communicate regularly
3. Ask questions early
4. Document your work
5. Continue contributing after program ends

---

## 3. Software Development Life Cycle (SDLC)

### 3.1 What is SDLC?

**Definition:**
A structured process for planning, creating, testing, and deploying software systems.

**Why SDLC Matters:**
- Ensures quality and consistency
- Reduces development costs
- Improves project management
- Minimizes risks
- Facilitates maintenance

### 3.2 SDLC Phases

**1. Planning & Requirement Analysis**
- Identify stakeholders
- Gather requirements
- Feasibility study
- Resource allocation
- Risk assessment

**Activities:**
- User interviews
- Market research
- Technical feasibility
- Cost-benefit analysis

**2. System Design**
- Architecture design
- Database design
- UI/UX design
- API design
- Security design

**Deliverables:**
- System architecture documents
- Database schemas
- Wireframes/mockups
- Technical specifications

**3. Implementation (Development)**
- Writing code
- Following coding standards
- Version control
- Unit testing
- Code reviews

**Best Practices:**
- Modular development
- Clean code principles
- Documentation
- Regular commits

**4. Testing**
- Unit testing
- Integration testing
- System testing
- Acceptance testing
- Performance testing

**Types of Testing:**
- Functional testing
- Non-functional testing
- Regression testing
- Security testing

**5. Deployment**
- Release planning
- Environment setup
- Deployment execution
- Monitoring setup
- User training

**Deployment Strategies:**
- Blue-green deployment
- Rolling deployment
- Canary deployment
- Feature flags

**6. Maintenance**
- Bug fixes
- Updates and patches
- Performance optimization
- Feature enhancements
- Technical support

**Types of Maintenance:**
- Corrective (fixing bugs)
- Adaptive (adapting to changes)
- Perfective (improvements)
- Preventive (future-proofing)

### 3.3 SDLC Models

**1. Waterfall Model**
```
Requirements → Design → Implementation → Testing → Deployment → Maintenance
```
**Pros:** Simple, structured, well-documented
**Cons:** Inflexible, late testing, high risk

**2. Agile Model**
```
Iterative cycles: Plan → Design → Develop → Test → Review → Deploy
```
**Pros:** Flexible, early delivery, customer collaboration
**Cons:** Less predictable, requires experience

**3. Iterative Model**
- Build in increments
- Each iteration is a mini-project
- Refine with each cycle

**4. Spiral Model**
- Risk-driven approach
- Combines waterfall and iterative
- Four phases: Planning, Risk Analysis, Engineering, Evaluation

**5. DevOps Model**
- Integration of Development and Operations
- Continuous Integration/Continuous Deployment (CI/CD)
- Automation and monitoring

### 3.4 Modern Development Practices

**Continuous Integration (CI):**
- Frequent code integration
- Automated builds
- Automated testing
- Quick feedback

**Continuous Deployment (CD):**
- Automated deployment
- Release automation
- Environment management
- Rollback capabilities

**Infrastructure as Code (IaC):**
- Version-controlled infrastructure
- Repeatable deployments
- Cloud resource management

---

## 4. Agile & Scrum Methodology

### 4.1 Agile Manifesto

**Four Core Values:**
1. **Individuals and interactions** over processes and tools
2. **Working software** over comprehensive documentation
3. **Customer collaboration** over contract negotiation
4. **Responding to change** over following a plan

**12 Principles:**
1. Satisfy customer through early and continuous delivery
2. Welcome changing requirements
3. Deliver working software frequently
4. Business and developers work together daily
5. Build projects around motivated individuals
6. Face-to-face conversation is best
7. Working software is primary measure of progress
8. Sustainable development pace
9. Technical excellence and good design
10. Simplicity is essential
11. Self-organizing teams
12. Regular reflection and adjustment

### 4.2 Scrum Framework

**What is Scrum?**
An Agile framework for managing complex projects through iterative progress.

**Scrum Roles:**

**1. Product Owner**
- Defines product vision
- Manages product backlog
- Prioritizes features
- Accepts/rejects work
- Stakeholder representative

**2. Scrum Master**
- Facilitates Scrum process
- Removes impediments
- Coaches the team
- Protects team from interruptions
- Ensures Scrum practices

**3. Development Team**
- Cross-functional (3-9 members)
- Self-organizing
- Delivers product increments
- Estimates work
- Commits to sprint goals

### 4.3 Scrum Artifacts

**1. Product Backlog**
- Ordered list of features
- User stories format
- Refined continuously
- Single source of requirements

**2. Sprint Backlog**
- Selected items for sprint
- Tasks breakdown
- Team commits to completion
- Visible to everyone

**3. Increment**
- Potentially shippable product
- Meets Definition of Done
- Integrated and tested
- Demonstrated at sprint review

### 4.4 Scrum Events

**1. Sprint**
- Time-boxed iteration (1-4 weeks)
- Consistent duration
- Produces increment
- No changes that endanger sprint goal

**2. Sprint Planning**
- Duration: 2-4 hours
- Define sprint goal
- Select backlog items
- Create sprint backlog
- Answer: "What" and "How"

**3. Daily Scrum (Stand-up)**
- Duration: 15 minutes
- Same time and place
- Three questions:
  - What did I do yesterday?
  - What will I do today?
  - Any impediments?

**4. Sprint Review**
- Duration: 1-2 hours
- Demo completed work
- Gather feedback
- Adapt product backlog
- Stakeholder attendance

**5. Sprint Retrospective**
- Duration: 1 hour
- Inspect team process
- Identify improvements
- Create action plan
- Occurs after sprint review

### 4.5 User Stories

**Format:**
```
As a [role]
I want [feature]
So that [benefit]
```

**Example:**
```
As a user
I want to reset my password
So that I can regain access to my account if I forget it
```

**INVEST Criteria:**
- **I**ndependent - Can be developed separately
- **N**egotiable - Details can be discussed
- **V**aluable - Delivers value to users
- **E**stimable - Can be estimated
- **S**mall - Completable in one sprint
- **T**estable - Clear acceptance criteria

**Acceptance Criteria:**
```
Given [context]
When [action]
Then [outcome]
```

### 4.6 Estimation Techniques

**Story Points:**
- Relative sizing
- Fibonacci sequence (1, 2, 3, 5, 8, 13)
- Considers complexity, effort, uncertainty

**Planning Poker:**
- Team-based estimation
- Each member votes
- Discuss differences
- Reach consensus

**T-Shirt Sizing:**
- XS, S, M, L, XL
- Quick, rough estimates
- Good for early planning

---

## 5. Git Fundamentals

### 5.1 What is Git?

**Definition:**
Git is a distributed version control system for tracking changes in source code during software development.

**Key Concepts:**
- **Version Control** - Track changes over time
- **Distributed** - Every developer has full history
- **Branching** - Work on features independently
- **Merging** - Combine different lines of development

**Why Git?**
- Collaboration made easy
- Complete history and audit trail
- Branching and merging capabilities
- Speed and performance
- Data integrity
- Industry standard

### 5.2 Git Installation

**Linux:**
```bash
# Debian/Ubuntu
sudo apt-get update
sudo apt-get install git

# Fedora/RHEL
sudo dnf install git
```

**macOS:**
```bash
# Using Homebrew
brew install git

# Or download from git-scm.com
```

**Windows:**
- Download from git-scm.com
- Git Bash provides Unix-like command line

**Verify Installation:**
```bash
git --version
```

### 5.3 Git Configuration

**Set User Information:**
```bash
# Set username
git config --global user.name "Your Name"

# Set email
git config --global user.email "your.email@example.com"

# Check configuration
git config --list

# View specific config
git config user.name
```

**Configuration Levels:**
- `--system` - All users on the system
- `--global` - All repositories for current user
- `--local` - Current repository only (default)

**Other Useful Configurations:**
```bash
# Set default branch name
git config --global init.defaultBranch main

# Set default editor
git config --global core.editor "code --wait"  # VS Code
git config --global core.editor "vim"          # Vim

# Enable color output
git config --global color.ui auto

# Set line ending handling
git config --global core.autocrlf input  # Linux/Mac
git config --global core.autocrlf true   # Windows
```

### 5.4 Git Repository Basics

**Creating a Repository:**

**Method 1: Initialize New Repository**
```bash
# Navigate to project directory
cd my-project

# Initialize Git repository
git init

# Result: Creates .git directory
```

**Method 2: Clone Existing Repository**
```bash
# Clone remote repository
git clone https://github.com/username/repository.git

# Clone with different name
git clone https://github.com/username/repository.git my-folder

# Clone specific branch
git clone -b branch-name https://github.com/username/repository.git
```

**Repository Structure:**
```
.git/
├── HEAD              # Points to current branch
├── config            # Repository configuration
├── description       # Repository description
├── hooks/            # Git hooks (scripts)
├── objects/          # Git objects (commits, trees, blobs)
├── refs/             # References (branches, tags)
│   ├── heads/        # Local branches
│   ├── remotes/      # Remote branches
│   └── tags/         # Tags
└── index             # Staging area
```

### 5.5 Git Workflow: Add, Commit, Push

**Git Areas:**
1. **Working Directory** - Your local files
2. **Staging Area (Index)** - Files ready to commit
3. **Repository (.git)** - Committed snapshots

**Basic Workflow:**

```bash
# 1. Check status
git status

# 2. Add files to staging area
git add filename.txt           # Add specific file
git add folder/                # Add folder
git add .                      # Add all changes
git add *.js                   # Add all JavaScript files
git add -A                     # Add all (new, modified, deleted)
git add -u                     # Add modified and deleted only

# 3. Review what's staged
git status
git diff                       # See unstaged changes
git diff --staged              # See staged changes

# 4. Commit changes
git commit -m "Add feature X"

# 5. Push to remote
git push origin main
```

**Understanding git add:**
```bash
# Interactive staging
git add -p                     # Patch mode: stage chunks interactively

# Add and commit together (for tracked files only)
git commit -am "Message"
```

**Undoing Add:**
```bash
# Unstage file
git restore --staged filename.txt
git reset HEAD filename.txt    # Older syntax

# Discard changes in working directory
git restore filename.txt
git checkout -- filename.txt   # Older syntax
```

### 5.6 Git Commit Deep Dive

**What is a Commit?**
A commit is a snapshot of your repository at a specific point in time.

**Commit Anatomy:**
- **SHA-1 hash** - Unique identifier (40 characters)
- **Author** - Who made the commit
- **Date** - When it was made
- **Message** - Description of changes
- **Parent(s)** - Previous commit(s)
- **Tree** - Snapshot of files

**Creating Commits:**
```bash
# Basic commit
git commit -m "Commit message"

# Multi-line commit message
git commit -m "Short summary" -m "Detailed explanation"

# Open editor for detailed message
git commit

# Amend last commit (add forgotten changes or fix message)
git commit --amend

# Commit with specific author
git commit --author="Name <email>" -m "Message"
```

**Viewing Commits:**
```bash
# Show commit history
git log

# One line per commit
git log --oneline

# Show graph
git log --graph --oneline --all

# Show last N commits
git log -5

# Show commits by author
git log --author="Name"

# Show commits in date range
git log --since="2024-01-01" --until="2024-12-31"

# Show commits affecting specific file
git log -- filename.txt

# Show commit details
git show commit-hash
git show HEAD               # Latest commit
git show HEAD~1             # Previous commit
```

**Pretty Log Formats:**
```bash
# Custom format
git log --pretty=format:"%h - %an, %ar : %s"

# Common alias
git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit
```

### 5.7 Git Diff

**Understanding Differences:**
```bash
# Show unstaged changes
git diff

# Show staged changes
git diff --staged
git diff --cached

# Show changes between commits
git diff commit1 commit2

# Show changes in specific file
git diff filename.txt

# Show changes between branches
git diff branch1 branch2

# Show word-level differences
git diff --word-diff

# Show statistics
git diff --stat
```

### 5.8 Git History & Time Travel

**Reverting Changes:**
```bash
# Discard changes in working directory
git restore filename.txt

# Unstage files
git restore --staged filename.txt

# Reset to previous commit (keep changes)
git reset --soft HEAD~1

# Reset to previous commit (discard changes)
git reset --hard HEAD~1

# Revert a commit (creates new commit)
git revert commit-hash
```

**HEAD Notation:**
```bash
HEAD        # Current commit
HEAD~1      # Parent commit (1 commit back)
HEAD~2      # Grandparent commit (2 commits back)
HEAD^       # Parent commit (alternative)
HEAD^^      # Grandparent commit
```

**Checkout Specific Commit:**
```bash
# Detached HEAD state (view-only)
git checkout commit-hash

# Return to branch
git checkout main
```

### 5.9 .gitignore

**Purpose:**
Specify files/directories that Git should ignore.

**Common Patterns:**
```gitignore
# Ignore specific file
secrets.txt

# Ignore all files with extension
*.log
*.tmp

# Ignore entire directory
node_modules/
build/
dist/

# Ignore files in any directory
**/temp

# Negation (don't ignore this specific file)
!important.log

# Ignore files only in current directory
/config.json

# Ignore based on pattern
**/logs/*.log
```

**Common .gitignore for Different Languages:**

**Python:**
```gitignore
__pycache__/
*.py[cod]
*.so
.env
venv/
*.egg-info/
dist/
build/
```

**Node.js:**
```gitignore
node_modules/
npm-debug.log
.env
dist/
build/
*.log
```

**Java:**
```gitignore
*.class
*.jar
*.war
target/
.idea/
*.iml
```

**Creating .gitignore:**
```bash
# Create .gitignore file
touch .gitignore

# Edit with your preferred editor
nano .gitignore

# Add and commit
git add .gitignore
git commit -m "Add .gitignore"
```

**Useful Resources:**
- [gitignore.io](https://gitignore.io) - Generate .gitignore templates
- GitHub's [gitignore templates](https://github.com/github/gitignore)

### 5.10 Git Remote

**Understanding Remotes:**
Remote repositories are versions of your project hosted on the internet or network.

**Common Commands:**
```bash
# List remotes
git remote
git remote -v                  # Verbose (show URLs)

# Add remote
git remote add origin https://github.com/username/repo.git

# Remove remote
git remote remove origin

# Rename remote
git remote rename old-name new-name

# Change remote URL
git remote set-url origin https://new-url.git

# Show remote details
git remote show origin
```

**Fetch vs Pull:**
```bash
# Fetch: Download objects and refs from remote (doesn't merge)
git fetch origin

# Pull: Fetch + Merge
git pull origin main

# Pull with rebase
git pull --rebase origin main
```

**Push:**
```bash
# Push to remote
git push origin main

# Push all branches
git push --all origin

# Push tags
git push --tags origin

# Force push (dangerous!)
git push --force origin main

# Safer force push
git push --force-with-lease origin main

# Set upstream branch
git push -u origin main        # After this, just 'git push' works
```

---

## 6. GitHub Essentials

### 6.1 What is GitHub?

**Definition:**
GitHub is a web-based hosting service for Git repositories with additional collaboration features.

**GitHub vs Git:**
- **Git** - Version control software (local)
- **GitHub** - Cloud hosting service for Git repositories

**GitHub Features:**
- Repository hosting
- Issue tracking
- Pull requests
- Project management
- GitHub Actions (CI/CD)
- GitHub Pages (hosting)
- Security features
- Social coding (stars, follows)

**Alternatives:**
- GitLab
- Bitbucket
- Gitea
- SourceForge

### 6.2 Creating GitHub Account

**Steps:**
1. Go to [github.com](https://github.com)
2. Click "Sign up"
3. Enter email, password, username
4. Verify email
5. Complete profile

**Profile Best Practices:**
- Use professional username
- Add profile picture
- Write bio
- Add location
- Link personal website
- Pin best repositories
- Keep activity consistent

### 6.3 SSH Keys Setup

**Why SSH?**
- More secure than HTTPS
- No password prompts
- Required for some operations

**Generate SSH Key:**
```bash
# Generate new SSH key
ssh-keygen -t ed25519 -C "your.email@example.com"

# For older systems
ssh-keygen -t rsa -b 4096 -C "your.email@example.com"

# Press Enter to accept default location
# Enter passphrase (optional but recommended)
```

**Add SSH Key to SSH Agent:**
```bash
# Start SSH agent
eval "$(ssh-agent -s)"

# Add key to agent
ssh-add ~/.ssh/id_ed25519
```

**Add SSH Key to GitHub:**
```bash
# Copy public key to clipboard
# Linux
cat ~/.ssh/id_ed25519.pub | xclip -selection clipboard

# macOS
cat ~/.ssh/id_ed25519.pub | pbcopy

# Windows Git Bash
cat ~/.ssh/id_ed25519.pub | clip
```

Then:
1. Go to GitHub Settings → SSH and GPG keys
2. Click "New SSH key"
3. Paste key and add title
4. Click "Add SSH key"

**Test Connection:**
```bash
ssh -T git@github.com
```

### 6.4 Creating Repositories on GitHub

**Method 1: Web Interface**
1. Click "+" icon → New repository
2. Enter repository name
3. Add description (optional)
4. Choose public/private
5. Initialize with README (optional)
6. Add .gitignore (optional)
7. Choose license (optional)
8. Click "Create repository"

**Method 2: GitHub CLI**
```bash
# Install GitHub CLI (gh)
# Then create repo
gh repo create my-repo --public
```

**Connecting Local to Remote:**
```bash
# If starting with local repository
cd my-project
git init
git add .
git commit -m "Initial commit"
git remote add origin git@github.com:username/repo.git
git branch -M main
git push -u origin main
```

### 6.5 Repository Structure

**Essential Files:**

**README.md:**
```markdown
# Project Title

Brief description of what this project does.

## Features
- Feature 1
- Feature 2

## Installation
```bash
npm install
```

## Usage
```bash
npm start
```

## Contributing
Pull requests are welcome!

## License
MIT
```

**LICENSE:**
- Choose appropriate license
- GitHub provides templates
- MIT, Apache 2.0, GPL are common

**CONTRIBUTING.md:**
```markdown
# Contributing Guidelines

## How to Contribute
1. Fork the repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Open pull request

## Code Style
- Use 2 spaces for indentation
- Follow ESLint rules
- Write meaningful commit messages

## Reporting Bugs
Use GitHub Issues with bug template
```

**CODE_OF_CONDUCT.md:**
- Defines community standards
- Use Contributor Covenant template

### 6.6 GitHub Issues

**What are Issues?**
Track bugs, enhancements, tasks, and discussions.

**Creating Good Issues:**
```markdown
**Description:**
Clear description of the issue

**Steps to Reproduce:**
1. Go to '...'
2. Click on '....'
3. See error

**Expected Behavior:**
What should happen

**Actual Behavior:**
What actually happens

**Environment:**
- OS: Ubuntu 22.04
- Browser: Chrome 120
- Version: 1.2.3

**Screenshots:**
[Attach if relevant]
```

**Issue Labels:**
- `bug` - Something isn't working
- `enhancement` - New feature request
- `documentation` - Documentation improvements
- `good first issue` - Good for newcomers
- `help wanted` - Need community help
- `duplicate` - Already reported
- `wontfix` - Won't be addressed

**Issue Templates:**
Create `.github/ISSUE_TEMPLATE/bug_report.md`

### 6.7 Forking & Cloning

**Forking:**
Creates personal copy of someone else's repository.

**Steps:**
1. Navigate to repository
2. Click "Fork" button
3. Fork created in your account

**Cloning:**
```bash
# Clone your fork
git clone git@github.com:yourusername/repo.git

# Add upstream remote
cd repo
git remote add upstream git@github.com:originalowner/repo.git

# Verify remotes
git remote -v
```

**Syncing Fork:**
```bash
# Fetch upstream changes
git fetch upstream

# Merge upstream into your branch
git checkout main
git merge upstream/main

# Push to your fork
git push origin main
```

### 6.8 GitHub Actions (Brief)

**What is GitHub Actions?**
CI/CD platform integrated into GitHub.

**Simple Example (.github/workflows/ci.yml):**
```yaml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    - name: Setup Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18'
    - name: Install dependencies
      run: npm ci
    - name: Run tests
      run: npm test
```

---

## 7. Branching Strategies

### 7.1 What are Branches?

**Definition:**
A branch is an independent line of development that diverges from the main codebase.

**Why Branches?**
- Work on features independently
- Isolate experiments
- Maintain multiple versions
- Enable collaboration
- Safe to try changes

**Branch Concepts:**
- **HEAD** - Pointer to current branch
- **main/master** - Default production branch
- **feature branches** - New features
- **bugfix branches** - Bug fixes
- **release branches** - Prepare releases

### 7.2 Basic Branch Operations

**Creating Branches:**
```bash
# Create new branch
git branch feature-login

# Create and switch to branch
git checkout -b feature-login
git switch -c feature-login      # Modern syntax

# List branches
git branch                        # Local branches
git branch -r                     # Remote branches
git branch -a                     # All branches

# Switch branches
git checkout main
git switch main                   # Modern syntax
```

**Deleting Branches:**
```bash
# Delete local branch (safe)
git branch -d feature-login

# Force delete (even if not merged)
git branch -D feature-login

# Delete remote branch
git push origin --delete feature-login
git push origin :feature-login    # Alternative syntax
```

**Renaming Branches:**
```bash
# Rename current branch
git branch -m new-name

# Rename specific branch
git branch -m old-name new-name
```

### 7.3 Merging Branches

**Fast-Forward Merge:**
When target branch hasn't diverged from source.
```bash
git checkout main
git merge feature-login
```

**Three-Way Merge:**
When branches have diverged.
```bash
git checkout main
git merge feature-login
# Creates merge commit
```

**Merge Options:**
```bash
# No fast-forward (always create merge commit)
git merge --no-ff feature-login

# Squash (combine all commits into one)
git merge --squash feature-login

# Abort merge
git merge --abort
```

### 7.4 Handling Merge Conflicts

**What is a Conflict?**
When Git can't automatically merge changes.

**Conflict Markers:**
```
<<<<<<< HEAD
Your changes
=======
Their changes
>>>>>>> feature-branch
```

**Resolving Conflicts:**
```bash
# 1. Attempt merge
git merge feature-branch

# 2. Git shows conflicts
# CONFLICT (content): Merge conflict in file.txt

# 3. Open file and resolve
# Edit file, choose which changes to keep

# 4. Mark as resolved
git add file.txt

# 5. Complete merge
git commit
```

**Using Merge Tools:**
```bash
# Configure merge tool
git config --global merge.tool vimdiff

# Use merge tool
git mergetool
```

**Best Practices:**
- Communicate with team
- Update frequently from main
- Keep branches short-lived
- Test after resolving conflicts
- Use descriptive commit messages

### 7.5 Git Rebase

**What is Rebase?**
Moves branch to new base commit, rewriting history.

**Basic Rebase:**
```bash
# Update feature branch with main
git checkout feature-branch
git rebase main
```

**Interactive Rebase:**
```bash
# Rebase last 5 commits
git rebase -i HEAD~5

# Options:
# pick - use commit
# reword - change commit message
# edit - modify commit
# squash - combine with previous
# fixup - like squash but discard message
# drop - remove commit
```

**Rebase vs Merge:**

**Merge:**
- Preserves history exactly
- Shows when branches merged
- Non-destructive
- Can clutter history

**Rebase:**
- Creates linear history
- Cleaner history
- Rewrites commits (new SHAs)
- Can be confusing

**Golden Rule of Rebase:**
Never rebase commits that exist outside your repository (already pushed and shared).

### 7.6 Popular Branching Strategies

**1. Git Flow**

**Branches:**
- `main` - Production-ready code
- `develop` - Integration branch
- `feature/*` - New features
- `release/*` - Release preparation
- `hotfix/*` - Emergency fixes

**Workflow:**
```bash
# Start feature
git checkout -b feature/user-auth develop

# Finish feature
git checkout develop
git merge --no-ff feature/user-auth
git branch -d feature/user-auth

# Start release
git checkout -b release/1.0.0 develop

# Finish release
git checkout main
git merge --no-ff release/1.0.0
git tag -a v1.0.0
git checkout develop
git merge --no-ff release/1.0.0
git branch -d release/1.0.0

# Hotfix
git checkout -b hotfix/security-patch main
# Fix bug
git checkout main
git merge --no-ff hotfix/security-patch
git tag -a v1.0.1
git checkout develop
git merge --no-ff hotfix/security-patch
```

**When to Use:**
- Traditional release cycles
- Multiple production versions
- Large teams
- Complex projects

**2. GitHub Flow**

**Branches:**
- `main` - Always deployable
- `feature/*` - All work branches

**Workflow:**
```bash
# 1. Create branch from main
git checkout -b feature/add-comments main

# 2. Make commits
git commit -m "Add comment model"
git commit -m "Add comment API"

# 3. Push and create PR
git push -u origin feature/add-comments

# 4. Discuss and review in PR

# 5. Deploy from branch for testing

# 6. Merge to main
# Done via PR on GitHub

# 7. Deploy main
```

**When to Use:**
- Continuous deployment
- Small to medium teams
- Web applications
- Simpler workflow needed

**3. Trunk-Based Development**

**Concept:**
Everyone commits to main (trunk) frequently.

**Rules:**
- Small, frequent commits to main
- Feature flags for incomplete features
- No long-lived branches
- Automated testing required
- Fast feedback loops

**Workflow:**
```bash
# Pull latest
git pull origin main

# Make small change
# Commit directly to main
git commit -m "Update button color"
git push origin main

# For larger features, use short-lived branches
git checkout -b feature-x
# Work for few hours
git checkout main
git merge feature-x
git push origin main
git branch -d feature-x
```

**When to Use:**
- Mature CI/CD pipeline
- Strong test automation
- Small, experienced teams
- High deployment frequency

**4. GitLab Flow**

**Combines:**
- GitHub Flow simplicity
- Git Flow environment branches

**Branches:**
- `main` - Development
- `pre-production` - Staging
- `production` - Production

**Workflow:**
```bash
# Feature development
git checkout -b feature/new-api main
# Develop
git push origin feature/new-api
# Create MR (Merge Request)

# After merge to main
git checkout pre-production
git merge main
git push origin pre-production

# After testing in staging
git checkout production
git merge pre-production
git push origin production
```

### 7.7 Branch Naming Conventions

**Common Patterns:**
```
feature/description
bugfix/description
hotfix/description
release/version
experimental/description
```

**Examples:**
```
feature/user-authentication
feature/shopping-cart
bugfix/login-error
bugfix/api-timeout
hotfix/security-vulnerability
release/v1.2.0
experimental/new-ui-design
```

**Best Practices:**
- Use lowercase
- Use hyphens, not spaces
- Be descriptive but concise
- Include issue number if applicable
  - `feature/123-user-auth`
- Use consistent prefixes

### 7.8 Remote Branch Operations

**Tracking Remote Branches:**
```bash
# List remote branches
git branch -r

# Create local branch tracking remote
git checkout -b local-name origin/remote-name

# Or simpler (if names match)
git checkout remote-name

# Set upstream for existing branch
git branch --set-upstream-to=origin/main
git push -u origin branch-name
```

**Fetching Changes:**
```bash
# Fetch all remotes
git fetch --all

# Fetch specific remote
git fetch origin

# Fetch and prune deleted branches
git fetch -p
```

**Pushing Branches:**
```bash
# Push new branch to remote
git push -u origin feature-branch

# Push all branches
git push --all origin

# Push branch with different remote name
git push origin local-branch:remote-branch
```

---

## 8. Professional Commit Messages

### 8.1 Why Good Commit Messages Matter

**Benefits:**
- Understand changes quickly
- Navigate history effectively
- Generate changelogs automatically
- Review code efficiently
- Debug issues faster
- Document decisions

**Poor vs Good Examples:**
```
❌ BAD:
- "fix bug"
- "update"
- "changes"
- "asdfasdf"
- "Fixed the thing"

✅ GOOD:
- "Fix null pointer exception in user login"
- "Add validation for email input field"
- "Refactor database connection pooling"
- "Update dependencies to patch security vulnerabilities"
```

### 8.2 Commit Message Structure

**Format:**
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Example:**
```
feat(auth): Add two-factor authentication

Implement 2FA using TOTP (Time-based One-Time Password).
Users can enable 2FA in security settings. Adds required
dependencies: speakeasy for token generation and qrcode
for QR code generation.

Closes #247
```

### 8.3 Conventional Commits

**Commit Types:**
- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation changes
- `style` - Code style/formatting (no logic change)
- `refactor` - Code restructuring (no functionality change)
- `perf` - Performance improvements
- `test` - Adding or updating tests
- `build` - Build system changes
- `ci` - CI/CD configuration
- `chore` - Maintenance tasks
- `revert` - Revert previous commit

**Scope (optional):**
Indicates part of codebase affected.
```
feat(api): Add user endpoint
fix(ui): Correct button alignment
docs(readme): Update installation steps
```

**Subject Line Rules:**
1. Use imperative mood ("Add" not "Added" or "Adds")
2. No period at the end
3. Capitalize first letter
4. Keep under 50 characters
5. Be specific and concise

**Examples:**
```bash
# Features
git commit -m "feat: Add user registration endpoint"
git commit -m "feat(api): Implement pagination for products"

# Bug Fixes
git commit -m "fix: Resolve memory leak in image processing"
git commit -m "fix(auth): Correct token expiration logic"

# Documentation
git commit -m "docs: Update API documentation for v2"
git commit -m "docs(readme): Add contribution guidelines"

# Refactoring
git commit -m "refactor: Simplify user validation logic"
git commit -m "refactor(database): Extract query builders"

# Tests
git commit -m "test: Add unit tests for payment service"
git commit -m "test(api): Increase coverage for auth endpoints"

# Performance
git commit -m "perf: Optimize database queries"
git commit -m "perf(images): Implement lazy loading"
```

### 8.4 Detailed Commit Messages

**When to Add Body:**
- Complex changes
- Multiple files affected
- Need context or reasoning
- Breaking changes
- Related to specific issue

**Body Guidelines:**
- Wrap at 72 characters
- Separate from subject with blank line
- Explain "what" and "why", not "how"
- Use bullet points for multiple items

**Example with Body:**
```
refactor(auth): Migrate from JWT to session-based authentication

The previous JWT implementation had several security concerns:
- Tokens couldn't be invalidated before expiration
- No way to track active sessions
- Difficult to implement logout functionality

This change:
- Implements Redis-backed sessions
- Adds session management interface
- Provides logout endpoint
- Improves security posture

Breaking Change: Clients must now maintain cookies
and handle CSRF tokens.

Refs: #312, #401
```

### 8.5 Breaking Changes

**Format:**
```
feat(api)!: Change response format

BREAKING CHANGE: API responses now use camelCase instead
of snake_case for property names. Clients must update their
response parsing logic.

Before: { user_name: "John" }
After: { userName: "John" }
```

**Alternative:**
```
feat(api): Change response format

BREAKING CHANGE: Response format updated to camelCase
```

### 8.6 Linking to Issues

**Footer References:**
```
Fixes #123
Closes #456
Resolves #789
Refs #234
Related to #567
```

**Multiple Issues:**
```
Fixes #123, #456
Closes #789, #890
```

**Full Example:**
```
fix(api): Resolve race condition in concurrent requests

Add mutex lock to prevent simultaneous access to shared
resource. This was causing intermittent 500 errors under
high load.

Fixes #445
Related to #423
```

### 8.7 Co-authored Commits

**Format:**
```
feat: Implement OAuth integration

Co-authored-by: John Doe <john@example.com>
Co-authored-by: Jane Smith <jane@example.com>
```

**Creating Co-authored Commits:**
```bash
git commit -m "feat: Implement OAuth" \
  -m "Co-authored-by: John Doe <john@example.com>"
```

### 8.8 Commit Message Templates

**Create Template File (~/.gitmessage):**
```
# <type>(<scope>): <subject>

# <body>

# <footer>

# Type: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert
# Scope: Optional, component affected
# Subject: Imperative, no period, max 50 chars
# Body: Explain what and why (wrap at 72 chars)
# Footer: Issue references, breaking changes
```

**Configure Git:**
```bash
git config --global commit.template ~/.gitmessage
```

### 8.9 Atomic Commits

**Concept:**
Each commit should be a single, logical change.

**Benefits:**
- Easier to review
- Easier to revert
- Clearer history
- Better for bisecting bugs

**Example - Bad:**
```
❌ ONE COMMIT: "Implement user feature"
- Add user model
- Add user controller
- Add user views
- Add user tests
- Fix unrelated bug in auth
- Update documentation
```

**Example - Good:**
```
✅ MULTIPLE ATOMIC COMMITS:
1. "feat(models): Add User model with validation"
2. "feat(api): Add User CRUD endpoints"
3. "feat(ui): Create user management interface"
4. "test(user): Add unit and integration tests"
5. "fix(auth): Resolve session timeout issue"
6. "docs: Update API documentation for user endpoints"
```

### 8.10 Commit Message Anti-patterns

**Avoid:**
```
❌ "WIP" - Work in progress
❌ "Fix stuff"
❌ "Updates"
❌ "Fixed typo" (if frequent, enable spell-checker)
❌ "asdf" or "testing"
❌ "Quick fix"
❌ Past tense: "Fixed bug" (use "Fix bug")
❌ Multiple unrelated changes in one commit
```

**Tools to Help:**
- commitlint - Enforce commit conventions
- husky - Git hooks for validation
- commitizen - Interactive commit message tool

---

## 9. Writing Effective Pull Requests

### 9.1 What is a Pull Request (PR)?

**Definition:**
A request to merge your changes from one branch into another, typically from your feature branch into the main branch.

**Purpose:**
- Code review
- Discussion platform
- Quality gate
- Knowledge sharing
- Documentation of changes

**PR Lifecycle:**
1. Create branch and make changes
2. Push branch to remote
3. Open pull request
4. Code review and discussion
5. Address feedback
6. Approval
7. Merge
8. Clean up (delete branch)

### 9.2 Before Creating a PR

**Checklist:**
- [ ] Code compiles/runs without errors
- [ ] All tests pass
- [ ] New tests added for new features
- [ ] Code follows project style guide
- [ ] No debug code or console logs
- [ ] Documentation updated
- [ ] Self-review completed
- [ ] Commits are clean and atomic
- [ ] Branch is up-to-date with base branch

**Update Your Branch:**
```bash
# Fetch latest changes
git fetch origin

# Rebase or merge main into your branch
git rebase origin/main
# OR
git merge origin/main

# Resolve conflicts if any
# Run tests
# Push updated branch
git push origin feature-branch --force-with-lease
```

### 9.3 Creating a Pull Request

**Via GitHub Interface:**
1. Navigate to repository
2. Click "Pull requests" tab
3. Click "New pull request"
4. Select base and compare branches
5. Click "Create pull request"
6. Fill in title and description
7. Add reviewers, assignees, labels
8. Click "Create pull request"

**Via GitHub CLI:**
```bash
# Create PR from current branch
gh pr create --title "Add user authentication" --body "Description here"

# Create PR with template
gh pr create --fill

# Create draft PR
gh pr create --draft
```

**Via Command Line (with hub):**
```bash
# Push branch
git push -u origin feature-branch

# Create PR
hub pull-request -m "Title" -m "Description"
```

### 9.4 PR Title Best Practices

**Format:**
```
[Type] Brief description of changes
```

**Examples:**
```
✅ GOOD:
- [Feature] Add two-factor authentication
- [Fix] Resolve memory leak in image processor
- [Refactor] Simplify database query logic
- [Docs] Update API documentation for v2
- [Performance] Optimize product search algorithm

❌ BAD:
- "Update files"
- "Fix"
- "Changes"
- "PR #123"
- "Work from yesterday"
```

**Tips:**
- Be specific and concise
- Use imperative mood
- Max 60 characters
- Include ticket number if applicable: "[Feature] Add 2FA (#234)"

### 9.5 PR Description Template

**Comprehensive Template:**
```markdown
## Description
Brief overview of what this PR does and why.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Refactoring (no functional changes)
- [ ] Performance improvement

## Related Issues
Closes #123
Relates to #456

## Changes Made
- Added user authentication flow
- Implemented JWT token generation
- Created login and signup endpoints
- Added password hashing with bcrypt

## Screenshots (if applicable)
Before:
[Screenshot]

After:
[Screenshot]

## Testing
### How to Test
1. Run `npm test`
2. Start server with `npm start`
3. Navigate to `/login`
4. Enter credentials
5. Verify successful login

### Test Coverage
- Unit tests added for auth service
- Integration tests for login endpoint
- E2E tests for complete flow
- Current coverage: 85%

## Checklist
- [ ] My code follows the project's style guidelines
- [ ] I have performed a self-review
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have updated the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally
- [ ] Any dependent changes have been merged and published

## Additional Notes
- Requires database migration (see migrations/001_add_users.sql)
- Environment variable AUTH_SECRET must be set
- Breaking change: Old API tokens will be invalid

## Reviewer Notes
Please pay special attention to:
- Security of password handling
- Token expiration logic
- Error handling in auth middleware
```

### 9.6 PR Size Best Practices

**Ideal PR Size:**
- **Small:** < 200 lines changed (ideal)
- **Medium:** 200-500 lines (acceptable)
- **Large:** 500-1000 lines (needs justification)
- **Huge:** > 1000 lines (should be broken down)

**Why Small PRs are Better:**
- Faster to review
- Easier to understand
- Less likely to introduce bugs
- Quicker feedback cycle
- Simpler to revert if needed

**Breaking Down Large PRs:**
```
❌ ONE LARGE PR: "Implement user management system" (2000 lines)

✅ MULTIPLE SMALL PRs:
1. "Add User model and database schema" (150 lines)
2. "Implement user CRUD API endpoints" (200 lines)
3. "Add user authentication middleware" (100 lines)
4. "Create user management UI components" (250 lines)
5. "Add user management tests" (180 lines)
```

### 9.7 Code Review Process

**As Author:**
1. **Respond promptly** to comments
2. **Be receptive** to feedback
3. **Explain your reasoning** when disagreeing
4. **Thank reviewers** for their time
5. **Mark conversations** as resolved appropriately
6. **Update PR** based on feedback

**Responding to Comments:**
```markdown
# Accepting feedback
✅ "Good catch! Fixed in commit abc123."
✅ "You're right, I'll refactor this approach."

# Explaining reasoning
✅ "I chose this approach because X. However, I'm open to Y if you think it's better."
✅ "This handles the edge case where Z, which wasn't covered before."

# Requesting clarification
✅ "Could you elaborate on what you mean by X?"
✅ "Do you have a suggestion for how to improve this?"
```

**As Reviewer:**
1. **Be respectful** and constructive
2. **Be specific** in comments
3. **Explain the "why"** not just "what"
4. **Approve** if changes are good
5. **Request changes** if issues exist
6. **Use suggestions** for small fixes

**Comment Types:**
```markdown
# Blocking issue
❗️ This will cause a security vulnerability. Must fix before merge.

# Suggestion
💡 Consider using Array.map() here for cleaner code.

# Question
❓ Why did you choose this approach over X?

# Nitpick (non-blocking)
🔍 Minor: Variable name could be more descriptive.

# Praise
✅ Great solution! Very clean implementation.
```

### 9.8 Draft Pull Requests

**When to Use:**
- Work in progress
- Want early feedback
- Exploring approach
- Pausing work temporarily

**Creating Draft PR:**
```bash
# Via GitHub CLI
gh pr create --draft

# Or click "Create draft pull request" in UI
```

**Converting to Ready:**
- Click "Ready for review" button
- Ensure all checklist items completed

### 9.9 PR Labels and Metadata

**Common Labels:**
- `bug` - Bug fix
- `feature` - New feature
- `documentation` - Docs only
- `refactor` - Code restructuring
- `performance` - Performance improvement
- `dependencies` - Dependency updates
- `breaking-change` - Breaking changes
- `needs-review` - Awaiting review
- `work-in-progress` - Not ready yet
- `ready-to-merge` - Approved
- `do-not-merge` - Hold merge

**Assignees:**
- Assign yourself (you're responsible)
- Assign pair programmer if collaborative

**Reviewers:**
- Request review from relevant team members
- Usually 1-3 reviewers
- Code owners automatically requested

**Projects:**
- Link to project board if applicable
- Tracks progress visually

**Milestones:**
- Associate with release milestone
- Helps track release progress

### 9.10 PR Review Comments

**GitHub Features:**
```markdown
# Single comment
Just add a comment on specific line

# Suggestion (code change)
```suggestion
const result = items.map(item => item.value);
```

# Multi-line comment
Select multiple lines → Click "+" → Add comment

# Starting a review
Click "Start a review" → Add all comments → "Finish review"
```

**Review Status:**
- **Comment** - General feedback, no explicit approval
- **Approve** - Changes look good
- **Request changes** - Issues must be addressed

### 9.11 Merging Strategies

**1. Merge Commit**
```bash
git merge --no-ff feature-branch
```
- Preserves all commits
- Creates merge commit
- Full history visible
- Can clutter history

**2. Squash and Merge**
```bash
git merge --squash feature-branch
```
- Combines all commits into one
- Clean history
- Loses individual commit messages
- Good for feature branches

**3. Rebase and Merge**
```bash
git rebase main
git merge --ff-only feature-branch
```
- Linear history
- Preserves commits
- No merge commit
- Cleaner than merge commit

**Choosing Strategy:**
- **Squash** - Small features, many WIP commits
- **Merge commit** - Want full history, multiple contributors
- **Rebase** - Clean history preferred, disciplined commits

### 9.12 After Merge

**Clean Up:**
```bash
# Delete local branch
git branch -d feature-branch

# Delete remote branch
git push origin --delete feature-branch

# Update main
git checkout main
git pull origin main

# GitHub can auto-delete branches after merge
```

**Post-Merge Activities:**
- Verify deployment/CI passes
- Close related issues
- Update project board
- Notify team if needed
- Monitor for issues

---

## 10. Advanced Git Topics

### 10.1 Git Stash

**Purpose:**
Temporarily store uncommitted changes without committing.

**Basic Usage:**
```bash
# Stash current changes
git stash

# Stash with message
git stash save "Work on feature X"

# List stashes
git stash list

# Apply most recent stash
git stash apply

# Apply and remove from stash
git stash pop

# Apply specific stash
git stash apply stash@{2}

# Show stash contents
git stash show
git stash show -p stash@{0}

# Drop specific stash
git stash drop stash@{1}

# Clear all stashes
git stash clear
```

**Advanced Stashing:**
```bash
# Stash including untracked files
git stash -u

# Stash including ignored files
git stash -a

# Stash specific files
git stash push -m "Message" file1.txt file2.txt

# Create branch from stash
git stash branch feature-branch stash@{0}

# Stash interactively
git stash -p
```

### 10.2 Git Tags

**Purpose:**
Mark specific points in history (usually releases).

**Types:**
- **Lightweight** - Just a pointer to commit
- **Annotated** - Full Git object with metadata (recommended)

**Creating Tags:**
```bash
# Lightweight tag
git tag v1.0.0

# Annotated tag
git tag -a v1.0.0 -m "Release version 1.0.0"

# Tag specific commit
git tag -a v1.0.0 abc123 -m "Release 1.0.0"

# List tags
git tag
git tag -l "v1.*"

# Show tag information
git show v1.0.0
```

**Pushing Tags:**
```bash
# Push specific tag
git push origin v1.0.0

# Push all tags
git push --tags

# Push annotated tags only
git push --follow-tags
```

**Deleting Tags:**
```bash
# Delete local tag
git tag -d v1.0.0

# Delete remote tag
git push origin --delete v1.0.0
git push origin :refs/tags/v1.0.0
```

**Checking Out Tags:**
```bash
# View tag (detached HEAD)
git checkout v1.0.0

# Create branch from tag
git checkout -b hotfix/v1.0.1 v1.0.0
```

### 10.3 Git Cherry-Pick

**Purpose:**
Apply specific commits from one branch to another.

**Basic Usage:**
```bash
# Cherry-pick single commit
git cherry-pick abc123

# Cherry-pick multiple commits
git cherry-pick abc123 def456

# Cherry-pick range
git cherry-pick abc123..def456

# Cherry-pick without committing
git cherry-pick -n abc123
```

**Handling Conflicts:**
```bash
# If conflicts occur
# 1. Resolve conflicts
# 2. Stage resolved files
git add .
# 3. Continue cherry-pick
git cherry-pick --continue

# Abort cherry-pick
git cherry-pick --abort
```

**Use Cases:**
- Apply hotfix to multiple branches
- Pick specific features
- Recover lost commits

### 10.4 Git Bisect

**Purpose:**
Binary search to find commit that introduced a bug.

**Usage:**
```bash
# Start bisecting
git bisect start

# Mark current commit as bad
git bisect bad

# Mark known good commit
git bisect good abc123

# Git checks out middle commit
# Test it, then mark as good or bad
git bisect good
# OR
git bisect bad

# Continue until bug is found
# Git will identify the culprit commit

# End bisect session
git bisect reset
```

**Automated Bisect:**
```bash
# With test script
git bisect start
git bisect bad
git bisect good abc123
git bisect run npm test

# Git automatically tests each commit
```

### 10.5 Git Reflog

**Purpose:**
Reference log of HEAD movements - safety net for "lost" commits.

**Usage:**
```bash
# Show reflog
git reflog

# Show detailed reflog
git reflog show HEAD

# Recover "lost" commit
git reflog
# Find commit SHA before mistake
git reset --hard abc123

# Or create branch from lost commit
git branch recovered-work abc123
```

**Use Cases:**
- Recover after bad reset
- Find deleted branches
- Undo rebase
- Recovery after force push

### 10.6 Git Hooks

**Purpose:**
Scripts that run automatically at certain Git events.

**Hook Types:**
- `pre-commit` - Before commit is created
- `prepare-commit-msg` - Before commit message editor
- `commit-msg` - Validate commit message
- `post-commit` - After commit is created
- `pre-push` - Before push
- `post-merge` - After merge

**Location:**
`.git/hooks/`

**Example: pre-commit Hook**
```bash
#!/bin/sh
# .git/hooks/pre-commit

# Run linter
npm run lint

# Check result
if [ $? -ne 0 ]; then
    echo "Linting failed. Commit aborted."
    exit 1
fi

# Run tests
npm test

if [ $? -ne 0 ]; then
    echo "Tests failed. Commit aborted."
    exit 1
fi
```

**Make Executable:**
```bash
chmod +x .git/hooks/pre-commit
```

**Using Husky (Recommended):**
```bash
# Install husky
npm install --save-dev husky

# Initialize
npx husky install

# Add pre-commit hook
npx husky add .git/hooks/pre-commit "npm test"
```

### 10.7 Git Submodules

**Purpose:**
Include other Git repositories within your repository.

**Adding Submodule:**
```bash
# Add submodule
git submodule add https://github.com/user/repo.git path/to/submodule

# Commit the submodule
git commit -m "Add submodule"
```

**Cloning Repository with Submodules:**
```bash
# Clone with submodules
git clone --recurse-submodules https://github.com/user/repo.git

# Or after cloning
git submodule init
git submodule update
```

**Updating Submodules:**
```bash
# Update to latest
cd path/to/submodule
git pull origin main

# Or from parent
git submodule update --remote

# Commit updated submodule reference
git add path/to/submodule
git commit -m "Update submodule"
```

**Alternative: Git Subtree**
```bash
# Add subtree
git subtree add --prefix=path/to/subtree https://github.com/user/repo.git main --squash

# Update subtree
git subtree pull --prefix=path/to/subtree https://github.com/user/repo.git main --squash
```

### 10.8 Git Worktrees

**Purpose:**
Work on multiple branches simultaneously.

**Usage:**
```bash
# Create worktree
git worktree add ../feature-x feature-branch

# List worktrees
git worktree list

# Remove worktree
git worktree remove ../feature-x

# Prune worktrees
git worktree prune
```

**Use Cases:**
- Review PR while working on feature
- Quick hotfix without stashing
- Compare implementations

### 10.9 Git Aliases

**Creating Aliases:**
```bash
# Status shortcut
git config --global alias.st status

# Commit shortcut
git config --global alias.cm commit

# Pretty log
git config --global alias.lg "log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"

# Undo last commit
git config --global alias.undo "reset HEAD~1 --mixed"

# List branches
git config --global alias.br branch

# Amend commit
git config --global alias.amend "commit --amend --no-edit"
```

**Complex Aliases:**
```bash
# Fancy log
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr)%C(bold blue)<%an>%Creset' --abbrev-commit"

# Show branches with latest commit
git config --global alias.br "for-each-ref --sort=-committerdate refs/heads/ --format='%(HEAD) %(color:yellow)%(refname:short)%(color:reset) - %(color:red)%(objectname:short)%(color:reset) - %(contents:subject) - %(authorname) (%(color:green)%(committerdate:relative)%(color:reset))'"
```

### 10.10 Git LFS (Large File Storage)

**Purpose:**
Handle large files (videos, datasets, binaries) efficiently.

**Installation:**
```bash
# Install Git LFS
# macOS
brew install git-lfs

# Linux
sudo apt-get install git-lfs

# Windows - included in Git for Windows

# Initialize
git lfs install
```

**Usage:**
```bash
# Track file types
git lfs track "*.psd"
git lfs track "*.mp4"
git lfs track "datasets/*"

# Track specific file
git lfs track "large-file.bin"

# Commit .gitattributes
git add .gitattributes
git commit -m "Configure Git LFS"

# Add and commit large files normally
git add large-file.mp4
git commit -m "Add video"
git push
```

**Viewing Tracked Files:**
```bash
# List tracked patterns
git lfs track

# List tracked files
git lfs ls-files
```

---

## 11. Best Practices & Real-world Workflows

### 11.1 Daily Workflow

**Morning Routine:**
```bash
# 1. Update main branch
git checkout main
git pull origin main

# 2. Create/switch to feature branch
git checkout -b feature/new-feature
# OR
git checkout feature/existing-feature

# 3. Update feature branch with main
git merge main
# OR
git rebase main
```

**During Development:**
```bash
# 1. Make changes
# 2. Check status frequently
git status

# 3. Review changes
git diff

# 4. Stage changes
git add -p  # Interactive staging

# 5. Commit atomically
git commit -m "type: descriptive message"

# 6. Push regularly
git push origin feature/new-feature
```

**End of Day:**
```bash
# 1. Ensure all work is committed
git status

# 2. Push to remote (backup)
git push origin feature/new-feature

# 3. If work is incomplete, consider
git stash  # For very temporary work
# OR
git commit -m "WIP: partial implementation"
```

### 11.2 Collaborative Workflow

**Starting Collaboration:**
```bash
# 1. Fork repository (on GitHub)

# 2. Clone your fork
git clone git@github.com:yourname/repo.git

# 3. Add upstream remote
git remote add upstream git@github.com:original/repo.git

# 4. Verify remotes
git remote -v
```

**Regular Sync:**
```bash
# 1. Fetch upstream changes
git fetch upstream

# 2. Update local main
git checkout main
git merge upstream/main

# 3. Push to your fork
git push origin main

# 4. Update feature branch
git checkout feature-branch
git rebase main
```

**Contributing:**
```bash
# 1. Create feature branch from main
git checkout -b feature/contribution main

# 2. Make changes and commit
git add .
git commit -m "feat: add new feature"

# 3. Push to your fork
git push origin feature/contribution

# 4. Create Pull Request on GitHub

# 5. Address review feedback
# Make changes
git add .
git commit -m "fix: address review comments"
git push origin feature/contribution

# 6. After merge, clean up
git checkout main
git pull upstream main
git branch -d feature/contribution
git push origin --delete feature/contribution
```

### 11.3 Hotfix Workflow

**Emergency Production Fix:**
```bash
# 1. Create hotfix branch from production
git checkout -b hotfix/critical-bug production

# 2. Fix the bug
# Make minimal, focused changes

# 3. Commit
git commit -m "fix: resolve critical security issue"

# 4. Test thoroughly

# 5. Merge to production
git checkout production
git merge --no-ff hotfix/critical-bug
git tag -a v1.2.1 -m "Hotfix release 1.2.1"
git push origin production --tags

# 6. Deploy

# 7. Merge back to main/develop
git checkout main
git merge --no-ff hotfix/critical-bug
git push origin main

# 8. Clean up
git branch -d hotfix/critical-bug
git push origin --delete hotfix/critical-bug
```

### 11.4 Release Workflow

**Preparing Release:**
```bash
# 1. Create release branch
git checkout -b release/v1.2.0 develop

# 2. Update version numbers
# Edit package.json, version files, etc.
git commit -am "chore: bump version to 1.2.0"

# 3. Run final tests
npm test

# 4. Update changelog
# Edit CHANGELOG.md
git commit -am "docs: update changelog for v1.2.0"

# 5. Merge to main
git checkout main
git merge --no-ff release/v1.2.0

# 6. Tag release
git tag -a v1.2.0 -m "Release version 1.2.0"

# 7. Push
git push origin main --tags

# 8. Merge back to develop
git checkout develop
git merge --no-ff release/v1.2.0
git push origin develop

# 9. Clean up
git branch -d release/v1.2.0
```

### 11.5 Code Review Best Practices

**Before Requesting Review:**
- [ ] Self-review completed
- [ ] All tests passing
- [ ] Documentation updated
- [ ] No debug code remaining
- [ ] Commits are clean
- [ ] PR description is clear

**As Reviewer:**
1. Understand the context
2. Check functionality
3. Review code quality
4. Verify tests
5. Check for security issues
6. Suggest improvements
7. Approve or request changes

**Review Checklist:**
```markdown
### Functionality
- [ ] Changes match requirements
- [ ] Edge cases handled
- [ ] Error handling present

### Code Quality
- [ ] Follows style guide
- [ ] No code duplication
- [ ] Functions are focused
- [ ] Variable names are clear

### Testing
- [ ] Tests included
- [ ] Tests pass
- [ ] Coverage adequate

### Security
- [ ] No sensitive data exposed
- [ ] Input validated
- [ ] Authentication/authorization correct

### Documentation
- [ ] Comments where needed
- [ ] README updated
- [ ] API docs updated
```

### 11.6 Handling Large Repositories

**Shallow Clone:**
```bash
# Clone with limited history
git clone --depth 1 https://github.com/user/repo.git

# Fetch more history later
git fetch --depth=100
```

**Sparse Checkout:**
```bash
# Clone without files
git clone --no-checkout https://github.com/user/repo.git
cd repo

# Enable sparse checkout
git sparse-checkout init --cone

# Specify directories to include
git sparse-checkout set src/components src/utils

# Checkout
git checkout main
```

**Partial Clone:**
```bash
# Clone without downloading all blobs
git clone --filter=blob:none https://github.com/user/repo.git
```

### 11.7 Git Performance Tips

**Optimize Repository:**
```bash
# Garbage collection
git gc

# Aggressive optimization
git gc --aggressive

# Prune old data
git prune

# Verify integrity
git fsck
```

**Optimize Fetch:**
```bash
# Fetch only specific branch
git fetch origin main

# Prune deleted branches
git fetch -p

# Fetch tags separately
git fetch --no-tags
```

### 11.8 Security Best Practices

**Protecting Credentials:**
```bash
# Use SSH keys instead of passwords

# Credential helper
git config --global credential.helper cache

# Never commit secrets
# Use environment variables
# Use .env files (added to .gitignore)
```

**Removing Sensitive Data:**
```bash
# Remove file from history
git filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch secrets.txt" \
  --prune-empty --tag-name-filter cat -- --all

# Modern alternative: git-filter-repo
pip install git-filter-repo
git filter-repo --path secrets.txt --invert-paths

# Force push
git push origin --force --all
```

**Signing Commits:**
```bash
# Generate GPG key
gpg --gen-key

# List keys
gpg --list-secret-keys --keyid-format LONG

# Configure Git
git config --global user.signingkey YOUR_KEY_ID

# Sign commits
git commit -S -m "Signed commit"

# Always sign
git config --global commit.gpgsign true
```

### 11.9 Troubleshooting Common Issues

**Resolving Detached HEAD:**
```bash
# Create branch from current state
git checkout -b temp-branch

# Or return to branch
git checkout main
```

**Undoing Mistakes:**
```bash
# Undo last commit (keep changes)
git reset --soft HEAD~1

# Undo last commit (discard changes)
git reset --hard HEAD~1

# Undo pushed commit
git revert HEAD
git push origin main

# Recover deleted branch
git reflog
git checkout -b recovered-branch abc123
```

**Fixing Merge Conflicts:**
```bash
# View conflicted files
git status

# Open and edit conflicted files
# Remove conflict markers

# Mark as resolved
git add conflicted-file.txt

# Complete merge
git commit

# If too complex, abort
git merge --abort
```

**Cleaning Up:**
```bash
# Remove untracked files (dry run)
git clean -n

# Remove untracked files
git clean -f

# Remove untracked directories
git clean -fd

# Remove ignored files too
git clean -fx
```

### 11.10 Git Configuration Best Practices

**Recommended Global Config:**
```bash
# User information
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# Default branch
git config --global init.defaultBranch main

# Editor
git config --global core.editor "code --wait"

# Merge tool
git config --global merge.tool vscode

# Color
git config --global color.ui auto

# Aliases
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status

# Pull strategy
git config --global pull.rebase false

# Push strategy
git config --global push.default simple

# Credential caching
git config --global credential.helper cache
```

**Project-specific Config:**
```bash
# In repository directory
git config user.email "work@company.com"
git config core.autocrlf input
```

---

## Practice Exercises

### Exercise 1: Basic Git Workflow
1. Create new directory and initialize Git
2. Create README.md file
3. Stage and commit the file
4. Make changes and commit again
5. View commit history

### Exercise 2: Branching
1. Create new branch for feature
2. Make changes and commit
3. Switch back to main
4. Merge feature branch
5. Delete feature branch

### Exercise 3: Collaboration Simulation
1. Fork a repository
2. Clone your fork
3. Create feature branch
4. Make changes and push
5. Create pull request

### Exercise 4: Conflict Resolution
1. Create two branches
2. Modify same line in both
3. Merge one branch
4. Attempt to merge second (conflict)
5. Resolve conflict manually

### Exercise 5: Advanced Operations
1. Use git stash to save work
2. Cherry-pick specific commit
3. Use interactive rebase
4. Create and push tags

---

## Additional Resources

### Documentation
- [Official Git Documentation](https://git-scm.com/doc)
- [GitHub Docs](https://docs.github.com)
- [Atlassian Git Tutorials](https://www.atlassian.com/git/tutorials)

### Interactive Learning
- [Learn Git Branching](https://learngitbranching.js.org/)
- [GitHub Learning Lab](https://lab.github.com/)
- [Git Immersion](http://gitimmersion.com/)

### Books
- *Pro Git* by Scott Chacon and Ben Straub (free online)
- *Git Pocket Guide* by Richard E. Silverman

### Videos
- [Git and GitHub for Beginners - freeCodeCamp](https://www.youtube.com/watch?v=RGOj5yH7evk)
- [Advanced Git - Frontend Masters](https://frontendmasters.com/courses/git-in-depth/)

### Tools
- [GitHub Desktop](https://desktop.github.com/)
- [GitKraken](https://www.gitkraken.com/)
- [SourceTree](https://www.sourcetreeapp.com/)
- [Git CLI - gh](https://cli.github.com/)

### Cheat Sheets
- [GitHub Git Cheat Sheet](https://education.github.com/git-cheat-sheet-education.pdf)
- [Atlassian Git Cheat Sheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)

---

## Quiz Questions

1. What is the difference between Git and GitHub?
2. Explain the three main areas in Git workflow.
3. What's the difference between `git fetch` and `git pull`?
4. When should you use rebase vs merge?
5. What is a detached HEAD state?
6. How do you undo the last commit?
7. What's the purpose of .gitignore?
8. Explain the difference between `git reset` and `git revert`.
9. What are the benefits of small, atomic commits?
10. How do you resolve merge conflicts?

---

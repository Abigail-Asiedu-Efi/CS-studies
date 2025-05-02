# Contributing Guide

Welcome! This guide outlines how we contribute to our **CS Learning Together** repository. Since we’re learning together, this process is designed to be beginner-friendly and collaborative.

---

## Purpose

We use this repository to:
- Share notes and code on CS topics
- Reinforce our learning by teaching each other
- Practice using Git, GitHub, and real-world workflows

---

## Tools We Use

- **Git & GitHub** for version control and collaboration
- **Markdown (`.md`)** for writing notes
- **Code Editors:** VS Code or any preferred IDE
- **Languages:** Python, JavaScript, Java, etc.

---

## Folder Structure

cs-learning-together/
├── programming/
├── computer-science/
├── cyber-security/
└── tools/

Place your content in the correct folder and subfolder. Each person can use their own file or folder inside a topic if needed.

---

## Our Workflow

1. **Pull the Latest Changes**  
    Always start by updating your local copy:

   ```bash
        git pull origin main
    ```

2. **Create or Edit Files**
    Add your notes, code, or questions in the appropriate folder. Use clear, descriptive filenames.

3. **Stage Your Changes**

    ```bash
        git add path/to/your-file
    ```

4. **Write a Clear Commit Message**

   ``` Use meaningful messages like:
    Add notes on recursion in algorithms
    Fix typo in sister-notes.md
    Add Python example for bubble sort```

5. **Commit and Push**

    ```bash
        git commit -m "Your message here"
        git push origin main
    ```

6. **Review Each Other’s Work**

- Leave comments in markdown files
- Use in-line comments in code for feedback
- Discuss unclear parts during weekly syncs

## Contribution Guidelines

- Keep your notes and code well-organized and readable
- Use headings, bullet points, and code blocks in Markdown
- If you use external content (books, websites), credit the source
- Use proper indentation and naming conventions in code
- Ask for help or clarification anytime—learning is the goal!

## Create a Branch

Once comfortable with Git:

```bash
    git checkout -b your-topic-name
    # make changes
    git commit -m "Work on topic"
    git push origin your-topic-name
```

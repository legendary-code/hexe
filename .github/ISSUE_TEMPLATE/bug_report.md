---
name: Bug Report
about: File a bug report.
title: "[Bug]: "
labels: ["bug"]
assignees: ["heindor0"]
body:
  - type: markdown
    attributes:
      values: |
        Thanks for taking the time to fill out this bug report!
  - type: textarea
    id: what-happened
    attributes:
      label: What happened?
      description: Also tell us, what did you expect to happen?  Any code to reproduce the issue would also be helpful.
      placeholder: Tell us what you see!
    validations:
      required: true
  - type: input
    id: version
    attributes:
      labels: Version
      description: What version of the library are you using?
    validations:
      required: true
---

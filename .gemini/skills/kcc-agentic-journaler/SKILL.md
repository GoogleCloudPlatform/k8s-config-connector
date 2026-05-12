---
name: kcc-agentic-journaler
description: Provides a structured logic for capturing and routing agentic learnings to prevent knowledge pollution.
---

# KCC Agentic Journaler

## Overview
As an agent, you must record "tribal knowledge" and technical breakthroughs. This skill ensures that your learnings are stored in the most contextually relevant location.

## Routing Logic

When you have a new learning, follow this hierarchy to choose the destination:

### Tier 1: General Mechanics
**Scenario**: You found a better way to do something that applies to *all* KCC resources (e.g., "A more efficient way to resolve Project IDs").
**Destination**: Update the relevant Skill's `SKILL.md` file directly.

### Tier 2: Domain Best Practices
**Scenario**: You found a trick for a specific implementation area (e.g., "How to handle OneOf fields in mappers").
**Destination**: Append to the relevant Skill's `journal.md` file (e.g., `.gemini/skills/kcc-mapper-implementer/journal.md`).

### Tier 3: Service Tribal Knowledge
**Scenario**: You found a quirk unique to a specific GCP service (e.g., "NetworkSecurity APIs require an explicit location in the body even if it is in the URL").
**Destination**: Create or append to `.gemini/journals/<service_name>.md`. 

## Journal Entry Template
Use this format for Tier 2 and Tier 3 entries:

```markdown
### [YYYY-MM-DD] <Brief Title>
- **Context**: <What were you implementing? Link to Kind/PR>
- **Problem**: <What was the unexpected behavior or hurdle?>
- **Solution**: <How did you solve it? Provide code snippets if relevant.>
- **Impact**: <Why should the next agent care about this?>
```

## Validation
In your final turn of the task, you MUST state which knowledge was captured and provide the file path.

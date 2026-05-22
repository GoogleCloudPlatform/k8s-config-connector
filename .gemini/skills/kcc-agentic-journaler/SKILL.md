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

### Tier 2: Domain and Service Tribal Knowledge
**Scenario**: You found a trick for a specific implementation area, or a quirk unique to a specific GCP service.
**Destination**: **ALWAYS** create or append to `.gemini/journals/<service_name>.md`. 

**CRITICAL**: Do NOT append to any `journal.md` inside a `.gemini/skills/` directory. Multiple agents working in parallel on different services will cause merge conflicts if they all write to a centralized skill journal. By keeping journals scoped to the `<service_name>`, we eliminate cross-service merge conflicts.

## Journal Entry Template
Use this format for your entries in `.gemini/journals/<service_name>.md`:

```markdown
### [YYYY-MM-DD] <Brief Title>
- **Context**: <What were you implementing? Link to Kind/PR>
- **Problem**: <What was the unexpected behavior or hurdle?>
- **Solution**: <How did you solve it? Provide code snippets if relevant.>
- **Impact**: <Why should the next agent care about this?>
```

## Validation
In your final turn of the task, you MUST state which knowledge was captured and provide the file path.
